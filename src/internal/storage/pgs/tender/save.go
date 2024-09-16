package tender

import (
	"context"
	"fmt"

	"tenders-management/internal/model/domain/tender"
)

func (r Repo) Save(ctx context.Context, t *tender.Tender) (*tender.Tender, error) {
	const comp = "storage.pgs.tender.Save"

	row := r.db.DBPool.QueryRow(ctx,
		`INSERT INTO tender (name, description, status, service_type, organization_id,
				creator_id, version, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;`,
		t.Name(), t.Description(), t.Status(), t.ServiceType(), t.Organization().ID().String(),
		t.Creator().ID().String(), t.Version(), t.CreatedAt(), t.UpdatedAt())

	var trow tenderRow
	err := row.Scan(
		&trow.id, &trow.name, &trow.description, &trow.status, &trow.serviceType, &trow.organizationID,
		&trow.creatorID, &trow.version, &trow.createdAt, &trow.updatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return trow.toModel()
}
