package dto

type ChangePasswordDTO struct {
	OldPassword string `json:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
} // @name ChangePasswordDTO
