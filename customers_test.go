package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/assertions"
	"github.com/MOHC-LTD/shopify"
)

// Tests customer can be built when date fields are not nil
func TestCustomerDTO_ToShopify(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var customerDTO = CustomerDTO{
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}

	customer := customerDTO.ToShopify()

	if !customer.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, customer.CreatedAt)
	}

	if !customer.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, customer.UpdatedAt)
	}
}

// Tests customer can be built when date fields are nil
func TestCustomerDTO_ToShopifyEmptyTimes(t *testing.T) {
	var createdAt *time.Time
	var updatedAt *time.Time

	var customerDTO = CustomerDTO{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	customer := customerDTO.ToShopify()

	if !customer.CreatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, createdAt, customer.CreatedAt)
	}

	if !customer.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, customer.UpdatedAt)
	}
}

// Tests that a customer DTO can be built when date fields are not nil
func TestBuildCustomerDTO(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var customer = shopify.Customer{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	customerDTO := BuildCustomerDTO(customer)

	if !customerDTO.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, customerDTO.CreatedAt)
	}

	if !customerDTO.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, customerDTO.UpdatedAt)
	}
}

// Tests that a customer DTO can be built when date fields are nil
func TestBuildCustomerDTOEmptyTimes(t *testing.T) {
	var createdAt time.Time
	var updatedAt time.Time

	var customer = shopify.Customer{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	customerDTO := BuildCustomerDTO(customer)

	if customerDTO.CreatedAt != nil {
		assertions.ValueAssertionFailure(t, createdAt, customerDTO.CreatedAt)
	}

	if customerDTO.UpdatedAt != nil {
		assertions.ValueAssertionFailure(t, updatedAt, customerDTO.UpdatedAt)
	}
}
