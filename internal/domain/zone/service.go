package zone

import (
	"spotsync/internal/domain/zone/dto"
	"time"
)

type service struct {
	repo Repository
}

func NewService(
	repo Repository,
) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateZone(req dto.CreateZoneRequest) (*dto.ZoneResponse, error) {

	zone := Zone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	if err := s.repo.Create(&zone); err != nil {
		return nil, err
	}

	return &dto.ZoneResponse{
		ID:            zone.ID,
		Name:          zone.Name,
		Type:          zone.Type,
		TotalCapacity: zone.TotalCapacity,
		PricePerHour:  zone.PricePerHour,
		CreatedAt:     zone.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *service) GetZone() ([]dto.ZoneResponse, error) {

	zones, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var response []dto.ZoneResponse

	for _, z := range zones {

		activeReservations, err := s.repo.CountActiveReservations(z.ID)
		if err != nil {
			return nil, err
		}

		response = append(response, dto.ZoneResponse{
			ID:             z.ID,
			Name:           z.Name,
			Type:           z.Type,
			TotalCapacity:  z.TotalCapacity,
			AvailableSpots: z.TotalCapacity - int(activeReservations),
			PricePerHour:   z.PricePerHour,
			CreatedAt:      z.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      z.UpdatedAt.Format(time.RFC3339),
		})
	}

	return response, nil
}
func (s *service) GetZoneByID(zoneId uint) (*dto.ZoneResponse, error) {

	zone, err := s.repo.GetByID(zoneId)
	if err != nil {
		return nil, err
	}

	activeReservations, err := s.repo.CountActiveReservations(zone.ID)
	if err != nil {
		return nil, err
	}
	response := &dto.ZoneResponse{
		ID:             zone.ID,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.TotalCapacity - int(activeReservations),
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      zone.UpdatedAt.Format(time.RFC3339),
	}

	return response, nil
}
func (s *service) UpdateZone(id uint, req dto.UpdateZoneRequest) (*dto.ZoneResponse, error) {

	zone, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	zone.Name = req.Name
	zone.Type = req.Type
	zone.TotalCapacity = req.TotalCapacity
	zone.PricePerHour = req.PricePerHour

	if err := s.repo.Update(zone); err != nil {
		return nil, err
	}

	return zone.ToResponse(), nil
}

func (s *service) DeleteZone(id uint) error {

	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
