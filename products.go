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
		CreatedAt:   &product.CreatedAt,
		Handle:      product.Handle,
		BodyHTML:    product.BodyHTML,
		ProductType: product.ProductType,
		Images:      BuildProductImageDTOs(product.Images),
		PublishedAt: &product.PublishedAt,
		Status:      product.Status,
		Tags:        string(product.Tags),
		Title:       product.Title,
		UpdatedAt:   &product.UpdatedAt,
		Variants:    BuildVariantDTOs(product.Variants),
		Vendor:      product.Vendor,
		Options:     BuildOptionsDTOs(product.Options),
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

func (repository productRepository) Update(product shopify.Product) (shopify.Product, error) {
	updateDTO := ProductDTO{
		ID:          product.ID,
		CreatedAt:   &product.CreatedAt,
		Handle:      product.Handle,
		BodyHTML:    product.BodyHTML,
		ProductType: product.ProductType,
		Images:      BuildProductImageDTOs(product.Images),
		PublishedAt: &product.PublishedAt,
		Status:      product.Status,
		Tags:        string(product.Tags),
		Title:       product.Title,
		UpdatedAt:   &product.UpdatedAt,
		Variants:    BuildVariantDTOs(product.Variants),
		Vendor:      product.Vendor,
		Options:     BuildOptionsDTOs(product.Options),
	}

	request := struct {
		Product ProductDTO `json:"product"`
	}{
		Product: updateDTO,
	}

	body, err := json.Marshal(request)

	if err != nil {
		return shopify.Product{}, err
	}

	url := repository.createURL(fmt.Sprintf("products/%d.json", updateDTO.ID))

	respBody, _, err := repository.client.Put(url, body, nil)
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
	BodyHTML    string            `json:"body_html,omitempty"`
	CreatedAt   *time.Time        `json:"created_at,omitempty"`
	Handle      string            `json:"handle,omitempty"`
	ID          int64             `json:"id,omitempty"`
	Images      ProductImageDTOs  `json:"images,omitempty"`
	ProductType string            `json:"product_type,omitempty"`
	PublishedAt *time.Time        `json:"published_at,omitempty"`
	Status      string            `json:"status,omitempty"`
	Tags        string            `json:"tags,omitempty"`
	Title       string            `json:"title,omitempty"`
	UpdatedAt   *time.Time        `json:"updated_at,omitempty"`
	Variants    VariantDTOs       `json:"variants,omitempty"`
	Vendor      string            `json:"vendor,omitempty"`
	Options     ProductOptionsDTO `json:"options,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ProductDTO) ToShopify() shopify.Product {
	var createdAt time.Time
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	var publishedAt time.Time
	if dto.PublishedAt != nil {
		publishedAt = *dto.PublishedAt
	}

	var updatedAt time.Time
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.Product{
		BodyHTML:    dto.BodyHTML,
		CreatedAt:   createdAt,
		Handle:      dto.Handle,
		ID:          dto.ID,
		Images:      dto.Images.ToShopify(),
		ProductType: dto.ProductType,
		PublishedAt: publishedAt,
		Status:      dto.Status,
		Tags:        shopify.Tags(dto.Tags),
		Title:       dto.Title,
		UpdatedAt:   updatedAt,
		Variants:    dto.Variants.ToShopify(),
		Vendor:      dto.Vendor,
		Options:     dto.Options.ToShopify(),
	}
}

// ProductOptionsDTO represents Shopify product options in HTTP requests and responses
type ProductOptionsDTO []ProductOptionDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos ProductOptionsDTO) ToShopify() shopify.ProductOptions {
	options := make(shopify.ProductOptions, 0, len(dtos))

	for _, dto := range dtos {
		options = append(options, dto.ToShopify())
	}

	return options
}

// ProductOptionDTO represents a Shopify product option in HTTP requests and responses
type ProductOptionDTO struct {
	ID       int64    `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Position int      `json:"position,omitempty"`
	Values   []string `json:"values,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ProductOptionDTO) ToShopify() shopify.ProductOption {
	return shopify.ProductOption{
		ID:       dto.ID,
		Name:     dto.Name,
		Position: dto.Position,
		Values:   dto.Values,
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
