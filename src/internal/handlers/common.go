package handlers

import (
	"errors"
	"net/http"
	"tenders-management/internal/model/domain/bid"
	"tenders-management/internal/model/domain/employee"
	"tenders-management/internal/model/domain/organization"

	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/service"
)

func MapErrorToStatusCode(err error) int {
	if errors.Is(err, service.ErrUnknownError) {
		return http.StatusInternalServerError
	} else if errors.Is(err, employee.ErrNotFound) ||
		errors.Is(err, organization.ErrOrganizationNotFound) ||
		errors.Is(err, organization.ErrResponsibleNotFound) {
		return http.StatusBadRequest
	} else if errors.Is(err, service.ErrUnprivileged) {
		return http.StatusForbidden
	} else if errors.Is(err, tender.ErrNotFound) ||
		errors.Is(err, bid.ErrNotFound) {
		return http.StatusNotFound
	} else if err != nil {
		return http.StatusBadRequest
	}
	return http.StatusOK
}
