package data

import (
	"context"
	"kratos-admin/internal/biz"
	"kratos-admin/internal/pkg/constant/e"
	"kratos-admin/pkg/util/hashx"
	"kratos-admin/pkg/util/pagination"
	"kratos-admin/pkg/util/uuidx"

	"github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

// UserPO UserPO 持久化对象，与数据库结构一一映射，它是数据持久化过程中的数据载体。
type UserPO struct {
	gorm.Model
	UserId   uint32 `gorm:"not null;index:idx_user_id;"`
	age      uint32 `gorm:"not null;"`
	UserName string `gorm:"not null;size:32;;index:idx_user_name;"`
	Password string `gorm:"not null;size:64;"`
	Email    string `gorm:"not null;size:128;unique;"`
	Phone    string `gorm:"not null;size:11;"`
	RoleName string `gorm:"not null;size:10;"`
}

// 入参 do-> po
// 响应 po -> do
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

func (u *UserPO) TableName() string {
	return "user"
}

func (u *userRepo) CreateUser(ctx context.Context, do *biz.UserDO) (userID uint32, err error) {
	var (
		po = UserPO{}
	)
	if err = copier.Copy(&po, do); err != nil {
		return
	}

	if po.Password, err = hashx.HashPassword(do.Password); err != nil {
		return
	}

	po.UserId = uuidx.GenID()

	if err = u.data.db.WithContext(ctx).Create(&po).Error; err != nil {
		err = errors.Wrap(err, "data: Create user failed")
		return
	}

	userID = po.UserId
	return
}

func (u *userRepo) UpdateUser(ctx context.Context, do *biz.UserDO) (res *biz.UserDO, err error) {
	var po = UserPO{}

	if err = copier.Copy(&po, do); err != nil {
		return
	}

	err = u.data.db.WithContext(ctx).
		Where("user_id = ? ", po.UserId).
		Updates(&po).Error

	res = new(biz.UserDO)
	if err = copier.Copy(&res, po); err != nil {
		return
	}

	return
}

func (u *userRepo) DeleteUser(ctx context.Context, userId uint32) error {
	result := u.data.db.WithContext(ctx).Where("user_id = ?", userId).Delete(&UserPO{})
	if result.Error != nil {
		return errors.Wrapf(result.Error, "data: deleted user failed, userID[%d]", userId)
	}

	if result.RowsAffected <= 0 {
		return e.ErrUserHasDeleted
	}
	return nil
}

func (u *userRepo) GetUser(ctx context.Context, userId uint32) (res *biz.UserDO, err error) {
	var (
		userPO UserPO
	)
	err = u.data.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Find(&userPO).Error
	res = new(biz.UserDO)
	switch err {
	case nil:
		if err = copier.Copy(&res, userPO); err != nil {
			return
		}

		return
	case gorm.ErrRecordNotFound:
		err = e.ErrNotFound
		return
	default:
		return
	}
}

func (u *userRepo) ListUser(ctx context.Context, pageNum, pageSize uint32) (doList []*biz.UserDO, err error) {
	var poList []UserPO
	result := u.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&poList)

	if result.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	if result.Error != nil {
		err = result.Error
		return
	}

	doList = make([]*biz.UserDO, 0)
	for _, po := range poList {
		doList = append(doList, &biz.UserDO{
			UserId:    po.UserId,
			UserName:  po.UserName,
			Password:  po.Password,
			Email:     po.Email,
			Phone:     po.Phone,
			RoleName:  po.RoleName,
			CreatedAt: po.CreatedAt,
			UpdatedAt: po.UpdatedAt,
		})
	}

	return
}

func (u *userRepo) VerifyPassword(ctx context.Context, do *biz.UserDO) (isCorrect bool, err error) {
	var po UserPO
	err = u.data.db.WithContext(ctx).Where("user_name = ?", do.UserName).Find(&po).Error

	switch err {
	case nil:
		isCorrect = hashx.CheckPasswordHash(do.Password, po.Password)
		return
	case gorm.ErrRecordNotFound:
		isCorrect = false
		err = e.ErrNotFound
	default:
		isCorrect = false
	}

	return

}
