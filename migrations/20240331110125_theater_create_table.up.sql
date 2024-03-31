CREATE TABLE IF NOT EXISTS theaters
(
    theater_id    SERIAL PRIMARY KEY,
    "name"        VARCHAR(255),
    "location"    VARCHAR(255),
    room_capacity SMALLINT
);