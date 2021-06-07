package httpshopify

import "github.com/MOHC-LTD/shopify"

type shippingLineDTOs []shippingLineDTO

type shippingLineDTO struct {
	ID    int64  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

func (dto shippingLineDTO) toDomain() shopify.ShippingLine {
	return shopify.ShippingLine{
		ID:    dto.ID,
		Code:  dto.Code,
		Title: dto.Title,
	}
}

func (dtos shippingLineDTOs) toDomain() []shopify.ShippingLine {
	shippingLines := make([]shopify.ShippingLine, 0, len(dtos))

	for _, dto := range dtos {
		shippingLines = append(shippingLines, dto.toDomain())
	}

	return shippingLines
}
