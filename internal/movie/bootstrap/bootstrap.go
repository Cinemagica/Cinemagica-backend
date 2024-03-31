package bootstrap

import (
	apiv1handler "cinemagica/internal/api/v1/movie"
	"cinemagica/internal/movie/movie"

	"github.com/jmoiron/sqlx"
)

func Bootstrap(db *sqlx.DB) *apiv1handler.Handler {
	movieRepository := movie.NewRepository(db)
	movieService := movie.NewService(movieRepository)
	movieHandler := apiv1handler.NewHandler(movieService)
	return movieHandler
}
