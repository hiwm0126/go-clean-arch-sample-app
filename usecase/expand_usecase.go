package usecase

import (
	"context"
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
	Expand(ctx context.Context, req *ExpandUseCaseReq) (*ExpandUseCaseRes, error)
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

func (e *expandUseCase) Expand(ctx context.Context, req *ExpandUseCaseReq) (*ExpandUseCaseRes, error) {
	// 追加出荷制限モデルの作成
	additionalShipmentLimit, err := model.NewAdditionalShipmentLimit(req.Quantity, req.From, req.To)
	if err != nil {
		return nil, err
	}

	// 追加出荷制限の保存
	err = e.additionalShipmentLimitRepo.Save(ctx, additionalShipmentLimit)
	if err != nil {
		return nil, err
	}
	return &ExpandUseCaseRes{req.ExpandRequestTime}, nil
}
