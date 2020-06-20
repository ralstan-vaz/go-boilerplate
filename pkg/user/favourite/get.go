package favourite

import (
	"context"
	"fmt"

	user "github.com/ralstan-vaz/go-boilerplate/apis/grpc/generated/user"
)

// GetRequest ..
type GetRequest struct {
	ID string
}

// GetResponse ..
type GetResponse struct {
	ID    string   `json:"id,omitempty"`
	Beers []string `json:"beers,omitempty"`
}

// Get Make the request to get Beer
func (f *Favourite) Get(req GetRequest) (*GetResponse, error) {
	favGrpcCon := f.grpcCon.GetFavourite()
	cli := user.NewUserServiceClient(favGrpcCon)
	resp, err := cli.GetAll(context.Background(), &user.UserGetRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Response", resp)
	res := GetResponse{}
	res.Beers = []string{"Moon Shine", "Bira", "Simba"}
	return &res, nil
}
