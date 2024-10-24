// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: feed_service.proto

package proto

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

type GetFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetFeedRequest) Reset() {
	*x = GetFeedRequest{}
	mi := &file_feed_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedRequest) ProtoMessage() {}

func (x *GetFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_feed_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedRequest.ProtoReflect.Descriptor instead.
func (*GetFeedRequest) Descriptor() ([]byte, []int) {
	return file_feed_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetFeedRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetFeedRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetFeedRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts     []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	FromCache bool    `protobuf:"varint,2,opt,name=fromCache,proto3" json:"fromCache,omitempty"`
}

func (x *GetFeedResponse) Reset() {
	*x = GetFeedResponse{}
	mi := &file_feed_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedResponse) ProtoMessage() {}

func (x *GetFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_feed_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedResponse.ProtoReflect.Descriptor instead.
func (*GetFeedResponse) Descriptor() ([]byte, []int) {
	return file_feed_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetFeedResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *GetFeedResponse) GetFromCache() bool {
	if x != nil {
		return x.FromCache
	}
	return false
}

var File_feed_service_proto protoreflect.FileDescriptor

var file_feed_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x4c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x63, 0x68, 0x65, 0x32, 0x3d,
	0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x44, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x62, 0x69, 0x6c,
	0x73, 0x61, 0x6e, 0x74, 0x61, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x66, 0x65, 0x65,
	0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feed_service_proto_rawDescOnce sync.Once
	file_feed_service_proto_rawDescData = file_feed_service_proto_rawDesc
)

func file_feed_service_proto_rawDescGZIP() []byte {
	file_feed_service_proto_rawDescOnce.Do(func() {
		file_feed_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_feed_service_proto_rawDescData)
	})
	return file_feed_service_proto_rawDescData
}

var file_feed_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_feed_service_proto_goTypes = []any{
	(*GetFeedRequest)(nil),  // 0: GetFeedRequest
	(*GetFeedResponse)(nil), // 1: GetFeedResponse
	(*Post)(nil),            // 2: Post
}
var file_feed_service_proto_depIdxs = []int32{
	2, // 0: GetFeedResponse.posts:type_name -> Post
	0, // 1: FeedService.GetFeed:input_type -> GetFeedRequest
	1, // 2: FeedService.GetFeed:output_type -> GetFeedResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_feed_service_proto_init() }
func file_feed_service_proto_init() {
	if File_feed_service_proto != nil {
		return
	}
	file_post_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feed_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feed_service_proto_goTypes,
		DependencyIndexes: file_feed_service_proto_depIdxs,
		MessageInfos:      file_feed_service_proto_msgTypes,
	}.Build()
	File_feed_service_proto = out.File
	file_feed_service_proto_rawDesc = nil
	file_feed_service_proto_goTypes = nil
	file_feed_service_proto_depIdxs = nil
}
