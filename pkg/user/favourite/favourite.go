package favourite

import (
	"google.golang.org/grpc"
)

// FavouriteInterface is implemented by any value that contains the required methods
// makes Favourite mockable
type FavouriteInterface interface {
	Get(GetRequest) (*GetResponse, error)
}

// grpcConnectioner contains methods to retrieve grpc connections
// this makes it possible to pass multiple grpc connectiions incase the package needs it
type grpcConnectioner interface {
	GetFavourite() *grpc.ClientConn
}
