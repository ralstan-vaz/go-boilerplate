package apis

import (
  "github.com/ralstan-vaz/go-boilerplate/pkg/user"
)


// PackageInterface contains methods which return dependencies that are used by the services.
// This aids in making it possible to send dependencies to the packages.
// All the subpackages will use this to resolve dependencies.
// An alternative approach would be to have `PackageInterface` in every sub package of `apis`
// But that would cause redundant interfaces with no benefit
// This interface may get larger as more packages(users etc...) come in but simplifies implementation.
type PackageInterface interface {
	NewUserPkg() *user.UserPkg
}
