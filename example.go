package main

import (
	"fmt"
	"io"
	"os"

	"github.com/qbantek/go-httpclient/gohttp"
)

func main() {
	client := gohttp.New()
	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	io.Copy(os.Stdout, response.Body)
}
