package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(connectionStr string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connectionStr)
	if err != nil {
		fmt.Println("Connection unsuccessful!: ", err)
		return nil, err
	}

	fmt.Println("Connection successful!")

	return db, nil
}
