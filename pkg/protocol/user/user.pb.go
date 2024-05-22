// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.15.0
// source: user/user.proto

package user

import (
	context "context"
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

type GetDesignateUsersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIDs []string `protobuf:"bytes,1,rep,name=userIDs,proto3" json:"userIDs"`
}

func (x *GetDesignateUsersReq) Reset() {
	*x = GetDesignateUsersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDesignateUsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDesignateUsersReq) ProtoMessage() {}

func (x *GetDesignateUsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDesignateUsersReq.ProtoReflect.Descriptor instead.
func (*GetDesignateUsersReq) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetDesignateUsersReq) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

type GetDesignateUsersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersInfo []*UserInfo `protobuf:"bytes,1,rep,name=usersInfo,proto3" json:"usersInfo"`
}

func (x *GetDesignateUsersResp) Reset() {
	*x = GetDesignateUsersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDesignateUsersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDesignateUsersResp) ProtoMessage() {}

func (x *GetDesignateUsersResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDesignateUsersResp.ProtoReflect.Descriptor instead.
func (*GetDesignateUsersResp) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetDesignateUsersResp) GetUsersInfo() []*UserInfo {
	if x != nil {
		return x.UsersInfo
	}
	return nil
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type UserRegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserInfo `protobuf:"bytes,2,rep,name=users,proto3" json:"users"`
}

func (x *UserRegisterReq) Reset() {
	*x = UserRegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterReq) ProtoMessage() {}

func (x *UserRegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterReq.ProtoReflect.Descriptor instead.
func (*UserRegisterReq) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserRegisterReq) GetUsers() []*UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserRegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserRegisterResp) Reset() {
	*x = UserRegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterResp) ProtoMessage() {}

func (x *UserRegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterResp.ProtoReflect.Descriptor instead.
func (*UserRegisterResp) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{4}
}

type UserLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaCode    string `protobuf:"bytes,1,opt,name=areaCode,proto3" json:"areaCode"`
	PhoneNumber string `protobuf:"bytes,2,opt,name=phoneNumber,proto3" json:"phoneNumber"`
	VerifyCode  string `protobuf:"bytes,3,opt,name=verifyCode,proto3" json:"verifyCode"`
	Account     string `protobuf:"bytes,4,opt,name=account,proto3" json:"account"`
	Password    string `protobuf:"bytes,5,opt,name=password,proto3" json:"password"`
	Platform    int32  `protobuf:"varint,6,opt,name=platform,proto3" json:"platform"`
	DeviceID    string `protobuf:"bytes,7,opt,name=deviceID,proto3" json:"deviceID"`
	Email       string `protobuf:"bytes,8,opt,name=email,proto3" json:"email"`
}

func (x *UserLoginReq) Reset() {
	*x = UserLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginReq) ProtoMessage() {}

func (x *UserLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginReq.ProtoReflect.Descriptor instead.
func (*UserLoginReq) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserLoginReq) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

func (x *UserLoginReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserLoginReq) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

func (x *UserLoginReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserLoginReq) GetPlatform() int32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *UserLoginReq) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *UserLoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserLoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token"`
	UserID string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID"`
}

func (x *UserLoginResp) Reset() {
	*x = UserLoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResp) ProtoMessage() {}

func (x *UserLoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResp.ProtoReflect.Descriptor instead.
func (*UserLoginResp) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserLoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserLoginResp) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetUserTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID"`
}

func (x *GetUserTokenReq) Reset() {
	*x = GetUserTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenReq) ProtoMessage() {}

func (x *GetUserTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenReq.ProtoReflect.Descriptor instead.
func (*GetUserTokenReq) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserTokenReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetUserTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
}

func (x *GetUserTokenResp) Reset() {
	*x = GetUserTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenResp) ProtoMessage() {}

func (x *GetUserTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenResp.ProtoReflect.Descriptor instead.
func (*GetUserTokenResp) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_user_user_proto protoreflect.FileDescriptor

var file_user_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0x30,
	0x0a, 0x14, 0x67, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73,
	0x22, 0x4c, 0x0a, 0x15, 0x67, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3e,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3e,
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x2b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x12,
	0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x22, 0xf0, 0x01, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3d, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0x28, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc0, 0x02, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x5a, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x69, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4b,
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69,
	0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x4b, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69,
	0x6d, 0x73, 0x64, 0x6b, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_user_proto_rawDescOnce sync.Once
	file_user_user_proto_rawDescData = file_user_user_proto_rawDesc
)

func file_user_user_proto_rawDescGZIP() []byte {
	file_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_proto_rawDescData)
	})
	return file_user_user_proto_rawDescData
}

var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_user_proto_goTypes = []interface{}{
	(*GetDesignateUsersReq)(nil),  // 0: openim.user.getDesignateUsersReq
	(*GetDesignateUsersResp)(nil), // 1: openim.user.getDesignateUsersResp
	(*UserInfo)(nil),              // 2: openim.user.UserInfo
	(*UserRegisterReq)(nil),       // 3: openim.user.userRegisterReq
	(*UserRegisterResp)(nil),      // 4: openim.user.userRegisterResp
	(*UserLoginReq)(nil),          // 5: openim.user.userLoginReq
	(*UserLoginResp)(nil),         // 6: openim.user.userLoginResp
	(*GetUserTokenReq)(nil),       // 7: openim.user.getUserTokenReq
	(*GetUserTokenResp)(nil),      // 8: openim.user.getUserTokenResp
}
var file_user_user_proto_depIdxs = []int32{
	2, // 0: openim.user.getDesignateUsersResp.usersInfo:type_name -> openim.user.UserInfo
	2, // 1: openim.user.userRegisterReq.users:type_name -> openim.user.UserInfo
	0, // 2: openim.user.user.getDesignateUsers:input_type -> openim.user.getDesignateUsersReq
	3, // 3: openim.user.user.userRegister:input_type -> openim.user.userRegisterReq
	5, // 4: openim.user.user.userLogin:input_type -> openim.user.userLoginReq
	7, // 5: openim.user.user.getUserToken:input_type -> openim.user.getUserTokenReq
	1, // 6: openim.user.user.getDesignateUsers:output_type -> openim.user.getDesignateUsersResp
	4, // 7: openim.user.user.userRegister:output_type -> openim.user.userRegisterResp
	6, // 8: openim.user.user.userLogin:output_type -> openim.user.userLoginResp
	8, // 9: openim.user.user.getUserToken:output_type -> openim.user.getUserTokenResp
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDesignateUsersReq); i {
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
		file_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDesignateUsersResp); i {
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
		file_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterReq); i {
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
		file_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterResp); i {
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
		file_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginReq); i {
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
		file_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginResp); i {
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
		file_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenReq); i {
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
		file_user_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenResp); i {
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
			RawDescriptor: file_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
		MessageInfos:      file_user_user_proto_msgTypes,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_rawDesc = nil
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	// Get the specified user information full field
	GetDesignateUsers(ctx context.Context, in *GetDesignateUsersReq, opts ...grpc.CallOption) (*GetDesignateUsersResp, error)
	// user registration
	UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
	// user login
	UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
	// get user token
	GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...grpc.CallOption) (*GetUserTokenResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetDesignateUsers(ctx context.Context, in *GetDesignateUsersReq, opts ...grpc.CallOption) (*GetDesignateUsersResp, error) {
	out := new(GetDesignateUsersResp)
	err := c.cc.Invoke(ctx, "/openim.user.user/getDesignateUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	out := new(UserRegisterResp)
	err := c.cc.Invoke(ctx, "/openim.user.user/userRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	out := new(UserLoginResp)
	err := c.cc.Invoke(ctx, "/openim.user.user/userLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserToken(ctx context.Context, in *GetUserTokenReq, opts ...grpc.CallOption) (*GetUserTokenResp, error) {
	out := new(GetUserTokenResp)
	err := c.cc.Invoke(ctx, "/openim.user.user/getUserToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	// Get the specified user information full field
	GetDesignateUsers(context.Context, *GetDesignateUsersReq) (*GetDesignateUsersResp, error)
	// user registration
	UserRegister(context.Context, *UserRegisterReq) (*UserRegisterResp, error)
	// user login
	UserLogin(context.Context, *UserLoginReq) (*UserLoginResp, error)
	// get user token
	GetUserToken(context.Context, *GetUserTokenReq) (*GetUserTokenResp, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetDesignateUsers(context.Context, *GetDesignateUsersReq) (*GetDesignateUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDesignateUsers not implemented")
}
func (*UnimplementedUserServer) UserRegister(context.Context, *UserRegisterReq) (*UserRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (*UnimplementedUserServer) UserLogin(context.Context, *UserLoginReq) (*UserLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedUserServer) GetUserToken(context.Context, *GetUserTokenReq) (*GetUserTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserToken not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetDesignateUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDesignateUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetDesignateUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.user.user/GetDesignateUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetDesignateUsers(ctx, req.(*GetDesignateUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.user.user/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.user.user/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*UserLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.user.user/GetUserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserToken(ctx, req.(*GetUserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openim.user.user",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getDesignateUsers",
			Handler:    _User_GetDesignateUsers_Handler,
		},
		{
			MethodName: "userRegister",
			Handler:    _User_UserRegister_Handler,
		},
		{
			MethodName: "userLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "getUserToken",
			Handler:    _User_GetUserToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
