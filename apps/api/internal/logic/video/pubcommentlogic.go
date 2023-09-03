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

type PubcommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPubcommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PubcommentLogic {
	return &PubcommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PubcommentLogic) Pubcomment(req *types.CommentActionReq) (resp *types.CommentActionResp, err error) {
	// todo: add your logic here and delete this line

	// token鉴权
	id, err := pkg.AuthToken(req.Token)
	if err != nil {
		return &types.CommentActionResp{
			Status_code: 1,
			Status_msg:  err.Error(),
			Comment:     types.Comments{},
		}, nil
	}

	// 发送给pubcomment的rpc服务处理
	videoid, _ := strconv.ParseInt(req.Video_id, 10, 64)
	actiontype, _ := strconv.ParseInt(req.Action_type, 10, 64)
	commentid, _ := strconv.ParseInt(req.Comment_id, 10, 64)
	commentactionresp, err := l.svcCtx.VideoRpc.CommentAction(l.ctx, &video.CommentActionReq{
		VideoId:     videoid,
		UserId:      id,
		ActionType:  int32(actiontype),
		CommentText: req.Comment_text,
		CommentId:   commentid,
	})

	if err != nil {
		return &types.CommentActionResp{
			Status_code: 1,
			Status_msg:  err.Error(),
			Comment:     types.Comments{},
		}, nil
	}

	if actiontype == 1 {
		return &types.CommentActionResp{
			Status_code: 0,
			Status_msg:  "发布评论成功",
			Comment: types.Comments{
				Id: commentactionresp.Id,
				User: types.Author{
					Id:               commentactionresp.User.Id,
					Name:             commentactionresp.User.Name,
					Follow_count:     commentactionresp.User.FollowCount,
					Follower_count:   commentactionresp.User.FollowerCount,
					Is_follow:        commentactionresp.User.IsFollow,
					Avatar:           commentactionresp.User.Avatar,
					Background_image: commentactionresp.User.BackgroundImage,
					Signature:        commentactionresp.User.Signature,
					Total_favorited:  commentactionresp.User.TotalFavorited,
					Work_count:       commentactionresp.User.WorkCount,
					Favorite_count:   commentactionresp.User.FavoriteCount,
				},
				Content:     commentactionresp.Content,
				Create_date: commentactionresp.CreateDate,
			},
		}, nil
	} else {
		return &types.CommentActionResp{
			Status_code: 0,
			Status_msg:  "删除评论成功",
			Comment:     types.Comments{},
		}, nil
	}

}
