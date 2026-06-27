package reservation

import (
	"github.com/google/uuid"
	"spotsync/internal/domain/reservation/dto"
	"spotsync/internal/domain/zone"
	"time"
)

type service struct {
	reservationRepo Repository
	zoneRepo        zone.Repository
}

func (r *Reservation) ToResponse() *dto.ReservationResponse {
	return &dto.ReservationResponse{
		ID:           r.ID,
		LicensePlate: r.LicensePlate,
		Status:       r.Status,
		Zone: dto.ZoneInfo{
			ID:   r.Zone.ID,
			Name: r.Zone.Name,
			Type: r.Zone.Type,
		},
		CreatedAt: r.CreatedAt.Format(time.RFC3339),
	}
}
func NewService(reservationRepo Repository, zoneRepo zone.Repository) *service {
	return &service{
		reservationRepo: reservationRepo,
		zoneRepo:        zoneRepo,
	}
}

func generateOrderCode() string {
	return "MG-" + uuid.New().String()
}

func (s *service) CreateOrder(userId uint, req dto.CreateReservationRequest) (*dto.ReservationResponse, error) {
	reservation, err := s.reservationRepo.CreateWithCapacityUpdate(
		userId,
		req.ZoneID,
		req.LicensePlate,
	)
	if err != nil {
		return nil, err
	}

	return reservation.ToResponse(), nil
}

func (s *service) GetMyReservations(userId uint) ([]*dto.ReservationResponse, error) {
	orders, err := s.reservationRepo.GetByUserID(userId)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.ReservationResponse, len(orders))

	for i, o := range orders {
		responses[i] = o.ToResponse()
	}

	return responses, nil
}
