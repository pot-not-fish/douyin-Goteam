package logic

import (
	"context"
	"errors"
	"strconv"

	"douyin/apps/rpc/user/internal/svc"
	"douyin/apps/rpc/user/types/user"
	userdb "douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserinfoLogic) Userinfo(in *user.UserinfoReq) (*user.UserinfoResp, error) {
	// todo: add your logic here and delete this line

	// 缓存连接
	redisDb := l.svcCtx.DbRedis

	// 数据库连接
	// db, err := userdb.MysqlInit()
	// if err != nil {
	// 	return nil, errors.New("数据库连接失败")
	// }
	db := l.svcCtx.DbEngin

	// 查询缓存数据
	Htable := "user_" + strconv.FormatInt(in.ToUserId, 10)
	data, _ := redisDb.HGetAll(Htable).Result()

	// 如果缓存没有该字段则返回数据库查询
	if len(data) == 0 {

		// 数据库中查询
		userinfo := userdb.User{}
		err := db.Where("id = ?", in.ToUserId).First(&userinfo).Error
		if err != nil {
			return nil, errors.New("查询失败")
		}

		// 写入缓存
		Hdata := make(map[string]interface{})
		Hdata["id"] = userinfo.ID
		Hdata["name"] = userinfo.Username
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
			return nil, errors.New("写入失败")
		}
		return &user.UserinfoResp{
			Id:              userinfo.ID,
			Name:            userinfo.Username,
			FollowCount:     userinfo.FansCount,
			FollowerCount:   userinfo.FollowCount,
			IsFollow:        userdb.IsFollow(db, in.UserId, in.ToUserId), // 如果后续有关注操作，这里应该进行一次关联表的查询是否关注，暂时默认不关注
			Avatar:          userinfo.Avatar,
			BackgroundImage: userinfo.BackgroundImage,
			Signature:       userinfo.Signature,
			TotalFavorited:  userinfo.TotalFavorited,
			WorkCount:       userinfo.WorkCount,
			FavoriteCount:   userinfo.FavoriteCount}, nil
	} else {
		id, _ := strconv.ParseInt(data["id"], 10, 64)
		follow_count, _ := strconv.ParseInt(data["follow_count"], 10, 64)
		follower_count, _ := strconv.ParseInt(data["follower_count"], 10, 64)
		work_count, _ := strconv.ParseInt(data["work_count"], 10, 64)
		favorite_count, _ := strconv.ParseInt(data["favorite_count"], 10, 64)
		return &user.UserinfoResp{
			Id:              id,
			Name:            data["name"],
			FollowCount:     follow_count,
			FollowerCount:   follower_count,
			IsFollow:        userdb.IsFollow(db, in.UserId, in.ToUserId), // 如果后续有关注操作，这里应该进行一次关联表的查询是否关注，暂时默认不关注
			Avatar:          data["avatar"],
			BackgroundImage: data["background_image"],
			Signature:       data["signature"],
			TotalFavorited:  data["total_favorited"],
			WorkCount:       work_count,
			FavoriteCount:   favorite_count,
		}, nil
	}
}
