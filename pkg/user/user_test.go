// +build unit

// go test -v -tags=unit
package user

import (
	"testing"

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
		t.Log("Return type matched")
	default:
		t.Error("Returned Data type did not match", u)
	}

	for i := 0; i < len(users); i++ {
		if users[i].ID != repoUsers[i].ID {
			t.Error("ID failed to match")
		}

		if users[i].Name != repoUsers[i].Name {
			t.Error("Name failed to match")
		}
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
