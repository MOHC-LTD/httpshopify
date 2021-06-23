package httpshopify

import "github.com/MOHC-LTD/shopify"

type TaxLinesDTO []TaxLineDTO

type TaxLineDTO struct {
	Title string  `json:"title"`
	Rate  float32 `json:"rate"`
}

func (dto TaxLineDTO) ToShopify() shopify.Tax {
	return shopify.Tax{
		Title: dto.Title,
		Rate:  dto.Rate,
	}
}
