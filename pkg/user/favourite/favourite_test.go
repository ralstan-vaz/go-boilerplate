package favourite

import (
	"context"
	"fmt"
	"testing"

	user "github.com/ralstan-vaz/go-boilerplate/apis/grpc/generated/user"
	"github.com/ralstan-vaz/go-boilerplate/config"
	grpcPkg "github.com/ralstan-vaz/go-boilerplate/pkg/grpc"
)

func TestBeerGet(t *testing.T) {
	conf := &config.Config{}
	err := grpcPkg.Initialize(conf)
	if err != nil {
		panic(err)
	}

	beerCon := grpcPkg.GetFavouriteConnection()
	cli := user.NewUserServiceClient(beerCon)
	fmt.Println("Client", beerCon)

	resp, err := cli.GetAll(context.Background(), &user.UserGetRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Response", resp)
	defer beerCon.Close()
}
