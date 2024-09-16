package tendersvc

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/service"

	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain"
	"tenders-management/internal/model/dto"
)

func (t *Svc) FindByID(ctx context.Context, tenderID string) (dto.Tender, error) {
	log := t.logger.With("comp", "service.tendrsvc.FindByID")

	tID, err := uuid.Parse(tenderID)
	if err != nil {
		log.Info("failed to parse tenderID", sl.Err(err))
		return dto.Tender{}, domain.ValidationErr(msg.ErrInvalidField("tenderID"))
	}

	tndr, err := t.getTenderByID(ctx, tID)
	if err != nil {
		return dto.Tender{}, err
	}
	if !isPrivileged(ctx, tndr) {
		return dto.Tender{}, service.ErrUnprivileged
	}

	return *dto.TenderFromModel(tndr), nil
}

func (t *Svc) getTenderByID(ctx context.Context, tenderID uuid.UUID) (*tender.Tender, error) {
	log := t.logger.With("comp", "service.tendrsvc.getTenderByID")

	tndr, err := t.tenderRepo.GetByID(ctx, tenderID)
	if errors.Is(err, tender.ErrNotFound) {
		log.Info("tender not found", "tenderID", tenderID)
		return nil, tender.ErrNotFound
	} else if err != nil {
		log.Error("error listing tenders", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return tndr, nil
}
