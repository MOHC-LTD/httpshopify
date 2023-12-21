package httpshopify

import (
	"github.com/MOHC-LTD/httpshopify/v2/internal/http"

	"github.com/MOHC-LTD/shopify/v2"
)

type blogRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newBlogRepository(client http.Client, createURL func(endpoint string) string) blogRepository {
	return blogRepository{
		client,
		createURL,
	}
}

func (repository blogRepository) Get(id int64) (shopify.Blog, error) {
	panic("Get has not been implement yet")
}

func (repository blogRepository) GetAll() (shopify.Blogs, error) {
	panic("GetAll has not been implement yet")
}
