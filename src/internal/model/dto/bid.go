package dto

import (
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/bid"
)

type Bid struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	AuthorType  string    `json:"authorType"`
	AuthorID    uuid.UUID `json:"authorId"`
	Version     int       `json:"version"`
	CreatedAt   string    `json:"createdAt"`
}

func BidFromModel(b *bid.Bid) Bid {
	var authorID uuid.UUID
	if b.AuthorType() == string(bid.AuthorTypeOrganization) {
		authorID = b.Organization().ID()
	} else if b.AuthorType() == string(bid.AuthorTypeUser) {
		authorID = b.User().ID()
	}
	return Bid{
		ID:          b.ID(),
		Name:        b.Name(),
		Description: b.Description(),
		Status:      toPascalCase(b.Status()),
		AuthorType:  toPascalCase(b.AuthorType()),
		AuthorID:    authorID,
		Version:     b.Version(),
		CreatedAt:   b.CreatedAt().Format(time.RFC3339),
	}
}

type NewBid struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TenderID    uuid.UUID `json:"tenderId"`
	AuthorType  string    `json:"authorType"`
	AuthorID    uuid.UUID `json:"authorId"`
}

type UpdateBid struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
