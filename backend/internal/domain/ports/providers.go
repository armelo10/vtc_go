package ports

import "context"

type PaymentProvider interface {
	CreatePayment(ctx context.Context, amountCents int64, currency, idempotencyKey string) (string, error)
	RefundPayment(ctx context.Context, paymentID string, amountCents int64) error
	HandleWebhook(ctx context.Context, payload []byte, signature string) error
}

type RoutingProvider interface {
	Geocode(ctx context.Context, address string) (lat, lon float64, err error)
	Estimate(ctx context.Context, fromLat, fromLon, toLat, toLon float64) (distanceKm, durationMin float64, err error)
}

type NotificationProvider interface {
	SendSMS(ctx context.Context, to, body string) error
	SendEmail(ctx context.Context, to, subject, body string) error
	SendPush(ctx context.Context, token, title, body string) error
}
