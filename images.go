package httpshopify

import (
	"time"

	"github.com/MOHC-LTD/shopify"
)

// ImageDTO represents a Shopify Image in HTTP requests and responses
type ImageDTO struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	SRC       string     `json:"src,omitempty"`
	Width     int        `json:"width,omitempty"`
	Height    int        `json:"height,omitempty"`
	Alt       string     `json:"alt,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ImageDTO) ToShopify() shopify.Image {
	var createdAt time.Time
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	return shopify.Image{
		CreatedAt: createdAt,
		SRC:       dto.SRC,
		Width:     dto.Width,
		Height:    dto.Height,
		Alt:       dto.Alt,
	}
}
