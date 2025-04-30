package commandline

import (
	"errors"
	"strconv"
	"theapp/constants"
	"theapp/domain/model"
	"time"

	"theapp/usecase"
)

type Parser interface {
	Parse(rawArgs [][]string) ([]interface{}, error)
}

type argsParser struct{}

func NewParser() Parser {
	return &argsParser{}
}

type QueryName = string

const (
	ORDER     QueryName = "ORDER"
	CANCEL    QueryName = "CANCEL"
	SHIP      QueryName = "SHIP"
	CHANGE    QueryName = "CHANGE"
	EXPAND    QueryName = "EXPAND"
	INIT_DATA QueryName = "INIT_DATA"
)

var validQueries = map[QueryName]bool{
	ORDER:  true,
	CANCEL: true,
	SHIP:   true,
	CHANGE: true,
	EXPAND: true,
}

func (ap *argsParser) Parse(rawArgs [][]string) ([]interface{}, error) {
	var reqParamList []interface{}
	var queryArgs [][]string
	var queryName = INIT_DATA

	for _, arg := range rawArgs {
		if len(arg) == 0 {
			continue // 空のスライスはスキップ
		}

		// クエリネームが含まれているかどうか
		// 含まれている場合、それまでのクエリの引数をリクエストパラメータに変換
		if ap.existsQueryName(arg[0]) {
			if err := ap.appendParam(&reqParamList, queryName, queryArgs); err != nil {
				return nil, err
			}

			// クエリネームを更新して、初期化
			queryName = QueryName(arg[0])
			queryArgs = nil
		}
		queryArgs = append(queryArgs, arg)
	}

	if err := ap.appendParam(&reqParamList, queryName, queryArgs); err != nil {
		return nil, err
	}

	return reqParamList, nil
}

func (ap *argsParser) existsQueryName(arg string) bool {
	_, exists := validQueries[arg]
	return exists
}

func (ap *argsParser) createParam(queryName QueryName, args [][]string) (reqParam interface{}, err error) {

	switch queryName {
	case ORDER:
		reqParam, err = ap.createOrderParam(args)
	case CANCEL:
		reqParam, err = ap.createCancelParam(args)
	case SHIP:
		reqParam, err = ap.createShipParam(args)
	case CHANGE:
		reqParam, err = ap.createChangeParam(args)
	case EXPAND:
		reqParam, err = ap.createExpandParam(args)
	case INIT_DATA:
		reqParam, err = ap.createInitDataParam(args)
	}

	return
}

func (ap *argsParser) createOrderParam(args [][]string) (*usecase.OrderUseCaseReq, error) {
	var (
		orderTimeIndexFirstLine        = 1
		orderNumberIndexSecondLine     = 0
		shipmentDueDateIndexSecondLine = 1
		productNumberIndexItemInfo     = 0
		itemQuantityItemInfo           = 1
	)

	firstLine := args[0]
	secondLine := args[1]

	// 商品情報のリストをargsから抽出
	itemInfoList := args[2:]

	itemInfos := make(map[string]int)
	for _, itemInfo := range itemInfoList {
		itemInfos[itemInfo[productNumberIndexItemInfo]] = ap.parseStringToInt(itemInfo[itemQuantityItemInfo])
	}

	return &usecase.OrderUseCaseReq{
		OrderTime:       ap.parseStringToTime(firstLine[orderTimeIndexFirstLine]),
		OrderNumber:     secondLine[orderNumberIndexSecondLine],
		ShipmentDueDate: secondLine[shipmentDueDateIndexSecondLine],
		ItemInfos:       itemInfos,
	}, nil
}

func (ap *argsParser) createCancelParam(args [][]string) (*usecase.CancelUseCaseReq, error) {
	var (
		cancelTimeIndexSecondLine  = 1
		orderNumberIndexSecondLine = 0
	)

	if len(args) != 2 {
		return nil, errors.New("invalid number of arguments")
	}
	firstLine := args[0]
	secondLine := args[1]

	return &usecase.CancelUseCaseReq{
		OrderNumber: secondLine[orderNumberIndexSecondLine],
		CancelTime:  ap.parseStringToTime(firstLine[cancelTimeIndexSecondLine]),
	}, nil
}

func (ap *argsParser) createShipParam(args [][]string) (*usecase.ShippingUseCaseReq, error) {
	var (
		shipmentRequestTimeIndexFirstLine = 1
		maxNumberOfParamLines             = 1
	)

	if len(args) != maxNumberOfParamLines {
		return nil, errors.New("invalid number of arguments")
	}

	return &usecase.ShippingUseCaseReq{
		ShipmentRequestTime: ap.parseStringToTime(args[0][shipmentRequestTimeIndexFirstLine]),
	}, nil
}

func (ap *argsParser) createChangeParam(args [][]string) (*usecase.ChangeUseCaseReq, error) {
	var (
		requestTimeIndexFirstLine        = 1
		orderNumberIndexSecondLine       = 0
		changeRequestDateIndexSecondLine = 1
		maxNumberOfParamLines            = 2
	)

	if len(args) != maxNumberOfParamLines {
		return nil, errors.New("invalid number of arguments")
	}

	firstLine := args[0]
	secondLine := args[1]

	return &usecase.ChangeUseCaseReq{
		OrderNumber:       secondLine[orderNumberIndexSecondLine],
		ChangeRequestDate: secondLine[changeRequestDateIndexSecondLine],
		RequestTime:       ap.parseStringToTime(firstLine[requestTimeIndexFirstLine]),
	}, nil
}

func (ap *argsParser) createExpandParam(args [][]string) (*usecase.ExpandUseCaseReq, error) {
	var (
		expandRequestTimeIndexFirstLine = 1
		FromIndexSecondLine             = 0
		toIndexSecondLine               = 1
		quantityIndexSecondLine         = 2
		maxNumberOfParamLines           = 2
	)

	if len(args) != maxNumberOfParamLines {
		return nil, errors.New("invalid number of arguments")
	}

	firstLine := args[0]
	secondLine := args[1]

	return &usecase.ExpandUseCaseReq{
		ExpandRequestTime: ap.parseStringToTime(firstLine[expandRequestTimeIndexFirstLine]),
		From:              secondLine[FromIndexSecondLine],
		To:                secondLine[toIndexSecondLine],
		Quantity:          ap.parseStringToInt(secondLine[quantityIndexSecondLine]),
	}, nil
}

func (ap *argsParser) createInitDataParam(args [][]string) (*usecase.DataInitializationUseCaseReq, error) {
	var (
		numOfProductIndexFirstLine             = 0
		shipmentLimitThresholdIndexFirstLine   = 1
		shipmentAcceptablePeriodIndexFirstLine = 2
		numOfQueryIndexForthLine               = 0
		maxNumberOfParamLines                  = 4
	)

	if len(args) != maxNumberOfParamLines {
		return nil, errors.New("invalid number of arguments")
	}

	firstLine := args[0]
	secondLine := args[1]
	thirdLine := args[2]
	forthLine := args[3]

	return &usecase.DataInitializationUseCaseReq{
		NumOfProduct:             ap.parseStringToInt(firstLine[numOfProductIndexFirstLine]),
		ShipmentLimitThreshold:   ap.parseStringToInt(firstLine[shipmentLimitThresholdIndexFirstLine]),
		ShipmentAcceptablePeriod: ap.parseStringToInt(firstLine[shipmentAcceptablePeriodIndexFirstLine]),
		ProductNumberList:        secondLine,
		ShipmentLimitFlags:       ap.parseToShipmentLimitFlags(thirdLine),
		NumOfQuery:               ap.parseStringToInt(forthLine[numOfQueryIndexForthLine]),
	}, nil
}

func (ap *argsParser) parseStringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

func (ap *argsParser) parseToShipmentLimitFlags(flagSlice []string) map[model.DayOfWeek]bool {
	// 一週間は7日間
	if len(flagSlice) != 7 {
		return nil
	}

	shipmentLimitFlags := make(map[model.DayOfWeek]bool, 7)
	for i, flagStr := range flagSlice {
		dayOfWeek := model.DayOfWeek(i)
		flag := ap.parseStringToInt(flagStr)
		shipmentLimitFlags[dayOfWeek] = flag == 1
	}
	return shipmentLimitFlags
}

func (ap *argsParser) parseStringToTime(str string) time.Time {
	result, _ := time.Parse(constants.DateTimeFormat, str)
	return result
}

func (ap *argsParser) appendParam(reqParamList *[]interface{}, queryName QueryName, queryArgs [][]string) error {
	if len(queryArgs) == 0 {
		return nil
	}
	reqParam, err := ap.createParam(queryName, queryArgs)
	if err != nil {
		return err
	}
	*reqParamList = append(*reqParamList, reqParam)
	return nil
}
