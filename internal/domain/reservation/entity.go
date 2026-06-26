package reservation

import (
	"gorm.io/gorm"
	"spotsync/internal/domain/user"
	"spotsync/internal/domain/zone"
)

type Reservation struct {
	gorm.Model

	UserID uint `gorm:"not null"`
	ZoneID uint `gorm:"not null"`

	LicensePlate string `gorm:"type:varchar(50);not null"`
	Status       string `gorm:"type:varchar(20);default:active"`

	User user.User `gorm:"foreignKey:UserID"`
	Zone zone.Zone `gorm:"foreignKey:ZoneID"`
}
