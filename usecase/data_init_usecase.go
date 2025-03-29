package usecase

import (
	"context"
	"theapp/domain/model"
	"theapp/domain/repository"
)

type DataInitializationUseCaseReq struct {
	NumOfProduct             int
	ShipmentLimitThreshold   int
	ShipmentAcceptablePeriod int
	ProductNumberList        []string
	ShipmentLimitFlags       map[model.DayOfWeek]bool
	NumOfQuery               int
}

type DataInitializationUseCase interface {
	InitData(ctx context.Context, req *DataInitializationUseCaseReq) error
}

type dataInitializationUseCase struct {
	productRepo                  repository.ProductRepository
	shipmentLimitRepo            repository.ShipmentLimitRepository
	shippingAcceptablePeriodRepo repository.ShippingAcceptablePeriodRepository
}

func NewImportDataUseCase(
	productRepo repository.ProductRepository,
	shipmentLimitRepo repository.ShipmentLimitRepository,
	shippingAcceptablePeriodRepo repository.ShippingAcceptablePeriodRepository,
) DataInitializationUseCase {
	return &dataInitializationUseCase{
		productRepo:                  productRepo,
		shipmentLimitRepo:            shipmentLimitRepo,
		shippingAcceptablePeriodRepo: shippingAcceptablePeriodRepo,
	}
}

// InitData 初期データをインポートする
func (u *dataInitializationUseCase) InitData(ctx context.Context, req *DataInitializationUseCaseReq) error {
	// 商品マスタを作成
	for _, productNumber := range req.ProductNumberList {
		product := model.NewProduct(productNumber)
		err := u.productRepo.Save(ctx, product)
		if err != nil {
			return err
		}
	}

	// 出荷可能数マスタを作成
	for dayOfWeek, shipmentFlag := range req.ShipmentLimitFlags {
		var quantity int
		if shipmentFlag {
			quantity = req.ShipmentLimitThreshold
		} else {
			quantity = 0
		}

		shipmentLimit := model.NewShipmentLimit(dayOfWeek, quantity)
		err := u.shipmentLimitRepo.Save(ctx, shipmentLimit)
		if err != nil {
			return err
		}
	}

	// 出荷可能期間マスタを作成
	shippingAcceptablePeriod := model.NewShippingAcceptablePeriod(req.ShipmentAcceptablePeriod)
	err := u.shippingAcceptablePeriodRepo.Save(ctx, shippingAcceptablePeriod)
	if err != nil {
		return err
	}

	return nil
}
