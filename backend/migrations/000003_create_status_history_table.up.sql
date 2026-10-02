CREATE TABLE status_history (
    id BIGSERIAL PRIMARY KEY,
    application_id BIGINT NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    from_status VARCHAR(20)
        CHECK (from_status IN ('wishlist', 'applied', 'interview', 'offer', 'accepted', 'rejected')),
    to_status VARCHAR(20) NOT NULL
        CHECK (to_status IN ('wishlist', 'applied', 'interview', 'offer', 'accepted', 'rejected')),
    changed_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_status_history_application_id ON status_history(application_id);