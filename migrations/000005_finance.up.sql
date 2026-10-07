CREATE TABLE pricing_versions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    code text NOT NULL UNIQUE,
    service_type text NOT NULL,
    valid_from timestamptz NOT NULL,
    valid_to timestamptz,
    rules jsonb NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE payments (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id uuid NOT NULL REFERENCES bookings(id),
    provider text NOT NULL,
    provider_payment_id text UNIQUE,
    amount_cents bigint NOT NULL,
    currency char(3) NOT NULL DEFAULT 'EUR',
    status text NOT NULL,
    idempotency_key text NOT NULL UNIQUE,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE refunds (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    payment_id uuid NOT NULL REFERENCES payments(id),
    amount_cents bigint NOT NULL,
    reason text,
    provider_refund_id text UNIQUE,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE invoices (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id uuid NOT NULL UNIQUE REFERENCES bookings(id),
    invoice_number text NOT NULL UNIQUE,
    amount_cents bigint NOT NULL,
    currency char(3) NOT NULL DEFAULT 'EUR',
    issued_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE ratings (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    booking_id uuid NOT NULL REFERENCES bookings(id),
    author_user_id uuid NOT NULL REFERENCES users(id),
    target_user_id uuid REFERENCES users(id),
    score smallint NOT NULL CHECK (score BETWEEN 1 AND 5),
    comment text,
    created_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (booking_id, author_user_id)
);
