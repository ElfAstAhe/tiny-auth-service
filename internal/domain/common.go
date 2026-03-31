package domain

import (
	"github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain/errs"
	"github.com/google/uuid"
)

func defaultBeforeCreate(entity domain.Entity[string]) error {
	newID, err := uuid.NewRandom()
	if err != nil {
		return errs.NewBllError("defaultBeforeCreate", "generate new id", err)
	}

	entity.SetID(newID.String())

	return nil
}
