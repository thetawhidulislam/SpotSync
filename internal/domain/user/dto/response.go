package dto

type UserResponse struct {
	ID           uint   `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	Role         string `json:"role,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

type APIResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}
