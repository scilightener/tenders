package tender

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/tender"
)

func (r Repo) GetByCreatorID(
	ctx context.Context, creatorID uuid.UUID, limit, offset int,
) ([]*tender.Tender, error) {
	const comp = "storage.pgs.tender.GetByCreatorID"

	tenders, err := r.db.DBPool.Query(ctx,
		`SELECT * FROM tender WHERE creator_id = $1 ORDER BY name LIMIT $2 OFFSET $3;`,
		creatorID, limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	var ts []*tender.Tender
	for tenders.Next() {
		var trow tenderRow
		err := tenders.Scan(
			&trow.id, &trow.name, &trow.description, &trow.status, &trow.serviceType, &trow.organizationID,
			&trow.creatorID, &trow.version, &trow.createdAt, &trow.updatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", comp, err)
		}

		t, err := trow.toModel()
		if err != nil {
			return nil, err
		}

		ts = append(ts, t)
	}

	return ts, nil
}
