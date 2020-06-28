package repo

import (
	"github.com/ralstan-vaz/go-boilerplate/pkg/clients/db"
)

// User contains the properties stored in the repo
type User struct {
	ID   string `json:"Id,omitempty"`
	Name string `json:"name,omitempty"`
}

// dbInstancer is implemented by any value that contains the GetMyDB method.
// Makes it possible to inject an instance
type dbInstancer interface {
	GetMyDB() db.Dber
}
