// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: proto/auth.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IntrospectReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IntrospectReq) Reset() {
	*x = IntrospectReq{}
	mi := &file_proto_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntrospectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntrospectReq) ProtoMessage() {}

func (x *IntrospectReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntrospectReq.ProtoReflect.Descriptor instead.
func (*IntrospectReq) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *IntrospectReq) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type IntrospectResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tid           string                 `protobuf:"bytes,1,opt,name=tid,proto3" json:"tid,omitempty"`
	Sub           string                 `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IntrospectResp) Reset() {
	*x = IntrospectResp{}
	mi := &file_proto_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntrospectResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntrospectResp) ProtoMessage() {}

func (x *IntrospectResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntrospectResp.ProtoReflect.Descriptor instead.
func (*IntrospectResp) Descriptor() ([]byte, []int) {
	return file_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *IntrospectResp) GetTid() string {
	if x != nil {
		return x.Tid
	}
	return ""
}

func (x *IntrospectResp) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

var File_proto_auth_proto protoreflect.FileDescriptor

var file_proto_auth_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x32, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73,
	0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x34, 0x0a, 0x0e, 0x49, 0x6e,
	0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62,
	0x32, 0x47, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x72, 0x6f,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x05, 0x5a, 0x03, 0x70, 0x62, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_auth_proto_rawDescOnce sync.Once
	file_proto_auth_proto_rawDescData []byte
)

func file_proto_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_auth_proto_rawDesc), len(file_proto_auth_proto_rawDesc)))
	})
	return file_proto_auth_proto_rawDescData
}

var file_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_auth_proto_goTypes = []any{
	(*IntrospectReq)(nil),  // 0: pb.IntrospectReq
	(*IntrospectResp)(nil), // 1: pb.IntrospectResp
}
var file_proto_auth_proto_depIdxs = []int32{
	0, // 0: pb.AuthService.IntrospectToken:input_type -> pb.IntrospectReq
	1, // 1: pb.AuthService.IntrospectToken:output_type -> pb.IntrospectResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_proto_init() }
func file_proto_auth_proto_init() {
	if File_proto_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_auth_proto_rawDesc), len(file_proto_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_proto_depIdxs,
		MessageInfos:      file_proto_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_proto = out.File
	file_proto_auth_proto_goTypes = nil
	file_proto_auth_proto_depIdxs = nil
}
