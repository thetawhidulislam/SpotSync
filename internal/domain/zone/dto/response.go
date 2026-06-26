package dto

type ZoneResponse struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	TotalCapacity int     `json:"total_capacity"`
	PricePerHour  float64 `json:"price_per_hour"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type APIResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    ZoneResponse `json:"data"`
}

type ListAPIResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    []ZoneResponse `json:"data"`
}
