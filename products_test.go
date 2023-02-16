package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/assertions"
)

// Tests product can be built when date fields are not nil
func TestProductDTO_ToShopify(t *testing.T) {
	createdAt := time.Now()
	publishedAt := time.Now()
	updatedAt := time.Now()

	var productDTO = ProductDTO{
		CreatedAt:   &createdAt,
		PublishedAt: &publishedAt,
		UpdatedAt:   &updatedAt,
	}

	product := productDTO.ToShopify()

	if !product.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, product.CreatedAt)
	}

	if !product.PublishedAt.Equal(publishedAt) {
		assertions.ValueAssertionFailure(t, publishedAt, product.PublishedAt)
	}

	if !product.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, product.UpdatedAt)
	}
}

// Tests product can be built when date fields are nil
func TestProductDTO_ToShopifyEmptyTimes(t *testing.T) {
	var createdAt *time.Time
	var publishedAt *time.Time
	var updatedAt *time.Time

	var productDTO = ProductDTO{
		CreatedAt:   createdAt,
		PublishedAt: publishedAt,
		UpdatedAt:   updatedAt,
	}

	product := productDTO.ToShopify()

	if !product.CreatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, createdAt, product.CreatedAt)
	}

	if !product.PublishedAt.IsZero() {
		assertions.ValueAssertionFailure(t, publishedAt, product.PublishedAt)
	}

	if !product.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, product.UpdatedAt)
	}
}
