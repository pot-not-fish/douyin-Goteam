package svc

import (
	"douyin/apps/api/internal/config"
	"douyin/apps/rpc/user/message"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc message.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: message.NewMessage(zrpc.MustNewClient(c.UserRpc)),
	}
}
