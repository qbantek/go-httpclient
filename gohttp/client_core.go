package gohttp

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
	"strings"

	"gopkg.in/square/go-jose.v2/json"
)

func (c *httpClient) getRequestBody(contentType string,
	body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) do(method string,
	url string,
	headers http.Header,
	body interface{}) (*http.Response, error) {

	requestHeaders := c.allHeaders(headers)
	requestBody, err := c.getRequestBody(requestHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	req, err := request(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header = requestHeaders

	client := http.Client{}
	return client.Do(req)
}

func request(method string,
	url string,
	body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *httpClient) allHeaders(headers http.Header) http.Header {
	result := make(http.Header)

	// Add common headers
	setHeaders(result, c.Headers)

	// Add custom headers
	setHeaders(result, headers)

	return result
}

func setHeaders(h http.Header, newHeaders http.Header) {
	for k, values := range newHeaders {
		if len(values) > 0 {
			h.Set(k, values[0])
		}
	}
}
