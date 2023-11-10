package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

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
	url := repository.createURL(fmt.Sprintf("blogs/%v/articles/%v.json", blogID, id))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.Article{}, err
	}

	var resultDTO struct {
		Article ArticleDTO `json:"Article"`
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.Article.ToShopify(), nil
}

func (repository articleRepository) GetAll(blogID int64) (shopify.Articles, error) {
	url := repository.createURL(fmt.Sprintf("blogs/%v/articles.json", blogID))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var resultDTO struct {
		Articles ArticleDTOs `json:"Articles"`
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.Articles.ToShopify(), nil
}

// ArticleDTOs represents a list of shopify Articles in HTTP requests and responses
type ArticleDTOs []ArticleDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos ArticleDTOs) ToShopify() shopify.Articles {
	Articles := make(shopify.Articles, 0, len(dtos))

	for _, dto := range dtos {
		Articles = append(Articles, dto.ToShopify())
	}

	return Articles
}

// ArticleDTO represents a Shopify Article in HTTP requests and responses
type ArticleDTO struct {
	Author      string     `json:"author,omitempty"`
	BlogID      string     `json:"blog_id,omitempty"`
	BodyHTML    string     `json:"body_html,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	ID          int64      `json:"id,omitempty"`
	Handle      string     `json:"handle,omitempty"`
	Image       ImageDTO   `json:"image,omitempty"`
	Published   bool       `json:"published,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	SummaryHTML string     `json:"summary_html,omitempty"`
	Tags        string     `json:"tags,omitempty"`
	Title       string     `json:"title,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	UserID      int64      `json:"user_id,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto ArticleDTO) ToShopify() shopify.Article {
	var createdAt time.Time
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	var publishedAt time.Time
	if dto.PublishedAt != nil {
		publishedAt = *dto.PublishedAt
	}

	var updatedAt time.Time
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.Article{
		Author:      dto.Author,
		BlogID:      dto.BlogID,
		BodyHTML:    dto.BodyHTML,
		CreatedAt:   createdAt,
		ID:          dto.ID,
		Handle:      dto.Handle,
		Image:       dto.Image.ToShopify(),
		Published:   dto.Published,
		PublishedAt: publishedAt,
		SummaryHTML: dto.SummaryHTML,
		Tags:        dto.Tags,
		Title:       dto.Title,
		UpdatedAt:   updatedAt,
		UserID:      dto.UserID,
	}
}
