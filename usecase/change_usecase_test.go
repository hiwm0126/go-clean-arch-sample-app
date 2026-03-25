package usecase

import (
	"context"
	"errors"
	"testing"

	"theapp/domain/model"
	"theapp/domain/repository"
)

type errOrderRepo struct{}

func (errOrderRepo) Save(_ context.Context, _ *model.Order) error { return nil }
func (errOrderRepo) Find(_ context.Context, _ int) (*model.Order, error) { return nil, nil }
func (errOrderRepo) FindByOrderNumber(_ context.Context, _ string) (*model.Order, error) {
	return nil, errors.New("repository unavailable")
}
func (errOrderRepo) GetOrdersByShipmentDueDate(_ context.Context, _ string) ([]*model.Order, error) {
	return nil, nil
}
func (errOrderRepo) UpdateStatus(_ context.Context, _ string, _ model.OrderStatus) error { return nil }

var _ repository.OrderRepository = errOrderRepo{}

func TestChangeUseCase_Change_repositoryError(t *testing.T) {
	t.Parallel()
	c := NewChangeUseCase(
		errOrderRepo{},
		nil,
		nil,
		nil,
		nil,
	)
	_, err := c.Change(context.Background(), &ChangeUseCaseReq{OrderNumber: "X"})
	if err == nil || err.Error() != "repository unavailable" {
		t.Fatalf("Change() err = %v, want repository unavailable", err)
	}
}
