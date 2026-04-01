package facade

import (
	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

func IsSubjectAdmin(subject *auth.Subject) bool {
	if subject == nil {
		return false
	}

	return subject.HasRole(domain.RoleAdmin)
}
