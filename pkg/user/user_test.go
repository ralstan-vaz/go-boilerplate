package user

import (
	"fmt"
	"testing"

	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/db"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/rating"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/repo"
)

func TestUserGet(t *testing.T) {
	conf := &config.Config{}
	database := db.NewMyDB()
	userRepo := repo.NewUserRepo(conf, database)
	userRating := rating.NewRating(conf)
	userPkg := NewUserPkg(conf, userRepo, userRating)
	fmt.Println(userPkg.GetOne("UniqueId"))
}
