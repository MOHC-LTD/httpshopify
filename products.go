package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/http"
	"github.com/MOHC-LTD/httpshopify/internal/slices"
	"github.com/MOHC-LTD/shopify"
)

type productRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newProductRepository(client http.Client, createURL func(endpoint string) string) productRepository {
	return productRepository{
		client,
		createURL,
	}
}

func (repository productRepository) List(query shopify.ProductQuery) (shopify.Products, error) {
	products := make(shopify.Products, 0)

	url := repository.createURL(fmt.Sprintf("products.json%v", parseProductQuery(query)))

	for {
		body, headers, err := repository.client.Get(url, nil)
		if err != nil {
			return nil, err
		}

		var resultDTO struct {
			Products ProductDTOs `json:"products"`
		}
		json.Unmarshal(body, &resultDTO)

		for _, dto := range resultDTO.Products {
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

// ProductDTOs is a collection of Product DTOs
type ProductDTOs []ProductDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos ProductDTOs) ToShopify() shopify.Products {
	products := make(shopify.Products, 0, len(dtos))

	for _, dto := range dtos {
		products = append(products, dto.ToShopify())
	}

	return products
}

// ProductDTO represents a Shopify product in HTTP requests and responses
type ProductDTO struct {
	ID          int64            `json:"id"`
	CreatedAt   time.Time        `json:"created_at"`
	BodyHTML    string           `json:"body_html"`
	ProductType string           `json:"product_type"`
	Images      ProductImageDTOs `json:"images"`
	PublishedAt time.Time        `json:"published_at"`
	Status      string           `json:"status"`
	Tags        string           `json:"tags"`
	Title       string           `json:"title"`
	UpdatedAt   time.Time        `json:"updated_at"`
	Variants    VariantDTOs      `json:"variants"`
	Vendor      string           `json:"vendor"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ProductDTO) ToShopify() shopify.Product {
	return shopify.Product{
		ID:          dto.ID,
		CreatedAt:   dto.CreatedAt,
		BodyHTML:    dto.BodyHTML,
		ProductType: dto.ProductType,
		PublishedAt: dto.PublishedAt,
		Images:      dto.Images.ToShopify(),
		Status:      dto.Status,
		Tags:        dto.Tags,
		Title:       dto.Title,
		UpdatedAt:   dto.UpdatedAt,
		Variants:    dto.Variants.ToShopify(),
		Vendor:      dto.Vendor,
	}
}

func parseProductQuery(query shopify.ProductQuery) string {
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
