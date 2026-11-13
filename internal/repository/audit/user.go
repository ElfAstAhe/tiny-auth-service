package audit

import (
	"context"

	libdom "github.com/ElfAstAhe/go-service-template/pkg/domain"
	libextdom "github.com/ElfAstAhe/go-service-template/pkg/ext/domain"
	libauderrs "github.com/ElfAstAhe/tiny-audit-service/pkg/errs"
	libaudrepo "github.com/ElfAstAhe/tiny-audit-service/pkg/repository"
	libaudutl "github.com/ElfAstAhe/tiny-audit-service/pkg/utils"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// UserAuditRepository implements the domain.UserRepository interface,
// acting as a non-invasive database auditing decorator component for user entities.
// It intercepts destructive data mutations to automatically extract, map, and stream audit records via a broker client.
type UserAuditRepository struct {
	*libaudrepo.AuditCRUD[*domain.User, *libextdom.DataAuditOutbox, string]
	next domain.UserRepository // Downstream storage persistence or tracing layer execution implementation path
}

// Compile-time interface compliance verification
var _ libdom.CRUDRepository[*domain.User, string] = (*UserAuditRepository)(nil)
var _ domain.UserRepository = (*UserAuditRepository)(nil)

// NewUserRepository acts as a factory constructor mounting audit logging proxies over standard user storage adapters.
func NewUserRepository(
	source string,
	next domain.UserRepository,
	outboxRepo libextdom.OutboxRepository[*libextdom.DataAuditOutbox, string],
) (*UserAuditRepository, error) {
	// instance
	res := &UserAuditRepository{
		next: next,
	}
	// embed
	embed, err := libaudrepo.NewAuditCRUD[*domain.User, *libextdom.DataAuditOutbox, string](
		libaudrepo.WithAuditSource[*domain.User, *libextdom.DataAuditOutbox, string](source),
		libaudrepo.WithAuditNext[*domain.User, *libextdom.DataAuditOutbox, string](next),
		libaudrepo.WithAuditOutboxBuilder[*domain.User, *libextdom.DataAuditOutbox, string](libaudutl.DefaultBuildDataAuditOutbox),
		libaudrepo.WithAuditOutboxRepository[*domain.User, *libextdom.DataAuditOutbox, string](outboxRepo),
	)
	if err != nil {
		return nil, libauderrs.NewAuditError("NewRoleRepository", "unable to create audit CRUD embed", err)
	}
	// setup
	res.AuditCRUD = embed

	return res, nil
}

// FindByName proxies name-based query execution straight to the downstream data adapter, bypassing change auditing layers.
func (uar *UserAuditRepository) FindByName(ctx context.Context, name string) (*domain.User, error) {
	return uar.next.FindByName(ctx, name)
}
