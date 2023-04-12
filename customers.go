package httpshopify

import (
	"encoding/json"
	"fmt"
	httpCode "net/http"
	"strings"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"

	"github.com/MOHC-LTD/shopify/v2"
)

// CustomerDTO represents a Shopify customer in HTTP requests and responses
type CustomerDTO struct {
	Addresses  CustomerAddressDTOs `json:"addresses,omitempty"`
	ID         int64               `json:"id,omitempty"`
	Email      string              `json:"email,omitempty"`
	Phone      string              `json:"phone,omitempty"`
	FirstName  string              `json:"first_name,omitempty"`
	LastName   string              `json:"last_name,omitempty"`
	Tags       string              `json:"tags,omitempty"`
	MetaFields []metafieldDTO      `json:"metafields,omitempty"`
	CreatedAt  *time.Time          `json:"created_at,omitempty"`
	UpdatedAt  *time.Time          `json:"updated_at,omitempty"`
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
		Addresses: dto.Addresses.ToShopify(),
		ID:        dto.ID,
		Email:     dto.Email,
		Phone:     dto.Phone,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Tags:      shopify.Tags(dto.Tags),
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

	metafields := make(metafieldsDTO, len(customer.Metafields))
	for _, metafield := range customer.Metafields {
		metafields = append(metafields, metafieldDTO{
			Key:       metafield.Key,
			Namespace: metafield.Namespace,
			Value:     metafield.Value,
			Type:      metafield.Type,
		})
	}

	customerDTO := CustomerDTO{
		ID:         customer.ID,
		Email:      customer.Email,
		Phone:      customer.Phone,
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Tags:       string(customer.Tags),
		MetaFields: metafields,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}

	return customerDTO
}

// CustomerDTOs is a collection of Customer DTOs
type CustomerDTOs []CustomerDTO

// ToShopify converts the DTOs to the Shopify equivalent
func (dtos CustomerDTOs) ToShopify() shopify.Customers {
	customers := make(shopify.Customers, 0, len(dtos))

	for _, dto := range dtos {
		customers = append(customers, dto.ToShopify())
	}

	return customers
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

func (c customerRepository) Get(id int64) (shopify.Customer, error) {
	url := c.createURL(fmt.Sprintf("customers/%v.json", id))

	body, _, err := c.client.Get(url, nil)
	if err != nil {
		switch err.(type) {
		// TODO This ErrHTTP feels like bloat now. Can probably simplify the http work
		case http.ErrHTTP:
			shopifyError := err.(http.ErrHTTP)
			switch shopifyError.Code {
			case httpCode.StatusNotFound:
				return shopify.Customer{}, shopify.NewErrCustomerNotFound(id)
			}
		}
		return shopify.Customer{}, err
	}

	var responseDTO struct {
		Customer CustomerDTO `json:"customer"`
	}

	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return shopify.Customer{}, err
	}

	if responseDTO.Customer.ID == 0 {
		return shopify.Customer{}, shopify.NewErrCustomerNotFound(id)
	}

	return responseDTO.Customer.ToShopify(), nil
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
		// TODO This ErrHTTP feels like bloat now. Can probably simplify the http work
		case http.ErrHTTP:
			shopifyError := err.(http.ErrHTTP)
			switch shopifyError.Code {
			case httpCode.StatusUnprocessableEntity:
				var unprocessableEntityDTO errCustomerUnprocessableEntityDTO
				e := json.Unmarshal([]byte(shopifyError.Body), &unprocessableEntityDTO)
				if e != nil {
					return shopify.Customer{}, err
				}

				return shopify.Customer{}, unprocessableEntityDTO.toError()
			}

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

func (c customerRepository) GetByQuery(fields []string, query shopify.CustomerSearchQuery) (shopify.Customers, error) {
	url := c.createURL(fmt.Sprintf("customers/search.json?fields=%v&query=%s", strings.Join(fields, ","), query.String()))

	body, _, err := c.client.Get(url, nil)
	if err != nil {
		switch err.(http.ErrHTTP).Code {
		case httpCode.StatusNotFound:
			return shopify.Customers{}, err
		default:
			return shopify.Customers{}, err
		}
	}

	var responseDTO struct {
		Customers CustomerDTOs `json:"customers"`
	}

	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		return shopify.Customers{}, err
	}

	return responseDTO.Customers.ToShopify(), nil
}

type errCustomerUnprocessableEntityDTO struct {
	Errors struct {
		Email []string `json:"email"`
		Phone []string `json:"phone"`
	} `json:"errors"`
}

func (dto errCustomerUnprocessableEntityDTO) toError() error {
	return ErrCustomerUnprocessableEntity{
		Email: dto.Errors.Email,
		Phone: dto.Errors.Phone,
	}
}

// ErrCustomerUnprocessableEntity is used to store unprocessable entity error responses for a customer
type ErrCustomerUnprocessableEntity struct {
	Email []string
	Phone []string
}

func (e ErrCustomerUnprocessableEntity) Error() string {
	var errorMessages []string

	// Get the phone errors
	for _, errorMessage := range e.Phone {
		errorMessages = append(errorMessages, fmt.Sprintf("Phone number: %s", errorMessage))
	}

	// Get the email errors
	for _, errorMessage := range e.Email {
		errorMessages = append(errorMessages, fmt.Sprintf("Email address: %s", errorMessage))
	}

	return strings.Join(errorMessages, ", ")
}
