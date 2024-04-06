package booking

type Repository interface {
	Create(reservation *storage) (uint64, error)
}

type BookingSeatService interface {
	Create(bookingID uint64, seatID uint64) error
}

type SeatService interface {
	Update(seatID uint64) error
}

type Service struct {
	repository         Repository
	bookingSeatService BookingSeatService
	seatService        SeatService
}

func NewService(repository Repository, bookingSeatService BookingSeatService, seatService SeatService) *Service {
	return &Service{
		repository:         repository,
		bookingSeatService: bookingSeatService,
		seatService:        seatService,
	}
}

func (s *Service) Create(dto *DTO) error {
	booking := dtoToStorage(dto)

	id, err := s.repository.Create(booking)
	if err != nil {
		return err
	}

	for _, seatID := range dto.SeatIDs {
		err = s.seatService.Update(seatID)
		if err != nil {
			return err
		}

		err := s.bookingSeatService.Create(id, seatID)
		if err != nil {
			return err
		}
	}

	return nil
}
