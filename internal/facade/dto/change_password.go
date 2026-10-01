package dto

// ChangePasswordDTO encapsulates the inbound user payload fields
// required to process account credential updates and password rotation requests.
type ChangePasswordDTO struct {
	OldPassword string `json:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
} // @name ChangePasswordDTO
