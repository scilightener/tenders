package bidv

import (
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/bid"
	"tenders-management/internal/model/domain/employee"
	"tenders-management/internal/model/domain/organization"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/storage/pgs"
)

type bidvRow struct {
	bidID          uuid.UUID
	version        int
	name           string
	description    string
	status         string
	authorType     string
	organizationID *uuid.UUID
	userID         *uuid.UUID
	tenderID       uuid.UUID
	createdAt      time.Time
	updatedAt      time.Time
}

func (bvrow bidvRow) toModel() (*bid.Versioned, error) {
	status, err := bid.StatusFromString(bvrow.status)
	if err != nil {
		return nil, err
	}
	svcType, err := bid.AuthorTypeFromString(bvrow.authorType)
	if err != nil {
		return nil, err
	}

	org, usr := new(organization.Organization), new(employee.Employee)
	if bvrow.organizationID != nil {
		org = organization.EmptyOrganizationWithID(*bvrow.organizationID)
	} else if bvrow.userID != nil {
		usr = employee.EmptyWithID(*bvrow.userID)
	}

	b, err := bid.New(bvrow.bidID, bvrow.name, bvrow.description, status, svcType,
		org, usr, bvrow.version, tender.EmptyWithID(bvrow.tenderID), time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	return bid.NewVersionedFromBid(b), nil
}

type Repo struct {
	db pgs.Storage
}

func NewRepo(db pgs.Storage) Repo {
	return Repo{db: db}
}
