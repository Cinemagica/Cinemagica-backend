package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type MovieHandler interface {
	GetAllWhichPremier(w http.ResponseWriter, r *http.Request)
}

type ReservationHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	movieHandler       MovieHandler
	reservationHandler ReservationHandler
}

func NewRouter(movieHandler MovieHandler, reservationHandler ReservationHandler) *Router {
	return &Router{
		movieHandler:       movieHandler,
		reservationHandler: reservationHandler,
	}
}

func (r *Router) Route() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger) // logger

	mux.Route("/api/v1", func(v1 chi.Router) {

		v1.Get("/movies", r.movieHandler.GetAllWhichPremier)

		v1.Post("/reservations", r.reservationHandler.Create)
	})

	return mux
}
