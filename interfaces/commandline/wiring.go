package commandline

import (
	"fmt"

	"theapp/domain/service"
	"theapp/infrastructure/datastore"
	"theapp/interfaces/commandline/handler"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

// appDeps は Composition Root で組み立てたアプリケーション依存（テストで差し替えしやすいよう集約）
type appDeps struct {
	dispatcher *Dispatcher
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
	initDataUseCase := usecase.NewDataInitializationUseCase(
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

	d := &Dispatcher{
		handlers: make(map[cmdname.CommandName]dispatchFn),
	}
	if err := registerDefaultCommands(d,
		initDataUseCase,
		orderUseCase,
		cancelUseCase,
		shipUseCase,
		changeUseCase,
		expandUseCase,
	); err != nil {
		panic(fmt.Sprintf("commandline: registerDefaultCommands: %v", err))
	}

	return &appDeps{dispatcher: d}
}

// registerDefaultCommands は CLI 既定の CommandHandler を Dispatcher に登録する（Composition Root / 唯一の DI 組み立て箇所）
func registerDefaultCommands(
	d *Dispatcher,
	initData usecase.DataInitializationUseCase,
	order usecase.OrderUseCase,
	cancel usecase.CancelUseCase,
	ship usecase.ShippingUseCase,
	change usecase.ChangeUseCase,
	expand usecase.ExpandUseCase,
) error {
	return registerHandlers(d,
		handler.NewInitDataCommandHandler(initData),
		handler.NewOrderCommandHandler(order),
		handler.NewCancelCommandHandler(cancel),
		handler.NewShippingCommandHandler(ship),
		handler.NewChangeCommandHandler(change),
		handler.NewExpandCommandHandler(expand),
	)
}

// registerHandlers は複数 CommandHandler を順に登録する（wiring 専用）
func registerHandlers(d *Dispatcher, hs ...handler.CommandHandler) error {
	for _, h := range hs {
		if err := d.RegisterHandler(h); err != nil {
			return err
		}
	}
	return nil
}
