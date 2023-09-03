package logic

import (
	"context"
	"errors"
	"strconv"

	"douyin/apps/rpc/video/internal/svc"
	"douyin/apps/rpc/video/types/video"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoListLogic) VideoList(in *video.VideoListReq) (*video.VideoListResp, error) {
	// todo: add your logic here and delete this line

	// video和user数据
	userinfo := pkg.User{}
	videoinfo := []pkg.Video{}

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

	// 查询user缓存数据
	Htable := "user_" + strconv.FormatInt(in.UserId, 10)
	data, _ := redisDb.HGetAll(Htable).Result()

	// 如果缓存没有该user字段则返回数据库查询
	if len(data) == 0 {

		// 数据库中查询
		err := db.Where("id = ?", in.UserId).First(&userinfo).Error
		if err != nil {
			return nil, errors.New("用户查询失败")
		}

		// 写入user缓存
		Hdata := make(map[string]interface{})
		Hdata["id"] = userinfo.ID
		Hdata["name"] = userinfo.Username
		Hdata["fans_count"] = userinfo.FansCount
		Hdata["follow_count"] = userinfo.FansCount
		Hdata["follower_count"] = userinfo.FollowCount
		Hdata["avatar"] = userinfo.Avatar
		Hdata["background_image"] = userinfo.BackgroundImage
		Hdata["signature"] = userinfo.Signature
		Hdata["total_favorited"] = userinfo.TotalFavorited
		Hdata["work_count"] = userinfo.WorkCount
		Hdata["favorite_count"] = userinfo.FavoriteCount
		err = redisDb.HMSet(Htable, Hdata).Err()
		if err != nil {
			return nil, errors.New("用户写入失败")
		}

		// 查询视频信息
		err = db.Where("user_id", in.UserId).Find(&videoinfo).Error
		if err != nil {
			return nil, errors.New("视频查询失败")
		}

		videolist := []*video.VideoList{}

		for i := len(videoinfo) - 1; i >= 0; i-- {
			videolist = append(videolist, &video.VideoList{
				Id: videoinfo[i].ID,
				Author: &video.Author{
					Id:              userinfo.ID,
					Name:            userinfo.Username,
					FollowCount:     userinfo.FollowCount,
					FollowerCount:   userinfo.FansCount,
					IsFollow:        pkg.IsFollow(db, in.UserId, userinfo.ID), // 后续添加关注操作需要判断
					Avatar:          userinfo.Avatar,
					BackgroundImage: userinfo.BackgroundImage,
					Signature:       userinfo.Signature,
					TotalFavorited:  userinfo.TotalFavorited,
					WorkCount:       userinfo.WorkCount,
					FavoriteCount:   userinfo.FavoriteCount,
				},
				PlayUrl:       videoinfo[i].PlayUrl,
				CoverUrl:      videoinfo[i].CoverUrl,
				FavoriteCount: videoinfo[i].FavoriteCount,
				CommentCount:  videoinfo[i].CommentCount,
				IsFavorite:    pkg.IsFavor(db, in.UserId, videoinfo[i].ID), // 后续添加喜欢操作需要判断
				Title:         videoinfo[i].Title,
			})
		}

		// for _, v := range videoinfo {
		// 	videolist = append(videolist, &video.VideoList{
		// 		Id: v.ID,
		// 		Author: &video.Author{
		// 			Id:              userinfo.ID,
		// 			Name:            userinfo.Username,
		// 			FollowCount:     userinfo.FollowCount,
		// 			FollowerCount:   userinfo.FansCount,
		// 			IsFollow:        pkg.IsFollow(db, in.UserId, userinfo.ID), // 后续添加关注操作需要判断
		// 			Avatar:          userinfo.Avatar,
		// 			BackgroundImage: userinfo.BackgroundImage,
		// 			Signature:       userinfo.Signature,
		// 			TotalFavorited:  userinfo.TotalFavorited,
		// 			WorkCount:       userinfo.WorkCount,
		// 			FavoriteCount:   userinfo.FavoriteCount,
		// 		},
		// 		PlayUrl:       v.PlayUrl,
		// 		CoverUrl:      v.CoverUrl,
		// 		FavoriteCount: v.FavoriteCount,
		// 		CommentCount:  v.CommentCount,
		// 		IsFavorite:    pkg.IsFavor(db, in.UserId, v.ID), // 后续添加喜欢操作需要判断
		// 		Title:         v.Title,
		// 	})
		// }
		return &video.VideoListResp{Videos: videolist}, nil
	} else {
		// 查询视频信息
		err := db.Where("user_id", in.UserId).Find(&videoinfo).Error
		if err != nil {
			return nil, errors.New("视频查询失败")
		}
		id, _ := strconv.ParseInt(data["id"], 10, 64)
		follow_count, _ := strconv.ParseInt(data["follow_count"], 10, 64)
		follower_count, _ := strconv.ParseInt(data["follower_count"], 10, 64)
		work_count, _ := strconv.ParseInt(data["work_count"], 10, 64)
		favorite_count, _ := strconv.ParseInt(data["favorite_count"], 10, 64)
		videolist := []*video.VideoList{}

		for i := len(videoinfo) - 1; i >= 0; i-- {
			videolist = append(videolist, &video.VideoList{
				Id: videoinfo[i].ID,
				Author: &video.Author{
					Id:              id,
					Name:            data["name"],
					FollowCount:     follow_count,
					FollowerCount:   follower_count,
					IsFollow:        pkg.IsFollow(db, in.UserId, id), // 后续添加关注操作需要判断
					Avatar:          data["avatar"],
					BackgroundImage: data["background_image"],
					Signature:       data["signature"],
					TotalFavorited:  data["total_favorited"],
					WorkCount:       work_count,
					FavoriteCount:   favorite_count,
				},
				PlayUrl:       videoinfo[i].PlayUrl,
				CoverUrl:      videoinfo[i].CoverUrl,
				FavoriteCount: videoinfo[i].FavoriteCount,
				CommentCount:  videoinfo[i].CommentCount,
				IsFavorite:    pkg.IsFavor(db, in.UserId, videoinfo[i].ID), // 后续添加喜欢操作需要判断
				Title:         videoinfo[i].Title,
			})
		}

		// for _, v := range videoinfo {
		// 	videolist = append(videolist, &video.VideoList{
		// 		Id: v.ID,
		// 		Author: &video.Author{
		// 			Id:              id,
		// 			Name:            data["name"],
		// 			FollowCount:     follow_count,
		// 			FollowerCount:   follower_count,
		// 			IsFollow:        pkg.IsFollow(db, in.UserId, id), // 后续添加关注操作需要判断
		// 			Avatar:          data["avatar"],
		// 			BackgroundImage: data["background_image"],
		// 			Signature:       data["signature"],
		// 			TotalFavorited:  data["total_favorited"],
		// 			WorkCount:       work_count,
		// 			FavoriteCount:   favorite_count,
		// 		},
		// 		PlayUrl:       v.PlayUrl,
		// 		CoverUrl:      v.CoverUrl,
		// 		FavoriteCount: v.FavoriteCount,
		// 		CommentCount:  v.CommentCount,
		// 		IsFavorite:    pkg.IsFavor(db, in.UserId, v.ID), // 后续添加喜欢操作需要判断
		// 		Title:         v.Title,
		// 	})
		// }
		return &video.VideoListResp{Videos: videolist}, nil
	}
}
