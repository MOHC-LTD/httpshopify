package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/slices"

	"github.com/MOHC-LTD/httpshopify/internal/http"

	"github.com/MOHC-LTD/shopify"
)

type orderRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newOrderRepository(client http.Client, createURL func(endpoint string) string) orderRepository {
	return orderRepository{
		client,
		createURL,
	}
}

func (repository orderRepository) List(query shopify.OrderQuery) (shopify.Orders, error) {
	orders := make(shopify.Orders, 0)

	url := repository.createURL(fmt.Sprintf("orders.json%v", parseOrderQuery(query)))

	for {
		body, headers, err := repository.client.Get(url, nil)
		if err != nil {
			return nil, err
		}

		var resultDTO struct {
			Orders []OrderDTO `json:"orders"`
		}
		json.Unmarshal(body, &resultDTO)

		for _, dto := range resultDTO.Orders {
			orders = append(orders, dto.ToShopify())
		}

		links := ParseLinkHeader(headers.Get("Link"))

		if !links.HasNext() {
			break
		}

		url = links.Next
	}

	return orders, nil
}

func (repository orderRepository) Get(id int64) (shopify.Order, error) {
	url := repository.createURL(fmt.Sprintf("orders/%v.json", id))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.Order{}, err
	}

	var response struct {
		Order OrderDTO `json:"order"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return shopify.Order{}, err
	}

	if response.Order.ID == 0 {
		return shopify.Order{}, shopify.NewErrOrderNotFound(id)
	}

	return response.Order.ToShopify(), nil
}

func (repository orderRepository) Close(id int64) error {
	url := repository.createURL(fmt.Sprintf("orders/%v/close.json", id))

	_, _, err := repository.client.Post(url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repository orderRepository) Create(order shopify.Order) (shopify.Order, error) {
	url := repository.createURL("orders.json")

	bodyData := struct {
		Order OrderDTO `json:"order"`
	}{
		Order: BuildOrderDTO(order),
	}

	body, err := json.Marshal(&bodyData)
	if err != nil {
		return shopify.Order{}, err
	}

	responseBody, _, err := repository.client.Post(url, body, nil)
	if err != nil {
		return shopify.Order{}, err
	}

	var response struct {
		Order OrderDTO `json:"order"`
	}
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return shopify.Order{}, err
	}

	return response.Order.ToShopify(), nil
}

// OrderDTO represents a Shopify order in HTTP requests and responses
type OrderDTO struct {
	BillingAddress           AddressDTO       `json:"billing_address,omitempty"`
	ClosedAt                 *time.Time       `json:"closed_at,omitempty"`
	CreatedAt                *time.Time       `json:"created_at,omitempty"`
	Currency                 string           `json:"currency,omitempty"`
	CurrentTotalDiscounts    string           `json:"current_total_discounts,omitempty"`
	CurrentTotalDiscountsSet PriceSetDTO      `json:"current_total_discounts_set,omitempty"`
	CurrentTotalPrice        string           `json:"current_total_price,omitempty"`
	CurrentTotalPriceSet     PriceSetDTO      `json:"current_total_price_set,omitempty"`
	CurrentSubtotalPrice     string           `json:"current_subtotal_price,omitempty"`
	CurrentSubtotalPriceSet  PriceSetDTO      `json:"current_subtotal_price_set,omitempty"`
	CurrentTotalTax          string           `json:"current_total_tax,omitempty"`
	CurrentTotalTaxSet       PriceSetDTO      `json:"current_total_tax_set,omitempty"`
	Customer                 CustomerDTO      `json:"customer,omitempty"`
	Email                    string           `json:"email,omitempty"`
	FinancialStatus          string           `json:"financial_status,omitempty"`
	Fulfillments             FulfillmentDTOs  `json:"fulfillments,omitempty"`
	FulfillmentStatus        string           `json:"fulfillment_status,omitempty"`
	ID                       int64            `json:"id,omitempty"`
	LineItems                LineItemDTOs     `json:"line_items,omitempty"`
	Name                     string           `json:"name,omitempty"`
	OrderNumber              int              `json:"order_number,omitempty"`
	PresentmentCurrency      string           `json:"presentment_currency,omitempty"`
	ProcessedAt              *time.Time       `json:"processed_at,omitempty"`
	ShippingAddress          AddressDTO       `json:"shipping_address,omitempty"`
	ShippingLines            ShippingLineDTOs `json:"shipping_lines,omitempty"`
	SubtotalPrice            string           `json:"subtotal_price,omitempty"`
	SubtotalPriceSet         PriceSetDTO      `json:"subtotal_price_set,omitempty"`
	TotalDiscounts           string           `json:"total_discounts,omitempty"`
	TotalDiscountsSet        PriceSetDTO      `json:"total_discounts_set,omitempty"`
	TotalLineItemsPrice      string           `json:"total_line_items_price,omitempty"`
	TotalLineItemsPriceSet   PriceSetDTO      `json:"total_line_items_price_set,omitempty"`
	TotalPrice               string           `json:"total_price,omitempty"`
	TotalPriceSet            PriceSetDTO      `json:"total_price_set,omitempty"`
	TotalTax                 string           `json:"total_tax,omitempty"`
	TotalTaxSet              PriceSetDTO      `json:"total_tax_set,omitempty"`
	UpdatedAt                *time.Time       `json:"updated_at,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto OrderDTO) ToShopify() shopify.Order {

	if dto.CreatedAt.IsZero() {
		dto.CreatedAt = nil
	}

	if dto.ClosedAt.IsZero() {
		dto.ClosedAt = nil
	}

	if dto.ProcessedAt.IsZero() {
		dto.ProcessedAt = nil
	}

	if dto.UpdatedAt.IsZero() {
		dto.UpdatedAt = nil
	}

	return shopify.Order{
		BillingAddress:           dto.BillingAddress.ToShopify(),
		ClosedAt:                 *dto.ClosedAt,
		CreatedAt:                *dto.CreatedAt,
		Currency:                 dto.Currency,
		CurrentTotalDiscounts:    dto.CurrentTotalDiscounts,
		CurrentTotalDiscountsSet: dto.CurrentTotalDiscountsSet.ToShopify(),
		CurrentTotalPrice:        dto.CurrentTotalPrice,
		CurrentTotalPriceSet:     dto.CurrentTotalPriceSet.ToShopify(),
		CurrentSubtotalPrice:     dto.CurrentSubtotalPrice,
		CurrentSubtotalPriceSet:  dto.CurrentSubtotalPriceSet.ToShopify(),
		CurrentTotalTax:          dto.CurrentTotalTax,
		CurrentTotalTaxSet:       dto.CurrentTotalTaxSet.ToShopify(),
		Customer:                 dto.Customer.ToShopify(),
		Email:                    dto.Email,
		FinancialStatus:          dto.FinancialStatus,
		Fulfillments:             dto.Fulfillments.ToShopify(),
		FulfillmentStatus:        dto.FulfillmentStatus,
		ID:                       dto.ID,
		LineItems:                dto.LineItems.ToShopify(),
		Name:                     dto.Name,
		OrderNumber:              dto.OrderNumber,
		PresentmentCurrency:      dto.PresentmentCurrency,
		ProcessedAt:              *dto.ProcessedAt,
		ShippingAddress:          dto.ShippingAddress.ToShopify(),
		ShippingLines:            dto.ShippingLines.ToShopify(),
		SubtotalPrice:            dto.SubtotalPrice,
		SubtotalPriceSet:         dto.SubtotalPriceSet.ToShopify(),
		TotalDiscounts:           dto.TotalDiscounts,
		TotalDiscountsSet:        dto.TotalDiscountsSet.ToShopify(),
		TotalLineItemsPrice:      dto.TotalLineItemsPrice,
		TotalLineItemsPriceSet:   dto.TotalLineItemsPriceSet.ToShopify(),
		TotalPrice:               dto.TotalPrice,
		TotalPriceSet:            dto.TotalPriceSet.ToShopify(),
		TotalTax:                 dto.TotalTax,
		TotalTaxSet:              dto.TotalTaxSet.ToShopify(),
		UpdatedAt:                *dto.UpdatedAt,
	}
}

// BuildOrderDTO converts a Shopify order to the DTO equivalent
func BuildOrderDTO(order shopify.Order) OrderDTO {

	orderDTO := OrderDTO{
		BillingAddress:           BuildAddressDTO(order.BillingAddress),
		Currency:                 order.Currency,
		CurrentTotalDiscounts:    order.CurrentTotalDiscounts,
		CurrentTotalDiscountsSet: BuildPriceSetDTO(order.CurrentTotalDiscountsSet),
		CurrentTotalPrice:        order.CurrentTotalPrice,
		CurrentTotalPriceSet:     BuildPriceSetDTO(order.CurrentTotalPriceSet),
		CurrentSubtotalPrice:     order.CurrentSubtotalPrice,
		CurrentSubtotalPriceSet:  BuildPriceSetDTO(order.CurrentSubtotalPriceSet),
		CurrentTotalTax:          order.CurrentTotalTax,
		CurrentTotalTaxSet:       BuildPriceSetDTO(order.CurrentTotalTaxSet),
		Customer:                 BuildCustomerDTO(order.Customer),
		Email:                    order.Email,
		FinancialStatus:          order.FinancialStatus,
		Fulfillments:             BuildFulfillmentDTOs(order.Fulfillments),
		FulfillmentStatus:        order.FulfillmentStatus,
		ID:                       order.ID,
		LineItems:                BuildLineItemDTOs(order.LineItems),
		Name:                     order.Name,
		OrderNumber:              order.OrderNumber,
		PresentmentCurrency:      order.PresentmentCurrency,
		ShippingAddress:          BuildAddressDTO(order.ShippingAddress),
		ShippingLines:            BuildShippingLineDTOs(order.ShippingLines),
		SubtotalPrice:            order.SubtotalPrice,
		SubtotalPriceSet:         BuildPriceSetDTO(order.SubtotalPriceSet),
		TotalDiscounts:           order.TotalDiscounts,
		TotalDiscountsSet:        BuildPriceSetDTO(order.TotalDiscountsSet),
		TotalLineItemsPrice:      order.TotalLineItemsPrice,
		TotalLineItemsPriceSet:   BuildPriceSetDTO(order.TotalLineItemsPriceSet),
		TotalPrice:               order.TotalPrice,
		TotalPriceSet:            BuildPriceSetDTO(order.TotalPriceSet),
		TotalTax:                 order.TotalTax,
		TotalTaxSet:              BuildPriceSetDTO(order.TotalTaxSet),
	}

	if order.CreatedAt.IsZero() {
		orderDTO.CreatedAt = nil
	}

	if order.ClosedAt.IsZero() {
		orderDTO.ClosedAt = nil
	}

	if order.ProcessedAt.IsZero() {
		orderDTO.ProcessedAt = nil
	}

	if order.UpdatedAt.IsZero() {
		orderDTO.UpdatedAt = nil
	}

	return orderDTO
}

func parseOrderQuery(query shopify.OrderQuery) string {
	queryStrings := make([]string, 0)

	if query.Status != "" {
		queryStrings = append(queryStrings, fmt.Sprintf("status=%v", query.Status))
	}

	if query.FinancialStatus != "" {
		queryStrings = append(queryStrings, fmt.Sprintf("financial_status=%v", query.FinancialStatus))
	}

	if query.FulfillmentStatus != "" {
		queryStrings = append(queryStrings, fmt.Sprintf("fulfillment_status=%v", query.FulfillmentStatus))
	}

	if query.SinceID != 0 {
		queryStrings = append(queryStrings, fmt.Sprintf("since_id=%v", query.SinceID))
	}

	if query.IDs != nil {
		queryStrings = append(queryStrings, fmt.Sprintf("ids=%v", slices.JoinInt64(query.IDs, ",")))
	}

	if len(queryStrings) == 0 {
		return ""
	}

	queryString := "?"

	for i, str := range queryStrings {
		if i != 0 {
			queryString = queryString + "&"
		}

		queryString = queryString + str
	}

	return queryString
}
