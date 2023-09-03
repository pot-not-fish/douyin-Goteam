package video

import (
	"context"
	"strconv"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/video/types/video"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteLogic) Favorite(req *types.VideoFavoriteReq) (resp *types.VideoFavoriteResp, err error) {
	// todo: add your logic here and delete this line

	// token鉴权
	id, err := pkg.AuthToken(req.Token)
	if err != nil {
		return &types.VideoFavoriteResp{
			Status_code: 1,
			Status_msg:  err.Error(),
		}, nil
	}

	// 发送给favoritevideo的rpc服务处理
	videoid, _ := strconv.ParseInt(req.Video_id, 10, 64)
	actiontype, _ := strconv.ParseInt(req.Action_type, 10, 32)
	videofavorResp, err := l.svcCtx.VideoRpc.VideoFavorite(l.ctx, &video.VideoFavoriteReq{VideoId: videoid, UserId: id, ActionType: int32(actiontype)})
	if err != nil {
		return &types.VideoFavoriteResp{
			Status_code: 1,
			Status_msg:  err.Error(),
		}, nil
	}
	return &types.VideoFavoriteResp{
		Status_code: 0,
		Status_msg:  videofavorResp.StatusMsg,
	}, nil
}
