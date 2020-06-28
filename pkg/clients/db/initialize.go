package db

import (
	"github.com/ralstan-vaz/go-boilerplate/config"
)

// DBInstances contains all the DB instances this app uses
type DBInstances struct {
	conf *config.Config
	myDB *MyDB
}

// newDBInstances creates an instance of DBInstances
// It does not initialize the connections.
func newDBInstances(conf *config.Config) *DBInstances {
	return &DBInstances{conf: conf}
}

// NewInitializedInstances creates an instance of initialized DBInstances
func NewInitializedInstances(conf *config.Config) (*DBInstances, error) {
	dbInstance := newDBInstances(conf)
	err := dbInstance.initialize()
	if err != nil {
		return nil, err
	}
	return dbInstance, nil
}

// initialize ..
func (i *DBInstances) initialize() error {
	err := i.myDBInit()
	if err != nil {
		return err
	}
	return nil
}

// myDBInit ..
func (i *DBInstances) myDBInit() error {
	i.myDB = NewMyDB(i.conf)
	return nil
}

// GetMyDB return the instance for MYDb
func (i *DBInstances) GetMyDB() Dber {
	return i.myDB
}
