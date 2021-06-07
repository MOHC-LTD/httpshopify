package httpshopify_test

import (
	"testing"

	"github.com/MOHC-LTD/httpshopify"
	"github.com/MOHC-LTD/shopify"
)

func Test_ShopImplementsShopify(t *testing.T) {
	var _ shopify.Shop = new(httpshopify.Shop)
}
