package model_test

import (
	"testing"

	"theapp/domain/model"
)

func TestNewProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		productNumber   string
		wantProductNum  string
	}{
		{
			name:           "商品番号がそのまま設定される",
			productNumber:  "P-001",
			wantProductNum: "P-001",
		},
		{
			name:           "空文字でも生成できる",
			productNumber:  "",
			wantProductNum: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := model.NewProduct(tt.productNumber)
			if got.ProductNumber != tt.wantProductNum {
				t.Fatalf("ProductNumber = %q, want %q", got.ProductNumber, tt.wantProductNum)
			}
		})
	}
}
