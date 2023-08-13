// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "douyin/apps/api/internal/handler/user"
	"douyin/apps/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/user",
				Handler: user.UserinfoHandler(serverCtx),
			},
		},
	)
}