// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: didcomm/issuer/api/canis-didcomm-issuer.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CredentialAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CredentialAttribute) Reset() {
	*x = CredentialAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialAttribute) ProtoMessage() {}

func (x *CredentialAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialAttribute.ProtoReflect.Descriptor instead.
func (*CredentialAttribute) Descriptor() ([]byte, []int) {
	return file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescGZIP(), []int{0}
}

func (x *CredentialAttribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CredentialAttribute) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Credential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaId   string                 `protobuf:"bytes,1,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	Comment    string                 `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Type       string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Attributes []*CredentialAttribute `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Credential) Reset() {
	*x = Credential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescGZIP(), []int{1}
}

func (x *Credential) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *Credential) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Credential) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Credential) GetAttributes() []*CredentialAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type IssueCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId    string      `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	SubjectId  string      `protobuf:"bytes,2,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	Credential *Credential `protobuf:"bytes,3,opt,name=credential,proto3" json:"credential,omitempty"`
}

func (x *IssueCredentialRequest) Reset() {
	*x = IssueCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCredentialRequest) ProtoMessage() {}

func (x *IssueCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCredentialRequest.ProtoReflect.Descriptor instead.
func (*IssueCredentialRequest) Descriptor() ([]byte, []int) {
	return file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescGZIP(), []int{2}
}

func (x *IssueCredentialRequest) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *IssueCredentialRequest) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *IssueCredentialRequest) GetCredential() *Credential {
	if x != nil {
		return x.Credential
	}
	return nil
}

type IssueCredentialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CredentialId string `protobuf:"bytes,1,opt,name=credential_id,json=credentialId,proto3" json:"credential_id,omitempty"`
}

func (x *IssueCredentialResponse) Reset() {
	*x = IssueCredentialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCredentialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCredentialResponse) ProtoMessage() {}

func (x *IssueCredentialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCredentialResponse.ProtoReflect.Descriptor instead.
func (*IssueCredentialResponse) Descriptor() ([]byte, []int) {
	return file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescGZIP(), []int{3}
}

func (x *IssueCredentialResponse) GetCredentialId() string {
	if x != nil {
		return x.CredentialId
	}
	return ""
}

var File_didcomm_issuer_api_canis_didcomm_issuer_proto protoreflect.FileDescriptor

var file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x64, 0x69, 0x64, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x6e, 0x69, 0x73, 0x2d, 0x64, 0x69, 0x64, 0x63, 0x6f,
	0x6d, 0x6d, 0x2d, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x69, 0x64, 0x63, 0x6f, 0x6d, 0x6d, 0x22, 0x3f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x69, 0x64, 0x63, 0x6f, 0x6d,
	0x6d, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x22, 0x87, 0x01, 0x0a, 0x16, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x64,
	0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x3e, 0x0a, 0x17, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x32, 0x60, 0x0a, 0x06, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x0f, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1f, 0x2e, 0x64, 0x69, 0x64, 0x63, 0x6f,
	0x6d, 0x6d, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x69, 0x64, 0x63,
	0x6f, 0x6d, 0x6d, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a,
	0x12, 0x64, 0x69, 0x64, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescOnce sync.Once
	file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescData = file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDesc
)

func file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescGZIP() []byte {
	file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescOnce.Do(func() {
		file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescData = protoimpl.X.CompressGZIP(file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescData)
	})
	return file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDescData
}

var file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_didcomm_issuer_api_canis_didcomm_issuer_proto_goTypes = []interface{}{
	(*CredentialAttribute)(nil),     // 0: didcomm.CredentialAttribute
	(*Credential)(nil),              // 1: didcomm.Credential
	(*IssueCredentialRequest)(nil),  // 2: didcomm.IssueCredentialRequest
	(*IssueCredentialResponse)(nil), // 3: didcomm.IssueCredentialResponse
}
var file_didcomm_issuer_api_canis_didcomm_issuer_proto_depIdxs = []int32{
	0, // 0: didcomm.Credential.attributes:type_name -> didcomm.CredentialAttribute
	1, // 1: didcomm.IssueCredentialRequest.credential:type_name -> didcomm.Credential
	2, // 2: didcomm.Issuer.IssueCredential:input_type -> didcomm.IssueCredentialRequest
	3, // 3: didcomm.Issuer.IssueCredential:output_type -> didcomm.IssueCredentialResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_didcomm_issuer_api_canis_didcomm_issuer_proto_init() }
func file_didcomm_issuer_api_canis_didcomm_issuer_proto_init() {
	if File_didcomm_issuer_api_canis_didcomm_issuer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialAttribute); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credential); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCredentialRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCredentialResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_didcomm_issuer_api_canis_didcomm_issuer_proto_goTypes,
		DependencyIndexes: file_didcomm_issuer_api_canis_didcomm_issuer_proto_depIdxs,
		MessageInfos:      file_didcomm_issuer_api_canis_didcomm_issuer_proto_msgTypes,
	}.Build()
	File_didcomm_issuer_api_canis_didcomm_issuer_proto = out.File
	file_didcomm_issuer_api_canis_didcomm_issuer_proto_rawDesc = nil
	file_didcomm_issuer_api_canis_didcomm_issuer_proto_goTypes = nil
	file_didcomm_issuer_api_canis_didcomm_issuer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IssuerClient is the client API for Issuer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IssuerClient interface {
	IssueCredential(ctx context.Context, in *IssueCredentialRequest, opts ...grpc.CallOption) (*IssueCredentialResponse, error)
}

type issuerClient struct {
	cc grpc.ClientConnInterface
}

func NewIssuerClient(cc grpc.ClientConnInterface) IssuerClient {
	return &issuerClient{cc}
}

func (c *issuerClient) IssueCredential(ctx context.Context, in *IssueCredentialRequest, opts ...grpc.CallOption) (*IssueCredentialResponse, error) {
	out := new(IssueCredentialResponse)
	err := c.cc.Invoke(ctx, "/didcomm.Issuer/IssueCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssuerServer is the server API for Issuer service.
type IssuerServer interface {
	IssueCredential(context.Context, *IssueCredentialRequest) (*IssueCredentialResponse, error)
}

// UnimplementedIssuerServer can be embedded to have forward compatible implementations.
type UnimplementedIssuerServer struct {
}

func (*UnimplementedIssuerServer) IssueCredential(context.Context, *IssueCredentialRequest) (*IssueCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueCredential not implemented")
}

func RegisterIssuerServer(s *grpc.Server, srv IssuerServer) {
	s.RegisterService(&_Issuer_serviceDesc, srv)
}

func _Issuer_IssueCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuerServer).IssueCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/didcomm.Issuer/IssueCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuerServer).IssueCredential(ctx, req.(*IssueCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Issuer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "didcomm.Issuer",
	HandlerType: (*IssuerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueCredential",
			Handler:    _Issuer_IssueCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "didcomm/issuer/api/canis-didcomm-issuer.proto",
}
