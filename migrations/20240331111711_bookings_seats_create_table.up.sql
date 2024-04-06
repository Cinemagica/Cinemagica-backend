CREATE TABLE bookings_seats
(
    booking_id BIGINT,
    seat_id    BIGINT,
    PRIMARY KEY (booking_id, seat_id)
);