package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"

	"github.com/MOHC-LTD/shopify/v2"
)

type fulfillmentOrderRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newFulfillmentOrderRepository(client http.Client, createURL func(endpoint string) string) fulfillmentOrderRepository {
	return fulfillmentOrderRepository{
		client,
		createURL,
	}
}

func (repository fulfillmentOrderRepository) Get(id int64) (shopify.FulfillmentOrder, error) {
	url := repository.createURL(fmt.Sprintf("fulfillment_orders/%d.json", id))

	respBody, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.FulfillmentOrder{}, err
	}

	response := struct {
		FulfillmentOrder FulfillmentOrderDTO `json:"fulfillment_order"`
	}{}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.FulfillmentOrder{}, err
	}

	return response.FulfillmentOrder.ToShopify(), nil
}

func (repository fulfillmentOrderRepository) List(orderID int64) (shopify.FulfillmentOrders, error) {
	url := repository.createURL(fmt.Sprintf("orders/%d/fulfillment_orders.json", orderID))

	respBody, _, err := repository.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	response := struct {
		FulfillmentOrders []FulfillmentOrderDTO `json:"fulfillment_orders"`
	}{}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}

	fulfillmentOrders := make(shopify.FulfillmentOrders, 0, len(response.FulfillmentOrders))
	for _, dto := range response.FulfillmentOrders {
		fulfillmentOrders = append(fulfillmentOrders, dto.ToShopify())
	}

	return fulfillmentOrders, nil
}

// FulfillmentOrderDTO represents a Shopify fulfillment order in HTTP requests and responses
type FulfillmentOrderDTO struct {
	ID                 int64                        `json:"id,omitempty"`
	OrderID            int64                        `json:"order_id,omitempty"`
	AssignedLocationID int64                        `json:"assigned_location_id,omitempty"`
	Status             string                       `json:"status,omitempty"`
	LineItems          FulfillmentOrderLineItemDTOs `json:"line_items,omitempty"`
	CreatedAt          *time.Time                   `json:"created_at,omitempty"`
	UpdatedAt          *time.Time                   `json:"updated_at,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto FulfillmentOrderDTO) ToShopify() shopify.FulfillmentOrder {
	createdAt := time.Time{}
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	updatedAt := time.Time{}
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.FulfillmentOrder{
		ID:                 dto.ID,
		OrderID:            dto.OrderID,
		AssignedLocationID: dto.AssignedLocationID,
		LineItems:          dto.LineItems.ToShopify(),
		Status:             dto.Status,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
	}
}

// FulfillmentOrderLineItemDTO represents a Shopify fulfillment order in HTTP requests and responses
type FulfillmentOrderLineItemDTO struct {
	ID                  int64 `json:"id,omitempty"`
	LineItemID          int64 `json:"line_item_id,omitempty"`
	VariantID           int64 `json:"variant_id,omitempty"`
	Quantity            int   `json:"quantity,omitempty"`
	FulfillableQuantity int   `json:"fulfillable_quantity,omitempty"`
}

func (dto FulfillmentOrderLineItemDTO) ToShopify() shopify.FulfillmentOrderLineItem {
	return shopify.FulfillmentOrderLineItem{
		ID:                  dto.ID,
		LineItemID:          dto.LineItemID,
		VariantID:           dto.VariantID,
		Quantity:            dto.Quantity,
		FulfillableQuantity: dto.FulfillableQuantity,
	}
}

// BuildFulfillmentOrderLineItemDTO converts a fulfillment order line item into a DTO equivalent
func BuildFulfillmentOrderLineItemDTO(lineItem shopify.FulfillmentOrderLineItem) FulfillmentOrderLineItemDTO {
	return FulfillmentOrderLineItemDTO{
		ID:                  lineItem.ID,
		Quantity:            lineItem.Quantity,
		LineItemID:          lineItem.LineItemID,
		VariantID:           lineItem.VariantID,
		FulfillableQuantity: lineItem.FulfillableQuantity,
	}
}

// FulfillmentOrderLineItemDTOs is a collection of fulfillment order line item DTOs
type FulfillmentOrderLineItemDTOs []FulfillmentOrderLineItemDTO

// ToShopify converts the DTOs to the Shopify equivalent
func (dtos FulfillmentOrderLineItemDTOs) ToShopify() shopify.FulfillmentOrderLineItems {
	lineItems := make(shopify.FulfillmentOrderLineItems, 0, len(dtos))
	for _, dto := range dtos {
		lineItems = append(lineItems, dto.ToShopify())
	}
	return lineItems
}

// BuildFulfillmentOrderLineItemDTOs converts many fulfillment order line items into a DTO equivalent
func BuildFulfillmentOrderLineItemDTOs(lineItems []shopify.FulfillmentOrderLineItem) []FulfillmentOrderLineItemDTO {
	dtos := make([]FulfillmentOrderLineItemDTO, 0, len(lineItems))
	for _, lineItem := range lineItems {
		dtos = append(dtos, BuildFulfillmentOrderLineItemDTO(lineItem))
	}
	return dtos
}
