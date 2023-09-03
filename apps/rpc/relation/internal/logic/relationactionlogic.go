package logic

import (
	"context"
	"errors"
	"strconv"

	"douyin/apps/rpc/relation/internal/svc"
	"douyin/apps/rpc/relation/types/relation"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type RelationActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func dbupdate(db *gorm.DB, fans *pkg.User, follow *pkg.User, UserId, ToUserId, followercount, fanscount int64) {
	err := db.Model(&fans).Where("id = ?", UserId).Update("follow_count", followercount).Error
	if err != nil {
		return
	}
	err = db.Model(&follow).Where("id = ?", ToUserId).Update("fans_count", fanscount).Error
	if err != nil {
		return
	}
}

func (l *RelationActionLogic) RelationAction(in *relation.RelationActionReq) (*relation.RelationActionResp, error) {
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

	// 创建关系字段
	userfollow := pkg.Fans_followers{
		FansId:   in.UserId,
		FollowId: in.ToUserId,
	}
	if in.ActionType == 1 {
		err := db.Where("fans_id = ? AND follow_id = ?", in.UserId, in.ToUserId).First(&userfollow).Error
		if err == nil {
			return nil, errors.New("已经关注过该用户")
		}
		err = db.Create(&userfollow).Error
		if err != nil {
			return nil, errors.New("failed to create fans_followers")
		}
	} else {
		err := db.Unscoped().Where("fans_id = ? AND follow_id = ?", in.UserId, in.ToUserId).Delete(&userfollow).Error
		if err != nil {
			return nil, errors.New("不能取关未关注的用户")
		}
	}

	// 查找fans目前的关注数 follower目前粉丝数
	var fans *pkg.User
	var follower *pkg.User
	fans, err := pkg.RedisUserRead(db, redisDb, in.UserId)
	if err != nil {
		return nil, errors.New("failed to search fans user mysql")
	}
	follower, err = pkg.RedisUserRead(db, redisDb, in.ToUserId)
	if err != nil {
		return nil, errors.New("failed to search follower user mysql")
	}

	// 判断关注还是取关操作发生的用户变化
	var fanscount int64
	var followercount int64
	if in.ActionType == 1 {
		fanscount = follower.FansCount + 1
		followercount = fans.FollowCount + 1
	} else {
		fanscount = follower.FansCount - 1
		followercount = fans.FollowCount - 1
	}

	// 异步操作
	// err = db.Model(&fans).Where("id = ?", in.UserId).Update("follow_count", followercount).Error
	// if err != nil {
	// 	return nil, errors.New("failed to update fans user mysql")
	// }
	// err = db.Model(&follower).Where("id = ?", in.ToUserId).Update("fans_count", fanscount).Error
	// if err != nil {
	// 	return nil, errors.New("failed to update follower user mysql")
	// }
	go dbupdate(db, fans, follower, in.UserId, in.ToUserId, followercount, fanscount)

	// 缓存修改fans的关注数和followers的粉丝数
	fans_str := strconv.FormatInt(in.UserId, 10)
	_, err = redisDb.HSet("user_"+fans_str, "follow_count", followercount).Result()
	if err != nil {
		return nil, errors.New("failed to update redis user follower")
	}
	follower_str := strconv.FormatInt(in.ToUserId, 10)
	_, err = redisDb.HSet("user_"+follower_str, "follower_count", fanscount).Result()
	if err != nil {
		return nil, errors.New("failed to update redis user fans")
	}

	if in.ActionType == 1 {
		return &relation.RelationActionResp{StatusMsg: "关注成功"}, nil
	} else {
		return &relation.RelationActionResp{StatusMsg: "取关成功"}, nil
	}
}
