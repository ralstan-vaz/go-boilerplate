package db

// Dber contains methods to operate on a DB
type Dber interface {
	getter
	inserter
}

// MimicUser Just to minic user collection
type MimicUser struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type getter interface {
	GetOne(id string) MimicUser
	Get(query string) []MimicUser
	GetAll() []MimicUser
}

type inserter interface {
	Insert(interface{}) error
}
