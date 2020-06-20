package grpc

import (
	pb "github.com/ralstan-vaz/go-boilerplate/apis/grpc/generated/user"
	"github.com/ralstan-vaz/go-boilerplate/apis/grpc/user"

	"google.golang.org/grpc"
)

func registerService(server *grpc.Server, pkg PackageInterface) {

	userService := user.NewUserService(pkg)
	// Bind the RPC services to the grpc server
	pb.RegisterUserServiceServer(server, userService)
}
