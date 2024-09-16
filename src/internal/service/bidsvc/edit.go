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

func (s *Svc) Edit(ctx context.Context, bidID string, updTndr dto.UpdateBid) (dto.Bid, error) {
	log := s.logger.With("comp", "service.bidsvc.Edit")

	bID, err := uuid.Parse(bidID)
	if err != nil {
		log.Info("failed to parse tenderID", sl.Err(err))
		return dto.Bid{}, domain.ValidationErr(msg.ErrInvalidField("tenderID"))
	}
	b, err := s.getBidByID(ctx, bID)
	if err != nil {
		return dto.Bid{}, err
	}
	if !isBidAuthor(ctx, b) {
		log.Info("unprivileged user", "username", api.UserInfo(ctx).Username)
		return dto.Bid{}, service.ErrUnprivileged
	}

	err = updateBid(b, updTndr)
	if err != nil {
		log.Info("failed to update tender", sl.Err(err))
		return dto.Bid{}, domain.ValidationErr(msg.ErrInvalidField(err.Error()))
	}

	// TODO: transaction
	b, err = s.bidRepo.Update(ctx, bID, b)
	if err != nil {
		log.Info("failed to update tender", sl.Err(err))
		return dto.Bid{}, service.ErrUnknownError
	}
	bv := bid.NewVersionedFromBid(b)
	_, err = s.bidVRepo.Save(ctx, bv)
	if err != nil {
		log.Info("failed to save tender version", sl.Err(err))
		return dto.Bid{}, service.ErrUnknownError
	}

	return dto.BidFromModel(b), nil
}

func updateBid(b *bid.Bid, updB dto.UpdateBid) error {
	if updB.Name != nil {
		err := b.SetName(*updB.Name)
		if err != nil {
			return err
		}
	}
	if updB.Description != nil {
		err := b.SetDescription(*updB.Description)
		if err != nil {
			return err
		}
	}

	if updB.Description != nil || updB.Name != nil {
		b.IncrementVersion()
	}
	return nil
}
