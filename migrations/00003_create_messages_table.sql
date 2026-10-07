-- +goose Up
SELECT 'up SQL query';
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sender_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL CHECK (length(content) > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CHECK (sender_id <> receiver_id)
);

CREATE INDEX idx_messages_sender_receiver ON messages(sender_id, receiver_id, created_at DESC);
CREATE INDEX idx_messages_receiver_sender ON messages(receiver_id, sender_id, created_at DESC);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE IF EXISTS messages;