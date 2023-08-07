// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.13.0
// source: ledgers/acc_configuration_message.proto

package ledgers

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

type AccConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AccConfiguration) Reset() {
	*x = AccConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledgers_acc_configuration_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccConfiguration) ProtoMessage() {}

func (x *AccConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_ledgers_acc_configuration_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccConfiguration.ProtoReflect.Descriptor instead.
func (*AccConfiguration) Descriptor() ([]byte, []int) {
	return file_ledgers_acc_configuration_message_proto_rawDescGZIP(), []int{0}
}

func (x *AccConfiguration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AccConfigurationDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccConfigurationId string          `protobuf:"bytes,2,opt,name=acc_configuration_id,json=accConfigurationId,proto3" json:"acc_configuration_id,omitempty"`
	ChartOfAccount     *ChartOfAccount `protobuf:"bytes,3,opt,name=chart_of_account,json=chartOfAccount,proto3" json:"chart_of_account,omitempty"`
}

func (x *AccConfigurationDetail) Reset() {
	*x = AccConfigurationDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledgers_acc_configuration_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccConfigurationDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccConfigurationDetail) ProtoMessage() {}

func (x *AccConfigurationDetail) ProtoReflect() protoreflect.Message {
	mi := &file_ledgers_acc_configuration_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccConfigurationDetail.ProtoReflect.Descriptor instead.
func (*AccConfigurationDetail) Descriptor() ([]byte, []int) {
	return file_ledgers_acc_configuration_message_proto_rawDescGZIP(), []int{1}
}

func (x *AccConfigurationDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccConfigurationDetail) GetAccConfigurationId() string {
	if x != nil {
		return x.AccConfigurationId
	}
	return ""
}

func (x *AccConfigurationDetail) GetChartOfAccount() *ChartOfAccount {
	if x != nil {
		return x.ChartOfAccount
	}
	return nil
}

var File_ledgers_acc_configuration_message_proto protoreflect.FileDescriptor

var file_ledgers_acc_configuration_message_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x77, 0x69, 0x72, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x1a, 0x26, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x16,
	0x41, 0x63, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x63, 0x63, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x63, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x3b, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72, 0x70, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x12, 0x70,
	0x62, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x3b, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ledgers_acc_configuration_message_proto_rawDescOnce sync.Once
	file_ledgers_acc_configuration_message_proto_rawDescData = file_ledgers_acc_configuration_message_proto_rawDesc
)

func file_ledgers_acc_configuration_message_proto_rawDescGZIP() []byte {
	file_ledgers_acc_configuration_message_proto_rawDescOnce.Do(func() {
		file_ledgers_acc_configuration_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_ledgers_acc_configuration_message_proto_rawDescData)
	})
	return file_ledgers_acc_configuration_message_proto_rawDescData
}

var file_ledgers_acc_configuration_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ledgers_acc_configuration_message_proto_goTypes = []interface{}{
	(*AccConfiguration)(nil),       // 0: wiradata.ledgers.AccConfiguration
	(*AccConfigurationDetail)(nil), // 1: wiradata.ledgers.AccConfigurationDetail
	(*ChartOfAccount)(nil),         // 2: wiradata.ledgers.ChartOfAccount
}
var file_ledgers_acc_configuration_message_proto_depIdxs = []int32{
	2, // 0: wiradata.ledgers.AccConfigurationDetail.chart_of_account:type_name -> wiradata.ledgers.ChartOfAccount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ledgers_acc_configuration_message_proto_init() }
func file_ledgers_acc_configuration_message_proto_init() {
	if File_ledgers_acc_configuration_message_proto != nil {
		return
	}
	file_ledgers_chart_of_account_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ledgers_acc_configuration_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccConfiguration); i {
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
		file_ledgers_acc_configuration_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccConfigurationDetail); i {
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
			RawDescriptor: file_ledgers_acc_configuration_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ledgers_acc_configuration_message_proto_goTypes,
		DependencyIndexes: file_ledgers_acc_configuration_message_proto_depIdxs,
		MessageInfos:      file_ledgers_acc_configuration_message_proto_msgTypes,
	}.Build()
	File_ledgers_acc_configuration_message_proto = out.File
	file_ledgers_acc_configuration_message_proto_rawDesc = nil
	file_ledgers_acc_configuration_message_proto_goTypes = nil
	file_ledgers_acc_configuration_message_proto_depIdxs = nil
}