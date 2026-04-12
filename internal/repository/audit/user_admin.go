package audit

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	auditlibdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type UserAdminAuditRepository struct {
	*repository.BaseAuditCRUDRepository[*domain.User, string]
	next domain.UserAdminRepository
}

var _ domain.UserAdminRepository = (*UserAdminAuditRepository)(nil)

func NewUserAdminRepository(
	source string,
	next domain.UserAdminRepository,
	auditClient client.DataAuditClient,
	log logger.Logger,
) *UserAdminAuditRepository {
	res := &UserAdminAuditRepository{
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

func (uaa *UserAdminAuditRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	return uaa.next.FindByName(ctx, name)
}

func (uaa *UserAdminAuditRepository) mapEntityToAuditable(entity *domain.User) auditlibdomain.Auditable {
	return entity
}
