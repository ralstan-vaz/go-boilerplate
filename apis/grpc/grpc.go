package grpc

import (
	"fmt"
	"net"
	"sync"

	"github.com/ralstan-vaz/go-boilerplate/config"
	userPkg "github.com/ralstan-vaz/go-boilerplate/pkg/user"
	"google.golang.org/grpc"
)

// PackageInterface contains methods which return dependencies that are used by the services.
// This aids in making it possible to send dependencies to the packages.
type PackageInterface interface {
	NewUserPkg() *userPkg.UserPkg
}

// StartServer starts the grpc server using the dependencies passed to it
func StartServer(conf *config.Config, pkg PackageInterface, wg *sync.WaitGroup) error {

	address := conf.Server.GRPC.Address

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	var server *grpc.Server
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
