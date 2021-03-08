package gohttp

import (
	"net/http"
)

func (c *httpClient) do(method string,
	url string,
	headers http.Header,
	body interface{}) (*http.Response, error) {

	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}
