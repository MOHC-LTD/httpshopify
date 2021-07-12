package http

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/time/rate"
)

// Client is a HTTP client
type Client struct {
	*http.Client
	defaultHeaders RequestHeaders
	limiter        *rate.Limiter
}

// NewClient builds a new HTTP client
func NewClient(defaultHeaders RequestHeaders, options ...Option) Client {
	client := Client{
		Client:         &http.Client{},
		defaultHeaders: defaultHeaders,
	}

	for _, option := range options {
		option.configure(&client)
	}

	return client
}

// AppendDefaultHeaders appends the default headers to the passed ones.
func (client Client) AppendDefaultHeaders(headers RequestHeaders) RequestHeaders {
	for _, header := range client.defaultHeaders {
		if !headers.Includes(header.Name) {
			headers = append(headers, header)
		}
	}

	return headers
}

// Do does a request
func (client Client) Do(method string, url string, headers RequestHeaders, body io.Reader) ([]byte, ResponseHeaders, error) {
	if client.limiter != nil {
		ctx := context.Background()
		client.limiter.Wait(ctx)
	}

	headers = client.AppendDefaultHeaders(headers)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	for _, header := range headers {
		req.Header.Set(header.Name, header.Value)
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	err = HandleStatus(resp.StatusCode)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ResponseHeaders{}, err
	}

	return responseBody, ResponseHeaders{resp.Header}, nil
}
