package facade

import (
	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// IsSubjectAdmin evaluates the validity of an inbound auth.Subject identity contract,
// executing a rapid role boundary check to confirm administrative RBAC privileges.
func IsSubjectAdmin(subject *auth.Subject) bool {
	if subject == nil {
		return false
	}

	return subject.HasRole(domain.RoleAdmin)
}
