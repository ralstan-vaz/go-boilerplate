package initiate

import (
	"github.com/ralstan-vaz/go-boilerplate/config"
	"github.com/ralstan-vaz/go-boilerplate/pkg/clients/db"
	grpcPkg "github.com/ralstan-vaz/go-boilerplate/pkg/clients/grpc"
	httpPkg "github.com/ralstan-vaz/go-boilerplate/pkg/clients/http"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/favourite"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user/rating"
	userRepo "github.com/ralstan-vaz/go-boilerplate/pkg/user/repo"
)

// PackageDeps acts as a dependency container for multiple packages in this project
type PackageDeps struct {
	conf          *config.Config
	dbInstances   *db.DBInstances
	grpcCons      *grpcPkg.GrpcConnections
	httpRequester *httpPkg.Request
}

// NewPackageDeps creates a new instance of PackageDeps
func NewPackageDeps(conf *config.Config, dbInstances *db.DBInstances, grpcCons *grpcPkg.GrpcConnections) *PackageDeps {
	pkgDeps := &PackageDeps{conf: conf, dbInstances: dbInstances, grpcCons: grpcCons}
	pkgDeps.httpRequester = httpPkg.NewRequest()
	return pkgDeps
}

// NewUserPkg Creates an instance of UserPkg by fulfilling all its dependencies
func (p *PackageDeps) NewUserPkg() *user.UserPkg {
	userRepo := userRepo.NewUserRepo(p.conf, p.dbInstances)
	userRating := rating.NewRating(p.conf, p.httpRequester)
	userFavourite := favourite.NewFavourite(p.conf, p.grpcCons)
	return user.NewUserPkg(p.conf, userRepo, userRating, userFavourite)
}
