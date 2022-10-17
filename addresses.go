package httpshopify

import "github.com/MOHC-LTD/shopify"

// AddressDTO represents a Shopify Address in HTTP requests and responses
type AddressDTO struct {
	Address1     string  `json:"address1,omitempty"`
	Address2     string  `json:"address2,omitempty"`
	City         string  `json:"city,omitempty"`
	Company      string  `json:"company,omitempty"`
	Country      string  `json:"country,omitempty"`
	CountryCode  string  `json:"country_code,omitempty"`
	FirstName    string  `json:"first_name,omitempty"`
	LastName     string  `json:"last_name,omitempty"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Name         string  `json:"name,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	Province     string  `json:"province,omitempty"`
	ProvinceCode string  `json:"province_code,omitempty"`
	Zip          string  `json:"zip,omitempty"`
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
		LastName:     dto.LastName,
		Latitude:     dto.Latitude,
		Longitude:    dto.Longitude,
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

type CustomerAddressDTO struct {
	Address1     string `json:"address1,omitempty"`
	Address2     string `json:"address2,omitempty"`
	City         string `json:"city,omitempty"`
	Company      string `json:"company,omitempty"`
	Country      string `json:"country,omitempty"`
	CountryCode  string `json:"country_code,omitempty"`
	CountryName  string `json:"country_name,omitempty"`
	CustomerID   uint64 `json:"customer_id,omitempty"`
	Default      bool   `json:"default,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Province     string `json:"province,omitempty"`
	ProvinceCode string `json:"province_code,omitempty"`
	Zip          string `json:"zip,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto CustomerAddressDTO) ToShopify() shopify.CustomerAddress {
	return shopify.CustomerAddress{
		Address1:     dto.Address1,
		Address2:     dto.Address2,
		City:         dto.City,
		Company:      dto.Company,
		Country:      dto.Country,
		CountryCode:  dto.CountryCode,
		CountryName:  dto.CountryName,
		CustomerID:   dto.CustomerID,
		Default:      dto.Default,
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		Name:         dto.Name,
		Phone:        dto.Phone,
		Province:     dto.Province,
		ProvinceCode: dto.ProvinceCode,
		Zip:          dto.Zip,
	}
}

// BuildCustomerAddressDTO converts a Shopify customer address to its DTO equivalent
func BuildCustomerAddressDTO(address shopify.CustomerAddress) CustomerAddressDTO {
	return CustomerAddressDTO(address)
}
