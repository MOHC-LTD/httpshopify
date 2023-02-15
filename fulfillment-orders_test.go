package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/assertions"
)

// Tests that fulfillment order can be built correctly when date fields are not nil
func TestFulfillmentOrderDTO_ToShopify(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var fulfillmentOrderDTO = FulfillmentOrderDTO{
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}

	fulfillmentOrder := fulfillmentOrderDTO.ToShopify()

	if !fulfillmentOrder.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, fulfillmentOrder.CreatedAt)
	}

	if !fulfillmentOrder.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, fulfillmentOrder.UpdatedAt)
	}
}

// Tests that fulfillment order can be built correctly when date fields are nil
func TestFulfillmentOrderDTO_ToShopifyEmptyFields(t *testing.T) {
	var createdAt *time.Time
	var updatedAt *time.Time

	var fulfillmentOrderDTO = FulfillmentOrderDTO{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	fulfillmentOrder := fulfillmentOrderDTO.ToShopify()

	if !fulfillmentOrder.CreatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, createdAt, fulfillmentOrder.CreatedAt)
	}

	if !fulfillmentOrder.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, fulfillmentOrder.UpdatedAt)
	}
}
