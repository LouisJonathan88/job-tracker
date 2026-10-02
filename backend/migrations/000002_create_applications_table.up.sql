CREATE TABLE applications (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    company VARCHAR(255) NOT NULL,
    position VARCHAR(255) NOT NULL,
    job_url TEXT,
    applied_at DATE,
    notes TEXT,
    current_status VARCHAR(20) NOT NULL DEFAULT 'wishlist'
        CHECK (current_status IN ('wishlist', 'applied', 'interview', 'offer', 'accepted', 'rejected')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_applications_user_id ON applications(user_id);