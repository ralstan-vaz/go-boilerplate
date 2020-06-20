package user

import "github.com/ralstan-vaz/go-boilerplate/pkg/user/repo"

// User ..
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

// Getter ..
type Getter interface {
	GetOne(id string) (*repo.User, error)
	Get(query string) ([]*repo.User, error)
	GetAll() ([]*repo.User, error)
}

// Modifier ..
type Modifier interface {
	Insert(repo.User) error
}

// Operator ..
type Operator interface {
	Modifier
	Getter
}
