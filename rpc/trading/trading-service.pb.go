// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: proto/trading-service.proto

package trading

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SetOrderStatusPaidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int32 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *SetOrderStatusPaidRequest) Reset() {
	*x = SetOrderStatusPaidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_trading_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOrderStatusPaidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOrderStatusPaidRequest) ProtoMessage() {}

func (x *SetOrderStatusPaidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trading_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOrderStatusPaidRequest.ProtoReflect.Descriptor instead.
func (*SetOrderStatusPaidRequest) Descriptor() ([]byte, []int) {
	return file_proto_trading_service_proto_rawDescGZIP(), []int{0}
}

func (x *SetOrderStatusPaidRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type SetOrderStatusPaidReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetOrderStatusPaidReply) Reset() {
	*x = SetOrderStatusPaidReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_trading_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOrderStatusPaidReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOrderStatusPaidReply) ProtoMessage() {}

func (x *SetOrderStatusPaidReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trading_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOrderStatusPaidReply.ProtoReflect.Descriptor instead.
func (*SetOrderStatusPaidReply) Descriptor() ([]byte, []int) {
	return file_proto_trading_service_proto_rawDescGZIP(), []int{1}
}

type GetOrderInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId int32 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *GetOrderInvoiceRequest) Reset() {
	*x = GetOrderInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_trading_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderInvoiceRequest) ProtoMessage() {}

func (x *GetOrderInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trading_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderInvoiceRequest.ProtoReflect.Descriptor instead.
func (*GetOrderInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_trading_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderInvoiceRequest) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type GetOrderInvoiceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total              int32            `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Commission         int32            `protobuf:"varint,2,opt,name=commission,proto3" json:"commission,omitempty"`
	Delivery           int32            `protobuf:"varint,3,opt,name=delivery,proto3" json:"delivery,omitempty"`
	PlantationRequests map[string]int32 `protobuf:"bytes,4,rep,name=plantation_requests,json=plantationRequests,proto3" json:"plantation_requests,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GetOrderInvoiceReply) Reset() {
	*x = GetOrderInvoiceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_trading_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderInvoiceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderInvoiceReply) ProtoMessage() {}

func (x *GetOrderInvoiceReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trading_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderInvoiceReply.ProtoReflect.Descriptor instead.
func (*GetOrderInvoiceReply) Descriptor() ([]byte, []int) {
	return file_proto_trading_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderInvoiceReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetOrderInvoiceReply) GetCommission() int32 {
	if x != nil {
		return x.Commission
	}
	return 0
}

func (x *GetOrderInvoiceReply) GetDelivery() int32 {
	if x != nil {
		return x.Delivery
	}
	return 0
}

func (x *GetOrderInvoiceReply) GetPlantationRequests() map[string]int32 {
	if x != nil {
		return x.PlantationRequests
	}
	return nil
}

var File_proto_trading_service_proto protoreflect.FileDescriptor

var file_proto_trading_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x19, 0x53, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50,
	0x61, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x33, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8f, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x12, 0x5e, 0x0a, 0x13, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x1a, 0x45, 0x0a, 0x17, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x9f, 0x01, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x53, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x69, 0x64,
	0x12, 0x1a, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x53,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x69,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x15, 0x5a, 0x13, 0x72, 0x70, 0x63,
	0x2f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x3b, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_trading_service_proto_rawDescOnce sync.Once
	file_proto_trading_service_proto_rawDescData = file_proto_trading_service_proto_rawDesc
)

func file_proto_trading_service_proto_rawDescGZIP() []byte {
	file_proto_trading_service_proto_rawDescOnce.Do(func() {
		file_proto_trading_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_trading_service_proto_rawDescData)
	})
	return file_proto_trading_service_proto_rawDescData
}

var file_proto_trading_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_trading_service_proto_goTypes = []interface{}{
	(*SetOrderStatusPaidRequest)(nil), // 0: SetOrderStatusPaidRequest
	(*SetOrderStatusPaidReply)(nil),   // 1: SetOrderStatusPaidReply
	(*GetOrderInvoiceRequest)(nil),    // 2: GetOrderInvoiceRequest
	(*GetOrderInvoiceReply)(nil),      // 3: GetOrderInvoiceReply
	nil,                               // 4: GetOrderInvoiceReply.PlantationRequestsEntry
}
var file_proto_trading_service_proto_depIdxs = []int32{
	4, // 0: GetOrderInvoiceReply.plantation_requests:type_name -> GetOrderInvoiceReply.PlantationRequestsEntry
	0, // 1: TradingService.SetOrderStatusPaid:input_type -> SetOrderStatusPaidRequest
	2, // 2: TradingService.GetOrderInvoice:input_type -> GetOrderInvoiceRequest
	1, // 3: TradingService.SetOrderStatusPaid:output_type -> SetOrderStatusPaidReply
	3, // 4: TradingService.GetOrderInvoice:output_type -> GetOrderInvoiceReply
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_trading_service_proto_init() }
func file_proto_trading_service_proto_init() {
	if File_proto_trading_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_trading_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOrderStatusPaidRequest); i {
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
		file_proto_trading_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOrderStatusPaidReply); i {
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
		file_proto_trading_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderInvoiceRequest); i {
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
		file_proto_trading_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderInvoiceReply); i {
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
			RawDescriptor: file_proto_trading_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_trading_service_proto_goTypes,
		DependencyIndexes: file_proto_trading_service_proto_depIdxs,
		MessageInfos:      file_proto_trading_service_proto_msgTypes,
	}.Build()
	File_proto_trading_service_proto = out.File
	file_proto_trading_service_proto_rawDesc = nil
	file_proto_trading_service_proto_goTypes = nil
	file_proto_trading_service_proto_depIdxs = nil
}
