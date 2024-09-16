package bid

import "github.com/google/uuid"

func EmptyWithID(id uuid.UUID) *Bid {
	return &Bid{
		id: id,
	}
}
