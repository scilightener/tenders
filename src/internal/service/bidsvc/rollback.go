package bidsvc

import (
	"context"
	"errors"
	"strconv"

	"github.com/google/uuid"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain"
	"tenders-management/internal/model/domain/bid"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/service"
)

func (s *Svc) Rollback(ctx context.Context, bidID string, version string) (dto.Bid, error) {
	log := s.logger.With("comp", "service.bidsvc.Rollback")

	tID, ver, err := parseRollbackInputs(bidID, version)
	if err != nil {
		log.Info("failed to parse inputs", sl.Err(err))
		return dto.Bid{}, err
	}

	b, err := s.getBidByID(ctx, tID)
	if err != nil {
		return dto.Bid{}, err
	}
	if !isBidAuthor(ctx, b) {
		log.Info("unprivileged user", "username", api.UserInfo(ctx).Username)
		return dto.Bid{}, service.ErrUnprivileged
	}

	bv, err := s.getVersionedByBidIDVersion(ctx, tID, ver)
	if err != nil {
		return dto.Bid{}, err
	}

	b, err = s.rollbackBidToVersioned(ctx, tID, b, bv)
	if err != nil {
		return dto.Bid{}, err
	}

	return dto.BidFromModel(b), nil
}

func parseRollbackInputs(bidID string, version string) (uuid.UUID, int, error) {
	bID, err := uuid.Parse(bidID)
	if err != nil {
		return bID, 0, domain.ValidationErr(msg.ErrInvalidField("bidID"))
	}
	ver, err := strconv.Atoi(version)
	if err != nil {
		return bID, ver, domain.ValidationErr(msg.ErrInvalidField("version"))
	}
	return bID, ver, nil
}

func (s *Svc) getVersionedByBidIDVersion(
	ctx context.Context, bidID uuid.UUID, version int,
) (*bid.Versioned, error) {
	log := s.logger.With("comp", "service.bidsvc.getVersionedByBidIDVersion")

	bv, err := s.bidVRepo.GetByBidIDVersion(ctx, bidID, version)
	if errors.Is(err, bid.ErrNotFound) {
		log.Info("version not found", "bidID", bidID, "version", version)
		return nil, bid.ErrNotFound
	} else if err != nil {
		log.Error("failed to get version", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return bv, nil
}

func (s *Svc) rollbackBidToVersioned(
	ctx context.Context, bidID uuid.UUID, b *bid.Bid, bv *bid.Versioned,
) (*bid.Bid, error) {
	log := s.logger.With("comp", "service.bidsvc.rollbackBidToVersioned")

	b.IncrementVersion()
	bv.Version = b.Version()
	b, err := bv.ToBid()
	if err != nil {
		log.Info("failed to convert to bid", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	// TODO: transaction
	b, err = s.bidRepo.Update(ctx, bidID, b)
	if err != nil {
		log.Info("failed to update bid", sl.Err(err))
		return nil, service.ErrUnknownError
	}
	_, err = s.bidVRepo.Save(ctx, bv)
	if err != nil {
		log.Info("failed to save bid version", sl.Err(err))
		return nil, service.ErrUnknownError
	}
	return b, nil
}
