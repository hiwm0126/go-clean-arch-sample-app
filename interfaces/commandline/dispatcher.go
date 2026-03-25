package commandline

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"theapp/constants"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

// Dispatcher は ParsedCommand をユースケースへ一意に振り分け、CLI 向け標準出力を行う。
type Dispatcher struct {
	initData usecase.DataInitializationUseCase
	order    usecase.OrderUseCase
	cancel   usecase.CancelUseCase
	ship     usecase.ShippingUseCase
	change   usecase.ChangeUseCase
	expand   usecase.ExpandUseCase
}

// NewDispatcher コンストラクタ
func NewDispatcher(
	initData usecase.DataInitializationUseCase,
	order usecase.OrderUseCase,
	cancel usecase.CancelUseCase,
	ship usecase.ShippingUseCase,
	change usecase.ChangeUseCase,
	expand usecase.ExpandUseCase,
) *Dispatcher {
	return &Dispatcher{
		initData: initData,
		order:    order,
		cancel:   cancel,
		ship:     ship,
		change:   change,
		expand:   expand,
	}
}

// Dispatch は1コマンドを処理する。
func (d *Dispatcher) Dispatch(ctx context.Context, cmd ParsedCommand) error {
	switch cmd.Kind {
	case cmdname.CommandNameInitData:
		if cmd.InitData == nil {
			return errors.New("commandline: missing InitData payload")
		}
		return d.initData.InitData(ctx, cmd.InitData)

	case cmdname.CommandNameOrder:
		if cmd.Order == nil {
			return errors.New("commandline: missing Order payload")
		}
		res, err := d.order.Order(ctx, cmd.Order)
		if err != nil {
			return err
		}
		if res.IsError {
			fmt.Printf("%s Ordered %s Error: the number of available shipments has been exceeded.\n",
				res.OrderTime.Format(constants.DateTimeFormat), res.OrderNumber)
		} else {
			fmt.Printf("%s Ordered %s\n", res.OrderTime.Format(constants.DateTimeFormat), res.OrderNumber)
		}
		return nil

	case cmdname.CommandNameCancel:
		if cmd.Cancel == nil {
			return errors.New("commandline: missing Cancel payload")
		}
		res, err := d.cancel.Cancel(ctx, cmd.Cancel)
		if err != nil {
			return err
		}
		fmt.Printf("%s Canceled %s\n", res.CancelTime.Format(constants.DateTimeFormat), res.OrderNumber)
		return nil

	case cmdname.CommandNameShip:
		if cmd.Ship == nil {
			return errors.New("commandline: missing Ship payload")
		}
		res, err := d.ship.Ship(ctx, cmd.Ship)
		if err != nil {
			return err
		}
		fmt.Printf("%s Shipped %v Orders\n", res.ShipmentRequestTime.Format(constants.DateTimeFormat), len(res.Orders))
		var orderNumbers []string
		for _, order := range res.Orders {
			orderNumbers = append(orderNumbers, order.OrderNumber)
		}
		slices.Sort(orderNumbers)
		fmt.Println(strings.Join(orderNumbers, " "))
		return nil

	case cmdname.CommandNameChange:
		if cmd.Change == nil {
			return errors.New("commandline: missing Change payload")
		}
		res, err := d.change.Change(ctx, cmd.Change)
		if err != nil {
			return err
		}
		if res.IsError {
			fmt.Printf("%s Changed %s Error: the number of available shipments has been exceeded.\n",
				res.RequestTime.Format(constants.DateTimeFormat), res.OrderNumber)
		} else {
			fmt.Printf("%s Changed %s\n", res.RequestTime.Format(constants.DateTimeFormat), res.OrderNumber)
		}
		return nil

	case cmdname.CommandNameExpand:
		if cmd.Expand == nil {
			return errors.New("commandline: missing Expand payload")
		}
		res, err := d.expand.Expand(ctx, cmd.Expand)
		if err != nil {
			return err
		}
		fmt.Printf("%s Expanded\n", res.ExpandRequestTime.Format(constants.DateTimeFormat))
		return nil

	default:
		return fmt.Errorf("commandline: unhandled command %q", cmd.Kind)
	}
}
