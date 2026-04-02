package trace

import (
	"context"
	"fmt"

	libdomain "github.com/ElfAstAhe/go-service-template/pkg/domain"
	"github.com/ElfAstAhe/go-service-template/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type RoleTraceRepository struct {
	*repository.BaseCRUDTraceRepository[*domain.Role, string]
	repo domain.RoleRepository
}

var _ libdomain.CRUDRepository[*domain.Role, string] = (*RoleTraceRepository)(nil)
var _ domain.RoleRepository = (*RoleTraceRepository)(nil)

func NewRoleTraceRepository(repo domain.RoleRepository) *RoleTraceRepository {
	return &RoleTraceRepository{
		repo:                    repo,
		BaseCRUDTraceRepository: repository.NewBaseCRUDTraceRepository[*domain.Role, string]("RoleRepository", repo),
	}
}

func (rtr *RoleTraceRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	ctx, span := rtr.StartSpan(ctx, fmt.Sprintf("%s.FindByName", rtr.BaseCRUDTraceRepository.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(attribute.String("param.name", name))

	res, err := rtr.repo.FindByName(ctx, name)
	if err != nil {
		span.AddEvent("FindByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
