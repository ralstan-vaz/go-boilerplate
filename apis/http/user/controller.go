package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user "github.com/ralstan-vaz/go-boilerplate/pkg/user"
)

// NewUserService Create a new instance of a UserService with the given dependencies.
func NewUserService(pkg PackageInterface) *UserService {
	return &UserService{pkg: pkg}
}

// UserService contains the methods required to perfom operation's on users
type UserService struct {
	pkg PackageInterface
}

func (u *UserService) getAll(c *gin.Context) {
	userPkg := u.pkg.NewUserPkg()
	users, err := userPkg.GetAll()
	if err != nil {
		errorCode := http.StatusBadRequest
		if err.Error() == "No such order found in database" {
			errorCode = http.StatusNotFound
		}
		c.JSON(errorCode, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (u *UserService) getOne(c *gin.Context) {
	userId := c.Param("userId")

	userPkg := u.pkg.NewUserPkg()
	users, err := userPkg.GetOne(userId)
	if err != nil {
		errorCode := http.StatusBadRequest
		if err.Error() == "No such order found in database" {
			errorCode = http.StatusNotFound
		}
		c.JSON(errorCode, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (u *UserService) getWithInfo(c *gin.Context) {
	userId := c.Param("userId")

	userPkg := u.pkg.NewUserPkg()
	users, err := userPkg.GetWithInfo(userId)
	if err != nil {
		errorCode := http.StatusBadRequest
		if err.Error() == "No such order found in database" {
			errorCode = http.StatusNotFound
		}
		c.JSON(errorCode, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (u *UserService) insert(c *gin.Context) {

	var user user.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userPkg := u.pkg.NewUserPkg()
	err := userPkg.Insert(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
	return
}
