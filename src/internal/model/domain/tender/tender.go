package tender

import (
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/organization"
)

type Tender struct {
	id           uuid.UUID
	name         string
	description  string
	status       status
	serviceType  serviceType
	version      int
	organization *organization.Organization
	creator      *organization.Responsible
	createdAt    time.Time
	updatedAt    time.Time
}

func New(
	id uuid.UUID,
	name, description string,
	status status,
	serviceType serviceType,
	version int,
	organization *organization.Organization,
	creator *organization.Responsible,
	createdAt, updatedAt time.Time,
) (*Tender, error) {
	t := &Tender{
		id:           id,
		name:         name,
		description:  description,
		status:       status,
		serviceType:  serviceType,
		version:      version,
		organization: organization,
		creator:      creator,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
	if err := t.validate(); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Tender) validate() error {
	if t.organization == nil {
		return ErrInvalidOrganization
	}
	if t.creator == nil {
		return ErrInvalidCreator
	}
	if len(t.name) > 100 {
		return ErrNameTooLong
	}
	if len(t.description) > 500 {
		return ErrDescriptionTooLong
	}
	if t.version < 1 {
		return ErrInvalidVersion(t.version)
	}
	return nil
}

func (t *Tender) ID() uuid.UUID {
	return t.id
}

func (t *Tender) SetID(id uuid.UUID) {
	t.id = id
}

func (t *Tender) Name() string {
	return t.name
}

func (t *Tender) SetName(name string) error {
	if len(t.name) > 100 {
		return ErrNameTooLong
	}
	t.name = name
	return nil
}

func (t *Tender) Description() string {
	return t.description
}

func (t *Tender) SetDescription(description string) error {
	if len(t.description) > 500 {
		return ErrDescriptionTooLong
	}
	t.description = description
	return nil
}

func (t *Tender) Status() string {
	return string(t.status)
}

func (t *Tender) TransitToStatus(s status) error {
	if !t.status.canTransitTo(s) {
		return ErrInvalidStatusTransition
	}
	t.status = s
	return nil
}

func (t *Tender) ServiceType() string {
	return string(t.serviceType)
}

func (t *Tender) SetServiceType(sType serviceType) {
	t.serviceType = sType
}

func (t *Tender) Version() int {
	return t.version
}

func (t *Tender) IncrementVersion() {
	t.version++
}

func (t *Tender) Organization() *organization.Organization {
	return t.organization
}

func (t *Tender) Creator() *organization.Responsible {
	return t.creator
}

func (t *Tender) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Tender) UpdatedAt() time.Time {
	return t.updatedAt
}
