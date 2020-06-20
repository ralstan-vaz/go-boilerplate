package rating

import (
	"fmt"
	"net/http"

	"github.com/ralstan-vaz/go-boilerplate/config"
)

// Rater ..
type Rater interface {
	Get(GetRequest) (*GetResponse, error)
}

// Rating ...
type Rating struct {
	config        *config.Config
	httpRequester httpRequester
}

// NewRating ...
func NewRating(conf *config.Config, httpRequester httpRequester) *Rating {
	return &Rating{config: conf, httpRequester: httpRequester}
}

// Pkg ...
type Pkg struct {
	rater Rater
}

// NewPkg ...
func NewPkg(rater Rater) *Pkg {
	return &Pkg{rater: rater}
}

// Get ..
func (p *Pkg) Get(req GetRequest) (*GetResponse, error) {
	ratings, err := p.rater.Get(req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting ratings", ratings)
	return ratings, nil
}

type httpRequester interface {
	Do(*http.Request) (*http.Response, error)
}
