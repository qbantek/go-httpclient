package gohttp

import (
	"net/http"
	"time"
)

type clientBuilder struct {
	disableTimeouts   bool
	maxIdleConns      int
	connectionTimeout time.Duration
	responseTimeout   time.Duration

	headers http.Header
}

// ClientBuilder is the top of the top
type ClientBuilder interface {
	SetMaxIdleConns(m int) ClientBuilder
	SetConnectionTimeout(t time.Duration) ClientBuilder
	SetResponseTimeout(t time.Duration) ClientBuilder
	SetHeaders(headers http.Header) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	Build() *httpClient
}

// NewBuilder returns a new clientBuilder
func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}

// Build...
func (c *clientBuilder) Build() *httpClient {
	client := httpClient{builder: c}
	return &client
}

// SetMaxIddleConns...
func (c *clientBuilder) SetMaxIdleConns(m int) ClientBuilder {
	c.maxIdleConns = m
	return c
}

// SetConnectionTimeout ...
func (c *clientBuilder) SetConnectionTimeout(d time.Duration) ClientBuilder {
	c.connectionTimeout = d
	return c
}

// SetRequestTimeout ...
func (c *clientBuilder) SetResponseTimeout(d time.Duration) ClientBuilder {
	c.responseTimeout = d
	return c
}

// SetHeaders ...
func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

// DisableTimeouts...
func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}
