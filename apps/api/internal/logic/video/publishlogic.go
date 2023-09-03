package video

import (
	"context"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.VideoPublishReq) (resp *types.VideoPublishResp, err error) {
	// todo: add your logic here and delete this line
	return &types.VideoPublishResp{Status_code: 0, Status_msg: "上传成功"}, nil
}
