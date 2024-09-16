package responsible

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/employee"
	"tenders-management/internal/model/domain/organization"
)

func (r Repo) GetByUserID(ctx context.Context, userID uuid.UUID) (*organization.Responsible, error) {
	const comp = "storage.pgs.responsible.GetByUserID"

	rows, err := r.db.DBPool.Query(ctx, "SELECT * FROM organization_responsible WHERE user_id = $1", userID.String())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, organization.ErrResponsibleNotFound
	}

	var rrow responsibleRow
	if err := rows.Scan(
		&rrow.id,
		&rrow.organizationID,
		&rrow.userID,
	); err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return organization.NewResponsible(rrow.id,
		organization.EmptyOrganizationWithID(rrow.organizationID),
		employee.EmptyWithID(rrow.userID))
}
