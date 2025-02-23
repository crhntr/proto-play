// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: api/playground/v1/playground.proto

package v1

import (
	v1 "github.com/crhntr/proto-play/api/authentication/v1"
	v3 "github.com/crhntr/proto-play/api/authentication/v3"
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

type Username struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Username) Reset() {
	*x = Username{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Username) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Username) ProtoMessage() {}

func (x *Username) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Username.ProtoReflect.Descriptor instead.
func (*Username) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{0}
}

func (x *Username) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PersonName struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	First         string                 `protobuf:"bytes,1,opt,name=first,proto3" json:"first,omitempty"`
	Last          string                 `protobuf:"bytes,2,opt,name=last,proto3" json:"last,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PersonName) Reset() {
	*x = PersonName{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PersonName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonName) ProtoMessage() {}

func (x *PersonName) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonName.ProtoReflect.Descriptor instead.
func (*PersonName) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{1}
}

func (x *PersonName) GetFirst() string {
	if x != nil {
		return x.First
	}
	return ""
}

func (x *PersonName) GetLast() string {
	if x != nil {
		return x.Last
	}
	return ""
}

type Identifier struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Id:
	//
	//	*Identifier_Un
	//	*Identifier_Pn
	Id            isIdentifier_Id `protobuf_oneof:"id"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Identifier) Reset() {
	*x = Identifier{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identifier) ProtoMessage() {}

func (x *Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identifier.ProtoReflect.Descriptor instead.
func (*Identifier) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{2}
}

func (x *Identifier) GetId() isIdentifier_Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Identifier) GetUn() *v1.Username {
	if x != nil {
		if x, ok := x.Id.(*Identifier_Un); ok {
			return x.Un
		}
	}
	return nil
}

func (x *Identifier) GetPn() *v3.PersonName {
	if x != nil {
		if x, ok := x.Id.(*Identifier_Pn); ok {
			return x.Pn
		}
	}
	return nil
}

type isIdentifier_Id interface {
	isIdentifier_Id()
}

type Identifier_Un struct {
	Un *v1.Username `protobuf:"bytes,1,opt,name=un,proto3,oneof"`
}

type Identifier_Pn struct {
	Pn *v3.PersonName `protobuf:"bytes,2,opt,name=pn,proto3,oneof"`
}

func (*Identifier_Un) isIdentifier_Id() {}

func (*Identifier_Pn) isIdentifier_Id() {}

type ListByIDKeyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *Identifier            `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListByIDKeyRequest) Reset() {
	*x = ListByIDKeyRequest{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListByIDKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByIDKeyRequest) ProtoMessage() {}

func (x *ListByIDKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByIDKeyRequest.ProtoReflect.Descriptor instead.
func (*ListByIDKeyRequest) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{3}
}

func (x *ListByIDKeyRequest) GetQuery() *Identifier {
	if x != nil {
		return x.Query
	}
	return nil
}

type ListByIDKeyResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []*Identifier          `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListByIDKeyResponse) Reset() {
	*x = ListByIDKeyResponse{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListByIDKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByIDKeyResponse) ProtoMessage() {}

func (x *ListByIDKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByIDKeyResponse.ProtoReflect.Descriptor instead.
func (*ListByIDKeyResponse) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{4}
}

func (x *ListByIDKeyResponse) GetMessages() []*Identifier {
	if x != nil {
		return x.Messages
	}
	return nil
}

type ExistsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *Identifier            `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExistsRequest) Reset() {
	*x = ExistsRequest{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsRequest) ProtoMessage() {}

func (x *ExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsRequest.ProtoReflect.Descriptor instead.
func (*ExistsRequest) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{5}
}

func (x *ExistsRequest) GetQuery() *Identifier {
	if x != nil {
		return x.Query
	}
	return nil
}

type ExistsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Found         bool                   `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExistsResponse) Reset() {
	*x = ExistsResponse{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsResponse) ProtoMessage() {}

func (x *ExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsResponse.ProtoReflect.Descriptor instead.
func (*ExistsResponse) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{6}
}

func (x *ExistsResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       *Identifier            `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{7}
}

func (x *CreateRequest) GetContent() *Identifier {
	if x != nil {
		return x.Content
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Number        int64                  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{8}
}

func (x *CreateResponse) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{9}
}

type ListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	List          []*Identifier          `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	mi := &file_api_playground_v1_playground_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_playground_v1_playground_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_api_playground_v1_playground_proto_rawDescGZIP(), []int{10}
}

func (x *ListResponse) GetList() []*Identifier {
	if x != nil {
		return x.List
	}
	return nil
}

var File_api_playground_v1_playground_proto protoreflect.FileDescriptor

var file_api_playground_v1_playground_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x72, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x61, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74,
	0x22, 0x7a, 0x0a, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x31,
	0x0a, 0x02, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x02, 0x75,
	0x6e, 0x12, 0x33, 0x0a, 0x02, 0x70, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x48, 0x00, 0x52, 0x02, 0x70, 0x6e, 0x42, 0x04, 0x0a, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x50, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x0d, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22,
	0x26, 0x0a, 0x0e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x48, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x28, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x0d, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xd3, 0x02,
	0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x72, 0x68, 0x6e, 0x74, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70,
	0x6c, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_playground_v1_playground_proto_rawDescOnce sync.Once
	file_api_playground_v1_playground_proto_rawDescData = file_api_playground_v1_playground_proto_rawDesc
)

func file_api_playground_v1_playground_proto_rawDescGZIP() []byte {
	file_api_playground_v1_playground_proto_rawDescOnce.Do(func() {
		file_api_playground_v1_playground_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_playground_v1_playground_proto_rawDescData)
	})
	return file_api_playground_v1_playground_proto_rawDescData
}

var file_api_playground_v1_playground_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_playground_v1_playground_proto_goTypes = []any{
	(*Username)(nil),            // 0: api.playground.v1.Username
	(*PersonName)(nil),          // 1: api.playground.v1.PersonName
	(*Identifier)(nil),          // 2: api.playground.v1.Identifier
	(*ListByIDKeyRequest)(nil),  // 3: api.playground.v1.ListByIDKeyRequest
	(*ListByIDKeyResponse)(nil), // 4: api.playground.v1.ListByIDKeyResponse
	(*ExistsRequest)(nil),       // 5: api.playground.v1.ExistsRequest
	(*ExistsResponse)(nil),      // 6: api.playground.v1.ExistsResponse
	(*CreateRequest)(nil),       // 7: api.playground.v1.CreateRequest
	(*CreateResponse)(nil),      // 8: api.playground.v1.CreateResponse
	(*ListRequest)(nil),         // 9: api.playground.v1.ListRequest
	(*ListResponse)(nil),        // 10: api.playground.v1.ListResponse
	(*v1.Username)(nil),         // 11: api.authentication.v1.Username
	(*v3.PersonName)(nil),       // 12: api.authentication.v3.PersonName
}
var file_api_playground_v1_playground_proto_depIdxs = []int32{
	11, // 0: api.playground.v1.Identifier.un:type_name -> api.authentication.v1.Username
	12, // 1: api.playground.v1.Identifier.pn:type_name -> api.authentication.v3.PersonName
	2,  // 2: api.playground.v1.ListByIDKeyRequest.query:type_name -> api.playground.v1.Identifier
	2,  // 3: api.playground.v1.ListByIDKeyResponse.messages:type_name -> api.playground.v1.Identifier
	2,  // 4: api.playground.v1.ExistsRequest.query:type_name -> api.playground.v1.Identifier
	2,  // 5: api.playground.v1.CreateRequest.content:type_name -> api.playground.v1.Identifier
	2,  // 6: api.playground.v1.ListResponse.list:type_name -> api.playground.v1.Identifier
	3,  // 7: api.playground.v1.StoreService.ListByIDKey:input_type -> api.playground.v1.ListByIDKeyRequest
	5,  // 8: api.playground.v1.StoreService.Exists:input_type -> api.playground.v1.ExistsRequest
	7,  // 9: api.playground.v1.StoreService.Create:input_type -> api.playground.v1.CreateRequest
	9,  // 10: api.playground.v1.StoreService.List:input_type -> api.playground.v1.ListRequest
	4,  // 11: api.playground.v1.StoreService.ListByIDKey:output_type -> api.playground.v1.ListByIDKeyResponse
	6,  // 12: api.playground.v1.StoreService.Exists:output_type -> api.playground.v1.ExistsResponse
	8,  // 13: api.playground.v1.StoreService.Create:output_type -> api.playground.v1.CreateResponse
	10, // 14: api.playground.v1.StoreService.List:output_type -> api.playground.v1.ListResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_playground_v1_playground_proto_init() }
func file_api_playground_v1_playground_proto_init() {
	if File_api_playground_v1_playground_proto != nil {
		return
	}
	file_api_playground_v1_playground_proto_msgTypes[2].OneofWrappers = []any{
		(*Identifier_Un)(nil),
		(*Identifier_Pn)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_playground_v1_playground_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_playground_v1_playground_proto_goTypes,
		DependencyIndexes: file_api_playground_v1_playground_proto_depIdxs,
		MessageInfos:      file_api_playground_v1_playground_proto_msgTypes,
	}.Build()
	File_api_playground_v1_playground_proto = out.File
	file_api_playground_v1_playground_proto_rawDesc = nil
	file_api_playground_v1_playground_proto_goTypes = nil
	file_api_playground_v1_playground_proto_depIdxs = nil
}
