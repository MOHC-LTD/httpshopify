package httpshopify

import "github.com/MOHC-LTD/shopify"

// TaxLineDTOs is a collection of tax line DTOs
type TaxLineDTOs []TaxLineDTO

// TaxLineDTO represents a Shopify tax line in HTTP requests and responses
type TaxLineDTO struct {
	Title string  `json:"title"`
	Rate  float32 `json:"rate"`
}

// ToShopify converts this DTO to the Shopify equivalent
func (dto TaxLineDTO) ToShopify() shopify.Tax {
	return shopify.Tax{
		Title: dto.Title,
		Rate:  dto.Rate,
	}
}
