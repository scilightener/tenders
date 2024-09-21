package tendersvc

import (
	"context"
	"tenders-management/internal/service"

	"github.com/google/uuid"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/model/dto"
)

func (t *Svc) EditStatus(ctx context.Context, tenderID string, newStatus string) (*dto.Tender, error) {
	log := t.logger.With("comp", "service.tendrsvc.EditStatus")

	tID, err := uuid.Parse(tenderID)
	if err != nil {
		log.Info("failed to parse tenderID", sl.Err(err))
		return nil, domain.ValidationErr(msg.ErrInvalidField("tenderID"))
	}
	status, err := tender.StatusFromString(newStatus)
	if err != nil {
		log.Info("failed to parse status", sl.Err(err))
		return nil, domain.ValidationErr(msg.ErrInvalidField("status"))
	}

	tndr, err := t.getTenderByID(ctx, tID)
	if err != nil {
		return nil, err
	}
	if !isAuthor(ctx, tndr) {
		log.Info("unprivileged user", "username", api.UserInfo(ctx).Username)
		return nil, service.ErrUnprivileged
	}

	if err = tndr.TransitToStatus(status); err != nil {
		log.Info("invalid status transition", "from", tndr.Status(), "to", status)
		return nil, err
	}

	tndr, err = t.tenderRepo.Update(ctx, tID, tndr)
	if err != nil {
		log.Error("failed to update tender", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return dto.TenderFromModel(tndr), nil
}
