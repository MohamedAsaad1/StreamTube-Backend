// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.6
// source: proto/transcoding.proto

package transcoding

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	VideoStatusService_StatusVideo_FullMethodName = "/transcoding.VideoStatusService/StatusVideo"
)

// VideoStatusServiceClient is the client API for VideoStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoStatusServiceClient interface {
	StatusVideo(ctx context.Context, in *VideoUuidRequest, opts ...grpc.CallOption) (VideoStatusService_StatusVideoClient, error)
}

type videoStatusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoStatusServiceClient(cc grpc.ClientConnInterface) VideoStatusServiceClient {
	return &videoStatusServiceClient{cc}
}

func (c *videoStatusServiceClient) StatusVideo(ctx context.Context, in *VideoUuidRequest, opts ...grpc.CallOption) (VideoStatusService_StatusVideoClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &VideoStatusService_ServiceDesc.Streams[0], VideoStatusService_StatusVideo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &videoStatusServiceStatusVideoClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VideoStatusService_StatusVideoClient interface {
	Recv() (*VideoStatusResponse, error)
	grpc.ClientStream
}

type videoStatusServiceStatusVideoClient struct {
	grpc.ClientStream
}

func (x *videoStatusServiceStatusVideoClient) Recv() (*VideoStatusResponse, error) {
	m := new(VideoStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VideoStatusServiceServer is the server API for VideoStatusService service.
// All implementations must embed UnimplementedVideoStatusServiceServer
// for forward compatibility
type VideoStatusServiceServer interface {
	StatusVideo(*VideoUuidRequest, VideoStatusService_StatusVideoServer) error
	mustEmbedUnimplementedVideoStatusServiceServer()
}

// UnimplementedVideoStatusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoStatusServiceServer struct {
}

func (UnimplementedVideoStatusServiceServer) StatusVideo(*VideoUuidRequest, VideoStatusService_StatusVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method StatusVideo not implemented")
}
func (UnimplementedVideoStatusServiceServer) mustEmbedUnimplementedVideoStatusServiceServer() {}

// UnsafeVideoStatusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoStatusServiceServer will
// result in compilation errors.
type UnsafeVideoStatusServiceServer interface {
	mustEmbedUnimplementedVideoStatusServiceServer()
}

func RegisterVideoStatusServiceServer(s grpc.ServiceRegistrar, srv VideoStatusServiceServer) {
	s.RegisterService(&VideoStatusService_ServiceDesc, srv)
}

func _VideoStatusService_StatusVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VideoUuidRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VideoStatusServiceServer).StatusVideo(m, &videoStatusServiceStatusVideoServer{ServerStream: stream})
}

type VideoStatusService_StatusVideoServer interface {
	Send(*VideoStatusResponse) error
	grpc.ServerStream
}

type videoStatusServiceStatusVideoServer struct {
	grpc.ServerStream
}

func (x *videoStatusServiceStatusVideoServer) Send(m *VideoStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

// VideoStatusService_ServiceDesc is the grpc.ServiceDesc for VideoStatusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoStatusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transcoding.VideoStatusService",
	HandlerType: (*VideoStatusServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StatusVideo",
			Handler:       _VideoStatusService_StatusVideo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/transcoding.proto",
}

const (
	Transcoder_NotifyUploadComplete_FullMethodName = "/transcoding.Transcoder/NotifyUploadComplete"
)

// TranscoderClient is the client API for Transcoder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranscoderClient interface {
	NotifyUploadComplete(ctx context.Context, in *VideoUuidRequest, opts ...grpc.CallOption) (*TranscodeResponse, error)
}

type transcoderClient struct {
	cc grpc.ClientConnInterface
}

func NewTranscoderClient(cc grpc.ClientConnInterface) TranscoderClient {
	return &transcoderClient{cc}
}

func (c *transcoderClient) NotifyUploadComplete(ctx context.Context, in *VideoUuidRequest, opts ...grpc.CallOption) (*TranscodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TranscodeResponse)
	err := c.cc.Invoke(ctx, Transcoder_NotifyUploadComplete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranscoderServer is the server API for Transcoder service.
// All implementations must embed UnimplementedTranscoderServer
// for forward compatibility
type TranscoderServer interface {
	NotifyUploadComplete(context.Context, *VideoUuidRequest) (*TranscodeResponse, error)
	mustEmbedUnimplementedTranscoderServer()
}

// UnimplementedTranscoderServer must be embedded to have forward compatible implementations.
type UnimplementedTranscoderServer struct {
}

func (UnimplementedTranscoderServer) NotifyUploadComplete(context.Context, *VideoUuidRequest) (*TranscodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyUploadComplete not implemented")
}
func (UnimplementedTranscoderServer) mustEmbedUnimplementedTranscoderServer() {}

// UnsafeTranscoderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranscoderServer will
// result in compilation errors.
type UnsafeTranscoderServer interface {
	mustEmbedUnimplementedTranscoderServer()
}

func RegisterTranscoderServer(s grpc.ServiceRegistrar, srv TranscoderServer) {
	s.RegisterService(&Transcoder_ServiceDesc, srv)
}

func _Transcoder_NotifyUploadComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoUuidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranscoderServer).NotifyUploadComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Transcoder_NotifyUploadComplete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranscoderServer).NotifyUploadComplete(ctx, req.(*VideoUuidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transcoder_ServiceDesc is the grpc.ServiceDesc for Transcoder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transcoder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transcoding.Transcoder",
	HandlerType: (*TranscoderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyUploadComplete",
			Handler:    _Transcoder_NotifyUploadComplete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transcoding.proto",
}
