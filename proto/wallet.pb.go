// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: wallet.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SendTransactionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account"`
	Amount  string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (x *SendTransactionInfo) Reset() {
	*x = SendTransactionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransactionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransactionInfo) ProtoMessage() {}

func (x *SendTransactionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransactionInfo.ProtoReflect.Descriptor instead.
func (*SendTransactionInfo) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *SendTransactionInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SendTransactionInfo) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confirmed   string `protobuf:"bytes,1,opt,name=confirmed,proto3" json:"confirmed"`
	Unconfirmed string `protobuf:"bytes,2,opt,name=unconfirmed,proto3" json:"unconfirmed"`
	Locked      string `protobuf:"bytes,3,opt,name=locked,proto3" json:"locked"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *Balance) GetConfirmed() string {
	if x != nil {
		return x.Confirmed
	}
	return ""
}

func (x *Balance) GetUnconfirmed() string {
	if x != nil {
		return x.Unconfirmed
	}
	return ""
}

func (x *Balance) GetLocked() string {
	if x != nil {
		return x.Locked
	}
	return ""
}

type Wallets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallets []string `protobuf:"bytes,1,rep,name=wallets,proto3" json:"wallets"`
}

func (x *Wallets) Reset() {
	*x = Wallets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallets) ProtoMessage() {}

func (x *Wallets) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallets.ProtoReflect.Descriptor instead.
func (*Wallets) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *Wallets) GetWallets() []string {
	if x != nil {
		return x.Wallets
	}
	return nil
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ImportWalletData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Key  *KeyPair `protobuf:"bytes,2,opt,name=key,proto3" json:"key"`
}

func (x *ImportWalletData) Reset() {
	*x = ImportWalletData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportWalletData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportWalletData) ProtoMessage() {}

func (x *ImportWalletData) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportWalletData.ProtoReflect.Descriptor instead.
func (*ImportWalletData) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *ImportWalletData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImportWalletData) GetKey() *KeyPair {
	if x != nil {
		return x.Key
	}
	return nil
}

var File_wallet_proto protoreflect.FileDescriptor

var file_wallet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x13, 0x53, 0x65,
	0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x61, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x23, 0x0a, 0x07, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x22, 0x1a, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x10, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4b,
	0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x32, 0xd6, 0x07, 0x0a, 0x06,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12,
	0x13, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x12, 0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x08, 0x2e, 0x4b, 0x65,
	0x79, 0x50, 0x61, 0x69, 0x72, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x3c, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x12, 0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x08, 0x2e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x3a, 0x01, 0x2a, 0x12, 0x49, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69,
	0x72, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12,
	0x3a, 0x0a, 0x0a, 0x44, 0x75, 0x6d, 0x70, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x06, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2f, 0x64, 0x75, 0x6d, 0x70, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x08, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x12, 0x12, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61,
	0x69, 0x72, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x52,
	0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f,
	0x73, 0x65, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x01, 0x2a, 0x12, 0x47, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x08, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x08,
	0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x22, 0x16, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x50, 0x0a, 0x12, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x75, 0x6c,
	0x6b, 0x12, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x73, 0x1a, 0x08, 0x2e, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a,
	0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x75, 0x6c, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x45, 0x0a,
	0x0d, 0x45, 0x78, 0x69, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x08,
	0x2e, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x2f, 0x65, 0x78, 0x69, 0x74, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x3a, 0x01, 0x2a, 0x12, 0x4e, 0x0a, 0x11, 0x45, 0x78, 0x69, 0x74, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x09, 0x2e, 0x4b, 0x65, 0x79, 0x50,
	0x61, 0x69, 0x72, 0x73, 0x1a, 0x08, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f,
	0x65, 0x78, 0x69, 0x74, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x75, 0x6c,
	0x6b, 0x3a, 0x01, 0x2a, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x79, 0x6d, 0x70, 0x75, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_proto_rawDescOnce sync.Once
	file_wallet_proto_rawDescData = file_wallet_proto_rawDesc
)

func file_wallet_proto_rawDescGZIP() []byte {
	file_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_proto_rawDescData)
	})
	return file_wallet_proto_rawDescData
}

var file_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_wallet_proto_goTypes = []interface{}{
	(*SendTransactionInfo)(nil), // 0: SendTransactionInfo
	(*Balance)(nil),             // 1: Balance
	(*Wallets)(nil),             // 2: Wallets
	(*Name)(nil),                // 3: Name
	(*ImportWalletData)(nil),    // 4: ImportWalletData
	(*KeyPair)(nil),             // 5: KeyPair
	(*Empty)(nil),               // 6: Empty
	(*KeyPairs)(nil),            // 7: KeyPairs
	(*Success)(nil),             // 8: Success
	(*ValidatorsRegistry)(nil),  // 9: ValidatorsRegistry
	(*Hash)(nil),                // 10: Hash
}
var file_wallet_proto_depIdxs = []int32{
	5,  // 0: ImportWalletData.key:type_name -> KeyPair
	6,  // 1: Wallet.ListWallets:input_type -> Empty
	3,  // 2: Wallet.CreateWallet:input_type -> Name
	3,  // 3: Wallet.OpenWallet:input_type -> Name
	4,  // 4: Wallet.ImportWallet:input_type -> ImportWalletData
	6,  // 5: Wallet.DumpWallet:input_type -> Empty
	6,  // 6: Wallet.CloseWallet:input_type -> Empty
	6,  // 7: Wallet.GetBalance:input_type -> Empty
	6,  // 8: Wallet.GetValidators:input_type -> Empty
	6,  // 9: Wallet.GetAccount:input_type -> Empty
	0,  // 10: Wallet.SendTransaction:input_type -> SendTransactionInfo
	5,  // 11: Wallet.StartValidator:input_type -> KeyPair
	7,  // 12: Wallet.StartValidatorBulk:input_type -> KeyPairs
	5,  // 13: Wallet.ExitValidator:input_type -> KeyPair
	7,  // 14: Wallet.ExitValidatorBulk:input_type -> KeyPairs
	2,  // 15: Wallet.ListWallets:output_type -> Wallets
	5,  // 16: Wallet.CreateWallet:output_type -> KeyPair
	8,  // 17: Wallet.OpenWallet:output_type -> Success
	5,  // 18: Wallet.ImportWallet:output_type -> KeyPair
	5,  // 19: Wallet.DumpWallet:output_type -> KeyPair
	8,  // 20: Wallet.CloseWallet:output_type -> Success
	1,  // 21: Wallet.GetBalance:output_type -> Balance
	9,  // 22: Wallet.GetValidators:output_type -> ValidatorsRegistry
	5,  // 23: Wallet.GetAccount:output_type -> KeyPair
	10, // 24: Wallet.SendTransaction:output_type -> Hash
	8,  // 25: Wallet.StartValidator:output_type -> Success
	8,  // 26: Wallet.StartValidatorBulk:output_type -> Success
	8,  // 27: Wallet.ExitValidator:output_type -> Success
	8,  // 28: Wallet.ExitValidatorBulk:output_type -> Success
	15, // [15:29] is the sub-list for method output_type
	1,  // [1:15] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_wallet_proto_init() }
func file_wallet_proto_init() {
	if File_wallet_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTransactionInfo); i {
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
		file_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallets); i {
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
		file_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportWalletData); i {
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
			RawDescriptor: file_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_proto_msgTypes,
	}.Build()
	File_wallet_proto = out.File
	file_wallet_proto_rawDesc = nil
	file_wallet_proto_goTypes = nil
	file_wallet_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletClient interface {
	ListWallets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Wallets, error)
	CreateWallet(ctx context.Context, in *Name, opts ...grpc.CallOption) (*KeyPair, error)
	OpenWallet(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Success, error)
	ImportWallet(ctx context.Context, in *ImportWalletData, opts ...grpc.CallOption) (*KeyPair, error)
	DumpWallet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*KeyPair, error)
	CloseWallet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Success, error)
	GetBalance(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Balance, error)
	GetValidators(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ValidatorsRegistry, error)
	GetAccount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*KeyPair, error)
	SendTransaction(ctx context.Context, in *SendTransactionInfo, opts ...grpc.CallOption) (*Hash, error)
	StartValidator(ctx context.Context, in *KeyPair, opts ...grpc.CallOption) (*Success, error)
	StartValidatorBulk(ctx context.Context, in *KeyPairs, opts ...grpc.CallOption) (*Success, error)
	ExitValidator(ctx context.Context, in *KeyPair, opts ...grpc.CallOption) (*Success, error)
	ExitValidatorBulk(ctx context.Context, in *KeyPairs, opts ...grpc.CallOption) (*Success, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) ListWallets(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Wallets, error) {
	out := new(Wallets)
	err := c.cc.Invoke(ctx, "/Wallet/ListWallets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) CreateWallet(ctx context.Context, in *Name, opts ...grpc.CallOption) (*KeyPair, error) {
	out := new(KeyPair)
	err := c.cc.Invoke(ctx, "/Wallet/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) OpenWallet(ctx context.Context, in *Name, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/Wallet/OpenWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ImportWallet(ctx context.Context, in *ImportWalletData, opts ...grpc.CallOption) (*KeyPair, error) {
	out := new(KeyPair)
	err := c.cc.Invoke(ctx, "/Wallet/ImportWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) DumpWallet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*KeyPair, error) {
	out := new(KeyPair)
	err := c.cc.Invoke(ctx, "/Wallet/DumpWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) CloseWallet(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/Wallet/CloseWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetBalance(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Balance, error) {
	out := new(Balance)
	err := c.cc.Invoke(ctx, "/Wallet/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetValidators(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ValidatorsRegistry, error) {
	out := new(ValidatorsRegistry)
	err := c.cc.Invoke(ctx, "/Wallet/GetValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetAccount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*KeyPair, error) {
	out := new(KeyPair)
	err := c.cc.Invoke(ctx, "/Wallet/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) SendTransaction(ctx context.Context, in *SendTransactionInfo, opts ...grpc.CallOption) (*Hash, error) {
	out := new(Hash)
	err := c.cc.Invoke(ctx, "/Wallet/SendTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) StartValidator(ctx context.Context, in *KeyPair, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/Wallet/StartValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) StartValidatorBulk(ctx context.Context, in *KeyPairs, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/Wallet/StartValidatorBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ExitValidator(ctx context.Context, in *KeyPair, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/Wallet/ExitValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) ExitValidatorBulk(ctx context.Context, in *KeyPairs, opts ...grpc.CallOption) (*Success, error) {
	out := new(Success)
	err := c.cc.Invoke(ctx, "/Wallet/ExitValidatorBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
type WalletServer interface {
	ListWallets(context.Context, *Empty) (*Wallets, error)
	CreateWallet(context.Context, *Name) (*KeyPair, error)
	OpenWallet(context.Context, *Name) (*Success, error)
	ImportWallet(context.Context, *ImportWalletData) (*KeyPair, error)
	DumpWallet(context.Context, *Empty) (*KeyPair, error)
	CloseWallet(context.Context, *Empty) (*Success, error)
	GetBalance(context.Context, *Empty) (*Balance, error)
	GetValidators(context.Context, *Empty) (*ValidatorsRegistry, error)
	GetAccount(context.Context, *Empty) (*KeyPair, error)
	SendTransaction(context.Context, *SendTransactionInfo) (*Hash, error)
	StartValidator(context.Context, *KeyPair) (*Success, error)
	StartValidatorBulk(context.Context, *KeyPairs) (*Success, error)
	ExitValidator(context.Context, *KeyPair) (*Success, error)
	ExitValidatorBulk(context.Context, *KeyPairs) (*Success, error)
}

// UnimplementedWalletServer can be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (*UnimplementedWalletServer) ListWallets(context.Context, *Empty) (*Wallets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWallets not implemented")
}
func (*UnimplementedWalletServer) CreateWallet(context.Context, *Name) (*KeyPair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (*UnimplementedWalletServer) OpenWallet(context.Context, *Name) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenWallet not implemented")
}
func (*UnimplementedWalletServer) ImportWallet(context.Context, *ImportWalletData) (*KeyPair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportWallet not implemented")
}
func (*UnimplementedWalletServer) DumpWallet(context.Context, *Empty) (*KeyPair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DumpWallet not implemented")
}
func (*UnimplementedWalletServer) CloseWallet(context.Context, *Empty) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseWallet not implemented")
}
func (*UnimplementedWalletServer) GetBalance(context.Context, *Empty) (*Balance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (*UnimplementedWalletServer) GetValidators(context.Context, *Empty) (*ValidatorsRegistry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidators not implemented")
}
func (*UnimplementedWalletServer) GetAccount(context.Context, *Empty) (*KeyPair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (*UnimplementedWalletServer) SendTransaction(context.Context, *SendTransactionInfo) (*Hash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTransaction not implemented")
}
func (*UnimplementedWalletServer) StartValidator(context.Context, *KeyPair) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartValidator not implemented")
}
func (*UnimplementedWalletServer) StartValidatorBulk(context.Context, *KeyPairs) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartValidatorBulk not implemented")
}
func (*UnimplementedWalletServer) ExitValidator(context.Context, *KeyPair) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitValidator not implemented")
}
func (*UnimplementedWalletServer) ExitValidatorBulk(context.Context, *KeyPairs) (*Success, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitValidatorBulk not implemented")
}

func RegisterWalletServer(s *grpc.Server, srv WalletServer) {
	s.RegisterService(&_Wallet_serviceDesc, srv)
}

func _Wallet_ListWallets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ListWallets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/ListWallets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ListWallets(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CreateWallet(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_OpenWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).OpenWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/OpenWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).OpenWallet(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ImportWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportWalletData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ImportWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/ImportWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ImportWallet(ctx, req.(*ImportWalletData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_DumpWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).DumpWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/DumpWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).DumpWallet(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_CloseWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).CloseWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/CloseWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).CloseWallet(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetBalance(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/GetValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetValidators(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAccount(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).SendTransaction(ctx, req.(*SendTransactionInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_StartValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).StartValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/StartValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).StartValidator(ctx, req.(*KeyPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_StartValidatorBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPairs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).StartValidatorBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/StartValidatorBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).StartValidatorBulk(ctx, req.(*KeyPairs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ExitValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ExitValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/ExitValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ExitValidator(ctx, req.(*KeyPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_ExitValidatorBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPairs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).ExitValidatorBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wallet/ExitValidatorBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).ExitValidatorBulk(ctx, req.(*KeyPairs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wallet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWallets",
			Handler:    _Wallet_ListWallets_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _Wallet_CreateWallet_Handler,
		},
		{
			MethodName: "OpenWallet",
			Handler:    _Wallet_OpenWallet_Handler,
		},
		{
			MethodName: "ImportWallet",
			Handler:    _Wallet_ImportWallet_Handler,
		},
		{
			MethodName: "DumpWallet",
			Handler:    _Wallet_DumpWallet_Handler,
		},
		{
			MethodName: "CloseWallet",
			Handler:    _Wallet_CloseWallet_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Wallet_GetBalance_Handler,
		},
		{
			MethodName: "GetValidators",
			Handler:    _Wallet_GetValidators_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Wallet_GetAccount_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _Wallet_SendTransaction_Handler,
		},
		{
			MethodName: "StartValidator",
			Handler:    _Wallet_StartValidator_Handler,
		},
		{
			MethodName: "StartValidatorBulk",
			Handler:    _Wallet_StartValidatorBulk_Handler,
		},
		{
			MethodName: "ExitValidator",
			Handler:    _Wallet_ExitValidator_Handler,
		},
		{
			MethodName: "ExitValidatorBulk",
			Handler:    _Wallet_ExitValidatorBulk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}
