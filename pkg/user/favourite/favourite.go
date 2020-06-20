package favourite

import (
	"fmt"

	"github.com/ralstan-vaz/go-boilerplate/config"
	"google.golang.org/grpc"
)

// Favourite ...
type Favourite struct {
	config  *config.Config
	grpcCon grpcConnectioner
}

// NewFavourite ...
func NewFavourite(conf *config.Config, grpcConn grpcConnectioner) *Favourite {
	return &Favourite{config: conf, grpcCon: grpcConn}
}

// Pkg ...
type Pkg struct {
	favourite FavouriteInterface
}

// NewPkg ...
func NewPkg(favourite FavouriteInterface) *Pkg {
	return &Pkg{favourite: favourite}
}

// FavouriteInterface ..
type FavouriteInterface interface {
	Get(GetRequest) (*GetResponse, error)
}

type grpcConnectioner interface {
	GetFavourite() *grpc.ClientConn
}

// Get ..
func (p *Pkg) Get(req GetRequest) (*GetResponse, error) {
	favourites, err := p.favourite.Get(req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Getting favourites", favourites)
	return favourites, nil
}
