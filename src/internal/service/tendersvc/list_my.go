package tendersvc

import (
	"context"
	"errors"
	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/service"
)

func (t *Svc) ListMy(ctx context.Context, lim, off string) ([]dto.Tender, error) {
	log := t.logger.With("comp", "service.tendersvc.ListMy")

	limit, offset, _, err := validateParams(lim, off, nil)
	if err != nil {
		log.Info("error validating params", sl.Err(err))
		return nil, err
	}

	employee := api.UserInfo(ctx)
	if !employee.IsResponsible {
		log.Info("user is not responsible for any organization", "username", employee.Username)
		return nil, service.ErrUnprivileged
	}
	creatorID := api.OrgRespInfo(ctx).ID
	tenders, err := t.tenderRepo.GetByCreatorID(ctx, creatorID, limit, offset)
	if errors.Is(err, tender.ErrNotFound) {
		log.Info("tenders not found", sl.Err(err))
		return []dto.Tender{}, nil
	} else if err != nil {
		log.Error("error listing tenders", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return mapTendersToDTOs(tenders), nil
}
