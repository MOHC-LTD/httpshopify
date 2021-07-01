package httpshopify

import "github.com/MOHC-LTD/shopify"

// TaxLineDTOs is a collection of tax line DTOs
type TaxLineDTOs []TaxLineDTO

// ToShopify converts this DTO to a list of Shopify tax lines
func (dto TaxLineDTOs) ToShopify() []shopify.TaxLine {
	taxes := make([]shopify.TaxLine, 0, len(dto))

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
func (dto TaxLineDTO) ToShopify() shopify.TaxLine {
	return shopify.TaxLine{
		Title: dto.Title,
		Rate:  dto.Rate,
	}
}

// BuildTaxLineDTOs builds the DTO from the Shopify equivalent
func BuildTaxLineDTOs(taxLines shopify.TaxLines) TaxLineDTOs {
	dtos := make(TaxLineDTOs, 0, len(taxLines))

	for _, taxLine := range taxLines {
		dtos = append(dtos, BuildTaxLineDTO(taxLine))
	}

	return dtos
}

// BuildTaxLineDTO builds the DTO from the Shopify equivalent
func BuildTaxLineDTO(taxLine shopify.TaxLine) TaxLineDTO {
	return TaxLineDTO{
		Title: taxLine.Title,
		Rate:  taxLine.Rate,
	}
}
