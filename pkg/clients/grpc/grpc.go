package grpc

import (
	"github.com/ralstan-vaz/go-boilerplate/config"
	"google.golang.org/grpc"
)

// GrpcConnections contains all the GRPC connections this app uses
type GrpcConnections struct {
	favouriteConnection *grpc.ClientConn
	conf                *config.Config
}

// newGrpcConnections creates an instance of GrpcConnections
// It does not initialize the connections.
func newGrpcConnections(conf *config.Config) *GrpcConnections {
	return &GrpcConnections{conf: conf}
}

// NewInitializeConnections creates an instance of initialized GrpcConnections
func NewInitializeConnections(conf *config.Config) (*GrpcConnections, error) {
	grpcCons := newGrpcConnections(conf)
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

// GetFavourite return the grpc connection for the favourite service
func (g *GrpcConnections) GetFavourite() *grpc.ClientConn {
	return g.favouriteConnection
}
