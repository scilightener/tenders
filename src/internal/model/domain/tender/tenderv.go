package tender

import (
	"tenders-management/internal/model/domain/organization"
	"time"

	"github.com/google/uuid"
)

type Versioned struct {
	TenderID       uuid.UUID
	Version        int
	Name           string
	Description    string
	Status         status
	ServiceType    serviceType
	OrganizationID uuid.UUID
	CreatorID      uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewVersioned(
	tenderID uuid.UUID, version int, name, description string,
	status status, serviceType serviceType, organizationID,
	creatorID uuid.UUID, createdAt, updatedAt time.Time,
) *Versioned {
	return &Versioned{
		TenderID:       tenderID,
		Version:        version,
		Name:           name,
		Description:    description,
		Status:         status,
		ServiceType:    serviceType,
		OrganizationID: organizationID,
		CreatorID:      creatorID,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}

func NewVersionedFromTender(t *Tender) *Versioned {
	return NewVersioned(t.ID(), t.Version(), t.Name(), t.Description(),
		status(t.Status()), serviceType(t.ServiceType()), t.Organization().ID(),
		t.Creator().ID(), t.CreatedAt(), t.UpdatedAt())
}

func (v *Versioned) ToTender() (*Tender, error) {
	return New(v.TenderID, v.Name, v.Description, v.Status,
		v.ServiceType, v.Version, organization.EmptyOrganizationWithID(v.OrganizationID),
		organization.EmptyResponsibleWithID(v.CreatorID), v.CreatedAt, v.UpdatedAt)
}
