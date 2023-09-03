package logic

import (
	"context"
	"errors"
	"fmt"

	"douyin/apps/rpc/video/internal/svc"
	"douyin/apps/rpc/video/types/video"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *video.CommentListReq) (*video.CommentListResp, error) {
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

	// 通过视频id查找视频的评论
	var comments []pkg.Comment
	err := db.Where("video_id = ?", in.VideoId).Find(&comments).Error
	if err != nil {
		return nil, errors.New("failed to find video comments")
	}

	// 通过视频评论的用户id查找用户
	var videoComments []*video.CommentActionResp
	for _, v := range comments {
		var userinfo *pkg.User
		userinfo, err = pkg.RedisUserRead(db, redisDb, v.UserId)
		if err != nil {
			return nil, err
		}
		_, mouth, day := v.CreatedAt.Date()
		t := fmt.Sprintf("%d-%d", mouth, day)
		videoComments = append(videoComments, &video.CommentActionResp{
			Id: v.ID,
			User: &video.Author{
				Id:              userinfo.ID,
				Name:            userinfo.Username,
				FollowCount:     userinfo.FollowCount,
				FollowerCount:   userinfo.FansCount,
				IsFollow:        pkg.IsFollow(db, in.UserId, userinfo.ID),
				Avatar:          userinfo.Avatar,
				BackgroundImage: userinfo.BackgroundImage,
				Signature:       userinfo.Signature,
				TotalFavorited:  userinfo.TotalFavorited,
				WorkCount:       userinfo.WorkCount,
				FavoriteCount:   userinfo.FavoriteCount,
			},
			Content:    v.Content,
			CreateDate: t,
		})
	}

	return &video.CommentListResp{Comments: videoComments}, nil
}
