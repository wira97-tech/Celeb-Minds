// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: question.proto

package pb

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
	Question_Get_FullMethodName    = "/protogrpc.Question/Get"
	Question_GetAll_FullMethodName = "/protogrpc.Question/GetAll"
)

// QuestionClient is the client API for Question service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionClient interface {
	Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResponseOne, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error)
}

type questionClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionClient(cc grpc.ClientConnInterface) QuestionClient {
	return &questionClient{cc}
}

func (c *questionClient) Get(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResponseOne, error) {
	out := new(ResponseOne)
	err := c.cc.Invoke(ctx, Question_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Responses, error) {
	out := new(Responses)
	err := c.cc.Invoke(ctx, Question_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionServer is the server API for Question service.
// All implementations must embed UnimplementedQuestionServer
// for forward compatibility
type QuestionServer interface {
	Get(context.Context, *Id) (*ResponseOne, error)
	GetAll(context.Context, *Empty) (*Responses, error)
	mustEmbedUnimplementedQuestionServer()
}

// UnimplementedQuestionServer must be embedded to have forward compatible implementations.
type UnimplementedQuestionServer struct {
}

func (UnimplementedQuestionServer) Get(context.Context, *Id) (*ResponseOne, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedQuestionServer) GetAll(context.Context, *Empty) (*Responses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedQuestionServer) mustEmbedUnimplementedQuestionServer() {}

// UnsafeQuestionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionServer will
// result in compilation errors.
type UnsafeQuestionServer interface {
	mustEmbedUnimplementedQuestionServer()
}

func RegisterQuestionServer(s grpc.ServiceRegistrar, srv QuestionServer) {
	s.RegisterService(&Question_ServiceDesc, srv)
}

func _Question_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Question_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServer).Get(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Question_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Question_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Question_ServiceDesc is the grpc.ServiceDesc for Question service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Question_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protogrpc.Question",
	HandlerType: (*QuestionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Question_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Question_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "question.proto",
}