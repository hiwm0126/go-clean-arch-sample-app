package commandline

import (
	"example.com/internship_27_test/domain/service"
	"example.com/internship_27_test/infrastructure/datastore"
	"example.com/internship_27_test/usecase"
)

type Controller struct {
	initDataUseCase usecase.DataInitializationUseCase
	orderUseCase    usecase.OrderUseCase
	cancelUseCase   usecase.CancelUseCase
	shipUseCase     usecase.ShippingUseCase
	changeUseCase   usecase.ChangeUseCase
	expandUseCase   usecase.ExpandUseCase
}

func NewController() *Controller {
	orderRepo := datastore.NewOrderRepository()
	orderItemRepo := datastore.NewOrderItemRepository()
	productRepo := datastore.NewProductRepository()
	shipmentLimitRepo := datastore.NewShipmentLimitRepository()
	shippingAcceptablePeriodRepo := datastore.NewShippingAcceptablePeriodRepository()
	additionalShipmentLimitRepo := datastore.NewAdditionalShipmentLimitRepository()
	orderValidationgService := service.NewOrderValidatingService(
		orderItemRepo,
		shipmentLimitRepo,
		shippingAcceptablePeriodRepo,
		additionalShipmentLimitRepo,
	)
	orderCreatingService := service.NewOrderCreatingService(
		orderRepo,
		orderItemRepo,
		productRepo,
		shipmentLimitRepo,
		shippingAcceptablePeriodRepo,
		additionalShipmentLimitRepo,
		orderValidationgService,
	)
	orderCancelService := service.NewOrderCancelService(orderRepo, orderItemRepo)

	return &Controller{
		initDataUseCase: usecase.NewImportDataUseCase(
			productRepo,
			shipmentLimitRepo,
			shippingAcceptablePeriodRepo,
		),
		orderUseCase: usecase.NewOrderUseCase(
			orderCreatingService,
		),
		cancelUseCase: usecase.NewCancelUseCase(
			orderRepo,
			orderCancelService,
		),
		shipUseCase: usecase.NewShippingUseCase(
			orderRepo,
		),
		changeUseCase: usecase.NewChangeUseCase(
			orderRepo,
			orderItemRepo,
			orderCreatingService,
			orderCancelService,
			orderValidationgService,
		),
		expandUseCase: usecase.NewExpandUseCase(
			additionalShipmentLimitRepo,
		),
	}
}

func (c *Controller) InitData(req *usecase.DataInitializationUseCaseReq) error {
	return c.initDataUseCase.InitData(req)
}

func (c *Controller) Cancel(req *usecase.CancelUseCaseReq) (*usecase.CancelUseCaseRes, error) {
	return c.cancelUseCase.Cancel(req)
}

func (c *Controller) Order(req *usecase.OrderUseCaseReq) (*usecase.OrderUseCaseRes, error) {
	return c.orderUseCase.Order(req)
}

func (c *Controller) Ship(req *usecase.ShippingUseCaseReq) (*usecase.ShippingUseCaseRes, error) {
	return c.shipUseCase.Ship(req)
}

func (c *Controller) Change(req *usecase.ChangeUseCaseReq) (*usecase.ChangeUseCaseRes, error) {
	return c.changeUseCase.Change(req)
}

func (c *Controller) Expand(req *usecase.ExpandUseCaseReq) (*usecase.ExpandUseCaseRes, error) {
	return c.expandUseCase.Expand(req)
}
