package db

import (
	"github.com/ralstan-vaz/go-boilerplate/config"
)

type DBInstances struct {
	conf *config.Config
	myDB *MyDB
}

func NewDBInstances(conf *config.Config) *DBInstances {
	return &DBInstances{conf: conf}
}

func NewInitializedInstances(conf *config.Config) (*DBInstances, error) {
	dbInstance := NewDBInstances(conf)
	err := dbInstance.Initialize()
	if err != nil {
		return nil, err
	}
	return dbInstance, nil
}

// Initialize ..
func (i *DBInstances) Initialize() error {
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

func (i *DBInstances) GetMyDB() Dber {
	return i.myDB
}
