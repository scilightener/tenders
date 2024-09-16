package bid

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/bid"
)

func (r Repo) GetByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*bid.Bid, error) {
	const comp = "storage.pgs.bid.GetByUserID"

	rows, err := r.db.DBPool.Query(ctx,
		"SELECT * FROM bid WHERE user_id = $1 ORDER BY name LIMIT $2 OFFSET $3;",
		userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}
	defer rows.Close()

	//if !rows.Next() {
	//	return nil, bid.ErrNotFound
	//}

	var bs []*bid.Bid
	for rows.Next() {
		var brow bidRow
		if err := rows.Scan(
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
		); err != nil {
			return nil, fmt.Errorf("%s: %w", comp, err)
		}

		b, err := brow.toModel()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", comp, err)
		}

		bs = append(bs, b)
	}

	return bs, nil
}
