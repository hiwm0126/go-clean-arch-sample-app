package repository

import (
	"context"
	"theapp/domain/model"
)

type ProductRepository interface {
	Save(ctx context.Context, product *model.Product) error
	FindByProductNumber(ctx context.Context, productNumber string) (*model.Product, error)
}
