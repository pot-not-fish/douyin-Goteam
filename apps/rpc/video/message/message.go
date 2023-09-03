// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package message

import (
	"context"

	"douyin/apps/rpc/video/types/video"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Author             = video.Author
	CommentActionReq   = video.CommentActionReq
	CommentActionResp  = video.CommentActionResp
	CommentListReq     = video.CommentListReq
	CommentListResp    = video.CommentListResp
	VideoFavorListReq  = video.VideoFavorListReq
	VideoFavorListResp = video.VideoFavorListResp
	VideoFavoriteReq   = video.VideoFavoriteReq
	VideoFavoriteResp  = video.VideoFavoriteResp
	VideoFeedReq       = video.VideoFeedReq
	VideoFeedResp      = video.VideoFeedResp
	VideoList          = video.VideoList
	VideoListReq       = video.VideoListReq
	VideoListResp      = video.VideoListResp
	VideoMiddleList    = video.VideoMiddleList

	Message interface {
		VideoList(ctx context.Context, in *VideoListReq, opts ...grpc.CallOption) (*VideoListResp, error)
		VideoFeed(ctx context.Context, in *VideoFeedReq, opts ...grpc.CallOption) (*VideoFeedResp, error)
		VideoFavorite(ctx context.Context, in *VideoFavoriteReq, opts ...grpc.CallOption) (*VideoFavoriteResp, error)
		VideoFavorList(ctx context.Context, in *VideoFavorListReq, opts ...grpc.CallOption) (*VideoFavorListResp, error)
		CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error)
		CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error)
	}

	defaultMessage struct {
		cli zrpc.Client
	}
)

func NewMessage(cli zrpc.Client) Message {
	return &defaultMessage{
		cli: cli,
	}
}

func (m *defaultMessage) VideoList(ctx context.Context, in *VideoListReq, opts ...grpc.CallOption) (*VideoListResp, error) {
	client := video.NewMessageClient(m.cli.Conn())
	return client.VideoList(ctx, in, opts...)
}

func (m *defaultMessage) VideoFeed(ctx context.Context, in *VideoFeedReq, opts ...grpc.CallOption) (*VideoFeedResp, error) {
	client := video.NewMessageClient(m.cli.Conn())
	return client.VideoFeed(ctx, in, opts...)
}

func (m *defaultMessage) VideoFavorite(ctx context.Context, in *VideoFavoriteReq, opts ...grpc.CallOption) (*VideoFavoriteResp, error) {
	client := video.NewMessageClient(m.cli.Conn())
	return client.VideoFavorite(ctx, in, opts...)
}

func (m *defaultMessage) VideoFavorList(ctx context.Context, in *VideoFavorListReq, opts ...grpc.CallOption) (*VideoFavorListResp, error) {
	client := video.NewMessageClient(m.cli.Conn())
	return client.VideoFavorList(ctx, in, opts...)
}

func (m *defaultMessage) CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionResp, error) {
	client := video.NewMessageClient(m.cli.Conn())
	return client.CommentAction(ctx, in, opts...)
}

func (m *defaultMessage) CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListResp, error) {
	client := video.NewMessageClient(m.cli.Conn())
	return client.CommentList(ctx, in, opts...)
}