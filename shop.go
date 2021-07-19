package httpshopify

import (
	"fmt"

	"github.com/MOHC-LTD/httpshopify/internal/http"

	"github.com/MOHC-LTD/shopify"
)

// Shop is an http shopify shop
type Shop struct {
	orders            orderRepository
	fulfillments      fulfillmentRepository
	fulfillmentEvents fulfillmentEventRepository
	variants          variantRepository
	products          productRepository
	inventoryLevels   inventoryLevelRepository
	collection        collectionRepository
}

// NewShop builds a shopify shop based on the shopify admin REST API
/*
	This constructor automatically determines the URL of the store from the store name.
	If you would like to use a custom store URL use the `NewCustomShop` constructor instread.
	Example:
	shop := shopify.NewShop("my-shop-name", "shppy_21u92h2184ho912h29r01")
	shippedOrders, err := shop.Orders().List(shopify.OrderQuery{FulfillmentStatus:"shipped"})
	For the full shopify admin REST API documentation see https://shopify.dev/docs/admin-api/rest/reference
*/
func NewShop(shop string, accessToken string) Shop {
	return NewCustomShop(
		fmt.Sprintf("https://%v.myshopify.com/admin/api/2021-04", shop),
		accessToken,
	)
}

// NewCustomShop builds a shopify shop based on the shopify admin REST API
/*
	This constructor automatically uses the URL passed to it to communicate with the store.
	If you would like to use an auto store URL use the `NewShop` constructor instread.
	Example:
	shop := shopify.NewCustomShop("https://my-shop-domain.com/foo/bar", "shppy_21u92h2184ho912h29r01")
	shippedOrders, err := shop.Orders().List(shopify.OrderQuery{FulfillmentStatus:"shipped"})
	For the full shopify admin REST API documentation see https://shopify.dev/docs/admin-api/rest/reference
*/
func NewCustomShop(url string, accessToken string) Shop {
	client := http.NewClient(http.RequestHeaders{
		{
			Name:  "X-Shopify-Access-Token",
			Value: accessToken,
		}, {
			Name:  "Content-Type",
			Value: "application/json",
		},
	})

	createURL := func(endpoint string) string {
		return fmt.Sprintf("%v/%v", url, endpoint)
	}

	return Shop{
		orders:            newOrderRepository(client, createURL),
		fulfillments:      newFulfillmentRepository(client, createURL),
		fulfillmentEvents: newFulfillmentEventRepository(client, createURL),
		variants:          newVariantRepository(client, createURL),
		products:          newProductRepository(client, createURL),
		inventoryLevels:   newInventoryLevelRepository(client, createURL),
		collection:        newCollectionRepository(client, createURL),
	}
}

// Orders returns an HTTP implementation of a Shopify order repository
func (shop Shop) Orders() shopify.OrderRepository {
	return shop.orders
}

// Fulfillments returns an HTTP implementation of a Shopify fulfillment repository
func (shop Shop) Fulfillments() shopify.FulfillmentRepository {
	return shop.fulfillments
}

// FulfillmentEvents returns an HTTP implementation of a Shopify fulfillment event repository
func (shop Shop) FulfillmentEvents() shopify.FulfillmentEventRepository {
	return shop.fulfillmentEvents
}

// Variants returns an HTTP implementation of a Shopify variant repository
func (shop Shop) Variants() shopify.VariantRepository {
	return shop.variants
}

// Products returns an HTTP implementation of a Shopify product repository
func (shop Shop) Products() shopify.ProductRepository {
	return shop.products
}

// InventoryLevels returns an HTTP implementation of a Shopify inventory level repository
func (shop Shop) InventoryLevels() shopify.InventoryLevelRepository {
	return shop.inventoryLevels
}

// Collections returns an HTTP implementation of a Shopify collection repository
func (shop Shop) Collections() shopify.CollectionRepository {
	return shop.collection
}
