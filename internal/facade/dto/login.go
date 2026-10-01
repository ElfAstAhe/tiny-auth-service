package dto

// LoginDTO encapsulates the inbound user credential payload
// required to process standard identity authentication requests.
type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
} // @name LoginDTO
