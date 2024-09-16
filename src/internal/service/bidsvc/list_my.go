package bidsvc

import (
	"context"
	"errors"
	"strconv"
	"tenders-management/internal/model/domain/organization"

	"github.com/google/uuid"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain"
	"tenders-management/internal/model/domain/bid"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/service"
)

func (s *Svc) ListMy(ctx context.Context, lim, off string) ([]dto.Bid, error) {
	log := s.logger.With("comp", "service.bidsvc.ListMy")

	limit, offset, err := validateParams(lim, off)
	if err != nil {
		log.Info("error validating params", sl.Err(err))
		return nil, err
	}

	bids, err := s.getBidsByUserOrOrganizationResponsible(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return mapBidsToDTOs(bids), nil
}

func (s *Svc) getBidsByUserOrOrganizationResponsible(
	ctx context.Context, limit, offset int,
) ([]*bid.Bid, error) {
	log := s.logger.With("comp", "service.bidsvc.getBidsByUserOrOrganizationResponsible")
	userInfo := api.UserInfo(ctx)
	//orgRespInfo := api.OrgRespInfo(ctx)
	var (
		userID uuid.UUID
		bids   []*bid.Bid
		err    error
	)
	//if userInfo.IsResponsible {
	//	userID = orgRespInfo.OrganizationID
	//	bids, err = s.bidRepo.GetByOrganizationID(ctx, userID, limit, offset)
	//} else {
	userID = userInfo.UserID
	user, err := s.getEmployeeByID(ctx, userID)
	if err != nil {
		log.Info("user not found", "userID", userID)
		return nil, err
	}
	userID = user.ID()
	bids, err = s.bidRepo.GetByUserID(ctx, userID, limit, offset)
	//}

	if errors.Is(err, bid.ErrNotFound) {
		log.Info("bids not found", "userID", userID)
		return nil, bid.ErrNotFound
	} else if errors.Is(err, organization.ErrOrganizationNotFound) {
		log.Info("organization not found", "userID", userID)
		return nil, organization.ErrOrganizationNotFound
	} else if err != nil {
		log.Error("error listing bids", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return bids, nil
}

func validateParams(lim, off string) (int, int, error) {
	var limit, offset int
	if lim == "" {
		limit = 5
	} else {
		l, err := strconv.Atoi(lim)
		if err != nil || l < 0 || l > 50 {
			return 0, 0, domain.ValidationErr(msg.ErrInvalidField("limit"))
		}
		limit = l
	}

	if off == "" {
		offset = 0
	} else {
		o, err := strconv.Atoi(off)
		if err != nil || o < 0 {
			return 0, 0, domain.ValidationErr(msg.ErrInvalidField("offset"))
		}
		offset = o
	}

	return limit, offset, nil
}
