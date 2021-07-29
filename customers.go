package httpshopify

import (
	"time"

	"github.com/MOHC-LTD/shopify"
)

// CustomerDTO represents a Shopify customer in HTTP requests and responses
type CustomerDTO struct {
	ID        int64      `json:"id,omitempty"`
	Email     string     `json:"email,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto CustomerDTO) ToShopify() shopify.Customer {
	var createdAt time.Time
	if !dto.CreatedAt.IsZero() {
		createdAt = *dto.CreatedAt
	}

	var updatedAt time.Time
	if !dto.UpdatedAt.IsZero() {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.Customer{
		ID:        dto.ID,
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// BuildCustomerDTO converts a Shopify customer to the DTO equivalent
func BuildCustomerDTO(customer shopify.Customer) CustomerDTO {
	var createdAt *time.Time
	if !customer.CreatedAt.IsZero() {
		createdAt = &customer.CreatedAt
	}

	var updatedAt *time.Time
	if !customer.UpdatedAt.IsZero() {
		updatedAt = &customer.UpdatedAt
	}

	customerDTO := CustomerDTO{
		ID:        customer.ID,
		Email:     customer.Email,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return customerDTO
}
