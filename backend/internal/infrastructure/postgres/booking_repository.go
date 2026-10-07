package postgres

import (
	"context"
	"time"

	"github.com/armelo10/vtc_go/backend/internal/domain/booking"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingRepository struct {
	pool *pgxpool.Pool
}

func NewBookingRepository(pool *pgxpool.Pool) *BookingRepository {
	return &BookingRepository{pool: pool}
}

func (r *BookingRepository) Create(ctx context.Context, b booking.Booking) error {
	_, err := r.pool.Exec(ctx,
		"INSERT INTO bookings (id, passenger_id, service_type, status, pickup_address, dropoff_address, scheduled_at, requested_at, estimated_amount_cents, currency) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		b.ID, b.PassengerID, b.ServiceType, b.Status, b.PickupAddress, b.DropoffAddress,
		b.ScheduledAt, b.RequestedAt, b.EstimatedCents, b.Currency,
	)
	return err
}

func (r *BookingRepository) FindByID(ctx context.Context, id string) (booking.Booking, error) {
	var b booking.Booking
	var serviceType, status string
	var requestedAt, scheduledAt time.Time

	err := r.pool.QueryRow(ctx,
		"SELECT id, passenger_id, service_type, status, pickup_address, dropoff_address, scheduled_at, requested_at, COALESCE(estimated_amount_cents, 0), currency FROM bookings WHERE id = $1",
		id,
	).Scan(
		&b.ID, &b.PassengerID, &serviceType, &status, &b.PickupAddress, &b.DropoffAddress,
		&scheduledAt, &requestedAt, &b.EstimatedCents, &b.Currency,
	)
	if err != nil {
		return booking.Booking{}, err
	}
	b.ServiceType = booking.ServiceType(serviceType)
	b.Status = booking.Status(status)
	b.ScheduledAt = scheduledAt
	b.RequestedAt = requestedAt
	return b, nil
}
