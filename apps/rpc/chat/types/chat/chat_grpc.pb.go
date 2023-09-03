// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Message_MessageAction_FullMethodName = "/chat.Message/MessageAction"
	Message_MessageChat_FullMethodName   = "/chat.Message/MessageChat"
)

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageClient interface {
	MessageAction(ctx context.Context, in *MessageActionReq, opts ...grpc.CallOption) (*MessageActionResp, error)
	MessageChat(ctx context.Context, in *MessageChatReq, opts ...grpc.CallOption) (*MessageChatResp, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) MessageAction(ctx context.Context, in *MessageActionReq, opts ...grpc.CallOption) (*MessageActionResp, error) {
	out := new(MessageActionResp)
	err := c.cc.Invoke(ctx, Message_MessageAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) MessageChat(ctx context.Context, in *MessageChatReq, opts ...grpc.CallOption) (*MessageChatResp, error) {
	out := new(MessageChatResp)
	err := c.cc.Invoke(ctx, Message_MessageChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
// All implementations must embed UnimplementedMessageServer
// for forward compatibility
type MessageServer interface {
	MessageAction(context.Context, *MessageActionReq) (*MessageActionResp, error)
	MessageChat(context.Context, *MessageChatReq) (*MessageChatResp, error)
	mustEmbedUnimplementedMessageServer()
}

// UnimplementedMessageServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (UnimplementedMessageServer) MessageAction(context.Context, *MessageActionReq) (*MessageActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageAction not implemented")
}
func (UnimplementedMessageServer) MessageChat(context.Context, *MessageChatReq) (*MessageChatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageChat not implemented")
}
func (UnimplementedMessageServer) mustEmbedUnimplementedMessageServer() {}

// UnsafeMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServer will
// result in compilation errors.
type UnsafeMessageServer interface {
	mustEmbedUnimplementedMessageServer()
}

func RegisterMessageServer(s grpc.ServiceRegistrar, srv MessageServer) {
	s.RegisterService(&Message_ServiceDesc, srv)
}

func _Message_MessageAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).MessageAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_MessageAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).MessageAction(ctx, req.(*MessageActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_MessageChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).MessageChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Message_MessageChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).MessageChat(ctx, req.(*MessageChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Message_ServiceDesc is the grpc.ServiceDesc for Message service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Message_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageAction",
			Handler:    _Message_MessageAction_Handler,
		},
		{
			MethodName: "MessageChat",
			Handler:    _Message_MessageChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
