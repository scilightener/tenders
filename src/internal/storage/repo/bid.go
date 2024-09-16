package repo

import (
	"context"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/bid"
)

type Bid interface {
	Save(ctx context.Context, bid *bid.Bid) (*bid.Bid, error)
	GetPublishedByTenderID(ctx context.Context, tenderID uuid.UUID, limit, offset int) ([]*bid.Bid, error)
	GetByID(ctx context.Context, id uuid.UUID) (*bid.Bid, error)
	GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*bid.Bid, error)
	GetByOrganizationID(ctx context.Context, orgID uuid.UUID, limit, offset int) ([]*bid.Bid, error)
	Update(ctx context.Context, bidID uuid.UUID, bid *bid.Bid) (*bid.Bid, error)
}

type BidVersions interface {
	Save(ctx context.Context, bid *bid.Versioned) (*bid.Versioned, error)
	GetByBidIDVersion(ctx context.Context, bidID uuid.UUID, version int) (*bid.Versioned, error)
}
