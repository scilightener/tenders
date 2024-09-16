package tenderv

import (
	"context"
	"fmt"

	"tenders-management/internal/model/domain/tender"
)

func (r Repo) Save(ctx context.Context, t *tender.Versioned) (*tender.Versioned, error) {
	const comp = "storage.pgs.tenderv.Save"

	row := r.db.DBPool.QueryRow(ctx,
		`INSERT INTO tender_versions (tender_id, version, name, description, status, service_type,
				 organization_id, creator_id, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *;`,
		t.TenderID, t.Version, t.Name, t.Description, string(t.Status), string(t.ServiceType), t.OrganizationID,
		t.CreatorID, t.CreatedAt, t.UpdatedAt)

	var trow tenderVRow
	err := row.Scan(
		&trow.tenderID, &trow.version, &trow.name, &trow.description, &trow.status, &trow.serviceType,
		&trow.organizationID, &trow.creatorID, &trow.createdAt, &trow.updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return trow.toModel()
}
