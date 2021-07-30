package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/http"

	"github.com/MOHC-LTD/shopify"
)

type fulfillmentRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newFulfillmentRepository(client http.Client, createURL func(endpoint string) string) fulfillmentRepository {
	return fulfillmentRepository{
		client,
		createURL,
	}
}

func (repository fulfillmentRepository) Create(orderID int64, fulfillment shopify.Fulfillment) (shopify.Fulfillment, error) {
	createDTO := FulfillmentDTO{
		LocationID:      fulfillment.LocationID,
		TrackingNumbers: fulfillment.TrackingNumbers,
		NotifyCustomer:  fulfillment.NotifyCustomer,
		LineItems:       BuildLineItemDTOs(fulfillment.LineItems),
	}

	request := struct {
		Fulfillment FulfillmentDTO `json:"fulfillment"`
	}{
		Fulfillment: createDTO,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	url := repository.createURL(fmt.Sprintf("orders/%v/fulfillments.json", orderID))

	respBody, _, err := repository.client.Post(url, body, nil)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	var response struct {
		Fulfillment FulfillmentDTO `json:"fulfillment"`
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	return response.Fulfillment.ToShopify(), nil
}

func (repository fulfillmentRepository) Update(orderID int64, fulfillmentID int64, update shopify.Fulfillment) (shopify.Fulfillment, error) {
	updateDTO := FulfillmentDTO{
		ID:              fulfillmentID,
		TrackingNumbers: update.TrackingNumbers,
		NotifyCustomer:  update.NotifyCustomer,
	}

	request := struct {
		Fulfillment FulfillmentDTO `json:"fulfillment"`
	}{
		Fulfillment: updateDTO,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	url := repository.createURL(fmt.Sprintf("orders/%v/fulfillments/%v.json", orderID, fulfillmentID))

	respBody, _, err := repository.client.Put(url, body, nil)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	var response struct {
		Fulfillment FulfillmentDTO `json:"fulfillment"`
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	return response.Fulfillment.ToShopify(), nil
}

func (repository fulfillmentRepository) Cancel(orderID int64, fulfillmentID int64) error {
	url := repository.createURL(fmt.Sprintf("orders/%v/fulfillments/%v/cancel.json", orderID, fulfillmentID))

	_, _, err := repository.client.Post(url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// FulfillmentDTOs is a collection of Fulfillment DTOs
type FulfillmentDTOs []FulfillmentDTO

// FulfillmentDTO represents and Shopify fulfillment in HTTP requests and responses
type FulfillmentDTO struct {
	ID              int64        `json:"id,omitempty"`
	OrderID         int64        `json:"order_id,omitempty"`
	TrackingNumbers []string     `json:"tracking_numbers,omitempty"`
	Status          string       `json:"status,omitempty"`
	CreatedAt       *time.Time   `json:"created_at,omitempty"`
	UpdatedAt       *time.Time   `json:"updated_at,omitempty"`
	NotifyCustomer  bool         `json:"notify_customer,omitempty"`
	ShipmentStatus  string       `json:"shipment_status,omitempty"`
	LocationID      int64        `json:"location_id,omitempty"`
	LineItems       LineItemDTOs `json:"lineItems,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto FulfillmentDTO) ToShopify() shopify.Fulfillment {
	var createdAt time.Time
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	var updatedAt time.Time
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.Fulfillment{
		ID:              dto.ID,
		OrderID:         dto.OrderID,
		TrackingNumbers: dto.TrackingNumbers,
		Status:          dto.Status,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		NotifyCustomer:  dto.NotifyCustomer,
		ShipmentStatus:  dto.ShipmentStatus,
		LocationID:      dto.LocationID,
		LineItems:       dto.LineItems.ToShopify(),
	}
}

// BuildFulfilmentDTO converts the fulfillment into its DTO equivalent
func BuildFulfilmentDTO(fulfillment shopify.Fulfillment) FulfillmentDTO {
	var createdAt *time.Time
	if !fulfillment.CreatedAt.IsZero() {
		createdAt = &fulfillment.CreatedAt
	}

	var updatedAt *time.Time
	if !fulfillment.UpdatedAt.IsZero() {
		updatedAt = &fulfillment.UpdatedAt
	}

	fulfillmentDTO := FulfillmentDTO{
		ID:              fulfillment.ID,
		OrderID:         fulfillment.OrderID,
		TrackingNumbers: fulfillment.TrackingNumbers,
		Status:          fulfillment.Status,
		NotifyCustomer:  fulfillment.NotifyCustomer,
		ShipmentStatus:  fulfillment.ShipmentStatus,
		LocationID:      fulfillment.LocationID,
		LineItems:       BuildLineItemDTOs(fulfillment.LineItems),
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}

	return fulfillmentDTO
}

// ToShopify converts the DTO to the Shopify equivalent
func (dtos FulfillmentDTOs) ToShopify() []shopify.Fulfillment {
	fulfillments := make([]shopify.Fulfillment, 0, len(dtos))

	for _, dto := range dtos {
		fulfillments = append(fulfillments, dto.ToShopify())
	}

	return fulfillments
}

// BuildFulfillmentDTOs converts the fulfillments into their DTO equivalents
func BuildFulfillmentDTOs(fulfillments []shopify.Fulfillment) FulfillmentDTOs {
	dtos := make(FulfillmentDTOs, 0, len(fulfillments))

	for _, fulfillment := range fulfillments {
		dtos = append(dtos, BuildFulfilmentDTO(fulfillment))
	}

	return dtos
}
