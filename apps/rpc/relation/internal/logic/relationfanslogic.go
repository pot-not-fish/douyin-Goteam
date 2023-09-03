package logic

import (
	"context"

	"douyin/apps/rpc/relation/internal/svc"
	"douyin/apps/rpc/relation/types/relation"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFansLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationFansLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFansLogic {
	return &RelationFansLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RelationFansLogic) RelationFans(in *relation.RelationFansReq) (*relation.RelationFansResp, error) {
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

	// 通过用户id查找粉丝的用户
	var follow []pkg.Fans_followers
	err := db.Where("follow_id = ?", in.UserId).Find(&follow).Error
	if err != nil {
		return &relation.RelationFansResp{UserList: nil}, nil
	}

	// 遍历关注的用户，查找用户信息
	var users []*relation.User
	var user *pkg.User
	for _, v := range follow {
		user, err = pkg.RedisUserRead(db, redisDb, v.FansId)
		if err != nil {
			return nil, err
		}
		users = append(users, &relation.User{
			Id:              user.ID,
			Name:            user.Username,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FansCount,
			IsFollow:        pkg.IsFollow(db, in.MeId, user.ID),
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		})
	}
	return &relation.RelationFansResp{UserList: users}, nil
}
