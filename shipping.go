package httpshopify

import "github.com/MOHC-LTD/shopify"

// ShippingLineDTOs is a collection of ShippingLine DTOs
type ShippingLineDTOs []ShippingLineDTO

// ShippingLineDTO represents a Shopify shipping line in HTTP requests and responses
type ShippingLineDTO struct {
	ID          int64       `json:"id"`
	Code        string      `json:"code"`
	Title       string      `json:"title"`
	PriceSet    PriceSetDTO `json:"price_set"`
	DiscountSet PriceSetDTO `json:"discount_set"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ShippingLineDTO) ToShopify() shopify.ShippingLine {
	return shopify.ShippingLine{
		ID:    dto.ID,
		Code:  dto.Code,
		Title: dto.Title,
	}
}

// ToShopify converts the DTO to the Shopify equivalent
func (dtos ShippingLineDTOs) ToShopify() []shopify.ShippingLine {
	shippingLines := make([]shopify.ShippingLine, 0, len(dtos))

	for _, dto := range dtos {
		shippingLines = append(shippingLines, dto.ToShopify())
	}

	return shippingLines
}
