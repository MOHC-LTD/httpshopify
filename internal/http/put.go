package http

import (
	"bytes"
	"net/http"
)

// Put performs a PUT request
func (client Client) Put(url string, requestBody []byte, headers RequestHeaders) ([]byte, ResponseHeaders, error) {
	responseBody, responseHeaders, err := client.Do(http.MethodPut, url, headers, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	return responseBody, responseHeaders, nil
}
