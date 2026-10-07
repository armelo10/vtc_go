# Persistance

PostgreSQL/PostGIS est la source de vérité. Le pool pgx est créé au démarrage de l'API.

readyz vérifie désormais une connexion réelle à PostgreSQL ; healthz reste un contrôle de processus.

Les repositories sont définis par le domaine puis implémentés dans infrastructure/postgres. Redis restera réservé au cache, à la disponibilité et aux verrous distribués, jamais aux données transactionnelles.

Les migrations sont actuellement exécutées par l'image PostgreSQL du Docker Compose local. Un runner de migrations versionné sera ajouté avant les déploiements staging/production.
