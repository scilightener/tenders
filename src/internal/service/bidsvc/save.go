package bidsvc

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"tenders-management/internal/lib/logger/sl"
	"tenders-management/internal/model/domain/bid"
	"tenders-management/internal/model/domain/employee"
	"tenders-management/internal/model/domain/organization"
	"tenders-management/internal/model/domain/tender"
	"tenders-management/internal/model/dto"
	"tenders-management/internal/service"
)

func (s *Svc) Save(ctx context.Context, newBid dto.NewBid) (dto.Bid, error) {
	log := s.logger.With("comp", "service.bidsvc.Save")

	b, err := s.getBidFromNewDto(ctx, newBid)
	if err != nil {
		return dto.Bid{}, err
	}

	// TODO: transaction
	b, err = s.bidRepo.Save(ctx, b)
	if err != nil {
		log.Error("failed to save bid", sl.Err(err))
		return dto.Bid{}, service.ErrUnknownError
	}
	bv := bid.NewVersionedFromBid(b)
	_, err = s.bidVRepo.Save(ctx, bv)
	if err != nil {
		log.Error("failed to save bid version", sl.Err(err))
		return dto.Bid{}, service.ErrUnknownError
	}

	return dto.BidFromModel(b), nil
}

func (s *Svc) getBidFromNewDto(ctx context.Context, newBid dto.NewBid) (*bid.Bid, error) {
	log := s.logger.With("comp", "service.bidsvc.getBidFromNewDto")

	status := bid.StatusCreated
	authorType, err := bid.AuthorTypeFromString(newBid.AuthorType)
	if err != nil {
		log.Info("failed to parse author type", sl.Err(err))
		return nil, err
	}

	tndr, err := s.getTenderByID(ctx, newBid.TenderID)
	if err != nil {
		return nil, err
	}

	var (
		org *organization.Organization
		usr *employee.Employee
	)

	if authorType == bid.AuthorTypeOrganization {
		org, err = s.getOrganizationByID(ctx, newBid.AuthorID)
		if err != nil {
			return nil, err
		}
	} else if authorType == bid.AuthorTypeUser {
		usr, err = s.getEmployeeByID(ctx, newBid.AuthorID)
		if err != nil {
			return nil, err
		}
	}

	return bid.New(uuid.Nil, newBid.Name, newBid.Description, status, authorType,
		org, usr, 1, tndr, time.Now(), time.Now())
}

func (s *Svc) getTenderByID(ctx context.Context, tenderID uuid.UUID) (*tender.Tender, error) {
	log := s.logger.With("comp", "service.bidsvc.getTenderByID")

	tndr, err := s.tenderRepo.GetByID(ctx, tenderID)
	if errors.Is(err, tender.ErrNotFound) {
		log.Info("tender not found", "id", tenderID)
		return nil, tender.ErrNotFound
	} else if err != nil {
		log.Error("failed to retrieve tender", sl.Err(err))
		return nil, service.ErrUnknownError
	} else if tndr.Status() != string(tender.StatusPublished) {
		log.Info("tender is not published", "status", tndr.Status())
		return nil, service.ErrUnprivileged
	}

	return tndr, nil
}

func (s *Svc) getOrganizationByID(ctx context.Context, orgID uuid.UUID) (*organization.Organization, error) {
	log := s.logger.With("comp", "service.bidsvc.getOrganizationByID")

	org, err := s.orgRepo.GetByID(ctx, orgID)
	if errors.Is(err, organization.ErrOrganizationNotFound) {
		log.Info("organization not found", "id", orgID)
		return nil, organization.ErrOrganizationNotFound
	} else if err != nil {
		log.Error("failed to retrieve organization", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return org, nil
}

func (s *Svc) getEmployeeByID(ctx context.Context, empID uuid.UUID) (*employee.Employee, error) {
	log := s.logger.With("comp", "service.bidsvc.getEmployeeByID")

	emp, err := s.emplRepo.GetByID(ctx, empID)
	if errors.Is(err, employee.ErrNotFound) {
		log.Info("employee not found", "id", empID)
		return nil, employee.ErrNotFound
	} else if err != nil {
		log.Error("failed to retrieve employee", sl.Err(err))
		return nil, service.ErrUnknownError
	}

	return emp, nil
}
