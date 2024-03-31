CREATE TABLE IF NOT EXISTS screens
(
    screen_id SERIAL PRIMARY KEY,
    start_time   TIMESTAMP,
    end_time     TIMESTAMP,
    date         DATE,
    movie_id     INTEGER,
    room_id      INTEGER,
    theater_id   INTEGER,
    deleted_at   TIMESTAMP
);