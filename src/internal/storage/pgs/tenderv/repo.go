package tenderv

import (
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/organization"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/storage/pgs"
)

type tenderVRow struct {
	tenderID       uuid.UUID
	version        int
	name           string
	description    string
	status         string
	serviceType    string
	organizationID uuid.UUID
	creatorID      uuid.UUID
	createdAt      time.Time
	updatedAt      time.Time
}

func (trow tenderVRow) toModel() (*tender.Versioned, error) {
	status, err := tender.StatusFromString(trow.status)
	if err != nil {
		return nil, err
	}
	svcType, err := tender.ServiceTypeFromString(trow.serviceType)
	if err != nil {
		return nil, err
	}
	t, err := tender.New(trow.tenderID, trow.name, trow.description, status, svcType,
		trow.version, organization.EmptyOrganizationWithID(trow.organizationID),
		organization.EmptyResponsibleWithID(trow.creatorID), trow.createdAt, trow.updatedAt)
	if err != nil {
		return nil, err
	}
	return tender.NewVersionedFromTender(t), nil
}

type Repo struct {
	db pgs.Storage
}

func NewRepo(db pgs.Storage) Repo {
	return Repo{db: db}
}
