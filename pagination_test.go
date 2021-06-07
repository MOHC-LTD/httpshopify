package httpshopify_test

import (
	"reflect"
	"testing"

	"github.com/MOHC-LTD/httpshopify"
	"github.com/MOHC-LTD/httpshopify/internal/assertions"
)

func TestParseLinkHeader(t *testing.T) {
	// Example lifted from https://shopify.dev/tutorials/make-paginated-requests-to-rest-admin-api
	linkHeader := "<https://{shop}.myshopify.com/admin/api/{version}/products.json?page_info={page_info}&limit={limit}>; rel={next}, <https://{shop}.myshopify.com/admin/api/{version}/products.json?page_info={page_info}&limit={limit}>; rel={previous}"

	expected := httpshopify.Pagination{
		Next: "https://{shop}.myshopify.com/admin/api/{version}/products.json?page_info={page_info}&limit={limit}",
		Prev: "https://{shop}.myshopify.com/admin/api/{version}/products.json?page_info={page_info}&limit={limit}",
	}

	actual := httpshopify.ParseLinkHeader(linkHeader)

	if !reflect.DeepEqual(expected, actual) {
		assertions.ValueAssertionFailure(t, expected, actual)
	}
}
