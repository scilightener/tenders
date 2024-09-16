package bid

import (
	"context"
	"fmt"
	"github.com/google/uuid"

	"tenders-management/internal/model/domain/bid"
)

func (r Repo) Save(ctx context.Context, b *bid.Bid) (*bid.Bid, error) {
	const comp = "storage.pgs.bid.Save"

	var organizationID, userID *uuid.UUID
	if b.AuthorType() == string(bid.AuthorTypeUser) {
		userID = new(uuid.UUID)
		*userID = b.User().ID()
	} else if b.AuthorType() == string(bid.AuthorTypeOrganization) {
		organizationID = new(uuid.UUID)
		*organizationID = b.Organization().ID()
	}

	row := r.db.DBPool.QueryRow(ctx,
		`INSERT INTO bid (name, description, status, author_type,
                 organization_id, user_id, version, tender_id, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *;`,
		b.Name(), b.Description(), b.Status(), b.AuthorType(),
		organizationID, userID, b.Version(),
		b.Tender().ID().String(), b.CreatedAt(), b.UpdatedAt())

	var brow bidRow
	err := row.Scan(
		&brow.id, &brow.name, &brow.description, &brow.status, &brow.authorType,
		&brow.organizationID, &brow.userID, &brow.version, &brow.tenderID, &brow.createdAt, &brow.updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return brow.toModel()
}
