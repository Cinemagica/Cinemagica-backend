package booking

import (
	"cinemagica/internal/booking/booking"
	"encoding/json"
	"net/http"
	"time"
)

type Service interface {
	Create(dto *booking.DTO) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validation
	parsedTime, err := time.Parse("2006-01-02 15:04", req.BookingTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto := booking.DTO{
		ClientName:   req.ClientName,
		PhoneNumber:  req.PhoneNumber,
		TotalPrice:   req.TotalPrice,
		BookingTime:  parsedTime,
		ProjectionID: req.ProjectionID,
		SeatIDs:      req.SeatIDs,
	}

	err = h.service.Create(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(dto)
	if err != nil {
		return
	}
}
