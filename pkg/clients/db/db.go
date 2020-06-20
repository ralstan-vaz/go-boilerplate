package db

type Dber interface {
	Getter
	Inserter
}

// MimicUser Just to minic user collection
type MimicUser struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Getter interface {
	GetOne(id string) MimicUser
	Get(query string) []MimicUser
	GetAll() []MimicUser
}

type Inserter interface {
	Insert(interface{}) error
}
