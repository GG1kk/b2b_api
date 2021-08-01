// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: catalog-service.proto

package catalog

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

type GetCatalogProductByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCatalogProductByIDRequest) Reset() {
	*x = GetCatalogProductByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogProductByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogProductByIDRequest) ProtoMessage() {}

func (x *GetCatalogProductByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogProductByIDRequest.ProtoReflect.Descriptor instead.
func (*GetCatalogProductByIDRequest) Descriptor() ([]byte, []int) {
	return file_catalog_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCatalogProductByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCurrentSeasonByPlantationIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlantationId string `protobuf:"bytes,1,opt,name=plantation_id,json=plantationId,proto3" json:"plantation_id,omitempty"`
}

func (x *GetCurrentSeasonByPlantationIDRequest) Reset() {
	*x = GetCurrentSeasonByPlantationIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentSeasonByPlantationIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentSeasonByPlantationIDRequest) ProtoMessage() {}

func (x *GetCurrentSeasonByPlantationIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentSeasonByPlantationIDRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentSeasonByPlantationIDRequest) Descriptor() ([]byte, []int) {
	return file_catalog_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrentSeasonByPlantationIDRequest) GetPlantationId() string {
	if x != nil {
		return x.PlantationId
	}
	return ""
}

type GetCurrentSeasonByPlantationIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseMultiplier float64                   `protobuf:"fixed64,1,opt,name=base_multiplier,json=baseMultiplier,proto3" json:"base_multiplier,omitempty"`
	Modifiers      map[string]*VarietyToSize `protobuf:"bytes,2,rep,name=modifiers,proto3" json:"modifiers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetCurrentSeasonByPlantationIDReply) Reset() {
	*x = GetCurrentSeasonByPlantationIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentSeasonByPlantationIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentSeasonByPlantationIDReply) ProtoMessage() {}

func (x *GetCurrentSeasonByPlantationIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentSeasonByPlantationIDReply.ProtoReflect.Descriptor instead.
func (*GetCurrentSeasonByPlantationIDReply) Descriptor() ([]byte, []int) {
	return file_catalog_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCurrentSeasonByPlantationIDReply) GetBaseMultiplier() float64 {
	if x != nil {
		return x.BaseMultiplier
	}
	return 0
}

func (x *GetCurrentSeasonByPlantationIDReply) GetModifiers() map[string]*VarietyToSize {
	if x != nil {
		return x.Modifiers
	}
	return nil
}

type VarietyToSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dictionary map[string]*SizeToModifier `protobuf:"bytes,1,rep,name=dictionary,proto3" json:"dictionary,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VarietyToSize) Reset() {
	*x = VarietyToSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VarietyToSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VarietyToSize) ProtoMessage() {}

func (x *VarietyToSize) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VarietyToSize.ProtoReflect.Descriptor instead.
func (*VarietyToSize) Descriptor() ([]byte, []int) {
	return file_catalog_service_proto_rawDescGZIP(), []int{3}
}

func (x *VarietyToSize) GetDictionary() map[string]*SizeToModifier {
	if x != nil {
		return x.Dictionary
	}
	return nil
}

type SizeToModifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dictionary map[string]*Modifier `protobuf:"bytes,1,rep,name=dictionary,proto3" json:"dictionary,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SizeToModifier) Reset() {
	*x = SizeToModifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SizeToModifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SizeToModifier) ProtoMessage() {}

func (x *SizeToModifier) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SizeToModifier.ProtoReflect.Descriptor instead.
func (*SizeToModifier) Descriptor() ([]byte, []int) {
	return file_catalog_service_proto_rawDescGZIP(), []int{4}
}

func (x *SizeToModifier) GetDictionary() map[string]*Modifier {
	if x != nil {
		return x.Dictionary
	}
	return nil
}

type Modifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price     int32 `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	InPack    int32 `protobuf:"varint,2,opt,name=in_pack,json=inPack,proto3" json:"in_pack,omitempty"`
	InFullBox int32 `protobuf:"varint,3,opt,name=in_full_box,json=inFullBox,proto3" json:"in_full_box,omitempty"`
}

func (x *Modifier) Reset() {
	*x = Modifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Modifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Modifier) ProtoMessage() {}

func (x *Modifier) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Modifier.ProtoReflect.Descriptor instead.
func (*Modifier) Descriptor() ([]byte, []int) {
	return file_catalog_service_proto_rawDescGZIP(), []int{5}
}

func (x *Modifier) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Modifier) GetInPack() int32 {
	if x != nil {
		return x.InPack
	}
	return 0
}

func (x *Modifier) GetInFullBox() int32 {
	if x != nil {
		return x.InFullBox
	}
	return 0
}

type GetSizeInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlantationId     string `protobuf:"bytes,1,opt,name=plantation_id,json=plantationId,proto3" json:"plantation_id,omitempty"`
	CatalogProductId string `protobuf:"bytes,2,opt,name=catalog_product_id,json=catalogProductId,proto3" json:"catalog_product_id,omitempty"`
	ProductVarietyId string `protobuf:"bytes,3,opt,name=product_variety_id,json=productVarietyId,proto3" json:"product_variety_id,omitempty"`
	Size             string `protobuf:"bytes,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *GetSizeInfoRequest) Reset() {
	*x = GetSizeInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSizeInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSizeInfoRequest) ProtoMessage() {}

func (x *GetSizeInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSizeInfoRequest.ProtoReflect.Descriptor instead.
func (*GetSizeInfoRequest) Descriptor() ([]byte, []int) {
	return file_catalog_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetSizeInfoRequest) GetPlantationId() string {
	if x != nil {
		return x.PlantationId
	}
	return ""
}

func (x *GetSizeInfoRequest) GetCatalogProductId() string {
	if x != nil {
		return x.CatalogProductId
	}
	return ""
}

func (x *GetSizeInfoRequest) GetProductVarietyId() string {
	if x != nil {
		return x.ProductVarietyId
	}
	return ""
}

func (x *GetSizeInfoRequest) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

type GetSizeInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price     int32 `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	InPack    int32 `protobuf:"varint,2,opt,name=in_pack,json=inPack,proto3" json:"in_pack,omitempty"`
	InFullBox int32 `protobuf:"varint,3,opt,name=in_full_box,json=inFullBox,proto3" json:"in_full_box,omitempty"`
}

func (x *GetSizeInfoReply) Reset() {
	*x = GetSizeInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSizeInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSizeInfoReply) ProtoMessage() {}

func (x *GetSizeInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSizeInfoReply.ProtoReflect.Descriptor instead.
func (*GetSizeInfoReply) Descriptor() ([]byte, []int) {
	return file_catalog_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetSizeInfoReply) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetSizeInfoReply) GetInPack() int32 {
	if x != nil {
		return x.InPack
	}
	return 0
}

func (x *GetSizeInfoReply) GetInFullBox() int32 {
	if x != nil {
		return x.InFullBox
	}
	return 0
}

type GetCatalogProductByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName      string            `protobuf:"bytes,1,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	VarietyNames     map[string]string `protobuf:"bytes,2,rep,name=variety_names,json=varietyNames,proto3" json:"variety_names,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	VarietyImages    map[string]string `protobuf:"bytes,3,rep,name=variety_images,json=varietyImages,proto3" json:"variety_images,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	VarietyColorHexs map[string]string `protobuf:"bytes,4,rep,name=variety_color_hexs,json=varietyColorHexs,proto3" json:"variety_color_hexs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetCatalogProductByIDReply) Reset() {
	*x = GetCatalogProductByIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_catalog_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogProductByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogProductByIDReply) ProtoMessage() {}

func (x *GetCatalogProductByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_catalog_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogProductByIDReply.ProtoReflect.Descriptor instead.
func (*GetCatalogProductByIDReply) Descriptor() ([]byte, []int) {
	return file_catalog_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetCatalogProductByIDReply) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *GetCatalogProductByIDReply) GetVarietyNames() map[string]string {
	if x != nil {
		return x.VarietyNames
	}
	return nil
}

func (x *GetCatalogProductByIDReply) GetVarietyImages() map[string]string {
	if x != nil {
		return x.VarietyImages
	}
	return nil
}

func (x *GetCatalogProductByIDReply) GetVarietyColorHexs() map[string]string {
	if x != nil {
		return x.VarietyColorHexs
	}
	return nil
}

var File_catalog_service_proto protoreflect.FileDescriptor

var file_catalog_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x25, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xef, 0x01, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a,
	0x0f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x4c, 0x0a, 0x0e, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x56,
	0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x54, 0x6f, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9f, 0x01, 0x0a, 0x0d, 0x56, 0x61, 0x72, 0x69,
	0x65, 0x74, 0x79, 0x54, 0x6f, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x54, 0x6f, 0x53, 0x69, 0x7a, 0x65, 0x2e, 0x44, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x1a, 0x4e, 0x0a, 0x0f, 0x44, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x53, 0x69, 0x7a, 0x65, 0x54, 0x6f, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9b, 0x01, 0x0a, 0x0e, 0x53, 0x69,
	0x7a, 0x65, 0x54, 0x6f, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0a,
	0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x54, 0x6f, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0a, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x1a, 0x48, 0x0a,
	0x0f, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x59, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6e, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x6e, 0x50, 0x61,
	0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x6e, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x62, 0x6f,
	0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x42,
	0x6f, 0x78, 0x22, 0xa9, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c,
	0x0a, 0x12, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x61,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6e, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x6e, 0x50, 0x61, 0x63,
	0x6b, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x6e, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x62, 0x6f, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x42, 0x6f,
	0x78, 0x22, 0x93, 0x04, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x76, 0x61, 0x72, 0x69, 0x65,
	0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x0e, 0x76, 0x61, 0x72, 0x69, 0x65,
	0x74, 0x79, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x56, 0x61, 0x72,
	0x69, 0x65, 0x74, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x5f,
	0x0a, 0x12, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f,
	0x68, 0x65, 0x78, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x48, 0x65, 0x78, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x76,
	0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x48, 0x65, 0x78, 0x73, 0x1a,
	0x3f, 0x0a, 0x11, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x40, 0x0a, 0x12, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x43, 0x0a, 0x15, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x48, 0x65, 0x78, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x8c, 0x02, 0x0a, 0x0e, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6e, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x26, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x15, 0x5a, 0x13, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catalog_service_proto_rawDescOnce sync.Once
	file_catalog_service_proto_rawDescData = file_catalog_service_proto_rawDesc
)

func file_catalog_service_proto_rawDescGZIP() []byte {
	file_catalog_service_proto_rawDescOnce.Do(func() {
		file_catalog_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_catalog_service_proto_rawDescData)
	})
	return file_catalog_service_proto_rawDescData
}

var file_catalog_service_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_catalog_service_proto_goTypes = []interface{}{
	(*GetCatalogProductByIDRequest)(nil),          // 0: GetCatalogProductByIDRequest
	(*GetCurrentSeasonByPlantationIDRequest)(nil), // 1: GetCurrentSeasonByPlantationIDRequest
	(*GetCurrentSeasonByPlantationIDReply)(nil),   // 2: GetCurrentSeasonByPlantationIDReply
	(*VarietyToSize)(nil),                         // 3: VarietyToSize
	(*SizeToModifier)(nil),                        // 4: SizeToModifier
	(*Modifier)(nil),                              // 5: Modifier
	(*GetSizeInfoRequest)(nil),                    // 6: GetSizeInfoRequest
	(*GetSizeInfoReply)(nil),                      // 7: GetSizeInfoReply
	(*GetCatalogProductByIDReply)(nil),            // 8: GetCatalogProductByIDReply
	nil,                                           // 9: GetCurrentSeasonByPlantationIDReply.ModifiersEntry
	nil,                                           // 10: VarietyToSize.DictionaryEntry
	nil,                                           // 11: SizeToModifier.DictionaryEntry
	nil,                                           // 12: GetCatalogProductByIDReply.VarietyNamesEntry
	nil,                                           // 13: GetCatalogProductByIDReply.VarietyImagesEntry
	nil,                                           // 14: GetCatalogProductByIDReply.VarietyColorHexsEntry
}
var file_catalog_service_proto_depIdxs = []int32{
	9,  // 0: GetCurrentSeasonByPlantationIDReply.modifiers:type_name -> GetCurrentSeasonByPlantationIDReply.ModifiersEntry
	10, // 1: VarietyToSize.dictionary:type_name -> VarietyToSize.DictionaryEntry
	11, // 2: SizeToModifier.dictionary:type_name -> SizeToModifier.DictionaryEntry
	12, // 3: GetCatalogProductByIDReply.variety_names:type_name -> GetCatalogProductByIDReply.VarietyNamesEntry
	13, // 4: GetCatalogProductByIDReply.variety_images:type_name -> GetCatalogProductByIDReply.VarietyImagesEntry
	14, // 5: GetCatalogProductByIDReply.variety_color_hexs:type_name -> GetCatalogProductByIDReply.VarietyColorHexsEntry
	3,  // 6: GetCurrentSeasonByPlantationIDReply.ModifiersEntry.value:type_name -> VarietyToSize
	4,  // 7: VarietyToSize.DictionaryEntry.value:type_name -> SizeToModifier
	5,  // 8: SizeToModifier.DictionaryEntry.value:type_name -> Modifier
	0,  // 9: CatalogService.GetCatalogProductByID:input_type -> GetCatalogProductByIDRequest
	6,  // 10: CatalogService.GetSizeInfo:input_type -> GetSizeInfoRequest
	1,  // 11: CatalogService.GetCurrentSeasonByPlantationID:input_type -> GetCurrentSeasonByPlantationIDRequest
	8,  // 12: CatalogService.GetCatalogProductByID:output_type -> GetCatalogProductByIDReply
	7,  // 13: CatalogService.GetSizeInfo:output_type -> GetSizeInfoReply
	2,  // 14: CatalogService.GetCurrentSeasonByPlantationID:output_type -> GetCurrentSeasonByPlantationIDReply
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_catalog_service_proto_init() }
func file_catalog_service_proto_init() {
	if File_catalog_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_catalog_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogProductByIDRequest); i {
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
		file_catalog_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentSeasonByPlantationIDRequest); i {
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
		file_catalog_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentSeasonByPlantationIDReply); i {
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
		file_catalog_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VarietyToSize); i {
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
		file_catalog_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SizeToModifier); i {
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
		file_catalog_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Modifier); i {
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
		file_catalog_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSizeInfoRequest); i {
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
		file_catalog_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSizeInfoReply); i {
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
		file_catalog_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogProductByIDReply); i {
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
			RawDescriptor: file_catalog_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catalog_service_proto_goTypes,
		DependencyIndexes: file_catalog_service_proto_depIdxs,
		MessageInfos:      file_catalog_service_proto_msgTypes,
	}.Build()
	File_catalog_service_proto = out.File
	file_catalog_service_proto_rawDesc = nil
	file_catalog_service_proto_goTypes = nil
	file_catalog_service_proto_depIdxs = nil
}
