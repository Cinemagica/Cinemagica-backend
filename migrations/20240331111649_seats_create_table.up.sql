CREATE TABLE IF NOT EXISTS seats
(
    seat_id      SERIAL PRIMARY KEY,
    "row"        SMALLINT,
    seat_number  SMALLINT,
    availability BOOLEAN,
    screening_id INTEGER
);