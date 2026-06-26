package zone

import (
	"spotsync/internal/domain/zone/dto"
	"time"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
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
	zone, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.ZoneResponse
	for _, m := range zone {
		responses = append(responses, *m.ToResponse())
	}
	return responses, nil
}

func (s *service) GetZoneByID(zoneId uint) (*dto.ZoneResponse, error) {
	zone, err := s.repo.GetByID(zoneId)
	if err != nil {
		return nil, err
	}
	return zone.ToResponse(), nil
}
