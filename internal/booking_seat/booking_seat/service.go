package booking_seat

type Repository interface {
	Create(bookingID int64, seatID uint8) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(bookingID int64, seatID uint8) error {
	err := s.repository.Create(bookingID, seatID)
	if err != nil {
		return err
	}

	return nil
}
