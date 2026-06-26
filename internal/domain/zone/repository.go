package zone

import (
	"errors"

	"gorm.io/gorm"
)

var ErrZoneNotFound = errors.New("zone not found")

type Repository interface {
	Create(mango *Zone) error
	// GetAll() ([]*Zone, error)
	// GetByID(mangoId uint) (*Zone, error)
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
