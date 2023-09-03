package chat

import (
	"context"

	"douyin/apps/api/internal/svc"
	"douyin/apps/api/internal/types"
	"douyin/apps/rpc/chat/types/chat"
	"douyin/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessagechatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessagechatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessagechatLogic {
	return &MessagechatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessagechatLogic) Messagechat(req *types.MessageChatReq) (resp *types.MessageChatResp, err error) {
	// todo: add your logic here and delete this line

	// token鉴权
	id, err := pkg.AuthToken(req.Token)
	if err != nil {
		return &types.MessageChatResp{
			Status_code:  1,
			Status_msg:   err.Error(),
			Message_list: nil,
		}, nil
	}

	// 发送给messagechat的rpc服务处理
	messagechatResp, err := l.svcCtx.ChatRpc.MessageChat(l.ctx, &chat.MessageChatReq{
		UserId:     id,
		ToUserId:   req.To_user_id,
		PreMsgTime: req.Pre_msg_time,
	})
	if err != nil {
		return &types.MessageChatResp{
			Status_code:  1,
			Status_msg:   err.Error(),
			Message_list: nil,
		}, nil
	}

	var chatlist []types.MessageList
	for _, v := range messagechatResp.MessageList {
		chatlist = append(chatlist, types.MessageList{
			Id:           v.Id,
			To_user_id:   v.ToUserId,
			From_user_id: v.FromUserId,
			Content:      v.Content,
			Create_time:  v.CreateTime,
		})
	}
	return &types.MessageChatResp{
		Status_code:  0,
		Status_msg:   "获取聊天信息列表成功",
		Message_list: chatlist,
	}, nil
}
