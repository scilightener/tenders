package bid

import (
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/bid"
	"tenders-management/internal/model/domain/employee"
	"tenders-management/internal/model/domain/organization"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/storage/pgs"
)

type bidRow struct {
	id             uuid.UUID
	name           string
	description    string
	status         string
	authorType     string
	organizationID *uuid.UUID
	userID         *uuid.UUID
	version        int
	tenderID       uuid.UUID
	createdAt      time.Time
	updatedAt      time.Time
}

func (brow bidRow) toModel() (*bid.Bid, error) {
	status, err := bid.StatusFromString(brow.status)
	if err != nil {
		return nil, err
	}
	authorType, err := bid.AuthorTypeFromString(brow.authorType)
	if err != nil {
		return nil, err
	}
	var (
		org *organization.Organization
		usr *employee.Employee
	)

	if authorType == bid.AuthorTypeOrganization {
		org = organization.EmptyOrganizationWithID(*brow.organizationID)
	} else if authorType == bid.AuthorTypeUser {
		usr = employee.EmptyWithID(*brow.userID)
	}

	return bid.New(brow.id, brow.name, brow.description, status, authorType,
		org, usr, brow.version, tender.EmptyWithID(brow.tenderID), brow.createdAt, brow.updatedAt)
}

type Repo struct {
	db pgs.Storage
}

func NewRepo(db pgs.Storage) Repo {
	return Repo{db: db}
}
