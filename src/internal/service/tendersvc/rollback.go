package tendersvc

import (
	"context"
	"errors"
	"strconv"
	"tenders-management/internal/service"

	"github.com/google/uuid"

	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/model/dto"
)

func (t *Svc) Rollback(ctx context.Context, tenderID string, version string) (dto.Tender, error) {
	log := t.logger.With("comp", "service.tendrsvc.Rollback")

	tID, ver, err := parseRollbackInputs(tenderID, version)
	if err != nil {
		log.Info("failed to parse inputs", sl.Err(err))
		return dto.Tender{}, err
	}

	tndr, err := t.getTenderByID(ctx, tID)
	if err != nil {
		return dto.Tender{}, err
	}
	if !isAuthor(ctx, tndr) {
		log.Info("unprivileged user", "username", api.UserInfo(ctx).Username)
		return dto.Tender{}, service.ErrUnprivileged
	}

	tndrv, err := t.getVersionedByTenderIDVersion(ctx, tID, ver)
	if err != nil {
		return dto.Tender{}, err
	}

	tndr, err = t.rollbackTenderToVersioned(ctx, tID, tndr, tndrv)
	if err != nil {
		return dto.Tender{}, err
	}

	return *dto.TenderFromModel(tndr), nil
}

func parseRollbackInputs(tenderID string, version string) (uuid.UUID, int, error) {
	tID, err := uuid.Parse(tenderID)
	if err != nil {
		return tID, 0, domain.ValidationErr(msg.ErrInvalidField("tenderID"))
	}
	ver, err := strconv.Atoi(version)
	if err != nil {
		return tID, ver, domain.ValidationErr(msg.ErrInvalidField("version"))
	}
	return tID, ver, nil
}

func (t *Svc) getVersionedByTenderIDVersion(
	ctx context.Context, tenderID uuid.UUID, version int,
) (*tender.Versioned, error) {
	log := t.logger.With("comp", "service.tendrsvc.getVersionedByTenderIDVersion")

	tndrv, err := t.tenderVRepo.GetByTenderIDVersion(ctx, tenderID, version)
	if errors.Is(err, tender.ErrNotFound) {
		log.Info("version not found", "tenderID", tenderID, "version", version)
		return nil, tender.ErrNotFound
	} else if err != nil {
		log.Error("failed to get version", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return tndrv, nil
}

func (t *Svc) rollbackTenderToVersioned(
	ctx context.Context, tenderID uuid.UUID, tndr *tender.Tender, tndrv *tender.Versioned,
) (*tender.Tender, error) {
	log := t.logger.With("comp", "service.tendrsvc.rollbackTenderToVersioned")

	tndr.IncrementVersion()
	tndrv.Version = tndr.Version()
	tndr, err := tndrv.ToTender()
	if err != nil {
		log.Info("failed to convert to tender", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	// TODO: transaction
	tndr, err = t.tenderRepo.Update(ctx, tenderID, tndr)
	if err != nil {
		log.Info("failed to update tender", sl.Err(err))
		return nil, service.ErrUnknownError
	}
	_, err = t.tenderVRepo.Save(ctx, tndrv)
	if err != nil {
		log.Info("failed to save tender version", sl.Err(err))
		return nil, service.ErrUnknownError
	}
	return tndr, nil
}
