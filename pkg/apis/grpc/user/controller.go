package user

import (
	"context"

	"github.com/ralstan-vaz/go-boilerplate/pkg/apis"
	proto "github.com/ralstan-vaz/go-boilerplate/pkg/apis/grpc/generated/user"
	utils "github.com/ralstan-vaz/go-boilerplate/pkg/apis/grpc/utils"

	"github.com/mitchellh/mapstructure"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user"
)

// NewUserService Create a new instance of a UserService with the given dependencies.
func NewUserService(pkg apis.PackageInterface) *UserService {
	return &UserService{pkg: pkg}
}

// UserService contains the methods required to perfom operation's on users
type UserService struct {
	pkg apis.PackageInterface
}

// GetAll gets all users
func (u *UserService) GetAll(ctx context.Context, req *proto.UserGetRequest) (res *proto.Users, err error) {

	userPkg := u.pkg.NewUserPkg()
	users, err := userPkg.GetAll()
	if err != nil {
		return nil, utils.HandleError(err)
	}

	res = &proto.Users{}
	err = mapstructure.Decode(users, &res.Users)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return res, nil
}

// GetOne gets one users
func (u *UserService) GetOne(ctx context.Context, req *proto.UserGetRequest) (res *proto.User, err error) {

	userPkg := u.pkg.NewUserPkg()

	userReq := user.User{}
	err = mapstructure.Decode(req, &userReq)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	users, err := userPkg.GetOne(userReq.ID)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	res = &proto.User{}
	err = mapstructure.Decode(users, &res)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return res, nil
}

// Insert stores a user in the datastore
func (u *UserService) Insert(ctx context.Context, req *proto.User) (res *proto.User, err error) {

	userPkg := u.pkg.NewUserPkg()
	userReq := user.User{}
	err = mapstructure.Decode(req, &userReq)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	err = userPkg.Insert(userReq)
	if err != nil {
		return nil, utils.HandleError(err)
	}
	res = nil

	return res, nil
}

// GetWithInfo gets a user from the database along with rating and favourites
func (u *UserService) GetWithInfo(ctx context.Context, req *proto.UserGetRequest) (res *proto.User, err error) {

	userPkg := u.pkg.NewUserPkg()

	userReq := user.User{}
	err = mapstructure.Decode(req, &userReq)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	users, err := userPkg.GetWithInfo(userReq.ID)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	res = &proto.User{}
	err = mapstructure.Decode(users, &res)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return res, nil
}
