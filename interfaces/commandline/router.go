package commandline

import (
	"context"
	"theapp/domain/service"
	"theapp/infrastructure/datastore"
	"theapp/interfaces/commandline/handler"
	"theapp/usecase"
)

type Router interface {
	Routing(args [][]string) error
}
type router struct {
	commandHandlers []handler.CommandHandler
}

func NewRouter() Router {
	router := &router{
		commandHandlers: make([]handler.CommandHandler, 0),
	}
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

	// コマンドハンドラーの登録
	router.registerCommandHandler(handler.NewInitDataCommandHandler(initDataUseCase))
	router.registerCommandHandler(handler.NewOrderCommandHandler(orderUseCase))
	router.registerCommandHandler(handler.NewCancelCommandHandler(cancelUseCase))
	router.registerCommandHandler(handler.NewShippingCommandHandler(shipUseCase))
	router.registerCommandHandler(handler.NewChangeCommandHandler(changeUseCase))
	router.registerCommandHandler(handler.NewExpandCommandHandler(expandUseCase))

	return router
}

func (r *router) Routing(args [][]string) error {
	// リクエストパラメーターを解析
	paramFactory := NewParamFactory()
	paramList, err := paramFactory.Create(args)
	if err != nil {
		return err
	}

	// 各リクエストパラメーターに対してコマンドハンドラーを適用
	for _, reqParam := range paramList {
		for _, h := range r.commandHandlers {
			if h.CanHandle(reqParam) {
				if err := h.Handle(context.Background(), reqParam); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (r *router) registerCommandHandler(h handler.CommandHandler) {
	r.commandHandlers = append(r.commandHandlers, h)
}
