package biz

import (
	"context"
	v1 "kratos-admin/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

// UserUsecase 调 repo 接口
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
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

func (uc *UserUsecase) Delete(ctx context.Context, userID int64) error {
	return uc.repo.DeleteUser(ctx, userID)
}

func (uc *UserUsecase) Get(ctx context.Context, ID int64) (*UserDO, error) {
	return uc.repo.SelectUserByID(ctx, ID)
}

func (uc *UserUsecase) GetByUID(ctx context.Context, userID int64) (*UserDO, error) {
	return uc.repo.SelectUserByUid(ctx, userID)
}

func (uc *UserUsecase) List(ctx context.Context, pageNum, pageSize uint32) ([]*UserDO, error) {
	return uc.repo.ListUser(ctx, int(pageNum), int(pageSize))
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
