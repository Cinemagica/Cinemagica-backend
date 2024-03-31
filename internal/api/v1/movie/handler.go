package movie

import (
	"cinemagica/internal/movie/movie"
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetAll(ctx context.Context) ([]*movie.DTO, error)
	GetPremiers(ctx context.Context) ([]*movie.DTO, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// GetAllWhichPremier godoc
// @Summary Get all movies or only premier movies
// @Description Get all movies or only premier movies
// @Tags movie
// @Accept  json
// @Produce  json
// @Param which-premier query string false "Get only premier movies"
// @Success 200 {array} movie.DTO
// @Failure 500 {object} internalHttp.ErrorResponse
// @Router /api/v1/movies?which_premier={} [get]
func (h *Handler) GetAllWhichPremier(w http.ResponseWriter, r *http.Request) {
	whichPremierQuery := r.URL.Query().Get("which-premier")

	var movies []*movie.DTO
	var err error

	if whichPremierQuery == "true" {
		movies, err = h.service.GetPremiers(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if whichPremierQuery == "false" {
		movies, err = h.service.GetAll(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
}
