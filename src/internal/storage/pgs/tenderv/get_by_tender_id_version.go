package tenderv

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/tender"
)

func (r Repo) GetByTenderIDVersion(ctx context.Context, tenderID uuid.UUID, version int) (*tender.Versioned, error) {
	const comp = "storage.pgs.tenderv.GetByTenderIDVersion"

	row := r.db.DBPool.QueryRow(ctx, `SELECT * FROM tender_versions
         WHERE tender_id = $1 AND version = $2`, tenderID, version)

	var trow tenderVRow
	err := row.Scan(&trow.tenderID, &trow.version, &trow.name, &trow.description, &trow.status, &trow.serviceType,
		&trow.organizationID, &trow.creatorID, &trow.createdAt, &trow.updatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}
	return trow.toModel()
}
