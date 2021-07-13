package http

import (
	"net/http"
)

// Get performs a GET request
func (client Client) Get(url string, headers RequestHeaders) ([]byte, ResponseHeaders, error) {
	responseBody, responseHeaders, err := client.Do(http.MethodGet, url, headers, nil)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	return responseBody, responseHeaders, nil
}
