package organization

import (
	"tenders-management/internal/model/domain/organization"
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/storage/pgs"
)

type organizationRow struct {
	id               uuid.UUID
	name             string
	description      string
	organizationType string
	createdAt        time.Time
	updatedAt        time.Time
}

func (or organizationRow) toModel() (*organization.Organization, error) {
	orgType, err := organization.TypeFromString(or.organizationType)
	if err != nil {
		return nil, err
	}

	return organization.NewOrganization(or.id, or.name, or.description, orgType, or.createdAt, or.updatedAt)
}

type Repo struct {
	db pgs.Storage
}

func NewRepo(db pgs.Storage) Repo {
	return Repo{db: db}
}
