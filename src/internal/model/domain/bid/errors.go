package bid

import (
	"errors"
	"strconv"

	"tenders-management/internal/model/domain"
)

var (
	ErrNameTooLong             = errors.New("name is too long")
	ErrDescriptionTooLong      = errors.New("description is too long")
	ErrInvalidTender           = errors.New("invalid tender")
	ErrInvalidStatusTransition = errors.New("invalid status transition")
	ErrNotFound                = errors.New("bid not found")
)

func ErrInvalidAuthorType(s string) error {
	return domain.ValidationErr("invalid author type: " + s)
}

func ErrAuthorTypeError(authorType string) error {
	return domain.ValidationErr(
		"author type is " + authorType + " but corresponding user or organization is not set or found")
}

func ErrInvalidVersion(version int) error {
	return domain.ValidationErr("invalid version: " + strconv.Itoa(version))
}

func ErrInvalidStatus(status string) error {
	return domain.ValidationErr("invalid status: " + status)
}
