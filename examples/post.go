package examples

import (
	"fmt"

	"github.com/qbantek/go-httpclient/gohttp"
)

// Repository represents a GitHub repository
type Repository struct {
	Name string `json:"name"`
}

// PostRepository creates a new respository
func PostRepository(repo *Repository) (*gohttp.Response, error) {
	response, err := httpClient.Post("https://api.github.com", nil, repo)
	fmt.Println(response)
	fmt.Println(err)

	return response, err
}
