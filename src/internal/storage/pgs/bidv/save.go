package bidv

import (
	"context"
	"fmt"

	"tenders-management/internal/model/domain/bid"
)

func (r Repo) Save(ctx context.Context, b *bid.Versioned) (*bid.Versioned, error) {
	const comp = "storage.pgs.bidv.Save"

	row := r.db.DBPool.QueryRow(ctx,
		`INSERT INTO bid_versions (bid_id, version, name, description, status, author_type,
				 organization_id, user_id, tender_id, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING *;`,
		b.BidID, b.Version, b.Name, b.Description, string(b.Status), string(b.AuthorType),
		b.OrganizationID, b.UserID, b.TenderID, b.CreatedAt, b.UpdatedAt)

	var brow bidvRow
	err := row.Scan(
		&brow.bidID, &brow.version, &brow.name, &brow.description, &brow.status, &brow.authorType,
		&brow.organizationID, &brow.userID, &brow.tenderID, &brow.createdAt, &brow.updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return brow.toModel()
}
