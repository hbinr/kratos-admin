package data

import (
	"database/sql"
	"kratos-admin/internal/conf"
	"os"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var (
	ProviderSet = wire.NewSet(NewData, NewUserRepo)
	sqlDB       *sql.DB
)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData 创建数据库连接, 基于 gorm + mysql
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	logx := log.NewHelper(log.With(logger, "module", "user-service/data"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), gormConfig(conf.Database.LogMode))
	if err != nil {
		logx.Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}

	if sqlDB, err = db.DB(); err != nil {
		logx.Error("db.DB() failed", err)
		return nil, nil, err
	}
	gormDBTables(db, logx)
	sqlDB.SetMaxIdleConns(int(conf.Database.MaxIdleConns))
	sqlDB.SetMaxOpenConns(int(conf.Database.MaxOpenConns))

	d := &Data{
		db: db,
	}
	return d, func() {
		//	gorm db 会自动释放资源, 无需处理
	}, nil
}

// gormDBTables 注册数据库表专用
func gormDBTables(db *gorm.DB, log *log.Helper) {
	err := db.AutoMigrate(&UserPO{})
	if err != nil {
		log.Errorf("data: gorm AutoMigrate tables failed, err:%v+", err)
		os.Exit(0)
	}
	log.Info("data: register table success")
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(mod bool) (c *gorm.Config) {
	if mod {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 表名不加复数形式，false默认加
			},
		}
	}
	return
}
