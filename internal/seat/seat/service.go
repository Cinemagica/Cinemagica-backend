package seat

type Repository interface {
	Update(seatID uint64) error
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
