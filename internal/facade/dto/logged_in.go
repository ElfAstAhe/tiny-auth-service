package dto

type LoggedInDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
