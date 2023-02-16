package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/assertions"
)

// Tests inventory level can be built when date fields are not nil
func TestInventoryLevelDTO_ToShopify(t *testing.T) {
	updatedAt := time.Now()

	var inventoryLevelDTO = InventoryLevelDTO{
		UpdatedAt: &updatedAt,
	}

	inventoryLevel := inventoryLevelDTO.ToShopify()

	if !inventoryLevel.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, inventoryLevel.UpdatedAt)
	}
}

// Tests inventory level can be built when date fields are nil
func TestInventoryLevelDTO_ToShopifyEmptyTimes(t *testing.T) {
	var updatedAt *time.Time

	var inventoryLevelDTO = InventoryLevelDTO{
		UpdatedAt: updatedAt,
	}

	inventoryLevel := inventoryLevelDTO.ToShopify()

	if !inventoryLevel.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, inventoryLevel.UpdatedAt)
	}
}
