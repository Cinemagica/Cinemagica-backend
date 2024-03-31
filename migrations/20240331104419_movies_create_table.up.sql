CREATE TABLE IF NOT EXISTS movies
(
    movie_id     SERIAL PRIMARY KEY,
    title        VARCHAR(255),
    genre_id     VARCHAR(255),
    description  VARCHAR(255),
    "cast"       VARCHAR(255),
    director     VARCHAR(255),
    age_rating   VARCHAR(255),
    duration     VARCHAR(20),
    release_date TIMESTAMP,
    poster_url   VARCHAR(255),
    trailer_url  VARCHAR(255)
);
