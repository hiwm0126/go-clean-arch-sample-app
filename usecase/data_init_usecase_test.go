package usecase

import (
	"context"
	"strings"
	"testing"

	"theapp/domain/model"
	"theapp/infrastructure/datastore"
)

func TestDataInitializationUseCase_InitData_productCountMismatch(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	u := NewImportDataUseCase(
		datastore.NewProductRepository(),
		datastore.NewShipmentLimitRepository(),
		datastore.NewShippingAcceptablePeriodRepository(),
	)
	err := u.InitData(ctx, &DataInitializationUseCaseReq{
		NumOfProduct:             2,
		ShipmentLimitThreshold:   1,
		ShipmentAcceptablePeriod: 1,
		ProductNumberList:        []string{"only-one"},
		ShipmentLimitFlags: map[model.DayOfWeek]bool{
			model.Sunday: true,
		},
		NumOfQuery: 0,
	})
	if err == nil || !strings.Contains(err.Error(), "product count mismatch") {
		t.Fatalf("InitData() err = %v, want product count mismatch", err)
	}
}
