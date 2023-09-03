package user

import (
	"context"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/user/types/user"
	authcrypto "douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.UserinfoReq) (resp *types.UserinfoResp, err error) {
	// todo: add your logic here and delete this line

	// token鉴权
	id, err := authcrypto.AuthToken(req.Token)
	if err != nil {
		return &types.UserinfoResp{Status_code: 1, Status_msg: err.Error(), User: nil}, nil
	}

	// 发送给rpc的userinfo处理
	userinfoResp, err := l.svcCtx.UserRpc.Userinfo(l.ctx, &user.UserinfoReq{UserId: id, ToUserId: req.User_id})
	if err != nil {
		return &types.UserinfoResp{Status_code: 1, Status_msg: err.Error(), User: nil}, nil
	} else {
		return &types.UserinfoResp{
			Status_code: 0,
			Status_msg:  "查询成功",
			User: &types.User{
				Id:               userinfoResp.Id,
				Name:             userinfoResp.Name,
				Follow_count:     userinfoResp.FollowCount,
				Follower_count:   userinfoResp.FollowerCount,
				Is_follow:        userinfoResp.IsFollow,
				Avatar:           userinfoResp.Avatar,
				Background_image: userinfoResp.BackgroundImage,
				Signature:        userinfoResp.Signature,
				Total_favorited:  userinfoResp.TotalFavorited,
				Work_count:       userinfoResp.WorkCount,
				Favorite_count:   userinfoResp.FavoriteCount}}, nil
	}
}
