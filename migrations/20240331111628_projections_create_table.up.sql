CREATE TABLE projections
(
    projection_id BIGSERIAL PRIMARY KEY,
    start_time    TIMESTAMP,
    end_time      TIMESTAMP,
    date          DATE,
    movie_id      BIGINT,
    theater_id    BIGINT,
    room_id       BIGINT
);