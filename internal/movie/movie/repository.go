package movie

import (
	"context"
	"database/sql"
	"time"
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

type dbEntity interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type repository struct {
	db dbEntity
}

func NewRepository(db dbEntity) *repository {
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
	query := `SELECT * FROM movies WHERE release_date > NOW()`

	movies := make([]*storage, 0, defaultSliceCapacityValue)

	err := r.db.SelectContext(ctx, &movies, query)
	if err != nil {
		return nil, err
	}

	return movies, nil
}
