package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	client *http.Client

	disableTimeouts   bool
	maxIdleConns      int
	connectionTimeout time.Duration
	responseTimeout   time.Duration

	Headers http.Header
}

// HTTPClient ...
type HTTPClient interface {
	DisableTimeouts(disable bool)
	SetMaxIdleConns(m int)
	SetConnectionTimeout(t time.Duration)
	SetResponseTimeout(t time.Duration)
	SetHeaders(headers http.Header)
	Get(string, http.Header) (*http.Response, error)
	Post(string, http.Header, interface{}) (*http.Response, error)
	Put(string, http.Header, interface{}) (*http.Response, error)
	Patch(string, http.Header, interface{}) (*http.Response, error)
	Delete(string, http.Header) (*http.Response, error)
}

// New ...
func New() HTTPClient {
	return &httpClient{}
}

// DisableTimeouts...
func (c *httpClient) DisableTimeouts(disable bool) {
	c.disableTimeouts = disable
}

// SetMaxIddleConns...
func (c *httpClient) SetMaxIdleConns(m int) {
	c.maxIdleConns = m
}

// SetConnectionTimeout ...
func (c *httpClient) SetConnectionTimeout(d time.Duration) {
	c.connectionTimeout = d
}

// SetRequestTimeout ...
func (c *httpClient) SetResponseTimeout(d time.Duration) {
	c.responseTimeout = d
}

// SetHeaders ...
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
