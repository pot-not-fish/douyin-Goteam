package relation

import (
	"context"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/relation/types/relation"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationfansLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationfansLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationfansLogic {
	return &RelationfansLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationfansLogic) Relationfans(req *types.RelationFansReq) (resp *types.RelationFansResp, err error) {
	// todo: add your logic here and delete this line

	// 获取id
	var id int64
	if len(req.Token) == 0 {
		id = 0
	} else {
		id, err = pkg.AuthToken(req.Token)
		if err != nil {
			return &types.RelationFansResp{
				Status_code: 1,
				Status_msg:  err.Error(),
				User_list:   []types.Author{},
			}, nil
		}
	}

	// 发送给relationfans的rpc处理
	relationfansResp, err := l.svcCtx.RelationRpc.RelationFans(l.ctx, &relation.RelationFansReq{UserId: req.User_id, MeId: id})
	if err != nil {
		return &types.RelationFansResp{
			Status_code: 1,
			Status_msg:  err.Error(),
			User_list:   []types.Author{},
		}, nil
	}

	var userlist []types.Author
	for _, v := range relationfansResp.UserList {
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
	return &types.RelationFansResp{
		Status_code: 0,
		Status_msg:  "获取粉丝列表成功",
		User_list:   userlist,
	}, nil
}
