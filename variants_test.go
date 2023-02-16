package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/assertions"
	"github.com/MOHC-LTD/shopify/v2"
)

// Tests variant can be built when date fields are not nil
func TestVariantDTO_ToShopify(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var variantDTO = VariantDTO{
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}

	variant := variantDTO.ToShopify()

	if !variant.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, variant.CreatedAt)
	}

	if !variant.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, variant.UpdatedAt)
	}
}

// Tests variant can be built when date fields are nil
func TestVariantDTO_ToShopifyEmptyTimes(t *testing.T) {
	var createdAt *time.Time
	var updatedAt *time.Time

	var variantDTO = VariantDTO{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	variant := variantDTO.ToShopify()

	if !variant.CreatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, createdAt, variant.CreatedAt)
	}

	if !variant.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, variant.UpdatedAt)
	}
}

// Tests that a variant DTO can be built when date fields are not nil
func TestBuildVariantDTOs(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var variants = []shopify.Variant{
		{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}

	variantDTOs := BuildVariantDTOs(variants)

	if !variantDTOs[0].CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, variantDTOs[0].CreatedAt)
	}

	if !variantDTOs[0].UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, variantDTOs[0].UpdatedAt)
	}
}

// Tests that a variant DTO can be built when date fields are nil
func TestBuildVariantDTOsEmptyTimes(t *testing.T) {
	var createdAt time.Time
	var updatedAt time.Time

	var variants = []shopify.Variant{
		{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}

	variantDTOs := BuildVariantDTOs(variants)

	if variantDTOs[0].CreatedAt != nil {
		assertions.ValueAssertionFailure(t, createdAt, variantDTOs[0].CreatedAt)
	}

	if variantDTOs[0].UpdatedAt != nil {
		assertions.ValueAssertionFailure(t, updatedAt, variantDTOs[0].UpdatedAt)
	}
}
