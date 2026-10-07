package httpapi

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/armelo10/vtc_go/backend/internal/domain/auth"
	"github.com/armelo10/vtc_go/backend/internal/infrastructure/postgres"
	"golang.org/x/crypto/bcrypt"
)

type authHandler struct {
	repo *postgres.AuthRepository
}

type registerRequest struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func newAuthHandler(repo *postgres.AuthRepository) *authHandler {
	return &authHandler{repo: repo}
}

func (h *authHandler) register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.Password) < 8 || req.Email == "" || req.FirstName == "" || req.LastName == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid registration"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "password hashing failed"})
		return
	}
	user, err := h.repo.CreatePassenger(r.Context(), strings.TrimSpace(req.Email), strings.TrimSpace(req.Phone), string(hash), strings.TrimSpace(req.FirstName), strings.TrimSpace(req.LastName))
	if err != nil {
		if err == postgres.ErrEmailExists {
			writeJSON(w, http.StatusConflict, map[string]string{"error": "email already registered"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "registration failed"})
		return
	}
	writeJSON(w, http.StatusCreated, user)
}

func (h *authHandler) login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid login"})
		return
	}
	hash, user, err := h.repo.FindByEmail(r.Context(), strings.TrimSpace(req.Email))
	if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)) != nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
		return
	}
	token, err := newToken()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "session creation failed"})
		return
	}
	expires := time.Now().UTC().Add(30 * 24 * time.Hour)
	if err := h.repo.CreateSession(r.Context(), user.ID, hashToken(token), expires); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "session creation failed"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"token": token, "expires_at": expires, "user": user})
}

func (h *authHandler) me(w http.ResponseWriter, r *http.Request) {
	user, ok := userFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "authentication required"})
		return
	}
	writeJSON(w, http.StatusOK, user)
}

func newToken() (string, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

type contextKey string

const userContextKey contextKey = "authenticated-user"

func withUser(r *http.Request, user auth.User) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userContextKey, user))
}

func userFromContext(ctx context.Context) (auth.User, bool) {
	user, ok := ctx.Value(userContextKey).(auth.User)
	return user, ok
}
