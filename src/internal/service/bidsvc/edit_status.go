package bidsvc

import (
	"context"

	"github.com/google/uuid"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain"
	"tenders-management/internal/model/domain/bid"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/service"
)

func (s *Svc) EditStatus(ctx context.Context, bidID string, newStatus string) (dto.Bid, error) {
	log := s.logger.With("comp", "service.tendrsvc.EditStatus")

	bID, err := uuid.Parse(bidID)
	if err != nil {
		log.Info("failed to parse bidID", sl.Err(err))
		return dto.Bid{}, domain.ValidationErr(msg.ErrInvalidField("bidID"))
	}
	status, err := bid.StatusFromString(newStatus)
	if err != nil {
		log.Info("failed to parse status", sl.Err(err))
		return dto.Bid{}, domain.ValidationErr(msg.ErrInvalidField("status"))
	}

	b, err := s.getBidByID(ctx, bID)
	if err != nil {
		return dto.Bid{}, err
	}
	if !isBidAuthor(ctx, b) {
		log.Info("unprivileged user", "username", api.UserInfo(ctx).Username)
		return dto.Bid{}, service.ErrUnprivileged
	}

	if err = b.TransitToStatus(status); err != nil {
		log.Info("invalid status transition", "from", b.Status(), "to", status)
		return dto.Bid{}, err
	}

	b, err = s.bidRepo.Update(ctx, bID, b)
	if err != nil {
		log.Error("failed to update tender", sl.Err(err))
		return dto.Bid{}, service.ErrUnknownError
	}

	return dto.BidFromModel(b), nil
}
