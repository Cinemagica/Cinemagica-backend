package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Route() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger) // logger

	mux.Route("/api/v1", func(v1 chi.Router) {

	})

	return mux
}
