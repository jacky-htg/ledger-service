// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.13.0
// source: ledgers/bank_account_service.proto

package ledgers

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

var File_ledgers_bank_account_service_proto protoreflect.FileDescriptor

var file_ledgers_bank_account_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x1a, 0x1d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2f, 0x62,
	0x61, 0x6e, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa9, 0x02, 0x0a, 0x12, 0x42, 0x61,
	0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x48, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x77, 0x69, 0x72,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x61,
	0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x1d, 0x2e, 0x77, 0x69, 0x72, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x61, 0x6e,
	0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x14, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x1d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x61, 0x6e, 0x6b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x64, 0x1a, 0x1d, 0x2e, 0x77, 0x69, 0x72, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x61, 0x6e,
	0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1e, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x1d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x3b, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72, 0x70, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x12,
	0x70, 0x62, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x3b, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ledgers_bank_account_service_proto_goTypes = []interface{}{
	(*BankAccount)(nil),  // 0: wiradata.ledgers.BankAccount
	(*Id)(nil),           // 1: wiradata.ledgers.Id
	(*EmptyMessage)(nil), // 2: wiradata.ledgers.EmptyMessage
}
var file_ledgers_bank_account_service_proto_depIdxs = []int32{
	0, // 0: wiradata.ledgers.BankAccountService.Create:input_type -> wiradata.ledgers.BankAccount
	1, // 1: wiradata.ledgers.BankAccountService.View:input_type -> wiradata.ledgers.Id
	1, // 2: wiradata.ledgers.BankAccountService.Update:input_type -> wiradata.ledgers.Id
	2, // 3: wiradata.ledgers.BankAccountService.List:input_type -> wiradata.ledgers.EmptyMessage
	0, // 4: wiradata.ledgers.BankAccountService.Create:output_type -> wiradata.ledgers.BankAccount
	0, // 5: wiradata.ledgers.BankAccountService.View:output_type -> wiradata.ledgers.BankAccount
	0, // 6: wiradata.ledgers.BankAccountService.Update:output_type -> wiradata.ledgers.BankAccount
	0, // 7: wiradata.ledgers.BankAccountService.List:output_type -> wiradata.ledgers.BankAccount
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ledgers_bank_account_service_proto_init() }
func file_ledgers_bank_account_service_proto_init() {
	if File_ledgers_bank_account_service_proto != nil {
		return
	}
	file_ledgers_generic_message_proto_init()
	file_ledgers_bank_account_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ledgers_bank_account_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ledgers_bank_account_service_proto_goTypes,
		DependencyIndexes: file_ledgers_bank_account_service_proto_depIdxs,
	}.Build()
	File_ledgers_bank_account_service_proto = out.File
	file_ledgers_bank_account_service_proto_rawDesc = nil
	file_ledgers_bank_account_service_proto_goTypes = nil
	file_ledgers_bank_account_service_proto_depIdxs = nil
}