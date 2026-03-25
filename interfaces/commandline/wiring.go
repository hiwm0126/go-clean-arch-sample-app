package commandline

import (
	"theapp/domain/service"
	"theapp/infrastructure/datastore"
	"theapp/usecase"
)

// appDeps は Composition Root で組み立てたアプリケーション依存（テストで差し替えしやすいよう集約）
type appDeps struct {
	Dispatcher *Dispatcher
}

func newAppDeps() *appDeps {
	orderRepo := datastore.NewOrderRepository()
	orderItemRepo := datastore.NewOrderItemRepository()
	productRepo := datastore.NewProductRepository()
	shipmentLimitRepo := datastore.NewShipmentLimitRepository()
	shippingAcceptablePeriodRepo := datastore.NewShippingAcceptablePeriodRepository()
	additionalShipmentLimitRepo := datastore.NewAdditionalShipmentLimitRepository()
	shipmentLimitProvider := service.NewShipmentLimitProvider(shipmentLimitRepo, additionalShipmentLimitRepo)
	orderValidationService := service.NewOrderValidatingService(
		orderItemRepo,
		shipmentLimitProvider,
		shippingAcceptablePeriodRepo,
	)
	orderFactory := service.NewOrderFactory(
		orderRepo,
		orderItemRepo,
		productRepo,
		shipmentLimitRepo,
		shippingAcceptablePeriodRepo,
		additionalShipmentLimitRepo,
	)
	orderCancelService := service.NewOrderCancelService(orderRepo, orderItemRepo)
	initDataUseCase := usecase.NewImportDataUseCase(
		productRepo,
		shipmentLimitRepo,
		shippingAcceptablePeriodRepo,
	)
	orderUseCase := usecase.NewOrderUseCase(
		orderFactory,
		orderValidationService,
	)
	cancelUseCase := usecase.NewCancelUseCase(
		orderRepo,
		orderCancelService,
	)
	shipUseCase := usecase.NewShippingUseCase(
		orderRepo,
	)
	changeUseCase := usecase.NewChangeUseCase(
		orderRepo,
		orderItemRepo,
		orderFactory,
		orderCancelService,
		orderValidationService,
	)
	expandUseCase := usecase.NewExpandUseCase(
		additionalShipmentLimitRepo,
	)

	d := NewDispatcher(
		initDataUseCase,
		orderUseCase,
		cancelUseCase,
		shipUseCase,
		changeUseCase,
		expandUseCase,
	)

	return &appDeps{Dispatcher: d}
}
