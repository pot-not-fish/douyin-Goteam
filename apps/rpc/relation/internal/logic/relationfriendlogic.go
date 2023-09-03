package logic

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"douyin/apps/rpc/relation/internal/svc"
	"douyin/apps/rpc/relation/types/relation"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFriendLogic {
	return &RelationFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RelationFriendLogic) RelationFriend(in *relation.RelationFriendReq) (*relation.RelationFriendResp, error) {
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

	// 通过用户id查找关注的用户
	var follow []pkg.Fans_followers
	err := db.Where("fans_id = ?", in.UserId).Find(&follow).Error
	if err != nil {
		return &relation.RelationFriendResp{UserList: nil}, nil
	}

	// 遍历follow的id，如果follow也关注了fans，则是好友关系
	var friendlist []*relation.FriendUser
	for _, v := range follow {
		var fansfollow pkg.Fans_followers
		var friendData *pkg.User
		err = db.Where("fans_id = ? AND follow_id = ?", v.FollowId, v.FansId).First(&fansfollow).Error
		if err != nil {
			continue
		}
		friendData, err = pkg.RedisUserRead(db, redisDb, v.FollowId)
		if err != nil {
			return nil, err
		}
		userid_str := strconv.FormatInt(in.UserId, 10)
		to_userid_str := strconv.FormatInt(v.FollowId, 10)
		// 查找最新的聊天消息
		var message *string
		var msgtype *int64
		var msgtypeInt int64
		var comment pkg.Chat
		messagelist1, err := redisDb.LLen(userid_str + "_" + to_userid_str).Result()
		if err != nil {
			return nil, errors.New("failed to find userid touserid redis")
		}
		messagelist2, err := redisDb.LLen(to_userid_str + "_" + userid_str).Result()
		if err != nil {
			return nil, errors.New("failed to find touserid userid redis")
		}
		if messagelist1 == 0 && messagelist2 == 0 {
			friendlist = append(friendlist, &relation.FriendUser{
				Id:              friendData.ID,
				Name:            friendData.Username,
				FollowCount:     friendData.FollowCount,
				FollowerCount:   friendData.FansCount,
				IsFollow:        true,
				Avatar:          friendData.Avatar,
				BackgroundImage: friendData.BackgroundImage,
				Signature:       friendData.Signature,
				TotalFavorited:  friendData.TotalFavorited,
				WorkCount:       friendData.WorkCount,
				FavoriteCount:   friendData.FavoriteCount,
			})
			continue
		} else if messagelist1 != 0 && messagelist2 == 0 {
			rawcomment, err := redisDb.LIndex(userid_str+"_"+to_userid_str, -1).Result()
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal([]byte(rawcomment), &comment)
			if err != nil {
				return nil, errors.New("fail to umarshal comment")
			}
			message = &comment.Content
			if comment.FromUserId == in.UserId {
				msgtypeInt = 1
			} else {
				msgtypeInt = 0
			}
			msgtype = &msgtypeInt
		} else if messagelist1 == 0 && messagelist2 != 0 {
			rawcomment, err := redisDb.LIndex(to_userid_str+"_"+userid_str, -1).Result()
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal([]byte(rawcomment), &comment)
			if err != nil {
				return nil, errors.New("fail to umarshal comment")
			}
			message = &comment.Content
			if comment.FromUserId == in.UserId {
				msgtypeInt = 1
			} else {
				msgtypeInt = 0
			}
			msgtype = &msgtypeInt
		}
		friendlist = append(friendlist, &relation.FriendUser{
			Id:              friendData.ID,
			Name:            friendData.Username,
			FollowCount:     friendData.FollowCount,
			FollowerCount:   friendData.FansCount,
			IsFollow:        true,
			Avatar:          friendData.Avatar,
			BackgroundImage: friendData.BackgroundImage,
			Signature:       friendData.Signature,
			TotalFavorited:  friendData.TotalFavorited,
			WorkCount:       friendData.WorkCount,
			FavoriteCount:   friendData.FavoriteCount,
			Message:         message,
			MsgType:         msgtype,
		})
	}
	return &relation.RelationFriendResp{UserList: friendlist}, nil
}
