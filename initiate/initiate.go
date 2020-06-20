package initiate

import (
	"fmt"
	"sync"

	"github.com/ralstan-vaz/go-boilerplate/apis/grpc"
	"github.com/ralstan-vaz/go-boilerplate/apis/http"
	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/clients/db"
	grpcPkg "github.com/ralstan-vaz/go-boilerplate/pkg/clients/grpc"
)

// Initialize the application
func Initialize() error {
	env, err := Env()
	if err != nil {
		return err
	}

	fmt.Println("Enviroment: ", env)

	// Gets config
	conf, err := config.NewConfig(env)
	if err != nil {
		return err
	}

	// Initializes the DB connections
	dbInstances, err := db.NewInitializedInstances(conf)
	if err != nil {
		return err
	}

	// Initializes the GRPC connections
	grpcCons, err := grpcPkg.NewInitializeConnections(conf)
	if err != nil {
		return err
	}

	InitServers(conf, dbInstances, grpcCons)

	return nil
}

// InitServers the HTTP and the gRPC servers
func InitServers(conf *config.Config, dbInstances *db.DBInstances, grpcCons *grpcPkg.GrpcConnections) {

	// Deps
	pkg := NewPackageDeps(conf, dbInstances, grpcCons)
	var wg sync.WaitGroup
	wg.Add(1)
	go http.StartServer(conf, pkg, &wg)

	wg.Add(1)
	go grpc.StartServer(conf, pkg, &wg)

	wg.Wait()
}