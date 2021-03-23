package gohttp

import (
	"net/http"

	"gopkg.in/square/go-jose.v2/json"
)

type Response struct {
	statusCode int
	status     string
	headers    http.Header
	body       []byte
}

func (r *Response) Status() string {
	return r.status
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) Headers() http.Header {
	return r.headers
}

func (r *Response) Bytes() []byte {
	return r.body
}

func (r *Response) String() string {
	return string(r.body)
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.body, target)
}
