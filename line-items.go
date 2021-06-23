package httpshopify

import "github.com/MOHC-LTD/shopify"

// LineItemDTO represents a Shopify line item in HTTP requests and responses
type LineItemDTO struct {
	ID               int64       `json:"id"`
	Title            string      `json:"title"`
	Name             string      `json:"name"`
	SKU              string      `json:"sku"`
	Quantity         int         `json:"quantity"`
	PriceSet         PriceSetDTO `json:"price_set"`
	TotalDiscountSet PriceSetDTO `json:"total_discount_set"`
	TaxLines         TaxLinesDTO `json:"tax_lines"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto LineItemDTO) ToShopify() shopify.LineItem {
	return shopify.LineItem{
		ID:       dto.ID,
		Title:    dto.Title,
		Name:     dto.Name,
		SKU:      dto.SKU,
		Quantity: dto.Quantity,
		Price:    dto.PriceSet.ShopMoney.ToShopify(),
		Discount: dto.TotalDiscountSet.ShopMoney.ToShopify(),
		Taxes:    dto.TaxLines.ToShopify(),
	}
}

func buildLineItemDTO(lineItem shopify.LineItem) LineItemDTO {
	return LineItemDTO{
		ID:       lineItem.ID,
		Title:    lineItem.Title,
		Name:     lineItem.Name,
		SKU:      lineItem.SKU,
		Quantity: lineItem.Quantity,
	}
}

// LineItemDTOs is a collection of LineItem DTOs
type LineItemDTOs []LineItemDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos LineItemDTOs) ToShopify() shopify.LineItems {
	lineItems := make(shopify.LineItems, 0, len(dtos))

	for _, dto := range dtos {
		lineItems = append(lineItems, dto.ToShopify())
	}

	return lineItems
}

func buildLineItemDTOs(lineItems shopify.LineItems) []LineItemDTO {
	dtos := make([]LineItemDTO, 0, len(lineItems))

	for _, lineItem := range lineItems {
		dtos = append(dtos, buildLineItemDTO(lineItem))
	}

	return dtos
}
