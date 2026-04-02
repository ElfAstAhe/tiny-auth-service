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

type UserTraceRepository struct {
	*repository.BaseCRUDTraceRepository[*domain.User, string]
	repo domain.UserRepository
}

var _ libdomain.CRUDRepository[*domain.User, string] = (*UserTraceRepository)(nil)
var _ domain.UserRepository = (*UserTraceRepository)(nil)

func NewUserTraceRepository(repo domain.UserRepository) *UserTraceRepository {
	return &UserTraceRepository{
		repo:                    repo,
		BaseCRUDTraceRepository: repository.NewBaseCRUDTraceRepository[*domain.User, string]("UserRepository", repo),
	}
}

//goland:noinspection DuplicatedCode
func (utr *UserTraceRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	ctx, span := utr.StartSpan(ctx, fmt.Sprintf("%s.FindByName", utr.BaseCRUDTraceRepository.GetRepositoryName()))
	defer span.End()

	span.SetAttributes(attribute.String("param.name", name))

	res, err := utr.repo.FindByName(ctx, name)
	if err != nil {
		span.AddEvent("FindByName_failed")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return res, err
}
