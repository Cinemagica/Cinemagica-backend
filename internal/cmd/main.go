package main

import (
	"cinemagica/internal/api"
	"cinemagica/internal/config"
	"fmt"
	"net/http"
)

const (
	connectionStr = "postgres://postgres:bogomil02@localhost:5432/cinema?sslmode=disable"
	port          = ":8080"
)

func main() {
	db, _ := config.ConnectDB(connectionStr)

	defer db.Close()

	// Handler initialization

	// Passing handlers to router
	router := api.NewRouter()

	// Starting server
	err := http.ListenAndServe(port, router.Route())
	if err != nil {
		fmt.Println("Server cannot be started!")
	}
}
