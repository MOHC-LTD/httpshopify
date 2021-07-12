package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/http"
	"github.com/MOHC-LTD/httpshopify/internal/slices"
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

func (repository CollectionRepository) List(query shopify.CollectionQuery) (shopify.Products, error) {
	products := make(shopify.Products, 0)

	url := repository.createURL(fmt.Sprintf("collections.json%v", parseProductQuery(query)))

	for {
		body, headers, err := repository.client.Get(url, nil)
		if err != nil {
			return nil, err
		}

		var resultDTO struct {
			Collections collectionDTO `json:"collections"`
		}
		json.Unmarshal(body, &resultDTO)

		for _, dto := range resultDTO.Collections {
			products = append(products, dto.ToShopify())
		}

		links := ParseLinkHeader(headers.Get("Link"))

		if !links.HasNext() {
			break
		}

		url = links.Next
	}

	return products, nil
}

// CollectionDTOs is a collection of Product DTOs
type CollectionDTOs []collectionDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos CollectionDTOs) ToShopify() shopify.Collections {
	collections := make(shopify.Collections, 0, len(dtos))

	for _, dto := range dtos {
		collections = append(collections, dto.ToShopify())
	}

	return collections
}

// ProductDTO represents a Shopify product in HTTP requests and responses
type collectionDTO struct {
	ID             int64     `json:"id"`
	Title          string    `json:"title"`
	BodyHTML       string    `json:"body_html"`
	Handle         string    `json:"handle"`
	PublishedAt    time.Time `json:"published_at"`
	PublishedScope string    `json:"published_scope"`
	SortOrder      string    `json:"sort_order"`
	TemplateSuffix string    `json:"template_suffix"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto collectionDTO) ToShopify() shopify.Collection {
	return shopify.Collection{
		ID:             dto.ID,
		Title:          dto.Title,
		BodyHTML:       dto.BodyHTML,
		Handle:         dto.Handle,
		PublishedAt:    dto.PublishedAt,
		PublishedScope: dto.PublishedScope,
		SortOrder:      dto.SortOrder,
		TemplateSuffix: dto.TemplateSuffix,
		UpdatedAt:      dto.UpdatedAt,
	}
}

func parseCollectionQuery(query shopify.CollectionQuery) string {
	queryStrings := make([]string, 0)

	if query.IDs != nil {
		queryStrings = append(queryStrings, fmt.Sprintf("ids=%v", slices.JoinInt64(query.IDs, ",")))
	}

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
