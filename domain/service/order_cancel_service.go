package service

import (
	"errors"
	"github.com/hiwm0126/internship_27_test/domain/model"
	"github.com/hiwm0126/internship_27_test/domain/repository"
	"time"
)

type OrderCancelService interface {
	Execute(order *model.Order, cancelTime time.Time) error
}

type orderCancelService struct {
	orderRepo     repository.OrderRepository
	orderItemRepo repository.OrderItemRepository
}

func NewOrderCancelService(
	orderRepo repository.OrderRepository,
	orderItemRepo repository.OrderItemRepository,
) OrderCancelService {
	return &orderCancelService{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
	}
}

func (o *orderCancelService) Execute(order *model.Order, cancelTime time.Time) error {

	// ステータス変更可能な日付かチェック
	if !order.CanChangeStatusDate(cancelTime) {
		return errors.New("this cancel date is not allowed")
	}

	// ステータス変更可能な状態かチェック
	if !order.CanChangeStatus() {
		return errors.New("this order status is not allowed")
	}

	// ステータスをキャンセル済みに変更
	err := o.orderRepo.UpdateStatus(order.OrderNumber, model.OrderStatusCanceled)
	if err != nil {
		return err
	}

	// 注文に紐づく商品を削除
	err = o.orderItemRepo.DeleteByOrderNumber(order.OrderNumber)
	if err != nil {
		return err
	}

	return nil
}
