# VTC Go

Plateforme VTC Paris / Île-de-France, conçue pour évoluer vers plusieurs villes et plusieurs services de transport.

## Stack

- Go
- PostgreSQL + PostGIS
- Redis
- Docker Compose
- GitHub Actions

## Local

```bash
docker compose up -d
go test ./...
go run ./backend/cmd/api
```

API de socle : `GET /healthz`, `GET /readyz`, `GET /api/v1`.

## Architecture

Le backend sépare domaine, infrastructure et interfaces. Les paiements, le routage et les notifications sont derrière des interfaces provider. Redis n'est pas une source de vérité.

Le cycle de réservation VTC est contrôlé par une machine d'état et la conformité du conducteur est une décision métier explicable.

## Base de données

Les migrations `migrations/` initialisent PostGIS puis les utilisateurs, chauffeurs, véhicules, réservations, courses, pricing, paiements, factures, conformité, notifications et audit.

## Documentation

- `docs/architecture.md`
- `docs/api.md`
- `docs/compliance.md`
- `docs/domain.md`

## Roadmap

1. Persistance/repositories
2. Authentification et RBAC
3. API réservation
4. Dispatch et disponibilité
5. Pricing versionné
6. Paiement et webhooks
7. Notifications
8. Facturation
9. Back-office
10. Applications passager/chauffeur
11. Observabilité, sécurité et déploiement
