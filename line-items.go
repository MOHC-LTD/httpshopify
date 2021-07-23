package httpshopify

import "github.com/MOHC-LTD/shopify"

// LineItemDTO represents a Shopify line item in HTTP requests and responses
type LineItemDTO struct {
	ID               int64       `json:"id"`
	Price            string      `json:"price"`
	PriceSet         PriceSetDTO `json:"price_set"`
	ProductID        int64       `json:"product_id"`
	Quantity         int         `json:"quantity"`
	SKU              string      `json:"sku"`
	Title            string      `json:"title"`
	VariantID        int64       `json:"variant_id"`
	VariantTitle     string      `json:"variant_title"`
	Name             string      `json:"name"`
	TaxLines         TaxLineDTOs `json:"tax_lines"`
	TotalDiscount    string      `json:"total_discount"`
	TotalDiscountSet PriceSetDTO `json:"total_discount_set"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto LineItemDTO) ToShopify() shopify.LineItem {
	return shopify.LineItem{
		ID:               dto.ID,
		Price:            dto.Price,
		PriceSet:         dto.PriceSet.ToShopify(),
		ProductID:        dto.ProductID,
		Quantity:         dto.Quantity,
		SKU:              dto.SKU,
		Title:            dto.Title,
		VariantID:        dto.VariantID,
		VariantTitle:     dto.VariantTitle,
		Name:             dto.Name,
		TaxLines:         dto.TaxLines.ToShopify(),
		TotalDiscount:    dto.TotalDiscount,
		TotalDiscountSet: dto.TotalDiscountSet.ToShopify(),
	}
}

func buildLineItemDTO(lineItem shopify.LineItem) LineItemDTO {
	return LineItemDTO{
		ID:               lineItem.ID,
		Price:            lineItem.Price,
		PriceSet:         BuildPriceSetDTO(lineItem.PriceSet),
		ProductID:        lineItem.ProductID,
		Quantity:         lineItem.Quantity,
		SKU:              lineItem.SKU,
		Title:            lineItem.Title,
		VariantID:        lineItem.VariantID,
		VariantTitle:     lineItem.VariantTitle,
		Name:             lineItem.Name,
		TaxLines:         BuildTaxLineDTOs(lineItem.TaxLines),
		TotalDiscount:    lineItem.TotalDiscount,
		TotalDiscountSet: BuildPriceSetDTO(lineItem.TotalDiscountSet),
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

func BuildLineItemDTOs(lineItems shopify.LineItems) []LineItemDTO {
	dtos := make([]LineItemDTO, 0, len(lineItems))

	for _, lineItem := range lineItems {
		dtos = append(dtos, buildLineItemDTO(lineItem))
	}

	return dtos
}
