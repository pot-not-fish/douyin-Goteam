package relation

import (
	"context"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/relation/types/relation"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationfriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationfriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationfriendLogic {
	return &RelationfriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationfriendLogic) Relationfriend(req *types.RelationFriendReq) (resp *types.RelationFriendResp, err error) {
	// todo: add your logic here and delete this line

	// token鉴权
	_, err = pkg.AuthToken(req.Token)
	if err != nil {
		return &types.RelationFriendResp{
			Status_code: 1,
			Status_msg:  err.Error(),
			User_list:   nil,
		}, nil
	}

	// 发送给relationfriend的rpc服务处理
	relationfriendResp, err := l.svcCtx.RelationRpc.RelationFriend(l.ctx, &relation.RelationFriendReq{UserId: req.User_id})
	if err != nil {
		return &types.RelationFriendResp{
			Status_code: 1,
			Status_msg:  err.Error(),
			User_list:   nil,
		}, nil
	}
	var friendlist []types.FreindUser
	for _, v := range relationfriendResp.UserList {
		if v.Message == nil || v.MsgType == nil {
			friendlist = append(friendlist, types.FreindUser{
				Id:               v.Id,
				Name:             v.Name,
				Follow_count:     v.FollowCount,
				Follower_count:   v.FollowerCount,
				Is_follow:        v.IsFollow,
				Avatar:           v.Avatar,
				Background_image: v.BackgroundImage,
				Signature:        v.Signature,
				Total_favorited:  v.TotalFavorited,
				Work_count:       v.WorkCount,
				Favorite_count:   v.FavoriteCount,
				Message:          "",
				MsgType:          0,
			})
		} else {
			friendlist = append(friendlist, types.FreindUser{
				Id:               v.Id,
				Name:             v.Name,
				Follow_count:     v.FollowCount,
				Follower_count:   v.FollowerCount,
				Is_follow:        v.IsFollow,
				Avatar:           v.Avatar,
				Background_image: v.BackgroundImage,
				Signature:        v.Signature,
				Total_favorited:  v.TotalFavorited,
				Work_count:       v.WorkCount,
				Favorite_count:   v.FavoriteCount,
				Message:          *v.Message,
				MsgType:          *v.MsgType,
			})
		}
	}
	return &types.RelationFriendResp{
		Status_code: 0,
		Status_msg:  "获取好友列表成功",
		User_list:   friendlist,
	}, nil
}
