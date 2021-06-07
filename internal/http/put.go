package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Put performs a PUT request
func (client Client) Put(url string, requestBody []byte, headers RequestHeaders) ([]byte, ResponseHeaders, error) {
	headers = client.AppendDefaultHeaders(headers)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	for _, header := range headers {
		req.Header.Set(header.Name, header.Value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	err = HandleStatus(resp.StatusCode)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	return body, ResponseHeaders{resp.Header}, nil
}
