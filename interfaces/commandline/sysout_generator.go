package commandline

import (
	"fmt"
	"slices"
	"strings"
	"theapp/constants"
	"theapp/usecase"
)

type SysOutGenerator struct{}

func NewSysOutGenerator() *SysOutGenerator {
	return &SysOutGenerator{}
}

// Generate レスポンスのtypeに応じて、標準出力を生成する
func (g *SysOutGenerator) Generate(res interface{}) {
	switch v := res.(type) {
	case *usecase.OrderUseCaseRes:
		if v.IsError {
			fmt.Printf("%s Ordered %s Error: the number of available shipments has been exceeded.\n", v.OrderTime.Format(constants.DateTimeFormat), v.OrderNumber)
			return
		}
		fmt.Printf("%s Ordered %s\n", v.OrderTime.Format(constants.DateTimeFormat), v.OrderNumber)
	case *usecase.CancelUseCaseRes:
		fmt.Printf("%s Canceled %s\n", v.CancelTime.Format(constants.DateTimeFormat), v.OrderNumber)
	case *usecase.ShippingUseCaseRes:
		g.generateShippingUseCaseSysOut(v)
	case *usecase.ChangeUseCaseRes:
		if v.IsError {
			fmt.Printf("%s Changed %s Error: the number of available shipments has been exceeded.\n", v.RequestTime.Format(constants.DateTimeFormat), v.OrderNumber)
			return
		}
		fmt.Printf("%s Changed %s\n", v.RequestTime.Format(constants.DateTimeFormat), v.OrderNumber)
	case *usecase.ExpandUseCaseRes:
		fmt.Printf("%s Expanded\n", v.ExpandRequestTime.Format(constants.DateTimeFormat))
	case error:
		fmt.Println("")
	default:
		println("Unknown type")
	}

}

// generateShippingUseCaseSysOut ShippingUseCaseResの標準出力を生成する
func (g *SysOutGenerator) generateShippingUseCaseSysOut(res *usecase.ShippingUseCaseRes) {
	fmt.Printf("%s Shipped %v Orders\n", res.ShipmentRequestTime.Format(constants.DateTimeFormat), len(res.Orders))
	var orderNumbers []string
	for _, order := range res.Orders {
		orderNumbers = append(orderNumbers, order.OrderNumber)
	}
	slices.Sort(orderNumbers)
	fmt.Println(strings.Join(orderNumbers, " "))
}
