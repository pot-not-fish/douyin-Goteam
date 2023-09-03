package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  // 主键（用户id）
	Username string // 用户名
	Password string // 密码

	FansCount       int64  `gorm:"default:0"`                                     // 粉丝数
	FollowCount     int64  `gorm:"default:0"`                                     // 关注数
	Avatar          string `gorm:"default:http://10.0.2.2/images/avatar.jpg"`     // 用户头像
	BackgroundImage string `gorm:"default:http://10.0.2.2/images/background.jpg"` // 顶部头图
	Signature       string `gorm:"default:这是一条个性签名~"`                             // 个人简介
	TotalFavorited  string `gorm:"default:0"`                                     // 获赞数量
	WorkCount       int64  `gorm:"default:0"`                                     // 作品数
	FavoriteCount   int64  `gorm:"default:0"`                                     // 喜欢数量
}

type Video struct {
	ID            int64  // 主键（视频id）
	UserId        int64  // 用户id
	PlayUrl       string // 视频播放地址
	CoverUrl      string // 视频封面地址
	FavoriteCount int64  `gorm:"default:0"` // 视频点赞总数
	CommentCount  int64  `gorm:"default:0"` // 评论总数
	Title         string // 视频标题
	CreatedAt     int    // 创建时间
}

type UsersFavorVideos struct {
	ID      int64 // 主键
	UserId  int64 //用户id
	VideoId int64 // 视频id
}

// 视频的评论  一对多  一个视频对应多个评论
type Comment struct {
	ID        int64     // 主键
	VideoId   int64     // 视频id
	UserId    int64     // 用户id
	Content   string    // 发布内容
	CreatedAt time.Time // 发布时间
}

// 是否关注这个人   一对多   一个用户关注多个人
type Fans_followers struct {
	ID       int64 // 主键
	FansId   int64 // 粉丝id
	FollowId int64 // 关注者id
}

// 是否进行聊天   多对多   发布者和接收者的id都需要操作
type Chat struct {
	ID         int64     // 主键
	Content    string    // 消息内容
	CreatedAt  time.Time // 发送时间
	ToUserId   int64     // 发布者id
	FromUserId int64     // 接收者id
}

func MysqlInit() (*gorm.DB, error) {
	// 数据库连接
	username := "searchdata" // 使用者名字 如root
	password := "123456"
	host := "127.0.0.1"
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
		Addr:     "127.0.0.1:6379", // redis地址
		Password: "123456",         // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})
	// 检测是否连接成功
	_, err := redisDb.Ping().Result()
	if err != nil {
		return nil, errors.New("redis连接失败")
	} else {
		return redisDb, nil
	}
}

// 通过id读取缓存中的用户信息，如果未命中，则从数据库中查找，并且增加到缓存中
func RedisUserRead(db *gorm.DB, redisDatabase *redis.Client, id int64) (*User, error) {
	// 查询user缓存数据
	Htable := "user_" + strconv.FormatInt(id, 10)
	data, _ := redisDatabase.HGetAll(Htable).Result()

	var userinfo User

	// 如果缓存没有该user字段则返回数据库查询 如果有则直接返回
	if len(data) == 0 {

		// 数据库中查询
		err := db.Where("id = ?", id).First(&userinfo).Error
		if err != nil {
			return nil, errors.New("用户查询失败")
		}

		// 写入user缓存
		Hdata := make(map[string]interface{})
		Hdata["id"] = userinfo.ID
		Hdata["name"] = userinfo.Username
		Hdata["follow_count"] = userinfo.FollowCount
		Hdata["follower_count"] = userinfo.FansCount
		Hdata["avatar"] = userinfo.Avatar
		Hdata["background_image"] = userinfo.BackgroundImage
		Hdata["signature"] = userinfo.Signature
		Hdata["total_favorited"] = userinfo.TotalFavorited
		Hdata["work_count"] = userinfo.WorkCount
		Hdata["favorite_count"] = userinfo.FavoriteCount
		err = redisDatabase.HMSet(Htable, Hdata).Err()
		if err != nil {
			return nil, errors.New("用户写入失败")
		}
		userinfo.Password = ""
		return &userinfo, nil
	} else {
		id, _ := strconv.ParseInt(data["id"], 10, 64)
		follow_count, _ := strconv.ParseInt(data["follow_count"], 10, 64)
		follower_count, _ := strconv.ParseInt(data["follower_count"], 10, 64)
		work_count, _ := strconv.ParseInt(data["work_count"], 10, 64)
		favorite_count, _ := strconv.ParseInt(data["favorite_count"], 10, 64)
		return &User{
			ID:              id,
			Username:        data["name"],
			FansCount:       follower_count,
			FollowCount:     follow_count,
			Avatar:          data["avatar"],
			BackgroundImage: data["background_image"],
			Signature:       data["signature"],
			TotalFavorited:  data["total_favorited"],
			WorkCount:       work_count,
			FavoriteCount:   favorite_count,
		}, nil
	}
}

func IsFavor(db *gorm.DB, fansId int64, videoId int64) bool {
	if fansId == 0 {
		return false
	}
	var userfavorvideo UsersFavorVideos
	err := db.Where("user_id = ? AND video_id = ?", fansId, videoId).First(&userfavorvideo).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func IsFollow(db *gorm.DB, fansId int64, followerId int64) bool {
	if fansId == 0 || followerId == 0 {
		return false
	}
	if fansId == followerId {
		return false
	}
	if fansId == 0 || followerId == 0 {
		return false
	}
	var userfollow Fans_followers
	err := db.Where("fans_id = ? AND follow_id = ?", fansId, followerId).First(&userfollow).Error
	if err != nil {
		return false
	} else {
		return true
	}
}
