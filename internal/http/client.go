package http

import "net/http"

// Client is a HTTP client
type Client struct {
	*http.Client
	defaultHeaders RequestHeaders
}

// NewClient builds a new HTTP client
func NewClient(defaultHeaders RequestHeaders) Client {
	return Client{
		&http.Client{},
		defaultHeaders,
	}
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
