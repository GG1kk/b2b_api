// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: setting-service.proto

package setting

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

type GetCountryCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryName string `protobuf:"bytes,1,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
}

func (x *GetCountryCodeRequest) Reset() {
	*x = GetCountryCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountryCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountryCodeRequest) ProtoMessage() {}

func (x *GetCountryCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountryCodeRequest.ProtoReflect.Descriptor instead.
func (*GetCountryCodeRequest) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCountryCodeRequest) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

type GetCountryCodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode string `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
}

func (x *GetCountryCodeReply) Reset() {
	*x = GetCountryCodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountryCodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountryCodeReply) ProtoMessage() {}

func (x *GetCountryCodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountryCodeReply.ProtoReflect.Descriptor instead.
func (*GetCountryCodeReply) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCountryCodeReply) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type GetCargoByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CargoId string `protobuf:"bytes,1,opt,name=cargo_id,json=cargoId,proto3" json:"cargo_id,omitempty"`
}

func (x *GetCargoByIDRequest) Reset() {
	*x = GetCargoByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCargoByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCargoByIDRequest) ProtoMessage() {}

func (x *GetCargoByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCargoByIDRequest.ProtoReflect.Descriptor instead.
func (*GetCargoByIDRequest) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCargoByIDRequest) GetCargoId() string {
	if x != nil {
		return x.CargoId
	}
	return ""
}

type GetCargoByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ToCity          string     `protobuf:"bytes,2,opt,name=to_city,json=toCity,proto3" json:"to_city,omitempty"`
	ToCountry       string     `protobuf:"bytes,3,opt,name=to_country,json=toCountry,proto3" json:"to_country,omitempty"`
	PreInvoicePrice int32      `protobuf:"varint,4,opt,name=pre_invoice_price,json=preInvoicePrice,proto3" json:"pre_invoice_price,omitempty"`
	PathTime        int32      `protobuf:"varint,5,opt,name=path_time,json=pathTime,proto3" json:"path_time,omitempty"`
	Handlers        []*Handler `protobuf:"bytes,6,rep,name=handlers,proto3" json:"handlers,omitempty"`
}

func (x *GetCargoByIDReply) Reset() {
	*x = GetCargoByIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCargoByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCargoByIDReply) ProtoMessage() {}

func (x *GetCargoByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCargoByIDReply.ProtoReflect.Descriptor instead.
func (*GetCargoByIDReply) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCargoByIDReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetCargoByIDReply) GetToCity() string {
	if x != nil {
		return x.ToCity
	}
	return ""
}

func (x *GetCargoByIDReply) GetToCountry() string {
	if x != nil {
		return x.ToCountry
	}
	return ""
}

func (x *GetCargoByIDReply) GetPreInvoicePrice() int32 {
	if x != nil {
		return x.PreInvoicePrice
	}
	return 0
}

func (x *GetCargoByIDReply) GetPathTime() int32 {
	if x != nil {
		return x.PathTime
	}
	return 0
}

func (x *GetCargoByIDReply) GetHandlers() []*Handler {
	if x != nil {
		return x.Handlers
	}
	return nil
}

type Handler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title           string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	PreCoolingPrice int32  `protobuf:"varint,3,opt,name=pre_cooling_price,json=preCoolingPrice,proto3" json:"pre_cooling_price,omitempty"`
}

func (x *Handler) Reset() {
	*x = Handler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handler) ProtoMessage() {}

func (x *Handler) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handler.ProtoReflect.Descriptor instead.
func (*Handler) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{4}
}

func (x *Handler) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Handler) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Handler) GetPreCoolingPrice() int32 {
	if x != nil {
		return x.PreCoolingPrice
	}
	return 0
}

type GetBasicSettingsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Periods            []*Period `protobuf:"bytes,2,rep,name=periods,proto3" json:"periods,omitempty"`
	WorkingDays        []bool    `protobuf:"varint,3,rep,packed,name=working_days,json=workingDays,proto3" json:"working_days,omitempty"`
	OrderPayTimer      int32     `protobuf:"varint,4,opt,name=order_pay_timer,json=orderPayTimer,proto3" json:"order_pay_timer,omitempty"`
	OrderAcceptTimer   int32     `protobuf:"varint,5,opt,name=order_accept_timer,json=orderAcceptTimer,proto3" json:"order_accept_timer,omitempty"`
	ScreenRefreshTimer int32     `protobuf:"varint,6,opt,name=screen_refresh_timer,json=screenRefreshTimer,proto3" json:"screen_refresh_timer,omitempty"`
}

func (x *GetBasicSettingsReply) Reset() {
	*x = GetBasicSettingsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBasicSettingsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBasicSettingsReply) ProtoMessage() {}

func (x *GetBasicSettingsReply) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBasicSettingsReply.ProtoReflect.Descriptor instead.
func (*GetBasicSettingsReply) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetBasicSettingsReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBasicSettingsReply) GetPeriods() []*Period {
	if x != nil {
		return x.Periods
	}
	return nil
}

func (x *GetBasicSettingsReply) GetWorkingDays() []bool {
	if x != nil {
		return x.WorkingDays
	}
	return nil
}

func (x *GetBasicSettingsReply) GetOrderPayTimer() int32 {
	if x != nil {
		return x.OrderPayTimer
	}
	return 0
}

func (x *GetBasicSettingsReply) GetOrderAcceptTimer() int32 {
	if x != nil {
		return x.OrderAcceptTimer
	}
	return 0
}

func (x *GetBasicSettingsReply) GetScreenRefreshTimer() int32 {
	if x != nil {
		return x.ScreenRefreshTimer
	}
	return 0
}

type Period struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartsAt int32 `protobuf:"varint,1,opt,name=starts_at,json=startsAt,proto3" json:"starts_at,omitempty"`
	EndsAt   int32 `protobuf:"varint,2,opt,name=ends_at,json=endsAt,proto3" json:"ends_at,omitempty"`
}

func (x *Period) Reset() {
	*x = Period{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Period) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Period) ProtoMessage() {}

func (x *Period) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Period.ProtoReflect.Descriptor instead.
func (*Period) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{6}
}

func (x *Period) GetStartsAt() int32 {
	if x != nil {
		return x.StartsAt
	}
	return 0
}

func (x *Period) GetEndsAt() int32 {
	if x != nil {
		return x.EndsAt
	}
	return 0
}

type GetPlantationIdsWithDisplayEnabledReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlantationIds []string `protobuf:"bytes,1,rep,name=plantation_ids,json=plantationIds,proto3" json:"plantation_ids,omitempty"`
}

func (x *GetPlantationIdsWithDisplayEnabledReply) Reset() {
	*x = GetPlantationIdsWithDisplayEnabledReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlantationIdsWithDisplayEnabledReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlantationIdsWithDisplayEnabledReply) ProtoMessage() {}

func (x *GetPlantationIdsWithDisplayEnabledReply) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlantationIdsWithDisplayEnabledReply.ProtoReflect.Descriptor instead.
func (*GetPlantationIdsWithDisplayEnabledReply) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetPlantationIdsWithDisplayEnabledReply) GetPlantationIds() []string {
	if x != nil {
		return x.PlantationIds
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{8}
}

var File_setting_service_proto protoreflect.FileDescriptor

var file_setting_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x30, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x49, 0x64, 0x22,
	0xca, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x63, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x43, 0x69, 0x74, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x2a, 0x0a,
	0x11, 0x70, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x74,
	0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x52, 0x08, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x22, 0x5b, 0x0a, 0x07,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2a, 0x0a,
	0x11, 0x70, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x43, 0x6f, 0x6f,
	0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xf5, 0x01, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x07, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x08, 0x52, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x79, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x72, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12,
	0x30, 0x0a, 0x14, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x73,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65,
	0x72, 0x22, 0x3e, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x73,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x73, 0x41,
	0x74, 0x22, 0x50, 0x0a, 0x27, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x73, 0x57, 0x69, 0x74, 0x68, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x97, 0x02, 0x0a,
	0x0f, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x56, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x73, 0x57, 0x69, 0x74, 0x68, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x67, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x67, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x32, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x15, 0x5a, 0x13, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x3b, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_setting_service_proto_rawDescOnce sync.Once
	file_setting_service_proto_rawDescData = file_setting_service_proto_rawDesc
)

func file_setting_service_proto_rawDescGZIP() []byte {
	file_setting_service_proto_rawDescOnce.Do(func() {
		file_setting_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_setting_service_proto_rawDescData)
	})
	return file_setting_service_proto_rawDescData
}

var file_setting_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_setting_service_proto_goTypes = []interface{}{
	(*GetCountryCodeRequest)(nil),                   // 0: GetCountryCodeRequest
	(*GetCountryCodeReply)(nil),                     // 1: GetCountryCodeReply
	(*GetCargoByIDRequest)(nil),                     // 2: GetCargoByIDRequest
	(*GetCargoByIDReply)(nil),                       // 3: GetCargoByIDReply
	(*Handler)(nil),                                 // 4: Handler
	(*GetBasicSettingsReply)(nil),                   // 5: GetBasicSettingsReply
	(*Period)(nil),                                  // 6: Period
	(*GetPlantationIdsWithDisplayEnabledReply)(nil), // 7: GetPlantationIdsWithDisplayEnabledReply
	(*Empty)(nil),                                   // 8: Empty
}
var file_setting_service_proto_depIdxs = []int32{
	4, // 0: GetCargoByIDReply.handlers:type_name -> Handler
	6, // 1: GetBasicSettingsReply.periods:type_name -> Period
	8, // 2: SettingsService.GetPlantationIdsWithDisplayEnabled:input_type -> Empty
	2, // 3: SettingsService.GetCargoByID:input_type -> GetCargoByIDRequest
	8, // 4: SettingsService.GetBasicSettings:input_type -> Empty
	0, // 5: SettingsService.GetCountryCode:input_type -> GetCountryCodeRequest
	7, // 6: SettingsService.GetPlantationIdsWithDisplayEnabled:output_type -> GetPlantationIdsWithDisplayEnabledReply
	3, // 7: SettingsService.GetCargoByID:output_type -> GetCargoByIDReply
	5, // 8: SettingsService.GetBasicSettings:output_type -> GetBasicSettingsReply
	1, // 9: SettingsService.GetCountryCode:output_type -> GetCountryCodeReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_setting_service_proto_init() }
func file_setting_service_proto_init() {
	if File_setting_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_setting_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountryCodeRequest); i {
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
		file_setting_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountryCodeReply); i {
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
		file_setting_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCargoByIDRequest); i {
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
		file_setting_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCargoByIDReply); i {
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
		file_setting_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handler); i {
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
		file_setting_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBasicSettingsReply); i {
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
		file_setting_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Period); i {
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
		file_setting_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlantationIdsWithDisplayEnabledReply); i {
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
		file_setting_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_setting_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_setting_service_proto_goTypes,
		DependencyIndexes: file_setting_service_proto_depIdxs,
		MessageInfos:      file_setting_service_proto_msgTypes,
	}.Build()
	File_setting_service_proto = out.File
	file_setting_service_proto_rawDesc = nil
	file_setting_service_proto_goTypes = nil
	file_setting_service_proto_depIdxs = nil
}
