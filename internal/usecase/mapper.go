package usecase

import (
	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

func ToSubjectRoles(roles []*domain.Role) []string {
	res := make([]string, len(roles), len(roles))
	for index, role := range roles {
		res[index] = role.Name
	}

	return res
}

func ToSubject(user *domain.User, metadata map[string]string) *auth.Subject {
	return auth.NewSubject(user.ID, user.Name, auth.SubjectUser, ToSubjectRoles(user.Roles), metadata)
}
