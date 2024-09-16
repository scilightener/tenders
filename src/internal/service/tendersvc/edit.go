package tendersvc

import (
	"context"

	"github.com/google/uuid"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/service"
)

func (t *Svc) Edit(ctx context.Context, tenderID string, updTndr dto.UpdateTender) (*dto.Tender, error) {
	log := t.logger.With("comp", "service.tendrsvc.Edit")

	tID, err := uuid.Parse(tenderID)
	if err != nil {
		log.Info("failed to parse tenderID", sl.Err(err))
		return nil, domain.ValidationErr(msg.ErrInvalidField("tenderID"))
	}
	tndr, err := t.getTenderByID(ctx, tID)
	if err != nil {
		return nil, err
	}
	if !isAuthor(ctx, tndr) {
		log.Info("unprivileged user", "username", api.UserInfo(ctx).Username)
		return nil, service.ErrUnprivileged
	}

	err = updateTender(tndr, updTndr)
	if err != nil {
		log.Info("failed to update tender", sl.Err(err))
		return nil, domain.ValidationErr(msg.ErrInvalidField(err.Error()))
	}

	// TODO: should be within a transaction
	tndr, err = t.tenderRepo.Update(ctx, tID, tndr)
	if err != nil {
		log.Info("failed to update tender", sl.Err(err))
		return nil, service.ErrUnknownError
	}
	tndrv := tender.NewVersionedFromTender(tndr)
	_, err = t.tenderVRepo.Save(ctx, tndrv)
	if err != nil {
		log.Info("failed to save tender version", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return dto.TenderFromModel(tndr), nil
}

func updateTender(t *tender.Tender, updTndr dto.UpdateTender) error {
	if updTndr.ServiceType != nil {
		serviceType, err := tender.ServiceTypeFromString(*updTndr.ServiceType)
		if err != nil {
			return err
		}
		t.SetServiceType(serviceType)
	}
	if updTndr.Name != nil {
		err := t.SetName(*updTndr.Name)
		if err != nil {
			return err
		}
	}
	if updTndr.Description != nil {
		err := t.SetDescription(*updTndr.Description)
		if err != nil {
			return err
		}
	}

	if updTndr.Description != nil || updTndr.Name != nil || updTndr.ServiceType != nil {
		t.IncrementVersion()
	}
	return nil
}
