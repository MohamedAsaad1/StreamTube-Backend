// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.6
// source: proto/upload.proto

package upload

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
	FileService_UploadVideo_FullMethodName = "/FileService/UploadVideo"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	UploadVideo(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadVideoClient, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadVideo(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadVideoClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], FileService_UploadVideo_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceUploadVideoClient{ClientStream: stream}
	return x, nil
}

type FileService_UploadVideoClient interface {
	Send(*UploadVideoRequest) error
	CloseAndRecv() (*UploadVideoResponse, error)
	grpc.ClientStream
}

type fileServiceUploadVideoClient struct {
	grpc.ClientStream
}

func (x *fileServiceUploadVideoClient) Send(m *UploadVideoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceUploadVideoClient) CloseAndRecv() (*UploadVideoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadVideoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	UploadVideo(FileService_UploadVideoServer) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) UploadVideo(FileService_UploadVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_UploadVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).UploadVideo(&fileServiceUploadVideoServer{ServerStream: stream})
}

type FileService_UploadVideoServer interface {
	SendAndClose(*UploadVideoResponse) error
	Recv() (*UploadVideoRequest, error)
	grpc.ServerStream
}

type fileServiceUploadVideoServer struct {
	grpc.ServerStream
}

func (x *fileServiceUploadVideoServer) SendAndClose(m *UploadVideoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceUploadVideoServer) Recv() (*UploadVideoRequest, error) {
	m := new(UploadVideoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadVideo",
			Handler:       _FileService_UploadVideo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/upload.proto",
}
