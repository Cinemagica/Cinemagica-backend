package reservation

import (
	"time"
)

type DTO struct {
	ClientName  string
	PhoneNumber string
	TotalPrice  float64
	BookingTime time.Time
	ScreenID    uint64
}
