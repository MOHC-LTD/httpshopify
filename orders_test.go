package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/assertions"
	"github.com/MOHC-LTD/shopify"
)

// Tests order can be built when date fields are not nil
func TestOrderDTO_ToShopify(t *testing.T) {
	createdAt := time.Now()
	closedAt := time.Now()
	processedAt := time.Now()
	updatedAt := time.Now()

	var orderDTO = OrderDTO{
		CreatedAt:   &createdAt,
		ClosedAt:    &closedAt,
		ProcessedAt: &processedAt,
		UpdatedAt:   &updatedAt,
	}

	order := orderDTO.ToShopify()

	if !order.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, order.CreatedAt)
	}

	if !order.ClosedAt.Equal(closedAt) {
		assertions.ValueAssertionFailure(t, closedAt, order.ClosedAt)
	}

	if !order.ProcessedAt.Equal(processedAt) {
		assertions.ValueAssertionFailure(t, processedAt, order.ProcessedAt)
	}

	if !order.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, order.UpdatedAt)
	}
}

// Tests order can be built when date fields are nil
func TestOrderDTO_ToShopifyEmptyTimes(t *testing.T) {
	var createdAt *time.Time
	var closedAt *time.Time
	var processedAt *time.Time
	var updatedAt *time.Time

	var orderDTO = OrderDTO{
		CreatedAt:   createdAt,
		ClosedAt:    closedAt,
		ProcessedAt: processedAt,
		UpdatedAt:   updatedAt,
	}

	order := orderDTO.ToShopify()

	if !order.CreatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, createdAt, order.CreatedAt)
	}

	if !order.ClosedAt.IsZero() {
		assertions.ValueAssertionFailure(t, closedAt, order.ClosedAt)
	}

	if !order.ProcessedAt.IsZero() {
		assertions.ValueAssertionFailure(t, processedAt, order.ProcessedAt)
	}

	if !order.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, order.UpdatedAt)
	}
}

// Tests that a order DTO can be built when date fields are not nil
func TestBuildOrderDTO(t *testing.T) {
	createdAt := time.Now()
	closedAt := time.Now()
	processedAt := time.Now()
	updatedAt := time.Now()

	var order = shopify.Order{
		CreatedAt:   createdAt,
		ClosedAt:    closedAt,
		ProcessedAt: processedAt,
		UpdatedAt:   updatedAt,
	}

	orderDTO := BuildOrderDTO(order)

	if !orderDTO.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, orderDTO.CreatedAt)
	}

	if !orderDTO.ClosedAt.Equal(closedAt) {
		assertions.ValueAssertionFailure(t, closedAt, orderDTO.ClosedAt)
	}

	if !orderDTO.ProcessedAt.Equal(processedAt) {
		assertions.ValueAssertionFailure(t, processedAt, orderDTO.ProcessedAt)
	}

	if !orderDTO.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, orderDTO.UpdatedAt)
	}
}

// Tests that a fulfillment DTO can be built when date fields are nil
func TestBuildOrderDTOEmptyTimes(t *testing.T) {
	var createdAt time.Time
	var closedAt time.Time
	var processedAt time.Time
	var updatedAt time.Time

	var order = shopify.Order{
		CreatedAt:   createdAt,
		ClosedAt:    closedAt,
		ProcessedAt: processedAt,
		UpdatedAt:   updatedAt,
	}

	orderDTO := BuildOrderDTO(order)

	if orderDTO.CreatedAt != nil {
		assertions.ValueAssertionFailure(t, createdAt, orderDTO.CreatedAt)
	}

	if orderDTO.ClosedAt != nil {
		assertions.ValueAssertionFailure(t, closedAt, orderDTO.ClosedAt)
	}

	if orderDTO.ProcessedAt != nil {
		assertions.ValueAssertionFailure(t, processedAt, orderDTO.ProcessedAt)
	}

	if orderDTO.UpdatedAt != nil {
		assertions.ValueAssertionFailure(t, updatedAt, orderDTO.UpdatedAt)
	}
}
