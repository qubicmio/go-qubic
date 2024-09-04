// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.23.4
// source: core_service.proto

package qubicpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetEntityInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEntityInfoRequest) Reset() {
	*x = GetEntityInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntityInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntityInfoRequest) ProtoMessage() {}

func (x *GetEntityInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntityInfoRequest.ProtoReflect.Descriptor instead.
func (*GetEntityInfoRequest) Descriptor() ([]byte, []int) {
	return file_core_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetEntityInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTickQuorumVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tick uint32 `protobuf:"varint,1,opt,name=tick,proto3" json:"tick,omitempty"`
}

func (x *GetTickQuorumVoteRequest) Reset() {
	*x = GetTickQuorumVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTickQuorumVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTickQuorumVoteRequest) ProtoMessage() {}

func (x *GetTickQuorumVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTickQuorumVoteRequest.ProtoReflect.Descriptor instead.
func (*GetTickQuorumVoteRequest) Descriptor() ([]byte, []int) {
	return file_core_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetTickQuorumVoteRequest) GetTick() uint32 {
	if x != nil {
		return x.Tick
	}
	return 0
}

type GetTickDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tick uint32 `protobuf:"varint,1,opt,name=tick,proto3" json:"tick,omitempty"`
}

func (x *GetTickDataRequest) Reset() {
	*x = GetTickDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTickDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTickDataRequest) ProtoMessage() {}

func (x *GetTickDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTickDataRequest.ProtoReflect.Descriptor instead.
func (*GetTickDataRequest) Descriptor() ([]byte, []int) {
	return file_core_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTickDataRequest) GetTick() uint32 {
	if x != nil {
		return x.Tick
	}
	return 0
}

type GetTickTransactionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tick uint32 `protobuf:"varint,1,opt,name=tick,proto3" json:"tick,omitempty"`
}

func (x *GetTickTransactionsRequest) Reset() {
	*x = GetTickTransactionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTickTransactionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTickTransactionsRequest) ProtoMessage() {}

func (x *GetTickTransactionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTickTransactionsRequest.ProtoReflect.Descriptor instead.
func (*GetTickTransactionsRequest) Descriptor() ([]byte, []int) {
	return file_core_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTickTransactionsRequest) GetTick() uint32 {
	if x != nil {
		return x.Tick
	}
	return 0
}

type GetTickTransactionsStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tick uint32 `protobuf:"varint,1,opt,name=tick,proto3" json:"tick,omitempty"`
}

func (x *GetTickTransactionsStatusRequest) Reset() {
	*x = GetTickTransactionsStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTickTransactionsStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTickTransactionsStatusRequest) ProtoMessage() {}

func (x *GetTickTransactionsStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTickTransactionsStatusRequest.ProtoReflect.Descriptor instead.
func (*GetTickTransactionsStatusRequest) Descriptor() ([]byte, []int) {
	return file_core_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTickTransactionsStatusRequest) GetTick() uint32 {
	if x != nil {
		return x.Tick
	}
	return 0
}

var File_core_service_proto protoreflect.FileDescriptor

var file_core_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x22, 0x28, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x22, 0x30, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x54, 0x69,
	0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x22, 0x36, 0x0a, 0x20, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x69, 0x63,
	0x6b, 0x32, 0xa2, 0x06, 0x0a, 0x0b, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x57, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67,
	0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x68, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x71, 0x75,
	0x62, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x71, 0x75,
	0x62, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x71,
	0x75, 0x62, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x6f, 0x72,
	0x73, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x6f, 0x72, 0x73,
	0x12, 0x74, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x72, 0x75,
	0x6d, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x71, 0x75, 0x62, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x56, 0x6f, 0x74, 0x65, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x51, 0x75, 0x6f, 0x72,
	0x75, 0x6d, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x60, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63,
	0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x69, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a,
	0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x80, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x24, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x98, 0x01, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x2e, 0x71, 0x75, 0x62, 0x69,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a,
	0x01, 0x2a, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x62, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2d, 0x71, 0x75,
	0x62, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x62, 0x69, 0x63, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_service_proto_rawDescOnce sync.Once
	file_core_service_proto_rawDescData = file_core_service_proto_rawDesc
)

func file_core_service_proto_rawDescGZIP() []byte {
	file_core_service_proto_rawDescOnce.Do(func() {
		file_core_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_service_proto_rawDescData)
	})
	return file_core_service_proto_rawDescData
}

var file_core_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_core_service_proto_goTypes = []interface{}{
	(*GetEntityInfoRequest)(nil),             // 0: qubic.v1.GetEntityInfoRequest
	(*GetTickQuorumVoteRequest)(nil),         // 1: qubic.v1.GetTickQuorumVoteRequest
	(*GetTickDataRequest)(nil),               // 2: qubic.v1.GetTickDataRequest
	(*GetTickTransactionsRequest)(nil),       // 3: qubic.v1.GetTickTransactionsRequest
	(*GetTickTransactionsStatusRequest)(nil), // 4: qubic.v1.GetTickTransactionsStatusRequest
	(*emptypb.Empty)(nil),                    // 5: google.protobuf.Empty
	(*TickInfo)(nil),                         // 6: qubic.v1.TickInfo
	(*EntityInfo)(nil),                       // 7: qubic.v1.EntityInfo
	(*Computors)(nil),                        // 8: qubic.v1.Computors
	(*QuorumVote)(nil),                       // 9: qubic.v1.QuorumVote
	(*TickData)(nil),                         // 10: qubic.v1.TickData
	(*TickTransactions)(nil),                 // 11: qubic.v1.TickTransactions
	(*TickTransactionsStatus)(nil),           // 12: qubic.v1.TickTransactionsStatus
}
var file_core_service_proto_depIdxs = []int32{
	5,  // 0: qubic.v1.CoreService.GetTickInfo:input_type -> google.protobuf.Empty
	0,  // 1: qubic.v1.CoreService.GetEntityInfo:input_type -> qubic.v1.GetEntityInfoRequest
	5,  // 2: qubic.v1.CoreService.GetComputors:input_type -> google.protobuf.Empty
	1,  // 3: qubic.v1.CoreService.GetTickQuorumVote:input_type -> qubic.v1.GetTickQuorumVoteRequest
	2,  // 4: qubic.v1.CoreService.GetTickData:input_type -> qubic.v1.GetTickDataRequest
	3,  // 5: qubic.v1.CoreService.GetTickTransactions:input_type -> qubic.v1.GetTickTransactionsRequest
	4,  // 6: qubic.v1.CoreService.GetTickTransactionsStatus:input_type -> qubic.v1.GetTickTransactionsStatusRequest
	6,  // 7: qubic.v1.CoreService.GetTickInfo:output_type -> qubic.v1.TickInfo
	7,  // 8: qubic.v1.CoreService.GetEntityInfo:output_type -> qubic.v1.EntityInfo
	8,  // 9: qubic.v1.CoreService.GetComputors:output_type -> qubic.v1.Computors
	9,  // 10: qubic.v1.CoreService.GetTickQuorumVote:output_type -> qubic.v1.QuorumVote
	10, // 11: qubic.v1.CoreService.GetTickData:output_type -> qubic.v1.TickData
	11, // 12: qubic.v1.CoreService.GetTickTransactions:output_type -> qubic.v1.TickTransactions
	12, // 13: qubic.v1.CoreService.GetTickTransactionsStatus:output_type -> qubic.v1.TickTransactionsStatus
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_core_service_proto_init() }
func file_core_service_proto_init() {
	if File_core_service_proto != nil {
		return
	}
	file_core_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntityInfoRequest); i {
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
		file_core_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTickQuorumVoteRequest); i {
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
		file_core_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTickDataRequest); i {
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
		file_core_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTickTransactionsRequest); i {
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
		file_core_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTickTransactionsStatusRequest); i {
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
			RawDescriptor: file_core_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_service_proto_goTypes,
		DependencyIndexes: file_core_service_proto_depIdxs,
		MessageInfos:      file_core_service_proto_msgTypes,
	}.Build()
	File_core_service_proto = out.File
	file_core_service_proto_rawDesc = nil
	file_core_service_proto_goTypes = nil
	file_core_service_proto_depIdxs = nil
}
