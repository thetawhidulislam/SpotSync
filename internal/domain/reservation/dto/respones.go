package dto

type ReservationResponse struct {
	ID           uint   `json:"id"`
	UserID       uint   `json:"user_id"`
	ZoneID       uint   `json:"zone_id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type APIResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    ReservationResponse `json:"data"`
}
