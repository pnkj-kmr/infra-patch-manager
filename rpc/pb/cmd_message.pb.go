// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: cmd_message.proto

package pb

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

type CmdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Pass string `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *CmdReq) Reset() {
	*x = CmdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdReq) ProtoMessage() {}

func (x *CmdReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdReq.ProtoReflect.Descriptor instead.
func (*CmdReq) Descriptor() ([]byte, []int) {
	return file_cmd_message_proto_rawDescGZIP(), []int{0}
}

func (x *CmdReq) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *CmdReq) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

type CmdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Out []byte `protobuf:"bytes,1,opt,name=out,proto3" json:"out,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *CmdResp) Reset() {
	*x = CmdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdResp) ProtoMessage() {}

func (x *CmdResp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdResp.ProtoReflect.Descriptor instead.
func (*CmdResp) Descriptor() ([]byte, []int) {
	return file_cmd_message_proto_rawDescGZIP(), []int{1}
}

func (x *CmdResp) GetOut() []byte {
	if x != nil {
		return x.Out
	}
	return nil
}

func (x *CmdResp) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_cmd_message_proto protoreflect.FileDescriptor

var file_cmd_message_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6d, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x06, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x07, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6f, 0x75, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x72, 0x72, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_message_proto_rawDescOnce sync.Once
	file_cmd_message_proto_rawDescData = file_cmd_message_proto_rawDesc
)

func file_cmd_message_proto_rawDescGZIP() []byte {
	file_cmd_message_proto_rawDescOnce.Do(func() {
		file_cmd_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_message_proto_rawDescData)
	})
	return file_cmd_message_proto_rawDescData
}

var file_cmd_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmd_message_proto_goTypes = []interface{}{
	(*CmdReq)(nil),  // 0: CmdReq
	(*CmdResp)(nil), // 1: CmdResp
}
var file_cmd_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_message_proto_init() }
func file_cmd_message_proto_init() {
	if File_cmd_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdReq); i {
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
		file_cmd_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdResp); i {
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
			RawDescriptor: file_cmd_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_message_proto_goTypes,
		DependencyIndexes: file_cmd_message_proto_depIdxs,
		MessageInfos:      file_cmd_message_proto_msgTypes,
	}.Build()
	File_cmd_message_proto = out.File
	file_cmd_message_proto_rawDesc = nil
	file_cmd_message_proto_goTypes = nil
	file_cmd_message_proto_depIdxs = nil
}
