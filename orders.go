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

	url := repository.createURL(fmt.Sprintf("orders.json%v", parseQuery(query)))

	for {
		body, headers, err := repository.client.Get(url, nil)
		if err != nil {
			return nil, err
		}

		var resultDTO struct {
			Orders []orderDTO `json:"orders"`
		}
		json.Unmarshal(body, &resultDTO)

		for _, dto := range resultDTO.Orders {
			orders = append(orders, dto.toDomain())
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
		Order orderDTO `json:"order"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return shopify.Order{}, err
	}

	if response.Order.ID == 0 {
		return shopify.Order{}, shopify.NewErrOrderNotFound(id)
	}

	return response.Order.toDomain(), nil
}

func (repository orderRepository) Close(id int64) error {
	url := repository.createURL(fmt.Sprintf("orders/%v/close.json", id))

	_, _, err := repository.client.Post(url, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

type orderDTO struct {
	ID                int64            `json:"id"`
	Name              string           `json:"name"`
	UpdatedAt         time.Time        `json:"updated_at"`
	CreatedAt         time.Time        `json:"created_at"`
	ClosedAt          time.Time        `json:"closed_at"`
	FulfillmentStatus string           `json:"fulfillment_status"`
	FinancialStatus   string           `json:"financial_status"`
	ShippingLines     shippingLineDTOs `json:"shipping_lines"`
	Customer          customerDTO      `json:"customer"`
	Fulfillments      fulfillmentDTOs  `json:"fulfillments"`
	LineItems         lineItemDTOs     `json:"line_items"`
}

func (dto orderDTO) toDomain() shopify.Order {
	return shopify.Order{
		ID:                dto.ID,
		Name:              dto.Name,
		UpdatedAt:         dto.UpdatedAt,
		CreatedAt:         dto.CreatedAt,
		ClosedAt:          dto.ClosedAt,
		FulfillmentStatus: dto.FulfillmentStatus,
		FinancialStatus:   dto.FinancialStatus,
		ShippingLines:     dto.ShippingLines.toDomain(),
		Customer:          dto.Customer.toDomain(),
		Fulfillments:      dto.Fulfillments.toDomain(),
		LineItems:         dto.LineItems.toDomain(),
	}
}

func parseQuery(query shopify.OrderQuery) string {
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
