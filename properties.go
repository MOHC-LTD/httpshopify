package httpshopify

import "github.com/MOHC-LTD/shopify/v2"

// PropertyDTOs is a collection of property DTOs.
type PropertyDTOs []PropertyDTO

// ToShopify converts this DTO to the Shopify equivalent.
func (dtos PropertyDTOs) ToShopify() []shopify.Property {
	properties := make([]shopify.Property, 0, len(dtos))
	for _, dto := range dtos {
		properties = append(properties, dto.ToShopify())
	}

	return properties
}

// PropertyDTO holds custom information for an entity.
type PropertyDTO struct {
	// Name is the name of the property.
	Name string `json:"name"`
	// Value is the value of the property
	Value interface{} `json:"value"`
}

// ToShopify converts this DTO to the Shopify equivalent.
func (dto PropertyDTO) ToShopify() shopify.Property {
	return shopify.Property{
		Name:  dto.Name,
		Value: dto.Value,
	}
}

// BuildPropertyDTOs builds the DTO from the Shopify equivalent.
func BuildPropertyDTOs(properties []shopify.Property) PropertyDTOs {
	dtos := make(PropertyDTOs, 0, len(properties))
	for _, property := range properties {
		dtos = append(dtos, BuildPropertyDTO(property))
	}

	return dtos
}

// BuildPropertyDTO builds the DTO from the Shopify equivalent.
func BuildPropertyDTO(property shopify.Property) PropertyDTO {
	return PropertyDTO{
		Name:  property.Name,
		Value: property.Value,
	}
}
