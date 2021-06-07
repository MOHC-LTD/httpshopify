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
	createDTO := fulfillmentDTO{
		LocationID:      fulfillment.LocationID,
		TrackingNumbers: fulfillment.TrackingNumbers,
		NotifyCustomer:  fulfillment.NotifyCustomer,
		LineItems:       buildLineItemDTOs(fulfillment.LineItems),
	}

	request := struct {
		Fulfillment fulfillmentDTO `json:"fulfillment"`
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
		Fulfillment fulfillmentDTO `json:"fulfillment"`
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	return response.Fulfillment.toDomain(), nil
}

func (repository fulfillmentRepository) Update(orderID int64, fulfillmentID int64, update shopify.Fulfillment) (shopify.Fulfillment, error) {
	updateDTO := fulfillmentDTO{
		ID:              fulfillmentID,
		TrackingNumbers: update.TrackingNumbers,
		NotifyCustomer:  update.NotifyCustomer,
	}

	request := struct {
		Fulfillment fulfillmentDTO `json:"fulfillment"`
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
		Fulfillment fulfillmentDTO `json:"fulfillment"`
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	return response.Fulfillment.toDomain(), nil
}

func (repository fulfillmentRepository) Cancel(orderID int64, fulfillmentID int64) error {
	url := repository.createURL(fmt.Sprintf("orders/%v/fulfillments/%v/cancel.json", orderID, fulfillmentID))

	_, _, err := repository.client.Post(url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

type fulfillmentDTOs []fulfillmentDTO

type fulfillmentDTO struct {
	ID              int64        `json:"id"`
	OrderID         int64        `json:"order_id"`
	TrackingNumbers []string     `json:"tracking_numbers"`
	Status          string       `json:"status"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	NotifyCustomer  bool         `json:"notify_customer"`
	ShipmentStatus  string       `json:"shipment_status"`
	LocationID      int64        `json:"location_id"`
	LineItems       lineItemDTOs `json:"lineItems"`
}

func (dto fulfillmentDTO) toDomain() shopify.Fulfillment {
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
		LineItems:       dto.LineItems.toDomain(),
	}
}

func (dtos fulfillmentDTOs) toDomain() []shopify.Fulfillment {
	fulfillments := make([]shopify.Fulfillment, 0, len(dtos))

	for _, dto := range dtos {
		fulfillments = append(fulfillments, dto.toDomain())
	}

	return fulfillments
}
