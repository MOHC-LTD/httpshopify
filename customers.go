package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/http"

	"github.com/MOHC-LTD/shopify"
)

// CustomerDTO represents a Shopify customer in HTTP requests and responses
type CustomerDTO struct {
	ID        int64      `json:"id,omitempty"`
	Email     string     `json:"email,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto CustomerDTO) ToShopify() shopify.Customer {
	var createdAt time.Time
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	var updatedAt time.Time
	if dto.UpdatedAt != nil {
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
		Phone:     customer.Phone,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return customerDTO
}

type customerRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newCustomerRepository(client http.Client, createURL func(endpoint string) string) customerRepository {
	return customerRepository{
		client,
		createURL,
	}
}

func (c customerRepository) Update(customer shopify.Customer) (shopify.Customer, error) {
	// Map to DTO
	customerDTO := BuildCustomerDTO(customer)

	request := struct {
		CustomerDTO `json:"customer"`
	}{
		customerDTO,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return shopify.Customer{}, err
	}

	url := c.createURL(fmt.Sprintf("customers/%d.json", customer.ID))

	respBody, _, err := c.client.Put(url, body, nil)
	if err != nil {
		switch err.(type) {
		case http.ErrHTTP:
			var responseError ResponseError
			e := json.Unmarshal([]byte(err.(http.ErrHTTP).Body), &responseError)
			if e != nil {
				return shopify.Customer{}, err
			}

			return shopify.Customer{}, responseError.getType(err)
		default:
			return shopify.Customer{}, err
		}
	}

	var response struct {
		Customer CustomerDTO `json:"customer"`
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.Customer{}, err
	}

	return response.Customer.ToShopify(), nil
}

// ResponseError is used to store the error response for a customer
type ResponseError struct {
	Errors Fields `json:"errors"`
}

// Fields holds all the possible fields for an error
type Fields struct {
	Email []interface{} `json:"email"`
	Phone []interface{} `json:"phone"`
}

func (e ResponseError) getType(err error) error {
	if len(e.Errors.Email) > 0 {
		return ErrCustomerEmail{
			Msg: err.Error(),
			Err: err,
		}
	}

	if len(e.Errors.Phone) > 0 {
		return ErrCustomerPhone{
			Msg: err.Error(),
			Err: err,
		}
	}

	return nil
}

// ErrCustomerEmail is thrown when there has been an error with a customers email
type ErrCustomerEmail struct {
	Msg string
	Err error
}

// Error returns the error
func (e ErrCustomerEmail) Error() string {
	return e.Msg
}

// Unwrap unwraps the error
func (e ErrCustomerEmail) Unwrap() error {
	return e.Err
}

// ErrCustomerPhone is thrown when there has been an error with a customers phone number
type ErrCustomerPhone struct {
	Msg string
	Err error
}

// Error returns the error
func (e ErrCustomerPhone) Error() string {
	return e.Msg
}

// Unwrap unwraps the error
func (e ErrCustomerPhone) Unwrap() error {
	return e.Err
}
