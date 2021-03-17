package gohttp

import (
	"bytes"
	"encoding/xml"
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

func (c *httpClient) do(method string,
	url string,
	headers http.Header,
	body interface{}) (*http.Response, error) {

	reqHeaders := c.allHeaders(headers)
	reqBody, err := c.requestBody(reqHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header = reqHeaders

	client := c.getHttpClient()
	return client.Do(req)
}

func (c *httpClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

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

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.maxIdleConns > 0 {
		return c.maxIdleConns
	}

	return http.DefaultMaxIdleConnsPerHost
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.disableTimeouts {
		return 0
	}

	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	}

	return defaultConnectionTimeout
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.disableTimeouts {
		return 0
	}

	if c.responseTimeout > 0 {
		return c.responseTimeout
	}

	return defaultResponseTimeout
}

func (c *httpClient) allHeaders(headers http.Header) http.Header {
	setHeaders := func(h http.Header, newHeaders http.Header) {
		for k, values := range newHeaders {
			if len(values) > 0 {
				h.Set(k, values[0])
			}
		}
	}

	h := make(http.Header)

	// Add common headers
	setHeaders(h, c.Headers)

	// Add custom headers
	setHeaders(h, headers)

	return h
}
