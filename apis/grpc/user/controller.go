package user

import (
	"context"

	"github.com/mitchellh/mapstructure"
	pb "github.com/ralstan-vaz/go-boilerplate/apis/grpc/generated/user"
	"github.com/ralstan-vaz/go-boilerplate/pkg/user"
	userRepo "github.com/ralstan-vaz/go-boilerplate/pkg/user/repo"
)

type PackageInterface interface {
	NewUserPkg() *user.UserPkg
}

func NewUserService(pkg PackageInterface) *UserService {

	return &UserService{pkg: pkg}
}

type UserService struct {
	pkg PackageInterface
}

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

func (u *UserService) Insert(ctx context.Context, req *pb.User) (res *pb.User, err error) {

	userPkg := u.pkg.NewUserPkg()
	userReq := userRepo.User{}
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

func (u *UserService) GetWithRating(ctx context.Context, req *pb.UserGetRequest) (res *pb.User, err error) {

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

	users, err := userPkg.GetWithRating(userReq.ID)
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
