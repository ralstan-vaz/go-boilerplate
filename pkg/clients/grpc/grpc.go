package grpc

import (
	"github.com/ralstan-vaz/go-boilerplate/config"
	"google.golang.org/grpc"
)

type GrpcConnections struct {
	favouriteConnection *grpc.ClientConn
	conf                *config.Config
}

func NewGrpcConnections(conf *config.Config) *GrpcConnections {
	return &GrpcConnections{conf: conf}
}

func NewInitializeConnections(conf *config.Config) (*GrpcConnections, error) {
	grpcCons := NewGrpcConnections(conf)
	err := grpcCons.initialize()
	if err != nil {
		return nil, err
	}
	return grpcCons, nil
}

// Initialize ..
func (g *GrpcConnections) initialize() error {
	err := g.favouriteInit()
	if err != nil {
		return err
	}
	return nil
}

// FavouriteInit ..
func (g *GrpcConnections) favouriteInit() error {
	favouriteCon, err := grpc.Dial(g.conf.User.FavouritesUrl, grpc.WithInsecure())
	if err != nil {
		return err
	}
	g.favouriteConnection = favouriteCon
	return nil
}

func (g *GrpcConnections) GetFavourite() *grpc.ClientConn {
	return g.favouriteConnection
}
