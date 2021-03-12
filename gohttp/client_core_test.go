package gohttp

import (
	"net/http"
	"testing"
)

func TestAllHeaders(t *testing.T) {
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "Go-HTTP-Client")
	client.Headers = commonHeaders

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.allHeaders(requestHeaders)

	if len(finalHeaders) != 3 {
		t.Error("Expected 3 headers, got", len(finalHeaders))
	}
}
