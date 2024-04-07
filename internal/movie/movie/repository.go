package movie

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	defaultSliceCapacityValue = 4
)

type storage struct {
	MovieID     int64     `db:"movie_id"`
	Title       string    `db:"title"`
	GenreID     int64     `db:"genre_id"`
	Description string    `db:"description"`
	Cast        string    `db:"cast"`
	Director    string    `db:"director"`
	AgeRating   string    `db:"age_rating"`
	Duration    string    `db:"duration"`
	ReleaseDate time.Time `db:"release_date"`
	PosterURL   string    `db:"poster_url"`
	TrailerURL  string    `db:"trailer_url"`
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]*storage, error) {
	query := `SELECT * FROM movies`

	movies := make([]*storage, 0, defaultSliceCapacityValue)

	err := r.db.SelectContext(ctx, &movies, query)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *repository) GetPremiers(ctx context.Context) ([]*storage, error) {
	query := `
	SELECT * 
	FROM movies 
	WHERE release_date > NOW()
	ORDER BY release_date DESC
	`

	movies := make([]*storage, 0, defaultSliceCapacityValue)

	err := r.db.SelectContext(ctx, &movies, query)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
