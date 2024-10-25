// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: follower_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddUserRequest) Reset() {
	*x = AddUserRequest{}
	mi := &file_follower_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRequest) ProtoMessage() {}

func (x *AddUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRequest.ProtoReflect.Descriptor instead.
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFollowersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFollowersRequest) Reset() {
	*x = GetFollowersRequest{}
	mi := &file_follower_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFollowersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersRequest) ProtoMessage() {}

func (x *GetFollowersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersRequest.ProtoReflect.Descriptor instead.
func (*GetFollowersRequest) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetFollowersRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFollowersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followers []*GraphUser `protobuf:"bytes,1,rep,name=followers,proto3" json:"followers,omitempty"`
}

func (x *GetFollowersResponse) Reset() {
	*x = GetFollowersResponse{}
	mi := &file_follower_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFollowersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowersResponse) ProtoMessage() {}

func (x *GetFollowersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowersResponse.ProtoReflect.Descriptor instead.
func (*GetFollowersResponse) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetFollowersResponse) GetFollowers() []*GraphUser {
	if x != nil {
		return x.Followers
	}
	return nil
}

type GetFollowingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFollowingRequest) Reset() {
	*x = GetFollowingRequest{}
	mi := &file_follower_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFollowingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingRequest) ProtoMessage() {}

func (x *GetFollowingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingRequest.ProtoReflect.Descriptor instead.
func (*GetFollowingRequest) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetFollowingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFollowingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Following []*GraphUser `protobuf:"bytes,1,rep,name=following,proto3" json:"following,omitempty"`
}

func (x *GetFollowingResponse) Reset() {
	*x = GetFollowingResponse{}
	mi := &file_follower_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFollowingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowingResponse) ProtoMessage() {}

func (x *GetFollowingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowingResponse.ProtoReflect.Descriptor instead.
func (*GetFollowingResponse) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetFollowingResponse) GetFollowing() []*GraphUser {
	if x != nil {
		return x.Following
	}
	return nil
}

type AddFollowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerID string `protobuf:"bytes,1,opt,name=followerID,proto3" json:"followerID,omitempty"`
	FollowedID string `protobuf:"bytes,2,opt,name=followedID,proto3" json:"followedID,omitempty"`
}

func (x *AddFollowerRequest) Reset() {
	*x = AddFollowerRequest{}
	mi := &file_follower_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddFollowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFollowerRequest) ProtoMessage() {}

func (x *AddFollowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFollowerRequest.ProtoReflect.Descriptor instead.
func (*AddFollowerRequest) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{5}
}

func (x *AddFollowerRequest) GetFollowerID() string {
	if x != nil {
		return x.FollowerID
	}
	return ""
}

func (x *AddFollowerRequest) GetFollowedID() string {
	if x != nil {
		return x.FollowedID
	}
	return ""
}

type DeleteFollowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerID string `protobuf:"bytes,1,opt,name=followerID,proto3" json:"followerID,omitempty"`
	FollowedID string `protobuf:"bytes,2,opt,name=followedID,proto3" json:"followedID,omitempty"`
}

func (x *DeleteFollowerRequest) Reset() {
	*x = DeleteFollowerRequest{}
	mi := &file_follower_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFollowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFollowerRequest) ProtoMessage() {}

func (x *DeleteFollowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFollowerRequest.ProtoReflect.Descriptor instead.
func (*DeleteFollowerRequest) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFollowerRequest) GetFollowerID() string {
	if x != nil {
		return x.FollowerID
	}
	return ""
}

func (x *DeleteFollowerRequest) GetFollowedID() string {
	if x != nil {
		return x.FollowedID
	}
	return ""
}

type GraphUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GraphUser) Reset() {
	*x = GraphUser{}
	mi := &file_follower_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GraphUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphUser) ProtoMessage() {}

func (x *GraphUser) ProtoReflect() protoreflect.Message {
	mi := &file_follower_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphUser.ProtoReflect.Descriptor instead.
func (*GraphUser) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{7}
}

func (x *GraphUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_follower_service_proto protoreflect.FileDescriptor

var file_follower_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x51, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x22, 0x54, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x49, 0x44, 0x22, 0x57, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x44, 0x22, 0x1b, 0x0a, 0x09,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xbe, 0x03, 0x0a, 0x0f, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x62, 0x69, 0x6c, 0x73, 0x61, 0x6e,
	0x74, 0x61, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follower_service_proto_rawDescOnce sync.Once
	file_follower_service_proto_rawDescData = file_follower_service_proto_rawDesc
)

func file_follower_service_proto_rawDescGZIP() []byte {
	file_follower_service_proto_rawDescOnce.Do(func() {
		file_follower_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_follower_service_proto_rawDescData)
	})
	return file_follower_service_proto_rawDescData
}

var file_follower_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_follower_service_proto_goTypes = []any{
	(*AddUserRequest)(nil),        // 0: follower_service.AddUserRequest
	(*GetFollowersRequest)(nil),   // 1: follower_service.GetFollowersRequest
	(*GetFollowersResponse)(nil),  // 2: follower_service.GetFollowersResponse
	(*GetFollowingRequest)(nil),   // 3: follower_service.GetFollowingRequest
	(*GetFollowingResponse)(nil),  // 4: follower_service.GetFollowingResponse
	(*AddFollowerRequest)(nil),    // 5: follower_service.AddFollowerRequest
	(*DeleteFollowerRequest)(nil), // 6: follower_service.DeleteFollowerRequest
	(*GraphUser)(nil),             // 7: follower_service.GraphUser
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_follower_service_proto_depIdxs = []int32{
	7, // 0: follower_service.GetFollowersResponse.followers:type_name -> follower_service.GraphUser
	7, // 1: follower_service.GetFollowingResponse.following:type_name -> follower_service.GraphUser
	0, // 2: follower_service.FollowerService.AddUser:input_type -> follower_service.AddUserRequest
	1, // 3: follower_service.FollowerService.GetFollowers:input_type -> follower_service.GetFollowersRequest
	3, // 4: follower_service.FollowerService.GetFollowing:input_type -> follower_service.GetFollowingRequest
	5, // 5: follower_service.FollowerService.AddFollower:input_type -> follower_service.AddFollowerRequest
	6, // 6: follower_service.FollowerService.DeleteFollower:input_type -> follower_service.DeleteFollowerRequest
	8, // 7: follower_service.FollowerService.AddUser:output_type -> google.protobuf.Empty
	2, // 8: follower_service.FollowerService.GetFollowers:output_type -> follower_service.GetFollowersResponse
	4, // 9: follower_service.FollowerService.GetFollowing:output_type -> follower_service.GetFollowingResponse
	8, // 10: follower_service.FollowerService.AddFollower:output_type -> google.protobuf.Empty
	8, // 11: follower_service.FollowerService.DeleteFollower:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_follower_service_proto_init() }
func file_follower_service_proto_init() {
	if File_follower_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_follower_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follower_service_proto_goTypes,
		DependencyIndexes: file_follower_service_proto_depIdxs,
		MessageInfos:      file_follower_service_proto_msgTypes,
	}.Build()
	File_follower_service_proto = out.File
	file_follower_service_proto_rawDesc = nil
	file_follower_service_proto_goTypes = nil
	file_follower_service_proto_depIdxs = nil
}
