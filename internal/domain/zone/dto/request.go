package dto

type CreateZoneRequest struct {
	Name          string  `json:"name" validate:"required"`
	Type          string  `json:"type" validate:"required,oneof=ev_charging regular"`
	TotalCapacity int     `json:"total_capacity" validate:"required,min=1"`
	PricePerHour  float64 `json:"price_per_hour" validate:"required,gt=0"`
}
type UpdateZoneRequest struct {
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	TotalCapacity int     `json:"total_capacity"`
	PricePerHour  float64 `json:"price_per_hour"`
}
