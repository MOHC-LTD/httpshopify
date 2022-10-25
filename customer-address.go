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

func (r customerAddressRepository) Delete(id int64, addressID int64) error {
	url := r.createURL(fmt.Sprintf("customers/%v/addresses/%v.json", id, addressID))

	_, _, err := r.client.Delete(url, nil)
	if err != nil {
		return err
	}

	return nil
}
