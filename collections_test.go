package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/assertions"
)

// Tests that collection cab be built correctly when date fields are not nil
func TestCollectionDTO_ToShopify(t *testing.T) {
	updatedAt := time.Now()
	publishedAt := time.Now()

	var collectionDTO = CollectionDTO{
		PublishedAt: &publishedAt,
		UpdatedAt:   &updatedAt,
	}

	collection := collectionDTO.ToShopify()

	if !collection.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, collection.UpdatedAt)
	}

	if !collection.PublishedAt.Equal(publishedAt) {
		assertions.ValueAssertionFailure(t, publishedAt, collection.PublishedAt)
	}
}

// Tests collection can be built when date fields are nil
func TestCollectionDTO_ToShopifyEmptyTimes(t *testing.T) {
	var updatedAt *time.Time
	var publishedAt *time.Time

	var collectionDTO = CollectionDTO{
		PublishedAt: publishedAt,
		UpdatedAt:   updatedAt,
	}

	collection := collectionDTO.ToShopify()

	if !collection.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, collection.UpdatedAt)
	}

	if !collection.PublishedAt.IsZero() {
		assertions.ValueAssertionFailure(t, publishedAt, collection.PublishedAt)
	}
}
