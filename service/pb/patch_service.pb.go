// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: patch_service.proto

package pb

import (
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

var File_patch_service_proto protoreflect.FileDescriptor

var file_patch_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x9e, 0x02, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x25, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x0c, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3a, 0x0a, 0x0b, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x13, 0x2e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x39, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x12, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x3c, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x13, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_patch_service_proto_goTypes = []interface{}{
	(*PingRequest)(nil),         // 0: PingRequest
	(*RightsCheckRequest)(nil),  // 1: RightsCheckRequest
	(*UploadFileRequest)(nil),   // 2: UploadFileRequest
	(*ApplyPatchRequest)(nil),   // 3: ApplyPatchRequest
	(*VerifyPatchRequest)(nil),  // 4: VerifyPatchRequest
	(*PingResponse)(nil),        // 5: PingResponse
	(*RightsCheckResponse)(nil), // 6: RightsCheckResponse
	(*UploadFileResponse)(nil),  // 7: UploadFileResponse
	(*ApplyPatchResponse)(nil),  // 8: ApplyPatchResponse
	(*VerifyPatchResponse)(nil), // 9: VerifyPatchResponse
}
var file_patch_service_proto_depIdxs = []int32{
	0, // 0: Patch.Ping:input_type -> PingRequest
	1, // 1: Patch.RightsCheck:input_type -> RightsCheckRequest
	2, // 2: Patch.UploadFile:input_type -> UploadFileRequest
	3, // 3: Patch.ApplyPatch:input_type -> ApplyPatchRequest
	4, // 4: Patch.VerifyPatch:input_type -> VerifyPatchRequest
	5, // 5: Patch.Ping:output_type -> PingResponse
	6, // 6: Patch.RightsCheck:output_type -> RightsCheckResponse
	7, // 7: Patch.UploadFile:output_type -> UploadFileResponse
	8, // 8: Patch.ApplyPatch:output_type -> ApplyPatchResponse
	9, // 9: Patch.VerifyPatch:output_type -> VerifyPatchResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_patch_service_proto_init() }
func file_patch_service_proto_init() {
	if File_patch_service_proto != nil {
		return
	}
	file_ping_message_proto_init()
	file_upload_message_proto_init()
	file_apply_message_proto_init()
	file_check_message_proto_init()
	file_verify_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_patch_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_patch_service_proto_goTypes,
		DependencyIndexes: file_patch_service_proto_depIdxs,
	}.Build()
	File_patch_service_proto = out.File
	file_patch_service_proto_rawDesc = nil
	file_patch_service_proto_goTypes = nil
	file_patch_service_proto_depIdxs = nil
}
