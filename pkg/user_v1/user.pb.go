// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.3
// source: user.proto

package user_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	Role_USER  Role = 0
	Role_ADMIN Role = 1
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "USER",
		1: "ADMIN",
	}
	Role_value = map[string]int32{
		"USER":  0,
		"ADMIN": 1,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type CreateUserIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email           string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password        string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirm string `protobuf:"bytes,4,opt,name=password_confirm,json=passwordConfirm,proto3" json:"password_confirm,omitempty"`
	Role            Role   `protobuf:"varint,5,opt,name=role,proto3,enum=user_v1.Role" json:"role,omitempty"`
}

func (x *CreateUserIn) Reset() {
	*x = CreateUserIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserIn) ProtoMessage() {}

func (x *CreateUserIn) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserIn.ProtoReflect.Descriptor instead.
func (*CreateUserIn) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserIn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserIn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserIn) GetPasswordConfirm() string {
	if x != nil {
		return x.PasswordConfirm
	}
	return ""
}

func (x *CreateUserIn) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_USER
}

type CreateUserOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateUserOut) Reset() {
	*x = CreateUserOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserOut) ProtoMessage() {}

func (x *CreateUserOut) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserOut.ProtoReflect.Descriptor instead.
func (*CreateUserOut) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserOut) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserIn) Reset() {
	*x = GetUserIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserIn) ProtoMessage() {}

func (x *GetUserIn) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserIn.ProtoReflect.Descriptor instead.
func (*GetUserIn) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email     string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role      Role                   `protobuf:"varint,4,opt,name=role,proto3,enum=user_v1.Role" json:"role,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetUserOut) Reset() {
	*x = GetUserOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserOut) ProtoMessage() {}

func (x *GetUserOut) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserOut.ProtoReflect.Descriptor instead.
func (*GetUserOut) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserOut) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserOut) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserOut) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserOut) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_USER
}

func (x *GetUserOut) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetUserOut) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UpdateUserIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Role  Role                    `protobuf:"varint,4,opt,name=role,proto3,enum=user_v1.Role" json:"role,omitempty"`
}

func (x *UpdateUserIn) Reset() {
	*x = UpdateUserIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserIn) ProtoMessage() {}

func (x *UpdateUserIn) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserIn.ProtoReflect.Descriptor instead.
func (*UpdateUserIn) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserIn) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UpdateUserIn) GetEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *UpdateUserIn) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_USER
}

type DeleteUserIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserIn) Reset() {
	*x = DeleteUserIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserIn) ProtoMessage() {}

func (x *DeleteUserIn) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserIn.ProtoReflect.Descriptor instead.
func (*DeleteUserIn) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUserIn) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdf, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4f, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x21, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x2a, 0x1b, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53,
	0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x32,
	0xe3, 0x01, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x1a, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x1a, 0x13,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x75, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6f, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x3b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_proto_goTypes = []interface{}{
	(Role)(0),                      // 0: user_v1.Role
	(*CreateUserIn)(nil),           // 1: user_v1.CreateUserIn
	(*CreateUserOut)(nil),          // 2: user_v1.CreateUserOut
	(*GetUserIn)(nil),              // 3: user_v1.GetUserIn
	(*GetUserOut)(nil),             // 4: user_v1.GetUserOut
	(*UpdateUserIn)(nil),           // 5: user_v1.UpdateUserIn
	(*DeleteUserIn)(nil),           // 6: user_v1.DeleteUserIn
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil), // 8: google.protobuf.StringValue
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: user_v1.CreateUserIn.role:type_name -> user_v1.Role
	0,  // 1: user_v1.GetUserOut.role:type_name -> user_v1.Role
	7,  // 2: user_v1.GetUserOut.created_at:type_name -> google.protobuf.Timestamp
	7,  // 3: user_v1.GetUserOut.updated_at:type_name -> google.protobuf.Timestamp
	8,  // 4: user_v1.UpdateUserIn.name:type_name -> google.protobuf.StringValue
	8,  // 5: user_v1.UpdateUserIn.email:type_name -> google.protobuf.StringValue
	0,  // 6: user_v1.UpdateUserIn.role:type_name -> user_v1.Role
	1,  // 7: user_v1.UserV1.Create:input_type -> user_v1.CreateUserIn
	3,  // 8: user_v1.UserV1.Get:input_type -> user_v1.GetUserIn
	5,  // 9: user_v1.UserV1.Update:input_type -> user_v1.UpdateUserIn
	6,  // 10: user_v1.UserV1.Delete:input_type -> user_v1.DeleteUserIn
	2,  // 11: user_v1.UserV1.Create:output_type -> user_v1.CreateUserOut
	4,  // 12: user_v1.UserV1.Get:output_type -> user_v1.GetUserOut
	9,  // 13: user_v1.UserV1.Update:output_type -> google.protobuf.Empty
	9,  // 14: user_v1.UserV1.Delete:output_type -> google.protobuf.Empty
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserIn); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserOut); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserIn); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserOut); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserIn); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserIn); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
