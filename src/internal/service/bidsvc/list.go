package bidsvc

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"tenders-management/internal/model/domain/bid"

	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/service"
)

func (s *Svc) List(ctx context.Context, lim, off, tID string) ([]dto.Bid, error) {
	log := s.logger.With("comp", "service.bidsvc.List")

	limit, offset, err := validateParams(lim, off)
	if err != nil {
		log.Info("error validating params", sl.Err(err))
		return nil, err
	}
	tenderID, err := uuid.Parse(tID)
	if err != nil {
		log.Info("failed to parse tenderID", sl.Err(err))
		return nil, domain.ValidationErr(msg.ErrInvalidField("tenderID"))
	}

	tender, err := s.getTenderByID(ctx, tenderID)
	if err != nil {
		return nil, err
	}

	tenders, err := s.bidRepo.GetPublishedByTenderID(ctx, tender.ID(), limit, offset)
	if errors.Is(err, bid.ErrNotFound) {
		log.Info("bids not found", sl.Err(err))
		return nil, bid.ErrNotFound
	} else if err != nil {
		log.Error("error listing tenders", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return mapBidsToDTOs(tenders), nil
}
