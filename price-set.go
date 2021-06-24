package httpshopify

// PriceSetDTO represents a price in Shopify HTTP requests and responses
type PriceSetDTO struct {
	ShopMoney MoneyDTO `json:"shop_money"`
}
