package httpshopify

import (
	"encoding/json"
	"fmt"

	"github.com/MOHC-LTD/httpshopify/v2/internal/http"
	"github.com/MOHC-LTD/shopify/v2"
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

func (r customerAddressRepository) List(id int64) (shopify.CustomerAddresses, error) {
	url := r.createURL(fmt.Sprintf("customers/%v/addresses.json", id))

	body, _, err := r.client.Get(url, nil)
	if err != nil {
		return shopify.CustomerAddresses{}, err
	}

	var response struct {
		Addresses CustomerAddressDTOs `json:"addresses"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return shopify.CustomerAddresses{}, err
	}

	return response.Addresses.ToShopify(), nil
}

func (r customerAddressRepository) Create(id int64, address shopify.CustomerAddress) (shopify.CustomerAddress, error) {
	url := r.createURL(fmt.Sprintf("customers/%v/addresses.json", id))

	addressDTO := BuildCustomerAddressDTO(address)

	// Wrapping Address object inside `address` property
	request := struct {
		CustomerAddressDTO `json:"address"`
	}{
		addressDTO,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return shopify.CustomerAddress{}, err
	}

	resBody, _, err := r.client.Post(url, body, nil)
	if err != nil {
		return shopify.CustomerAddress{}, err
	}

	// Wrapping Address object inside `address` property
	responseDTO := struct {
		CustomerAddressDTO `json:"customer_address"`
	}{}

	err = json.Unmarshal(resBody, &responseDTO)
	if err != nil {
		return shopify.CustomerAddress{}, err
	}

	return responseDTO.ToShopify(), nil
}

func (r customerAddressRepository) Update(id int64, address shopify.CustomerAddress) (shopify.CustomerAddress, error) {
	addressDTO := BuildCustomerAddressDTO(address)

	request := struct {
		CustomerAddressDTO `json:"address"`
	}{
		addressDTO,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return shopify.CustomerAddress{}, err
	}

	url := r.createURL(fmt.Sprintf("customers/%v/addresses/%v.json", id, address.ID))

	resBody, _, err := r.client.Put(url, body, nil)
	if err != nil {
		return shopify.CustomerAddress{}, err
	}

	responseDTO := struct {
		CustomerAddressDTO `json:"customer_address"`
	}{}

	err = json.Unmarshal(resBody, &responseDTO)
	if err != nil {
		return shopify.CustomerAddress{}, err
	}

	return responseDTO.ToShopify(), nil
}

func (r customerAddressRepository) SetDefault(id int64, addressID int64) (shopify.CustomerAddress, error) {
	url := r.createURL(fmt.Sprintf("customers/%v/addresses/%v/default.json", id, addressID))

	resBody, _, err := r.client.Put(url, nil, nil)
	if err != nil {
		return shopify.CustomerAddress{}, err
	}

	responseDTO := struct {
		CustomerAddressDTO `json:"customer_address"`
	}{}

	err = json.Unmarshal(resBody, &responseDTO)
	if err != nil {
		return shopify.CustomerAddress{}, err
	}

	return responseDTO.ToShopify(), nil
}

func (r customerAddressRepository) Delete(id int64, addressID int64) error {
	url := r.createURL(fmt.Sprintf("customers/%v/addresses/%v.json", id, addressID))

	_, _, err := r.client.Delete(url, nil)
	if err != nil {
		return err
	}

	return nil
}
