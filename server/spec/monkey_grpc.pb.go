// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package spec

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

// MonkeyClient is the client API for Monkey service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonkeyClient interface {
	ListArtists(ctx context.Context, in *ArtistsRequest, opts ...grpc.CallOption) (*ArtistsReply, error)
	ListAlbums(ctx context.Context, in *AlbumsRequest, opts ...grpc.CallOption) (*AlbumsReply, error)
	ListAllAlbums(ctx context.Context, in *AllAlbumsRequest, opts ...grpc.CallOption) (*AllAlbumsReply, error)
	ListSongs(ctx context.Context, in *SongsRequest, opts ...grpc.CallOption) (*SongsReply, error)
}

type monkeyClient struct {
	cc grpc.ClientConnInterface
}

func NewMonkeyClient(cc grpc.ClientConnInterface) MonkeyClient {
	return &monkeyClient{cc}
}

func (c *monkeyClient) ListArtists(ctx context.Context, in *ArtistsRequest, opts ...grpc.CallOption) (*ArtistsReply, error) {
	out := new(ArtistsReply)
	err := c.cc.Invoke(ctx, "/spec.Monkey/ListArtists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monkeyClient) ListAlbums(ctx context.Context, in *AlbumsRequest, opts ...grpc.CallOption) (*AlbumsReply, error) {
	out := new(AlbumsReply)
	err := c.cc.Invoke(ctx, "/spec.Monkey/ListAlbums", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monkeyClient) ListAllAlbums(ctx context.Context, in *AllAlbumsRequest, opts ...grpc.CallOption) (*AllAlbumsReply, error) {
	out := new(AllAlbumsReply)
	err := c.cc.Invoke(ctx, "/spec.Monkey/ListAllAlbums", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monkeyClient) ListSongs(ctx context.Context, in *SongsRequest, opts ...grpc.CallOption) (*SongsReply, error) {
	out := new(SongsReply)
	err := c.cc.Invoke(ctx, "/spec.Monkey/ListSongs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonkeyServer is the server API for Monkey service.
// All implementations must embed UnimplementedMonkeyServer
// for forward compatibility
type MonkeyServer interface {
	ListArtists(context.Context, *ArtistsRequest) (*ArtistsReply, error)
	ListAlbums(context.Context, *AlbumsRequest) (*AlbumsReply, error)
	ListAllAlbums(context.Context, *AllAlbumsRequest) (*AllAlbumsReply, error)
	ListSongs(context.Context, *SongsRequest) (*SongsReply, error)
	mustEmbedUnimplementedMonkeyServer()
}

// UnimplementedMonkeyServer must be embedded to have forward compatible implementations.
type UnimplementedMonkeyServer struct {
}

func (UnimplementedMonkeyServer) ListArtists(context.Context, *ArtistsRequest) (*ArtistsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtists not implemented")
}
func (UnimplementedMonkeyServer) ListAlbums(context.Context, *AlbumsRequest) (*AlbumsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlbums not implemented")
}
func (UnimplementedMonkeyServer) ListAllAlbums(context.Context, *AllAlbumsRequest) (*AllAlbumsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllAlbums not implemented")
}
func (UnimplementedMonkeyServer) ListSongs(context.Context, *SongsRequest) (*SongsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSongs not implemented")
}
func (UnimplementedMonkeyServer) mustEmbedUnimplementedMonkeyServer() {}

// UnsafeMonkeyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonkeyServer will
// result in compilation errors.
type UnsafeMonkeyServer interface {
	mustEmbedUnimplementedMonkeyServer()
}

func RegisterMonkeyServer(s grpc.ServiceRegistrar, srv MonkeyServer) {
	s.RegisterService(&Monkey_ServiceDesc, srv)
}

func _Monkey_ListArtists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonkeyServer).ListArtists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.Monkey/ListArtists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonkeyServer).ListArtists(ctx, req.(*ArtistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monkey_ListAlbums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonkeyServer).ListAlbums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.Monkey/ListAlbums",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonkeyServer).ListAlbums(ctx, req.(*AlbumsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monkey_ListAllAlbums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllAlbumsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonkeyServer).ListAllAlbums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.Monkey/ListAllAlbums",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonkeyServer).ListAllAlbums(ctx, req.(*AllAlbumsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monkey_ListSongs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonkeyServer).ListSongs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spec.Monkey/ListSongs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonkeyServer).ListSongs(ctx, req.(*SongsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Monkey_ServiceDesc is the grpc.ServiceDesc for Monkey service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Monkey_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spec.Monkey",
	HandlerType: (*MonkeyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListArtists",
			Handler:    _Monkey_ListArtists_Handler,
		},
		{
			MethodName: "ListAlbums",
			Handler:    _Monkey_ListAlbums_Handler,
		},
		{
			MethodName: "ListAllAlbums",
			Handler:    _Monkey_ListAllAlbums_Handler,
		},
		{
			MethodName: "ListSongs",
			Handler:    _Monkey_ListSongs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spec/monkey.proto",
}
