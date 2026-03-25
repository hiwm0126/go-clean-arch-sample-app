package commandline

import (
	"context"
	"testing"
	"time"

	"theapp/interfaces/commandline/cli"
	"theapp/interfaces/commandline/handler"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type stubInit struct{ called bool }

func (s *stubInit) InitData(_ context.Context, _ *usecase.DataInitializationUseCaseReq) error {
	s.called = true
	return nil
}

type stubOrder struct{}

func (stubOrder) Order(_ context.Context, _ *usecase.OrderUseCaseReq) (*usecase.OrderUseCaseRes, error) {
	return &usecase.OrderUseCaseRes{}, nil
}

type stubCancel struct{}

func (stubCancel) Cancel(_ context.Context, _ *usecase.CancelUseCaseReq) (*usecase.CancelUseCaseRes, error) {
	return &usecase.CancelUseCaseRes{}, nil
}

type stubShip struct{}

func (stubShip) Ship(_ context.Context, _ *usecase.ShippingUseCaseReq) (*usecase.ShippingUseCaseRes, error) {
	return &usecase.ShippingUseCaseRes{}, nil
}

type stubChange struct{}

func (stubChange) Change(_ context.Context, _ *usecase.ChangeUseCaseReq) (*usecase.ChangeUseCaseRes, error) {
	return &usecase.ChangeUseCaseRes{}, nil
}

type stubExpand struct{}

func (stubExpand) Expand(_ context.Context, _ *usecase.ExpandUseCaseReq) (*usecase.ExpandUseCaseRes, error) {
	return &usecase.ExpandUseCaseRes{}, nil
}

func newTestDispatcher(
	initData usecase.DataInitializationUseCase,
	order usecase.OrderUseCase,
	cancel usecase.CancelUseCase,
	ship usecase.ShippingUseCase,
	change usecase.ChangeUseCase,
	expand usecase.ExpandUseCase,
) Dispatcher {
	d := Dispatcher{handlers: make(map[cmdname.CommandName]dispatchFn)}
	if err := registerHandlers(&d,
		handler.NewInitDataCommandHandler(initData),
		handler.NewOrderCommandHandler(order),
		handler.NewCancelCommandHandler(cancel),
		handler.NewShippingCommandHandler(ship),
		handler.NewChangeCommandHandler(change),
		handler.NewExpandCommandHandler(expand),
	); err != nil {
		panic(err)
	}
	return d
}

func TestDispatcher_Dispatch_routesInitData(t *testing.T) {
	t.Parallel()
	init := &stubInit{}
	d := newTestDispatcher(init, stubOrder{}, stubCancel{}, stubShip{}, stubChange{}, stubExpand{})
	err := d.Dispatch(context.Background(), cli.ParsedCommand{
		Kind:     cmdname.CommandNameInitData,
		InitData: &usecase.DataInitializationUseCaseReq{},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !init.called {
		t.Fatal("expected InitData to be called")
	}
}

func TestDispatcher_Dispatch_rejectsNilPayload(t *testing.T) {
	t.Parallel()
	d := newTestDispatcher(&stubInit{}, stubOrder{}, stubCancel{}, stubShip{}, stubChange{}, stubExpand{})
	err := d.Dispatch(context.Background(), cli.ParsedCommand{Kind: cmdname.CommandNameOrder})
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDispatcher_Dispatch_orderSuccessPrints(t *testing.T) {
	t.Parallel()
	d := newTestDispatcher(&stubInit{}, stubOrder{}, stubCancel{}, stubShip{}, stubChange{}, stubExpand{})
	err := d.Dispatch(context.Background(), cli.ParsedCommand{
		Kind: cmdname.CommandNameOrder,
		Order: &usecase.OrderUseCaseReq{
			OrderTime:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			OrderNumber: "N1",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDispatcher_Register_duplicate(t *testing.T) {
	t.Parallel()
	d := newTestDispatcher(&stubInit{}, stubOrder{}, stubCancel{}, stubShip{}, stubChange{}, stubExpand{})
	err := d.Register(cmdname.CommandNameOrder, func(context.Context, cli.ParsedCommand) error { return nil })
	if err == nil {
		t.Fatal("expected duplicate registration error")
	}
}

func TestDispatcher_Register_nilHandler(t *testing.T) {
	t.Parallel()
	d := Dispatcher{handlers: make(map[cmdname.CommandName]dispatchFn)}
	err := d.Register(cmdname.CommandNameOrder, nil)
	if err == nil {
		t.Fatal("expected error for nil handler")
	}
}
