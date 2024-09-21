package organization

import (
	"errors"

	"tenders-management/internal/model/domain"
)

var (
	ErrNameTooLong = errors.New("name is too long")
)

var (
	ErrInvalidOrganization  = errors.New("invalid organization")
	ErrInvalidUser          = errors.New("invalid user")
	ErrOrganizationNotFound = errors.New("organization not found")
)

var (
	ErrResponsibleNotFound = errors.New("responsible not found")
)

func ErrInvalidOrganizationType(serviceType string) error {
	return domain.ValidationErr("invalid service type: " + serviceType)
}
