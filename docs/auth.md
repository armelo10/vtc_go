# Authentification

Le MVP utilise des sessions opaques plutôt que des JWT.

- Le mot de passe est hashé avec bcrypt.
- Le token aléatoire n'est jamais stocké en clair : seul son SHA-256 est conservé.
- Les sessions expirent après 30 jours.
- Les endpoints protégés attendent Authorization: Bearer <token>.
- /api/v1/me restitue l'utilisateur authentifié.
- Le rôle est conservé sur users et sera utilisé par les politiques RBAC des prochains lots.

Avant production : rotation/révocation explicite des sessions, limitation de tentatives, vérification email/téléphone, politique de durée de session par niveau de risque, journalisation d'audit et secrets hors dépôt.
