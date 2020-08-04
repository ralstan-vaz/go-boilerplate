// +build unit

// go test -v -tags=unit
package user

import (
	"reflect"
	"testing"

	"github.com/ralstan-vaz/go-boilerplate/pkg/user/favourite"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/rating"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/repo"
)

func TestUserGet(t *testing.T) {

	repoUsers := []*repo.User{
		&repo.User{ID: "1", Name: "Miranda"},
		&repo.User{ID: "2", Name: "Shepard"},
	}

	mockUser := NewMockUser(repoUsers...)
	pkg := UserPkg{user: mockUser}
	users, err := pkg.Get("")
	if err != nil {
		return
	}

	var iuser interface{}
	iuser = users

	switch u := iuser.(type) {
	case []*User:
	default:
		t.Error("Returned Data type did not match", u)
	}

	matchUsers(t, users, repoUsers)
}
func TestUserGetOne(t *testing.T) {

	repoUsers := []*repo.User{
		&repo.User{ID: "1", Name: "Miranda"},
		&repo.User{ID: "2", Name: "Shepard"},
	}

	mockUser := NewMockUser(repoUsers...)
	pkg := UserPkg{user: mockUser}
	user, err := pkg.GetOne("")
	if err != nil {
		return
	}

	var iuser interface{}
	iuser = user

	switch u := iuser.(type) {
	case *User:
	default:
		t.Error("Returned Data type did not match", u)
	}
	matchUser(t, user, repoUsers[0])
}

func TestUserGetAll(t *testing.T) {

	repoUsers := []*repo.User{
		&repo.User{ID: "1", Name: "Miranda"},
		&repo.User{ID: "2", Name: "Shepard"},
		&repo.User{ID: "3", Name: "Skywalker"},
	}

	mockUser := NewMockUser(repoUsers...)
	pkg := UserPkg{user: mockUser}
	users, err := pkg.GetAll()
	if err != nil {
		return
	}

	var iuser interface{}
	iuser = users

	switch u := iuser.(type) {
	case []*User:
	default:
		t.Error("Returned Data type did not match", u)
	}

	matchUsers(t, users, repoUsers)
}

func TestUserInsert(t *testing.T) {

	userReq := User{ID: "2", Name: "Tali"}
	mockUser := NewMockUser()
	pkg := UserPkg{user: mockUser}
	err := pkg.Insert(userReq)
	if err != nil {
		return
	}

	repoUser := &repo.User{ID: userReq.ID, Name: userReq.Name}

	user, err := pkg.GetOne(userReq.ID)
	if err != nil {
		return
	}

	matchUser(t, user, repoUser)
}

func TestUserGetWithInfo(t *testing.T) {

	repoUsers := []*repo.User{
		&repo.User{ID: "1", Name: "Miranda"},
		&repo.User{ID: "2", Name: "Shepard"},
	}

	ratingRes := &rating.GetResponse{ID: "1", Stars: "5"}
	favouriteRes := &favourite.GetResponse{ID: "1", Beers: []string{"moonshine"}}

	mockUser := NewMockUser(repoUsers...)
	mockRating := NewMockRating(ratingRes)
	mockfavourite := NewMockFavourite(favouriteRes)

	pkg := UserPkg{
		user:      mockUser,
		rating:    mockRating,
		favourite: mockfavourite,
	}

	user, err := pkg.GetWithInfo("1")
	if err != nil {
		return
	}

	var iuser interface{}
	iuser = user

	switch u := iuser.(type) {
	case *User:
	default:
		t.Error("Returned Data type did not match", u)
	}

	matchUser(t, user, repoUsers[0])
	matchRating(t, user, ratingRes)

	// Since the response of Favourite's is hardcoded in the app this test would always fail
	// matchFavourite would demonstrate failure
	matchFavourite(t, user, favouriteRes)
}

func matchUsers(t *testing.T, users []*User, repoUsers []*repo.User) {
	for i := 0; i < len(users); i++ {
		matchUser(t, users[i], repoUsers[i])
	}
}

func matchUser(t *testing.T, user *User, repoUser *repo.User) {
	if user.ID != repoUser.ID {
		t.Error("ID failed to match")
	}

	if user.Name != repoUser.Name {
		t.Error("Name failed to match")
	}

}

func matchRating(t *testing.T, user *User, ratingRes *rating.GetResponse) {
	if user.ID != ratingRes.ID {
		t.Error("ID failed to match")
	}

	if user.Stars != ratingRes.Stars {
		t.Error("Stars failed to match")
	}
}

func matchFavourite(t *testing.T, user *User, favouriteRes *favourite.GetResponse) {
	if user.ID != favouriteRes.ID {
		t.Error("ID failed to match")
	}

	if !reflect.DeepEqual(user.Favourite.Beers, favouriteRes.Beers) {
		t.Error("Beers failed to match")
	}
}

type MockUser struct {
	users []*repo.User
}

func NewMockUser(users ...*repo.User) *MockUser {
	return &MockUser{users: users}
}

// Get Gets users using a query
func (mu *MockUser) Get(query string) ([]*repo.User, error) {
	return mu.users, nil
}

// GetOne Gets the 1st user in the array
func (mu *MockUser) GetOne(id string) (*repo.User, error) {
	return mu.users[0], nil
}

// GetAll Gets all the users
func (mu *MockUser) GetAll() ([]*repo.User, error) {
	return mu.users, nil
}

// Insert Inserts a User
func (mu *MockUser) Insert(u repo.User) error {
	mu.users = append(mu.users, &u)
	return nil
}

type MockRating struct {
	response *rating.GetResponse
}

// NewMockRating creates a new instance of MockRating
func NewMockRating(response *rating.GetResponse) *MockRating {
	return &MockRating{response: response}
}

// Get fetches the MockRating
func (mr *MockRating) Get(req rating.GetRequest) (*rating.GetResponse, error) {
	return mr.response, nil
}

// MockFavourite contains methods to peform operations on users Mockfavourites
type MockFavourite struct {
	response *favourite.GetResponse
}

// NewMockFavourite create an instance of Mockfavourite
func NewMockFavourite(response *favourite.GetResponse) *MockFavourite {
	return &MockFavourite{response: response}
}

// Get Make the request to Mockfavourites
func (f *MockFavourite) Get(req favourite.GetRequest) (*favourite.GetResponse, error) {
	res := favourite.GetResponse{}
	res.Beers = []string{"Moon Shine", "Bira", "Simba"}
	return &res, nil
}
