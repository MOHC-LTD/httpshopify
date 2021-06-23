package httpshopify

import "github.com/MOHC-LTD/shopify"

type MoneyDTO struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

func (dto MoneyDTO) ToShopify() shopify.Money {
	return shopify.Money{
		Amount:       dto.Amount,
		CurrencyCode: dto.CurrencyCode,
	}
}
