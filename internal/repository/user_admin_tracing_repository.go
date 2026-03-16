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
	*repository.BaseCRUDTraceRepository[*domain.User, string]
	repo domain.UserAdminRepository
}

func NewUserAdminTraceRepository(repo domain.UserRepository) *UserAdminTraceRepository {
	return &UserAdminTraceRepository{
		repo:                    repo,
		BaseCRUDTraceRepository: repository.NewBaseCRUDTraceRepository[*domain.User, string]("UserAdminRepository", repo),
	}
}

func (uat *UserAdminTraceRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	ctx, span := uat.StartSpan(ctx, fmt.Sprintf("%s.FindByName", uat.BaseCRUDTraceRepository.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(attribute.String("param.name", name))

	res, err := uat.repo.FindByName(ctx, name)
	if err != nil {
		span.AddEvent("FindByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
