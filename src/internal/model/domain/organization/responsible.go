package organization

import (
	"github.com/google/uuid"

	"tenders-management/internal/model/domain/employee"
)

type Responsible struct {
	id           uuid.UUID
	organization *Organization
	user         *employee.Employee
}

func NewResponsible(id uuid.UUID, organization *Organization, user *employee.Employee) (*Responsible, error) {
	if organization == nil {
		return nil, ErrInvalidOrganization
	} else if user == nil {
		return nil, ErrInvalidUser
	}
	return &Responsible{
		id:           id,
		organization: organization,
		user:         user,
	}, nil
}

func (r *Responsible) ID() uuid.UUID {
	return r.id
}

func (r *Responsible) Organization() *Organization {
	return r.organization
}

func (r *Responsible) User() *employee.Employee {
	return r.user
}
