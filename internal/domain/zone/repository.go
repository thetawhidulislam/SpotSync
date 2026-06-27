package zone

import (
	"errors"

	"gorm.io/gorm"
)

var ErrZoneNotFound = errors.New("zone not found")

type Repository interface {
	Create(mango *Zone) error
	GetAll() ([]*Zone, error)
	GetByID(mangoId uint) (*Zone, error)
	CountActiveReservations(zoneID uint) (int64, error)
	// Update(mango *Zone) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(zone *Zone) error {
	return r.db.Create(zone).Error
}

func (r *repository) GetAll() ([]*Zone, error) {
	var zone []*Zone
	if err := r.db.Find(&zone).Error; err != nil {
		return nil, err
	}
	return zone, nil
}
func (r *repository) GetByID(zoneId uint) (*Zone, error) {
	var zone Zone
	err := r.db.First(&zone, zoneId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrZoneNotFound
		}
		return nil, err
	}
	return &zone, nil
}

func (r *repository) CountActiveReservations(zoneID uint) (int64, error) {
	var count int64

	err := r.db.
		Table("reservations").
		Where("zone_id = ? AND status = ?", zoneID, "active").
		Count(&count).Error

	return count, err
}