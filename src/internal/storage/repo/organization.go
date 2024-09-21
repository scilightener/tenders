package repo

import (
	"context"
	"tenders-management/internal/model/domain/organization"

	"github.com/google/uuid"
)

type Organization interface {
	GetByID(ctx context.Context, id uuid.UUID) (*organization.Organization, error)
}
