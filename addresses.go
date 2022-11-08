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

// CustomerAddressDTO represents a Shopify customer address in HTTP requests and responses
type CustomerAddressDTO struct {
	Address1     string  `json:"address1,omitempty"`
	Address2     *string `json:"address2"`
	City         string  `json:"city,omitempty"`
	Company      string  `json:"company,omitempty"`
	Country      string  `json:"country,omitempty"`
	CountryCode  string  `json:"country_code,omitempty"`
	CountryName  string  `json:"country_name,omitempty"`
	CustomerID   uint64  `json:"customer_id,omitempty"`
	Default      bool    `json:"default,omitempty"`
	FirstName    string  `json:"first_name,omitempty"`
	ID           uint64  `json:"id,omitempty"`
	LastName     string  `json:"last_name,omitempty"`
	Name         string  `json:"name,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	Province     string  `json:"province,omitempty"`
	ProvinceCode string  `json:"province_code,omitempty"`
	Zip          *string `json:"zip"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto CustomerAddressDTO) ToShopify() shopify.CustomerAddress {
	var address2 string
	var zip string

	if dto.Address2 != nil {
		address2 = *dto.Address2
	}

	if dto.Zip != nil {
		zip = *dto.Zip
	}

	return shopify.CustomerAddress{
		Address1:     dto.Address1,
		Address2:     address2,
		City:         dto.City,
		Company:      dto.Company,
		Country:      dto.Country,
		CountryCode:  dto.CountryCode,
		CountryName:  dto.CountryName,
		CustomerID:   dto.CustomerID,
		Default:      dto.Default,
		FirstName:    dto.FirstName,
		ID:           dto.ID,
		LastName:     dto.LastName,
		Name:         dto.Name,
		Phone:        dto.Phone,
		Province:     dto.Province,
		ProvinceCode: dto.ProvinceCode,
		Zip:          zip,
	}
}

// BuildCustomerAddressDTO converts a Shopify customer address to its DTO equivalent
func BuildCustomerAddressDTO(address shopify.CustomerAddress) CustomerAddressDTO {
	var address2 *string
	var zip *string

	if address.Address2 != "" {
		address2 = &address.Address2
	}

	if address.Zip != "" {
		zip = &address.Zip
	}

	return CustomerAddressDTO{
		Address1:     address.Address1,
		Address2:     address2,
		City:         address.City,
		Company:      address.Company,
		Country:      address.Country,
		CountryCode:  address.CountryCode,
		CountryName:  address.CountryName,
		CustomerID:   address.CustomerID,
		Default:      address.Default,
		FirstName:    address.FirstName,
		ID:           address.ID,
		LastName:     address.LastName,
		Name:         address.Name,
		Phone:        address.Phone,
		Province:     address.Province,
		ProvinceCode: address.ProvinceCode,
		Zip:          zip,
	}
}
