package seat

import (
	"cinemagica/internal/seat/seat"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Service interface {
	Update(seatID uint64) error
	GetSeatsByID(roomID uint64) ([]*seat.DTO, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetSeatsByID(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomID")

	parsedInteger, err := strconv.Atoi(roomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	seats, err := h.service.GetSeatsByID(uint64(parsedInteger))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(seats)
	if err != nil {
		return
	}
}
