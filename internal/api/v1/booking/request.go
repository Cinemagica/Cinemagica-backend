package booking

type Request struct {
	ClientName   string   `json:"client_name"`
	PhoneNumber  string   `json:"phone_number"`
	TotalPrice   float64  `json:"total_price"`
	BookingTime  string   `json:"booking_time"`
	ProjectionID uint64   `json:"projection_id"`
	SeatIDs      []uint64 `json:"seat_ids"`
}
