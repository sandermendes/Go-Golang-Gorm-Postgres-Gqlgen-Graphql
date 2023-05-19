// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0
// source: account/v1/account.proto

package accountv1

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

type AccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Redirect string `protobuf:"bytes,2,opt,name=redirect,proto3" json:"redirect,omitempty"`
}

func (x *AccountResponse) Reset() {
	*x = AccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountResponse) ProtoMessage() {}

func (x *AccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountResponse.ProtoReflect.Descriptor instead.
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return file_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *AccountResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AccountResponse) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

var File_account_v1_account_proto protoreflect.FileDescriptor

var file_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x43, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x42, 0x68, 0x5a, 0x66, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x6f, 0x2d, 0x47, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2d, 0x47, 0x6f, 0x72, 0x6d, 0x2d, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x2d, 0x47, 0x71, 0x6c, 0x67, 0x65, 0x6e, 0x2d, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c,
	0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_v1_account_proto_rawDescOnce sync.Once
	file_account_v1_account_proto_rawDescData = file_account_v1_account_proto_rawDesc
)

func file_account_v1_account_proto_rawDescGZIP() []byte {
	file_account_v1_account_proto_rawDescOnce.Do(func() {
		file_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_v1_account_proto_rawDescData)
	})
	return file_account_v1_account_proto_rawDescData
}

var file_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_account_v1_account_proto_goTypes = []interface{}{
	(*AccountResponse)(nil), // 0: account.v1.AccountResponse
}
var file_account_v1_account_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_account_v1_account_proto_init() }
func file_account_v1_account_proto_init() {
	if File_account_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountResponse); i {
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
			RawDescriptor: file_account_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_v1_account_proto_goTypes,
		DependencyIndexes: file_account_v1_account_proto_depIdxs,
		MessageInfos:      file_account_v1_account_proto_msgTypes,
	}.Build()
	File_account_v1_account_proto = out.File
	file_account_v1_account_proto_rawDesc = nil
	file_account_v1_account_proto_goTypes = nil
	file_account_v1_account_proto_depIdxs = nil
}
