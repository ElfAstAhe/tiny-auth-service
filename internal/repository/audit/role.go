package audit

import (
	"context"

	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/client"
	auditlibdomain "github.com/ElfAstAhe/tiny-audit-service/pkg/domain"
	"github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

type RoleAuditRepository struct {
	*repository.BaseAuditCRUDRepository[*domain.Role, string]
	next domain.RoleRepository
}

var _ domain.RoleRepository = (*RoleAuditRepository)(nil)

func NewRoleRepository(
	source string,
	next domain.RoleRepository,
	auditClient client.DataAuditClient,
	log logger.Logger,
) *RoleAuditRepository {
	res := &RoleAuditRepository{
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

func (rar *RoleAuditRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	return rar.next.FindByName(ctx, name)
}

func (rar *RoleAuditRepository) mapEntityToAuditable(entity *domain.Role) auditlibdomain.Auditable {
	return entity
}
