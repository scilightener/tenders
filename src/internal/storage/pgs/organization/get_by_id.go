package organization

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"tenders-management/internal/model/domain/organization"
)

func (r Repo) GetByID(ctx context.Context, id uuid.UUID) (*organization.Organization, error) {
	const comp = "storage.pgs.organization.GetByID"
	row := r.db.DBPool.QueryRow(ctx, `
		SELECT id, name, description, type, created_at, updated_at
		FROM organization
		WHERE id = $1
	`, id)

	var orow organizationRow
	err := row.Scan(&orow.id, &orow.name, &orow.description, &orow.organizationType, &orow.createdAt, &orow.updatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, organization.ErrOrganizationNotFound
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return orow.toModel()
}
