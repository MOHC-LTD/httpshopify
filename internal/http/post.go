package http

import (
	"bytes"
	"net/http"
)

// Post performs a POST request
func (client Client) Post(url string, requestBody []byte, headers RequestHeaders) ([]byte, ResponseHeaders, error) {
	responseBody, responseHeaders, err := client.Do(http.MethodPost, url, headers, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	return responseBody, responseHeaders, nil
}
