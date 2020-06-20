package http

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	httpUser "github.com/ralstan-vaz/go-boilerplate/apis/http/user"
	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user"
)

// StartServer It will initialize the routes and the server
func StartServer(conf *config.Config, pkg PackageInterface, wg *sync.WaitGroup) error {
	defer wg.Done()

	address := conf.Server.HTTP.Address

	router := gin.Default()

	// Initialize all the routes
	httpUser.NewUserRoute(conf, router, pkg)

	fmt.Println("HTTP Server listening on : ", address)

	// Start the server
	router.Run(address)

	return nil
}

type PackageInterface interface {
	NewUserPkg() *user.UserPkg
}
