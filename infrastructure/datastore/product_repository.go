package datastore

import (
	"context"
	"example.com/internship_27_test/domain/model"
	"example.com/internship_27_test/domain/repository"
)

type Product struct {
	ID            int
	ProductNumber string
}

type productRepository struct {
	index        int
	productTable map[string]*Product
}

func NewProductRepository() repository.ProductRepository {
	return &productRepository{
		index:        1,
		productTable: make(map[string]*Product),
	}
}

// SaveProduct 商品マスタ情報を保存する
func (r *productRepository) Save(_ context.Context, product *model.Product) error {
	data := &Product{
		ID:            r.index,
		ProductNumber: product.ProductNumber,
	}
	r.productTable[data.ProductNumber] = data
	r.index++
	return nil
}

// FindByProductNumber 商品マスタ情報を取得する
func (r *productRepository) FindByProductNumber(_ context.Context, productNumber string) (*model.Product, error) {
	data, ok := r.productTable[productNumber]
	if !ok {
		return nil, nil
	}
	return &model.Product{
		ProductNumber: data.ProductNumber,
	}, nil
}
