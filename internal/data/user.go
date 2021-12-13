package data

import (
	"context"
	v1 "kratos-admin/api/user/service/v1"
	"kratos-admin/internal/biz"
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

// UserPO  持久化对象，与数据库结构一一映射，它是数据持久化过程中的数据载体。
type UserPO struct {
	Id        uint   `gorm:"primarykey"`
	UserId    uint32 `gorm:"not null;index:idx_user_id;"`
	Age       uint32 `gorm:"not null;"`
	UserName  string `gorm:"not null;size:32;;index:idx_user_name;"`
	Password  string `gorm:"not null;size:64;"`
	Email     string `gorm:"not null;size:128;unique;"`
	Phone     string `gorm:"not null;size:11;"`
	RoleName  string `gorm:"not null;size:10;"`
	CreatedAt time.Time
	UpdatedAt time.Time
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

func (po *UserPO) DOFactory(do *biz.UserDO) {
	do.Id = po.Id
	do.UserId = po.UserId
	do.Age = po.Age
	do.UserName = po.UserName
	do.Password = po.Password
	do.Email = po.Email
	do.Phone = po.Phone
	do.RoleName = po.RoleName
	do.CreatedAt = timex.DateToString(po.CreatedAt)
	do.UpdatedAt = timex.DateToString(po.UpdatedAt)

}

func (po *UserPO) POFactory(do *biz.UserDO) {
	po.Id = do.Id
	po.UserId = do.UserId
	po.Age = do.Age
	po.UserName = do.UserName
	po.Password = do.Password
	po.Email = do.Email
	po.Phone = do.Phone
	po.RoleName = do.RoleName
}

func (po *UserPO) TableName() string {
	return "user"
}

func (u *userRepo) CreateUser(ctx context.Context, do *biz.UserDO) (userID uint32, err error) {
	po := new(UserPO)
	po.POFactory(do)

	po.Password = hashx.MD5String(do.Password)
	po.UserId = uuidx.GenID()
	if err = u.data.db.WithContext(ctx).Create(po).Error; err != nil {
		return
	}

	userID = po.UserId
	return
}

func (u *userRepo) UpdateUser(ctx context.Context, do *biz.UserDO) error {
	po := new(UserPO)
	po.POFactory(do)
	po.UpdatedAt = time.Now()

	return u.data.db.WithContext(ctx).Updates(po).Error
}

func (u *userRepo) DeleteUser(ctx context.Context, userId uint32) error {
	result := u.data.db.WithContext(ctx).Where("user_id = ?", userId).Delete(&UserPO{})
	if result.Error != nil {
		return errors.Wrapf(result.Error, "data: deleted user failed, userID[%d]", userId)
	}

	if result.RowsAffected <= 0 {
		// 响应结果中 message 字段值就是 format 参数
		return v1.ErrorUserNotFound("data: user_id = %d", userId)
	}
	return nil
}

func (u *userRepo) SelectUserByUid(ctx context.Context, userId uint32) (do *biz.UserDO, err error) {
	var (
		userPO UserPO
	)
	if err = u.data.db.WithContext(ctx).Where("user_id = ?", userId).Take(&userPO).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = v1.ErrorUserNotFound("data: userId = %d", userId)
		}
		return
	}

	do = new(biz.UserDO)
	userPO.DOFactory(do)
	return
}

func (u *userRepo) SelectUserByID(ctx context.Context, id uint) (do *biz.UserDO, err error) {
	var (
		userPO UserPO
	)
	if err = u.data.db.WithContext(ctx).Where("id = ?", id).First(&userPO).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = v1.ErrorUserNotFound("data: id = %d", id)
		}
		return
	}

	do = new(biz.UserDO)
	userPO.DOFactory(do)
	return
}
func (u *userRepo) ListUser(ctx context.Context, pageNum, pageSize uint32) (doList []*biz.UserDO, err error) {
	var poList []UserPO
	result := u.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&poList)

	if result.Error != nil {
		err = result.Error
		return
	}

	doList = make([]*biz.UserDO, 0, len(poList))
	for _, po := range poList {
		doList = append(doList, &biz.UserDO{
			Id:        po.Id,
			UserId:    po.UserId,
			UserName:  po.UserName,
			Password:  po.Password,
			Email:     po.Email,
			Phone:     po.Phone,
			Age:       po.Age,
			RoleName:  po.RoleName,
			CreatedAt: timex.DateToString(po.CreatedAt),
			UpdatedAt: timex.DateToString(po.CreatedAt),
		})
	}
	return
}

func (u *userRepo) VerifyPassword(ctx context.Context, do *biz.UserDO) (isCorrect bool, err error) {
	var po UserPO
	if err = u.data.db.WithContext(ctx).Where("user_name = ?", do.UserName).Take(&po).Error; err != nil {
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
	var (
		userPO UserPO
	)
	if err = u.data.db.WithContext(ctx).Where("email = ?", email).Take(&userPO).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = v1.ErrorUserNotFound("data: email = %s", email)
		}
		return
	}

	do = new(biz.UserDO)
	userPO.DOFactory(do)
	return
}

func (u *userRepo) SelectUserByName(ctx context.Context, userName string) (do *biz.UserDO, err error) {
	var (
		userPO UserPO
	)
	if err = u.data.db.WithContext(ctx).Where("user_name = ?", userName).Take(&userPO).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = v1.ErrorUserNotFound("data: user_name = %s", userName)
		}
		return
	}

	do = new(biz.UserDO)
	userPO.DOFactory(do)
	return
}
