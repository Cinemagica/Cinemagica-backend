package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type MovieHandler interface {
	GetAllWhichPremier(w http.ResponseWriter, r *http.Request)
}

type BookingHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	movieHandler   MovieHandler
	bookingHandler BookingHandler
}

func NewRouter(movieHandler MovieHandler, bookingHandler BookingHandler) *Router {
	return &Router{
		movieHandler:   movieHandler,
		bookingHandler: bookingHandler,
	}
}

func (r *Router) Route() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger) // logger

	mux.Route("/api/v1", func(v1 chi.Router) {

		v1.Get("/movies", r.movieHandler.GetAllWhichPremier)

		v1.Post("/projections/booking", r.bookingHandler.Create)
	})

	return mux
}
