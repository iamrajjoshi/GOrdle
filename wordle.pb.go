// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: wordle.proto

package main

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

type Color int32

const (
	Color_GRAY   Color = 0
	Color_YELLOW Color = 1
	Color_GREEN  Color = 2
)

// Enum value maps for Color.
var (
	Color_name = map[int32]string{
		0: "GRAY",
		1: "YELLOW",
		2: "GREEN",
	}
	Color_value = map[string]int32{
		"GRAY":   0,
		"YELLOW": 1,
		"GREEN":  2,
	}
)

func (x Color) Enum() *Color {
	p := new(Color)
	*p = x
	return p
}

func (x Color) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Color) Descriptor() protoreflect.EnumDescriptor {
	return file_wordle_proto_enumTypes[0].Descriptor()
}

func (Color) Type() protoreflect.EnumType {
	return &file_wordle_proto_enumTypes[0]
}

func (x Color) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Color.Descriptor instead.
func (Color) EnumDescriptor() ([]byte, []int) {
	return file_wordle_proto_rawDescGZIP(), []int{0}
}

// Server -> Client
// Assigns a unique id to the client
type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Uid       uint64                 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wordle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wordle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_wordle_proto_rawDescGZIP(), []int{0}
}

func (x *JoinResponse) GetTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeStamp
	}
	return nil
}

func (x *JoinResponse) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

// Client -> Server
// User ready to start game
type ReadyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Uid       uint64                 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ReadyRequest) Reset() {
	*x = ReadyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wordle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyRequest) ProtoMessage() {}

func (x *ReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wordle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyRequest.ProtoReflect.Descriptor instead.
func (*ReadyRequest) Descriptor() ([]byte, []int) {
	return file_wordle_proto_rawDescGZIP(), []int{1}
}

func (x *ReadyRequest) GetTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeStamp
	}
	return nil
}

func (x *ReadyRequest) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

// Server -> Client
// Game is starting
type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Uid       uint64                 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wordle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wordle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_wordle_proto_rawDescGZIP(), []int{2}
}

func (x *StartResponse) GetTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeStamp
	}
	return nil
}

func (x *StartResponse) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

// Client -> Server
// User's guess
type GuessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Uid       uint64                 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Guess     string                 `protobuf:"bytes,3,opt,name=guess,proto3" json:"guess,omitempty"`
}

func (x *GuessRequest) Reset() {
	*x = GuessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wordle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuessRequest) ProtoMessage() {}

func (x *GuessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wordle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuessRequest.ProtoReflect.Descriptor instead.
func (*GuessRequest) Descriptor() ([]byte, []int) {
	return file_wordle_proto_rawDescGZIP(), []int{3}
}

func (x *GuessRequest) GetTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeStamp
	}
	return nil
}

func (x *GuessRequest) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GuessRequest) GetGuess() string {
	if x != nil {
		return x.Guess
	}
	return ""
}

// Server -> Client
// Color response to user's guess
type GuessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Uid       uint64                 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Colors    Color                  `protobuf:"varint,3,opt,name=colors,proto3,enum=main.Color" json:"colors,omitempty"`
}

func (x *GuessResponse) Reset() {
	*x = GuessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wordle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuessResponse) ProtoMessage() {}

func (x *GuessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wordle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuessResponse.ProtoReflect.Descriptor instead.
func (*GuessResponse) Descriptor() ([]byte, []int) {
	return file_wordle_proto_rawDescGZIP(), []int{4}
}

func (x *GuessResponse) GetTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeStamp
	}
	return nil
}

func (x *GuessResponse) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GuessResponse) GetColors() Color {
	if x != nil {
		return x.Colors
	}
	return Color_GRAY
}

// Server -> Client
// Color response to opponent's guess
type OpponentUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Uid       uint64                 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Colors    Color                  `protobuf:"varint,3,opt,name=colors,proto3,enum=main.Color" json:"colors,omitempty"`
}

func (x *OpponentUpdateResponse) Reset() {
	*x = OpponentUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wordle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpponentUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpponentUpdateResponse) ProtoMessage() {}

func (x *OpponentUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wordle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpponentUpdateResponse.ProtoReflect.Descriptor instead.
func (*OpponentUpdateResponse) Descriptor() ([]byte, []int) {
	return file_wordle_proto_rawDescGZIP(), []int{5}
}

func (x *OpponentUpdateResponse) GetTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeStamp
	}
	return nil
}

func (x *OpponentUpdateResponse) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *OpponentUpdateResponse) GetColors() Color {
	if x != nil {
		return x.Colors
	}
	return Color_GRAY
}

// Client -> Server
// Game is over
type FinishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	Uid       uint64                 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Win       bool                   `protobuf:"varint,3,opt,name=win,proto3" json:"win,omitempty"`
}

func (x *FinishResponse) Reset() {
	*x = FinishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wordle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishResponse) ProtoMessage() {}

func (x *FinishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wordle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishResponse.ProtoReflect.Descriptor instead.
func (*FinishResponse) Descriptor() ([]byte, []int) {
	return file_wordle_proto_rawDescGZIP(), []int{6}
}

func (x *FinishResponse) GetTimeStamp() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeStamp
	}
	return nil
}

func (x *FinishResponse) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FinishResponse) GetWin() bool {
	if x != nil {
		return x.Win
	}
	return false
}

var File_wordle_proto protoreflect.FileDescriptor

var file_wordle_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x22, 0x5b, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x5c, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x71, 0x0a,
	0x0c, 0x47, 0x75, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x75,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x75, 0x65, 0x73, 0x73,
	0x22, 0x81, 0x01, 0x0a, 0x0d, 0x47, 0x75, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x23, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x16, 0x4f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x06,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x73, 0x22, 0x6f, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x77,
	0x69, 0x6e, 0x2a, 0x28, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x47,
	0x52, 0x41, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x45, 0x4c, 0x4c, 0x4f, 0x57, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45, 0x45, 0x4e, 0x10, 0x02, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wordle_proto_rawDescOnce sync.Once
	file_wordle_proto_rawDescData = file_wordle_proto_rawDesc
)

func file_wordle_proto_rawDescGZIP() []byte {
	file_wordle_proto_rawDescOnce.Do(func() {
		file_wordle_proto_rawDescData = protoimpl.X.CompressGZIP(file_wordle_proto_rawDescData)
	})
	return file_wordle_proto_rawDescData
}

var file_wordle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wordle_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_wordle_proto_goTypes = []interface{}{
	(Color)(0),                     // 0: main.Color
	(*JoinResponse)(nil),           // 1: main.JoinResponse
	(*ReadyRequest)(nil),           // 2: main.ReadyRequest
	(*StartResponse)(nil),          // 3: main.StartResponse
	(*GuessRequest)(nil),           // 4: main.GuessRequest
	(*GuessResponse)(nil),          // 5: main.GuessResponse
	(*OpponentUpdateResponse)(nil), // 6: main.OpponentUpdateResponse
	(*FinishResponse)(nil),         // 7: main.FinishResponse
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
}
var file_wordle_proto_depIdxs = []int32{
	8, // 0: main.JoinResponse.time_stamp:type_name -> google.protobuf.Timestamp
	8, // 1: main.ReadyRequest.time_stamp:type_name -> google.protobuf.Timestamp
	8, // 2: main.StartResponse.time_stamp:type_name -> google.protobuf.Timestamp
	8, // 3: main.GuessRequest.time_stamp:type_name -> google.protobuf.Timestamp
	8, // 4: main.GuessResponse.time_stamp:type_name -> google.protobuf.Timestamp
	0, // 5: main.GuessResponse.colors:type_name -> main.Color
	8, // 6: main.OpponentUpdateResponse.time_stamp:type_name -> google.protobuf.Timestamp
	0, // 7: main.OpponentUpdateResponse.colors:type_name -> main.Color
	8, // 8: main.FinishResponse.time_stamp:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_wordle_proto_init() }
func file_wordle_proto_init() {
	if File_wordle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wordle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_wordle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyRequest); i {
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
		file_wordle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
		file_wordle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuessRequest); i {
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
		file_wordle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuessResponse); i {
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
		file_wordle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpponentUpdateResponse); i {
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
		file_wordle_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishResponse); i {
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
			RawDescriptor: file_wordle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wordle_proto_goTypes,
		DependencyIndexes: file_wordle_proto_depIdxs,
		EnumInfos:         file_wordle_proto_enumTypes,
		MessageInfos:      file_wordle_proto_msgTypes,
	}.Build()
	File_wordle_proto = out.File
	file_wordle_proto_rawDesc = nil
	file_wordle_proto_goTypes = nil
	file_wordle_proto_depIdxs = nil
}
