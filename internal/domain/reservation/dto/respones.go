package dto

type ZoneInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type UserInfo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type ReservationResponse struct {
	ID           uint     `json:"id"`
	LicensePlate string   `json:"license_plate"`
	Status       string   `json:"status"`
	Zone         ZoneInfo `json:"zone"`
	CreatedAt    string   `json:"created_at"`
}

type AdminReservationResponse struct {
	ID           uint   `json:"id"`
	LicensePlate string `json:"license_plate"`
	Status       string `json:"status"`

	User UserInfo `json:"user"`
	Zone ZoneInfo `json:"zone"`

	CreatedAt string `json:"created_at"`
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
type AdminListAPIResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    []*AdminReservationResponse `json:"data"`
}
