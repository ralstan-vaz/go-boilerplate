package user

import (
  "github.com/gin-gonic/gin"
  "github.com/ralstan-vaz/go-boilerplate/config"
  "github.com/ralstan-vaz/go-boilerplate/pkg/apis"
)

// NewUserRoute Creates and initializes user routes
func NewUserRoute(conf *config.Config, router *gin.Engine, pkg apis.PackageInterface) {
	bindRoutes(router, pkg)
}

func bindRoutes(router *gin.Engine, pkg apis.PackageInterface) {
	service := NewUserService(pkg)
	userAPI := router.Group("/users")
	{
		userAPI.GET("/", service.getAll)
		userAPI.GET("/:userId", service.getOne)
		userAPI.GET("/:userId/rating", service.getWithInfo)
		userAPI.POST("/", service.insert)
	}
}
