package video

import (
	"context"
	"strconv"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/video/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavorlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorlistLogic {
	return &FavorlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavorlistLogic) Favorlist(req *types.VideoFavorListReq) (resp *types.VideoFavorListResp, err error) {
	// todo: add your logic here and delete this line

	// 发送给videofavorlist的rpc服务处理
	id, _ := strconv.ParseInt(req.User_id, 10, 64)
	videofavorlistResp, err := l.svcCtx.VideoRpc.VideoFavorList(l.ctx, &video.VideoFavorListReq{
		UserId: id,
	})
	if err != nil {
		return &types.VideoFavorListResp{
			Status_code: 1,
			Status_msg:  err.Error(),
			Video_list:  nil,
		}, nil
	}

	var list []types.Video
	for _, v := range videofavorlistResp.Videos {
		list = append(list, types.Video{
			Id: v.Id,
			Author: types.Author{
				Id:               v.Author.Id,
				Name:             v.Author.Name,
				Follow_count:     v.Author.FollowCount,
				Follower_count:   v.Author.FollowerCount,
				Is_follow:        v.Author.IsFollow,
				Avatar:           v.Author.Avatar,
				Background_image: v.Author.BackgroundImage,
				Signature:        v.Author.Signature,
				Total_favorited:  v.Author.TotalFavorited,
				Work_count:       v.Author.WorkCount,
				Favorite_count:   v.Author.FavoriteCount,
			},
			Play_url:       "http://10.0.2.2/" + v.PlayUrl,
			Cover_url:      "http://10.0.2.2/images/" + v.CoverUrl,
			Favorite_count: v.FavoriteCount,
			Comment_count:  v.CommentCount,
			Is_favorite:    v.IsFavorite,
			Title:          v.Title,
		})
	}
	return &types.VideoFavorListResp{Status_code: 0, Status_msg: "视频列表获取成功", Video_list: list}, nil
}
