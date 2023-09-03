package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"douyin/apps/rpc/video/internal/svc"
	"douyin/apps/rpc/video/types/video"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type VideoFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoFavoriteLogic {
	return &VideoFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func dbupdateVideoFavor(db *gorm.DB, userdata, touserdata *pkg.User, total, count int64) {
	err := db.Model(&touserdata).Where("id = ?", touserdata.ID).Update("TotalFavorited", total).Error
	if err != nil {
		fmt.Println("failed to increase mysql user total favorite")
		return
	}
	err = db.Model(&userdata).Where("id = ?", userdata.ID).Update("favorite_count", count).Error
	if err != nil {
		fmt.Println("failed to increase mysql user favorite count")
		return
	}
}

func dbupdateVideo(db *gorm.DB, videodata *pkg.Video, VideoId int64) {
	err := db.Model(&videodata).Where("ID = ?", VideoId).Update("FavoriteCount", videodata.FavoriteCount).Error
	if err != nil {
		fmt.Println("failed to update video")
		return
	}
}

func (l *VideoFavoriteLogic) VideoFavorite(in *video.VideoFavoriteReq) (*video.VideoFavoriteResp, error) {
	// todo: add your logic here and delete this line

	// 连接redis
	// redisDb, err := pkg.RedisInit()
	// if err != nil {
	// 	return nil, errors.New("failed to link redis")
	// }
	redisDb := l.svcCtx.DbRedis

	// 数据库连接
	// db, err := pkg.MysqlInit()
	// if err != nil {
	// 	return nil, errors.New("failed to link mysql")
	// }
	db := l.svcCtx.DbEngin

	// 创建关系字段
	userfavorvideo := pkg.UsersFavorVideos{
		UserId:  in.UserId,
		VideoId: in.VideoId,
	}
	if in.ActionType == 1 {
		err := db.Where("user_id = ? AND video_id = ?", in.UserId, in.VideoId).First(&userfavorvideo).Error
		if err == nil {
			return nil, errors.New("已经点赞过该视频")
		}
		err = db.Create(&userfavorvideo).Error
		if err != nil {
			return nil, errors.New("failed to create userfavorvideo")
		}
	} else {
		err := db.Unscoped().Where("user_id = ? AND video_id = ?", in.UserId, in.VideoId).Delete(&userfavorvideo).Error
		if err != nil {
			return nil, errors.New("failed to delete userfavorvideo")
		}
	}

	// 遍历list 查找相应的视频 修改喜欢数量 (如果前端提供视频时间戳，通过时间戳分片的查询效率更高)
	VideoList, err := redisDb.LRange("videos", 0, -1).Result()
	if err != nil {
		return nil, errors.New("failed to retrieve video list")
	}
	var videodata pkg.Video
	for k, v := range VideoList {
		err = json.Unmarshal([]byte(v), &videodata)
		if err != nil {
			return nil, errors.New("failed to unmarshal video")
		}
		if videodata.ID == in.VideoId {
			if in.ActionType == 1 { // 点赞
				videodata.FavoriteCount++
			} else {
				videodata.FavoriteCount--
			}
			data, err := json.Marshal(videodata)
			if err != nil {
				return nil, errors.New("failed to marshal")
			}
			redisDb.LSet("videos", int64(k), data)
			break
		}
	}

	// 数据库喜欢列表的视频操作
	// 异步处理
	// err = db.Model(&videodata).Where("ID = ?", in.VideoId).Update("FavoriteCount", videodata.FavoriteCount).Error
	// if err != nil {
	// 	return nil, errors.New("failed to update video")
	// }
	go dbupdateVideo(db, &videodata, in.VideoId)

	// 点赞用户的喜欢数+1  获赞用户的获赞量+1
	var userdata *pkg.User
	var touserdata *pkg.User
	userdata, err = pkg.RedisUserRead(db, redisDb, in.UserId)
	if err != nil {
		return nil, errors.New("failed to search mysql user")
	}
	touserdata, err = pkg.RedisUserRead(db, redisDb, videodata.UserId)
	if err != nil {
		return nil, errors.New("failed to search mysql video user")
	}
	totalfavor, _ := strconv.ParseInt(touserdata.TotalFavorited, 10, 64)
	var total int64 // 获赞量
	var count int64 // 喜欢数
	if in.ActionType == 1 {
		total = totalfavor + 1
		count = userdata.FavoriteCount + 1
	} else {
		total = totalfavor - 1
		count = userdata.FavoriteCount - 1
	}

	// 异步操作
	// err = db.Model(&touserdata).Where("id = ?", touserdata.ID).Update("TotalFavorited", total).Error
	// if err != nil {
	// 	return nil, errors.New("failed to increase mysql user total favorite")
	// }
	// err = db.Model(&userdata).Where("id = ?", userdata.ID).Update("favorite_count", count).Error
	// if err != nil {
	// 	return nil, errors.New("failed to increase mysql user favorite count")
	// }
	go dbupdateVideoFavor(db, touserdata, userdata, total, count)

	// 缓存用户获赞量操作
	user_str := strconv.FormatInt(userdata.ID, 10)
	touser_str := strconv.FormatInt(touserdata.ID, 10)
	_, err = redisDb.HSet("user_"+touser_str, "total_favorited", total).Result()
	if err != nil {
		return nil, errors.New("failed to increase redis user total favorite")
	}
	_, err = redisDb.HSet("user_"+user_str, "favorite_count", count).Result()
	if err != nil {
		return nil, errors.New("failed to increase redis user ")
	}

	if in.ActionType == 1 {
		return &video.VideoFavoriteResp{StatusMsg: "点赞成功"}, nil
	} else {
		return &video.VideoFavoriteResp{StatusMsg: "取消点赞成功"}, nil
	}
}
