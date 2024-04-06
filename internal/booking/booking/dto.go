package booking

import (
	"time"
)

type DTO struct {
	ClientName   string
	PhoneNumber  string
	TotalPrice   float64
	BookingTime  time.Time
	ProjectionID uint64
	SeatIDs      []uint64
}
