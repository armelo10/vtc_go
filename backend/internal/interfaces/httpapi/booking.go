package httpapi

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/armelo10/vtc_go/backend/internal/domain/booking"
	"github.com/armelo10/vtc_go/backend/internal/infrastructure/postgres"
)

type bookingHandler struct {
	repo *postgres.BookingRepository
}

type createBookingRequest struct {
	PickupAddress  string              `json:"pickup_address"`
	DropoffAddress string              `json:"dropoff_address"`
	ServiceType    booking.ServiceType `json:"service_type"`
	ScheduledAt   time.Time           `json:"scheduled_at"`
}

func newBookingHandler(repo *postgres.BookingRepository) *bookingHandler {
	return &bookingHandler{repo: repo}
}

func (h *bookingHandler) create(w http.ResponseWriter, r *http.Request) {
	user, ok := userFromContext(r.Context())
	if !ok || user.Role != "PASSENGER" {
		writeJSON(w, http.StatusForbidden, map[string]string{"error": "passenger role required"})
		return
	}

	var req createBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid booking"})
		return
	}

	serviceType := req.ServiceType
	if serviceType == "" {
		serviceType = booking.ServiceVTC
	}
	if serviceType != booking.ServiceVTC && serviceType != booking.ServiceTaxi {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid service_type"})
		return
	}

	now := time.Now().UTC()
	entity := booking.Booking{
		ID:             newUUID(),
		PassengerID:    user.ID,
		PickupAddress: strings.TrimSpace(req.PickupAddress),
		DropoffAddress: strings.TrimSpace(req.DropoffAddress),
		ServiceType:    serviceType,
		Status:         booking.Requested,
		RequestedAt:    now,
		ScheduledAt:    req.ScheduledAt.UTC(),
		Currency:       "EUR",
	}
	if entity.ScheduledAt.IsZero() {
		entity.ScheduledAt = now
	}
	if err := entity.Validate(); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err := h.repo.Create(r.Context(), entity); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "booking creation failed"})
		return
	}
	writeJSON(w, http.StatusCreated, entity)
}

func (h *bookingHandler) get(w http.ResponseWriter, r *http.Request) {
	user, ok := userFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "authentication required"})
		return
	}

	entity, err := h.repo.FindByID(r.Context(), r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "booking not found"})
		return
	}
	if entity.PassengerID != user.ID && user.Role != "ADMIN" {
		writeJSON(w, http.StatusForbidden, map[string]string{"error": "forbidden"})
		return
	}
	writeJSON(w, http.StatusOK, entity)
}

func newUUID() string {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	buf[6] = (buf[6] & 0x0f) | 0x40
	buf[8] = (buf[8] & 0x3f) | 0x80
	raw := hex.EncodeToString(buf)
	return raw[0:8] + "-" + raw[8:12] + "-" + raw[12:16] + "-" + raw[16:20] + "-" + raw[20:32]
}
