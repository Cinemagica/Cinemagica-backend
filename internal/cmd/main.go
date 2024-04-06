package main

import (
	"cinemagica/internal/api"
	bookingBootstrap "cinemagica/internal/booking/bootstrap"
	bookingSeatBootstrap "cinemagica/internal/booking_seat/bootstrap"
	"cinemagica/internal/config"
	movieBootstrap "cinemagica/internal/movie/bootstrap"
	seatBootstrap "cinemagica/internal/seat/bootstrap"
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
	seatService := seatBootstrap.Bootstrap(db)
	bookingHandler := bookingBootstrap.Bootstrap(db, bookingSeatService, seatService)

	// Passing handlers to router
	router := api.NewRouter(movieHandler, bookingHandler)

	// Starting server
	err := http.ListenAndServe(port, router.Route())
	if err != nil {
		fmt.Println("Server cannot be started!")
	}
}
