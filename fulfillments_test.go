package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/assertions"
	"github.com/MOHC-LTD/shopify"
)

// Tests that fulfillment can be built correctly when date fields are not nil
func TestFulfillmentDTO_ToShopify(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var fulfillmentDTO = FulfillmentDTO{
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}

	fulfillment := fulfillmentDTO.ToShopify()

	if !fulfillment.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, fulfillment.CreatedAt)
	}

	if !fulfillment.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, fulfillment.UpdatedAt)
	}
}

// Tests that fulfillment event can be built correctly when date fields are nil
func TestFulfillmentDTO_ToShopifyEmptyFields(t *testing.T) {
	var createdAt *time.Time
	var updatedAt *time.Time

	var fulfillmentDTO = FulfillmentDTO{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	fulfillment := fulfillmentDTO.ToShopify()

	if !fulfillment.CreatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, createdAt, fulfillment.CreatedAt)
	}

	if !fulfillment.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, fulfillment.UpdatedAt)
	}
}

// Tests that a fulfillment DTO can be built when date fields are not nil
func TestBuildFulfilmentDTO(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var fulfillment = shopify.Fulfillment{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	fulfillmentDTO := BuildFulfilmentDTO(fulfillment)

	if !fulfillmentDTO.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, fulfillmentDTO.CreatedAt)
	}

	if !fulfillmentDTO.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, fulfillmentDTO.UpdatedAt)
	}
}

// Tests that a fulfillment DTO can be built when date fields are nil
func TestBuildFulfilmentDTOEmptyTimes(t *testing.T) {
	var createdAt time.Time
	var updatedAt time.Time

	var fulfillment = shopify.Fulfillment{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	fulfillmentDTO := BuildFulfilmentDTO(fulfillment)

	if fulfillmentDTO.CreatedAt != nil {
		assertions.ValueAssertionFailure(t, createdAt, fulfillmentDTO.CreatedAt)
	}

	if fulfillmentDTO.UpdatedAt != nil {
		assertions.ValueAssertionFailure(t, updatedAt, fulfillmentDTO.UpdatedAt)
	}
}
