package video

import (
	"context"
	"strconv"
	"time"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/video/types/video"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.VideoFeedReq) (resp *types.VideoFeedResp, err error) {
	// todo: add your logic here and delete this line

	// 如果返回时间戳为空，则按照当前时间判断
	if len(req.Latest_time) == 0 {
		t := strconv.FormatInt(time.Now().Unix(), 10)
		req.Latest_time = t
	}

	// 如果token为空，则默认id为0
	var id int64
	if len(req.Token) == 0 {
		id = 0
	} else {
		id, err = pkg.AuthToken(req.Token)
		if err != nil {
			return &types.VideoFeedResp{
				Status_code: 1,
				Status_msg:  err.Error(),
			}, nil
		}
	}

	// 发送给rpc处理
	l_t, _ := strconv.ParseInt(req.Latest_time, 10, 64)
	videofeedResp, err := l.svcCtx.VideoRpc.VideoFeed(l.ctx, &video.VideoFeedReq{LatestTime: l_t, UserId: id})
	if err != nil {
		return &types.VideoFeedResp{
			Status_code: 1,
			Status_msg:  err.Error(),
		}, nil
	}

	var videolist []types.Video
	for _, v := range videofeedResp.Videos {
		videolist = append(videolist, types.Video{
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
	return &types.VideoFeedResp{
		Status_code: 0,
		Status_msg:  "获取视频流成功",
		Next_time:   int32(videofeedResp.NextTime),
		Video_list:  videolist,
	}, nil
}
