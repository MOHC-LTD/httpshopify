package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"
	"github.com/MOHC-LTD/shopify/v2"
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
		return nil, err
	}

	var resultDTO struct {
		Collection CollectionDTO `json:"collection"`
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.Collection.ToShopify(), nil
}

func (repository collectionRepository) GetSmartCollectionsList() (shopify.Collections, error) {
	url := repository.createURL("smart_collections.json")

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var resultDTO struct {
		Collections CollectionDTOs `json:"smart_collections"`
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.Collections.ToShopify(), nil
}

func (repository collectionRepository) GetCustomCollectionsList() (shopify.Collections, error) {
	url := repository.createURL("custom_collections.json")

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var resultDTO struct {
		Collections CollectionDTOs `json:"custom_collections"`
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.Collections.ToShopify(), nil
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
	CollectionType string     `json:"collection_type,omitempty"`
	Handle         string     `json:"handle,omitempty"`
	Image          ImageDTO   `json:"image,omitempty"`
	ID             int64      `json:"id,omitempty"`
	PublishedAt    *time.Time `json:"published_at,omitempty"`
	PublishedScope string     `json:"published_scope,omitempty"`
	Rules          RuleDTOs   `json:"rules,omitempty"`
	Disjunctive    bool       `json:"disjunctive,omitempty"`
	SortOrder      string     `json:"sort_order,omitempty"`
	TemplateSuffix string     `json:"template_suffix,omitempty"`
	ProductsCount  int        `json:"products_count,omitempty"`
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

	switch dto.CollectionType {
	case shopify.CollectionTypeSmart:
		return shopify.NewSmartCollection(
			dto.BodyHTML,
			dto.CollectionType,
			dto.Handle,
			dto.ID,
			dto.Image.ToShopify(),
			dto.ProductsCount,
			publishedAt,
			dto.PublishedScope,
			dto.Rules.ToShopify(),
			dto.Disjunctive,
			dto.SortOrder,
			dto.TemplateSuffix,
			dto.Title,
			updatedAt,
		)
	default:
		return shopify.NewCustomCollection(
			dto.BodyHTML,
			dto.CollectionType,
			dto.Handle,
			dto.ID,
			dto.Image.ToShopify(),
			dto.ProductsCount,
			publishedAt,
			dto.PublishedScope,
			dto.SortOrder,
			dto.TemplateSuffix,
			dto.Title,
			updatedAt,
		)
	}
}
