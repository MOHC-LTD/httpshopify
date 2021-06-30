package httpshopify

import "github.com/MOHC-LTD/shopify"

// PriceSetDTO represents a price set in Shopify HTTP requests and responses
type PriceSetDTO struct {
	ShopMoney        MoneyDTO `json:"shop_money"`
	PresentmentMoney MoneyDTO `json:"presentment_money"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto PriceSetDTO) ToShopify() shopify.PriceSet {
	return shopify.PriceSet{
		ShopMoney:        dto.ShopMoney.ToShopify(),
		PresentmentMoney: dto.PresentmentMoney.ToShopify(),
	}
}

// BuildPriceSetDTO builds the DTO from the Shopify equivalent
func BuildPriceSetDTO(priceSet shopify.PriceSet) PriceSetDTO {
	return PriceSetDTO{
		ShopMoney:        BuildMoneyDTO(priceSet.ShopMoney),
		PresentmentMoney: BuildMoneyDTO(priceSet.PresentmentMoney),
	}
}

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

// BuildMoneyDTO builds the DTO from the Shopify equivalent
func BuildMoneyDTO(money shopify.Money) MoneyDTO {
	return MoneyDTO{
		Amount:       money.Amount,
		CurrencyCode: money.CurrencyCode,
	}
}
