package repository

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type RoleAdminTraceRepository struct {
	*repository.BaseCRUDTraceRepository[*domain.Role, string]
	repo domain.RoleAdminRepository
}

func NewRoleAdminTraceRepository(repo domain.RoleAdminRepository) *RoleAdminTraceRepository {
	return &RoleAdminTraceRepository{
		repo:                    repo,
		BaseCRUDTraceRepository: repository.NewBaseCRUDTraceRepository[*domain.Role, string]("RoleAdminRepository", repo),
	}
}

func (rat *RoleAdminTraceRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	ctx, span := rat.StartSpan(ctx, fmt.Sprintf("%s.FindByName", rat.BaseCRUDTraceRepository.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(attribute.String("param.name", name))

	res, err := rat.repo.FindByName(ctx, name)
	if err != nil {
		span.AddEvent("FindByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	return res, nil
}
