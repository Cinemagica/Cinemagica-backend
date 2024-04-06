package bootstrap

import (
	"cinemagica/internal/seat/seat"

	"github.com/jmoiron/sqlx"
)

func Bootstrap(db *sqlx.DB) *seat.Service {
	repository := seat.NewRepository(db)
	service := seat.NewService(repository)

	return service
}
