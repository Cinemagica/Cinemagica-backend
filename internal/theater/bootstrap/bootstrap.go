package bootstrap

import (
	"cinemagica/internal/theater/theater"

	"github.com/jmoiron/sqlx"

	apiv1handler "cinemagica/internal/api/v1/theater"
)

func Bootstrap(db *sqlx.DB) *apiv1handler.Handler {
	repository := theater.NewRepository(db)
	service := theater.NewService(repository)
	handler := apiv1handler.NewHandler(service)

	return handler
}
