CREATE TABLE rooms
(
    room_id     BIGSERIAL PRIMARY KEY,
    room_number SMALLINT,
    seats       SMALLINT,
    rows        SMALLINT,
    theater_id  BIGINT
);