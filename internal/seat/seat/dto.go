package seat

type DTO struct {
	ID           int64
	SeatNumber   int16
	Row          int16
	Availability bool
	RoomID       int64
}
