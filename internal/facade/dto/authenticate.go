package dto

type AuthenticateDto struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
