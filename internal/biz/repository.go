package biz

import "context"

type UserRepo interface {
	CreateUser(context.Context, *UserDO) (int64, error)
	UpdateUser(context.Context, *UserDO) error
	DeleteUser(context.Context, int64) error
	SelectUserByID(ctx context.Context, ID int64) (*UserDO, error)
	SelectUserByUid(ctx context.Context, userID int64) (*UserDO, error)
	ListUser(ctx context.Context, pageNum, pageSize int) ([]*UserDO, error)
	VerifyPassword(context.Context, *UserDO) (bool, error)
	SelectUserByEmail(context.Context, string) (*UserDO, error)
	SelectUserByName(context.Context, string) (*UserDO, error)
}
