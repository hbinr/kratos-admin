package biz

import (
	"context"
	v1 "kratos-admin/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// UserDO  领域对象（ Domain Object， DO），微服务运行时核心业务对象的载体， DO 一般包括实体或值对象。
type UserDO struct {
	Id        int64
	Age       int32
	UserId    int64
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
	CreateUser(context.Context, *UserDO) (int64, error)
	UpdateUser(context.Context, *UserDO) error
	DeleteUser(context.Context, int64) error
	SelectUserByID(ctx context.Context, id int64) (*UserDO, error)
	SelectUserByUid(ctx context.Context, userId int64) (*UserDO, error)
	ListUser(ctx context.Context, pageNum, pageSize uint32) ([]*UserDO, error)
	VerifyPassword(context.Context, *UserDO) (bool, error)
	SelectUserByEmail(context.Context, string) (*UserDO, error)
	SelectUserByName(context.Context, string) (*UserDO, error)
}

func NewUserBiz(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, user *UserDO) (int64, error) {
	if err := uc.Validate(ctx, user); err != nil {
		return 0, err
	}
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsecase) Update(ctx context.Context, user *UserDO) error {
	_, err := uc.repo.SelectUserByID(ctx, user.Id)
	if err != nil {
		return err
	}

	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUsecase) Delete(ctx context.Context, userId int64) error {
	return uc.repo.DeleteUser(ctx, userId)
}

func (uc *UserUsecase) Get(ctx context.Context, id int64) (*UserDO, error) {
	return uc.repo.SelectUserByID(ctx, id)
}

func (uc *UserUsecase) GetByUID(ctx context.Context, userId int64) (*UserDO, error) {
	return uc.repo.SelectUserByUid(ctx, userId)
}

func (uc *UserUsecase) List(ctx context.Context, pageNum, pageSize uint32) ([]*UserDO, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc UserUsecase) Validate(ctx context.Context, user *UserDO) (err error) {
	var userDO *UserDO
	userDO, err = uc.repo.SelectUserByName(ctx, user.UserName)
	if err != nil && !v1.IsUserNotFound(err) {
		return err
	}

	if userDO != nil {
		return v1.ErrorUserHasExist("%s has exist", user.UserName)
	}

	userDO, err = uc.repo.SelectUserByEmail(ctx, user.Email)
	if err != nil && !v1.IsUserNotFound(err) {
		return err
	}

	if userDO != nil {
		return v1.ErrorEmailHasExist("%s has exist", user.Email)
	}

	return nil
}
