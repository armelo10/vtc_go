# Architecture

## Cibles
Paris et Île-de-France d'abord, puis multi-ville.

## Couches
- domain : règles métier pures
- application : cas d'usage et transactions
- infrastructure : PostgreSQL/PostGIS, Redis, providers
- interfaces : HTTP, webhooks et futures apps

## Flux
Passenger → Booking → Compliance/Dispatch → Trip → Payment → Invoice → Rating.

## Principes
1. Réservation préalable obligatoire pour un service VTC.
2. Course dispatchée uniquement vers un conducteur conforme.
3. Événements de course historisés.
4. Pricing versionné et explicable.
5. Redis n'est jamais la source de vérité.
6. Paiement, cartographie et notifications sont remplaçables par interfaces.
7. Taxi et VTC partagent le socle technique mais leurs politiques métier restent séparées.
