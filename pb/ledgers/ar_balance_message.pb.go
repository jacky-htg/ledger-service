// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.13.0
// source: ledgers/ar_balance_message.proto

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

type ArBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId        string  `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Period           string  `protobuf:"bytes,3,opt,name=period,proto3" json:"period,omitempty"`
	BeginningBalance float64 `protobuf:"fixed64,4,opt,name=beginning_balance,json=beginningBalance,proto3" json:"beginning_balance,omitempty"`
	Amount           float64 `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	EndingBalance    float64 `protobuf:"fixed64,6,opt,name=ending_balance,json=endingBalance,proto3" json:"ending_balance,omitempty"`
	ClosingAt        string  `protobuf:"bytes,7,opt,name=closing_at,json=closingAt,proto3" json:"closing_at,omitempty"`
	ClosingBy        string  `protobuf:"bytes,8,opt,name=closing_by,json=closingBy,proto3" json:"closing_by,omitempty"`
}

func (x *ArBalance) Reset() {
	*x = ArBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledgers_ar_balance_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArBalance) ProtoMessage() {}

func (x *ArBalance) ProtoReflect() protoreflect.Message {
	mi := &file_ledgers_ar_balance_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArBalance.ProtoReflect.Descriptor instead.
func (*ArBalance) Descriptor() ([]byte, []int) {
	return file_ledgers_ar_balance_message_proto_rawDescGZIP(), []int{0}
}

func (x *ArBalance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArBalance) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *ArBalance) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *ArBalance) GetBeginningBalance() float64 {
	if x != nil {
		return x.BeginningBalance
	}
	return 0
}

func (x *ArBalance) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *ArBalance) GetEndingBalance() float64 {
	if x != nil {
		return x.EndingBalance
	}
	return 0
}

func (x *ArBalance) GetClosingAt() string {
	if x != nil {
		return x.ClosingAt
	}
	return ""
}

func (x *ArBalance) GetClosingBy() string {
	if x != nil {
		return x.ClosingBy
	}
	return ""
}

type ArBalances struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance []*ArBalance `protobuf:"bytes,1,rep,name=balance,proto3" json:"balance,omitempty"`
}

func (x *ArBalances) Reset() {
	*x = ArBalances{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledgers_ar_balance_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArBalances) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArBalances) ProtoMessage() {}

func (x *ArBalances) ProtoReflect() protoreflect.Message {
	mi := &file_ledgers_ar_balance_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArBalances.ProtoReflect.Descriptor instead.
func (*ArBalances) Descriptor() ([]byte, []int) {
	return file_ledgers_ar_balance_message_proto_rawDescGZIP(), []int{1}
}

func (x *ArBalances) GetBalance() []*ArBalance {
	if x != nil {
		return x.Balance
	}
	return nil
}

var File_ledgers_ar_balance_message_proto protoreflect.FileDescriptor

var file_ledgers_ar_balance_message_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x72, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x73, 0x22, 0xfc, 0x01, 0x0a, 0x09, 0x41, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x62, 0x65, 0x67,
	0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x69,
	0x6e, 0x67, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x5f,
	0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e,
	0x67, 0x42, 0x79, 0x22, 0x43, 0x0a, 0x0a, 0x41, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x35, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x3b, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e,
	0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65, 0x72, 0x70, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x50,
	0x01, 0x5a, 0x12, 0x70, 0x62, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x3b, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ledgers_ar_balance_message_proto_rawDescOnce sync.Once
	file_ledgers_ar_balance_message_proto_rawDescData = file_ledgers_ar_balance_message_proto_rawDesc
)

func file_ledgers_ar_balance_message_proto_rawDescGZIP() []byte {
	file_ledgers_ar_balance_message_proto_rawDescOnce.Do(func() {
		file_ledgers_ar_balance_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_ledgers_ar_balance_message_proto_rawDescData)
	})
	return file_ledgers_ar_balance_message_proto_rawDescData
}

var file_ledgers_ar_balance_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ledgers_ar_balance_message_proto_goTypes = []interface{}{
	(*ArBalance)(nil),  // 0: wiradata.ledgers.ArBalance
	(*ArBalances)(nil), // 1: wiradata.ledgers.ArBalances
}
var file_ledgers_ar_balance_message_proto_depIdxs = []int32{
	0, // 0: wiradata.ledgers.ArBalances.balance:type_name -> wiradata.ledgers.ArBalance
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ledgers_ar_balance_message_proto_init() }
func file_ledgers_ar_balance_message_proto_init() {
	if File_ledgers_ar_balance_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ledgers_ar_balance_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArBalance); i {
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
		file_ledgers_ar_balance_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArBalances); i {
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
			RawDescriptor: file_ledgers_ar_balance_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ledgers_ar_balance_message_proto_goTypes,
		DependencyIndexes: file_ledgers_ar_balance_message_proto_depIdxs,
		MessageInfos:      file_ledgers_ar_balance_message_proto_msgTypes,
	}.Build()
	File_ledgers_ar_balance_message_proto = out.File
	file_ledgers_ar_balance_message_proto_rawDesc = nil
	file_ledgers_ar_balance_message_proto_goTypes = nil
	file_ledgers_ar_balance_message_proto_depIdxs = nil
}
