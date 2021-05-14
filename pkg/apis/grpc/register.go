package grpc

import (
	"github.com/ralstan-vaz/go-boilerplate/pkg/apis"
	"github.com/ralstan-vaz/go-boilerplate/pkg/apis/grpc/generated/user"
	user "github.com/ralstan-vaz/go-boilerplate/pkg/apis/grpc/user"

	"google.golang.org/grpc"
)

func registerService(server *grpc.Server, pkg apis.PackageInterface) {

	userService := user.NewUserService(pkg)
	proto.RegisterUserServiceServer(server, userService)
}
