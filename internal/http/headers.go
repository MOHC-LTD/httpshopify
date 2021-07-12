package http

import (
	"net/http"
	"strings"
)

// RequestHeader is a header on a request
type RequestHeader struct {
	// The name of the header
	Name string
	// The value of the header
	Value string
}

// RequestHeaders is a collection of headers to go on a request
type RequestHeaders []RequestHeader

// Includes returns whether the header is in the collection
func (headers RequestHeaders) Includes(name string) bool {
	for _, header := range headers {
		if strings.EqualFold(header.Name, name) {
			return true
		}
	}

	return false
}

// ResponseHeaders are headers on a response
type ResponseHeaders struct {
	header http.Header
}

// Get gets a header
func (h ResponseHeaders) Get(name string) string {
	return h.header.Get(name)
}

// OptionDefaultHeader holds configuration for a default header that will be placed on the requests made by the client.
type OptionDefaultHeader struct {
	name  string
	value string
}

func (option OptionDefaultHeader) configure(client *Client) error {
	client.defaultHeaders = append(client.defaultHeaders, RequestHeader{
		Name:  option.name,
		Value: option.value,
	})

	return nil
}

// WithDefaultHeader allows configuring the default headers of the requests made by the client.
func WithDefaultHeader(name string, value string) OptionDefaultHeader {
	return OptionDefaultHeader{
		name,
		value,
	}
}
