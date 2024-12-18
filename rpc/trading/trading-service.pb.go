// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.11
// source: proto/trading-service.proto

package trading

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

type GetOrderByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOrderByIdRequest) Reset() {
	*x = GetOrderByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_trading_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByIdRequest) ProtoMessage() {}

func (x *GetOrderByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trading_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderByIdRequest.ProtoReflect.Descriptor instead.
func (*GetOrderByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_trading_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOrderByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId           int32       `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	ClientCompanyId   string      `protobuf:"bytes,2,opt,name=clientCompanyId,proto3" json:"clientCompanyId,omitempty"`
	PaymentType       int32       `protobuf:"varint,3,opt,name=paymentType,proto3" json:"paymentType,omitempty"`
	OrderStatus       int32       `protobuf:"varint,4,opt,name=orderStatus,proto3" json:"orderStatus,omitempty"`
	ProcessedAt       string      `protobuf:"bytes,5,opt,name=processedAt,proto3" json:"processedAt,omitempty"`
	CargoId           string      `protobuf:"bytes,6,opt,name=cargoId,proto3" json:"cargoId,omitempty"`
	PreInvoiceId      string      `protobuf:"bytes,7,opt,name=preInvoiceId,proto3" json:"preInvoiceId,omitempty"`
	InvoiceId         string      `protobuf:"bytes,8,opt,name=invoiceId,proto3" json:"invoiceId,omitempty"`
	DeliveryPrice     int32       `protobuf:"varint,9,opt,name=deliveryPrice,proto3" json:"deliveryPrice,omitempty"`
	FullBoxesQuantity int32       `protobuf:"varint,10,opt,name=fullBoxesQuantity,proto3" json:"fullBoxesQuantity,omitempty"`
	Invoices          []*Invoices `protobuf:"bytes,11,rep,name=invoices,proto3" json:"invoices,omitempty"`
}

func (x *GetOrderByIdResponse) Reset() {
	*x = GetOrderByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_trading_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByIdResponse) ProtoMessage() {}

func (x *GetOrderByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trading_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderByIdResponse.ProtoReflect.Descriptor instead.
func (*GetOrderByIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_trading_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderByIdResponse) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *GetOrderByIdResponse) GetClientCompanyId() string {
	if x != nil {
		return x.ClientCompanyId
	}
	return ""
}

func (x *GetOrderByIdResponse) GetPaymentType() int32 {
	if x != nil {
		return x.PaymentType
	}
	return 0
}

func (x *GetOrderByIdResponse) GetOrderStatus() int32 {
	if x != nil {
		return x.OrderStatus
	}
	return 0
}

func (x *GetOrderByIdResponse) GetProcessedAt() string {
	if x != nil {
		return x.ProcessedAt
	}
	return ""
}

func (x *GetOrderByIdResponse) GetCargoId() string {
	if x != nil {
		return x.CargoId
	}
	return ""
}

func (x *GetOrderByIdResponse) GetPreInvoiceId() string {
	if x != nil {
		return x.PreInvoiceId
	}
	return ""
}

func (x *GetOrderByIdResponse) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *GetOrderByIdResponse) GetDeliveryPrice() int32 {
	if x != nil {
		return x.DeliveryPrice
	}
	return 0
}

func (x *GetOrderByIdResponse) GetFullBoxesQuantity() int32 {
	if x != nil {
		return x.FullBoxesQuantity
	}
	return 0
}

func (x *GetOrderByIdResponse) GetInvoices() []*Invoices {
	if x != nil {
		return x.Invoices
	}
	return nil
}

type Invoices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlantationId string      `protobuf:"bytes,1,opt,name=plantationId,proto3" json:"plantationId,omitempty"`
	Positions    []*Position `protobuf:"bytes,2,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *Invoices) Reset() {
	*x = Invoices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_trading_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoices) ProtoMessage() {}

func (x *Invoices) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trading_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoices.ProtoReflect.Descriptor instead.
func (*Invoices) Descriptor() ([]byte, []int) {
	return file_proto_trading_service_proto_rawDescGZIP(), []int{6}
}

func (x *Invoices) GetPlantationId() string {
	if x != nil {
		return x.PlantationId
	}
	return ""
}

func (x *Invoices) GetPositions() []*Position {
	if x != nil {
		return x.Positions
	}
	return nil
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AmountOfStems int32  `protobuf:"varint,1,opt,name=amountOfStems,proto3" json:"amountOfStems,omitempty"`
	ProductName   string `protobuf:"bytes,2,opt,name=productName,proto3" json:"productName,omitempty"`
	VarietyName   string `protobuf:"bytes,3,opt,name=varietyName,proto3" json:"varietyName,omitempty"`
	VarietySize   string `protobuf:"bytes,4,opt,name=varietySize,proto3" json:"varietySize,omitempty"`
	TypeOfLot     int32  `protobuf:"varint,5,opt,name=typeOfLot,proto3" json:"typeOfLot,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_trading_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_proto_trading_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_proto_trading_service_proto_rawDescGZIP(), []int{7}
}

func (x *Position) GetAmountOfStems() int32 {
	if x != nil {
		return x.AmountOfStems
	}
	return 0
}

func (x *Position) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Position) GetVarietyName() string {
	if x != nil {
		return x.VarietyName
	}
	return ""
}

func (x *Position) GetVarietySize() string {
	if x != nil {
		return x.VarietySize
	}
	return ""
}

func (x *Position) GetTypeOfLot() int32 {
	if x != nil {
		return x.TypeOfLot
	}
	return 0
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
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x97,
	0x03, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70,
	0x72, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x75, 0x6c, 0x6c, 0x42, 0x6f, 0x78, 0x65, 0x73,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x66, 0x75, 0x6c, 0x6c, 0x42, 0x6f, 0x78, 0x65, 0x73, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x25, 0x0a, 0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x08, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xb4, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24,
	0x0a, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x66, 0x53, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x66, 0x53,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x61, 0x72,
	0x69, 0x65, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x61, 0x72, 0x69,
	0x65, 0x74, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76,
	0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x79,
	0x70, 0x65, 0x4f, 0x66, 0x4c, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74,
	0x79, 0x70, 0x65, 0x4f, 0x66, 0x4c, 0x6f, 0x74, 0x32, 0xdf, 0x01, 0x0a, 0x0e, 0x54, 0x72, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x12, 0x53,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x69,
	0x64, 0x12, 0x1a, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x53, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61,
	0x69, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x72, 0x70,
	0x63, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x3b, 0x74, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_trading_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_trading_service_proto_goTypes = []interface{}{
	(*SetOrderStatusPaidRequest)(nil), // 0: SetOrderStatusPaidRequest
	(*SetOrderStatusPaidReply)(nil),   // 1: SetOrderStatusPaidReply
	(*GetOrderInvoiceRequest)(nil),    // 2: GetOrderInvoiceRequest
	(*GetOrderInvoiceReply)(nil),      // 3: GetOrderInvoiceReply
	(*GetOrderByIdRequest)(nil),       // 4: GetOrderByIdRequest
	(*GetOrderByIdResponse)(nil),      // 5: GetOrderByIdResponse
	(*Invoices)(nil),                  // 6: Invoices
	(*Position)(nil),                  // 7: Position
	nil,                               // 8: GetOrderInvoiceReply.PlantationRequestsEntry
}
var file_proto_trading_service_proto_depIdxs = []int32{
	8, // 0: GetOrderInvoiceReply.plantation_requests:type_name -> GetOrderInvoiceReply.PlantationRequestsEntry
	6, // 1: GetOrderByIdResponse.invoices:type_name -> Invoices
	7, // 2: Invoices.positions:type_name -> Position
	0, // 3: TradingService.SetOrderStatusPaid:input_type -> SetOrderStatusPaidRequest
	2, // 4: TradingService.GetOrderInvoice:input_type -> GetOrderInvoiceRequest
	2, // 5: TradingService.GetOrderById:input_type -> GetOrderInvoiceRequest
	1, // 6: TradingService.SetOrderStatusPaid:output_type -> SetOrderStatusPaidReply
	3, // 7: TradingService.GetOrderInvoice:output_type -> GetOrderInvoiceReply
	5, // 8: TradingService.GetOrderById:output_type -> GetOrderByIdResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
		file_proto_trading_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderByIdRequest); i {
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
		file_proto_trading_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderByIdResponse); i {
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
		file_proto_trading_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoices); i {
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
		file_proto_trading_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
			NumMessages:   9,
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
