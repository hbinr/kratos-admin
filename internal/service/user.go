package service

import (
	"context"
	"fmt"
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

func (us *UserService) Register(ctx context.Context, req *pb.RegisterReq) (reply *pb.RegisterReply, err error) {
	var (
		userDO biz.UserDO
		userId uint32
	)

	if err = copier.Copy(&userDO, req); err != nil {
		return nil, errors.Wrap(err, "service: copier.Copy(&userDO, req) failed")
	}

	if userId, err = us.userBiz.Create(ctx, &userDO); err != nil {
		return nil, errors.WithMessagef(err, "service: Create User failed, userName: [%s]", req.UserName)
	}
	reply = new(pb.RegisterReply)
	reply.UserId = userId
	return
}

func (us *UserService) Login(ctx context.Context, req *pb.LoginReq) (reply *pb.LoginReply, err error) {
	return
}

func (us *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (reply *pb.UpdateUserReply, err error) {
	var userDO biz.UserDO

	if err = copier.Copy(&userDO, req); err != nil {
		return
	}

	result, err := us.userBiz.Update(ctx, &userDO)

	if err != nil {
		return
	}

	reply = new(pb.UpdateUserReply)
	if err = copier.Copy(&reply, result); err != nil {
		return
	}

	return
}

func (us *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (reply *pb.DeleteUserReply, err error) {
	reply = &pb.DeleteUserReply{Ok: true}
	if err = us.userBiz.Delete(ctx, req.UserId); err != nil {
		reply.Ok = false
		return
	}

	return
}

func (us *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (reply *pb.GetUserReply, err error) {
	var (
		userRes *biz.UserDO
	)
	if userRes, err = us.userBiz.Get(ctx, req.GetUserId()); err != nil {
		return nil, err
	}

	reply = new(pb.GetUserReply)
	if err = copier.Copy(&reply, userRes); err != nil {
		return nil, errors.Wrap(err, "service: GetUser copier.Copy(&userReply, userDO) failed")
	}

	reply.CreatedAt = timex.DateToString(userRes.CreatedAt)
	reply.UpdatedAt = timex.DateToString(userRes.UpdatedAt)
	return
}

func (us *UserService) ListUser(ctx context.Context, req *pb.ListUserReq) (reply *pb.ListUserReply, err error) {
	var (
		userDOList []*biz.UserDO
		resList    []*pb.ListUserReply_User
	)
	fmt.Println("------")

	if userDOList, err = us.userBiz.List(ctx, req.GetPageNum(), req.GetPageSize()); err != nil {
		return nil, err
	}

	if err = copier.Copy(&resList, userDOList); err != nil {
		return nil, errors.Wrap(err, "service: ListUser copier.Copy(&resList, userDOList) failed")
	}

	reply = new(pb.ListUserReply)
	reply.Users = resList
	return
}
