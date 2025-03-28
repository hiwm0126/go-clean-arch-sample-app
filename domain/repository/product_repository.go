package repository

import "github.com/hiwm0126/internship_27_test/domain/model"

type ProductRepository interface {
	Save(product *model.Product) error
	FindByProductNumber(productNumber string) (*model.Product, error)
}
