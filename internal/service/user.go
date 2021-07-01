package service

import (
	"context"
	"kratos-admin/internal/biz"
	"kratos-admin/pkg/util/timex"

	"github.com/pkg/errors"

	"github.com/jinzhu/copier"

	"github.com/go-kratos/kratos/v2/log"

	pb "kratos-admin/api/user/service/v1"
)

//	req -> po
// 	调 biz 接口
//	响应 -> reply
type UserService struct {
	pb.UnimplementedUserServer

	userBiz *biz.UserBiz
	log     *log.Helper
}

func NewUserService(uc *biz.UserBiz, logger log.Logger) *UserService {

	return &UserService{userBiz: uc, log: log.NewHelper(logger)}
}

func (us *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
	var userDO biz.UserDO

	if err := copier.Copy(&userDO, req); err != nil {
		return nil, errors.Wrap(err, "service: copier.Copy(&userDO, req) failed")
	}

	userID, err := us.userBiz.Create(ctx, &userDO)

	if err != nil {
		return nil, errors.WithMessagef(err, "service: Create User failed, userName: [%s]", req.UserName)
	}

	return &pb.CreateUserReply{UserId: userID}, nil
}

func (us *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserReply, error) {
	var userDO biz.UserDO
	if err := copier.Copy(&userDO, req); err != nil {
		return nil, err
	}

	result, err := us.userBiz.Update(ctx, &userDO)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserReply{
		Id:       int64(result.Id),
		UserName: result.UserName,
	}, nil
}

func (us *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}

func (us *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	var (
		err       error
		userRes   *biz.UserDO
		userReply pb.GetUserReply
	)
	if userRes, err = us.userBiz.Get(ctx, req.GetUserId()); err != nil {
		return nil, errors.Wrap(err, "service: Get user failed")
	}

	if err := copier.Copy(&userReply, userRes); err != nil {
		return nil, errors.Wrap(err, "service: copier.Copy(&userReply, userDO) failed")
	}

	userReply.CreatedAt = timex.DateToString(userRes.CreatedAt)
	userReply.UpdatedAt = timex.DateToString(userRes.UpdatedAt)
	return &userReply, nil
}
func (us *UserService) ListUser(ctx context.Context, req *pb.ListUserReq) (*pb.ListUserReply, error) {
	userDOList, err := us.userBiz.List(ctx, req.GetPageNum(), req.GetPageSize())

	if err != nil {
		return nil, err
	}

	var resList []*pb.ListUserReply_User
	if err = copier.Copy(&resList, userDOList); err != nil {

		return nil, err
	}
	return &pb.ListUserReply{
		Users: resList,
	}, nil
}
