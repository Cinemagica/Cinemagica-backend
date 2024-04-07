package theater

type Repository interface {
	GetTheaters() ([]*storage, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetTheaters() ([]*DTO, error) {
	theaters, err := s.repository.GetTheaters()
	if err != nil {
		return nil, err
	}

	var dtos []*DTO
	for _, theater := range theaters {
		dtos = append(dtos, &DTO{
			ID:       theater.TheaterID,
			Name:     theater.Name,
			Location: theater.Location,
			Country:  theater.Country,
			Rooms:    theater.Rooms,
		})
	}

	return dtos, nil
}
