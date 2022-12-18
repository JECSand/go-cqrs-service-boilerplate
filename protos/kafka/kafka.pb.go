// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: kafka.proto

package kafkaMessages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Email     string                 `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Username  string                 `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Password  string                 `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	Root      bool                   `protobuf:"varint,5,opt,name=Root,proto3" json:"Root,omitempty"`
	Active    bool                   `protobuf:"varint,6,opt,name=Active,proto3" json:"Active,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetRoot() bool {
	if x != nil {
		return x.Root
	}
	return false
}

func (x *User) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UserCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *UserCreate) Reset() {
	*x = UserCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreate) ProtoMessage() {}

func (x *UserCreate) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreate.ProtoReflect.Descriptor instead.
func (*UserCreate) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{1}
}

func (x *UserCreate) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UserCreate) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserCreate) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserCreate) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserCreated) Reset() {
	*x = UserCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreated) ProtoMessage() {}

func (x *UserCreated) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreated.ProtoReflect.Descriptor instead.
func (*UserCreated) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{2}
}

func (x *UserCreated) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *UserUpdate) Reset() {
	*x = UserUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdate) ProtoMessage() {}

func (x *UserUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdate.ProtoReflect.Descriptor instead.
func (*UserUpdate) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{3}
}

func (x *UserUpdate) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UserUpdate) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserUpdate) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserUpdated) Reset() {
	*x = UserUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdated) ProtoMessage() {}

func (x *UserUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdated.ProtoReflect.Descriptor instead.
func (*UserUpdated) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{4}
}

func (x *UserUpdated) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UserDelete) Reset() {
	*x = UserDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDelete) ProtoMessage() {}

func (x *UserDelete) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDelete.ProtoReflect.Descriptor instead.
func (*UserDelete) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{5}
}

func (x *UserDelete) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type UserDeleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UserDeleted) Reset() {
	*x = UserDeleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDeleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDeleted) ProtoMessage() {}

func (x *UserDeleted) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDeleted.ProtoReflect.Descriptor instead.
func (*UserDeleted) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{6}
}

func (x *UserDeleted) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type PasswordUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CurrentPassword string `protobuf:"bytes,2,opt,name=CurrentPassword,proto3" json:"CurrentPassword,omitempty"`
	NewPassword     string `protobuf:"bytes,3,opt,name=NewPassword,proto3" json:"NewPassword,omitempty"`
}

func (x *PasswordUpdate) Reset() {
	*x = PasswordUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordUpdate) ProtoMessage() {}

func (x *PasswordUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordUpdate.ProtoReflect.Descriptor instead.
func (*PasswordUpdate) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{7}
}

func (x *PasswordUpdate) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *PasswordUpdate) GetCurrentPassword() string {
	if x != nil {
		return x.CurrentPassword
	}
	return ""
}

func (x *PasswordUpdate) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type PasswordUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64 `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *PasswordUpdated) Reset() {
	*x = PasswordUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordUpdated) ProtoMessage() {}

func (x *PasswordUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordUpdated.ProtoReflect.Descriptor instead.
func (*PasswordUpdated) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{8}
}

func (x *PasswordUpdated) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_kafka_proto protoreflect.FileDescriptor

var file_kafka_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x02,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x6a, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x36, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x27, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x4e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x1c, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x1d,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x6c, 0x0a,
	0x0e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x28, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x3b, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_kafka_proto_rawDescOnce sync.Once
	file_kafka_proto_rawDescData = file_kafka_proto_rawDesc
)

func file_kafka_proto_rawDescGZIP() []byte {
	file_kafka_proto_rawDescOnce.Do(func() {
		file_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_kafka_proto_rawDescData)
	})
	return file_kafka_proto_rawDescData
}

var file_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_kafka_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: kafkaMessages.User
	(*UserCreate)(nil),            // 1: kafkaMessages.UserCreate
	(*UserCreated)(nil),           // 2: kafkaMessages.UserCreated
	(*UserUpdate)(nil),            // 3: kafkaMessages.UserUpdate
	(*UserUpdated)(nil),           // 4: kafkaMessages.UserUpdated
	(*UserDelete)(nil),            // 5: kafkaMessages.UserDelete
	(*UserDeleted)(nil),           // 6: kafkaMessages.UserDeleted
	(*PasswordUpdate)(nil),        // 7: kafkaMessages.PasswordUpdate
	(*PasswordUpdated)(nil),       // 8: kafkaMessages.PasswordUpdated
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_kafka_proto_depIdxs = []int32{
	9, // 0: kafkaMessages.User.CreatedAt:type_name -> google.protobuf.Timestamp
	9, // 1: kafkaMessages.User.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 2: kafkaMessages.UserCreated.User:type_name -> kafkaMessages.User
	0, // 3: kafkaMessages.UserUpdated.User:type_name -> kafkaMessages.User
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_kafka_proto_init() }
func file_kafka_proto_init() {
	if File_kafka_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kafka_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_kafka_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreate); i {
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
		file_kafka_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreated); i {
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
		file_kafka_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdate); i {
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
		file_kafka_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdated); i {
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
		file_kafka_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDelete); i {
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
		file_kafka_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDeleted); i {
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
		file_kafka_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordUpdate); i {
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
		file_kafka_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordUpdated); i {
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
			RawDescriptor: file_kafka_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kafka_proto_goTypes,
		DependencyIndexes: file_kafka_proto_depIdxs,
		MessageInfos:      file_kafka_proto_msgTypes,
	}.Build()
	File_kafka_proto = out.File
	file_kafka_proto_rawDesc = nil
	file_kafka_proto_goTypes = nil
	file_kafka_proto_depIdxs = nil
}
