CREATE TABLE movies
(
    movie_id     BIGSERIAL PRIMARY KEY,
    title        VARCHAR(255),
    genre_id     BIGINT,
    description  TEXT,
    release_date DATE,
    age_rating   VARCHAR(20),
    duration     VARCHAR(20),
    director     VARCHAR(255),
    "cast"       TEXT,
    poster_url   TEXT,
    trailer_url  TEXT
);
