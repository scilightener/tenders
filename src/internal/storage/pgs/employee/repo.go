package employee

import (
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/storage/pgs"
)

type employeeRow struct {
	id        uuid.UUID
	username  string
	firstName *string
	lastName  *string
	createdAt time.Time
	updatedAt time.Time
}

type Repo struct {
	db pgs.Storage
}

func NewRepo(db pgs.Storage) Repo {
	return Repo{db: db}
}
