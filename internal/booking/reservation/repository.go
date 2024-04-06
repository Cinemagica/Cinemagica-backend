package reservation

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type storage struct {
	BookingID   int64     `db:"booking_id"`
	ClientName  string    `db:"client_name"`
	PhoneNumber string    `db:"phone_number"`
	TotalPrice  float64   `db:"total_price"`
	BookingTime time.Time `db:"booking_time"`
	ScreenID    uint64    `db:"screen_id"`
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(booking *storage) (int64, error) {
	query := `
 	INSERT INTO bookings (client_name, phone_number, total_price, booking_time, screen_id)
 	VALUES (:client_name, :phone_number, :total_price, :booking_time, :screen_id)
 	RETURNING booking_id;
 `

	rows, err := r.db.NamedQuery(query, booking)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, errors.New("no rows returned")
	}

	var id int64
	if err := rows.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
