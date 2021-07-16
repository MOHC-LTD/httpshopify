package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MOHC-LTD/httpshopify/internal/http"

	"github.com/MOHC-LTD/shopify"
)

type productImagesRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newProductImagesRepository(client http.Client, createURL func(endpoint string) string) productRepository {
	return productRepository{
		client,
		createURL,
	}
}

// ProductImagesDTOs is a collection of productImage DTOs
type ProductImageDTOs []ProductImageDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos ProductImageDTOs) ToShopify() shopify.ProductImages {
	productImages := make(shopify.ProductImages, 0, len(dtos))

	for _, dto := range dtos {
		productImages = append(productImages, dto.ToShopify())
	}

	return productImages
}

func (repository productImagesRepository) List(productID int64) (shopify.ProductImages, error) {
	productImages := make(shopify.ProductImages, 0)

	url := repository.createURL(fmt.Sprintf("products/%v/images.json%v", productID, parseProductImagesQuery(query)))

	for {
		body, headers, err := repository.client.Get(url, nil)
		if err != nil {
			return nil, err
		}

		var resultDTO struct {
			Images ProductImageDTOs `json:"images"`
		}
		json.Unmarshal(body, &resultDTO)

		for _, dto := range resultDTO.Images {
			productImages = append(productImages, dto.ToShopify())
		}

		links := ParseLinkHeader(headers.Get("Link"))

		if !links.HasNext() {
			break
		}

		url = links.Next
	}

	return productImages, nil
}

// ProductImagesDTO represents a Shopify product images in HTTP requests and responses
type ProductImageDTO struct {
	ImageDTO
	ID         int64     `json:"id"`
	Position   int       `json:"position"`
	ProductID  int64     `json:"product_id"`
	VariantIDs []int64   `json:"variant_ids"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ProductImageDTO) ToShopify() shopify.ProductImage {
	return shopify.ProductImage{
		ID: dto.ID,
		Image: shopify.Image{
			CreatedAt: dto.CreatedAt,
			SRC:       dto.SRC,
			Width:     dto.Width,
			Height:    dto.Height,
			Alt:       dto.Alt,
		},
		Position:   dto.Position,
		ProductID:  dto.ProductID,
		VariantIDs: dto.VariantIDs,
		UpdatedAt:  dto.UpdatedAt,
	}
}

func parseProductImagesQuery(query shopify.ProductImageQuery) string {
	queryStrings := make([]string, 0)

	if query.SinceID != 0 {
		queryStrings = append(queryStrings, fmt.Sprintf("since_ids=%v", query.SinceID))
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
