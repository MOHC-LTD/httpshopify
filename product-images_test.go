package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/assertions"
	"github.com/MOHC-LTD/shopify/v2"
)

// Tests product image can be built when date fields are not nil
func TestProductImageDTO_ToShopify(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var productImageDTO = ProductImageDTO{
		ImageDTO: ImageDTO{
			CreatedAt: &createdAt,
		},
		UpdatedAt: &updatedAt,
	}

	productImage := productImageDTO.ToShopify()

	if !productImage.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, productImage.CreatedAt)
	}

	if !productImage.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, productImage.UpdatedAt)
	}
}

// Tests product image can be built when date fields are nil
func TestProductImageDTO_ToShopifyEmptyTimes(t *testing.T) {
	var createdAt *time.Time
	var updatedAt *time.Time

	var productImageDTO = ProductImageDTO{
		ImageDTO: ImageDTO{
			CreatedAt: createdAt,
		},
		UpdatedAt: updatedAt,
	}

	productImage := productImageDTO.ToShopify()

	if !productImage.CreatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, createdAt, productImage.CreatedAt)
	}

	if !productImage.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, productImage.UpdatedAt)
	}
}

// Tests that a product image DTO can be built when date fields are not nil
func TestBuildProductImageDTO(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var productImage = shopify.ProductImage{
		Image: shopify.Image{
			CreatedAt: createdAt,
		},
		UpdatedAt: updatedAt,
	}

	productImageDTO := BuildProductImageDTO(productImage)

	if !productImageDTO.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, productImageDTO.CreatedAt)
	}

	if !productImageDTO.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, productImageDTO.UpdatedAt)
	}
}

// Tests that a product image DTO can be built when date fields are nil
func TestBuildProductImageDTOEmptyTimes(t *testing.T) {
	var createdAt time.Time
	var updatedAt time.Time

	var productImage = shopify.ProductImage{
		Image: shopify.Image{
			CreatedAt: createdAt,
		},
		UpdatedAt: updatedAt,
	}

	productImageDTO := BuildProductImageDTO(productImage)

	if productImageDTO.CreatedAt != nil {
		assertions.ValueAssertionFailure(t, createdAt, productImageDTO.CreatedAt)
	}

	if productImageDTO.UpdatedAt != nil {
		assertions.ValueAssertionFailure(t, updatedAt, productImageDTO.UpdatedAt)
	}
}
