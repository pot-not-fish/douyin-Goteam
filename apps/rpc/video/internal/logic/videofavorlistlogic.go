package logic

import (
	"context"
	"errors"

	"douyin/apps/rpc/video/internal/svc"
	"douyin/apps/rpc/video/types/video"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoFavorListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoFavorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoFavorListLogic {
	return &VideoFavorListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoFavorListLogic) VideoFavorList(in *video.VideoFavorListReq) (*video.VideoFavorListResp, error) {
	// todo: add your logic here and delete this line

	// 数据库连接
	// db, err := pkg.MysqlInit()
	// if err != nil {
	// 	return nil, err
	// }
	db := l.svcCtx.DbEngin

	// 缓存连接
	// redisDb, err := pkg.RedisInit()
	// if err != nil {
	// 	return nil, err
	// }
	redisDb := l.svcCtx.DbRedis

	// 查询喜欢的视频的id
	var videofavor []pkg.UsersFavorVideos
	err := db.Where("user_id = ?", in.UserId).Find(&videofavor).Error
	if err != nil {
		return nil, errors.New("failed to search mysql video")
	}

	// 根据id查找相应的视频和用户 进行数据包装
	var videolist []*video.VideoList

	for i := len(videofavor) - 1; i >= 0; i-- {
		var userdata *pkg.User
		var videodata *pkg.Video
		err = db.Where("ID = ?", videofavor[i].VideoId).First(&videodata).Error
		if err != nil {
			return nil, err
		}
		userdata, err = pkg.RedisUserRead(db, redisDb, videodata.UserId)
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
				IsFollow:        pkg.IsFollow(db, in.UserId, userdata.ID),
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
			IsFavorite:    true,
			Title:         videodata.Title,
		})
	}

	// for _, v := range videofavor {
	// 	var userdata *pkg.User
	// 	var videodata *pkg.Video
	// 	err = db.Where("ID = ?", v.VideoId).First(&videodata).Error
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	userdata, err = pkg.RedisUserRead(db, redisDb, videodata.UserId)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	videolist = append(videolist, &video.VideoList{
	// 		Id: videodata.ID,
	// 		Author: &video.Author{
	// 			Id:              userdata.ID,
	// 			Name:            userdata.Username,
	// 			FollowCount:     userdata.FollowCount,
	// 			FollowerCount:   userdata.FansCount,
	// 			IsFollow:        pkg.IsFollow(db, in.UserId, userdata.ID),
	// 			Avatar:          userdata.Avatar,
	// 			BackgroundImage: userdata.BackgroundImage,
	// 			Signature:       userdata.Signature,
	// 			TotalFavorited:  userdata.TotalFavorited,
	// 			WorkCount:       userdata.WorkCount,
	// 			FavoriteCount:   userdata.FavoriteCount,
	// 		},
	// 		PlayUrl:       videodata.PlayUrl,
	// 		CoverUrl:      videodata.CoverUrl,
	// 		FavoriteCount: videodata.FavoriteCount,
	// 		CommentCount:  videodata.CommentCount,
	// 		IsFavorite:    true,
	// 		Title:         videodata.Title,
	// 	})
	// }
	return &video.VideoFavorListResp{Videos: videolist}, nil
}
