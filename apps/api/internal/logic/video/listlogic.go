package video

import (
	"context"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/video/types/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.VideoListReq) (resp *types.VideoListResp, err error) {
	// todo: add your logic here and delete this line

	// 发送给rpc的videolist处理
	videolistResp, err := l.svcCtx.VideoRpc.VideoList(l.ctx, &video.VideoListReq{UserId: req.User_id})
	if err != nil {
		return &types.VideoListResp{Status_code: 1, Status_msg: err.Error(), Video_list: nil}, nil
	} else {
		var list []types.Video
		for _, v := range videolistResp.Videos {
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
		return &types.VideoListResp{Status_code: 0, Status_msg: "视频列表获取成功", Video_list: list}, nil
	}
}
