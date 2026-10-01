package transport

import (
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/errs"
)

// IsBadRequest evaluates if the inbound error unwraps into any validation, schema mapping, or invalid criteria types.
func IsBadRequest(err error) bool {
	var (
		errInvalidArgument *errs.InvalidArgumentError
		errBllValidate     *errs.BllValidateError
		errTrMapping       *errs.TlMappingError
	)

	return errors.As(err, &errInvalidArgument) ||
		errors.As(err, &errBllValidate) ||
		errors.As(err, &errTrMapping)
}

// IsUnauthorized checks if the root or cascaded error identity maps directly to an unauthenticated session state.
func IsUnauthorized(err error) bool {
	var (
		errBllUnauthorized *errs.BllUnauthorizedError
	)

	return errors.As(err, &errBllUnauthorized)
}

// IsForbidden checks if the structural error tree represents access control or permission privilege enforcement rejections.
func IsForbidden(err error) bool {
	var (
		errBllForbidden *errs.BllForbiddenError
	)

	return errors.As(err, &errBllForbidden)
}

// IsNotFound determines whether the error chain indicates a missing application aggregate or storage entry node.
func IsNotFound(err error) bool {
	var (
		errBllNotFound *errs.BllNotFoundError
		errDalNotFound *errs.DalNotFoundError
	)

	return errors.As(err, &errBllNotFound) ||
		errors.As(err, &errDalNotFound)
}

// IsConflict tests if the error footprint targets unique constraints database violations or concurrency collision states.
func IsConflict(err error) bool {
	var (
		errBllUnique        *errs.BllUniqueError
		errDalAlreadyExists *errs.DalAlreadyExistsError
	)

	return errors.As(err, &errBllUnique) ||
		errors.As(err, &errDalAlreadyExists)
}

// IsGone verifies if the structural error payload flags an asset that was previously archived or softly deleted.
func IsGone(err error) bool {
	var (
		errDalSoftDeleted *errs.DalSoftDeletedError
	)

	return errors.As(err, &errDalSoftDeleted)
}
