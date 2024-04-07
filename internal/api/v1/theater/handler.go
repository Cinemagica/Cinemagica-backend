package theater

import (
	"cinemagica/internal/theater/theater"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetTheaters() ([]*theater.DTO, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetTheaters(w http.ResponseWriter, r *http.Request) {
	theaters, err := h.service.GetTheaters()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(theaters)
	if err != nil {
		return
	}
}
