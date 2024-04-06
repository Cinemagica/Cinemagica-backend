package booking

type Request struct {
	ClientName   string  `json:"client_name"`
	PhoneNumber  string  `json:"phone_number"`
	TotalPrice   float64 `json:"total_price"`
	BookingTime  string  `json:"booking_time"`
	ProjectionID uint64  `json:"screen_id"`
	Seats        []uint8 `json:"seats"`
}
