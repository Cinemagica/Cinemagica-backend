package reservation

import (
	"fmt"
)

type Repository interface {
	Create(reservation *storage) (int64, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(dto *DTO) error {
	booking := dtoToStorage(dto)

	id, err := s.repository.Create(booking)
	if err != nil {
		return err
	}

	fmt.Println("Created booking with ID:", id)

	return nil
}
