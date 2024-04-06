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
