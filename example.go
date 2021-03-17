package main

import (
	"io"
	"net/http"
	"os"

	"github.com/qbantek/go-httpclient/gohttp"
)

type user struct {
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

	client := gohttp.New()

	client.DisableTimeouts(true)
	// client.SetMaxIdleConns(4)
	// client.SetConnectionTimeout(1 * time.Second)
	// client.SetResponseTimeout(2 * time.Millisecond)
	// client.SetHeaders(commonHeaders)

	return client
}

func createUser(user user) {
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
