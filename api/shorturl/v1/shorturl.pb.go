// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: api/shorturl/v1/shorturl.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateShorturlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateShorturlRequest) Reset() {
	*x = CreateShorturlRequest{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShorturlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShorturlRequest) ProtoMessage() {}

func (x *CreateShorturlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShorturlRequest.ProtoReflect.Descriptor instead.
func (*CreateShorturlRequest) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShorturlRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type CreateShorturlReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateShorturlReply) Reset() {
	*x = CreateShorturlReply{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShorturlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShorturlReply) ProtoMessage() {}

func (x *CreateShorturlReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShorturlReply.ProtoReflect.Descriptor instead.
func (*CreateShorturlReply) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{1}
}

type UpdateShorturlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateShorturlRequest) Reset() {
	*x = UpdateShorturlRequest{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateShorturlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShorturlRequest) ProtoMessage() {}

func (x *UpdateShorturlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShorturlRequest.ProtoReflect.Descriptor instead.
func (*UpdateShorturlRequest) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{2}
}

type UpdateShorturlReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateShorturlReply) Reset() {
	*x = UpdateShorturlReply{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateShorturlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShorturlReply) ProtoMessage() {}

func (x *UpdateShorturlReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShorturlReply.ProtoReflect.Descriptor instead.
func (*UpdateShorturlReply) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{3}
}

type DeleteShorturlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteShorturlRequest) Reset() {
	*x = DeleteShorturlRequest{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteShorturlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShorturlRequest) ProtoMessage() {}

func (x *DeleteShorturlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShorturlRequest.ProtoReflect.Descriptor instead.
func (*DeleteShorturlRequest) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{4}
}

type DeleteShorturlReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteShorturlReply) Reset() {
	*x = DeleteShorturlReply{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteShorturlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShorturlReply) ProtoMessage() {}

func (x *DeleteShorturlReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShorturlReply.ProtoReflect.Descriptor instead.
func (*DeleteShorturlReply) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{5}
}

type GetShorturlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShorturlRequest) Reset() {
	*x = GetShorturlRequest{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShorturlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShorturlRequest) ProtoMessage() {}

func (x *GetShorturlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShorturlRequest.ProtoReflect.Descriptor instead.
func (*GetShorturlRequest) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{6}
}

type GetShorturlReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetShorturlReply) Reset() {
	*x = GetShorturlReply{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetShorturlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShorturlReply) ProtoMessage() {}

func (x *GetShorturlReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShorturlReply.ProtoReflect.Descriptor instead.
func (*GetShorturlReply) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{7}
}

type ListShorturlRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListShorturlRequest) Reset() {
	*x = ListShorturlRequest{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListShorturlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShorturlRequest) ProtoMessage() {}

func (x *ListShorturlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShorturlRequest.ProtoReflect.Descriptor instead.
func (*ListShorturlRequest) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{8}
}

type ListShorturlReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListShorturlReply) Reset() {
	*x = ListShorturlReply{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListShorturlReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShorturlReply) ProtoMessage() {}

func (x *ListShorturlReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShorturlReply.ProtoReflect.Descriptor instead.
func (*ListShorturlReply) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{9}
}

type HelloReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{10}
}

func (x *HelloReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloResp) Reset() {
	*x = HelloResp{}
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResp) ProtoMessage() {}

func (x *HelloResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_shorturl_v1_shorturl_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResp.ProtoReflect.Descriptor instead.
func (*HelloResp) Descriptor() ([]byte, []int) {
	return file_api_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{11}
}

func (x *HelloResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_shorturl_v1_shorturl_proto protoreflect.FileDescriptor

var file_api_shorturl_v1_shorturl_proto_rawDesc = string([]byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x29, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x75, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x75, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75,
	0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd0,
	0x04, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x12, 0x77, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x12, 0x26, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x75, 0x72, 0x6c, 0x12, 0x5e, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x5e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x55, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x75, 0x72, 0x6c, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75,
	0x72, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x58, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x12, 0x24, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5a, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x42, 0x32, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72,
	0x6c, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x2d, 0x6a, 0x75,
	0x6d, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_shorturl_v1_shorturl_proto_rawDescOnce sync.Once
	file_api_shorturl_v1_shorturl_proto_rawDescData []byte
)

func file_api_shorturl_v1_shorturl_proto_rawDescGZIP() []byte {
	file_api_shorturl_v1_shorturl_proto_rawDescOnce.Do(func() {
		file_api_shorturl_v1_shorturl_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_shorturl_v1_shorturl_proto_rawDesc), len(file_api_shorturl_v1_shorturl_proto_rawDesc)))
	})
	return file_api_shorturl_v1_shorturl_proto_rawDescData
}

var file_api_shorturl_v1_shorturl_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_shorturl_v1_shorturl_proto_goTypes = []any{
	(*CreateShorturlRequest)(nil), // 0: api.shorturl.v1.CreateShorturlRequest
	(*CreateShorturlReply)(nil),   // 1: api.shorturl.v1.CreateShorturlReply
	(*UpdateShorturlRequest)(nil), // 2: api.shorturl.v1.UpdateShorturlRequest
	(*UpdateShorturlReply)(nil),   // 3: api.shorturl.v1.UpdateShorturlReply
	(*DeleteShorturlRequest)(nil), // 4: api.shorturl.v1.DeleteShorturlRequest
	(*DeleteShorturlReply)(nil),   // 5: api.shorturl.v1.DeleteShorturlReply
	(*GetShorturlRequest)(nil),    // 6: api.shorturl.v1.GetShorturlRequest
	(*GetShorturlReply)(nil),      // 7: api.shorturl.v1.GetShorturlReply
	(*ListShorturlRequest)(nil),   // 8: api.shorturl.v1.ListShorturlRequest
	(*ListShorturlReply)(nil),     // 9: api.shorturl.v1.ListShorturlReply
	(*HelloReq)(nil),              // 10: api.shorturl.v1.HelloReq
	(*HelloResp)(nil),             // 11: api.shorturl.v1.HelloResp
}
var file_api_shorturl_v1_shorturl_proto_depIdxs = []int32{
	0,  // 0: api.shorturl.v1.Shorturl.CreateShorturl:input_type -> api.shorturl.v1.CreateShorturlRequest
	2,  // 1: api.shorturl.v1.Shorturl.UpdateShorturl:input_type -> api.shorturl.v1.UpdateShorturlRequest
	4,  // 2: api.shorturl.v1.Shorturl.DeleteShorturl:input_type -> api.shorturl.v1.DeleteShorturlRequest
	6,  // 3: api.shorturl.v1.Shorturl.GetShorturl:input_type -> api.shorturl.v1.GetShorturlRequest
	8,  // 4: api.shorturl.v1.Shorturl.ListShorturl:input_type -> api.shorturl.v1.ListShorturlRequest
	10, // 5: api.shorturl.v1.Shorturl.Hello:input_type -> api.shorturl.v1.HelloReq
	1,  // 6: api.shorturl.v1.Shorturl.CreateShorturl:output_type -> api.shorturl.v1.CreateShorturlReply
	3,  // 7: api.shorturl.v1.Shorturl.UpdateShorturl:output_type -> api.shorturl.v1.UpdateShorturlReply
	5,  // 8: api.shorturl.v1.Shorturl.DeleteShorturl:output_type -> api.shorturl.v1.DeleteShorturlReply
	7,  // 9: api.shorturl.v1.Shorturl.GetShorturl:output_type -> api.shorturl.v1.GetShorturlReply
	9,  // 10: api.shorturl.v1.Shorturl.ListShorturl:output_type -> api.shorturl.v1.ListShorturlReply
	11, // 11: api.shorturl.v1.Shorturl.Hello:output_type -> api.shorturl.v1.HelloResp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_shorturl_v1_shorturl_proto_init() }
func file_api_shorturl_v1_shorturl_proto_init() {
	if File_api_shorturl_v1_shorturl_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_shorturl_v1_shorturl_proto_rawDesc), len(file_api_shorturl_v1_shorturl_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_shorturl_v1_shorturl_proto_goTypes,
		DependencyIndexes: file_api_shorturl_v1_shorturl_proto_depIdxs,
		MessageInfos:      file_api_shorturl_v1_shorturl_proto_msgTypes,
	}.Build()
	File_api_shorturl_v1_shorturl_proto = out.File
	file_api_shorturl_v1_shorturl_proto_goTypes = nil
	file_api_shorturl_v1_shorturl_proto_depIdxs = nil
}
