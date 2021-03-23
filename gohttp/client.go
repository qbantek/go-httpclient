package gohttp

import (
	"net/http"
	"sync"
)

type httpClient struct {
	client  *http.Client
	builder *clientBuilder
	once    sync.Once
}

// Client ...
type Client interface {
	Get(string, http.Header) (*Response, error)
	Post(string, http.Header, interface{}) (*Response, error)
	Put(string, http.Header, interface{}) (*Response, error)
	Patch(string, http.Header, interface{}) (*Response, error)
	Delete(string, http.Header) (*Response, error)
}

// Get issues a GET to the specified URL.
func (c *httpClient) Get(url string,
	headers http.Header) (*Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

// Post ...
func (c *httpClient) Post(url string,
	headers http.Header,
	body interface{}) (*Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

// Put ...
func (c *httpClient) Put(url string,
	headers http.Header,
	body interface{}) (*Response, error) {
	return c.do(http.MethodPut, url, headers, nil)
}

// Patch ...
func (c *httpClient) Patch(url string,
	headers http.Header,
	body interface{}) (*Response, error) {
	return c.do(http.MethodPatch, url, headers, nil)
}

// Delete ...
func (c *httpClient) Delete(url string,
	headers http.Header) (*Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
