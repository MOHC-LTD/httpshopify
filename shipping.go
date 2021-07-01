package httpshopify

import "github.com/MOHC-LTD/shopify"

// ShippingLineDTOs is a collection of ShippingLine DTOs
type ShippingLineDTOs []ShippingLineDTO

// ShippingLineDTO represents a Shopify shipping line in HTTP requests and responses
type ShippingLineDTO struct {
	Code               string      `json:"code"`
	Price              string      `json:"price"`
	PriceSet           PriceSetDTO `json:"price_set"`
	DiscountedPrice    string      `json:"discounted_price"`
	DiscountedPriceSet PriceSetDTO `json:"discounted_price_set"`
	ID                 int64       `json:"id"`
	Title              string      `json:"title"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ShippingLineDTO) ToShopify() shopify.ShippingLine {
	return shopify.ShippingLine{
		Code:               dto.Code,
		Price:              dto.Price,
		PriceSet:           dto.PriceSet.ToShopify(),
		DiscountedPrice:    dto.DiscountedPrice,
		DiscountedPriceSet: dto.DiscountedPriceSet.ToShopify(),
		ID:                 dto.ID,
		Title:              dto.Title,
	}
}

// ToShopify converts the DTO to the Shopify equivalent
func (dtos ShippingLineDTOs) ToShopify() shopify.ShippingLines {
	shippingLines := make(shopify.ShippingLines, 0, len(dtos))

	for _, dto := range dtos {
		shippingLines = append(shippingLines, dto.ToShopify())
	}

	return shippingLines
}
