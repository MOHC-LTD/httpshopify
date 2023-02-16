package httpshopify

import (
	"github.com/MOHC-LTD/shopify/v2"
)

// DiscountApplicationDTO represents a Shopify discount application in HTTP requests and responses
type DiscountApplicationDTO struct {
	AllocationMethod string `json:"allocation_method,omitempty"`
	Code             string `json:"code,omitempty"`
	Type             string `json:"type,omitempty"`
	Title            string `json:"title,omitempty"`
	Value            string `json:"value,omitempty"`
	ValueType        string `json:"value_type,omitempty"`
	Description      string `json:"description,omitempty"`
	TargetType       string `json:"target_type,omitempty"`
	TargetSelection  string `json:"target_selection,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto DiscountApplicationDTO) ToShopify() shopify.DiscountApplication {
	return shopify.DiscountApplication{
		AllocationMethod: dto.AllocationMethod,
		Code:             dto.Code,
		Type:             dto.Type,
		Title:            dto.Title,
		Value:            dto.Value,
		ValueType:        dto.ValueType,
		Description:      dto.Description,
		TargetType:       dto.TargetType,
		TargetSelection:  dto.TargetSelection,
	}
}

// BuildDiscountApplicationDTO builds the DTO from the Shopify equivalent
func BuildDiscountApplicationDTO(application shopify.DiscountApplication) DiscountApplicationDTO {
	return DiscountApplicationDTO{
		AllocationMethod: application.AllocationMethod,
		Code:             application.Code,
		Type:             application.Type,
		Title:            application.Title,
		Value:            application.Title,
		ValueType:        application.ValueType,
		Description:      application.Description,
		TargetType:       application.TargetType,
		TargetSelection:  application.TargetSelection,
	}
}

// DiscountApplicationDTOs is a collection of DiscountApplication DTOs
type DiscountApplicationDTOs []DiscountApplicationDTO

// ToShopify converts the DTOs to the Shopify equivalent
func (dtos DiscountApplicationDTOs) ToShopify() shopify.DiscountApplications {
	applications := make(shopify.DiscountApplications, 0, len(dtos))

	for _, dto := range dtos {
		applications = append(applications, dto.ToShopify())
	}

	return applications
}

// BuildDiscountApplicationDTOs builds the DTOs from the Shopify equivalent
func BuildDiscountApplicationDTOs(applications shopify.DiscountApplications) DiscountApplicationDTOs {
	dtos := make(DiscountApplicationDTOs, 0, len(applications))

	for _, application := range applications {
		dtos = append(dtos, BuildDiscountApplicationDTO(application))
	}

	return dtos
}

// DiscountAllocationDTO represents a Shopify discount allocation in HTTP requests and responses
type DiscountAllocationDTO struct {
	Amount                   string      `json:"amount,omitempty"`
	AmountSet                PriceSetDTO `json:"amount_set,omitempty"`
	DiscountApplicationIndex int         `json:"discount_application_index,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto DiscountAllocationDTO) ToShopify() shopify.DiscountAllocation {
	return shopify.DiscountAllocation{
		Amount:                   dto.Amount,
		AmountSet:                dto.AmountSet.ToShopify(),
		DiscountApplicationIndex: dto.DiscountApplicationIndex,
	}
}

// BuildDiscountAllocationDTO builds the DTO from the Shopify equivalent
func BuildDiscountAllocationDTO(allocation shopify.DiscountAllocation) DiscountAllocationDTO {
	return DiscountAllocationDTO{
		Amount:                   allocation.Amount,
		AmountSet:                BuildPriceSetDTO(allocation.AmountSet),
		DiscountApplicationIndex: allocation.DiscountApplicationIndex,
	}
}

// DiscountAllocationDTOs is a collection of DiscountAllocation DTOs
type DiscountAllocationDTOs []DiscountAllocationDTO

// ToShopify converts the DTOs to the Shopify equivalent
func (dtos DiscountAllocationDTOs) ToShopify() shopify.DiscountAllocations {
	allocations := make(shopify.DiscountAllocations, 0, len(dtos))

	for _, dto := range dtos {
		allocations = append(allocations, dto.ToShopify())
	}

	return allocations
}

// BuildDiscountAllocationDTOs builds the DTOs from the Shopify equivalent
func BuildDiscountAllocationDTOs(allocations shopify.DiscountAllocations) DiscountAllocationDTOs {
	dtos := make(DiscountAllocationDTOs, 0, len(allocations))

	for _, allocation := range allocations {
		dtos = append(dtos, BuildDiscountAllocationDTO(allocation))
	}

	return dtos
}
