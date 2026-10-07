CREATE TABLE trips (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id uuid NOT NULL UNIQUE REFERENCES bookings(id),
    driver_id uuid NOT NULL REFERENCES drivers(user_id),
    vehicle_id uuid NOT NULL REFERENCES vehicles(id),
    started_at timestamptz,
    completed_at timestamptz,
    distance_m integer,
    duration_seconds integer,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE trip_events (
    id bigserial PRIMARY KEY,
    trip_id uuid NOT NULL REFERENCES trips(id) ON DELETE CASCADE,
    event_type text NOT NULL,
    occurred_at timestamptz NOT NULL DEFAULT now(),
    latitude double precision,
    longitude double precision,
    metadata jsonb NOT NULL DEFAULT '{}'::jsonb
);

CREATE INDEX trip_events_trip_time_idx ON trip_events (trip_id, occurred_at);
