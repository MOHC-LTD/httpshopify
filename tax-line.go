package httpshopify

import "github.com/MOHC-LTD/shopify"

// TaxLineDTOs is a collection of tax line DTOs
type TaxLineDTOs []TaxLineDTO

// ToShopify converts this DTO to a list of Shopify tax lines
func (dto TaxLineDTOs) ToShopify() []shopify.Tax {
	taxes := make([]shopify.Tax, 0, len(dto))

	for _, tax := range dto {
		taxes = append(taxes, tax.ToShopify())
	}

	return taxes
}

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
