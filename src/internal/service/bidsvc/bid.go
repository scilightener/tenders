package bidsvc

import (
	"context"
	"log/slog"
	"tenders-management/internal/lib/api"
	"tenders-management/internal/model/domain/bid"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/storage/repo"
)

type Svc struct {
	tenderRepo repo.Tender
	bidRepo    repo.Bid
	bidVRepo   repo.BidVersions
	emplRepo   repo.Employee
	orgRepo    repo.Organization
	respRepo   repo.OrganizationResponsible
	logger     *slog.Logger
}

func NewBid(logger *slog.Logger, emplRepo repo.Employee, orgRepo repo.Organization,
	respRepo repo.OrganizationResponsible,
	tenderRepo repo.Tender, bidRepo repo.Bid, bidVRepo repo.BidVersions,
) *Svc {
	return &Svc{
		tenderRepo: tenderRepo,
		bidRepo:    bidRepo,
		bidVRepo:   bidVRepo,
		emplRepo:   emplRepo,
		orgRepo:    orgRepo,
		respRepo:   respRepo,
		logger:     logger,
	}
}

func isBidAuthor(ctx context.Context, b *bid.Bid) bool {
	userInfo := api.UserInfo(ctx)
	orgRespInfo := api.OrgRespInfo(ctx)
	isUserAuthor := b.AuthorType() == string(bid.AuthorTypeUser) &&
		b.UserID() != nil && *b.UserID() == userInfo.UserID
	isOrgRespAuthor := b.AuthorType() == string(bid.AuthorTypeOrganization) &&
		b.OrganizationID() != nil && *b.OrganizationID() == orgRespInfo.OrganizationID
	return isUserAuthor || isOrgRespAuthor
}

func mapBidsToDTOs(bids []*bid.Bid) []dto.Bid {
	dtos := make([]dto.Bid, 0, len(bids))
	for _, b := range bids {
		dtos = append(dtos, dto.BidFromModel(b))
	}

	return dtos
}
