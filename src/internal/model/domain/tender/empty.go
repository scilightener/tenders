package tender

import "github.com/google/uuid"

func EmptyWithID(id uuid.UUID) *Tender {
	return &Tender{
		id: id,
	}
}
