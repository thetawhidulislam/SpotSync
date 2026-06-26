package zone

import (
	"spotsync/internal/domain/zone/dto"

	"gorm.io/gorm"
)

type Zone struct {
	gorm.Model
	Name          string  `gorm:"type:varchar(100);not null"`
	Type          string  `gorm:"type:varchar(50);not null"`
	TotalCapacity int     `gorm:"not null"`
	PricePerHour  float64 `gorm:"type:numeric(10,2);not null"`
}

func (m *Zone) ToResponse() *dto.ZoneResponse {
	return &dto.ZoneResponse{
		ID:            m.ID,
		Name:          m.Name,
		Type:          m.Type,
		TotalCapacity: m.TotalCapacity,
		PricePerHour:  m.PricePerHour,
		CreatedAt:     m.CreatedAt.String(),
	}
}
