package httpshopify

import (
	"github.com/MOHC-LTD/httpshopify/v2/internal/http"

	"github.com/MOHC-LTD/shopify/v2"
)

type inventoryItemRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newInventoryItemRepository(client http.Client, createURL func(endpoint string) string) inventoryItemRepository {
	return inventoryItemRepository{
		client,
		createURL,
	}
}

func (repository inventoryItemRepository) Get(id int64) (shopify.InventoryItem, error) {
	panic("Get has not been implement yet")
}
