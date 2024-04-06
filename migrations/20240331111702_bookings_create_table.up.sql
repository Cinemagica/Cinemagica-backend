CREATE TABLE bookings
(
    booking_id    BIGSERIAL PRIMARY KEY,
    client_name   VARCHAR(255),
    phone_number  VARCHAR(20),
    booking_time  TIMESTAMP,
    total_price   NUMERIC(10, 2),
    projection_id BIGINT
);