package seat

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

type Storage struct {
	SeatID       int64 `db:"seat_id"`
	SeatNumber   int16 `db:"seat_number"`
	Row          int16 `db:"row"`
	Availability bool  `db:"availability"`
	RoomID       int64 `db:"room_id"`
}

func (r *repository) Update(seatID uint64) error {
	query := `
		UPDATE seats
		SET availability =  NOT availability
		WHERE seat_id = $1
	`

	_, err := r.db.Exec(query, seatID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetSeatsByID(roomID uint64) ([]Storage, error) {
	query := `
		SELECT seat_id, seat_number, row, availability, room_id
		FROM seats
		WHERE room_id = $1
	`

	var seats []Storage
	err := r.db.Select(&seats, query, roomID)
	if err != nil {
		return nil, err
	}

	return seats, nil
}
