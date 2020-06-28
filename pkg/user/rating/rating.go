package rating

import (
	"net/http"
)

// Rater is implemented by any value that contains the required methods
type Rater interface {
	Get(GetRequest) (*GetResponse, error)
}

// httpRequester makes it possible to mock or intercept http request
type httpRequester interface {
	Do(*http.Request) (*http.Response, error)
}
