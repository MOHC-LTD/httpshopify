package httpshopify

import "github.com/MOHC-LTD/shopify"

// MoneyDTO represents a quantity of money in HTTP requests and responses
type MoneyDTO struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

// ToShopify converts this DTO to the Shopify equivalent
func (dto MoneyDTO) ToShopify() shopify.Money {
	return shopify.Money{
		Amount:       dto.Amount,
		CurrencyCode: dto.CurrencyCode,
	}
}
