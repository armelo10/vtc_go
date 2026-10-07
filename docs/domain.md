# Domaine métier

Cycle réservation : DRAFT -> REQUESTED -> SEARCHING_DRIVER -> DRIVER_ASSIGNED -> DRIVER_ARRIVING -> DRIVER_AT_PICKUP -> PASSENGER_ONBOARD -> IN_PROGRESS -> COMPLETED.

Les états terminaux sont COMPLETED, CANCELLED, EXPIRED, NO_SHOW et FAILED.

Les transitions sont contrôlées par le domaine. Le dispatch doit vérifier la conformité du conducteur avant affectation. Le pricing est versionné et exprimé en centimes d'euro.

VTC et taxi partagent le socle technique mais conservent des politiques métier distinctes.