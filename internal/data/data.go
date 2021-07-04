package data

import (
	"kratos-admin/internal/conf"
	"kratos-admin/internal/pkg/cache"
	"kratos-admin/internal/pkg/database"

	"github.com/go-redis/redis"

	"gorm.io/gorm"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var (
	ProviderSet = wire.NewSet(NewData, NewUserRepo)
	entities    = []interface{}{&UserPO{}}
)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewData(conf *conf.Data) *Data {
	data := new(Data)
	data.db = database.Init(conf)
	data.rdb = cache.Init(conf)
	if err := data.AutoMigrate(); err != nil {
		panic(err)
	}
	return data
}

func (d *Data) AutoMigrate() error {
	return d.db.AutoMigrate(entities...)
}
