// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: relation.proto

package relation

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RelationActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ToUserId   int64 `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`
	ActionType int32 `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
}

func (x *RelationActionReq) Reset() {
	*x = RelationActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionReq) ProtoMessage() {}

func (x *RelationActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionReq.ProtoReflect.Descriptor instead.
func (*RelationActionReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{0}
}

func (x *RelationActionReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationActionReq) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *RelationActionReq) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type RelationActionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusMsg string `protobuf:"bytes,1,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *RelationActionResp) Reset() {
	*x = RelationActionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationActionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationActionResp) ProtoMessage() {}

func (x *RelationActionResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationActionResp.ProtoReflect.Descriptor instead.
func (*RelationActionResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{1}
}

func (x *RelationActionResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type RelationFollowReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MeId   int64 `protobuf:"varint,2,opt,name=me_id,json=meId,proto3" json:"me_id,omitempty"`
}

func (x *RelationFollowReq) Reset() {
	*x = RelationFollowReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowReq) ProtoMessage() {}

func (x *RelationFollowReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowReq.ProtoReflect.Descriptor instead.
func (*RelationFollowReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{2}
}

func (x *RelationFollowReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationFollowReq) GetMeId() int64 {
	if x != nil {
		return x.MeId
	}
	return 0
}

type RelationFollowResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*User `protobuf:"bytes,1,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *RelationFollowResp) Reset() {
	*x = RelationFollowResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFollowResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFollowResp) ProtoMessage() {}

func (x *RelationFollowResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFollowResp.ProtoReflect.Descriptor instead.
func (*RelationFollowResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{3}
}

func (x *RelationFollowResp) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type RelationFansReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MeId   int64 `protobuf:"varint,2,opt,name=me_id,json=meId,proto3" json:"me_id,omitempty"`
}

func (x *RelationFansReq) Reset() {
	*x = RelationFansReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFansReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFansReq) ProtoMessage() {}

func (x *RelationFansReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFansReq.ProtoReflect.Descriptor instead.
func (*RelationFansReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{4}
}

func (x *RelationFansReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RelationFansReq) GetMeId() int64 {
	if x != nil {
		return x.MeId
	}
	return 0
}

type RelationFansResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*User `protobuf:"bytes,1,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *RelationFansResp) Reset() {
	*x = RelationFansResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFansResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFansResp) ProtoMessage() {}

func (x *RelationFansResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFansResp.ProtoReflect.Descriptor instead.
func (*RelationFansResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{5}
}

func (x *RelationFansResp) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type RelationFriendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RelationFriendReq) Reset() {
	*x = RelationFriendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFriendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFriendReq) ProtoMessage() {}

func (x *RelationFriendReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFriendReq.ProtoReflect.Descriptor instead.
func (*RelationFriendReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{6}
}

func (x *RelationFriendReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RelationFriendResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*FriendUser `protobuf:"bytes,1,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *RelationFriendResp) Reset() {
	*x = RelationFriendResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationFriendResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationFriendResp) ProtoMessage() {}

func (x *RelationFriendResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationFriendResp.ProtoReflect.Descriptor instead.
func (*RelationFriendResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{7}
}

func (x *RelationFriendResp) GetUserList() []*FriendUser {
	if x != nil {
		return x.UserList
	}
	return nil
}

type FriendUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FollowCount     int64   `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`
	FollowerCount   int64   `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	IsFollow        bool    `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
	Avatar          string  `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	BackgroundImage string  `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"`
	Signature       string  `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	TotalFavorited  string  `protobuf:"bytes,9,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"`
	WorkCount       int64   `protobuf:"varint,10,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`
	FavoriteCount   int64   `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`
	Message         *string `protobuf:"bytes,12,opt,name=message,proto3,oneof" json:"message,omitempty"`
	MsgType         *int64  `protobuf:"varint,13,opt,name=msg_type,json=msgType,proto3,oneof" json:"msg_type,omitempty"`
}

func (x *FriendUser) Reset() {
	*x = FriendUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendUser) ProtoMessage() {}

func (x *FriendUser) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendUser.ProtoReflect.Descriptor instead.
func (*FriendUser) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{8}
}

func (x *FriendUser) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FriendUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FriendUser) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *FriendUser) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *FriendUser) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *FriendUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *FriendUser) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *FriendUser) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *FriendUser) GetTotalFavorited() string {
	if x != nil {
		return x.TotalFavorited
	}
	return ""
}

func (x *FriendUser) GetWorkCount() int64 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

func (x *FriendUser) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *FriendUser) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *FriendUser) GetMsgType() int64 {
	if x != nil && x.MsgType != nil {
		return *x.MsgType
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FollowCount     int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`
	FollowerCount   int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	IsFollow        bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
	Avatar          string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	BackgroundImage string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"`
	Signature       string `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	TotalFavorited  string `protobuf:"bytes,9,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"`
	WorkCount       int64  `protobuf:"varint,10,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`
	FavoriteCount   int64  `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{9}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *User) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *User) GetTotalFavorited() string {
	if x != nil {
		return x.TotalFavorited
	}
	return ""
}

func (x *User) GetWorkCount() int64 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

func (x *User) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

var File_relation_proto protoreflect.FileDescriptor

var file_relation_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6b, 0x0a, 0x11, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x33, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x41, 0x0a, 0x11,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6d, 0x65, 0x49, 0x64, 0x22,
	0x41, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x3f, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x13,
	0x0a, 0x05, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6d,
	0x65, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x47, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xbf, 0x03, 0x0a, 0x0a,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x29, 0x0a,
	0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe1, 0x02,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0xb7, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a,
	0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4b, 0x0a, 0x0e, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1b, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x61, 0x6e, 0x73, 0x12, 0x19, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4b,
	0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x12, 0x1b, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_relation_proto_rawDescOnce sync.Once
	file_relation_proto_rawDescData = file_relation_proto_rawDesc
)

func file_relation_proto_rawDescGZIP() []byte {
	file_relation_proto_rawDescOnce.Do(func() {
		file_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_relation_proto_rawDescData)
	})
	return file_relation_proto_rawDescData
}

var file_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_relation_proto_goTypes = []interface{}{
	(*RelationActionReq)(nil),  // 0: relation.RelationActionReq
	(*RelationActionResp)(nil), // 1: relation.RelationActionResp
	(*RelationFollowReq)(nil),  // 2: relation.RelationFollowReq
	(*RelationFollowResp)(nil), // 3: relation.RelationFollowResp
	(*RelationFansReq)(nil),    // 4: relation.RelationFansReq
	(*RelationFansResp)(nil),   // 5: relation.RelationFansResp
	(*RelationFriendReq)(nil),  // 6: relation.RelationFriendReq
	(*RelationFriendResp)(nil), // 7: relation.RelationFriendResp
	(*FriendUser)(nil),         // 8: relation.FriendUser
	(*User)(nil),               // 9: relation.User
}
var file_relation_proto_depIdxs = []int32{
	9, // 0: relation.RelationFollowResp.user_list:type_name -> relation.User
	9, // 1: relation.RelationFansResp.user_list:type_name -> relation.User
	8, // 2: relation.RelationFriendResp.user_list:type_name -> relation.FriendUser
	0, // 3: relation.Message.RelationAction:input_type -> relation.RelationActionReq
	2, // 4: relation.Message.RelationFollow:input_type -> relation.RelationFollowReq
	4, // 5: relation.Message.RelationFans:input_type -> relation.RelationFansReq
	6, // 6: relation.Message.RelationFriend:input_type -> relation.RelationFriendReq
	1, // 7: relation.Message.RelationAction:output_type -> relation.RelationActionResp
	3, // 8: relation.Message.RelationFollow:output_type -> relation.RelationFollowResp
	5, // 9: relation.Message.RelationFans:output_type -> relation.RelationFansResp
	7, // 10: relation.Message.RelationFriend:output_type -> relation.RelationFriendResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_relation_proto_init() }
func file_relation_proto_init() {
	if File_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationActionResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFollowResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFansReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFansResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFriendReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationFriendResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_relation_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relation_proto_goTypes,
		DependencyIndexes: file_relation_proto_depIdxs,
		MessageInfos:      file_relation_proto_msgTypes,
	}.Build()
	File_relation_proto = out.File
	file_relation_proto_rawDesc = nil
	file_relation_proto_goTypes = nil
	file_relation_proto_depIdxs = nil
}
