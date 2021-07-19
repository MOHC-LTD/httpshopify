package http

import (
	"net/http"
)

// Delete performs a DELETE request
func (client Client) Delete(url string, headers RequestHeaders) ([]byte, ResponseHeaders, error) {
	responseBody, responseHeaders, err := client.Do(http.MethodDelete, url, headers, nil)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	return responseBody, responseHeaders, nil
}
