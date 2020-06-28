package user

import (
	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/favourite"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/rating"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/repo"
)

// UserPkg provides a way to perform operations on a user
type UserPkg struct {
	config    *config.Config
	user      operator
	rating    rating.Rater
	favourite favourite.FavouriteInterface
}

// NewUserPkg creates an instance of UserPkg using the dependencies passed
// The dependency params can be moved to an interface to reduce to make it clean
func NewUserPkg(conf *config.Config, user operator, rating rating.Rater, favourite favourite.FavouriteInterface) *UserPkg {
	return &UserPkg{config: conf, user: user, rating: rating, favourite: favourite}
}

// Get gets users from the store using the query passed
func (pkg *UserPkg) Get(query string) ([]*repo.User, error) {
	users, err := pkg.user.Get(query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetOne gets a user from the store using the query
func (pkg *UserPkg) GetOne(id string) (*repo.User, error) {
	user, err := pkg.user.GetOne(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAll gets all the users
func (pkg *UserPkg) GetAll() ([]*repo.User, error) {
	users, err := pkg.user.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Insert stores a user
func (pkg *UserPkg) Insert(u repo.User) error {
	err := pkg.user.Insert(u)
	if err != nil {
		return err
	}

	return nil
}

// GetWithRating get a user from the store along with the ratings
func (pkg *UserPkg) GetWithRating(id string) (*User, error) {
	repoUser, err := pkg.user.GetOne(id)
	if err != nil {
		return nil, err
	}

	rating, err := pkg.rating.Get(rating.GetRequest{ID: id})
	if err != nil {
		return nil, err
	}

	favourite, err := pkg.favourite.Get(favourite.GetRequest{ID: id})
	if err != nil {
		return nil, err
	}

	fav := Favourite{Beers: favourite.Beers}
	user := &User{
		ID:        repoUser.ID,
		Name:      repoUser.Name,
		Stars:     rating.Stars,
		Favourite: fav,
	}

	return user, nil
}
