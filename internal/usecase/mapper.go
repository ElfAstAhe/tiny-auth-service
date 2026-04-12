package usecase

import (
	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

func ToSubjectRoles(roles []*domain.Role) []string {
	res := make([]string, 0, len(roles))
	for _, role := range roles {
		res = append(res, role.Name)
	}

	return res
}

func ToSubject(user *domain.User, metadata map[string]string) *auth.Subject {
	return auth.NewSubject(user.ID, user.Name, ToSubjectType(user.Type), ToSubjectRoles(user.Roles), metadata)
}

func ToSubjectType(userType string) auth.SubjectType {
	switch userType {
	case domain.UserTypeUser:
		return auth.SubjectUser
	case domain.UserTypeService:
		return auth.SubjectService
	default:
		return auth.SubjectGuest
	}
}
