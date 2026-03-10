package repository

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type RoleTraceRepository struct {
	*repository.BaseCRUDTraceRepository[*domain.Role, string]
	repo domain.RoleRepository
}

func NewRoleTraceRepository(repo domain.RoleRepository) *RoleTraceRepository {
	return &RoleTraceRepository{
		repo:                    repo,
		BaseCRUDTraceRepository: repository.NewBaseCRUDTraceRepository[*domain.Role, string]("RoleRepository", repo),
	}
}

func (rtr *RoleTraceRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	ctx, span := rtr.GetTracer().Start(ctx, fmt.Sprintf("%s.FindByName", rtr.BaseCRUDTraceRepository.GetRepositoryName()))
	span.SetAttributes(attribute.String("param.name", name))
	defer span.End()

	res, err := rtr.repo.FindByName(ctx, name)
	if err != nil {
		span.AddEvent("FindByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	return res, nil
}
