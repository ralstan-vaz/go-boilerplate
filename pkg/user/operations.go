package user

import (
	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/favourite"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/rating"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/repo"
)

// UserPkg ..
type UserPkg struct {
	config    *config.Config
	user      Operator
	rating    rating.Rater
	favourite favourite.FavouriteInterface
}

// NewUserPkg ...
func NewUserPkg(conf *config.Config, user Operator, rating rating.Rater, favourite favourite.FavouriteInterface) *UserPkg {
	return &UserPkg{config: conf, user: user, rating: rating, favourite: favourite}
}

// Get ...
func (pkg *UserPkg) Get(query string) ([]*repo.User, error) {
	users, err := pkg.user.Get(query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetOne ...
func (pkg *UserPkg) GetOne(id string) (*repo.User, error) {
	user, err := pkg.user.GetOne(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAll ...
func (pkg *UserPkg) GetAll() ([]*repo.User, error) {
	users, err := pkg.user.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Insert ...
func (pkg *UserPkg) Insert(u repo.User) error {
	err := pkg.user.Insert(u)
	if err != nil {
		return err
	}

	return nil
}

// GetWithRating ...
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
