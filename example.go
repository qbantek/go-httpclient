package main

import (
	"fmt"

	"github.com/qbantek/go-httpclient/gohttp"
)

func main() {
	client := gohttp.New()
	client.Get()

	fmt.Println(client)
}
