package repo

import (
	"context"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/tender"
)

type Tender interface {
	Save(ctx context.Context, tender *tender.Tender) (*tender.Tender, error)
	GetPublishedBySvcType(ctx context.Context, limit, offset int, svcTypes []string) ([]*tender.Tender, error)
	GetByCreatorID(ctx context.Context, creatorID uuid.UUID, limit, offset int) ([]*tender.Tender, error)
	GetByID(ctx context.Context, tenderID uuid.UUID) (*tender.Tender, error)
	Update(ctx context.Context, tenderID uuid.UUID, tender *tender.Tender) (*tender.Tender, error)
}

type TenderVersions interface {
	Save(ctx context.Context, tender *tender.Versioned) (*tender.Versioned, error)
	GetByTenderIDVersion(ctx context.Context, tenderID uuid.UUID, version int) (*tender.Versioned, error)
}
