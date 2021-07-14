package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/http"
	"github.com/MOHC-LTD/shopify"
)

type CollectionRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newCollectionRepository(client http.Client, createURL func(endpoint string) string) CollectionRepository {
	return CollectionRepository{
		client,
		createURL,
	}
}

func (repository CollectionRepository) Get(ID int64) (shopify.Collection, error) {

	url := repository.createURL(fmt.Sprintf("collections/%v.json%v", ID, parseCollectionQuery()))

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

func (repository CollectionRepository) Products(ID int64) (shopify.Products, error) {

	url := repository.createURL(fmt.Sprintf("collections/%v/products.json%v", ID, parseCollectionQuery()))

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

// CollectionDTOs is a collection of Product DTOs
type CollectionDTOs []CollectionDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos CollectionDTOs) ToShopify() shopify.Collections {
	collections := make(shopify.Collections, 0, len(dtos))

	for _, dto := range dtos {
		collections = append(collections, dto.ToShopify())
	}

	return collections
}

// ProductDTO represents a Shopify product in HTTP requests and responses
type CollectionDTO struct {
	ID             int64     `json:"id"`
	Title          string    `json:"title"`
	BodyHTML       string    `json:"body_html"`
	Handle         string    `json:"handle"`
	PublishedAt    time.Time `json:"published_at"`
	PublishedScope string    `json:"published_scope"`
	SortOrder      string    `json:"sort_order"`
	Image          ImageDTO  `json:"image"`
	TemplateSuffix string    `json:"template_suffix"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto CollectionDTO) ToShopify() shopify.Collection {
	return shopify.Collection{
		ID:             dto.ID,
		Title:          dto.Title,
		BodyHTML:       dto.BodyHTML,
		Handle:         dto.Handle,
		PublishedAt:    dto.PublishedAt,
		PublishedScope: dto.PublishedScope,
		Image:          dto.Image.ToShopify(),
		SortOrder:      dto.SortOrder,
		TemplateSuffix: dto.TemplateSuffix,
		UpdatedAt:      dto.UpdatedAt,
	}
}

func parseCollectionQuery() string {
	queryStrings := make([]string, 0)

	if len(queryStrings) == 0 {
		return ""
	}

	queryString := "?"

	for i, str := range queryStrings {
		if i != 0 {
			queryString = queryString + "&"
		}

		queryString = queryString + str
	}

	return queryString
}
