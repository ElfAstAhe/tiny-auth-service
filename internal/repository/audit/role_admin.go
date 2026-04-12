package audit

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	auditlibdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type RoleAdminAuditRepository struct {
	*repository.BaseAuditCRUDRepository[*domain.Role, string]
	next domain.RoleAdminRepository
}

var _ domain.RoleAdminRepository = (*RoleAdminAuditRepository)(nil)

func NewRoleAdminRepository(
	source string,
	next domain.RoleAdminRepository,
	auditClient client.DataAuditClient,
	log logger.Logger,
) *RoleAdminAuditRepository {
	res := &RoleAdminAuditRepository{
		next: next,
	}

	res.BaseAuditCRUDRepository = repository.NewBaseAuditCRUDRepository[*domain.Role, string](
		next,
		source,
		res.mapEntityToAuditable,
		auditClient,
		log,
	)

	return res
}

func (raa *RoleAdminAuditRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	return raa.next.FindByName(ctx, name)
}

func (raa *RoleAdminAuditRepository) mapEntityToAuditable(entity *domain.Role) auditlibdomain.Auditable {
	return entity
}
