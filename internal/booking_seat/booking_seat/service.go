package booking_seat

type Repository interface {
	Create(bookingID uint64, seatID uint64) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(bookingID uint64, seatID uint64) error {
	err := s.repository.Create(bookingID, seatID)
	if err != nil {
		return err
	}

	return nil
}
