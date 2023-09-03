package chat

import (
	"context"
	"strconv"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/chat/types/chat"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageactionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageactionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageactionLogic {
	return &MessageactionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageactionLogic) Messageaction(req *types.MessageActionReq) (resp *types.MessageActionResp, err error) {
	// todo: add your logic here and delete this line

	// token鉴权
	id, err := pkg.AuthToken(req.Token)
	if err != nil {
		return &types.MessageActionResp{
			Status_code: 1,
			Status_msg:  err.Error(),
		}, nil
	}

	// 发送给messageaction的rpc服务处理
	actiontype, _ := strconv.ParseInt(req.Action_type, 10, 64)
	messageactionResp, err := l.svcCtx.ChatRpc.MessageAction(l.ctx, &chat.MessageActionReq{
		UserId:     id,
		ToUserId:   req.To_user_id,
		ActionType: int32(actiontype),
		Content:    req.Content,
	})
	if err != nil {
		return &types.MessageActionResp{
			Status_code: 1,
			Status_msg:  err.Error(),
		}, nil
	} else {
		return &types.MessageActionResp{
			Status_code: 0,
			Status_msg:  messageactionResp.StatusMsg,
		}, nil
	}
}
