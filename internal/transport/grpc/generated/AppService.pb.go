// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: AppService.proto

package generated

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

type SubnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *SubnetRequest) Reset() {
	*x = SubnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AppService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetRequest) ProtoMessage() {}

func (x *SubnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_AppService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetRequest.ProtoReflect.Descriptor instead.
func (*SubnetRequest) Descriptor() ([]byte, []int) {
	return file_AppService_proto_rawDescGZIP(), []int{0}
}

func (x *SubnetRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AppService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_AppService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_AppService_proto_rawDescGZIP(), []int{1}
}

type ServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ServerResponse) Reset() {
	*x = ServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AppService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerResponse) ProtoMessage() {}

func (x *ServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_AppService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerResponse.ProtoReflect.Descriptor instead.
func (*ServerResponse) Descriptor() ([]byte, []int) {
	return file_AppService_proto_rawDescGZIP(), []int{2}
}

func (x *ServerResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_AppService_proto protoreflect.FileDescriptor

var file_AppService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x41, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x24, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xc3, 0x02, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x74, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x13, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c,
	0x2e, 0x2f, 0x3b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AppService_proto_rawDescOnce sync.Once
	file_AppService_proto_rawDescData = file_AppService_proto_rawDesc
)

func file_AppService_proto_rawDescGZIP() []byte {
	file_AppService_proto_rawDescOnce.Do(func() {
		file_AppService_proto_rawDescData = protoimpl.X.CompressGZIP(file_AppService_proto_rawDescData)
	})
	return file_AppService_proto_rawDescData
}

var file_AppService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_AppService_proto_goTypes = []interface{}{
	(*SubnetRequest)(nil),  // 0: api.SubnetRequest
	(*EmptyRequest)(nil),   // 1: api.EmptyRequest
	(*ServerResponse)(nil), // 2: api.ServerResponse
}
var file_AppService_proto_depIdxs = []int32{
	1, // 0: api.AppService.ResetBucket:input_type -> api.EmptyRequest
	0, // 1: api.AppService.AddToWhitelist:input_type -> api.SubnetRequest
	0, // 2: api.AppService.AddToBlacklist:input_type -> api.SubnetRequest
	0, // 3: api.AppService.RemoveFromWhitelist:input_type -> api.SubnetRequest
	0, // 4: api.AppService.RemoveFromBlacklist:input_type -> api.SubnetRequest
	2, // 5: api.AppService.ResetBucket:output_type -> api.ServerResponse
	2, // 6: api.AppService.AddToWhitelist:output_type -> api.ServerResponse
	2, // 7: api.AppService.AddToBlacklist:output_type -> api.ServerResponse
	2, // 8: api.AppService.RemoveFromWhitelist:output_type -> api.ServerResponse
	2, // 9: api.AppService.RemoveFromBlacklist:output_type -> api.ServerResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AppService_proto_init() }
func file_AppService_proto_init() {
	if File_AppService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AppService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubnetRequest); i {
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
		file_AppService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_AppService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerResponse); i {
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
			RawDescriptor: file_AppService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_AppService_proto_goTypes,
		DependencyIndexes: file_AppService_proto_depIdxs,
		MessageInfos:      file_AppService_proto_msgTypes,
	}.Build()
	File_AppService_proto = out.File
	file_AppService_proto_rawDesc = nil
	file_AppService_proto_goTypes = nil
	file_AppService_proto_depIdxs = nil
}