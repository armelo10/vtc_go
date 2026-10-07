# Pricing

## Endpoint

POST /api/v1/pricing/estimate

Authentication is required.

Request:

    {
      "pickup_lat": 48.8566,
      "pickup_lng": 2.3522,
      "dropoff_lat": 48.8584,
      "dropoff_lng": 2.2945
    }

The API currently uses a provider-neutral routing interface with a straight-line fallback for development. The response explicitly marks this as provisional.

The pricing MVP uses:
- 150 cents/km
- 50 cents/minute
- EUR
- version mvp-1

The next production step is to replace the straight-line provider with a real routing/geocoding provider and move tariff parameters to versioned PostgreSQL pricing rules. Client-provided distance and duration are never trusted as pricing inputs.
