package bootstrap

import (
	"cinemagica/internal/booking_seat/booking_seat"

	"github.com/jmoiron/sqlx"
)

func Bootstrap(db *sqlx.DB) *booking_seat.Service {
	repository := booking_seat.NewRepository(db)
	service := booking_seat.NewService(repository)

	return service
}
