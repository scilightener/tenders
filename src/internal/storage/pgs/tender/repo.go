package tender

import (
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/organization"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/storage/pgs"
)

type tenderRow struct {
	id             uuid.UUID
	name           string
	description    string
	status         string
	serviceType    string
	organizationID uuid.UUID
	creatorID      uuid.UUID
	version        int
	createdAt      time.Time
	updatedAt      time.Time
}

func (trow tenderRow) toModel() (*tender.Tender, error) {
	status, err := tender.StatusFromString(trow.status)
	if err != nil {
		return nil, err
	}
	svcType, err := tender.ServiceTypeFromString(trow.serviceType)
	if err != nil {
		return nil, err
	}
	return tender.New(trow.id, trow.name, trow.description, status, svcType,
		trow.version, organization.EmptyOrganizationWithID(trow.organizationID),
		organization.EmptyResponsibleWithID(trow.creatorID), trow.createdAt, trow.updatedAt)
}

type Repo struct {
	db pgs.Storage
}

func NewRepo(db pgs.Storage) Repo {
	return Repo{db: db}
}
