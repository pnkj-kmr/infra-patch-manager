// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: download_message.proto

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

type DownloadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *DownloadReq) Reset() {
	*x = DownloadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_download_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadReq) ProtoMessage() {}

func (x *DownloadReq) ProtoReflect() protoreflect.Message {
	mi := &file_download_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadReq.ProtoReflect.Descriptor instead.
func (*DownloadReq) Descriptor() ([]byte, []int) {
	return file_download_message_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type DownloadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DownloadInfo) Reset() {
	*x = DownloadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_download_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadInfo) ProtoMessage() {}

func (x *DownloadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_download_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadInfo.ProtoReflect.Descriptor instead.
func (*DownloadInfo) Descriptor() ([]byte, []int) {
	return file_download_message_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DownloadInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DownloadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*DownloadResp_File
	//	*DownloadResp_ChunkData
	Data isDownloadResp_Data `protobuf_oneof:"data"`
}

func (x *DownloadResp) Reset() {
	*x = DownloadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_download_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResp) ProtoMessage() {}

func (x *DownloadResp) ProtoReflect() protoreflect.Message {
	mi := &file_download_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResp.ProtoReflect.Descriptor instead.
func (*DownloadResp) Descriptor() ([]byte, []int) {
	return file_download_message_proto_rawDescGZIP(), []int{2}
}

func (m *DownloadResp) GetData() isDownloadResp_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *DownloadResp) GetFile() *DownloadInfo {
	if x, ok := x.GetData().(*DownloadResp_File); ok {
		return x.File
	}
	return nil
}

func (x *DownloadResp) GetChunkData() []byte {
	if x, ok := x.GetData().(*DownloadResp_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isDownloadResp_Data interface {
	isDownloadResp_Data()
}

type DownloadResp_File struct {
	File *DownloadInfo `protobuf:"bytes,1,opt,name=file,proto3,oneof"`
}

type DownloadResp_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*DownloadResp_File) isDownloadResp_Data() {}

func (*DownloadResp_ChunkData) isDownloadResp_Data() {}

var File_download_message_proto protoreflect.FileDescriptor

var file_download_message_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0b, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x5c, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x23, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_download_message_proto_rawDescOnce sync.Once
	file_download_message_proto_rawDescData = file_download_message_proto_rawDesc
)

func file_download_message_proto_rawDescGZIP() []byte {
	file_download_message_proto_rawDescOnce.Do(func() {
		file_download_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_download_message_proto_rawDescData)
	})
	return file_download_message_proto_rawDescData
}

var file_download_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_download_message_proto_goTypes = []interface{}{
	(*DownloadReq)(nil),  // 0: DownloadReq
	(*DownloadInfo)(nil), // 1: DownloadInfo
	(*DownloadResp)(nil), // 2: DownloadResp
}
var file_download_message_proto_depIdxs = []int32{
	1, // 0: DownloadResp.file:type_name -> DownloadInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_download_message_proto_init() }
func file_download_message_proto_init() {
	if File_download_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_download_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadReq); i {
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
		file_download_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadInfo); i {
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
		file_download_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadResp); i {
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
	file_download_message_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DownloadResp_File)(nil),
		(*DownloadResp_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_download_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_download_message_proto_goTypes,
		DependencyIndexes: file_download_message_proto_depIdxs,
		MessageInfos:      file_download_message_proto_msgTypes,
	}.Build()
	File_download_message_proto = out.File
	file_download_message_proto_rawDesc = nil
	file_download_message_proto_goTypes = nil
	file_download_message_proto_depIdxs = nil
}
