// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.1
// source: api/cart/cart.proto

package cart

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

type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User     string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Quantity int64  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	mi := &file_api_cart_cart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddToCartRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AddToCartRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AddToCartRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type AddToCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddToCartResponse) Reset() {
	*x = AddToCartResponse{}
	mi := &file_api_cart_cart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponse) ProtoMessage() {}

func (x *AddToCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponse.ProtoReflect.Descriptor instead.
func (*AddToCartResponse) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddToCartResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddToCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RemoveFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *RemoveFromCartRequest) Reset() {
	*x = RemoveFromCartRequest{}
	mi := &file_api_cart_cart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartRequest) ProtoMessage() {}

func (x *RemoveFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartRequest.ProtoReflect.Descriptor instead.
func (*RemoveFromCartRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveFromCartRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RemoveFromCartRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type RemoveFromCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoveFromCartResponse) Reset() {
	*x = RemoveFromCartResponse{}
	mi := &file_api_cart_cart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartResponse) ProtoMessage() {}

func (x *RemoveFromCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartResponse.ProtoReflect.Descriptor instead.
func (*RemoveFromCartResponse) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveFromCartResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemoveFromCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	mi := &file_api_cart_cart_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{4}
}

func (x *GetCartRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *GetCartRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type GetCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart []*CartItem `protobuf:"bytes,1,rep,name=cart,proto3" json:"cart,omitempty"`
}

func (x *GetCartResponse) Reset() {
	*x = GetCartResponse{}
	mi := &file_api_cart_cart_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartResponse) ProtoMessage() {}

func (x *GetCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartResponse.ProtoReflect.Descriptor instead.
func (*GetCartResponse) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{5}
}

func (x *GetCartResponse) GetCart() []*CartItem {
	if x != nil {
		return x.Cart
	}
	return nil
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid          string `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Laptop       string `protobuf:"bytes,2,opt,name=Laptop,proto3" json:"Laptop,omitempty"`
	Quantity     int64  `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	CPU          string `protobuf:"bytes,4,opt,name=CPU,proto3" json:"CPU,omitempty"`
	GPU          string `protobuf:"bytes,5,opt,name=GPU,proto3" json:"GPU,omitempty"`
	Display      string `protobuf:"bytes,6,opt,name=Display,proto3" json:"Display,omitempty"`
	Memory       string `protobuf:"bytes,7,opt,name=Memory,proto3" json:"Memory,omitempty"`
	Storage      string `protobuf:"bytes,8,opt,name=Storage,proto3" json:"Storage,omitempty"`
	WirelessCard string `protobuf:"bytes,9,opt,name=WirelessCard,proto3" json:"WirelessCard,omitempty"`
	Motherboard  string `protobuf:"bytes,10,opt,name=Motherboard,proto3" json:"Motherboard,omitempty"`
	Chasis       string `protobuf:"bytes,11,opt,name=Chasis,proto3" json:"Chasis,omitempty"`
	Battery      string `protobuf:"bytes,12,opt,name=Battery,proto3" json:"Battery,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	mi := &file_api_cart_cart_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{6}
}

func (x *CartItem) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CartItem) GetLaptop() string {
	if x != nil {
		return x.Laptop
	}
	return ""
}

func (x *CartItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartItem) GetCPU() string {
	if x != nil {
		return x.CPU
	}
	return ""
}

func (x *CartItem) GetGPU() string {
	if x != nil {
		return x.GPU
	}
	return ""
}

func (x *CartItem) GetDisplay() string {
	if x != nil {
		return x.Display
	}
	return ""
}

func (x *CartItem) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *CartItem) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *CartItem) GetWirelessCard() string {
	if x != nil {
		return x.WirelessCard
	}
	return ""
}

func (x *CartItem) GetMotherboard() string {
	if x != nil {
		return x.Motherboard
	}
	return ""
}

func (x *CartItem) GetChasis() string {
	if x != nil {
		return x.Chasis
	}
	return ""
}

func (x *CartItem) GetBattery() string {
	if x != nil {
		return x.Battery
	}
	return ""
}

type AddToWishlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AddToWishlistRequest) Reset() {
	*x = AddToWishlistRequest{}
	mi := &file_api_cart_cart_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToWishlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToWishlistRequest) ProtoMessage() {}

func (x *AddToWishlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToWishlistRequest.ProtoReflect.Descriptor instead.
func (*AddToWishlistRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{7}
}

func (x *AddToWishlistRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *AddToWishlistRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type AddToWishlistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddToWishlistResponse) Reset() {
	*x = AddToWishlistResponse{}
	mi := &file_api_cart_cart_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToWishlistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToWishlistResponse) ProtoMessage() {}

func (x *AddToWishlistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToWishlistResponse.ProtoReflect.Descriptor instead.
func (*AddToWishlistResponse) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{8}
}

func (x *AddToWishlistResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddToWishlistResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RemoveFromWishlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *RemoveFromWishlistRequest) Reset() {
	*x = RemoveFromWishlistRequest{}
	mi := &file_api_cart_cart_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromWishlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromWishlistRequest) ProtoMessage() {}

func (x *RemoveFromWishlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromWishlistRequest.ProtoReflect.Descriptor instead.
func (*RemoveFromWishlistRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveFromWishlistRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RemoveFromWishlistRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type RemoveFromWishlistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RemoveFromWishlistResponse) Reset() {
	*x = RemoveFromWishlistResponse{}
	mi := &file_api_cart_cart_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromWishlistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromWishlistResponse) ProtoMessage() {}

func (x *RemoveFromWishlistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromWishlistResponse.ProtoReflect.Descriptor instead.
func (*RemoveFromWishlistResponse) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveFromWishlistResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemoveFromWishlistResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetWishlistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetWishlistRequest) Reset() {
	*x = GetWishlistRequest{}
	mi := &file_api_cart_cart_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWishlistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWishlistRequest) ProtoMessage() {}

func (x *GetWishlistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWishlistRequest.ProtoReflect.Descriptor instead.
func (*GetWishlistRequest) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{11}
}

func (x *GetWishlistRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *GetWishlistRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type GetWishlistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wishlist []*WishlistItem `protobuf:"bytes,1,rep,name=wishlist,proto3" json:"wishlist,omitempty"`
}

func (x *GetWishlistResponse) Reset() {
	*x = GetWishlistResponse{}
	mi := &file_api_cart_cart_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWishlistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWishlistResponse) ProtoMessage() {}

func (x *GetWishlistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWishlistResponse.ProtoReflect.Descriptor instead.
func (*GetWishlistResponse) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{12}
}

func (x *GetWishlistResponse) GetWishlist() []*WishlistItem {
	if x != nil {
		return x.Wishlist
	}
	return nil
}

type WishlistItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Laptop string `protobuf:"bytes,2,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *WishlistItem) Reset() {
	*x = WishlistItem{}
	mi := &file_api_cart_cart_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WishlistItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishlistItem) ProtoMessage() {}

func (x *WishlistItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_cart_cart_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishlistItem.ProtoReflect.Descriptor instead.
func (*WishlistItem) Descriptor() ([]byte, []int) {
	return file_api_cart_cart_proto_rawDescGZIP(), []int{13}
}

func (x *WishlistItem) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *WishlistItem) GetLaptop() string {
	if x != nil {
		return x.Laptop
	}
	return ""
}

var File_api_cart_cart_proto protoreflect.FileDescriptor

var file_api_cart_cart_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x54, 0x0a, 0x10, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0x47, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3d, 0x0a, 0x15, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x4c, 0x0a, 0x16, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x36, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0xb8, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x50, 0x55,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x50, 0x55, 0x12, 0x10, 0x0a, 0x03, 0x47,
	0x50, 0x55, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x47, 0x50, 0x55, 0x12, 0x18, 0x0a,
	0x07, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x57, 0x69, 0x72,
	0x65, 0x6c, 0x65, 0x73, 0x73, 0x43, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x57, 0x69, 0x72, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x43, 0x61, 0x72, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x4d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x4d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x73, 0x69, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x43, 0x68, 0x61, 0x73, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x79, 0x22, 0x3c, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x4b, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x41, 0x0a, 0x19,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x50, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x69, 0x73,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x3a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x45, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x57, 0x69,
	0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x77, 0x69, 0x73, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x0c, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x32, 0xb7,
	0x03, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1b,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x72,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x48, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57,
	0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x57, 0x69, 0x73, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x72, 0x6f, 0x6d, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x57, 0x69, 0x73, 0x68, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x69,
	0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x62, 0x7a, 0x72, 0x61, 0x65, 0x6c, 0x2f,
	0x72, 0x61, 0x65, 0x6c, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cart_cart_proto_rawDescOnce sync.Once
	file_api_cart_cart_proto_rawDescData = file_api_cart_cart_proto_rawDesc
)

func file_api_cart_cart_proto_rawDescGZIP() []byte {
	file_api_cart_cart_proto_rawDescOnce.Do(func() {
		file_api_cart_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cart_cart_proto_rawDescData)
	})
	return file_api_cart_cart_proto_rawDescData
}

var file_api_cart_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_cart_cart_proto_goTypes = []any{
	(*AddToCartRequest)(nil),           // 0: cart.AddToCartRequest
	(*AddToCartResponse)(nil),          // 1: cart.AddToCartResponse
	(*RemoveFromCartRequest)(nil),      // 2: cart.RemoveFromCartRequest
	(*RemoveFromCartResponse)(nil),     // 3: cart.RemoveFromCartResponse
	(*GetCartRequest)(nil),             // 4: cart.GetCartRequest
	(*GetCartResponse)(nil),            // 5: cart.GetCartResponse
	(*CartItem)(nil),                   // 6: cart.CartItem
	(*AddToWishlistRequest)(nil),       // 7: cart.AddToWishlistRequest
	(*AddToWishlistResponse)(nil),      // 8: cart.AddToWishlistResponse
	(*RemoveFromWishlistRequest)(nil),  // 9: cart.RemoveFromWishlistRequest
	(*RemoveFromWishlistResponse)(nil), // 10: cart.RemoveFromWishlistResponse
	(*GetWishlistRequest)(nil),         // 11: cart.GetWishlistRequest
	(*GetWishlistResponse)(nil),        // 12: cart.GetWishlistResponse
	(*WishlistItem)(nil),               // 13: cart.WishlistItem
}
var file_api_cart_cart_proto_depIdxs = []int32{
	6,  // 0: cart.GetCartResponse.cart:type_name -> cart.CartItem
	13, // 1: cart.GetWishlistResponse.wishlist:type_name -> cart.WishlistItem
	0,  // 2: cart.CartService.AddToCart:input_type -> cart.AddToCartRequest
	2,  // 3: cart.CartService.RemoveFromCart:input_type -> cart.RemoveFromCartRequest
	4,  // 4: cart.CartService.GetCart:input_type -> cart.GetCartRequest
	7,  // 5: cart.CartService.AddToWishlist:input_type -> cart.AddToWishlistRequest
	9,  // 6: cart.CartService.RemoveFromWishlist:input_type -> cart.RemoveFromWishlistRequest
	11, // 7: cart.CartService.GetWishlist:input_type -> cart.GetWishlistRequest
	1,  // 8: cart.CartService.AddToCart:output_type -> cart.AddToCartResponse
	3,  // 9: cart.CartService.RemoveFromCart:output_type -> cart.RemoveFromCartResponse
	5,  // 10: cart.CartService.GetCart:output_type -> cart.GetCartResponse
	8,  // 11: cart.CartService.AddToWishlist:output_type -> cart.AddToWishlistResponse
	10, // 12: cart.CartService.RemoveFromWishlist:output_type -> cart.RemoveFromWishlistResponse
	12, // 13: cart.CartService.GetWishlist:output_type -> cart.GetWishlistResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_cart_cart_proto_init() }
func file_api_cart_cart_proto_init() {
	if File_api_cart_cart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_cart_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_cart_cart_proto_goTypes,
		DependencyIndexes: file_api_cart_cart_proto_depIdxs,
		MessageInfos:      file_api_cart_cart_proto_msgTypes,
	}.Build()
	File_api_cart_cart_proto = out.File
	file_api_cart_cart_proto_rawDesc = nil
	file_api_cart_cart_proto_goTypes = nil
	file_api_cart_cart_proto_depIdxs = nil
}
