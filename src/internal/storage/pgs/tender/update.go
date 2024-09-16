package tender

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/tender"
)

func (r Repo) Update(ctx context.Context, tenderID uuid.UUID, tender *tender.Tender) (*tender.Tender, error) {
	const comp = "storage.pgs.tender.Update"

	row := r.db.DBPool.QueryRow(ctx,
		`UPDATE tender SET name = $1, description = $2, status = $3, service_type = $4,
                  organization_id = $5, creator_id = $6, version = $7, created_at = $8, updated_at = $9
              WHERE id = $10 RETURNING *;`,
		tender.Name(), tender.Description(), tender.Status(), tender.ServiceType(),
		tender.Organization().ID().String(), tender.Creator().ID().String(), tender.Version(),
		tender.CreatedAt(), tender.UpdatedAt(), tenderID)

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
