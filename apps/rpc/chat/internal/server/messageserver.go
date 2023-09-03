// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package server

import (
	"context"

	"douyin/apps/rpc/chat/internal/logic"
	"douyin/apps/rpc/chat/internal/svc"
	"douyin/apps/rpc/chat/types/chat"
)

type MessageServer struct {
	svcCtx *svc.ServiceContext
	chat.UnimplementedMessageServer
}

func NewMessageServer(svcCtx *svc.ServiceContext) *MessageServer {
	return &MessageServer{
		svcCtx: svcCtx,
	}
}

func (s *MessageServer) MessageAction(ctx context.Context, in *chat.MessageActionReq) (*chat.MessageActionResp, error) {
	l := logic.NewMessageActionLogic(ctx, s.svcCtx)
	return l.MessageAction(in)
}

func (s *MessageServer) MessageChat(ctx context.Context, in *chat.MessageChatReq) (*chat.MessageChatResp, error) {
	l := logic.NewMessageChatLogic(ctx, s.svcCtx)
	return l.MessageChat(in)
}
