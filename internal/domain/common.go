package domain

import (
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
	"github.com/google/uuid"
)

func defaultBeforeCreate(entity domain.Entity[string]) error {
	newID, err := uuid.NewV7()
	if err != nil {
		return errs.NewBllError("defaultBeforeCreate", "generate new id", err)
	}

	entity.SetID(newID.String())

	return nil
}

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
