package dto

type LoginDto struct {
	Username          string `json:"username"`
	EncryptedPassword string `json:"encrypted_password"`
}
