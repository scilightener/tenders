package organization

import "github.com/google/uuid"

func EmptyOrganizationWithID(id uuid.UUID) *Organization {
	return &Organization{
		id: id,
	}
}

func EmptyResponsibleWithID(id uuid.UUID) *Responsible {
	return &Responsible{
		id: id,
	}
}
