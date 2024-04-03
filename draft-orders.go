package httpshopify

import (
	"encoding/json"
	"fmt"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"

	"github.com/MOHC-LTD/shopify/v2"
)

type draftOrderRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newDraftOrderRepository(client http.Client, createURL func(endpoint string) string) draftOrderRepository {
	return draftOrderRepository{
		client,
		createURL,
	}
}

func (repository draftOrderRepository) Get(id int64) (shopify.DraftOrder, error) {
	url := repository.createURL(fmt.Sprintf("draft_orders/%v.json", id))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.DraftOrder{}, err
	}

	var response struct {
		DraftOrder DraftOrderDTO `json:"draft_order"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return shopify.DraftOrder{}, err
	}

	if response.DraftOrder.ID == 0 {
		return shopify.DraftOrder{}, shopify.NewErrOrderNotFound(id)
	}

	return response.DraftOrder.ToShopify(), nil
}

// DraftOrderDTO represents a Shopify draft order in HTTP requests and responses
type DraftOrderDTO struct {
	ID              int64      `json:"id,omitempty"`
	ShippingAddress AddressDTO `json:"shipping_address,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto DraftOrderDTO) ToShopify() shopify.DraftOrder {
	return shopify.DraftOrder{
		ShippingAddress: dto.ShippingAddress.ToShopify(),
	}
}

// DraftOrderDTOs is a collection of Draft Order DTOs
type DraftOrderDTOs []DraftOrderDTO

// ToShopify converts the DTOs to the Shopify equivalent
func (dtos DraftOrderDTOs) ToShopify() shopify.DraftOrders {
	orders := make(shopify.DraftOrders, 0, len(dtos))

	for _, dto := range dtos {
		orders = append(orders, dto.ToShopify())
	}

	return orders
}
