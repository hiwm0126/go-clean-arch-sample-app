package commandline

import (
	"theapp/domain/service"
	"theapp/infrastructure/datastore"
	"theapp/interfaces/commandline/handler"
	"theapp/usecase"
)

type Router interface {
	Routing(args [][]string) error
}
type router struct {
	handlers []handler.Handler
}

func NewRouter() Router {
	router := &router{
		handlers: make([]handler.Handler, 0),
	}
	orderRepo := datastore.NewOrderRepository()
	orderItemRepo := datastore.NewOrderItemRepository()
	productRepo := datastore.NewProductRepository()
	shipmentLimitRepo := datastore.NewShipmentLimitRepository()
	shippingAcceptablePeriodRepo := datastore.NewShippingAcceptablePeriodRepository()
	additionalShipmentLimitRepo := datastore.NewAdditionalShipmentLimitRepository()
	shipmentLimitFactory := service.NewShipmentLimitFactory(shipmentLimitRepo, additionalShipmentLimitRepo)
	orderValidationService := service.NewOrderValidatingService(
		orderItemRepo,
		shipmentLimitFactory,
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

	// ハンドラーの登録
	router.registerHandler(handler.NewInitDataHandler(initDataUseCase))
	router.registerHandler(handler.NewOrderHandler(orderUseCase))
	router.registerHandler(handler.NewCancelHandler(cancelUseCase))
	router.registerHandler(handler.NewShippingHandler(shipUseCase))
	router.registerHandler(handler.NewChangeHandler(changeUseCase))
	router.registerHandler(handler.NewExpandHandler(expandUseCase))

	return router
}

func (r *router) Routing(args [][]string) error {
	// リクエストパラメーターを解析
	paramFactory := NewParamFactory()
	paramList, err := paramFactory.Create(args)
	if err != nil {
		return err
	}

	// 各リクエストパラメーターに対してハンドラーを適用
	for _, reqParam := range paramList {
		for _, h := range r.handlers {
			if h.CanHandle(reqParam) {
				if err := h.Handle(reqParam); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (r *router) registerHandler(h handler.Handler) {
	r.handlers = append(r.handlers, h)
}
