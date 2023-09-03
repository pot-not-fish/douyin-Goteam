package svc

import (
	"douyin/apps/rpc/video/internal/config"
	"fmt"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
	DbRedis *redis.Client
}

// 声明一个全局的redisDb变量
var redisDb *redis.Client

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库连接
	username := "searchdata" // 使用者名字 如root
	password := "123456"
	host := "127.0.0.1"
	port := 3306
	dbname := "douyin" // 数据库名字
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 缓存连接
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // redis地址
		Password: "123456",         // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})
	// 检测是否连接成功
	_, err = redisDb.Ping().Result()
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:  c,
		DbEngin: db,
		DbRedis: redisDb,
	}
}
