package seat

type Repository interface {
	Update(seatID uint64) error
	GetSeatsByID(roomID uint64) ([]Storage, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Update(seatID uint64) error {
	err := s.repository.Update(seatID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetSeatsByID(roomID uint64) ([]*DTO, error) {
	seats, err := s.repository.GetSeatsByID(roomID)
	if err != nil {
		return nil, err
	}

	var seatsDTO []*DTO

	for _, seat := range seats {
		seatsDTO = append(seatsDTO, &DTO{
			ID:           seat.SeatID,
			SeatNumber:   seat.SeatNumber,
			RoomID:       seat.RoomID,
			Row:          seat.Row,
			Availability: seat.Availability,
		})
	}

	return seatsDTO, nil
}
