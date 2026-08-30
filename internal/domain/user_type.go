package domain

import (
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

const (
	UserTypeGuest   string = "guest"
	UserTypeUser    string = "user"
	UserTypeService string = "service"
)

var (
	userTypes = map[string]struct{}{
		UserTypeGuest:   {},
		UserTypeUser:    {},
		UserTypeService: {},
	}
)

func validateUserType(userType string) error {
	_, ok := userTypes[userType]

	if !ok {
		return errs.NewBllValidateError("validateUserType", fmt.Sprintf("user type '%s' is not allowed", userType), nil)
	}

	return nil
}
