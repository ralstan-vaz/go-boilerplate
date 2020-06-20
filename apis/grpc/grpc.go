package grpc

import (
	"fmt"
	"net"
	"sync"

	"github.com/ralstan-vaz/go-boilerplate/config"
	userPkg "github.com/ralstan-vaz/go-boilerplate/pkg/user"
	"google.golang.org/grpc"
)

type PackageInterface interface {
	NewUserPkg() *userPkg.UserPkg
}

// StartServer ...
func StartServer(conf *config.Config, pkg PackageInterface, wg *sync.WaitGroup) error {

	address := conf.Server.GRPC.Address

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	var server *grpc.Server
	// Add required opts
	opts := []grpc.ServerOption{}
	server = grpc.NewServer(opts...)

	registerService(server, pkg)

	fmt.Println("GRPC Server listening on : ", address)

	err = server.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
