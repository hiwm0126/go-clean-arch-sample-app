package repository

import (
	"context"
	"example.com/internship_27_test/domain/model"
)

type ProductRepository interface {
	Save(ctx context.Context, product *model.Product) error
	FindByProductNumber(ctx context.Context, productNumber string) (*model.Product, error)
}
