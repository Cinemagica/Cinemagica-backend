package main

import (
	"cinemagica/internal/api"
	reservationBootstrap "cinemagica/internal/booking/bootstrap"
	bookingSeatBootstrap "cinemagica/internal/booking_seat/bootstrap"
	"cinemagica/internal/config"
	movieBootstrap "cinemagica/internal/movie/bootstrap"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	connectionStr = "postgres://postgres:bogomil02@localhost:5432/cinema?sslmode=disable"
	port          = ":8080"
)

func main() {
	db, _ := config.ConnectDB(connectionStr)

	defer db.Close()

	// Handler initialization
	movieHandler := movieBootstrap.Bootstrap(db)

	bookingSeatService := bookingSeatBootstrap.Bootstrap(db)
	reservationHandler := reservationBootstrap.Bootstrap(db, bookingSeatService)

	// Passing handlers to router
	router := api.NewRouter(movieHandler, reservationHandler)

	// Starting server
	err := http.ListenAndServe(port, router.Route())
	if err != nil {
		fmt.Println("Server cannot be started!")
	}
}
