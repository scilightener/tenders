package bid

import (
	"tenders-management/internal/model/domain/employee"
	"tenders-management/internal/model/domain/organization"
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/tender"
)

type Bid struct {
	id           uuid.UUID
	name         string
	description  string
	status       status
	authorType   authorType
	organization *organization.Organization
	user         *employee.Employee
	version      int
	tender       *tender.Tender
	createdAt    time.Time
	updatedAt    time.Time
}

func New(id uuid.UUID,
	name string,
	description string,
	status status,
	authorType authorType,
	organization *organization.Organization,
	user *employee.Employee,
	version int,
	tender *tender.Tender,
	createdAt, updatedAt time.Time,
) (*Bid, error) {
	b := &Bid{
		id:           id,
		name:         name,
		description:  description,
		status:       status,
		authorType:   authorType,
		organization: organization,
		user:         user,
		version:      version,
		tender:       tender,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
	if err := b.validate(); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Bid) validate() error {
	if b.tender == nil {
		return ErrInvalidTender
	}
	if len(b.name) > 100 {
		return ErrNameTooLong
	}
	if len(b.description) > 500 {
		return ErrDescriptionTooLong
	}
	if b.version < 1 {
		return ErrInvalidVersion(b.version)
	}
	if b.authorType == AuthorTypeUser && b.user == nil ||
		b.authorType == AuthorTypeOrganization && b.organization == nil {
		return ErrAuthorTypeError(b.AuthorType())
	}
	return nil
}

func (b *Bid) ID() uuid.UUID {
	return b.id
}

func (b *Bid) Name() string {
	return b.name
}

func (b *Bid) SetName(name string) error {
	if len(name) > 100 {
		return ErrNameTooLong
	}
	b.name = name
	return nil
}

func (b *Bid) Description() string {
	return b.description
}

func (b *Bid) SetDescription(description string) error {
	if len(description) > 500 {
		return ErrDescriptionTooLong
	}
	b.description = description
	return nil
}

func (b *Bid) AuthorType() string {
	return string(b.authorType)
}

func (b *Bid) Status() string {
	return string(b.status)
}

func (b *Bid) TransitToStatus(s status) error {
	if !b.status.canTransitTo(s) {
		return ErrInvalidStatusTransition
	}
	b.status = s
	return nil
}

func (b *Bid) Organization() *organization.Organization {
	return b.organization
}

func (b *Bid) OrganizationID() *uuid.UUID {
	if b.organization == nil {
		return nil
	}
	id := new(uuid.UUID)
	*id = b.organization.ID()
	return id
}

func (b *Bid) User() *employee.Employee {
	return b.user
}

func (b *Bid) UserID() *uuid.UUID {
	if b.user == nil {
		return nil
	}
	id := new(uuid.UUID)
	*id = b.user.ID()
	return id
}

func (b *Bid) Version() int {
	return b.version
}

func (b *Bid) IncrementVersion() {
	b.version++
}

func (b *Bid) Tender() *tender.Tender {
	return b.tender
}

func (b *Bid) CreatedAt() time.Time {
	return b.createdAt
}

func (b *Bid) UpdatedAt() time.Time {
	return b.updatedAt
}
