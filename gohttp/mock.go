package gohttp

import (
	"fmt"
	"net/http"
)

// Mock allow us to mock Requests
type Mock struct {
	Method      string
	URL         string
	RequestBody string

	ResponseBody       string
	Error              error
	ResponseStatusCode int
}

func (m *Mock) Response() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	response := Response{
		body:       []byte(m.ResponseBody),
		statusCode: m.ResponseStatusCode,
		status: fmt.Sprintf(
			"%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
	}
	return &response, nil
}
