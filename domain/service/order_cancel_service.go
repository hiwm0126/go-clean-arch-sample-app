package service

import (
	"context"
	"errors"
	"example.com/internship_27_test/domain/model"
	"example.com/internship_27_test/domain/repository"
	"time"
)

type OrderCancelService interface {
	Execute(ctx context.Context, order *model.Order, cancelTime time.Time) error
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

func (o *orderCancelService) Execute(ctx context.Context, order *model.Order, cancelTime time.Time) error {

	// ステータス変更可能な日付かチェック
	if !order.CanChangeStatusDate(cancelTime) {
		return errors.New("this cancel date is not allowed")
	}

	// ステータス変更可能な状態かチェック
	if !order.CanChangeStatus() {
		return errors.New("this order status is not allowed")
	}

	// ステータスをキャンセル済みに変更
	err := o.orderRepo.UpdateStatus(ctx, order.OrderNumber, model.OrderStatusCanceled)
	if err != nil {
		return err
	}

	// 注文に紐づく商品を削除
	err = o.orderItemRepo.DeleteByOrderNumber(ctx, order.OrderNumber)
	if err != nil {
		return err
	}

	return nil
}
