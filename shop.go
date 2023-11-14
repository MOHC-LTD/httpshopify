package httpshopify

import (
	"fmt"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"

	"github.com/MOHC-LTD/shopify/v2"
)

// Shop is an http shopify shop
type Shop struct {
	orders            orderRepository
	fulfillments      fulfillmentRepository
	fulfillmentEvents fulfillmentEventRepository
	fulfillmentOrders fulfillmentOrderRepository
	variants          variantRepository
	products          productRepository
	inventoryLevels   inventoryLevelRepository
	collections       collectionRepository
	productImages     productImagesRepository
	metafields        metafieldRepository
	customers         customerRepository
	customerAddresses customerAddressRepository
	blogs             blogRepository
	articles          articleRepository
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
		fmt.Sprintf("https://%v.myshopify.com/admin/api/2022-07", shop),
		accessToken,
		IsDefault,
	)
}

// NewPlusShop builds a shopify plus shop based on the shopify admin REST API
/*
	This shop uses the Shopify plus rate limits allowing for higher throughput.
	This constructor automatically determines the URL of the store from the store name.
	If you would like to use a custom store URL use the `NewCustomShop` constructor instread.
	Example:
	shop := shopify.NewShop("my-shop-name", "shppy_21u92h2184ho912h29r01")
	shippedOrders, err := shop.Orders().List(shopify.OrderQuery{FulfillmentStatus:"shipped"})
	For the full shopify admin REST API documentation see https://shopify.dev/docs/admin-api/rest/reference
*/
func NewPlusShop(shop string, accessToken string) Shop {
	return NewCustomShop(
		fmt.Sprintf("https://%v.myshopify.com/admin/api/2022-07", shop),
		accessToken,
		IsPlus,
	)
}

// NewCustomShop builds a shopify shop based on the shopify admin REST API
/*
	This constructor automatically uses the URL passed to it to communicate with the store.
	If you would like to use an auto store URL use the `NewShop` constructor instead.
	Example:
	shop := shopify.NewCustomShop("https://my-shop-domain.com/foo/bar", "shppy_21u92h2184ho912h29r01", httpshopify.IsPlus)
	shippedOrders, err := shop.Orders().List(shopify.OrderQuery{FulfillmentStatus:"shipped"})
	For the full shopify admin REST API documentation see https://shopify.dev/docs/admin-api/rest/reference
*/
func NewCustomShop(url string, accessToken string, isPlus bool) Shop {
	var rateLimitOption http.Option
	if isPlus {
		rateLimitOption = RateLimitPlus()
	} else {
		rateLimitOption = RateLimitDefault()
	}

	client := http.NewClient(
		http.WithDefaultHeader("X-Shopify-Access-Token", accessToken),
		http.WithDefaultHeader("Content-Type", "application/json"),
		rateLimitOption,
	)

	createURL := func(endpoint string) string {
		return fmt.Sprintf("%v/%v", url, endpoint)
	}

	return Shop{
		orders:            newOrderRepository(client, createURL),
		fulfillments:      newFulfillmentRepository(client, createURL),
		fulfillmentEvents: newFulfillmentEventRepository(client, createURL),
		fulfillmentOrders: newFulfillmentOrderRepository(client, createURL),
		variants:          newVariantRepository(client, createURL),
		products:          newProductRepository(client, createURL),
		inventoryLevels:   newInventoryLevelRepository(client, createURL),
		collections:       newCollectionRepository(client, createURL),
		productImages:     newProductImagesRepository(client, createURL),
		metafields:        newMetafieldRepository(client, createURL),
		customers:         newCustomerRepository(client, createURL),
		customerAddresses: newCustomerAddressRepository(client, createURL),
		blogs:             newBlogRepository(client, createURL),
		articles:          newArticleRepository(client, createURL),
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

// FulfillmentOrders returns an HTTP implementation of a Shopify fulfillment order repository
func (shop Shop) FulfillmentOrders() shopify.FulfillmentOrderRepository {
	return shop.fulfillmentOrders
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
	return shop.collections
}

// ProductImages returns an HTTP implementation of a Shopify product images repository
func (shop Shop) ProductImages() shopify.ProductImageRepository {
	return shop.productImages
}

// Metafields returns an HTTP implementation of a Shopify metafield repository
func (shop Shop) Metafields() shopify.MetafieldRepository {
	return shop.metafields
}

// Customers returns an HTTP implementation of a Shopify customers repository
func (shop Shop) Customers() shopify.CustomerRepository {
	return shop.customers
}

// CustomerAddresses returns an HTTP implementation of a Shopify customer addresses repository
func (shop Shop) CustomerAddresses() shopify.CustomerAddressRepository {
	return shop.customerAddresses
}

// Blogs returns an HTTP implementation of a Shopify blog repository
func (shop Shop) Blogs() shopify.BlogRepository {
	return shop.blogs
}

// Articles returns an HTTP implementation of a Shopify article repository
func (shop Shop) Articles() shopify.ArticleRepository {
	return shop.articles
}
