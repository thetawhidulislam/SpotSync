package reservation

import (
	"errors"
	"spotsync/internal/domain/zone"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrOrderNotFound         = errors.New("order not found")
	ErrNotEnoughStock        = errors.New("not enough stock available")
	ErrOrderAlreadyCancelled = errors.New("order already cancelled")
	ErrForbiddenOrderAccess  = errors.New("you do not own this order")

	ErrZoneFull = errors.New("parking zone is full")
)

const (
	ReservationActive    = "active"
	ReservationCancelled = "cancelled"
)

type Repository interface {
	Create(order *Reservation) error
	// GetByID(orderId uint) (*Reservation, error)
	GetByUserID(userId uint) ([]*Reservation, error)
	// Update(order *Reservation) error
	CreateWithCapacityUpdate(userId, ZoneID uint, LicensePlate string) (*Reservation, error)
	CountActiveReservations(zoneID uint) (int64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(order *Reservation) error {
	return r.db.Create(order).Error
}

func (r *repository) CreateWithCapacityUpdate(
	userId uint,
	zoneId uint,
	licensePlate string,
) (*Reservation, error) {

	var reservation Reservation

	err := r.db.Transaction(func(tx *gorm.DB) error {

		// Step 1: Lock the zone row
		var zoneData zone.Zone

		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&zoneData, zoneId).Error; err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return zone.ErrZoneNotFound
			}
			return err
		}

		// Step 2: Count active reservations
		var activeCount int64

		if err := tx.
			Model(&Reservation{}).
			Where("zone_id = ? AND status = ?", zoneId, ReservationActive).
			Count(&activeCount).Error; err != nil {
			return err
		}

		// Step 3: Check capacity
		if activeCount >= int64(zoneData.TotalCapacity) {
			return ErrZoneFull
		}

		// Step 4: Create reservation
		reservation = Reservation{
			UserID:       userId,
			ZoneID:       zoneId,
			LicensePlate: licensePlate,
			Status:       ReservationActive,
		}

		if err := tx.Create(&reservation).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *repository) GetByUserID(userId uint) ([]*Reservation, error) {

	var reservations []*Reservation

	err := r.db.
		Preload("Zone").
		Where("user_id = ?", userId).
		Find(&reservations).Error

	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *repository) Update(order *Reservation) error {
	return r.db.Save(order).Error
}

func (r *repository) CountActiveReservations(zoneID uint) (int64, error) {
	var count int64

	err := r.db.
		Model(&Reservation{}).
		Where("zone_id = ? AND status = ?", zoneID, ReservationActive).
		Count(&count).Error

	return count, err
}
