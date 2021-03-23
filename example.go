package main

import (
	"fmt"
	"time"

	"github.com/qbantek/go-httpclient/gohttp"
)

type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	client = githubClient()
)

func main() {
	getUrls()
	// createUser(user{})
}

func githubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetMaxIdleConns(4).
		SetConnectionTimeout(1 * time.Second).
		SetResponseTimeout(4 * time.Second).
		Build()

	return client
}

func createUser(user user) {
	response, err := client.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.String())
}

func getUrls() {
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.String())
}
