package http

import (
	"io/ioutil"
	"net/http"
)

// Get performs a GET request
func (client Client) Get(url string, headers RequestHeaders) ([]byte, ResponseHeaders, error) {
	headers = client.AppendDefaultHeaders(headers)

	req, err := http.NewRequest(http.MethodGet, url, nil)
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
