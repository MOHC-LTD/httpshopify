package httpshopify

import (
	"strings"
)

// Pagination represents the pagination details returned by shopify
// for an endpoint that supports pagination.
//
// See here for more: https://shopify.dev/tutorials/make-paginated-requests-to-rest-admin-api
type Pagination struct {
	// The next page url
	Next string
	// The prev page url
	Prev string
}

// HasNext returns whether or not there is a next page
func (pagination Pagination) HasNext() bool {
	return pagination.Next != ""
}

// HasPrev returns whether or not there is a previous page
func (pagination Pagination) HasPrev() bool {
	return pagination.Prev != ""
}

// ParseLinkHeader parses the link header string returned by pagination
// enabled shopify endpoints
func ParseLinkHeader(linkHeader string) Pagination {
	if linkHeader == "" {
		return Pagination{}
	}

	var pagination Pagination

	parts := strings.Split(linkHeader, ",")

	for _, part := range parts {
		subParts := strings.Split(part, ";")

		url := TrimBetween(subParts[0], "<", ">")

		isNext := strings.Contains(subParts[1], "next")

		if isNext {
			pagination.Next = url
		} else {
			pagination.Prev = url
		}
	}

	return pagination
}

// TrimBetween returns the substring between the start and end strings
func TrimBetween(str string, start string, end string) string {
	startIndex := strings.Index(str, start)
	if startIndex == -1 {
		return ""
	}
	result := str[startIndex+len(start):]

	endIndex := strings.Index(result, end)
	if endIndex == -1 {
		return ""
	}

	return result[:endIndex]
}
