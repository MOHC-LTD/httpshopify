package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"

	"github.com/MOHC-LTD/shopify/v2"
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

func (repository fulfillmentRepository) Create(fulfillment shopify.Fulfillment) (shopify.Fulfillment, error) {
	createDTO := FulfillmentDTO{
		NotifyCustomer:              fulfillment.NotifyCustomer,
		TrackingInfo:                BuildTrackingInfoDTO(fulfillment),
		LineItemsByFulfillmentOrder: BuildLineItemsByFulfillmentOrderDTOs(fulfillment.LineItemsByFulfillmentOrder),
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

	url := repository.createURL("fulfillments.json")

	respBody, _, err := repository.client.Post(url, body, nil)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	response := struct {
		Fulfillment FulfillmentDTO `json:"fulfillment"`
	}{}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	return response.Fulfillment.ToShopify(), nil
}

func (repository fulfillmentRepository) UpdateTracking(update shopify.Fulfillment) (shopify.Fulfillment, error) {
	updateTrackingDTO := FulfillmentDTO{
		TrackingInfo:   BuildTrackingInfoDTO(update),
		NotifyCustomer: update.NotifyCustomer,
	}

	request := struct {
		Fulfillment FulfillmentDTO `json:"fulfillment"`
	}{
		Fulfillment: updateTrackingDTO,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	url := repository.createURL(fmt.Sprintf("fulfillments/%d/update_tracking.json", update.ID))

	respBody, _, err := repository.client.Post(url, body, nil)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	response := struct {
		Fulfillment FulfillmentDTO `json:"fulfillment"`
	}{}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.Fulfillment{}, err
	}

	return response.Fulfillment.ToShopify(), nil
}

func (repository fulfillmentRepository) Cancel(id int64) error {
	url := repository.createURL(fmt.Sprintf("fulfillments/%d/cancel.json", id))

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
	ID                          int64                            `json:"id,omitempty"`
	OrderID                     int64                            `json:"order_id,omitempty"`
	TrackingCompany             string                           `json:"tracking_company,omitempty"`
	TrackingNumber              string                           `json:"tracking_number,omitempty"`
	TrackingNumbers             []string                         `json:"tracking_numbers,omitempty"`
	TrackingURL                 string                           `json:"tracking_url,omitempty"`
	TrackingURLs                []string                         `json:"tracking_urls,omitempty"`
	Status                      string                           `json:"status,omitempty"`
	CreatedAt                   *time.Time                       `json:"created_at,omitempty"`
	UpdatedAt                   *time.Time                       `json:"updated_at,omitempty"`
	NotifyCustomer              bool                             `json:"notify_customer,omitempty"`
	ShipmentStatus              string                           `json:"shipment_status,omitempty"`
	LocationID                  int64                            `json:"location_id,omitempty"`
	LineItems                   LineItemDTOs                     `json:"line_items,omitempty"`
	TrackingInfo                *TrackingInfoDTO                 `json:"tracking_info,omitempty"`
	LineItemsByFulfillmentOrder []LineItemsByFulfillmentOrderDTO `json:"line_items_by_fulfillment_order,omitempty"`
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
		TrackingCompany: dto.TrackingCompany,
		TrackingNumber:  dto.TrackingNumber,
		TrackingNumbers: dto.TrackingNumbers,
		TrackingURL:     dto.TrackingURL,
		TrackingURLs:    dto.TrackingURLs,
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
		TrackingCompany: fulfillment.TrackingCompany,
		TrackingNumber:  fulfillment.TrackingNumber,
		TrackingNumbers: fulfillment.TrackingNumbers,
		TrackingURL:     fulfillment.TrackingURL,
		TrackingURLs:    fulfillment.TrackingURLs,
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

// TrackingInfoDTO represents Shopify fulfillment tracking info in HTTP requests and responses
type TrackingInfoDTO struct {
	Company string `json:"company,omitempty"`
	Number  string `json:"number,omitempty"`
	URL     string `json:"url,omitempty"`
}

// BuildTrackingInfoDTO converts fulfillment tracking info into a DTO equivalent
func BuildTrackingInfoDTO(fulfillment shopify.Fulfillment) *TrackingInfoDTO {
	if fulfillment.TrackingNumber == "" && fulfillment.TrackingURL == "" && fulfillment.TrackingCompany == "" {
		return nil
	}

	return &TrackingInfoDTO{
		Company: fulfillment.TrackingCompany,
		Number:  fulfillment.TrackingNumber,
		URL:     fulfillment.TrackingURL,
	}
}

// LineItemsByFulfillmentOrderDTO represents Shopify fulfillment order line items input in HTTP requests and responses
type LineItemsByFulfillmentOrderDTO struct {
	FulfillmentOrderID        int64                         `json:"fulfillment_order_id,omitempty"`
	FulfillmentOrderLineItems []FulfillmentOrderLineItemDTO `json:"fulfillment_order_line_items,omitempty"`
}

// BuildLineItemsByFulfillmentOrderDTO converts a fulfillment order into a DTO equivalent
func BuildLineItemsByFulfillmentOrderDTO(fulfillmentOrder shopify.FulfillmentOrder) LineItemsByFulfillmentOrderDTO {
	return LineItemsByFulfillmentOrderDTO{
		FulfillmentOrderID:        fulfillmentOrder.ID,
		FulfillmentOrderLineItems: BuildFulfillmentOrderLineItemDTOs(fulfillmentOrder.LineItems),
	}
}

// BuildLineItemsByFulfillmentOrderDTOs converts many fulfillment orders into a DTO equivalent
func BuildLineItemsByFulfillmentOrderDTOs(fulfillmentOrders []shopify.FulfillmentOrder) []LineItemsByFulfillmentOrderDTO {
	dtos := make([]LineItemsByFulfillmentOrderDTO, 0, len(fulfillmentOrders))
	for _, fulfillmentOrder := range fulfillmentOrders {
		dtos = append(dtos, BuildLineItemsByFulfillmentOrderDTO(fulfillmentOrder))
	}
	return dtos
}
