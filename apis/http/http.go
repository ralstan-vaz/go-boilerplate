package http

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	httpUser "github.com/ralstan-vaz/go-boilerplate/apis/http/user"
	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user"
)

// StartServer starts the http server using the dependencies passed to it.
// It also initializes the routes
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

// PackageInterface contains methods which return dependencies that are used by the services.
// This aids in making it possible to send dependencies to the packages.
type PackageInterface interface {
	NewUserPkg() *user.UserPkg
}
