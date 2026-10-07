package httpapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/armelo10/vtc_go/backend/internal/infrastructure/config"
	pricingapp "github.com/armelo10/vtc_go/backend/internal/application/pricing"
	"github.com/armelo10/vtc_go/backend/internal/domain/pricing"
	"github.com/armelo10/vtc_go/backend/internal/infrastructure/postgres"
	"github.com/armelo10/vtc_go/backend/internal/infrastructure/routing"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	cfg  config.Config
	pool *pgxpool.Pool
	mux  *http.ServeMux
}

func NewServer(cfg config.Config, pool *pgxpool.Pool) *Server {
	authRepo := postgres.NewAuthRepository(pool)
	auth := newAuthHandler(authRepo)
	bookingRepo := postgres.NewBookingRepository(pool)
	bookings := newBookingHandler(bookingRepo)
	pricing := newPricingHandler(pricingapp.NewService(routing.NewStraightLineProvider(30), pricing.Engine{BaseCentsPerKm: 150, MinuteCents: 50}))
	server := &Server{cfg: cfg, pool: pool, mux: http.NewServeMux()}
	server.mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"status": "ok", "service": "vtc-api"})
	})
	server.mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		if err := postgres.Ping(r.Context(), server.pool); err != nil {
			writeJSON(w, http.StatusServiceUnavailable, map[string]any{"status": "not_ready", "reason": "database_unavailable"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]any{"status": "ready", "environment": cfg.Environment})
	})
	server.mux.HandleFunc("POST /api/v1/auth/register", auth.register)
	server.mux.HandleFunc("POST /api/v1/auth/login", auth.login)
	server.mux.Handle("GET /api/v1/me", requireAuth(authRepo, http.HandlerFunc(auth.me)))
	server.mux.Handle("POST /api/v1/bookings", requireAuth(authRepo, http.HandlerFunc(bookings.create)))
	server.mux.Handle("GET /api/v1/bookings/{id}", requireAuth(authRepo, http.HandlerFunc(bookings.get)))
	server.mux.Handle("POST /api/v1/pricing/estimate", requireAuth(authRepo, http.HandlerFunc(pricing.estimate)))
	server.mux.HandleFunc("GET /api/v1", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"version": "v1"})
	})
	return server
}

func (s *Server) Handler() http.Handler {
	return requestID(s.mux)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = time.Now().UTC().Format("20060102T150405.000000000Z")
		}
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r)
	})
}
