package db

import (
	"github.com/ralstan-vaz/go-boilerplate/config"
)

// NewMyDB ..
func NewMyDB(conf *config.Config) *MyDB {
	return &MyDB{connection: "connection established"}
}

// MyDB ..
type MyDB struct {
	connection string
}

// GetOne ..
func (m *MyDB) GetOne(id string) MimicUser {
	return MimicUser{ID: id, Name: "Shepard"}
}

// Get ..
func (m *MyDB) Get(query string) []MimicUser {

	return []MimicUser{{ID: "1", Name: "Shepard"}, {ID: "2", Name: "Miranda"}}
}

// GetAll ..
func (m *MyDB) GetAll() []MimicUser {
	return []MimicUser{{ID: "1", Name: "Shepard"},
		{ID: "2", Name: "Miranda"}, {ID: "3", Name: "Tali"}}
}

// Insert ..
func (m *MyDB) Insert(obj interface{}) error {
	return nil
}
