package tendersvc

import (
	"context"
	"errors"
	"tenders-management/internal/service"
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain/employee"
	"tenders-management/internal/model/domain/organization"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/model/dto"
)

func (t *Svc) Save(ctx context.Context, newTender dto.NewTender) (*dto.Tender, error) {
	log := t.logger.With("comp", "service.tendersvc.Save")

	status := tender.StatusCreated
	serviceType, err := tender.ServiceTypeFromString(newTender.ServiceType)
	if err != nil {
		log.Info("failed to parse service type", sl.Err(err))
		return nil, err
	}

	responsible, err := t.getResponsibleByEmployeeUsername(ctx, newTender.CreatorUsername)
	if err != nil {
		return nil, err
	} else if responsible == nil || responsible.Organization().ID() != newTender.OrganizationID {
		log.Info("unprivileged user", "username", newTender.CreatorUsername)
		return nil, service.ErrUnprivileged
	}

	tndr, err := tender.New(uuid.Nil, newTender.Name, newTender.Description, status, serviceType,
		1, organization.EmptyOrganizationWithID(newTender.OrganizationID),
		organization.EmptyResponsibleWithID(responsible.ID()), time.Now(), time.Now())
	if err != nil {
		log.Info("failed to create tender", sl.Err(err))
		return nil, err
	}

	tndr, err = t.saveTender(ctx, tndr)
	if err != nil {
		return nil, err
	}

	return dto.TenderFromModel(tndr), nil
}

func (t *Svc) getResponsibleByEmployeeUsername(
	ctx context.Context, username string,
) (*organization.Responsible, error) {
	log := t.logger.With("comp", "service.tendersvc.getResponsibleByEmployeeUsername")

	empl, err := t.emplRepo.GetByUsername(ctx, username)
	if err != nil && !errors.Is(err, employee.ErrNotFound) {
		log.Error("error retrieving employee", sl.Err(err))
		return nil, service.ErrUnknownError
	} else if errors.Is(err, employee.ErrNotFound) || empl == nil {
		log.Info("employee not found", "username", username)
		return nil, employee.ErrNotFound
	}
	responsible, err := t.respRepo.GetByUserID(ctx, empl.ID())
	if errors.Is(err, organization.ErrResponsibleNotFound) {
		log.Info("unprivileged user", "username", username)
		return nil, service.ErrUnprivileged
	} else if err != nil {
		log.Error("error retrieving responsible", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return responsible, nil
}

func (t *Svc) saveTender(ctx context.Context, tndr *tender.Tender) (*tender.Tender, error) {
	log := t.logger.With("comp", "service.tendersvc.saveTender")

	// TODO: should be within one transaction
	tndr, err := t.tenderRepo.Save(ctx, tndr)
	if err != nil {
		log.Error("failed to save tender", sl.Err(err))
		return nil, service.ErrUnknownError
	}
	tndrv := tender.NewVersionedFromTender(tndr)
	_, err = t.tenderVRepo.Save(ctx, tndrv)
	if err != nil {
		log.Error("failed to save tender version", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return tndr, nil
}
