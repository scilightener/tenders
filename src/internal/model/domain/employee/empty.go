package employee

import "github.com/google/uuid"

func EmptyWithID(id uuid.UUID) *Employee {
	return &Employee{
		id: id,
	}
}
