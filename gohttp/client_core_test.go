package gohttp

import (
	"net/http"
	"testing"
)

func TestRequestHeaders(t *testing.T) {
	commonHeaders := func() http.Header {
		t.Helper()
		h := make(http.Header)
		h.Set("Content-Type", "application/json")
		h.Set("User-Agent", "Go-HTTP-Client")

		return h
	}

	c := NewBuilder().
		SetHeaders(commonHeaders()).
		Build()

	customHeaders := func() http.Header {
		t.Helper()
		h := make(http.Header)
		h.Set("Content-Type", "application/xml")
		h.Set("X-Request-Id", "ABC-123")

		return h
	}
	headers := c.requestHeaders(customHeaders())

	t.Run("custom headers override matching common headers", func(t *testing.T) {
		if got := headers.Get("Content-Type"); got != "application/xml" {
			t.Errorf("got %v, want %v", got, "application/xml")
		}
	})

	t.Run("custom and common headers get merged", func(t *testing.T) {
		if got := len(headers); got != 3 {
			t.Errorf("got %v headers, want %v", got, 3)
		}
	})
}

func TestRequestBody(t *testing.T) {
	c := httpClient{}

	t.Run("Null body", func(t *testing.T) {
		got, err := c.requestBody("anything", nil)
		if err != nil {
			t.Errorf("got error: %v, want no error", err)
		}

		if got != nil {
			t.Errorf("got %v, want %v", got, nil)
		}
	})

	type s struct{ A string }
	b := &s{A: "b"}
	tests := map[string]struct {
		contentType string
		body        *s
		want        string
	}{
		"xml content type":     {"application/xml", b, "<s><A>b</A></s>"},
		"json content type":    {"application/json", b, "{\"A\":\"b\"}"},
		"default content type": {"any-other", b, "{\"A\":\"b\"}"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := c.requestBody(tc.contentType, tc.body)
			if err != nil {
				t.Errorf("got error: %v, want no error", err)
			}

			if string(got) != tc.want {
				t.Errorf("got %v, want %v", string(got), tc.want)
			}
		})
	}
}
