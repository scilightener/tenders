package bidsvc

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain"
	"tenders-management/internal/model/domain/bid"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/service"
)

func (s *Svc) FindByID(ctx context.Context, bidID string) (dto.Bid, error) {
	log := s.logger.With("comp", "service.bidsvc.FindByID")

	bID, err := uuid.Parse(bidID)
	if err != nil {
		log.Info("failed to parse bidID", sl.Err(err))
		return dto.Bid{}, domain.ValidationErr(msg.ErrInvalidField("bidID"))
	}

	b, err := s.getBidByID(ctx, bID)
	if err != nil {
		return dto.Bid{}, err
	}
	if isBidAuthor(ctx, b) {
		return dto.BidFromModel(b), nil
	}

	return dto.Bid{}, service.ErrUnprivileged
}

func (s *Svc) getBidByID(ctx context.Context, id uuid.UUID) (*bid.Bid, error) {
	log := s.logger.With("comp", "service.bidsvc.getBidByID")

	b, err := s.bidRepo.GetByID(ctx, id)
	if errors.Is(err, bid.ErrNotFound) {
		log.Info("bid not found", "bidID", id)
		return nil, bid.ErrNotFound
	} else if err != nil {
		log.Error("error listing bids", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return b, nil
}
