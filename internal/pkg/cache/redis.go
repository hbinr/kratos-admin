package cache

import (
	"kratos-admin/internal/conf"
	"log"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

// Init 初始化 redis 连接
func Init(conf *conf.Data) *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password, // no password set
		DB:           int(conf.Redis.Db),  // use default db
		PoolSize:     int(conf.Redis.PoolSize),
		MinIdleConns: int(conf.Redis.MinIdleConns),
	})
	if _, err := rdb.Ping().Result(); err != nil {
		log.Fatal(err)
	}

	return rdb
}

// Close 关闭redis clent连接资源
func Close() {
	_ = rdb.Close()
}
