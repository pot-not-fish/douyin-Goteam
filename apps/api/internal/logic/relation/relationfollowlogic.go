package relation

import (
	"context"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/relation/types/relation"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationfollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationfollowLogic {
	return &RelationfollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationfollowLogic) Relationfollow(req *types.RelationFollowReq) (resp *types.RelationFollowResp, err error) {
	// todo: add your logic here and delete this line

	// 获取id
	var id int64
	if len(req.Token) == 0 {
		id = 0
	} else {
		id, err = pkg.AuthToken(req.Token)
		if err != nil {
			return &types.RelationFollowResp{
				Status_code: 1,
				Status_msg:  err.Error(),
				User_list:   []types.Author{},
			}, nil
		}
	}

	// 发送给relationfans的rpc处理
	relationfollowResp, err := l.svcCtx.RelationRpc.RelationFollow(l.ctx, &relation.RelationFollowReq{UserId: req.User_id, MeId: id})
	if err != nil {
		return &types.RelationFollowResp{
			Status_code: 1,
			Status_msg:  err.Error(),
			User_list:   []types.Author{},
		}, nil
	}

	var userlist []types.Author
	for _, v := range relationfollowResp.UserList {
		userlist = append(userlist, types.Author{
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
		})
	}
	return &types.RelationFollowResp{
		Status_code: 0,
		Status_msg:  "关注列表获取成功",
		User_list:   userlist,
	}, nil
}
