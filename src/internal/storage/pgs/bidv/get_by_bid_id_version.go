package bidv

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/bid"
)

func (r Repo) GetByBidIDVersion(ctx context.Context, bidID uuid.UUID, version int) (*bid.Versioned, error) {
	const comp = "storage.pgs.bidv.GetByBidIDVersion"
	row := r.db.DBPool.QueryRow(ctx, `SELECT * FROM bid_versions WHERE bid_id = $1 AND version = $2`,
		bidID, version)

	var brow bidvRow
	err := row.Scan(
		&brow.bidID, &brow.version, &brow.name, &brow.description, &brow.status, &brow.authorType,
		&brow.organizationID, &brow.userID, &brow.tenderID, &brow.createdAt, &brow.updatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}
	return brow.toModel()
}
