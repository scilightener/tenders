package organization

import (
	"time"

	"github.com/google/uuid"
)

// Organization represents an organization (legal entity).
type Organization struct {
	id               uuid.UUID
	name             string
	description      string
	organizationType _type
	createdAt        time.Time
	updatedAt        time.Time
}

// NewOrganization creates a new instance of the Organization.
func NewOrganization(
	id uuid.UUID, name, description string, orgType _type,
	createdAt, updatedAt time.Time,
) (*Organization, error) {
	t := &Organization{
		id:               id,
		name:             name,
		description:      description,
		organizationType: orgType,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
	}
	if err := t.validate(); err != nil {
		return nil, err
	}

	return t, nil
}

func (o *Organization) validate() error {
	if len(o.name) > 100 {
		return ErrNameTooLong
	}
	return nil
}

// ID returns the organization's ID.
func (o *Organization) ID() uuid.UUID {
	return o.id
}

// Name returns the organization's name.
func (o *Organization) Name() string {
	return o.name
}

// Description returns the organization's description.
func (o *Organization) Description() string {
	return o.description
}

// Type returns the organization's type.
func (o *Organization) Type() string {
	return string(o.organizationType)
}

func (o *Organization) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Organization) UpdatedAt() time.Time {
	return o.updatedAt
}
