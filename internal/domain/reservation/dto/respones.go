package dto

type ZoneInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ReservationResponse struct {
	ID           uint     `json:"id"`
	LicensePlate string   `json:"license_plate"`
	Status       string   `json:"status"`
	Zone         ZoneInfo `json:"zone"`
	CreatedAt    string   `json:"created_at"`
}

type APIResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    ReservationResponse `json:"data"`
}
type ListAPIResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    []*ReservationResponse `json:"data"`
}
