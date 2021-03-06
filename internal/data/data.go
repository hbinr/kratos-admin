package data

import (
	"kratos-admin/internal/conf"
	"kratos-admin/internal/data/dao"
	"kratos-admin/internal/pkg/cache"
	"kratos-admin/internal/pkg/database"

	"github.com/go-redis/redis"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var (
	ProviderSet = wire.NewSet(NewData, NewUserRepo)
)

// Data .
type Data struct {
	sqlClient *dao.Query
	rdb       *redis.Client
}

func NewData(conf *conf.Data) *Data {
	data := new(Data)
	data.sqlClient = dao.Use(database.Init(conf))
	data.rdb = cache.Init(conf)
	return data
}
