package theater

import (
	"github.com/jmoiron/sqlx"
)

type storage struct {
	TheaterID uint64 `db:"theater_id"`
	Name      string `db:"name"`
	Location  string `db:"location"`
	Country   string `db:"country"`
	Rooms     uint64 `db:"rooms"`
}

type repository struct {
	db sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: *db,
	}
}

func (r *repository) GetTheaters() ([]*storage, error) {
	query := `
	SELECT * FROM theaters
	ORDER BY name
	`

	theaters := make([]*storage, 0, 4)

	err := r.db.Select(&theaters, query)
	if err != nil {
		return nil, err
	}

	return theaters, nil
}
