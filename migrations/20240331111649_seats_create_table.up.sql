CREATE TABLE seats
(
    seat_id      BIGSERIAL PRIMARY KEY,
    seat_number  SMALLINT,
    row          SMALLINT,
    availability BOOLEAN,
    room_id      BIGINT
);