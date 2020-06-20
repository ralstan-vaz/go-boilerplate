package repo

import (
	"fmt"
	"testing"

	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/db"
)

func TestUserGet(t *testing.T) {
	conf := config.Config{}
	database := db.NewMyDB()
	userRepo := NewUserRepo(&conf, database)
	user, err := userRepo.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
