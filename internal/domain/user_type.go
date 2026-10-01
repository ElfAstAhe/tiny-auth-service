package domain

import (
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

const (
	// UserTypeGuest defines the standard fallback structural token representation for anonymous or unauthenticated request flows.
	UserTypeGuest string = "guest"
	// UserTypeUser defines the standard operational credential criteria identifier for authenticated physical persons.
	UserTypeUser string = "user"
	// UserTypeService defines the non-human programmatic access account classification tailored for secure machine-to-machine integrations.
	UserTypeService string = "service"
)

var (
	// userTypes encapsulates a fast O(1) memory lookup table schema to execute rapid identity configuration verification bounds.
	userTypes = map[string]struct{}{
		UserTypeGuest:   {},
		UserTypeUser:    {},
		UserTypeService: {},
	}
)

// validateUserType evaluates an inbound system string token against active business configuration allowance boundaries.
func validateUserType(userType string) error {
	_, ok := userTypes[userType]

	if !ok {
		return errs.NewBllValidateError("validateUserType", fmt.Sprintf("user type '%s' is not allowed", userType), nil)
	}

	return nil
}
