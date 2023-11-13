package httpshopify_test

import (
	"testing"

	"github.com/MOHC-LTD/httpshopify/v2"
	"github.com/MOHC-LTD/shopify/v2"
)

func Test_ShopImplementsShopify(t *testing.T) {
	var _ shopify.Shop = new(httpshopify.Shop)
}
