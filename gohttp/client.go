package gohttp

import (
	"net/http"
)

type httpClient struct {
	Headers http.Header
}

// HTTPClient ...
type HTTPClient interface {
	SetHeaders(headers http.Header)
	Get(string, http.Header) (*http.Response, error)
	Post(string, http.Header, interface{}) (*http.Response, error)
	Put(string, http.Header, interface{}) (*http.Response, error)
	Patch(string, http.Header, interface{}) (*http.Response, error)
	Delete(string, http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

// Get issues a GET to the specified URL.
func (c *httpClient) Get(url string,
	headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

// Post ...
func (c *httpClient) Post(url string,
	headers http.Header,
	body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

// Put ...
func (c *httpClient) Put(url string,
	headers http.Header,
	body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, nil)
}

// Patch ...
func (c *httpClient) Patch(url string,
	headers http.Header,
	body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, nil)
}

// Delete ...
func (c *httpClient) Delete(url string,
	headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

// New ...
func New() HTTPClient {
	client := &httpClient{}
	return client
}
