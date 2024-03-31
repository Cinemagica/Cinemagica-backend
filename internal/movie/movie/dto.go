package movie

import (
	"time"
)

type DTO struct {
	ID          int64
	Title       string
	Genre       int64
	Description string
	Cast        string
	Director    string
	AgeRating   string
	Duration    string
	ReleaseDate time.Time
	PosterURL   string
	TrailerURL  string
}
