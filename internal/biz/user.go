package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// UserDO  领域对象（ Domain Object， DO），微服务运行时核心业务对象的载体， DO 一般包括实体或值对象。
type UserDO struct {
	Id        uint
	Age       uint8
	UserId    string
	UserName  string
	Password  string
	Email     string
	Phone     string
	RoleName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// UserBiz
// 调 repo 接口
type UserBiz struct {
	repo UserRepo
	log  *log.Helper
}

type UserRepo interface {
	CreateUser(context.Context, *UserDO) (userId string, err error)
	UpdateUser(context.Context, *UserDO) (*UserDO, error)
	DeleteUser(context.Context, string) error
	GetUser(ctx context.Context, userId string) (*UserDO, error)
	ListUser(ctx context.Context, pageNum, pageSize int64) ([]*UserDO, error)
	VerifyPassword(context.Context, *UserDO) (bool, error)
}

func NewUserBiz(repo UserRepo, logger log.Logger) *UserBiz {
	return &UserBiz{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserBiz) Create(ctx context.Context, user *UserDO) (userId string, err error) {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserBiz) Update(ctx context.Context, user *UserDO) (*UserDO, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserBiz) Delete(ctx context.Context, userId string) error {
	return uc.repo.DeleteUser(ctx, userId)
}

func (uc *UserBiz) Get(ctx context.Context, userId string) (*UserDO, error) {
	return uc.repo.GetUser(ctx, userId)
}

func (uc *UserBiz) List(ctx context.Context, pageNum, pageSize int64) ([]*UserDO, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}
