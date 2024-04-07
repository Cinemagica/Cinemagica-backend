CREATE TABLE seats
(
    seat_id      SERIAL PRIMARY KEY,
    seat_number  INTEGER,
    row          INTEGER,
    availability BOOLEAN,
    room_id      INTEGER
);