package service

import (
	"context"
	"theapp/domain/model"
	"theapp/domain/repository"
)

type ShipmentLimitFactory interface {
	Create(ctx context.Context, shipmentDueDate string) (*model.ShipmentLimit, error)
}

type shipmentLimitFactory struct {
	shipmentLimitRepo           repository.ShipmentLimitRepository
	additionalShipmentLimitRepo repository.AdditionalShipmentLimitRepository
}

func NewShipmentLimitFactory(
	shipmentLimitRepo repository.ShipmentLimitRepository,
	additionalShipmentLimitRepo repository.AdditionalShipmentLimitRepository,
) ShipmentLimitFactory {
	return &shipmentLimitFactory{
		shipmentLimitRepo:           shipmentLimitRepo,
		additionalShipmentLimitRepo: additionalShipmentLimitRepo,
	}
}

func (f *shipmentLimitFactory) Create(ctx context.Context, shipmentDueDate string) (*model.ShipmentLimit, error) {
	// 出荷可能数取得
	shipmentLimit, err := f.shipmentLimitRepo.GetShipmentLimitByDate(ctx, shipmentDueDate)
	if err != nil {
		return nil, err
	}
	// ShipmentDueDateが含まれる期間で有効な、追加出荷可能数を取得
	asl, err := f.additionalShipmentLimitRepo.GetByShipmentDueDate(ctx, shipmentDueDate)
	if err != nil {
		return nil, err
	}
	// 出荷制限数に追加出荷可能数を設定
	shipmentLimit.SetAdditionalShipmentLimits(asl)

	return shipmentLimit, nil
}
