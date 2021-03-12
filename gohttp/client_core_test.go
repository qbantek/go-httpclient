package gohttp

import (
	"net/http"
	"testing"
)

func TestAllHeaders(t *testing.T) {
	c := httpClient{}
	c.Headers = commonHeaders()

	headers := c.allHeaders(customHeaders())

	// matching common headers are overwritten by custom ones
	const wantCt = "application/xml"
	if ct := headers.Get("Content-Type"); ct != wantCt {
		t.Errorf("got %v Content Type, want %v", ct, wantCt)
	}

	// all headers = common + custom (-overwritten)
	const wantTotal = 3
	if a := len(headers); a != wantTotal {
		t.Errorf("got %v headers, want %v", a, wantTotal)
	}
}

func commonHeaders() http.Header {
	common := make(http.Header)
	common.Set("Content-Type", "application/json")
	common.Set("User-Agent", "Go-HTTP-Client")

	return common
}

func customHeaders() http.Header {
	custom := make(http.Header)
	custom.Set("Content-Type", "application/xml")
	custom.Set("X-Request-Id", "ABC-123")

	return custom
}
