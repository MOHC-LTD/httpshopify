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
	setRequest := InventoryLevelSetRequest{
		InventoryItemID: inventoryItemID,
		Available:       quantity,
		LocationID:      locationID,
	}

	body, err := json.Marshal(setRequest)
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

// InventoryLevelDTO represents a Shopify inventory level in HTTP requests and responses - READ ONLY
type InventoryLevelDTO struct {
	InventoryItemID int64      `json:"inventory_item_id"`
	Available       int        `json:"available"`
	LocationID      int64      `json:"location_id"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

// InventoryLevelSetRequest represents a Shopify inventory level in HTTP requests and responses - WRITE ONLY
type InventoryLevelSetRequest struct {
	InventoryItemID int64 `json:"inventory_item_id,omitempty"`
	Available       int   `json:"available"`
	LocationID      int64 `json:"location_id,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto InventoryLevelDTO) ToShopify() shopify.InventoryLevel {
	var updatedAt time.Time
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.InventoryLevel{
		InventoryItemID: dto.InventoryItemID,
		Available:       dto.Available,
		LocationID:      dto.LocationID,
		UpdatedAt:       updatedAt,
	}
}
