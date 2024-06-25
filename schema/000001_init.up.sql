CREATE TABLE IF NOT EXISTS images
(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    filename TEXT NOT NULL,
    original_url TEXT NOT NULL,
    size_large TEXT NOT NULL,
    size_medium TEXT NOT NULL,
    size_small TEXT NOT NULL,
    is_done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);