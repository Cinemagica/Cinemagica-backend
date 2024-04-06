package booking

type Repository interface {
	Create(reservation *storage) (int64, error)
}

type BookingSeatService interface {
	Create(bookingID int64, seatID uint8) error
}

type Service struct {
	repository         Repository
	bookingSeatService BookingSeatService
}

func NewService(repository Repository, bookingSeatService BookingSeatService) *Service {
	return &Service{
		repository:         repository,
		bookingSeatService: bookingSeatService,
	}
}

func (s *Service) Create(dto *DTO) error {
	booking := dtoToStorage(dto)

	id, err := s.repository.Create(booking)
	if err != nil {
		return err
	}

	// Creating booking - seat association
	for _, number := range dto.Seats {
		err := s.bookingSeatService.Create(id, number)
		if err != nil {
			return err
		}
	}

	return nil
}
