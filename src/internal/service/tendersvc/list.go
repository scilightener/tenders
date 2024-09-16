package tendersvc

import (
	"context"
	"errors"
	"strconv"
	"tenders-management/internal/service"

	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/model/dto"
)

func (t *Svc) List(ctx context.Context, lim, off string, svcTypes []string) ([]dto.Tender, error) {
	log := t.logger.With("comp", "service.tendrsvc.List")

	limit, offset, svcTypes, err := validateParams(lim, off, svcTypes)
	if err != nil {
		log.Info("error validating params", sl.Err(err))
		return nil, err
	}

	tenders, err := t.tenderRepo.GetPublishedBySvcType(ctx, limit, offset, svcTypes)
	if errors.Is(err, tender.ErrNotFound) {
		log.Info("tenders not found", sl.Err(err))
		return []dto.Tender{}, nil
	} else if err != nil {
		log.Error("error listing tenders", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	dtos := mapTendersToDTOs(tenders)

	return dtos, nil
}

// validateParams is officially garbage.
func validateParams(lim, off string, svcTypes []string) (int, int, []string, error) {
	var limit, offset int
	if lim == "" {
		limit = 5
	} else {
		l, err := strconv.Atoi(lim)
		if err != nil || l < 0 || l > 50 {
			return 0, 0, nil, domain.ValidationErr(msg.ErrInvalidField("limit"))
		}
		limit = l
	}

	if off == "" {
		offset = 0
	} else {
		o, err := strconv.Atoi(off)
		if err != nil || o < 0 {
			return 0, 0, nil, domain.ValidationErr(msg.ErrInvalidField("offset"))
		}
		offset = o
	}

	if err := validateSvcTypes(svcTypes); err != nil {
		return 0, 0, nil, err
	}

	return limit, offset, svcTypes, nil
}

func validateSvcTypes(svcTypes []string) error {
	for i, st := range svcTypes {
		s, err := tender.ServiceTypeFromString(st)
		if err != nil {
			return err
		}
		svcTypes[i] = string(s)
	}

	return nil
}
