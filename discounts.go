package httpshopify

import (
	"github.com/MOHC-LTD/shopify"
)

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

type DiscountApplicationDTOs []DiscountApplicationDTO

func (dtos DiscountApplicationDTOs) ToShopify() shopify.DiscountApplications {
	applications := make(shopify.DiscountApplications, 0, len(dtos))

	for _, dto := range dtos {
		applications = append(applications, dto.ToShopify())
	}

	return applications
}

func BuildDiscountApplicationDTOs(applications shopify.DiscountApplications) DiscountApplicationDTOs {
	dtos := make(DiscountApplicationDTOs, 0, len(applications))

	for _, application := range applications {
		dtos = append(dtos, BuildDiscountApplicationDTO(application))
	}

	return dtos
}

type DiscountAllocationDTO struct {
	Amount                   string      `json:"amount,omitempty"`
	AmountSet                PriceSetDTO `json:"amount_set,omitempty"`
	DiscountApplicationIndex int         `json:"discount_application_index,omitempty"`
}

func (dto DiscountAllocationDTO) ToShopify() shopify.DiscountAllocation {
	return shopify.DiscountAllocation{
		Amount:                   dto.Amount,
		AmountSet:                dto.AmountSet.ToShopify(),
		DiscountApplicationIndex: dto.DiscountApplicationIndex,
	}
}

func BuildDiscountAllocationDTO(allocation shopify.DiscountAllocation) DiscountAllocationDTO {
	return DiscountAllocationDTO{
		Amount:                   allocation.Amount,
		AmountSet:                BuildPriceSetDTO(allocation.AmountSet),
		DiscountApplicationIndex: allocation.DiscountApplicationIndex,
	}
}

type DiscountAllocationDTOs []DiscountAllocationDTO

func (dtos DiscountAllocationDTOs) ToShopify() shopify.DiscountAllocations {
	allocations := make(shopify.DiscountAllocations, 0, len(dtos))

	for _, dto := range dtos {
		allocations = append(allocations, dto.ToShopify())
	}

	return allocations
}

func BuildDiscountAllocationDTOs(allocations shopify.DiscountAllocations) DiscountAllocationDTOs {
	dtos := make(DiscountAllocationDTOs, 0, len(allocations))

	for _, allocation := range allocations {
		dtos = append(dtos, BuildDiscountAllocationDTO(allocation))
	}

	return dtos
}
