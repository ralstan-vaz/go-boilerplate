package repo

import (
	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/clients/db"
)

// UserRepo Contains methods to action on the User Repository
type UserRepo struct {
	config *config.Config
	db     db.Dber
}

// NewUserRepo Create's an instance of a User Repository
func NewUserRepo(conf *config.Config, dbInstances dbInstancer) *UserRepo {
	return &UserRepo{config: conf, db: dbInstances.GetMyDB()}
}

// Get Gets users using a query
func (ur *UserRepo) Get(query string) ([]*User, error) {
	u := ur.db.Get(query)
	users := bindToUsers(u)
	return users, nil
}

// GetOne Gets a user user an Id
func (ur *UserRepo) GetOne(id string) (*User, error) {
	u := ur.db.GetOne(id)
	user := User(u)
	return &user, nil
}

// GetAll Gets all the users
func (ur *UserRepo) GetAll() ([]*User, error) {
	u := ur.db.GetAll()
	users := bindToUsers(u)
	return users, nil
}

// Insert Inserts a User
func (ur *UserRepo) Insert(u User) error {
	ur.db.Insert(u)
	return nil
}

func bindToUsers(u []db.MimicUser) []*User {
	user := []*User{}
	for i := 0; i < len(u); i++ {
		user = append(user, &User{ID: u[i].ID, Name: u[i].Name})
	}
	return user
}
