package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"douyin/apps/rpc/video/internal/svc"
	"douyin/apps/rpc/video/types/video"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func dbupdate(db *gorm.DB, videodata *pkg.Video) {
	err := db.Model(&videodata).Where("ID = ?", videodata.ID).Update("CommentCount", videodata.CommentCount).Error
	if err != nil {
		return
	}
}

func (l *CommentActionLogic) CommentAction(in *video.CommentActionReq) (*video.CommentActionResp, error) {
	// todo: add your logic here and delete this line

	// 连接数据库
	// db, err := pkg.MysqlInit()
	// if err != nil {
	// 	return nil, err
	// }
	db := l.svcCtx.DbEngin

	// 创建关系字段
	comment := pkg.Comment{
		UserId:  in.UserId,
		VideoId: in.VideoId,
		Content: in.CommentText,
	}
	if in.ActionType == 1 { // 判断发布评论 还是删除评论
		err := db.Create(&comment).Error
		if err != nil {
			return nil, errors.New("failed to create comment")
		}
	} else {
		err := db.Unscoped().Where("ID = ?", in.CommentId).Delete(&comment).Error
		if err != nil {
			return nil, errors.New("failed to delete comment")
		}
	}

	// 连接缓存
	redisDb, err := pkg.RedisInit()
	if err != nil {
		return nil, err
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
			if in.ActionType == 1 { // 评论
				videodata.CommentCount++
			} else {
				videodata.CommentCount--
			}
			data, err := json.Marshal(videodata)
			if err != nil {
				return nil, errors.New("failed to marshal")
			}
			redisDb.LSet("videos", int64(k), data)
		}
	}

	// 数据库喜欢列表的视频操作
	// 异步操作
	// err = db.Model(&videodata).Where("ID = ?", videodata.ID).Update("CommentCount", videodata.CommentCount).Error
	// if err != nil {
	// 	return nil, errors.New("failed to update video")
	// }
	go dbupdate(db, &videodata)

	// 获取用户信息
	var user *pkg.User
	user, err = pkg.RedisUserRead(db, redisDb, in.UserId)
	if err != nil {
		return nil, err
	}

	_, mouth, day := comment.CreatedAt.Date()
	t := fmt.Sprintf("%d-%d", mouth, day)

	return &video.CommentActionResp{
		Id: comment.ID,
		User: &video.Author{
			Id:              user.ID,
			Name:            user.Username,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FansCount,
			IsFollow:        false, // 自己发布评论不能关注自己
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		},
		Content:    in.CommentText,
		CreateDate: t,
	}, nil
}
