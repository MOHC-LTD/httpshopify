package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/http"
	"github.com/MOHC-LTD/shopify"
)

type collectionRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newCollectionRepository(client http.Client, createURL func(endpoint string) string) collectionRepository {
	return collectionRepository{
		client,
		createURL,
	}
}

func (repository collectionRepository) Get(id int64) (shopify.Collection, error) {
	url := repository.createURL(fmt.Sprintf("collections/%v.json", id))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.Collection{}, err
	}

	var resultDTO struct {
		Collection CollectionDTO `json:"collection"`
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.Collection.ToShopify(), nil
}

func (repository collectionRepository) Products(id int64) (shopify.Products, error) {

	url := repository.createURL(fmt.Sprintf("collections/%v/products.json", id))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var resultDTO struct {
		Products ProductDTOs `json:"products"`
	}
	json.Unmarshal(body, &resultDTO)

	return resultDTO.Products.ToShopify(), nil
}

// CollectionDTOs represents a list of shopify collections in HTTP requests and responses
type CollectionDTOs []CollectionDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos CollectionDTOs) ToShopify() shopify.Collections {
	collections := make(shopify.Collections, 0, len(dtos))

	for _, dto := range dtos {
		collections = append(collections, dto.ToShopify())
	}

	return collections
}

// CollectionDTO represents a Shopify collection in HTTP requests and responses
type CollectionDTO struct {
	BodyHTML       string     `json:"body_html,omitempty"`
	Handle         string     `json:"handle,omitempty"`
	Image          ImageDTO   `json:"image,omitempty"`
	ID             int64      `json:"id,omitempty"`
	PublishedAt    *time.Time `json:"published_at,omitempty"`
	PublishedScope string     `json:"published_scope,omitempty"`
	SortOrder      string     `json:"sort_order,omitempty"`
	TemplateSuffix string     `json:"template_suffix,omitempty"`
	Title          string     `json:"title,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto CollectionDTO) ToShopify() shopify.Collection {
	var publishedAt time.Time
	if dto.PublishedAt != nil {
		publishedAt = *dto.PublishedAt
	}

	var updatedAt time.Time
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.Collection{
		BodyHTML:       dto.BodyHTML,
		Handle:         dto.Handle,
		Image:          dto.Image.ToShopify(),
		ID:             dto.ID,
		PublishedAt:    publishedAt,
		PublishedScope: dto.PublishedScope,
		SortOrder:      dto.SortOrder,
		TemplateSuffix: dto.TemplateSuffix,
		Title:          dto.Title,
		UpdatedAt:      updatedAt,
	}
}
