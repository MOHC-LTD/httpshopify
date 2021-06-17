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
		LineItems:       buildLineItemDTOs(fulfillment.LineItems),
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
	ID              int64        `json:"id"`
	OrderID         int64        `json:"order_id"`
	TrackingNumbers []string     `json:"tracking_numbers"`
	Status          string       `json:"status"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	NotifyCustomer  bool         `json:"notify_customer"`
	ShipmentStatus  string       `json:"shipment_status"`
	LocationID      int64        `json:"location_id"`
	LineItems       LineItemDTOs `json:"lineItems"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto FulfillmentDTO) ToShopify() shopify.Fulfillment {
	return shopify.Fulfillment{
		ID:              dto.ID,
		OrderID:         dto.OrderID,
		TrackingNumbers: dto.TrackingNumbers,
		Status:          dto.Status,
		CreatedAt:       dto.CreatedAt,
		UpdatedAt:       dto.UpdatedAt,
		NotifyCustomer:  dto.NotifyCustomer,
		ShipmentStatus:  dto.ShipmentStatus,
		LocationID:      dto.LocationID,
		LineItems:       dto.LineItems.ToShopify(),
	}
}

// ToShopify converts the DTO to the Shopify equivalent
func (dtos FulfillmentDTOs) ToShopify() []shopify.Fulfillment {
	fulfillments := make([]shopify.Fulfillment, 0, len(dtos))

	for _, dto := range dtos {
		fulfillments = append(fulfillments, dto.ToShopify())
	}

	return fulfillments
}
