package usecase

import (
	"github.com/ElfAstAhe/go-service-template/pkg/auth"
	"github.com/ElfAstAhe/tiny-auth-service/internal/domain"
)

// ToSubjectRoles maps a slice of domain Role entities into a flat slice of string names.
func ToSubjectRoles(roles []*domain.Role) []string {
	res := make([]string, 0, len(roles))
	for _, role := range roles {
		res = append(res, role.Name)
	}

	return res
}

// ToSubject encapsulates transformation mechanics converting a domain User aggregate and contextual metadata into an auth.Subject identity contract.
func ToSubject(user *domain.User, metadata map[string]string) *auth.Subject {
	return auth.NewSubject(user.ID, user.Name, ToSubjectType(user.Type), ToSubjectRoles(user.Roles), metadata)
}

// ToSubjectType evaluates a raw domain user type string criterion and maps it into a strongly-typed auth.SubjectType system constant.
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
