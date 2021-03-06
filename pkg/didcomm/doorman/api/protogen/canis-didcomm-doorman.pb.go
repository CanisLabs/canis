// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: canis-didcomm-doorman.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	common "github.com/scoir/canis/pkg/protogen/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_canis_didcomm_doorman_proto protoreflect.FileDescriptor

var file_canis_didcomm_doorman_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x61, 0x6e, 0x69, 0x73, 0x2d, 0x64, 0x69, 0x64, 0x63, 0x6f, 0x6d, 0x6d, 0x2d,
	0x64, 0x6f, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64,
	0x69, 0x64, 0x63, 0x6f, 0x6d, 0x6d, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xac, 0x01,
	0x0a, 0x07, 0x44, 0x6f, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13,
	0x64, 0x69, 0x64, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x64, 0x6f, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_canis_didcomm_doorman_proto_goTypes = []interface{}{
	(*common.InvitationRequest)(nil),        // 0: common.InvitationRequest
	(*common.AcceptInvitationRequest)(nil),  // 1: common.AcceptInvitationRequest
	(*common.InvitationResponse)(nil),       // 2: common.InvitationResponse
	(*common.AcceptInvitationResponse)(nil), // 3: common.AcceptInvitationResponse
}
var file_canis_didcomm_doorman_proto_depIdxs = []int32{
	0, // 0: didcomm.Doorman.GetInvitation:input_type -> common.InvitationRequest
	1, // 1: didcomm.Doorman.AcceptInvitation:input_type -> common.AcceptInvitationRequest
	2, // 2: didcomm.Doorman.GetInvitation:output_type -> common.InvitationResponse
	3, // 3: didcomm.Doorman.AcceptInvitation:output_type -> common.AcceptInvitationResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_canis_didcomm_doorman_proto_init() }
func file_canis_didcomm_doorman_proto_init() {
	if File_canis_didcomm_doorman_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_canis_didcomm_doorman_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_canis_didcomm_doorman_proto_goTypes,
		DependencyIndexes: file_canis_didcomm_doorman_proto_depIdxs,
	}.Build()
	File_canis_didcomm_doorman_proto = out.File
	file_canis_didcomm_doorman_proto_rawDesc = nil
	file_canis_didcomm_doorman_proto_goTypes = nil
	file_canis_didcomm_doorman_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DoormanClient is the client API for Doorman service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DoormanClient interface {
	GetInvitation(ctx context.Context, in *common.InvitationRequest, opts ...grpc.CallOption) (*common.InvitationResponse, error)
	AcceptInvitation(ctx context.Context, in *common.AcceptInvitationRequest, opts ...grpc.CallOption) (*common.AcceptInvitationResponse, error)
}

type doormanClient struct {
	cc grpc.ClientConnInterface
}

func NewDoormanClient(cc grpc.ClientConnInterface) DoormanClient {
	return &doormanClient{cc}
}

func (c *doormanClient) GetInvitation(ctx context.Context, in *common.InvitationRequest, opts ...grpc.CallOption) (*common.InvitationResponse, error) {
	out := new(common.InvitationResponse)
	err := c.cc.Invoke(ctx, "/didcomm.Doorman/GetInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *doormanClient) AcceptInvitation(ctx context.Context, in *common.AcceptInvitationRequest, opts ...grpc.CallOption) (*common.AcceptInvitationResponse, error) {
	out := new(common.AcceptInvitationResponse)
	err := c.cc.Invoke(ctx, "/didcomm.Doorman/AcceptInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoormanServer is the server API for Doorman service.
type DoormanServer interface {
	GetInvitation(context.Context, *common.InvitationRequest) (*common.InvitationResponse, error)
	AcceptInvitation(context.Context, *common.AcceptInvitationRequest) (*common.AcceptInvitationResponse, error)
}

// UnimplementedDoormanServer can be embedded to have forward compatible implementations.
type UnimplementedDoormanServer struct {
}

func (*UnimplementedDoormanServer) GetInvitation(context.Context, *common.InvitationRequest) (*common.InvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvitation not implemented")
}
func (*UnimplementedDoormanServer) AcceptInvitation(context.Context, *common.AcceptInvitationRequest) (*common.AcceptInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptInvitation not implemented")
}

func RegisterDoormanServer(s *grpc.Server, srv DoormanServer) {
	s.RegisterService(&_Doorman_serviceDesc, srv)
}

func _Doorman_GetInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.InvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoormanServer).GetInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/didcomm.Doorman/GetInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoormanServer).GetInvitation(ctx, req.(*common.InvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Doorman_AcceptInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.AcceptInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoormanServer).AcceptInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/didcomm.Doorman/AcceptInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoormanServer).AcceptInvitation(ctx, req.(*common.AcceptInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Doorman_serviceDesc = grpc.ServiceDesc{
	ServiceName: "didcomm.Doorman",
	HandlerType: (*DoormanServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInvitation",
			Handler:    _Doorman_GetInvitation_Handler,
		},
		{
			MethodName: "AcceptInvitation",
			Handler:    _Doorman_AcceptInvitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "canis-didcomm-doorman.proto",
}
