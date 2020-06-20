package repo

import (
	"github.com/ralstan-vaz/go-boilerplate/pkg/clients/db"
)

type User struct {
	ID   string `json:"Id,omitempty"`
	Name string `json:"name,omitempty"`
}

type DBInstancer interface {
	GetMyDB() db.Dber
}
