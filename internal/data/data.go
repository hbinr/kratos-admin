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
)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewData(conf *conf.Data) *Data {
	return &Data{
		db:  database.Init(conf),
		rdb: cache.Init(conf),
	}
}
