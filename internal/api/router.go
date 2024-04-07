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

type SeatHandler interface {
	GetSeatsByID(w http.ResponseWriter, r *http.Request)
}

type TheaterHandler interface {
	GetTheaters(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	movieHandler   MovieHandler
	bookingHandler BookingHandler
	seatHandler    SeatHandler
	theaterHandler TheaterHandler
}

func NewRouter(movieHandler MovieHandler, bookingHandler BookingHandler, seatHandler SeatHandler, theaterHandler TheaterHandler) *Router {
	return &Router{
		movieHandler:   movieHandler,
		bookingHandler: bookingHandler,
		seatHandler:    seatHandler,
		theaterHandler: theaterHandler,
	}
}

func (r *Router) Route() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger) // logger

	mux.Route("/api/v1", func(v1 chi.Router) {

		v1.Get("/movies", r.movieHandler.GetAllWhichPremier)

		v1.Post("/projections/booking", r.bookingHandler.Create)

		v1.Get("/seats/{roomID}", r.seatHandler.GetSeatsByID)

		v1.Get("/theaters", r.theaterHandler.GetTheaters)
	})

	return mux
}
