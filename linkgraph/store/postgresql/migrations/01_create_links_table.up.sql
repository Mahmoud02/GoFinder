CREATE TABLE IF NOT EXISTS links (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    url text UNIQUE,
    retrieved_at TIMESTAMP
);