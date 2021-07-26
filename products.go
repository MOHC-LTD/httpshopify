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

func (repository productRepository) Create(product shopify.Product) (shopify.Product, error) {
	createDTO := ProductDTO{
		ID:          product.ID,
		CreatedAt:   product.CreatedAt,
		BodyHTML:    product.BodyHTML,
		ProductType: product.ProductType,
		Images:      BuildProductImageDTOs(product.Images),
		PublishedAt: product.PublishedAt,
		Status:      product.Status,
		Tags:        product.Tags,
		Title:       product.Title,
		UpdatedAt:   product.UpdatedAt,
		Variants:    BuildVariantDTOs(product.Variants),
		Vendor:      product.Vendor,
	}

	request := struct {
		Product ProductDTO `json:"product"`
	}{
		Product: createDTO,
	}

	body, err := json.Marshal(request)

	if err != nil {
		return shopify.Product{}, err
	}

	url := repository.createURL("products.json")

	respBody, _, err := repository.client.Post(url, body, nil)
	if err != nil {
		return shopify.Product{}, err
	}

	var response struct {
		Product ProductDTO `json:"product"`
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return shopify.Product{}, err
	}

	return response.Product.ToShopify(), nil
}

func (repository productRepository) Get(id int64) (shopify.Product, error) {
	url := repository.createURL(fmt.Sprintf("products/%v.json", id))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.Product{}, err
	}

	var response struct {
		Product ProductDTO `json:"product"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return shopify.Product{}, err
	}

	if response.Product.ID == 0 {
		return shopify.Product{}, shopify.NewErrProductNotFound(id)
	}

	return response.Product.ToShopify(), nil
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
	BodyHTML    string           `json:"body_html,omitempty"`
	CreatedAt   time.Time        `json:"created_at,omitempty"`
	Handle      string           `json:"handle,omitempty"`
	ID          int64            `json:"id,omitempty"`
	Images      ProductImageDTOs `json:"images,omitempty"`
	ProductType string           `json:"product_type,omitempty"`
	PublishedAt time.Time        `json:"published_at,omitempty"`
	Status      string           `json:"status,omitempty"`
	Tags        string           `json:"tags,omitempty"`
	Title       string           `json:"title,omitempty"`
	UpdatedAt   time.Time        `json:"updated_at,omitempty"`
	Variants    VariantDTOs      `json:"variants,omitempty"`
	Vendor      string           `json:"vendor,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ProductDTO) ToShopify() shopify.Product {
	return shopify.Product{
		BodyHTML:  dto.BodyHTML,
		CreatedAt: dto.CreatedAt,
		// Handle:      dto.Handle,
		ID:          dto.ID,
		Images:      dto.Images.ToShopify(),
		ProductType: dto.ProductType,
		PublishedAt: dto.PublishedAt,
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
