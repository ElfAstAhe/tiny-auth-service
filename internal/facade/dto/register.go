package dto

// RegisterDTO encapsulates inbound credential criteria parameters payload
// required to process public user identity account creation requests.
type RegisterDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
} // @name RegisterDTO
