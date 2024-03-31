CREATE TABLE genres
(
    genre_id SERIAL PRIMARY KEY,
    name     VARCHAR(255) UNIQUE
);