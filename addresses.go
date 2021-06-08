package httpshopify

import "github.com/MOHC-LTD/shopify"

type addressDTO struct {
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Company      string `json:"company"`
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Province     string `json:"province"`
	ProvinceCode string `json:"province_code"`
	Zip          string `json:"zip"`
}

func (dto addressDTO) toDomain() shopify.Address {
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
