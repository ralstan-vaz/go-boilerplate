package user

import "github.com/ralstan-vaz/go-boilerplate/pkg/user/repo"

// User contains all the properties of a user
type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// Stars come from an HTTP API
	Stars     string    `json:"stars,omitempty"`
	Favourite Favourite `json:"favourite,omitempty"`
}

// Favourite contains the list of users favourite stuff
type Favourite struct {
	// Beers come from a GRPC API
	Beers []string `json:"beers,omitempty"`
}

// Create interfaces only where they are being used

// getter ..
type getter interface {
	GetOne(id string) (*repo.User, error)
	Get(query string) ([]*repo.User, error)
	GetAll() ([]*repo.User, error)
}

// modifier ..
type modifier interface {
	Insert(repo.User) error
}

// operator is implemented by any value that contains getter and modifier interface methods
// Gives a way to mock the user repo
type operator interface {
	modifier
	getter
}
