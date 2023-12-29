package httpshopify

import (
	"encoding/json"
	"fmt"
	"time"

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
	url := repository.createURL(fmt.Sprintf("blogs/%v.json", id))

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return shopify.Blog{}, err
	}

	var resultDTO struct {
		Blog BlogDTO `json:"blog"`
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.Blog.ToShopify(), nil
}

func (repository blogRepository) GetAll() (shopify.Blogs, error) {
	url := repository.createURL("blogs.json")

	body, _, err := repository.client.Get(url, nil)
	if err != nil {
		return nil, err
	}

	var resultDTO struct {
		Blogs BlogDTOs `json:"blogs"`
	}

	json.Unmarshal(body, &resultDTO)

	return resultDTO.Blogs.ToShopify(), nil
}

// BlogDTOs represents a list of shopify blogs in HTTP requests and responses
type BlogDTOs []BlogDTO

// ToShopify converts the DTO to the Shopify equivalent
func (dtos BlogDTOs) ToShopify() shopify.Blogs {
	blogs := make(shopify.Blogs, 0, len(dtos))

	for _, dto := range dtos {
		blogs = append(blogs, dto.ToShopify())
	}

	return blogs
}

// BlogDTO represents a Shopify blog in HTTP requests and responses
type BlogDTO struct {
	Commentable string     `json:"commentable,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	Handle      string     `json:"handle,omitempty"`
	ID          int64      `json:"id,omitempty"`
	Tags        string     `json:"tags,omitempty"`
	Title       string     `json:"title,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// ToShopify converts the DTO to the Shopify equivalent
func (dto BlogDTO) ToShopify() shopify.Blog {
	var createdAt time.Time
	if dto.CreatedAt != nil {
		createdAt = *dto.CreatedAt
	}

	var updatedAt time.Time
	if dto.UpdatedAt != nil {
		updatedAt = *dto.UpdatedAt
	}

	return shopify.Blog{
		Commentable: dto.Commentable,
		CreatedAt:   createdAt,
		Handle:      dto.Handle,
		ID:          dto.ID,
		Tags:        dto.Tags,
		Title:       dto.Title,
		UpdatedAt:   updatedAt,
	}
}
