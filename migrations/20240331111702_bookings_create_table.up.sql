CREATE TABLE IF NOT EXISTS bookings
(
    booking_id   SERIAL PRIMARY KEY,
    client_name  VARCHAR(255),
    phone_number VARCHAR(13),
    total_price  DECIMAL(10, 2),
    booking_time TIMESTAMP,
    screen_id    INTEGER,
    deleted_at   TIMESTAMP
);