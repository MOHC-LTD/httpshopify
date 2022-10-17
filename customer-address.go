package httpshopify

import (
	"encoding/json"
	"fmt"

	"github.com/MOHC-LTD/httpshopify/internal/http"
	"github.com/MOHC-LTD/shopify"
)

type customerAddressRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newCustomerAddressRepository(client http.Client, createURL func(endpoint string) string) customerAddressRepository {
	return customerAddressRepository{
		client,
		createURL,
	}
}

func (r customerAddressRepository) Create(id int64, address shopify.Address) (shopify.Address, error) {
	url := r.createURL(fmt.Sprintf("customers/%v/addresses.json", id))

	addressDTO := AddressDTO{
		Address1:     address.Address1,
		Address2:     address.Address2,
		City:         address.City,
		Company:      address.Company,
		Country:      address.Country,
		CountryCode:  address.CountryCode,
		FirstName:    address.FirstName,
		ID:           address.ID,
		LastName:     address.LastName,
		Name:         address.Name,
		Phone:        address.Phone,
		Province:     address.Province,
		ProvinceCode: address.ProvinceCode,
		Zip:          address.Zip,
	}

	// Wrapping Address object inside `address` property
	request := struct {
		Address AddressDTO `json:"address"`
	}{
		Address: addressDTO,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return shopify.Address{}, err
	}

	resBody, _, err := r.client.Post(url, body, nil)
	if err != nil {
		return shopify.Address{}, err
	}

	// Wrapping Address object inside `address` property
	responseDTO := struct {
		CustomerAddress AddressDTO `json:"customer_address"`
	}{}

	err = json.Unmarshal(resBody, &responseDTO)
	if err != nil {
		return shopify.Address{}, err
	}

	return responseDTO.CustomerAddress.ToShopify(), nil
}

func (r customerAddressRepository) Delete(id int64, addressID int64) error {
	url := r.createURL(fmt.Sprintf("customers/%v/addresses/%v.json", id, addressID))

	_, _, err := r.client.Delete(url, nil)
	if err != nil {
		return err
	}

	return nil
}
