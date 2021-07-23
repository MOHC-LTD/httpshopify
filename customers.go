package httpshopify

import (
	"time"

	"github.com/MOHC-LTD/shopify"
)

// CustomerDTO represents a Shopify customer in HTTP requests and responses
type CustomerDTO struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto CustomerDTO) ToShopify() shopify.Customer {
	return shopify.Customer{
		ID:        dto.ID,
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

// BuildCustomerDTO converts a Shopify customer to the DTO equivalent
func BuildCustomerDTO(customer shopify.Customer) CustomerDTO {
	return CustomerDTO{
		ID:        customer.ID,
		Email:     customer.Email,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}
}
