package http

import (
	"net/http"
)

func NewRequest() *Request {
	return &Request{}
}

type Request struct {
}

func (r *Request) Do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
