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

type CommentlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentlistLogic {
	return &CommentlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentlistLogic) Commentlist(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	// todo: add your logic here and delete this line

	// token鉴权
	var id int64
	if len(req.Token) == 0 {
		id = 0
	} else {
		id, err = pkg.AuthToken(req.Token)
		if err != nil {
			return &types.CommentListResp{
				Status_code:  1,
				Status_msg:   err.Error(),
				Comment_list: nil,
			}, err
		}
	}

	// 发送给commentlist的rpc服务处理
	videoid, _ := strconv.ParseInt(req.Video_id, 10, 64)
	commentlistResp, err := l.svcCtx.VideoRpc.CommentList(l.ctx, &video.CommentListReq{VideoId: videoid, UserId: id})
	if err != nil {
		return nil, err
	}

	var comments []types.Comments
	for _, v := range commentlistResp.Comments {
		comments = append(comments, types.Comments{
			Id: v.Id,
			User: types.Author{
				Id:               v.User.Id,
				Name:             v.User.Name,
				Follow_count:     v.User.FollowCount,
				Follower_count:   v.User.FollowerCount,
				Is_follow:        v.User.IsFollow,
				Avatar:           v.User.Avatar,
				Background_image: v.User.BackgroundImage,
				Signature:        v.User.Signature,
				Total_favorited:  v.User.TotalFavorited,
				Work_count:       v.User.WorkCount,
				Favorite_count:   v.User.FavoriteCount,
			},
			Content:     v.Content,
			Create_date: v.CreateDate,
		})
	}
	return &types.CommentListResp{
		Status_code:  0,
		Status_msg:   "获取评论列表成功",
		Comment_list: comments,
	}, err
}
