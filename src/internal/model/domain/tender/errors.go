package tender

import (
	"errors"
	"strconv"

	"tenders-management/internal/model/domain"
)

var (
	ErrInvalidOrganization     = errors.New("invalid organization")
	ErrInvalidCreator          = errors.New("invalid creator")
	ErrNameTooLong             = errors.New("name is too long")
	ErrDescriptionTooLong      = errors.New("description is too long")
	ErrInvalidStatusTransition = errors.New("invalid status transition")

	ErrNotFound = errors.New("tender not found")
)

func ErrInvalidStatus(status string) error {
	return domain.ValidationErr("invalid status: " + status)
}

func ErrInvalidServiceType(serviceType string) error {
	return domain.ValidationErr("invalid service type: " + serviceType)
}

func ErrInvalidVersion(version int) error {
	return domain.ValidationErr("invalid version: " + strconv.Itoa(version))
}
