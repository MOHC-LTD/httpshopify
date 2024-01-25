package httpshopify

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"
	"github.com/MOHC-LTD/shopify/v2"
)

type metafieldRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newMetafieldRepository(client http.Client, createURL func(endpoint string) string) metafieldRepository {
	return metafieldRepository{client, createURL}
}

func (repository metafieldRepository) List(query shopify.MetafieldQuery) (shopify.Metafields, error) {

	url := repository.createURL(fmt.Sprintf("metafields.json?%s", parseMetafieldQuery(query)))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.Metafields{}, err
	}

	var response struct {
		Metafields metafieldsDTO `json:"metafields"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return shopify.Metafields{}, err
	}

	return response.Metafields.toShopify(), nil
}

func parseMetafieldQuery(query shopify.MetafieldQuery) string {
	params := url.Values{}

	if query.Resource.OwnerID != 0 {
		params.Add("metafield[owner_id]", strconv.FormatInt(query.Resource.OwnerID, 10))
	}

	if query.Resource.OwnerResource != "" {
		params.Add("metafield[owner_resource]", string(query.Resource.OwnerResource))
	}

	return params.Encode()
}

// MetafieldDTO represents a Shopify metafield in HTTP requests and responses
type MetafieldDTO struct {
	Key       string      `json:"key,omitempty"`
	Namespace string      `json:"namespace,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Type      string      `json:"type,omitempty"`
}

// MetafieldsDTO represents Shopify metafields in HTTP requests and responses
type MetafieldsDTO []MetafieldDTO

// BuildMetafieldDTOs builds the DTOs from the Shopify equivalent
func BuildMetafieldDTOs(metafields shopify.Metafields) MetafieldsDTO {
	dtos := make(MetafieldsDTO, 0, len(metafields))

	for _, metafield := range metafields {
		metafieldDTO := MetafieldDTO{
			Key:       metafield.Key,
			Namespace: metafield.Namespace,
			Value:     metafield.Value,
			Type:      metafield.Type,
		}

		dtos = append(dtos, metafieldDTO)
	}

	return dtos
}

type metafieldsDTO []metafieldDTO

func (dto metafieldsDTO) toShopify() shopify.Metafields {
	metafields := shopify.Metafields{}

	for _, metafieldDTO := range dto {
		metafields = append(metafields, metafieldDTO.toShopify())
	}

	return metafields
}

type metafieldDTO struct {
	ID          int64  `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Key         string `json:"key"`
	Namespace   string `json:"namespace"`
	OwnerID     int64  `json:"owner_id,omitempty"`
	// TODO This is better in a later version as it is always a string. Consider versioning this package with a new Shopify version.
	Value         interface{} `json:"value"`
	Type          string      `json:"type"`
	OwnerResource string      `json:"owner_resource,omitempty"`
	CreatedAt     *time.Time  `json:"created_at,omitempty"`
	UpdatedAt     *time.Time  `json:"updated_at,omitempty"`
}

func (dto metafieldDTO) toShopify() shopify.Metafield {
	var createdAt time.Time
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	var updatedAt time.Time
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.Metafield{
		ID:          dto.ID,
		Description: dto.Description,
		Key:         dto.Key,
		Namespace:   dto.Namespace,
		Resource: shopify.MetafieldResource{
			OwnerID:       dto.OwnerID,
			OwnerResource: dto.OwnerResource,
		},
		Value:     fmt.Sprintf("%v", dto.Value),
		Type:      dto.Type,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
