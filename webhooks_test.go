package httpshopify

import (
	"testing"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/assertions"
)

// Tests that webhook can be built correctly when date fields are not nil
func TestWebhookDTO_ToShopify(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now()

	var webhookDTO = WebhookDTO{
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}

	webhook := webhookDTO.ToShopify()

	if !webhook.CreatedAt.Equal(createdAt) {
		assertions.ValueAssertionFailure(t, createdAt, webhook.CreatedAt)
	}

	if !webhook.UpdatedAt.Equal(updatedAt) {
		assertions.ValueAssertionFailure(t, updatedAt, webhook.UpdatedAt)
	}
}

// Tests that webhook can be built correctly when date fields are nil
func TestWebhookDTO_ToShopifyEmptyFields(t *testing.T) {
	var createdAt *time.Time
	var updatedAt *time.Time

	var webhookDTO = WebhookDTO{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	webhook := webhookDTO.ToShopify()

	if !webhook.CreatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, createdAt, webhook.CreatedAt)
	}

	if !webhook.UpdatedAt.IsZero() {
		assertions.ValueAssertionFailure(t, updatedAt, webhook.UpdatedAt)
	}
}
