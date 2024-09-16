package employee

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"tenders-management/internal/model/domain/employee"
)

func (r Repo) GetByID(ctx context.Context, id uuid.UUID) (*employee.Employee, error) {
	const comp = "storage.pgs.employee.GetByID"

	rows, err := r.db.DBPool.Query(ctx, "SELECT * FROM employee WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, employee.ErrNotFound
	}

	var erow employeeRow
	if err := rows.Scan(
		&erow.id,
		&erow.username,
		&erow.firstName,
		&erow.lastName,
		&erow.createdAt,
		&erow.updatedAt,
	); err != nil {
		return nil, fmt.Errorf("%s: %w", comp, err)
	}

	return employee.New(erow.id, erow.username, erow.firstName, erow.lastName, erow.createdAt, erow.updatedAt)
}
