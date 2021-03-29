package gohttp

import (
	"encoding/json"
	"net/http"
)

// Response represents the response from an HTTP request.
type Response struct {
	statusCode int
	status     string
	headers    http.Header
	body       []byte
}

// Status returns the status of the Response
func (r *Response) Status() string {
	return r.status
}

// StatusCode returns the status code of the Response
func (r *Response) StatusCode() int {
	return r.statusCode
}

// Headers returns the headers of the Response
func (r *Response) Headers() http.Header {
	return r.headers
}

// Bytes returns the body of the Response as a slice of bytes
func (r *Response) Bytes() []byte {
	return r.body
}

// Bytes returns the body of the Response as a string
func (r *Response) String() string {
	return string(r.body)
}

// UnmarshalJSONBody unmarshals and returns the body of the Response using the json
// package
func (r *Response) UnmarshalJSONBody(target interface{}) error {
	return json.Unmarshal(r.body, target)
}
