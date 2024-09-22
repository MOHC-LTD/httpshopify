package httpshopify

import "github.com/MOHC-LTD/shopify/v2"

// LineItemDTO represents a Shopify line item in HTTP requests and responses
type LineItemDTO struct {
	ID                  int64                  `json:"id,omitempty"`
	Price               string                 `json:"price,omitempty"`
	PriceSet            PriceSetDTO            `json:"price_set,omitempty"`
	ProductID           int64                  `json:"product_id,omitempty"`
	Quantity            int                    `json:"quantity,omitempty"`
	RequiresShipping    bool                   `json:"requires_shipping,omitempty"`
	SKU                 string                 `json:"sku,omitempty"`
	Title               string                 `json:"title,omitempty"`
	VariantID           int64                  `json:"variant_id,omitempty"`
	VariantTitle        string                 `json:"variant_title,omitempty"`
	Name                string                 `json:"name,omitempty"`
	TaxLines            TaxLineDTOs            `json:"tax_lines,omitempty"`
	TotalDiscount       string                 `json:"total_discount,omitempty"`
	TotalDiscountSet    PriceSetDTO            `json:"total_discount_set,omitempty"`
	DiscountAllocations DiscountAllocationDTOs `json:"discount_allocations,omitempty"`
	Properties          PropertyDTOs           `json:"properties,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto LineItemDTO) ToShopify() shopify.LineItem {
	return shopify.LineItem{
		ID:                  dto.ID,
		Price:               dto.Price,
		PriceSet:            dto.PriceSet.ToShopify(),
		ProductID:           dto.ProductID,
		Quantity:            dto.Quantity,
		RequiresShipping:    dto.RequiresShipping,
		SKU:                 dto.SKU,
		Title:               dto.Title,
		VariantID:           dto.VariantID,
		VariantTitle:        dto.VariantTitle,
		Name:                dto.Name,
		TaxLines:            dto.TaxLines.ToShopify(),
		TotalDiscount:       dto.TotalDiscount,
		TotalDiscountSet:    dto.TotalDiscountSet.ToShopify(),
		DiscountAllocations: dto.DiscountAllocations.ToShopify(),
		Properties:          dto.Properties.ToShopify(),
	}
}

// BuildLineItemDTO builds the DTO from the Shopify equivalent
func BuildLineItemDTO(lineItem shopify.LineItem) LineItemDTO {
	return LineItemDTO{
		ID:                  lineItem.ID,
		Price:               lineItem.Price,
		PriceSet:            BuildPriceSetDTO(lineItem.PriceSet),
		ProductID:           lineItem.ProductID,
		Quantity:            lineItem.Quantity,
		RequiresShipping:    lineItem.RequiresShipping,
		SKU:                 lineItem.SKU,
		Title:               lineItem.Title,
		VariantID:           lineItem.VariantID,
		VariantTitle:        lineItem.VariantTitle,
		Name:                lineItem.Name,
		TaxLines:            BuildTaxLineDTOs(lineItem.TaxLines),
		TotalDiscount:       lineItem.TotalDiscount,
		TotalDiscountSet:    BuildPriceSetDTO(lineItem.TotalDiscountSet),
		DiscountAllocations: BuildDiscountAllocationDTOs(lineItem.DiscountAllocations),
		Properties:          BuildPropertyDTOs(lineItem.Properties),
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

// BuildLineItemDTOs builds the DTOs from the Shopify equivalent
func BuildLineItemDTOs(lineItems shopify.LineItems) []LineItemDTO {
	dtos := make([]LineItemDTO, 0, len(lineItems))

	for _, lineItem := range lineItems {
		dtos = append(dtos, BuildLineItemDTO(lineItem))
	}

	return dtos
}
