package svc

import (
	"douyin/apps/rpc/chat/internal/config"

	"github.com/go-redis/redis"
)

type ServiceContext struct {
	Config  config.Config
	DbRedis *redis.Client
}

// 声明一个全局的redisDb变量
var redisDb *redis.Client

func NewServiceContext(c config.Config) *ServiceContext {

	// 缓存连接
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // redis地址
		Password: "123456",         // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})
	// 检测是否连接成功
	_, err := redisDb.Ping().Result()
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:  c,
		DbRedis: redisDb,
	}
}
