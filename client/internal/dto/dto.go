package dto

type Response struct {
	Status       string `json:"status"`
	Error        string `json:"error,omitempty"`
	UserID       string `json:"user_id,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type RegisterRequest struct {
	Avatar   string `json:"avatar"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	UserID string `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Birth  string `json:"birthday,omitempty"`
	Email  string `json:"email,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}
