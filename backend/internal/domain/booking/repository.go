package booking

import "context"

type Repository interface {
	Create(context.Context, Booking) error
	FindByID(context.Context, string) (Booking, error)
}
