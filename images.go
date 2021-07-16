package httpshopify

import (
	"time"

	"github.com/MOHC-LTD/shopify"
)

// ImageDTO represents a Shopify Image
type ImageDTO struct {
	CreatedAt time.Time `json:"created_at"`
	SRC       string    `json:"src"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Alt       string    `json:"alt"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ImageDTO) ToShopify() shopify.Image {
	return shopify.Image{
		CreatedAt: dto.CreatedAt,
		SRC:       dto.SRC,
		Width:     dto.Width,
		Height:    dto.Height,
		Alt:       dto.Alt,
	}
}
