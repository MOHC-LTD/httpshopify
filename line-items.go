package httpshopify

import "github.com/MOHC-LTD/shopify"

type lineItemDTO struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Name     string `json:"name"`
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

func (dto lineItemDTO) toDomain() shopify.LineItem {
	return shopify.LineItem{
		ID:       dto.ID,
		Title:    dto.Title,
		Name:     dto.Name,
		SKU:      dto.SKU,
		Quantity: dto.Quantity,
	}
}

func buildLineItemDTO(lineItem shopify.LineItem) lineItemDTO {
	return lineItemDTO{
		ID:       lineItem.ID,
		Title:    lineItem.Title,
		Name:     lineItem.Name,
		SKU:      lineItem.SKU,
		Quantity: lineItem.Quantity,
	}
}

type lineItemDTOs []lineItemDTO

func (dtos lineItemDTOs) toDomain() shopify.LineItems {
	lineItems := make(shopify.LineItems, 0, len(dtos))

	for _, dto := range dtos {
		lineItems = append(lineItems, dto.toDomain())
	}

	return lineItems
}

func buildLineItemDTOs(lineItems shopify.LineItems) []lineItemDTO {
	dtos := make([]lineItemDTO, 0, len(lineItems))

	for _, lineItem := range lineItems {
		dtos = append(dtos, buildLineItemDTO(lineItem))
	}

	return dtos
}
