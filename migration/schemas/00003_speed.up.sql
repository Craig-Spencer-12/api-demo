CREATE TABLE IF NOT EXISTS speeds (
    id SERIAL PRIMARY KEY,
    truck_id TEXT NOT NULL,
    speed TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
