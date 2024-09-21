package repo

import (
	"context"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/employee"
)

type Employee interface {
	GetByUsername(ctx context.Context, username string) (*employee.Employee, error)
	GetByID(ctx context.Context, id uuid.UUID) (*employee.Employee, error)
}
