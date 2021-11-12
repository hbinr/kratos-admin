package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// UserDO  领域对象（ Domain Object， DO），微服务运行时核心业务对象的载体， DO 一般包括实体或值对象。
type UserDO struct {
	ID        uint
	Age       uint32
	UserId    uint32
	UserName  string
	Password  string
	Email     string
	Phone     string
	RoleName  string
	CreatedAt string
	UpdatedAt string
}

// UserUsecase 调 repo 接口
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

type UserRepo interface {
	CreateUser(context.Context, *UserDO) (uint32, error)
	UpdateUser(context.Context, *UserDO) (*UserDO, error)
	DeleteUser(context.Context, uint32) error
	GetUser(ctx context.Context, userId uint32) (*UserDO, error)
	ListUser(ctx context.Context, pageNum, pageSize uint32) ([]*UserDO, error)
	VerifyPassword(context.Context, *UserDO) (bool, error)
}

func NewUserBiz(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, user *UserDO) (uint32, error) {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsecase) Update(ctx context.Context, user *UserDO) (*UserDO, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUsecase) Delete(ctx context.Context, userId uint32) error {
	return uc.repo.DeleteUser(ctx, userId)
}

func (uc *UserUsecase) Get(ctx context.Context, userId uint32) (*UserDO, error) {
	return uc.repo.GetUser(ctx, userId)
}

func (uc *UserUsecase) List(ctx context.Context, pageNum, pageSize uint32) ([]*UserDO, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}
