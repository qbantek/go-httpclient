package examples

import (
	"time"

	"github.com/qbantek/go-httpclient/gohttp"
)

var (
	httpClient = getHTTPClient()
)

func getHTTPClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(1 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()

	return client
}
