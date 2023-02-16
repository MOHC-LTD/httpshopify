package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/assertions"
)

// Tests that fulfillment event can be built correctly when date fields are not nil
func TestFulfillmentEventDTO_ToShopify(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var fulfillmentEventDTO = FulfillmentEventDTO{
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}

	fulfillmentEvent := fulfillmentEventDTO.ToShopify()

	if !fulfillmentEvent.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, fulfillmentEvent.CreatedAt)
	}

	if !fulfillmentEvent.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, fulfillmentEvent.UpdatedAt)
	}
}

// Tests that fulfillment event can be built correctly when date fields are nil
func TestFulfillmentEventDTO_ToShopifyEmptyFields(t *testing.T) {
	var createdAt *time.Time
	var updatedAt *time.Time

	var fulfillmentEventDTO = FulfillmentEventDTO{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	fulfillmentEvent := fulfillmentEventDTO.ToShopify()

	if !fulfillmentEvent.CreatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, createdAt, fulfillmentEvent.CreatedAt)
	}

	if !fulfillmentEvent.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, fulfillmentEvent.UpdatedAt)
	}
}
