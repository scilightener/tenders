package bid

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/bid"
)

func (r Repo) Update(ctx context.Context, bidID uuid.UUID, b *bid.Bid) (*bid.Bid, error) {
	const comp = "storage.pgs.bid.Update"

	row := r.db.DBPool.QueryRow(ctx,
		`UPDATE bid SET name = $1, description = $2, status = $3, author_type = $4,
                  organization_id = $5, user_id = $6, version = $7,
                  tender_id = $8, created_at = $9, updated_at = $10
              WHERE id = $11 RETURNING *;`,
		b.Name(), b.Description(), b.Status(), b.AuthorType(),
		b.OrganizationID(), b.UserID(), b.Version(),
		b.Tender().ID(), b.CreatedAt(), b.UpdatedAt(), bidID)

	var brow bidRow
	err := row.Scan(
		&brow.id, &brow.name, &brow.description, &brow.status, &brow.authorType,
		&brow.organizationID, &brow.userID, &brow.version,
		&brow.tenderID, &brow.createdAt, &brow.updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return brow.toModel()
}
