package datastore

import (
	"context"
	"testing"

	"theapp/domain/model"
)

func TestProductRepository_Save_and_FindByProductNumber(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name          string
		save          []*model.Product
		findNumber    string
		wantFound     bool
		wantProductNo string
	}{
		{
			name:          "保存後に商品番号で取得できる",
			save:          []*model.Product{model.NewProduct("P-A")},
			findNumber:    "P-A",
			wantFound:     true,
			wantProductNo: "P-A",
		},
		{
			name:          "未登録の商品番号は nil",
			save:          []*model.Product{model.NewProduct("P-X")},
			findNumber:    "P-Y",
			wantFound:     false,
			wantProductNo: "",
		},
		{
			name: "同じ商品番号を上書き保存しても最新で取得できる",
			save: []*model.Product{
				model.NewProduct("P-SAME"),
				model.NewProduct("P-SAME"),
			},
			findNumber:    "P-SAME",
			wantFound:     true,
			wantProductNo: "P-SAME",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := NewProductRepository()
			for _, p := range tt.save {
				if err := repo.Save(ctx, p); err != nil {
					t.Fatalf("Save: %v", err)
				}
			}
			got, err := repo.FindByProductNumber(ctx, tt.findNumber)
			if err != nil {
				t.Fatalf("FindByProductNumber: %v", err)
			}
			if !tt.wantFound {
				if got != nil {
					t.Fatalf("got %+v, want nil", got)
				}
				return
			}
			if got == nil || got.ProductNumber != tt.wantProductNo {
				t.Fatalf("got %+v, want ProductNumber %q", got, tt.wantProductNo)
			}
		})
	}
}
