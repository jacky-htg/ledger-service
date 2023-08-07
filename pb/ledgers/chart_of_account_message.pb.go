// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.13.0
// source: ledgers/chart_of_account_message.proto

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

type AccountType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AccountType) Reset() {
	*x = AccountType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledgers_chart_of_account_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountType) ProtoMessage() {}

func (x *AccountType) ProtoReflect() protoreflect.Message {
	mi := &file_ledgers_chart_of_account_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountType.ProtoReflect.Descriptor instead.
func (*AccountType) Descriptor() ([]byte, []int) {
	return file_ledgers_chart_of_account_message_proto_rawDescGZIP(), []int{0}
}

func (x *AccountType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ChartOfAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId   string          `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	AccountType *AccountType    `protobuf:"bytes,3,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	Parent      *ChartOfAccount `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	Code        string          `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	Name        string          `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt   string          `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy   string          `protobuf:"bytes,8,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt   string          `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy   string          `protobuf:"bytes,10,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
}

func (x *ChartOfAccount) Reset() {
	*x = ChartOfAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledgers_chart_of_account_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartOfAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartOfAccount) ProtoMessage() {}

func (x *ChartOfAccount) ProtoReflect() protoreflect.Message {
	mi := &file_ledgers_chart_of_account_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartOfAccount.ProtoReflect.Descriptor instead.
func (*ChartOfAccount) Descriptor() ([]byte, []int) {
	return file_ledgers_chart_of_account_message_proto_rawDescGZIP(), []int{1}
}

func (x *ChartOfAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChartOfAccount) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *ChartOfAccount) GetAccountType() *AccountType {
	if x != nil {
		return x.AccountType
	}
	return nil
}

func (x *ChartOfAccount) GetParent() *ChartOfAccount {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *ChartOfAccount) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ChartOfAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChartOfAccount) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ChartOfAccount) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *ChartOfAccount) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ChartOfAccount) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

type ChartOfAccounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChartOfAccount []*ChartOfAccount `protobuf:"bytes,1,rep,name=chart_of_account,json=chartOfAccount,proto3" json:"chart_of_account,omitempty"`
}

func (x *ChartOfAccounts) Reset() {
	*x = ChartOfAccounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledgers_chart_of_account_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartOfAccounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartOfAccounts) ProtoMessage() {}

func (x *ChartOfAccounts) ProtoReflect() protoreflect.Message {
	mi := &file_ledgers_chart_of_account_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartOfAccounts.ProtoReflect.Descriptor instead.
func (*ChartOfAccounts) Descriptor() ([]byte, []int) {
	return file_ledgers_chart_of_account_message_proto_rawDescGZIP(), []int{2}
}

func (x *ChartOfAccounts) GetChartOfAccount() []*ChartOfAccount {
	if x != nil {
		return x.ChartOfAccount
	}
	return nil
}

var File_ledgers_chart_of_account_message_proto protoreflect.FileDescriptor

var file_ledgers_chart_of_account_message_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f,
	0x6f, 0x66, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x22, 0x31, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xdf, 0x02,
	0x0a, 0x0e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12,
	0x40, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x38, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22,
	0x5d, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x4a, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x77,
	0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x3b,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x72, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x65,
	0x72, 0x70, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x73, 0x50, 0x01, 0x5a, 0x12, 0x70, 0x62, 0x2f, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x73, 0x3b, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ledgers_chart_of_account_message_proto_rawDescOnce sync.Once
	file_ledgers_chart_of_account_message_proto_rawDescData = file_ledgers_chart_of_account_message_proto_rawDesc
)

func file_ledgers_chart_of_account_message_proto_rawDescGZIP() []byte {
	file_ledgers_chart_of_account_message_proto_rawDescOnce.Do(func() {
		file_ledgers_chart_of_account_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_ledgers_chart_of_account_message_proto_rawDescData)
	})
	return file_ledgers_chart_of_account_message_proto_rawDescData
}

var file_ledgers_chart_of_account_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ledgers_chart_of_account_message_proto_goTypes = []interface{}{
	(*AccountType)(nil),     // 0: wiradata.ledgers.AccountType
	(*ChartOfAccount)(nil),  // 1: wiradata.ledgers.ChartOfAccount
	(*ChartOfAccounts)(nil), // 2: wiradata.ledgers.ChartOfAccounts
}
var file_ledgers_chart_of_account_message_proto_depIdxs = []int32{
	0, // 0: wiradata.ledgers.ChartOfAccount.account_type:type_name -> wiradata.ledgers.AccountType
	1, // 1: wiradata.ledgers.ChartOfAccount.parent:type_name -> wiradata.ledgers.ChartOfAccount
	1, // 2: wiradata.ledgers.ChartOfAccounts.chart_of_account:type_name -> wiradata.ledgers.ChartOfAccount
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ledgers_chart_of_account_message_proto_init() }
func file_ledgers_chart_of_account_message_proto_init() {
	if File_ledgers_chart_of_account_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ledgers_chart_of_account_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountType); i {
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
		file_ledgers_chart_of_account_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartOfAccount); i {
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
		file_ledgers_chart_of_account_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChartOfAccounts); i {
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
			RawDescriptor: file_ledgers_chart_of_account_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ledgers_chart_of_account_message_proto_goTypes,
		DependencyIndexes: file_ledgers_chart_of_account_message_proto_depIdxs,
		MessageInfos:      file_ledgers_chart_of_account_message_proto_msgTypes,
	}.Build()
	File_ledgers_chart_of_account_message_proto = out.File
	file_ledgers_chart_of_account_message_proto_rawDesc = nil
	file_ledgers_chart_of_account_message_proto_goTypes = nil
	file_ledgers_chart_of_account_message_proto_depIdxs = nil
}