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

// VariantDTO represents a Shopify variant in HTTP requests and responses
type VariantDTO struct {
	ID              int64     `json:"id"`
	SKU             string    `json:"sku"`
	Title           string    `json:"title"`
	InventoryItemID int64     `json:"inventory_item_id"`
	Price           string    `json:"price"`
	ProductID       int64     `json:"product_id"`
	Barcode         string    `json:"barcode"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto VariantDTO) ToShopify() shopify.Variant {
	return shopify.Variant{
		ID:              dto.ID,
		SKU:             dto.SKU,
		Title:           dto.Title,
		InventoryItemID: dto.InventoryItemID,
		Price:           dto.Price,
		ProductID:       dto.ProductID,
		Barcode:         dto.Barcode,
		CreatedAt:       dto.CreatedAt,
		UpdatedAt:       dto.UpdatedAt,
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
