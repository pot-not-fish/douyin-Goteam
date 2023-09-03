package logic

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"douyin/apps/rpc/video/internal/svc"
	"douyin/apps/rpc/video/types/video"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoFeedLogic {
	return &VideoFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoFeedLogic) VideoFeed(in *video.VideoFeedReq) (*video.VideoFeedResp, error) {
	// todo: add your logic here and delete this line

	// 缓存连接
	// redisDb, err := pkg.RedisInit()
	// if err != nil {
	// 	return nil, err
	// }
	redisDb := l.svcCtx.DbRedis

	// 数据库连接
	// db, err := pkg.MysqlInit()
	// if err != nil {
	// 	return nil, errors.New("数据库连接失败")
	// }
	db := l.svcCtx.DbEngin

	// redis time列表获取
	timeset, err := redisDb.LRange("time", 0, -1).Result()
	if err != nil {
		return nil, errors.New("time列表查找失败")
	}

	// 判断视频流需要返回的位置
	var left, right int64
	var next_time int64
	VLen, err := redisDb.LLen("videos").Result()
	if err != nil {
		return nil, errors.New("vidoes长度获取失败")
	}
	TLen, err := redisDb.LLen("time").Result()
	if err != nil {
		return nil, errors.New("time长度获取失败")
	}
	for k, v := range timeset {
		t, _ := strconv.ParseInt(v, 10, 64)
		if in.LatestTime >= t {
			right = VLen - 10*(TLen-int64(k)-1) - 1
			if k != 0 {
				left = right - 9
			} else {
				left = 0
			}
			if k != int(TLen-1) {
				n_t, _ := strconv.ParseInt(timeset[k+1], 10, 64)
				next_time = n_t
			} else {
				next_time = time.Now().Unix()
			}
			break
		}
	}
	result, err := redisDb.LRange("videos", left, right).Result()
	if err != nil {
		return nil, errors.New("获取video列表失败")
	}

	var videolist []*video.VideoList
	for _, v := range result {
		var videodata pkg.Video
		var userdata *pkg.User
		err := json.Unmarshal([]byte(v), &videodata)
		if err != nil {
			return nil, errors.New("视频反序列化失败")
		}
		err = db.Where("Id = ?", videodata.UserId).First(&userdata).Error
		if err != nil {
			return nil, errors.New("查找视频对应用户失败")
		}

		// 获取user数据
		userdata, err = pkg.RedisUserRead(db, redisDb, userdata.ID)
		if err != nil {
			return nil, err
		}
		videolist = append(videolist, &video.VideoList{
			Id: videodata.ID,
			Author: &video.Author{
				Id:              userdata.ID,
				Name:            userdata.Username,
				FollowCount:     userdata.FollowCount,
				FollowerCount:   userdata.FansCount,
				IsFollow:        pkg.IsFollow(db, in.UserId, userdata.ID), // 后续添加关注操作需要判断
				Avatar:          userdata.Avatar,
				BackgroundImage: userdata.BackgroundImage,
				Signature:       userdata.Signature,
				TotalFavorited:  userdata.TotalFavorited,
				WorkCount:       userdata.WorkCount,
				FavoriteCount:   userdata.FavoriteCount,
			},
			PlayUrl:       videodata.PlayUrl,
			CoverUrl:      videodata.CoverUrl,
			FavoriteCount: videodata.FavoriteCount,
			CommentCount:  videodata.CommentCount,
			IsFavorite:    pkg.IsFavor(db, in.UserId, videodata.ID), // 后续添加喜欢操作需要判断
			Title:         videodata.Title,
		})
	}
	return &video.VideoFeedResp{Videos: videolist, NextTime: next_time}, nil
}
