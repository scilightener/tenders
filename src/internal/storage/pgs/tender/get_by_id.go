package tender

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"tenders-management/internal/model/domain/tender"
)

func (r Repo) GetByID(ctx context.Context, tenderID uuid.UUID) (*tender.Tender, error) {
	const comp = "storage.pgs.tender.GetByID"

	tndrRow := r.db.DBPool.QueryRow(ctx,
		`SELECT * FROM tender WHERE id = $1;`,
		tenderID,
	)
	var trow tenderRow
	err := tndrRow.Scan(
		&trow.id, &trow.name, &trow.description, &trow.status, &trow.serviceType, &trow.organizationID,
		&trow.creatorID, &trow.version, &trow.createdAt, &trow.updatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, tender.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return trow.toModel()
}
