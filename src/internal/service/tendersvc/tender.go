package tendersvc

import (
	"context"
	"log/slog"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/storage/repo"
)

type Svc struct {
	tenderRepo  repo.Tender
	tenderVRepo repo.TenderVersions
	emplRepo    repo.Employee
	respRepo    repo.OrganizationResponsible
	logger      *slog.Logger
}

func NewTender(
	tenderRepo repo.Tender, tenderVRepo repo.TenderVersions, emplRepo repo.Employee,
	respRepo repo.OrganizationResponsible, logger *slog.Logger,
) *Svc {
	return &Svc{
		tenderRepo:  tenderRepo,
		tenderVRepo: tenderVRepo,
		emplRepo:    emplRepo,
		respRepo:    respRepo,
		logger:      logger,
	}
}

func mapTendersToDTOs(tenders []*tender.Tender) []dto.Tender {
	dtos := make([]dto.Tender, 0, len(tenders))
	for _, tndr := range tenders {
		dtos = append(dtos, *dto.TenderFromModel(tndr))
	}

	return dtos
}

func isPrivileged(ctx context.Context, tndr *tender.Tender) bool {
	emplInfo := api.UserInfo(ctx)
	return emplInfo.IsResponsible && isAuthor(ctx, tndr) ||
		tndr.Status() == string(tender.StatusPublished)
}

func isAuthor(ctx context.Context, tndr *tender.Tender) bool {
	orgRespInfo := api.OrgRespInfo(ctx)
	return tndr.Organization().ID() == orgRespInfo.OrganizationID
}
