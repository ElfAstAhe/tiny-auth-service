package dto

// LoggedInDTO encapsulates the outbound token session assets payload
// returned to consumers upon a successful authentication event stream execution.
type LoggedInDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token,omitempty"`
} // @name LoggedInDTO
