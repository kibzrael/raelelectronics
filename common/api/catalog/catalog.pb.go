// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.1
// source: api/catalog/catalog.proto

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

type FeedSort int32

const (
	FeedSort_Featured  FeedSort = 0
	FeedSort_New       FeedSort = 1
	FeedSort_PriceAsc  FeedSort = 2
	FeedSort_PriceDesc FeedSort = 3
)

// Enum value maps for FeedSort.
var (
	FeedSort_name = map[int32]string{
		0: "Featured",
		1: "New",
		2: "PriceAsc",
		3: "PriceDesc",
	}
	FeedSort_value = map[string]int32{
		"Featured":  0,
		"New":       1,
		"PriceAsc":  2,
		"PriceDesc": 3,
	}
)

func (x FeedSort) Enum() *FeedSort {
	p := new(FeedSort)
	*p = x
	return p
}

func (x FeedSort) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedSort) Descriptor() protoreflect.EnumDescriptor {
	return file_api_catalog_catalog_proto_enumTypes[0].Descriptor()
}

func (FeedSort) Type() protoreflect.EnumType {
	return &file_api_catalog_catalog_proto_enumTypes[0]
}

func (x FeedSort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedSort.Descriptor instead.
func (FeedSort) EnumDescriptor() ([]byte, []int) {
	return file_api_catalog_catalog_proto_rawDescGZIP(), []int{0}
}

type LaptopCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         string  `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Brand       string  `protobuf:"bytes,3,opt,name=Brand,proto3" json:"Brand,omitempty"`
	Thumbnail   string  `protobuf:"bytes,4,opt,name=Thumbnail,proto3" json:"Thumbnail,omitempty"`
	Launched    string  `protobuf:"bytes,5,opt,name=Launched,proto3" json:"Launched,omitempty"`
	PriceMin    float32 `protobuf:"fixed32,6,opt,name=PriceMin,proto3" json:"PriceMin,omitempty"`
	PriceMax    float32 `protobuf:"fixed32,7,opt,name=PriceMax,proto3" json:"PriceMax,omitempty"`
	Colors      string  `protobuf:"bytes,8,opt,name=Colors,proto3" json:"Colors,omitempty"`
	Size        float32 `protobuf:"fixed32,9,opt,name=Size,proto3" json:"Size,omitempty"`
	Cpu         string  `protobuf:"bytes,10,opt,name=Cpu,proto3" json:"Cpu,omitempty"`
	Cores       int64   `protobuf:"varint,11,opt,name=Cores,proto3" json:"Cores,omitempty"`
	BaseSpeed   float32 `protobuf:"fixed32,12,opt,name=BaseSpeed,proto3" json:"BaseSpeed,omitempty"`
	Memory      int64   `protobuf:"varint,13,opt,name=Memory,proto3" json:"Memory,omitempty"`
	Storage     int64   `protobuf:"varint,14,opt,name=Storage,proto3" json:"Storage,omitempty"`
	BatteryLife float32 `protobuf:"fixed32,15,opt,name=BatteryLife,proto3" json:"BatteryLife,omitempty"`
	Featured    float32 `protobuf:"fixed32,16,opt,name=Featured,proto3" json:"Featured,omitempty"`
}

func (x *LaptopCard) Reset() {
	*x = LaptopCard{}
	mi := &file_api_catalog_catalog_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LaptopCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaptopCard) ProtoMessage() {}

func (x *LaptopCard) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_catalog_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaptopCard.ProtoReflect.Descriptor instead.
func (*LaptopCard) Descriptor() ([]byte, []int) {
	return file_api_catalog_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *LaptopCard) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *LaptopCard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LaptopCard) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *LaptopCard) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *LaptopCard) GetLaunched() string {
	if x != nil {
		return x.Launched
	}
	return ""
}

func (x *LaptopCard) GetPriceMin() float32 {
	if x != nil {
		return x.PriceMin
	}
	return 0
}

func (x *LaptopCard) GetPriceMax() float32 {
	if x != nil {
		return x.PriceMax
	}
	return 0
}

func (x *LaptopCard) GetColors() string {
	if x != nil {
		return x.Colors
	}
	return ""
}

func (x *LaptopCard) GetSize() float32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *LaptopCard) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *LaptopCard) GetCores() int64 {
	if x != nil {
		return x.Cores
	}
	return 0
}

func (x *LaptopCard) GetBaseSpeed() float32 {
	if x != nil {
		return x.BaseSpeed
	}
	return 0
}

func (x *LaptopCard) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *LaptopCard) GetStorage() int64 {
	if x != nil {
		return x.Storage
	}
	return 0
}

func (x *LaptopCard) GetBatteryLife() float32 {
	if x != nil {
		return x.BatteryLife
	}
	return 0
}

func (x *LaptopCard) GetFeatured() float32 {
	if x != nil {
		return x.Featured
	}
	return 0
}

type FeedFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *FeedFilter) Reset() {
	*x = FeedFilter{}
	mi := &file_api_catalog_catalog_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedFilter) ProtoMessage() {}

func (x *FeedFilter) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_catalog_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedFilter.ProtoReflect.Descriptor instead.
func (*FeedFilter) Descriptor() ([]byte, []int) {
	return file_api_catalog_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *FeedFilter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FeedFilter) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type FeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters []*FeedFilter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
	Sort    FeedSort      `protobuf:"varint,2,opt,name=sort,proto3,enum=catalog.FeedSort" json:"sort,omitempty"`
	Page    int64         `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *FeedRequest) Reset() {
	*x = FeedRequest{}
	mi := &file_api_catalog_catalog_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedRequest) ProtoMessage() {}

func (x *FeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_catalog_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedRequest.ProtoReflect.Descriptor instead.
func (*FeedRequest) Descriptor() ([]byte, []int) {
	return file_api_catalog_catalog_proto_rawDescGZIP(), []int{2}
}

func (x *FeedRequest) GetFilters() []*FeedFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *FeedRequest) GetSort() FeedSort {
	if x != nil {
		return x.Sort
	}
	return FeedSort_Featured
}

func (x *FeedRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type FeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptops []*LaptopCard `protobuf:"bytes,1,rep,name=laptops,proto3" json:"laptops,omitempty"`
}

func (x *FeedResponse) Reset() {
	*x = FeedResponse{}
	mi := &file_api_catalog_catalog_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedResponse) ProtoMessage() {}

func (x *FeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_catalog_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedResponse.ProtoReflect.Descriptor instead.
func (*FeedResponse) Descriptor() ([]byte, []int) {
	return file_api_catalog_catalog_proto_rawDescGZIP(), []int{3}
}

func (x *FeedResponse) GetLaptops() []*LaptopCard {
	if x != nil {
		return x.Laptops
	}
	return nil
}

var File_api_catalog_catalog_proto protoreflect.FileDescriptor

var file_api_catalog_catalog_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x22, 0x9c, 0x03, 0x0a, 0x0a, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x4d, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x4d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4d, 0x61,
	0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4d, 0x61,
	0x78, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x43, 0x70, 0x75, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x70, 0x75, 0x12,
	0x14, 0x0a, 0x05, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65,
	0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x42, 0x61, 0x73, 0x65, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x4c, 0x69, 0x66, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x42, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x79, 0x4c, 0x69, 0x66, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x64, 0x22, 0x30, 0x0a, 0x0a, 0x46, 0x65, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x77, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x46, 0x65, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x46, 0x65, 0x65, 0x64,
	0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x3d,
	0x0a, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d,
	0x0a, 0x07, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x07, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x73, 0x2a, 0x3e, 0x0a,
	0x08, 0x46, 0x65, 0x65, 0x64, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x64, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x65, 0x77, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x73, 0x63, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x10, 0x03, 0x32, 0x4b, 0x0a,
	0x0e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x46, 0x65, 0x65, 0x64, 0x12, 0x14, 0x2e,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x62, 0x7a, 0x72, 0x61, 0x65,
	0x6c, 0x2f, 0x72, 0x61, 0x65, 0x6c, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_catalog_catalog_proto_rawDescOnce sync.Once
	file_api_catalog_catalog_proto_rawDescData = file_api_catalog_catalog_proto_rawDesc
)

func file_api_catalog_catalog_proto_rawDescGZIP() []byte {
	file_api_catalog_catalog_proto_rawDescOnce.Do(func() {
		file_api_catalog_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_catalog_catalog_proto_rawDescData)
	})
	return file_api_catalog_catalog_proto_rawDescData
}

var file_api_catalog_catalog_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_catalog_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_catalog_catalog_proto_goTypes = []any{
	(FeedSort)(0),        // 0: catalog.FeedSort
	(*LaptopCard)(nil),   // 1: catalog.LaptopCard
	(*FeedFilter)(nil),   // 2: catalog.FeedFilter
	(*FeedRequest)(nil),  // 3: catalog.FeedRequest
	(*FeedResponse)(nil), // 4: catalog.FeedResponse
}
var file_api_catalog_catalog_proto_depIdxs = []int32{
	2, // 0: catalog.FeedRequest.filters:type_name -> catalog.FeedFilter
	0, // 1: catalog.FeedRequest.sort:type_name -> catalog.FeedSort
	1, // 2: catalog.FeedResponse.laptops:type_name -> catalog.LaptopCard
	3, // 3: catalog.CatalogService.laptopFeed:input_type -> catalog.FeedRequest
	4, // 4: catalog.CatalogService.laptopFeed:output_type -> catalog.FeedResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_catalog_catalog_proto_init() }
func file_api_catalog_catalog_proto_init() {
	if File_api_catalog_catalog_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_catalog_catalog_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_catalog_catalog_proto_goTypes,
		DependencyIndexes: file_api_catalog_catalog_proto_depIdxs,
		EnumInfos:         file_api_catalog_catalog_proto_enumTypes,
		MessageInfos:      file_api_catalog_catalog_proto_msgTypes,
	}.Build()
	File_api_catalog_catalog_proto = out.File
	file_api_catalog_catalog_proto_rawDesc = nil
	file_api_catalog_catalog_proto_goTypes = nil
	file_api_catalog_catalog_proto_depIdxs = nil
}
