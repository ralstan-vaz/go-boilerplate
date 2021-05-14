package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/apis"
	user "github.com/ralstan-vaz/go-boilerplate/pkg/apis/http/user"
)

// Start starts the http server using the dependencies passed to it.
// It also initializes the routes
func Start(conf *config.Config, pkg apis.PackageInterface) error {

	address := conf.Server.HTTP.Address

	router := gin.Default()

	// Initialize all the routes
	user.NewUserRoute(conf, router, pkg)

	fmt.Println("HTTP Server listening on : ", address)

	// Start the server
	router.Run(address)

	return nil
}
