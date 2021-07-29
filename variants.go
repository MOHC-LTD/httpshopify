package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/http"

	"github.com/MOHC-LTD/shopify"
)

type variantRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newVariantRepository(client http.Client, createURL func(endpoint string) string) variantRepository {
	return variantRepository{
		client,
		createURL,
	}
}

// VariantDTOs is a collection of Variant DTOs
type VariantDTOs []VariantDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos VariantDTOs) ToShopify() shopify.Variants {
	variants := make(shopify.Variants, 0, len(dtos))

	for _, dto := range dtos {
		variants = append(variants, dto.ToShopify())
	}

	return variants
}

// BuildVariantDTOs builds the DTOs from the Shopify equivalent
func BuildVariantDTOs(variants shopify.Variants) VariantDTOs {
	dtos := make([]VariantDTO, 0, len(variants))

	for _, variant := range variants {
		var createdAt *time.Time
		if !variant.CreatedAt.IsZero() {
			createdAt = &variant.CreatedAt
		}

		var updatedAt *time.Time
		if !variant.UpdatedAt.IsZero() {
			updatedAt = &variant.UpdatedAt
		}

		variantDTO := VariantDTO{
			ID:              variant.ID,
			SKU:             variant.SKU,
			Title:           variant.Title,
			InventoryItemID: variant.InventoryItemID,
			Price:           variant.Price,
			Barcode:         variant.Barcode,
			CreatedAt:       createdAt,
			UpdatedAt:       updatedAt,
		}

		dtos = append(dtos, variantDTO)
	}

	return dtos
}

// VariantDTO represents a Shopify variant in HTTP requests and responses
type VariantDTO struct {
	ID              int64      `json:"id,omitempty"`
	SKU             string     `json:"sku,omitempty"`
	Title           string     `json:"title,omitempty"`
	InventoryItemID int64      `json:"inventory_item_id,omitempty"`
	Price           string     `json:"price,omitempty"`
	Barcode         string     `json:"barcode,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto VariantDTO) ToShopify() shopify.Variant {
	var createdAt time.Time
	if !dto.CreatedAt.IsZero() {
		createdAt = *dto.CreatedAt
	}

	var updatedAt time.Time
	if !dto.UpdatedAt.IsZero() {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.Variant{
		ID:              dto.ID,
		SKU:             dto.SKU,
		Title:           dto.Title,
		InventoryItemID: dto.InventoryItemID,
		Price:           dto.Price,
		Barcode:         dto.Barcode,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}

func (repository variantRepository) Get(id int64) (shopify.Variant, error) {
	url := repository.createURL(fmt.Sprintf("variants/%v.json", id))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.Variant{}, err
	}

	var response struct {
		Variant VariantDTO `json:"variant"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return shopify.Variant{}, err
	}

	if response.Variant.ID == 0 {
		return shopify.Variant{}, shopify.NewErrVariantNotFound(id)
	}

	return response.Variant.ToShopify(), nil
}
