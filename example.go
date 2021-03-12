package main

import (
	"io"
	"net/http"
	"os"

	"github.com/qbantek/go-httpclient/gohttp"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	client = getGithubClient()
)

func main() {
	getUrls()
}

func getGithubClient() gohttp.HTTPClient {
	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")

	githubClient := gohttp.New()
	githubClient.SetHeaders(commonHeaders)

	return githubClient
}

func createUser(user User) {
	response, err := client.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		panic(err)
	}
}

func getUrls() {
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		panic(err)
	}
}
