package favourite

import (
	"context"
	"fmt"

	user "github.com/ralstan-vaz/go-boilerplate/apis/grpc/generated/user"
	"github.com/ralstan-vaz/go-boilerplate/config"
)

// Favourite contains methods to peform operations on users favourites
type Favourite struct {
	config  *config.Config
	grpcCon grpcConnectioner
}

// NewFavourite create an instance of favourite
func NewFavourite(conf *config.Config, grpcConn grpcConnectioner) *Favourite {
	return &Favourite{config: conf, grpcCon: grpcConn}
}

// GetRequest request for getting a favourite
type GetRequest struct {
	ID string
}

// GetResponse response after getting a favourite
type GetResponse struct {
	ID    string   `json:"id,omitempty"`
	Beers []string `json:"beers,omitempty"`
}

// Get Make the request to favourites
func (f *Favourite) Get(req GetRequest) (*GetResponse, error) {
	// Note: The below call is only made to demonstrate how a GRPC call would would work
	// The response of the call is ignored and a hard coded response is sent
	favGrpcCon := f.grpcCon.GetFavourite()
	cli := user.NewUserServiceClient(favGrpcCon)
	resp, err := cli.GetAll(context.Background(), &user.UserGetRequest{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Response", resp)

	// Hard coded response
	res := GetResponse{}
	res.Beers = []string{"Moon Shine", "Bira", "Simba"}
	return &res, nil
}
