// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fmpb

import (
	context "context"
	model "github.com/DataWorkbench/gproto/pkg/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FileManagerClient is the client API for FileManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileManagerClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileManager_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileManager_DownloadFileClient, error)
	ListFiles(ctx context.Context, in *FilesFilterRequest, opts ...grpc.CallOption) (*FileListResponse, error)
	UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
	DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error)
}

type fileManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewFileManagerClient(cc grpc.ClientConnInterface) FileManagerClient {
	return &fileManagerClient{cc}
}

func (c *fileManagerClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileManager_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileManager_serviceDesc.Streams[0], "/fmpb.FileManager/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileManagerUploadFileClient{stream}
	return x, nil
}

type FileManager_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*FileInfoResponse, error)
	grpc.ClientStream
}

type fileManagerUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileManagerUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileManagerUploadFileClient) CloseAndRecv() (*FileInfoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileInfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileManagerClient) DownloadFile(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FileManager_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileManager_serviceDesc.Streams[1], "/fmpb.FileManager/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileManagerDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileManager_DownloadFileClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type fileManagerDownloadFileClient struct {
	grpc.ClientStream
}

func (x *fileManagerDownloadFileClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileManagerClient) ListFiles(ctx context.Context, in *FilesFilterRequest, opts ...grpc.CallOption) (*FileListResponse, error) {
	out := new(FileListResponse)
	err := c.cc.Invoke(ctx, "/fmpb.FileManager/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/fmpb.FileManager/UpdateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileManagerClient) DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*model.EmptyStruct, error) {
	out := new(model.EmptyStruct)
	err := c.cc.Invoke(ctx, "/fmpb.FileManager/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileManagerServer is the server API for FileManager service.
// All implementations must embed UnimplementedFileManagerServer
// for forward compatibility
type FileManagerServer interface {
	UploadFile(FileManager_UploadFileServer) error
	DownloadFile(*DownloadRequest, FileManager_DownloadFileServer) error
	ListFiles(context.Context, *FilesFilterRequest) (*FileListResponse, error)
	UpdateFile(context.Context, *UpdateFileRequest) (*model.EmptyStruct, error)
	DeleteFile(context.Context, *DeleteFileRequest) (*model.EmptyStruct, error)
	mustEmbedUnimplementedFileManagerServer()
}

// UnimplementedFileManagerServer must be embedded to have forward compatible implementations.
type UnimplementedFileManagerServer struct {
}

func (UnimplementedFileManagerServer) UploadFile(FileManager_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileManagerServer) DownloadFile(*DownloadRequest, FileManager_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileManagerServer) ListFiles(context.Context, *FilesFilterRequest) (*FileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFileManagerServer) UpdateFile(context.Context, *UpdateFileRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedFileManagerServer) DeleteFile(context.Context, *DeleteFileRequest) (*model.EmptyStruct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFileManagerServer) mustEmbedUnimplementedFileManagerServer() {}

// UnsafeFileManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileManagerServer will
// result in compilation errors.
type UnsafeFileManagerServer interface {
	mustEmbedUnimplementedFileManagerServer()
}

func RegisterFileManagerServer(s grpc.ServiceRegistrar, srv FileManagerServer) {
	s.RegisterService(&_FileManager_serviceDesc, srv)
}

func _FileManager_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileManagerServer).UploadFile(&fileManagerUploadFileServer{stream})
}

type FileManager_UploadFileServer interface {
	SendAndClose(*FileInfoResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type fileManagerUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileManagerUploadFileServer) SendAndClose(m *FileInfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileManagerUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileManager_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileManagerServer).DownloadFile(m, &fileManagerDownloadFileServer{stream})
}

type FileManager_DownloadFileServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type fileManagerDownloadFileServer struct {
	grpc.ServerStream
}

func (x *fileManagerDownloadFileServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileManager_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilesFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fmpb.FileManager/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).ListFiles(ctx, req.(*FilesFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fmpb.FileManager/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).UpdateFile(ctx, req.(*UpdateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileManager_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileManagerServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fmpb.FileManager/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileManagerServer).DeleteFile(ctx, req.(*DeleteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fmpb.FileManager",
	HandlerType: (*FileManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _FileManager_ListFiles_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _FileManager_UpdateFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _FileManager_DeleteFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileManager_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _FileManager_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/filemanager.proto",
}
