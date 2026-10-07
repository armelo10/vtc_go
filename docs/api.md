# API v1

Base: /api/v1

MVP prévu:
- POST /auth/register
- POST /auth/login
- GET /me
- POST /bookings
- GET /bookings
- GET /bookings/{id}
- POST /bookings/{id}/cancel
- GET /drivers/me
- PUT /drivers/me/availability
- GET /drivers/me/bookings
- POST /drivers/me/bookings/{id}/accept
- POST /payments
- POST /payments/webhook
- GET /invoices
- POST /ratings

Les endpoints seront ajoutés avec RBAC, validation, idempotence et audit.
