package initiate

import (
	"fmt"

	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/apis/grpc"
	"github.com/ralstan-vaz/go-boilerplate/pkg/apis/http"
	"github.com/ralstan-vaz/go-boilerplate/pkg/clients/db"
	grpcPkg "github.com/ralstan-vaz/go-boilerplate/pkg/clients/grpc"
)

// Initialize will initialize all the dependencies and the servers.
// Dependencies include config, external connections(grpc, http)
func Initialize() error {
	env, err := Env()
	if err != nil {
		return err
	}

	fmt.Println("Enviroment: ", env)

	conf, err := config.NewConfig(env)
	if err != nil {
		return err
	}

	dbInstances, err := db.NewInitializedInstances(conf)
	if err != nil {
		return err
	}

	grpcCons, err := grpcPkg.NewInitializeConnections(conf)
	if err != nil {
		return err
	}

	err = StartServers(conf, dbInstances, grpcCons)
	if err != nil {
		return err
	}

	return nil
}

// StartServers will pass the dependencies to the servers.
// The servers will start in an individual goroutine
// Wait group is used to wait for all the goroutines launched here to finish.
// In in ideal scenario the routines would run indefinitely
func StartServers(conf *config.Config, dbInstances *db.DBInstances, grpcCons *grpcPkg.GrpcConnections) error {

	pkg := NewPackageDeps(conf, dbInstances, grpcCons)

	var errCh = make(chan error)

	go startHTTP(conf, pkg, errCh)
	go startGRPC(conf, pkg, errCh)

	// Wait until an error is received through the channel
	// We want the program to exit if anyone of the servers return an error
	select {
	case err := <-errCh:
		close(errCh)
		return err
	}
}

func startHTTP(conf *config.Config, pkg *PackageDeps, errCh chan error) {
	err := http.Start(conf, pkg)
	if err != nil {
		errCh <- err
	}
}

func startGRPC(conf *config.Config, pkg *PackageDeps, errCh chan error) {
	err := grpc.Start(conf, pkg)
	if err != nil {
		errCh <- err
	}
}
