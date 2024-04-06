package bootstrap

import (
	apiv1handler "cinemagica/internal/api/v1/booking"
	"cinemagica/internal/booking/reservation"

	"github.com/jmoiron/sqlx"
)

func Bootstrap(db *sqlx.DB) *apiv1handler.Handler {
	repository := reservation.NewRepository(db)
	service := reservation.NewService(repository)
	handler := apiv1handler.NewHandler(service)

	return handler
}
