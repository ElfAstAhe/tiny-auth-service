package audit

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	auditlibdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserAuditRepository struct {
	*repository.BaseAuditCRUDRepository[*domain.User, string]
	next domain.UserRepository
}

var _ domain.UserRepository = (*UserAuditRepository)(nil)

func NewUserRepository(
	source string,
	next domain.UserRepository,
	auditClient client.DataAuditClient,
	log logger.Logger,
) *UserAuditRepository {
	res := &UserAuditRepository{
		next: next,
	}

	res.BaseAuditCRUDRepository = repository.NewBaseAuditCRUDRepository[*domain.User, string](
		next,
		source,
		res.mapEntityToAuditable,
		auditClient,
		log,
	)

	return res
}

func (uar *UserAuditRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	return uar.next.FindByName(ctx, name)
}

func (uar *UserAuditRepository) mapEntityToAuditable(entity *domain.User) auditlibdomain.Auditable {
	return entity
}
