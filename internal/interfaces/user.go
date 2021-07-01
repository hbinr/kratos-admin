package interfaces

import (
	"context"
	pb "kratos-admin/api/user/service/v1"
	"kratos-admin/internal/pkg/constant/e"
	"kratos-admin/internal/pkg/ginx"
	"kratos-admin/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/gin-gonic/gin"
)

type UserUseCase struct {
	userService *service.UserService
	log         *log.Helper
}

func NewUserUseCase(us *service.UserService, logger log.Logger) *UserUseCase {
	return &UserUseCase{userService: us, log: log.NewHelper(logger)}
}

func (u *UserUseCase) SayHi(c *gin.Context) {
	ginx.RespSuccess(c, "Hello World")
}

func (u *UserUseCase) Signup(c *gin.Context) {
	var (
		req pb.CreateUserReq
	)
	if err := c.BindJSON(&req); err != nil {
		ginx.RespError(c, e.CodeInvalidParams)
		return
	}
	res, err := u.userService.CreateUser(context.Background(), &req)

	if err != nil {
		ginx.RespError(c, e.CodeInternalError)
		return
	}
	ginx.RespSuccess(c, res)
}

func (u *UserUseCase) Get(c *gin.Context) {
	var (
		req pb.GetUserReq
	)
	if err := c.BindJSON(&req); err != nil {
		ginx.RespError(c, e.CodeInvalidParams)
		return
	}
	res, err := u.userService.GetUser(context.Background(), &req)

	switch err {
	case nil:
		ginx.RespSuccess(c, res)
	case e.ErrNotFound:
		ginx.RespError(c, e.CodeNotFound)
	default:
		ginx.RespError(c, e.CodeInternalError)
	}
}

func (u *UserUseCase) List(c *gin.Context) {
	var (
		req pb.ListUserReq
	)
	if err := c.BindJSON(&req); err != nil {
		ginx.RespError(c, e.CodeInvalidParams)
		return
	}
	res, err := u.userService.ListUser(context.Background(), &req)

	if err != nil {
		ginx.RespError(c, e.CodeInternalError)
		return
	}
	ginx.RespSuccess(c, res)
}

func (u *UserUseCase) Delte(c *gin.Context) {
	var (
		req pb.DeleteUserReq
	)
	if err := c.BindJSON(&req); err != nil {
		ginx.RespError(c, e.CodeInvalidParams)
		return
	}
	res, err := u.userService.DeleteUser(context.Background(), &req)

	if err != nil {
		ginx.RespError(c, e.CodeInternalError)
		return
	}
	ginx.RespSuccess(c, res)
}
