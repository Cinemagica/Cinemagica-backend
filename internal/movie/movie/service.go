package movie

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*storage, error)
	GetPremiers(ctx context.Context) ([]*storage, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetAll(ctx context.Context) ([]*DTO, error) {
	storages, _ := s.repository.GetAll(ctx)

	return storagesToDTOs(storages), nil
}

func (s *Service) GetPremiers(ctx context.Context) ([]*DTO, error) {
	storages, _ := s.repository.GetPremiers(ctx)

	return storagesToDTOs(storages), nil
}
