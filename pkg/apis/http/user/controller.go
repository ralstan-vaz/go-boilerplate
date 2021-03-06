package user

import (
	"net/http"

	"github.com/ralstan-vaz/go-boilerplate/pkg/apis"
	utils "github.com/ralstan-vaz/go-boilerplate/pkg/apis/http/utils"

	"github.com/gin-gonic/gin"
	user "github.com/ralstan-vaz/go-boilerplate/pkg/user"
)

// NewUserService Create a new instance of a UserService with the given dependencies.
func NewUserService(pkg apis.PackageInterface) *UserService {
	return &UserService{pkg: pkg}
}

// UserService contains the methods required to perform operation's on users
type UserService struct {
	pkg apis.PackageInterface
}

func (u *UserService) getAll(c *gin.Context) {
	userPkg := u.pkg.NewUserPkg()
	users, err := userPkg.GetAll()
	if err != nil {
		utils.HandleError(c, err)
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserService) getOne(c *gin.Context) {
	userID := c.Param("userID")

	userPkg := u.pkg.NewUserPkg()
	users, err := userPkg.GetOne(userID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserService) getWithInfo(c *gin.Context) {
	userID := c.Param("userID")

	userPkg := u.pkg.NewUserPkg()
	users, err := userPkg.GetWithInfo(userID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserService) insert(c *gin.Context) {

	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.HandleError(c, err)
		return
	}

	userPkg := u.pkg.NewUserPkg()
	err := userPkg.Insert(user)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
