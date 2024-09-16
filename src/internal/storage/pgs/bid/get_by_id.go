package bid

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/bid"
)

func (r Repo) GetByID(ctx context.Context, id uuid.UUID) (*bid.Bid, error) {
	const comp = "storage.pgs.bid.GetByID"

	row := r.db.DBPool.QueryRow(ctx,
		"SELECT * FROM bid WHERE id = $1;", id)

	var brow bidRow
	err := row.Scan(
		&brow.id,
		&brow.name,
		&brow.description,
		&brow.status,
		&brow.authorType,
		&brow.organizationID,
		&brow.userID,
		&brow.version,
		&brow.tenderID,
		&brow.createdAt,
		&brow.updatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, bid.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return brow.toModel()
}
