package rating

import (
	"encoding/json"
	"io/ioutil"

	"net/http"
)

// GetRequest ..
type GetRequest struct {
	ID string
}

// GetResponse ..
type GetResponse struct {
	ID    string `json:"id,omitempty"`
	Stars string `json:"stars,omitempty"`
}

// Get makes the request to get the ratings
func (r *Rating) Get(req GetRequest) (*GetResponse, error) {

	httpReq, err := http.NewRequest("GET", r.config.User.RatingsUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.httpRequester.Do(httpReq)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := GetResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
