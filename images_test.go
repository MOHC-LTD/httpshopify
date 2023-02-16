package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/assertions"
)

// Tests image can be built when date fields are not nil
func TestImageDTO_ToShopify(t *testing.T) {
	createdAt := time.Now()

	var imageDTO = ImageDTO{
		CreatedAt: &createdAt,
	}

	image := imageDTO.ToShopify()

	if !image.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, image.CreatedAt)
	}
}

// Tests image can be built when date fields are nil
func TestImageDTO_ToShopifyEmptyTimes(t *testing.T) {
	var createdAt *time.Time

	var imageDTO = ImageDTO{
		CreatedAt: createdAt,
	}

	image := imageDTO.ToShopify()

	if !image.CreatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, createdAt, image.CreatedAt)
	}
}
