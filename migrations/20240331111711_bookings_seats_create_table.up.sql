CREATE TABLE IF NOT EXISTS bookings_seats
(
    booking_id INTEGER,
    seat_id    INTEGER,
    PRIMARY KEY (booking_id, seat_id)
);