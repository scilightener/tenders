package responsible

import (
	"github.com/google/uuid"

	"tenders-management/internal/storage/pgs"
)

type responsibleRow struct {
	id             uuid.UUID
	organizationID uuid.UUID
	userID         uuid.UUID
}

type Repo struct {
	db pgs.Storage
}

func NewRepo(db pgs.Storage) Repo {
	return Repo{db: db}
}
