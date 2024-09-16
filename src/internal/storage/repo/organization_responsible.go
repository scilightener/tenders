package repo

import (
	"context"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/organization"
)

type OrganizationResponsible interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*organization.Responsible, error)
}
