package dto

import (
	"strings"
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/model/domain/tender"
)

type Tender struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ServiceType    string    `json:"serviceType"`
	Status         string    `json:"status"`
	OrganizationID uuid.UUID `json:"organizationId"`
	Version        int       `json:"version"`
	CreatedAt      string    `json:"createdAt"`
}

func TenderFromModel(t *tender.Tender) *Tender {
	return &Tender{
		ID:             t.ID(),
		Name:           t.Name(),
		Description:    t.Description(),
		ServiceType:    toPascalCase(t.ServiceType()),
		Status:         toPascalCase(t.Status()),
		OrganizationID: t.Organization().ID(),
		Version:        t.Version(),
		CreatedAt:      t.CreatedAt().Format(time.RFC3339),
	}
}

func toPascalCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
}

type NewTender struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ServiceType     string    `json:"serviceType"`
	OrganizationID  uuid.UUID `json:"organizationID"`
	CreatorUsername string    `json:"creatorUsername"`
}

type UpdateTender struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	ServiceType *string `json:"serviceType"`
}
