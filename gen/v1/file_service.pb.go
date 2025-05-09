// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: file_service.proto

package fsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_file_service_proto protoreflect.FileDescriptor

const file_file_service_proto_rawDesc = "" +
	"\n" +
	"\x12file_service.proto\x121github.redbloodpixel.filesexchange.fileservice.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\vmodel.proto2\xbc\x03\n" +
	"\vFileService\x12\x9b\x01\n" +
	"\n" +
	"UploadFile\x12D.github.redbloodpixel.filesexchange.fileservice.v1.UploadFileRequest\x1aE.github.redbloodpixel.filesexchange.fileservice.v1.UploadFileResponse\"\x00\x12k\n" +
	"\tListFiles\x12\x16.google.protobuf.Empty\x1aD.github.redbloodpixel.filesexchange.fileservice.v1.ListFilesResponse\"\x00\x12\xa1\x01\n" +
	"\fDownloadFile\x12F.github.redbloodpixel.filesexchange.fileservice.v1.DownloadFileRequest\x1aG.github.redbloodpixel.filesexchange.fileservice.v1.DownloadFileResponse\"\x00B6Z4github.com/redblood-pixel/FilesExchanger/gen/v1;fsv1b\x06proto3"

var file_file_service_proto_goTypes = []any{
	(*UploadFileRequest)(nil),    // 0: github.redbloodpixel.filesexchange.fileservice.v1.UploadFileRequest
	(*emptypb.Empty)(nil),        // 1: google.protobuf.Empty
	(*DownloadFileRequest)(nil),  // 2: github.redbloodpixel.filesexchange.fileservice.v1.DownloadFileRequest
	(*UploadFileResponse)(nil),   // 3: github.redbloodpixel.filesexchange.fileservice.v1.UploadFileResponse
	(*ListFilesResponse)(nil),    // 4: github.redbloodpixel.filesexchange.fileservice.v1.ListFilesResponse
	(*DownloadFileResponse)(nil), // 5: github.redbloodpixel.filesexchange.fileservice.v1.DownloadFileResponse
}
var file_file_service_proto_depIdxs = []int32{
	0, // 0: github.redbloodpixel.filesexchange.fileservice.v1.FileService.UploadFile:input_type -> github.redbloodpixel.filesexchange.fileservice.v1.UploadFileRequest
	1, // 1: github.redbloodpixel.filesexchange.fileservice.v1.FileService.ListFiles:input_type -> google.protobuf.Empty
	2, // 2: github.redbloodpixel.filesexchange.fileservice.v1.FileService.DownloadFile:input_type -> github.redbloodpixel.filesexchange.fileservice.v1.DownloadFileRequest
	3, // 3: github.redbloodpixel.filesexchange.fileservice.v1.FileService.UploadFile:output_type -> github.redbloodpixel.filesexchange.fileservice.v1.UploadFileResponse
	4, // 4: github.redbloodpixel.filesexchange.fileservice.v1.FileService.ListFiles:output_type -> github.redbloodpixel.filesexchange.fileservice.v1.ListFilesResponse
	5, // 5: github.redbloodpixel.filesexchange.fileservice.v1.FileService.DownloadFile:output_type -> github.redbloodpixel.filesexchange.fileservice.v1.DownloadFileResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_service_proto_init() }
func file_file_service_proto_init() {
	if File_file_service_proto != nil {
		return
	}
	file_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_file_service_proto_rawDesc), len(file_file_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_service_proto_goTypes,
		DependencyIndexes: file_file_service_proto_depIdxs,
	}.Build()
	File_file_service_proto = out.File
	file_file_service_proto_goTypes = nil
	file_file_service_proto_depIdxs = nil
}
