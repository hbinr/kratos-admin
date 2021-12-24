package data

import (
	"context"
	v1 "kratos-admin/api/user/service/v1"
	"kratos-admin/internal/biz"
	"kratos-admin/internal/data/entity"
	"kratos-admin/pkg/util/hashx"
	"kratos-admin/pkg/util/pagination"
	"kratos-admin/pkg/util/timex"
	"kratos-admin/pkg/util/uuidx"
	"time"

	"github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 入参 do -> po
// 响应 po -> do

// NewUserPO UserPO  持久化对象，与数据库结构一一映射，它是数据持久化过程中的数据载体。
func (u userRepo) NewUserPO(do *biz.UserDO) *entity.User {
	return &entity.User{
		ID:       do.Id,
		UserID:   do.UserId,
		Age:      do.Age,
		UserName: do.UserName,
		Password: do.Password,
		Email:    do.Email,
		Phone:    do.Phone,
		RoleName: do.RoleName,
	}
}

func (u userRepo) NewUserDO(po *entity.User) *biz.UserDO {
	return &biz.UserDO{
		Id:        po.ID,
		Age:       po.Age,
		UserId:    po.UserID,
		UserName:  po.UserName,
		Password:  po.Password,
		Email:     po.Email,
		Phone:     po.Phone,
		RoleName:  po.RoleName,
		CreatedAt: timex.DateToString(po.CreatedAt),
		UpdatedAt: timex.DateToString(po.UpdatedAt),
	}
}

func (u *userRepo) CreateUser(ctx context.Context, do *biz.UserDO) (userID int64, err error) {
	po := u.NewUserPO(do)
	po.Password = hashx.MD5String(do.Password)
	po.UserID = int64(uuidx.GenID())
	if err = u.data.sqlClient.User.WithContext(ctx).Create(po); err != nil {
		return
	}

	userID = po.UserID
	return
}

func (u *userRepo) UpdateUser(ctx context.Context, do *biz.UserDO) error {
	user := u.data.sqlClient.User

	po := u.NewUserPO(do)
	po.UpdatedAt = time.Now()
	_, err := user.WithContext(ctx).Where(user.ID.Eq(po.ID)).Updates(po)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepo) DeleteUser(ctx context.Context, userID int64) error {
	user := u.data.sqlClient.User
	if _, err := user.WithContext(ctx).Where(user.UserID.Eq(userID)).Delete(); err != nil {
		return errors.Wrapf(err, "data: deleted user failed, userID = %d", userID)
	}

	return nil
}

func (u *userRepo) SelectUserByUid(ctx context.Context, userID int64) (do *biz.UserDO, err error) {
	user := u.data.sqlClient.User
	res, err := user.WithContext(ctx).Where(user.UserID.Eq(userID)).Take()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = v1.ErrorUserNotFound("data: userID = %d", userID)
		}
		return
	}

	do = u.NewUserDO(res)
	return
}

func (u *userRepo) SelectUserByID(ctx context.Context, id int64) (do *biz.UserDO, err error) {
	user := u.data.sqlClient.User
	res, err := user.WithContext(ctx).Where(user.ID.Eq(id)).Take()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = v1.ErrorUserNotFound("data: id = %d", id)
		}
		return
	}

	do = u.NewUserDO(res)
	return
}
func (u *userRepo) ListUser(ctx context.Context, pageNum, pageSize int) (doList []*biz.UserDO, err error) {
	user := u.data.sqlClient.User
	poList, err := user.WithContext(ctx).Limit(pageSize).Offset(pagination.GetPageOffset(pageNum, pageSize)).Find()

	if err != nil {
		return
	}

	doList = make([]*biz.UserDO, 0, len(poList))
	for _, po := range poList {
		doList = append(doList, u.NewUserDO(po))
	}
	return
}

func (u *userRepo) VerifyPassword(ctx context.Context, do *biz.UserDO) (isCorrect bool, err error) {
	user := u.data.sqlClient.User
	var po *entity.User
	if po, err = user.WithContext(ctx).Where(user.UserName.Eq(do.UserName)).Take(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = v1.ErrorUserNotFound("data: user_name = %s", do.UserName)
		}
		isCorrect = false
		return
	}

	if po.Password == hashx.MD5String(do.Password) {
		isCorrect = true
	}

	return
}

func (u *userRepo) SelectUserByEmail(ctx context.Context, email string) (do *biz.UserDO, err error) {
	user := u.data.sqlClient.User
	var po *entity.User
	if po, err = user.WithContext(ctx).Where(user.Email.Eq(email)).Take(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = v1.ErrorUserNotFound("data: email = %s", email)
		}
		return
	}

	do = u.NewUserDO(po)
	return
}

func (u *userRepo) SelectUserByName(ctx context.Context, userName string) (do *biz.UserDO, err error) {
	user := u.data.sqlClient.User
	var po *entity.User
	if po, err = user.WithContext(ctx).Where(user.UserName.Eq(userName)).Take(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = v1.ErrorUserNotFound("data: user_name = %s", userName)
		}
		return
	}
	do = u.NewUserDO(po)
	return
}
