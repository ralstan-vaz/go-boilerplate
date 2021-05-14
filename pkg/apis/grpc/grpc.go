package grpc

import (
  "fmt"
  "github.com/ralstan-vaz/go-boilerplate/config"
  "github.com/ralstan-vaz/go-boilerplate/pkg/apis"
  "google.golang.org/grpc"
  "net"
)

// Start starts the grpc server using the dependencies passed to it
func Start(conf *config.Config, pkg apis.PackageInterface) error {

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
