package bid

import (
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/employee"
	"tenders-management/internal/model/domain/organization"
	"tenders-management/internal/model/domain/tender"
)

type Versioned struct {
	BidID          uuid.UUID
	Version        int
	Name           string
	Description    string
	Status         status
	AuthorType     authorType
	OrganizationID *uuid.UUID
	UserID         *uuid.UUID
	TenderID       uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewVersioned(
	bidID uuid.UUID, version int, name, description string,
	status status, authorType authorType,
	organizationID, userID *uuid.UUID, tenderID uuid.UUID, createdAt, updatedAt time.Time,
) *Versioned {
	return &Versioned{
		BidID:          bidID,
		Version:        version,
		Name:           name,
		Description:    description,
		Status:         status,
		AuthorType:     authorType,
		OrganizationID: organizationID,
		UserID:         userID,
		TenderID:       tenderID,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}

func NewVersionedFromBid(b *Bid) *Versioned {
	var orgID, userID *uuid.UUID

	if b.AuthorType() == string(AuthorTypeOrganization) {
		orgID = new(uuid.UUID)
		*orgID = b.Organization().ID()
	} else if b.AuthorType() == string(AuthorTypeUser) {
		userID = new(uuid.UUID)
		*userID = b.User().ID()
	}

	return NewVersioned(b.ID(), b.Version(), b.Name(), b.Description(),
		status(b.Status()), authorType(b.AuthorType()), orgID, userID, b.Tender().ID(),
		b.CreatedAt(), b.UpdatedAt())
}

func (v *Versioned) ToBid() (*Bid, error) {
	org, usr := new(organization.Organization), new(employee.Employee)

	if v.OrganizationID != nil {
		org = organization.EmptyOrganizationWithID(*v.OrganizationID)
		usr = nil
	} else if v.UserID != nil {
		usr = employee.EmptyWithID(*v.UserID)
		org = nil
	}

	return New(v.BidID, v.Name, v.Description, v.Status, v.AuthorType, org, usr,
		v.Version, tender.EmptyWithID(v.TenderID), v.CreatedAt, v.UpdatedAt)
}
