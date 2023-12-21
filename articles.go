package httpshopify

import (
	"github.com/MOHC-LTD/httpshopify/v2/internal/http"

	"github.com/MOHC-LTD/shopify/v2"
)

type articleRepository struct {
	client    http.Client
	createURL func(endpoint string) string
}

func newArticleRepository(client http.Client, createURL func(endpoint string) string) articleRepository {
	return articleRepository{
		client,
		createURL,
	}
}

func (repository articleRepository) Get(blogID, id int64) (shopify.Article, error) {
	panic("Get has not been implement yet")
}

func (repository articleRepository) GetAll(blogID int64) (shopify.Articles, error) {
	panic("GetAll has not been implement yet")
}
