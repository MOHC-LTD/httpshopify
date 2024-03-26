package httpshopify

import (
	"encoding/json"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"
	"github.com/MOHC-LTD/shopify/v2"
)

type webhookRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newWebhookRepository(client http.Client, createURL func(endpoint string) string) webhookRepository {
	return webhookRepository{
		client,
		createURL,
	}
}

func (r webhookRepository) List() (shopify.Webhooks, error) {
	url := r.createURL("webhooks.json")

	body, _, err := r.client.Get(url, nil)
	if err != nil {
		return shopify.Webhooks{}, err
	}

	var response struct {
		Webhooks WebhookDTOs `json:"webhooks"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return shopify.Webhooks{}, err
	}

	return response.Webhooks.ToShopify(), nil
}

func (repository webhookRepository) Create(webhook shopify.Webhook) (shopify.Webhook, error) {
	createDTO := WebhookDTO{
		Topic:   webhook.Topic,
		Address: webhook.Address,
	}

	request := struct {
		Webhook WebhookDTO `json:"webhook"`
	}{
		Webhook: createDTO,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return shopify.Webhook{}, err
	}

	url := repository.createURL("webhooks.json")

	respBody, _, err := repository.client.Post(url, body, nil)
	if err != nil {
		return shopify.Webhook{}, err
	}

	responseDTO := struct {
		WebhookDTO `json:"webhook"`
	}{}

	err = json.Unmarshal(respBody, &responseDTO)
	if err != nil {
		return shopify.Webhook{}, err
	}

	return responseDTO.ToShopify(), nil
}

// WebhookDTOs represents a list of webhooks in HTTP requests and responses
type WebhookDTOs []WebhookDTO

// WebhookDTO represents a webhook in HTTP requests and responses
type WebhookDTO struct {
	ID        int64      `json:"id,omitempty"`
	Address   string     `json:"address,omitempty"`
	Topic     string     `json:"topic,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dtos WebhookDTOs) ToShopify() shopify.Webhooks {
	Webhooks := make(shopify.Webhooks, 0, len(dtos))

	for _, dto := range dtos {
		Webhooks = append(Webhooks, dto.ToShopify())
	}

	return Webhooks
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto WebhookDTO) ToShopify() shopify.Webhook {
	var createdAt time.Time
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	var updatedAt time.Time
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.Webhook{
		ID:        dto.ID,
		Address:   dto.Address,
		Topic:     dto.Topic,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
