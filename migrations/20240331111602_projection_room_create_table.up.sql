CREATE TABLE IF NOT EXISTS projection_rooms
(
    room_id     SERIAL PRIMARY KEY,
    room_number SMALLINT,
    seats       SMALLINT,
    "rows"      SMALLINT,
    theater_id  INTEGER
);