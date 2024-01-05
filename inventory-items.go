package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"

	"github.com/MOHC-LTD/shopify/v2"
)

type inventoryItemRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newInventoryItemRepository(client http.Client, createURL func(endpoint string) string) inventoryItemRepository {
	return inventoryItemRepository{
		client,
		createURL,
	}
}

func (repository inventoryItemRepository) Get(id int64) (shopify.InventoryItem, error) {
	url := repository.createURL(fmt.Sprintf("inventory_items/%v.json", id))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.InventoryItem{}, err
	}

	var resultDTO struct {
		InventoryItem InventoryItemDTO `json:"inventory_item"`
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.InventoryItem.ToShopify(), nil
}

// CountryHarmonizedSystemCodeDTO represents a Shopify country system code in HTTP requests and responses
type CountryHarmonizedSystemCodeDTO struct {
	HarmonizedSystemCode string `json:"harmonized_system_code,omitempty"`
	CountryCode          string `json:"country_code,omitempty"`
}

func (dto CountryHarmonizedSystemCodeDTO) ToShopify() shopify.CountryHarmonizedSystemCode {
	return shopify.CountryHarmonizedSystemCode{
		HarmonizedSystemCode: dto.HarmonizedSystemCode,
		CountryCode:          dto.CountryCode,
	}
}

// CountryHarmonizedSystemCodeDTOs is a collection of Country Harmonized System Code DTOs
type CountryHarmonizedSystemCodeDTOs []CountryHarmonizedSystemCodeDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos CountryHarmonizedSystemCodeDTOs) ToShopify() []shopify.CountryHarmonizedSystemCode {
	countryHarmonizedSystemCode := make([]shopify.CountryHarmonizedSystemCode, 0, len(dtos))

	for _, dto := range dtos {
		countryHarmonizedSystemCode = append(countryHarmonizedSystemCode, dto.ToShopify())
	}

	return countryHarmonizedSystemCode
}

// InventoryItemDTO represents a Shopify inventory item in HTTP requests and responses
type InventoryItemDTO struct {
	Cost                         string                          `json:"cost,omitempty"`
	CountryCodeOfOrigin          string                          `json:"country_code_of_origin,omitempty"`
	CountryHarmonizedSystemCodes CountryHarmonizedSystemCodeDTOs `json:"country_harmonized_system_codes,omitempty"`
	CreatedAt                    *time.Time                      `json:"created_at,omitempty"`
	HarmonizedSystemCode         int64                           `json:"harmonized_system_code,omitempty"`
	ID                           int64                           `json:"id,omitempty"`
	ProvinceCodeOfOrigin         string                          `json:"province_code_of_origin,omitempty"`
	SKU                          string                          `json:"sku,omitempty"`
	Tracked                      bool                            `json:"tracked,omitempty"`
	UpdatedAt                    *time.Time                      `json:"updated_at,omitempty"`
	RequiresShipping             bool                            `json:"requires_shipping,omitempty"`
	AdminGraphqlApiId            string                          `json:"admin_graphql_api_id,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto InventoryItemDTO) ToShopify() shopify.InventoryItem {
	var createdAt time.Time
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	var updatedAt time.Time
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.InventoryItem{
		ID:                           dto.ID,
		SKU:                          dto.SKU,
		CreatedAt:                    createdAt,
		UpdatedAt:                    updatedAt,
		RequiresShipping:             dto.RequiresShipping,
		Cost:                         dto.Cost,
		CountryCodeOfOrigin:          dto.CountryCodeOfOrigin,
		ProvinceCodeOfOrigin:         dto.ProvinceCodeOfOrigin,
		HarmonizedSystemCode:         dto.HarmonizedSystemCode,
		Tracked:                      dto.Tracked,
		CountryHarmonizedSystemCodes: dto.CountryHarmonizedSystemCodes.ToShopify(),
		AdminGraphqlApiId:            dto.AdminGraphqlApiId,
	}
}
