// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: api/microservice.proto

package pkg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_microservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_microservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_api_microservice_proto_rawDescGZIP(), []int{0}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_microservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_microservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_api_microservice_proto_rawDescGZIP(), []int{1}
}

func (x *SignInResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_microservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_microservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_api_microservice_proto_rawDescGZIP(), []int{2}
}

func (x *SignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_microservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_microservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_api_microservice_proto_rawDescGZIP(), []int{3}
}

func (x *SignUpResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_microservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_microservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_api_microservice_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32              `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	PerPage int32              `protobuf:"varint,2,opt,name=PerPage,proto3" json:"PerPage,omitempty"`
	Total   int32              `protobuf:"varint,3,opt,name=Total,proto3" json:"Total,omitempty"`
	Users   []*GetUserResponse `protobuf:"bytes,4,rep,name=Users,proto3" json:"Users,omitempty"`
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_microservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_microservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_microservice_proto_rawDescGZIP(), []int{5}
}

func (x *GetUsersResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetUsersResponse) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *GetUsersResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetUsersResponse) GetUsers() []*GetUserResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Email                  string                 `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password               string                 `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	AccessToken            string                 `protobuf:"bytes,4,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken           string                 `protobuf:"bytes,5,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	LastResetSentAt        string                 `protobuf:"bytes,6,opt,name=LastResetSentAt,proto3" json:"LastResetSentAt,omitempty"`
	LastVerificationSentAt string                 `protobuf:"bytes,7,opt,name=LastVerificationSentAt,proto3" json:"LastVerificationSentAt,omitempty"`
	Verified               bool                   `protobuf:"varint,8,opt,name=Verified,proto3" json:"Verified,omitempty"`
	CreatedAt              *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt              *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_microservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_microservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_api_microservice_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetUserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *GetUserResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GetUserResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *GetUserResponse) GetLastResetSentAt() string {
	if x != nil {
		return x.LastResetSentAt
	}
	return ""
}

func (x *GetUserResponse) GetLastVerificationSentAt() string {
	if x != nil {
		return x.LastVerificationSentAt
	}
	return ""
}

func (x *GetUserResponse) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *GetUserResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetUserResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_api_microservice_proto protoreflect.FileDescriptor

var file_api_microservice_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a,
	0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x8b, 0x03, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x53, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4c,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x36,
	0x0a, 0x16, 0x4c, 0x61, 0x73, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16,
	0x4c, 0x61, 0x73, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x9c, 0x02, 0x0a, 0x0c, 0x4d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x12, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06,
	0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1b, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x70, 0x6b, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_microservice_proto_rawDescOnce sync.Once
	file_api_microservice_proto_rawDescData = file_api_microservice_proto_rawDesc
)

func file_api_microservice_proto_rawDescGZIP() []byte {
	file_api_microservice_proto_rawDescOnce.Do(func() {
		file_api_microservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_microservice_proto_rawDescData)
	})
	return file_api_microservice_proto_rawDescData
}

var file_api_microservice_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_microservice_proto_goTypes = []interface{}{
	(*SignInRequest)(nil),         // 0: microservice.SignInRequest
	(*SignInResponse)(nil),        // 1: microservice.SignInResponse
	(*SignUpRequest)(nil),         // 2: microservice.SignUpRequest
	(*SignUpResponse)(nil),        // 3: microservice.SignUpResponse
	(*LogoutResponse)(nil),        // 4: microservice.LogoutResponse
	(*GetUsersResponse)(nil),      // 5: microservice.GetUsersResponse
	(*GetUserResponse)(nil),       // 6: microservice.GetUserResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_api_microservice_proto_depIdxs = []int32{
	6, // 0: microservice.GetUsersResponse.Users:type_name -> microservice.GetUserResponse
	7, // 1: microservice.GetUserResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	7, // 2: microservice.GetUserResponse.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 3: microservice.Microservice.SignIn:input_type -> microservice.SignInRequest
	2, // 4: microservice.Microservice.SignUp:input_type -> microservice.SignUpRequest
	8, // 5: microservice.Microservice.Logout:input_type -> google.protobuf.Empty
	8, // 6: microservice.Microservice.GetUsers:input_type -> google.protobuf.Empty
	1, // 7: microservice.Microservice.SignIn:output_type -> microservice.SignInResponse
	3, // 8: microservice.Microservice.SignUp:output_type -> microservice.SignUpResponse
	4, // 9: microservice.Microservice.Logout:output_type -> microservice.LogoutResponse
	5, // 10: microservice.Microservice.GetUsers:output_type -> microservice.GetUsersResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_microservice_proto_init() }
func file_api_microservice_proto_init() {
	if File_api_microservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_microservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_api_microservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
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
		file_api_microservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
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
		file_api_microservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
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
		file_api_microservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResponse); i {
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
		file_api_microservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResponse); i {
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
		file_api_microservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
			RawDescriptor: file_api_microservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_microservice_proto_goTypes,
		DependencyIndexes: file_api_microservice_proto_depIdxs,
		MessageInfos:      file_api_microservice_proto_msgTypes,
	}.Build()
	File_api_microservice_proto = out.File
	file_api_microservice_proto_rawDesc = nil
	file_api_microservice_proto_goTypes = nil
	file_api_microservice_proto_depIdxs = nil
}
