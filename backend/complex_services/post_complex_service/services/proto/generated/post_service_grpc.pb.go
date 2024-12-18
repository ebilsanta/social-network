// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: post_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PostService_CreatePost_FullMethodName        = "/post_service.PostService/CreatePost"
	PostService_GetPost_FullMethodName           = "/post_service.PostService/GetPost"
	PostService_GetPosts_FullMethodName          = "/post_service.PostService/GetPosts"
	PostService_GetPostsByPostIds_FullMethodName = "/post_service.PostService/GetPostsByPostIds"
	PostService_GetPostsByUserId_FullMethodName  = "/post_service.PostService/GetPostsByUserId"
	PostService_GetPostsByUserIds_FullMethodName = "/post_service.PostService/GetPostsByUserIds"
)

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error)
	GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error)
	GetPostsByPostIds(ctx context.Context, in *GetPostsByIdsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error)
	GetPostsByUserId(ctx context.Context, in *GetPostsByUserRequest, opts ...grpc.CallOption) (*GetPostsPaginatedResponse, error)
	GetPostsByUserIds(ctx context.Context, in *GetPostsByUsersRequest, opts ...grpc.CallOption) (*GetPostsResponse, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, PostService_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*Post, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Post)
	err := c.cc.Invoke(ctx, PostService_GetPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostsResponse)
	err := c.cc.Invoke(ctx, PostService_GetPosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostsByPostIds(ctx context.Context, in *GetPostsByIdsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostsResponse)
	err := c.cc.Invoke(ctx, PostService_GetPostsByPostIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostsByUserId(ctx context.Context, in *GetPostsByUserRequest, opts ...grpc.CallOption) (*GetPostsPaginatedResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostsPaginatedResponse)
	err := c.cc.Invoke(ctx, PostService_GetPostsByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostsByUserIds(ctx context.Context, in *GetPostsByUsersRequest, opts ...grpc.CallOption) (*GetPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostsResponse)
	err := c.cc.Invoke(ctx, PostService_GetPostsByUserIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility.
type PostServiceServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*Post, error)
	GetPost(context.Context, *GetPostRequest) (*Post, error)
	GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error)
	GetPostsByPostIds(context.Context, *GetPostsByIdsRequest) (*GetPostsResponse, error)
	GetPostsByUserId(context.Context, *GetPostsByUserRequest) (*GetPostsPaginatedResponse, error)
	GetPostsByUserIds(context.Context, *GetPostsByUsersRequest) (*GetPostsResponse, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPostServiceServer struct{}

func (UnimplementedPostServiceServer) CreatePost(context.Context, *CreatePostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedPostServiceServer) GetPost(context.Context, *GetPostRequest) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedPostServiceServer) GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedPostServiceServer) GetPostsByPostIds(context.Context, *GetPostsByIdsRequest) (*GetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsByPostIds not implemented")
}
func (UnimplementedPostServiceServer) GetPostsByUserId(context.Context, *GetPostsByUserRequest) (*GetPostsPaginatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsByUserId not implemented")
}
func (UnimplementedPostServiceServer) GetPostsByUserIds(context.Context, *GetPostsByUsersRequest) (*GetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostsByUserIds not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}
func (UnimplementedPostServiceServer) testEmbeddedByValue()                     {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	// If the following call pancis, it indicates UnimplementedPostServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPosts(ctx, req.(*GetPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostsByPostIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostsByPostIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPostsByPostIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostsByPostIds(ctx, req.(*GetPostsByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPostsByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostsByUserId(ctx, req.(*GetPostsByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostsByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsByUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostsByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PostService_GetPostsByUserIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostsByUserIds(ctx, req.(*GetPostsByUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post_service.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _PostService_GetPost_Handler,
		},
		{
			MethodName: "GetPosts",
			Handler:    _PostService_GetPosts_Handler,
		},
		{
			MethodName: "GetPostsByPostIds",
			Handler:    _PostService_GetPostsByPostIds_Handler,
		},
		{
			MethodName: "GetPostsByUserId",
			Handler:    _PostService_GetPostsByUserId_Handler,
		},
		{
			MethodName: "GetPostsByUserIds",
			Handler:    _PostService_GetPostsByUserIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post_service.proto",
}
