# Réservations

Les endpoints de réservation sont authentifiés.

- POST /api/v1/bookings
- GET /api/v1/bookings/{id}

Le passenger_id n'est jamais accepté dans la requête de création : il vient de la session Bearer.

Le service par défaut est VTC. La réservation VTC exige une date/heure de réservation. Le dispatch et le pricing complet seront ajoutés dans les lots suivants.

Un accès à une réservation est limité à son propriétaire, sauf rôle ADMIN.
