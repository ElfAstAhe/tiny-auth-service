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
	*repository.BaseTraceRepository[*domain.Role, string]

	repo domain.RoleAdminRepository
}

func NewRoleAdminTraceRepository(repo domain.RoleAdminRepository) *RoleAdminTraceRepository {
	return &RoleAdminTraceRepository{
		repo:                repo,
		BaseTraceRepository: repository.NewBaseTraceRepository[*domain.Role, string]("RoleAdminRepository", repo),
	}
}

func (rat *RoleAdminTraceRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	ctx, span := rat.GetTracer().Start(ctx, fmt.Sprintf("%s.FindByName", rat.BaseTraceRepository.GetRepositoryName()))
	span.SetAttributes(attribute.String("param.name", name))
	defer span.End()

	res, err := rat.repo.FindByName(ctx, name)
	if err != nil {
		span.AddEvent("findByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	return res, nil
}
