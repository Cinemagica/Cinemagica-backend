package bootstrap

import (
	apiv1handler "cinemagica/internal/api/v1/seat"
	"cinemagica/internal/seat/seat"

	"github.com/jmoiron/sqlx"
)

func Bootstrap(db *sqlx.DB) (*apiv1handler.Handler, *seat.Service) {
	repository := seat.NewRepository(db)
	service := seat.NewService(repository)
	handler := apiv1handler.NewHandler(service)

	return handler, service
}
