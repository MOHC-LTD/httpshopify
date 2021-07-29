package httpshopify

import (
	"encoding/json"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/http"

	"github.com/MOHC-LTD/shopify"
)

type inventoryLevelRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newInventoryLevelRepository(client http.Client, createURL func(endpoint string) string) inventoryLevelRepository {
	return inventoryLevelRepository{
		client,
		createURL,
	}
}

func (repository inventoryLevelRepository) Set(inventoryItemID int64, locationID int64, quantity int) (shopify.InventoryLevel, error) {
	setDTO := InventoryLevelDTO{
		InventoryItemID: inventoryItemID,
		Available:       quantity,
		LocationID:      locationID,
	}

	body, err := json.Marshal(setDTO)
	if err != nil {
		return shopify.InventoryLevel{}, err
	}

	url := repository.createURL("inventory_levels/set.json")

	respBody, _, err := repository.client.Post(url, body, nil)
	if err != nil {
		return shopify.InventoryLevel{}, err
	}

	var response struct {
		InventoryLevel InventoryLevelDTO `json:"inventory_level"`
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.InventoryLevel{}, err
	}

	return response.InventoryLevel.ToShopify(), nil
}

// InventoryLevelDTO represents a Shopify inventory level in HTTP requests and responses
type InventoryLevelDTO struct {
	InventoryItemID int64      `json:"inventory_item_id,omitempty"`
	Available       int        `json:"available,omitempty"`
	LocationID      int64      `json:"location_id,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto InventoryLevelDTO) ToShopify() shopify.InventoryLevel {
	var updatedAt time.Time
	if !dto.UpdatedAt.IsZero() {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.InventoryLevel{
		InventoryItemID: dto.InventoryItemID,
		Available:       dto.Available,
		LocationID:      dto.LocationID,
		UpdatedAt:       updatedAt,
	}
}
