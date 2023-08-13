package pkg

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  // 主键（用户id）
	Username string // 用户名
	Password string // 密码

	FansCount       int64  `gorm:"default:0"`                                              // 粉丝数
	FollowerCount   int64  `gorm:"default:0"`                                              // 关注数
	Avatar          string `gorm:"default:https://img1.imgtp.com/2023/08/10/L3PtLRbx.png"` // 用户头像
	BackgroundImage string `gorm:"default:https://img1.imgtp.com/2023/08/10/FFUYtwsH.jpg"` // 顶部头图
	Signature       string `gorm:"default:这是一条个性签名~"`                                      // 个人简介
	TotalFavorited  string `gorm:"default:0"`                                              // 获赞数量
	WorkCount       int64  `gorm:"default:0"`                                              // 作品数
	FavoriteCount   int64  `gorm:"default:0"`                                              // 喜欢数量
}

func MysqlInit() (*gorm.DB, error) {
	// 数据库连接
	username := "searchdata" // 使用者名字 如root
	password := "123456Xk@"
	host := "8.130.24.85"
	port := 3306
	dbname := "douyin" // 数据库名字
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("数据库连接失败")
	} else {
		return db, nil
	}
}

// 声明一个全局的redisDb变量
var redisDb *redis.Client

func RedisInit() (*redis.Client, error) {
	// 缓存连接
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "8.130.24.85:6379", // redis地址
		Password: "",                 // redis密码，没有则留空
		DB:       0,                  // 默认数据库，默认是0
	})
	// 检测是否连接成功
	_, err := redisDb.Ping().Result()
	if err != nil {
		return nil, errors.New("redis连接失败")
	} else {
		return redisDb, nil
	}
}
