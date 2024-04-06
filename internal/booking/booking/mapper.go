package booking

func dtoToStorage(dto *DTO) *storage {
	return &storage{
		ClientName:   dto.ClientName,
		PhoneNumber:  dto.PhoneNumber,
		TotalPrice:   dto.TotalPrice,
		BookingTime:  dto.BookingTime,
		ProjectionID: dto.ProjectionID,
	}
}
