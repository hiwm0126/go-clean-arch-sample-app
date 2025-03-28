package model

type Product struct {
	ProductNumber string
}

func NewProduct(productNumber string) *Product {
	return &Product{
		ProductNumber: productNumber,
	}
}
