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
	UserId   string `gorm:"not null;size:64;index:idx_user_id;"`
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

func (u *userRepo) CreateUser(ctx context.Context, do *biz.UserDO) (userId string, err error) {
	var po = UserPO{}
	if err = copier.Copy(&po, do); err != nil {
		return
	}

	po.Password, err = hashx.HashPassword(do.Password)
	po.UserId = uuidx.GenID()

	if err = u.data.db.WithContext(ctx).Create(&po).Error; err != nil {
		err = errors.Wrap(err, "data: Create user failed")
		return
	}
	userId = po.UserId
	return
}

func (u *userRepo) UpdateUser(ctx context.Context, do *biz.UserDO) (*biz.UserDO, error) {
	var po = UserPO{}
	if err := copier.Copy(&po, do); err != nil {
		return nil, err
	}
	err := u.data.db.WithContext(ctx).
		Where("user_id = ? ", po.UserId).
		Updates(&po).Error

	return &biz.UserDO{
		Id:       po.ID,
		UserId:   po.UserId,
		UserName: po.UserName,
	}, err
}

func (u *userRepo) DeleteUser(ctx context.Context, userId string) error {
	result := u.data.db.WithContext(ctx).Where("user_id = ?", userId).Delete(&UserPO{})
	if result.Error != nil {
		return errors.Wrapf(result.Error, "data: deleted user failed, userID[%s]", userId)
	}

	if result.RowsAffected <= 0 {
		return e.ErrUserHasDeleted
	}
	return nil
}

func (u *userRepo) GetUser(ctx context.Context, userId string) (*biz.UserDO, error) {
	var (
		userPO UserPO
		do     biz.UserDO
	)
	result := u.data.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Find(&userPO)

	if result.RowsAffected == 0 {
		return nil, e.ErrNotFound
	}

	switch result.Error {
	case nil:

		if err := copier.Copy(&do, userPO); err != nil {
			return nil, err
		}

		return &do, nil
	case gorm.ErrRecordNotFound:
		return nil, e.ErrNotFound
	default:
		return nil, result.Error
	}
}

func (u *userRepo) ListUser(ctx context.Context, pageNum, pageSize int64) (doList []*biz.UserDO, err error) {
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

func (u *userRepo) VerifyPassword(ctx context.Context, do *biz.UserDO) (bool, error) {
	var po UserPO
	res := u.data.db.WithContext(ctx).Where("user_name = ?", do.UserName).Find(&po)
	if res.Error != nil {
		return false, res.Error
	}

	return hashx.CheckPasswordHash(do.Password, po.Password), nil

}
