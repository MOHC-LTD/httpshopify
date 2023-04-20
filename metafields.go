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

	return response.Metafields.toShopify()
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

type metafieldsDTO []metafieldDTO

func (dto metafieldsDTO) toShopify() (shopify.Metafields, error) {
	metafields := shopify.Metafields{}

	for _, metafieldDTO := range dto {
		m, err := metafieldDTO.toShopify()
		if err != nil {
			return nil, err
		}

		metafields = append(metafields, m)
	}

	return metafields, nil
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

func (dto metafieldDTO) toShopify() (shopify.Metafield, error) {
	var createdAt time.Time
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	var updatedAt time.Time
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	value := dto.Value
	switch dto.Type {
	case shopify.ListSingleLineTextFieldMetaFieldType:
		v, ok := dto.Value.(string)
		if !ok {
			return shopify.Metafield{}, ErrMetafieldValueType
		}

		var d []string
		err := json.Unmarshal([]byte(v), &d)
		if err != nil {
			return shopify.Metafield{}, err
		}

		value = d
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
		Value:     value,
		Type:      dto.Type,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
