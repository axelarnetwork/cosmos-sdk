// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: cosmos/auth/v1beta1/auth.proto

package authv1beta1

import (
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BaseAccount defines a base account type. It contains all the necessary fields
// for basic account functionality. Any custom account type should extend this
// type for additional functionality (e.g. vesting).
type BaseAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address       string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey        *anypb.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	AccountNumber uint64     `protobuf:"varint,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64     `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *BaseAccount) Reset() {
	*x = BaseAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseAccount) ProtoMessage() {}

func (x *BaseAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseAccount.ProtoReflect.Descriptor instead.
func (*BaseAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_auth_v1beta1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *BaseAccount) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BaseAccount) GetPubKey() *anypb.Any {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *BaseAccount) GetAccountNumber() uint64 {
	if x != nil {
		return x.AccountNumber
	}
	return 0
}

func (x *BaseAccount) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

// ModuleAccount defines an account for modules that holds coins on a pool.
type ModuleAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAccount *BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3" json:"base_account,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Permissions []string     `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *ModuleAccount) Reset() {
	*x = ModuleAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleAccount) ProtoMessage() {}

func (x *ModuleAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleAccount.ProtoReflect.Descriptor instead.
func (*ModuleAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_auth_v1beta1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ModuleAccount) GetBaseAccount() *BaseAccount {
	if x != nil {
		return x.BaseAccount
	}
	return nil
}

func (x *ModuleAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModuleAccount) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// Params defines the parameters for the auth module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxMemoCharacters      uint64 `protobuf:"varint,1,opt,name=max_memo_characters,json=maxMemoCharacters,proto3" json:"max_memo_characters,omitempty"`
	TxSigLimit             uint64 `protobuf:"varint,2,opt,name=tx_sig_limit,json=txSigLimit,proto3" json:"tx_sig_limit,omitempty"`
	TxSizeCostPerByte      uint64 `protobuf:"varint,3,opt,name=tx_size_cost_per_byte,json=txSizeCostPerByte,proto3" json:"tx_size_cost_per_byte,omitempty"`
	SigVerifyCostEd25519   uint64 `protobuf:"varint,4,opt,name=sig_verify_cost_ed25519,json=sigVerifyCostEd25519,proto3" json:"sig_verify_cost_ed25519,omitempty"`
	SigVerifyCostSecp256K1 uint64 `protobuf:"varint,5,opt,name=sig_verify_cost_secp256k1,json=sigVerifyCostSecp256k1,proto3" json:"sig_verify_cost_secp256k1,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_auth_v1beta1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_cosmos_auth_v1beta1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *Params) GetMaxMemoCharacters() uint64 {
	if x != nil {
		return x.MaxMemoCharacters
	}
	return 0
}

func (x *Params) GetTxSigLimit() uint64 {
	if x != nil {
		return x.TxSigLimit
	}
	return 0
}

func (x *Params) GetTxSizeCostPerByte() uint64 {
	if x != nil {
		return x.TxSizeCostPerByte
	}
	return 0
}

func (x *Params) GetSigVerifyCostEd25519() uint64 {
	if x != nil {
		return x.SigVerifyCostEd25519
	}
	return 0
}

func (x *Params) GetSigVerifyCostSecp256K1() uint64 {
	if x != nil {
		return x.SigVerifyCostSecp256K1
	}
	return 0
}

var File_cosmos_auth_v1beta1_auth_proto protoreflect.FileDescriptor

var file_cosmos_auth_v1beta1_auth_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x47, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x18, 0xea, 0xde,
	0x1f, 0x14, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x2c, 0x6f, 0x6d, 0x69,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x3a, 0x18, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0xca,
	0xb4, 0x2d, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x22, 0xac, 0x01, 0x0a, 0x0d,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x49, 0x0a,
	0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x04, 0xd0, 0xde, 0x1f, 0x01, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x1a,
	0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0, 0x1f, 0x00, 0xca, 0xb4, 0x2d, 0x0e, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x22, 0xbe, 0x02, 0x0a, 0x06, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x65, 0x6d,
	0x6f, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x78, 0x5f, 0x73, 0x69, 0x67, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x78, 0x53,
	0x69, 0x67, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x15, 0x74, 0x78, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x74, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x6f,
	0x73, 0x74, 0x50, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x12, 0x4f, 0x0a, 0x17, 0x73, 0x69, 0x67,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x65, 0x64, 0x32,
	0x35, 0x35, 0x31, 0x39, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x18, 0xe2, 0xde, 0x1f, 0x14,
	0x53, 0x69, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x45, 0x44, 0x32,
	0x35, 0x35, 0x31, 0x39, 0x52, 0x14, 0x73, 0x69, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x73, 0x74, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x12, 0x55, 0x0a, 0x19, 0x73, 0x69,
	0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x1a, 0xe2,
	0xde, 0x1f, 0x16, 0x53, 0x69, 0x67, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x73, 0x74,
	0x53, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x52, 0x16, 0x73, 0x69, 0x67, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b,
	0x31, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x01, 0x42, 0xd4, 0x01, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x13, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x75, 0x74, 0x68,
	0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x43, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_auth_v1beta1_auth_proto_rawDescOnce sync.Once
	file_cosmos_auth_v1beta1_auth_proto_rawDescData = file_cosmos_auth_v1beta1_auth_proto_rawDesc
)

func file_cosmos_auth_v1beta1_auth_proto_rawDescGZIP() []byte {
	file_cosmos_auth_v1beta1_auth_proto_rawDescOnce.Do(func() {
		file_cosmos_auth_v1beta1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_auth_v1beta1_auth_proto_rawDescData)
	})
	return file_cosmos_auth_v1beta1_auth_proto_rawDescData
}

var file_cosmos_auth_v1beta1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cosmos_auth_v1beta1_auth_proto_goTypes = []interface{}{
	(*BaseAccount)(nil),   // 0: cosmos.auth.v1beta1.BaseAccount
	(*ModuleAccount)(nil), // 1: cosmos.auth.v1beta1.ModuleAccount
	(*Params)(nil),        // 2: cosmos.auth.v1beta1.Params
	(*anypb.Any)(nil),     // 3: google.protobuf.Any
}
var file_cosmos_auth_v1beta1_auth_proto_depIdxs = []int32{
	3, // 0: cosmos.auth.v1beta1.BaseAccount.pub_key:type_name -> google.protobuf.Any
	0, // 1: cosmos.auth.v1beta1.ModuleAccount.base_account:type_name -> cosmos.auth.v1beta1.BaseAccount
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cosmos_auth_v1beta1_auth_proto_init() }
func file_cosmos_auth_v1beta1_auth_proto_init() {
	if File_cosmos_auth_v1beta1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_auth_v1beta1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseAccount); i {
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
		file_cosmos_auth_v1beta1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleAccount); i {
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
		file_cosmos_auth_v1beta1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
			RawDescriptor: file_cosmos_auth_v1beta1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_auth_v1beta1_auth_proto_goTypes,
		DependencyIndexes: file_cosmos_auth_v1beta1_auth_proto_depIdxs,
		MessageInfos:      file_cosmos_auth_v1beta1_auth_proto_msgTypes,
	}.Build()
	File_cosmos_auth_v1beta1_auth_proto = out.File
	file_cosmos_auth_v1beta1_auth_proto_rawDesc = nil
	file_cosmos_auth_v1beta1_auth_proto_goTypes = nil
	file_cosmos_auth_v1beta1_auth_proto_depIdxs = nil
}
