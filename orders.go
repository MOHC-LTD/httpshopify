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

// OrderDTO represents a Shopify order in HTTP requests and responses
type OrderDTO struct {
	BillingAddress           AddressDTO       `json:"billing_address"`
	ClosedAt                 time.Time        `json:"closed_at"`
	CreatedAt                time.Time        `json:"created_at"`
	Currency                 string           `json:"currency"`
	CurrentTotalDiscounts    string           `json:"current_total_discounts"`
	CurrentTotalDiscountsSet PriceSetDTO      `json:"current_total_discounts_set"`
	CurrentTotalPrice        string           `json:"current_total_price"`
	CurrentTotalPriceSet     PriceSetDTO      `json:"current_total_price_set"`
	CurrentSubtotalPrice     string           `json:"current_subtotal_price"`
	CurrentSubtotalPriceSet  PriceSetDTO      `json:"current_subtotal_price_set"`
	CurrentTotalTax          string           `json:"current_total_tax"`
	CurrentTotalTaxSet       PriceSetDTO      `json:"current_total_tax_set"`
	Customer                 CustomerDTO      `json:"customer"`
	Email                    string           `json:"email"`
	FinancialStatus          string           `json:"financial_status"`
	Fulfillments             FulfillmentDTOs  `json:"fulfillments"`
	FulfillmentStatus        string           `json:"fulfillment_status"`
	ID                       int64            `json:"id"`
	LineItems                LineItemDTOs     `json:"line_items"`
	Name                     string           `json:"name"`
	PresentmentCurrency      string           `json:"presentment_currency"`
	ShippingAddress          AddressDTO       `json:"shipping_address"`
	ShippingLines            ShippingLineDTOs `json:"shipping_lines"`
	SubtotalPrice            string           `json:"subtotal_price"`
	SubtotalPriceSet         PriceSetDTO      `json:"subtotal_price_set"`
	TotalDiscounts           string           `json:"total_discounts"`
	TotalDiscountsSet        PriceSetDTO      `json:"total_discounts_set"`
	TotalLineItemsPrice      string           `json:"total_line_items_price"`
	TotalLineItemsPriceSet   PriceSetDTO      `json:"total_line_items_price_set"`
	TotalPrice               string           `json:"total_price"`
	TotalPriceSet            PriceSetDTO      `json:"total_price_set"`
	TotalTax                 string           `json:"total_tax"`
	TotalTaxSet              PriceSetDTO      `json:"total_tax_set"`
	UpdatedAt                time.Time        `json:"updated_at"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto OrderDTO) ToShopify() shopify.Order {
	return shopify.Order{
		ID:                dto.ID,
		Name:              dto.Name,
		UpdatedAt:         dto.UpdatedAt,
		CreatedAt:         dto.CreatedAt,
		ClosedAt:          dto.ClosedAt,
		FulfillmentStatus: dto.FulfillmentStatus,
		FinancialStatus:   dto.FinancialStatus,
		ShippingLines:     dto.ShippingLines.ToShopify(),
		Customer:          dto.Customer.ToShopify(),
		Fulfillments:      dto.Fulfillments.ToShopify(),
		LineItems:         dto.LineItems.ToShopify(),
		BillingAddress:    dto.BillingAddress.ToShopify(),
		ShippingAddress:   dto.ShippingAddress.ToShopify(),
	}
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
