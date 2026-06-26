package zone

import "gorm.io/gorm"

type Zone struct {
	gorm.Model
	Name          string  `gorm:"type:varchar(100);not null"`
	Type          string  `gorm:"type:varchar(50);not null"`
	TotalCapacity int     `gorm:"not null"`
	PricePerHour  float64 `gorm:"type:numeric(10,2);not null"`
}
