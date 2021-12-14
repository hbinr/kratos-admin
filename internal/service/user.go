package service

import (
	"context"
	"kratos-admin/internal/biz"

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

	userBiz *biz.UserUsecase
	log     *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {

	return &UserService{userBiz: uc, log: log.NewHelper(logger)}
}

func (us *UserService) Register(ctx context.Context, req *pb.RegisterReq) (reply *pb.RegisterReply, err error) {
	var (
		userDO biz.UserDO
		userID uint32
	)

	if err = copier.Copy(&userDO, req); err != nil {
		err = errors.Wrap(err, "service: Register data copy failed")
		return
	}

	if userID, err = us.userBiz.Create(ctx, &userDO); err != nil {
		return
	}

	reply = new(pb.RegisterReply)
	reply.UserId = userID
	return
}

func (us *UserService) Login(ctx context.Context, req *pb.LoginReq) (reply *pb.LoginReply, err error) {
	return
}

func (us *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (reply *pb.UpdateUserReply, err error) {
	var (
		userDO biz.UserDO
	)

	if err = copier.Copy(&userDO, req); err != nil {
		return
	}

	if err = us.userBiz.Update(ctx, &userDO); err != nil {
		return
	}

	userRes, err := us.userBiz.Get(ctx, uint(req.Id))
	if err != nil {
		return
	}

	reply = new(pb.UpdateUserReply)
	if err = copier.Copy(&reply, userRes); err != nil {
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
	if userRes, err = us.userBiz.GetByUID(ctx, req.GetUserId()); err != nil {
		return nil, err
	}

	reply = new(pb.GetUserReply)
	if err = copier.Copy(&reply, userRes); err != nil {
		return
	}

	return
}

func (us *UserService) ListUser(ctx context.Context, req *pb.ListUserReq) (reply *pb.ListUserReply, err error) {
	var (
		userDOList = make([]*biz.UserDO, 0, 10)
		resList    = make([]*pb.ListUserReply_User, 0, 10)
	)
	if req.GetPageNum() == 0 {
		req.PageNum = 1
	}
	if req.GetPageSize() == 0 {
		req.PageSize = 10
	}

	if userDOList, err = us.userBiz.List(ctx, req.PageNum, req.GetPageSize()); err != nil {
		return
	}

	if err = copier.Copy(&resList, userDOList); err != nil {
		return
	}

	reply = new(pb.ListUserReply)
	reply.Users = resList
	return
}
