// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: contracts/api.proto

package contracts

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyName string `protobuf:"bytes,1,opt,name=my_name,json=myName,proto3" json:"my_name,omitempty"`
	MyAge  int32  `protobuf:"varint,2,opt,name=my_age,json=myAge,proto3" json:"my_age,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_contracts_api_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetMyName() string {
	if x != nil {
		return x.MyName
	}
	return ""
}

func (x *HelloRequest) GetMyAge() int32 {
	if x != nil {
		return x.MyAge
	}
	return 0
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseText      string `protobuf:"bytes,1,opt,name=response_text,json=responseText,proto3" json:"response_text,omitempty"`
	ServerUnixTimeI   int32  `protobuf:"varint,2,opt,name=server_unix_time_i,json=serverUnixTimeI,proto3" json:"server_unix_time_i,omitempty"`
	ServerUnixTimeStr string `protobuf:"bytes,3,opt,name=server_unix_time_str,json=serverUnixTimeStr,proto3" json:"server_unix_time_str,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contracts_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contracts_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_contracts_api_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetResponseText() string {
	if x != nil {
		return x.ResponseText
	}
	return ""
}

func (x *HelloResponse) GetServerUnixTimeI() int32 {
	if x != nil {
		return x.ServerUnixTimeI
	}
	return 0
}

func (x *HelloResponse) GetServerUnixTimeStr() string {
	if x != nil {
		return x.ServerUnixTimeStr
	}
	return ""
}

var File_contracts_api_proto protoreflect.FileDescriptor

var file_contracts_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15,
	0x0a, 0x06, 0x6d, 0x79, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6d, 0x79, 0x41, 0x67, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2b, 0x0a, 0x12,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x55, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x12, 0x2f, 0x0a, 0x14, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x55,
	0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x32, 0x68, 0x0a, 0x0e, 0x4d, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x05,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x0d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_contracts_api_proto_rawDescOnce sync.Once
	file_contracts_api_proto_rawDescData = file_contracts_api_proto_rawDesc
)

func file_contracts_api_proto_rawDescGZIP() []byte {
	file_contracts_api_proto_rawDescOnce.Do(func() {
		file_contracts_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_contracts_api_proto_rawDescData)
	})
	return file_contracts_api_proto_rawDescData
}

var file_contracts_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contracts_api_proto_goTypes = []any{
	(*HelloRequest)(nil),  // 0: HelloRequest
	(*HelloResponse)(nil), // 1: HelloResponse
}
var file_contracts_api_proto_depIdxs = []int32{
	0, // 0: MyHelloService.Hello:input_type -> HelloRequest
	0, // 1: MyHelloService.HelloStream:input_type -> HelloRequest
	1, // 2: MyHelloService.Hello:output_type -> HelloResponse
	1, // 3: MyHelloService.HelloStream:output_type -> HelloResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contracts_api_proto_init() }
func file_contracts_api_proto_init() {
	if File_contracts_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contracts_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HelloRequest); i {
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
		file_contracts_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_contracts_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contracts_api_proto_goTypes,
		DependencyIndexes: file_contracts_api_proto_depIdxs,
		MessageInfos:      file_contracts_api_proto_msgTypes,
	}.Build()
	File_contracts_api_proto = out.File
	file_contracts_api_proto_rawDesc = nil
	file_contracts_api_proto_goTypes = nil
	file_contracts_api_proto_depIdxs = nil
}
