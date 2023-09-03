// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	chat "douyin/apps/api/internal/handler/chat"
	relation "douyin/apps/api/internal/handler/relation"
	user "douyin/apps/api/internal/handler/user"
	video "douyin/apps/api/internal/handler/video"
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

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/publish/action",
				Handler: video.PublishHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/publish/list",
				Handler: video.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/feed",
				Handler: video.FeedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/favorite/action",
				Handler: video.FavoriteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/favorite/list",
				Handler: video.FavorlistHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/comment/action",
				Handler: video.PubcommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/comment/list",
				Handler: video.CommentlistHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/relation/action",
				Handler: relation.RelationactionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/follow/list",
				Handler: relation.RelationfollowHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/follower/list",
				Handler: relation.RelationfansHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/friend/list",
				Handler: relation.RelationfriendHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/message/action",
				Handler: chat.MessageactionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/message/chat",
				Handler: chat.MessagechatHandler(serverCtx),
			},
		},
	)
}
