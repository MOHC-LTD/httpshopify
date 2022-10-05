package httpshopify

import (
	"fmt"

	"github.com/MOHC-LTD/httpshopify/internal/http"
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

func (r customerAddressRepository) Delete(id int64, addressID int64) error {
	url := r.createURL(fmt.Sprintf("customers/%v/addresses/%v.json", id, addressID))

	_, _, err := r.client.Delete(url, nil)
	if err != nil {
		return err
	}

	return nil
}
