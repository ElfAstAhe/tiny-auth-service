package repository

import (
	"context"
	"fmt"

	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type UserAdminTraceRepository struct {
	*repository.BaseTraceRepository[*domain.User, string]

	repo domain.UserAdminRepository
}

func NewUserAdminTraceRepository(repo domain.UserRepository) *UserAdminTraceRepository {
	return &UserAdminTraceRepository{
		repo:                repo,
		BaseTraceRepository: repository.NewBaseTraceRepository[*domain.User, string]("UserAdminRepository", repo),
	}
}

func (uat *UserAdminTraceRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	ctx, span := uat.GetTracer().Start(ctx, fmt.Sprintf("%s.FindByName", uat.BaseTraceRepository.GetRepositoryName()))
	span.SetAttributes(attribute.String("param.name", name))
	defer span.End()

	res, err := uat.repo.FindByName(ctx, name)
	if err != nil {
		span.AddEvent("findByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())

		return nil, err
	}

	return res, nil
}
