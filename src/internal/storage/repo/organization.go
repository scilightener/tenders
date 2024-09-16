package repo

import (
	"context"
	"github.com/google/uuid"
	"tenders-management/internal/model/domain/organization"
)

type Organization interface {
	GetByID(ctx context.Context, id uuid.UUID) (*organization.Organization, error)
}
