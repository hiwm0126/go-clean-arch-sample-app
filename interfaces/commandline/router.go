package commandline

import "github.com/hiwm0126/internship_27_test/usecase"

type Router interface {
	Routing(args [][]string) error
}
type router struct {
	controller      *Controller
	sysOutGenerator *SysOutGenerator
}

func NewRouter(con *Controller, sysOutGenerator *SysOutGenerator) Router {
	return &router{
		controller: con,
	}
}

func (r *router) Routing(args [][]string) error {
	// リクエストパラメーターを解析
	parser := NewParser()
	reqParamList, err := parser.Parse(args)
	if err != nil {
		return err
	}

	// リクエストパラメーターに応じた処理を実行
	for _, reqParam := range reqParamList {
		switch param := reqParam.(type) {
		case *usecase.DataInitializationUseCaseReq:
			err := r.controller.InitData(param)
			if err != nil {
				return err
			}
		case *usecase.OrderUseCaseReq:
			res, err := r.controller.Order(param)
			if err != nil {
				return err
			}
			r.sysOutGenerator.Generate(res)
		case *usecase.CancelUseCaseReq:
			res, err := r.controller.Cancel(param)
			if err != nil {
				return err
			}
			r.sysOutGenerator.Generate(res)
		case *usecase.ShippingUseCaseReq:
			res, err := r.controller.Ship(param)
			if err != nil {
				return err
			}
			r.sysOutGenerator.Generate(res)
		case *usecase.ChangeUseCaseReq:
			res, err := r.controller.Change(param)
			if err != nil {
				return err
			}
			r.sysOutGenerator.Generate(res)
		case *usecase.ExpandUseCaseReq:
			res, err := r.controller.Expand(param)
			if err != nil {
				return err
			}
			r.sysOutGenerator.Generate(res)
		}
	}
	return nil
}
