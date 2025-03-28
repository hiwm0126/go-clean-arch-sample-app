package usecase

import (
	"example.com/internship_27_test/domain/model"
	"example.com/internship_27_test/domain/repository"
	"time"
)

type ExpandUseCaseReq struct {
	ExpandRequestTime time.Time
	From              string
	To                string
	Quantity          int
}

type ExpandUseCaseRes struct {
	ExpandRequestTime time.Time
}

type ExpandUseCase interface {
	Expand(req *ExpandUseCaseReq) (*ExpandUseCaseRes, error)
}

type expandUseCase struct {
	additionalShipmentLimitRepo repository.AdditionalShipmentLimitRepository
}

func NewExpandUseCase(
	additionalShipmentLimitRepo repository.AdditionalShipmentLimitRepository,
) ExpandUseCase {
	return &expandUseCase{
		additionalShipmentLimitRepo: additionalShipmentLimitRepo,
	}
}

func (e *expandUseCase) Expand(req *ExpandUseCaseReq) (*ExpandUseCaseRes, error) {
	additionalShipmentLimit := model.NewAdditionalShipmentLimit(req.Quantity, req.From, req.To)
	err := e.additionalShipmentLimitRepo.Save(additionalShipmentLimit)
	if err != nil {
		return nil, err
	}
	return &ExpandUseCaseRes{req.ExpandRequestTime}, nil
}
