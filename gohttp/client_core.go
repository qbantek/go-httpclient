package gohttp

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"gopkg.in/square/go-jose.v2/json"
)

const (
	defaultConnectionTimeout = 2 * time.Second
	defaultResponseTimeout   = 1 * time.Second
)

func (c *httpClient) do(method string,
	url string,
	headers http.Header,
	body interface{}) (*Response, error) {

	reqHeaders := c.requestHeaders(headers)
	reqBody, err := c.requestBody(reqHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header = reqHeaders

	client := c.getHTTPClient()

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	finalResponse := Response{
		statusCode: response.StatusCode,
		status:     response.Status,
		headers:    response.Header,
		body:       responseBody,
	}
	return &finalResponse, nil
}

func (c *httpClient) getHTTPClient() *http.Client {
	c.once.Do(func() {
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConns > 0 {
		return c.builder.maxIdleConns
	}

	return http.DefaultMaxIdleConnsPerHost
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.disableTimeouts {
		return 0
	}

	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}

	return defaultConnectionTimeout
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.disableTimeouts {
		return 0
	}

	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}

	return defaultResponseTimeout
}

func (c *httpClient) requestBody(
	contentType string,
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

func (c *httpClient) requestHeaders(headers http.Header) http.Header {
	setHeaders := func(h http.Header, newHeaders http.Header) {
		for k, values := range newHeaders {
			if len(values) > 0 {
				h.Set(k, values[0])
			}
		}
	}

	h := make(http.Header)

	// Add common headers
	setHeaders(h, c.builder.headers)

	// Add custom headers
	setHeaders(h, headers)

	return h
}
