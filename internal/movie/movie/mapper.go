package movie

func storageToDTO(storage *storage) *DTO {
	return &DTO{
		ID:          storage.MovieID,
		Title:       storage.Title,
		Genre:       storage.GenreID,
		Description: storage.Description,
		Cast:        storage.Cast,
		Director:    storage.Director,
		AgeRating:   storage.AgeRating,
		Duration:    storage.Duration,
		ReleaseDate: storage.ReleaseDate,
		PosterURL:   storage.PosterURL,
		TrailerURL:  storage.TrailerURL,
	}
}

func storagesToDTOs(storage []*storage) []*DTO {
	DTOs := make([]*DTO, 0, len(storage))

	for _, s := range storage {
		DTOs = append(DTOs, storageToDTO(s))
	}
	return DTOs
}
