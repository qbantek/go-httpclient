package examples

import (
	"github.com/qbantek/go-httpclient/gohttp"
)

// Repository represents a GitHub repository
type Repository struct {
	Name string `json:"name"`
}

// PostRepository creates a new respository
func PostRepository(repo *Repository) (*gohttp.Response, error) {
	response, err := httpClient.Post("https://api.github.com", nil, repo)
	return response, err
}
