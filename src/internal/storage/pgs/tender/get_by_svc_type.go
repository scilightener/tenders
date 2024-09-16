package tender

import (
	"context"
	"fmt"

	"tenders-management/internal/model/domain/tender"
)

func (r Repo) GetPublishedBySvcType(
	ctx context.Context, limit, offset int, svcTypes []string,
) ([]*tender.Tender, error) {
	const comp = "storage.pgs.tender.GetPublishedBySvcType"

	query, args := getPublishedBySvcTypeQueryAndArgs(limit, offset, svcTypes)
	tenders, err := r.db.DBPool.Query(ctx,
		query, args...,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	if !tenders.Next() {
		return nil, tender.ErrNotFound
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

func getPublishedBySvcTypeQueryAndArgs(limit, offset int, svcTypes []string) (string, []any) {
	var (
		query string
		args  = make([]interface{}, 0, 3)
	)
	if len(svcTypes) != 0 {
		query = `SELECT * FROM tender WHERE status = 'PUBLISHED' AND service_type = ANY($1) ORDER BY name LIMIT $2 OFFSET $3;`
		args = append(args, svcTypes)
	} else {
		query = `SELECT * FROM tender WHERE status = 'PUBLISHED' ORDER BY name LIMIT $1 OFFSET $2;`
	}

	args = append(args, limit, offset)
	return query, args
}
