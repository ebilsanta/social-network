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
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Page  int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
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

func (x *GetFollowersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFollowersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetFollowersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []string                    `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Pagination *FollowerPaginationMetadata `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
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

func (x *GetFollowersResponse) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetFollowersResponse) GetPagination() *FollowerPaginationMetadata {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type GetFollowingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Page  int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
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

func (x *GetFollowingRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFollowingRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetFollowingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []string                    `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Pagination *FollowerPaginationMetadata `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
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

func (x *GetFollowingResponse) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetFollowingResponse) GetPagination() *FollowerPaginationMetadata {
	if x != nil {
		return x.Pagination
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

type FollowerPaginationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalRecords int32                  `protobuf:"varint,1,opt,name=totalRecords,proto3" json:"totalRecords,omitempty"`
	CurrentPage  int32                  `protobuf:"varint,2,opt,name=currentPage,proto3" json:"currentPage,omitempty"`
	TotalPages   int32                  `protobuf:"varint,3,opt,name=totalPages,proto3" json:"totalPages,omitempty"`
	NextPage     *wrapperspb.Int32Value `protobuf:"bytes,4,opt,name=nextPage,proto3" json:"nextPage,omitempty"`
	PrevPage     *wrapperspb.Int32Value `protobuf:"bytes,5,opt,name=prevPage,proto3" json:"prevPage,omitempty"`
}

func (x *FollowerPaginationMetadata) Reset() {
	*x = FollowerPaginationMetadata{}
	mi := &file_follower_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FollowerPaginationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerPaginationMetadata) ProtoMessage() {}

func (x *FollowerPaginationMetadata) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FollowerPaginationMetadata.ProtoReflect.Descriptor instead.
func (*FollowerPaginationMetadata) Descriptor() ([]byte, []int) {
	return file_follower_service_proto_rawDescGZIP(), []int{7}
}

func (x *FollowerPaginationMetadata) GetTotalRecords() int32 {
	if x != nil {
		return x.TotalRecords
	}
	return 0
}

func (x *FollowerPaginationMetadata) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *FollowerPaginationMetadata) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *FollowerPaginationMetadata) GetNextPage() *wrapperspb.Int32Value {
	if x != nil {
		return x.NextPage
	}
	return nil
}

func (x *FollowerPaginationMetadata) GetPrevPage() *wrapperspb.Int32Value {
	if x != nil {
		return x.PrevPage
	}
	return nil
}

var File_follower_service_proto protoreflect.FileDescriptor

var file_follower_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x78, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4c, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x78, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x4c, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x54, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x49, 0x44, 0x22, 0x57, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x44, 0x22, 0xf4,
	0x01, 0x0a, 0x1a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a,
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x08,
	0x70, 0x72, 0x65, 0x76, 0x50, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x70, 0x72, 0x65,
	0x76, 0x50, 0x61, 0x67, 0x65, 0x32, 0xbe, 0x03, 0x0a, 0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x5f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x12, 0x25, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x12, 0x25, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x12, 0x24, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x53, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x62, 0x69, 0x6c, 0x73, 0x61, 0x6e, 0x74, 0x61, 0x2f, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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
	(*AddUserRequest)(nil),             // 0: follower_service.AddUserRequest
	(*GetFollowersRequest)(nil),        // 1: follower_service.GetFollowersRequest
	(*GetFollowersResponse)(nil),       // 2: follower_service.GetFollowersResponse
	(*GetFollowingRequest)(nil),        // 3: follower_service.GetFollowingRequest
	(*GetFollowingResponse)(nil),       // 4: follower_service.GetFollowingResponse
	(*AddFollowerRequest)(nil),         // 5: follower_service.AddFollowerRequest
	(*DeleteFollowerRequest)(nil),      // 6: follower_service.DeleteFollowerRequest
	(*FollowerPaginationMetadata)(nil), // 7: follower_service.FollowerPaginationMetadata
	(*wrapperspb.Int32Value)(nil),      // 8: google.protobuf.Int32Value
	(*emptypb.Empty)(nil),              // 9: google.protobuf.Empty
}
var file_follower_service_proto_depIdxs = []int32{
	7, // 0: follower_service.GetFollowersResponse.pagination:type_name -> follower_service.FollowerPaginationMetadata
	7, // 1: follower_service.GetFollowingResponse.pagination:type_name -> follower_service.FollowerPaginationMetadata
	8, // 2: follower_service.FollowerPaginationMetadata.nextPage:type_name -> google.protobuf.Int32Value
	8, // 3: follower_service.FollowerPaginationMetadata.prevPage:type_name -> google.protobuf.Int32Value
	0, // 4: follower_service.FollowerService.AddUser:input_type -> follower_service.AddUserRequest
	1, // 5: follower_service.FollowerService.GetFollowers:input_type -> follower_service.GetFollowersRequest
	3, // 6: follower_service.FollowerService.GetFollowing:input_type -> follower_service.GetFollowingRequest
	5, // 7: follower_service.FollowerService.AddFollower:input_type -> follower_service.AddFollowerRequest
	6, // 8: follower_service.FollowerService.DeleteFollower:input_type -> follower_service.DeleteFollowerRequest
	9, // 9: follower_service.FollowerService.AddUser:output_type -> google.protobuf.Empty
	2, // 10: follower_service.FollowerService.GetFollowers:output_type -> follower_service.GetFollowersResponse
	4, // 11: follower_service.FollowerService.GetFollowing:output_type -> follower_service.GetFollowingResponse
	9, // 12: follower_service.FollowerService.AddFollower:output_type -> google.protobuf.Empty
	9, // 13: follower_service.FollowerService.DeleteFollower:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
