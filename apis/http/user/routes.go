package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user"
)

func NewUserRoute(conf *config.Config, router *gin.Engine, pkg PackageInterface) {
	BindRoutes(router, pkg)
}

func BindRoutes(router *gin.Engine, pkg PackageInterface) {
	service := NewUserService(pkg)
	userAPI := router.Group("/users")
	{
		userAPI.GET("/", service.getAll)
		userAPI.GET("/:userId", service.getOne)
		userAPI.GET("/:userId/rating", service.getWithRating)
		userAPI.POST("/", service.insert)
	}
}

type PackageInterface interface {
	NewUserPkg() *user.UserPkg
}
