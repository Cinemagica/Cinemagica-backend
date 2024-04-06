package bootstrap

import (
	apiv1handler "cinemagica/internal/api/v1/booking"
	"cinemagica/internal/booking/booking"
	"cinemagica/internal/booking_seat/booking_seat"
	"cinemagica/internal/seat/seat"

	"github.com/jmoiron/sqlx"
)

func Bootstrap(db *sqlx.DB, bookingSeatService *booking_seat.Service, seatService *seat.Service) *apiv1handler.Handler {
	repository := booking.NewRepository(db)
	service := booking.NewService(repository, bookingSeatService, seatService)
	handler := apiv1handler.NewHandler(service)

	return handler
}
