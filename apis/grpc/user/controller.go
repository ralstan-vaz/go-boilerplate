package user

import (
	"context"

	"github.com/mitchellh/mapstructure"
	pb "github.com/ralstan-vaz/go-boilerplate/apis/grpc/generated/user"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user"
)

// PackageInterface contains methods which return dependencies that are used by the services.
type PackageInterface interface {
	NewUserPkg() *user.UserPkg
}

// NewUserService Create a new instance of a UserService with the given dependencies.
func NewUserService(pkg PackageInterface) *UserService {
	return &UserService{pkg: pkg}
}

// UserService contains the methods required to perfom operation's on users
type UserService struct {
	pkg PackageInterface
}

// GetAll gets all users
func (u *UserService) GetAll(ctx context.Context, req *pb.UserGetRequest) (res *pb.Users, err error) {

	userPkg := u.pkg.NewUserPkg()
	users, err := userPkg.GetAll()
	if err != nil {
		return nil, err
	}

	res = &pb.Users{}
	err = mapstructure.Decode(users, &res.Users)
	if err != nil {
		panic(err)
	}

	return res, nil
}

// GetOne gets one users
func (u *UserService) GetOne(ctx context.Context, req *pb.UserGetRequest) (res *pb.User, err error) {

	userPkg := u.pkg.NewUserPkg()
	if err != nil {
		return nil, err
	}

	userReq := user.User{}
	// Need to decode to user.User since User is an embedded struct
	err = mapstructure.Decode(req, &userReq)
	if err != nil {
		panic(err)
	}

	users, err := userPkg.GetOne(userReq.ID)
	if err != nil {
		return nil, err
	}

	res = &pb.User{}
	err = mapstructure.Decode(users, &res)
	if err != nil {
		panic(err)
	}

	return res, nil
}

// Insert stores a user in the datastore
func (u *UserService) Insert(ctx context.Context, req *pb.User) (res *pb.User, err error) {

	userPkg := u.pkg.NewUserPkg()
	userReq := user.User{}
	// Need to decode to user.User since User is an embedded struct
	err = mapstructure.Decode(req, &userReq)
	if err != nil {
		panic(err)
	}

	err = userPkg.Insert(userReq)
	if err != nil {
		return nil, err
	}
	res = nil

	return res, nil
}

// GetWithInfo gets a user from the database along with rating and favourites
func (u *UserService) GetWithInfo(ctx context.Context, req *pb.UserGetRequest) (res *pb.User, err error) {

	userPkg := u.pkg.NewUserPkg()
	if err != nil {
		return nil, err
	}

	userReq := user.User{}
	// Need to decode to user.User since User is an embedded struct
	err = mapstructure.Decode(req, &userReq)
	if err != nil {
		panic(err)
	}

	users, err := userPkg.GetWithInfo(userReq.ID)
	if err != nil {
		return nil, err
	}

	res = &pb.User{}
	err = mapstructure.Decode(users, &res)
	if err != nil {
		panic(err)
	}

	return res, nil
}
