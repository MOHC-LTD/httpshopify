package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"
)

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

// ProductImagesDTO represents a Shopify product images in HTTP requests and responses
type ProductImageDTO struct {
	ID int64 `json:"id"`
	ImageDTO
	Position   int       `json:"position"`
	ProductID  int64     `json:"product_id"`
	VariantIDs []int64   `json:"variant_ids"`
	updatedAt  time.Time `json:"updated_at"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ProductImageDTO) ToShopify() shopify.ProductImage {
	return Shopify.ProductImage{
		ID:         dto.ID,
		CreatedAt:  dto.CreatedAt,
		SRC:        dto.SRC,
		Width:      dto.Width,
		Height:     dto.Height,
		Alt:        dto.Alt,
		Position:   dto.Position,
		ProductID:  dto.ProductID,
		VariantIDs: dto.VariantIDs,
		UpdatedAt:  dto.updatedAt,
	}
}

func (repository shopify.ProductImagesRepository) List(productID int64, query shopify.ProductImageQuery) (shopify.ProductImages, error) {
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

func parseProductImagesQuery(query shopify.ProductImagesQuery) string {
	queryStrings := make([]string, 0)

	if query.ID != 0 {
		queryStrings = append(queryStrings, fmt.Sprintf("since_ids=%v", slices.JoinInt64(query.IDs, ",")))
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
