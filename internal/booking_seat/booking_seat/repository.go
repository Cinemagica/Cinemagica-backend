package booking_seat

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(bookingID int64, seatID uint8) error {
	query := `
		INSERT INTO bookings_seats (booking_id, seat_id)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(query, bookingID, seatID)
	if err != nil {
		return err
	}

	return nil
}
