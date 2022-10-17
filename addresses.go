package httpshopify

import "github.com/MOHC-LTD/shopify"

// AddressDTO represents a Shopify Address in HTTP requests and responses
type AddressDTO struct {
	Address1     string `json:"address1,omitempty"`
	Address2     string `json:"address2,omitempty"`
	City         string `json:"city,omitempty"`
	Company      string `json:"company,omitempty"`
	Country      string `json:"country,omitempty"`
	CountryCode  string `json:"country_code,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	ID           uint64 `json:"id,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Province     string `json:"province,omitempty"`
	ProvinceCode string `json:"province_code,omitempty"`
	Zip          string `json:"zip,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto AddressDTO) ToShopify() shopify.Address {
	return shopify.Address{
		Address1:     dto.Address1,
		Address2:     dto.Address2,
		City:         dto.City,
		Company:      dto.Company,
		Country:      dto.Country,
		CountryCode:  dto.CountryCode,
		FirstName:    dto.FirstName,
		ID:           dto.ID,
		LastName:     dto.LastName,
		Name:         dto.Name,
		Phone:        dto.Phone,
		Province:     dto.Province,
		ProvinceCode: dto.ProvinceCode,
		Zip:          dto.Zip,
	}
}

// BuildAddressDTO converts a Shopify address to its DTO equivalent
func BuildAddressDTO(address shopify.Address) AddressDTO {
	return AddressDTO(address)
}
