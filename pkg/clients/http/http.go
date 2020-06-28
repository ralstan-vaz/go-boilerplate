package http

import (
	"net/http"
)

// NewRequest Creates an instance if a request
func NewRequest() *Request {
	return &Request{}
}

// Request contains a method to perform an HTTP request
type Request struct {
}

// Do makes an http request
// Interceptors can be added here for tracing etc
func (r *Request) Do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
