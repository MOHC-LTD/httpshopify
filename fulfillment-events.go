package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/http"

	"github.com/MOHC-LTD/shopify"
)

type fulfillmentEventRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newFulfillmentEventRepository(client http.Client, createURL func(endpoint string) string) fulfillmentEventRepository {
	return fulfillmentEventRepository{
		client,
		createURL,
	}
}

func (repository fulfillmentEventRepository) Create(orderID int64, fulfillmentID int64, event shopify.FulfillmentEvent) (shopify.FulfillmentEvent, error) {
	createDTO := FulfillmentEventDTO{
		Status: event.Status,
	}

	request := struct {
		FulfillmentEvent FulfillmentEventDTO `json:"event"`
	}{
		FulfillmentEvent: createDTO,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return shopify.FulfillmentEvent{}, err
	}

	url := repository.createURL(fmt.Sprintf("orders/%v/fulfillments/%v/events.json", orderID, fulfillmentID))

	respBody, _, err := repository.client.Post(url, body, nil)
	if err != nil {
		return shopify.FulfillmentEvent{}, err
	}

	var response struct {
		FulfillmentEvent FulfillmentEventDTO `json:"fulfillment_event"`
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.FulfillmentEvent{}, err
	}

	return response.FulfillmentEvent.ToShopify(), nil
}

func (repository fulfillmentEventRepository) Delete(orderID int64, fulfillmentID int64, eventID int64) error {
	url := repository.createURL(fmt.Sprintf("orders/%v/fulfillments/%v/events/%v.json", orderID, fulfillmentID, eventID))

	_, _, err := repository.client.Delete(url, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repository fulfillmentEventRepository) List(orderID int64, fulfillmentID int64) ([]shopify.FulfillmentEvent, error) {
	url := repository.createURL(fmt.Sprintf("orders/%v/fulfillments/%v/events.json", orderID, fulfillmentID))

	respBody, _, err := repository.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		FulfillmentEvents []FulfillmentEventDTO `json:"fulfillment_events"`
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, err
	}

	events := make([]shopify.FulfillmentEvent, 0, len(response.FulfillmentEvents))
	for _, dto := range response.FulfillmentEvents {
		events = append(events, dto.ToShopify())
	}

	return events, nil
}

// FulfillmentEventDTO represents a Shopify fulfillment event in HTTP requests and responses
type FulfillmentEventDTO struct {
	ID            int64      `json:"id,omitempty"`
	FulfillmentID int64      `json:"fulfillment_id,omitempty"`
	Status        string     `json:"status,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto FulfillmentEventDTO) ToShopify() shopify.FulfillmentEvent {

	if dto.CreatedAt.IsZero() {
		dto.CreatedAt = nil
	}

	if dto.UpdatedAt.IsZero() {
		dto.UpdatedAt = nil
	}

	return shopify.FulfillmentEvent{
		ID:            dto.ID,
		FulfillmentID: dto.FulfillmentID,
		Status:        dto.Status,
		CreatedAt:     *dto.CreatedAt,
		UpdatedAt:     *dto.UpdatedAt,
	}
}
