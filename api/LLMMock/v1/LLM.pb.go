// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.6
// source: LLMMock/v1/LLM.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RootRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RootRequest) Reset() {
	*x = RootRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LLMMock_v1_LLM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootRequest) ProtoMessage() {}

func (x *RootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_LLMMock_v1_LLM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootRequest.ProtoReflect.Descriptor instead.
func (*RootRequest) Descriptor() ([]byte, []int) {
	return file_LLMMock_v1_LLM_proto_rawDescGZIP(), []int{0}
}

type RootReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RootReply) Reset() {
	*x = RootReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LLMMock_v1_LLM_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RootReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootReply) ProtoMessage() {}

func (x *RootReply) ProtoReflect() protoreflect.Message {
	mi := &file_LLMMock_v1_LLM_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootReply.ProtoReflect.Descriptor instead.
func (*RootReply) Descriptor() ([]byte, []int) {
	return file_LLMMock_v1_LLM_proto_rawDescGZIP(), []int{1}
}

func (x *RootReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ModelName string `protobuf:"bytes,3,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LLMMock_v1_LLM_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_LLMMock_v1_LLM_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_LLMMock_v1_LLM_proto_rawDescGZIP(), []int{2}
}

func (x *SendMessageRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *SendMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendMessageRequest) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

type SessionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role    string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`       // user or bot
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"` // 用户消息 or 机器人回复
}

func (x *SessionMessage) Reset() {
	*x = SessionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LLMMock_v1_LLM_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionMessage) ProtoMessage() {}

func (x *SessionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_LLMMock_v1_LLM_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionMessage.ProtoReflect.Descriptor instead.
func (*SessionMessage) Descriptor() ([]byte, []int) {
	return file_LLMMock_v1_LLM_proto_rawDescGZIP(), []int{3}
}

func (x *SessionMessage) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *SessionMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response        string          `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	SessionMessages *SessionMessage `protobuf:"bytes,2,opt,name=session_messages,json=sessionMessages,proto3" json:"session_messages,omitempty"` // user
	BotMessages     *SessionMessage `protobuf:"bytes,3,opt,name=bot_messages,json=botMessages,proto3" json:"bot_messages,omitempty"`             // bot
}

func (x *SendMessageReply) Reset() {
	*x = SendMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LLMMock_v1_LLM_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReply) ProtoMessage() {}

func (x *SendMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_LLMMock_v1_LLM_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReply.ProtoReflect.Descriptor instead.
func (*SendMessageReply) Descriptor() ([]byte, []int) {
	return file_LLMMock_v1_LLM_proto_rawDescGZIP(), []int{4}
}

func (x *SendMessageReply) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *SendMessageReply) GetSessionMessages() *SessionMessage {
	if x != nil {
		return x.SessionMessages
	}
	return nil
}

func (x *SendMessageReply) GetBotMessages() *SessionMessage {
	if x != nil {
		return x.BotMessages
	}
	return nil
}

type GetHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *GetHistoryRequest) Reset() {
	*x = GetHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LLMMock_v1_LLM_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryRequest) ProtoMessage() {}

func (x *GetHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_LLMMock_v1_LLM_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetHistoryRequest) Descriptor() ([]byte, []int) {
	return file_LLMMock_v1_LLM_proto_rawDescGZIP(), []int{5}
}

func (x *GetHistoryRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type GetHistoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionMessages []*SessionMessage `protobuf:"bytes,1,rep,name=session_messages,json=sessionMessages,proto3" json:"session_messages,omitempty"` // user
	BotMessages     []*SessionMessage `protobuf:"bytes,2,rep,name=bot_messages,json=botMessages,proto3" json:"bot_messages,omitempty"`             // bot
}

func (x *GetHistoryReply) Reset() {
	*x = GetHistoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LLMMock_v1_LLM_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryReply) ProtoMessage() {}

func (x *GetHistoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_LLMMock_v1_LLM_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryReply.ProtoReflect.Descriptor instead.
func (*GetHistoryReply) Descriptor() ([]byte, []int) {
	return file_LLMMock_v1_LLM_proto_rawDescGZIP(), []int{6}
}

func (x *GetHistoryReply) GetSessionMessages() []*SessionMessage {
	if x != nil {
		return x.SessionMessages
	}
	return nil
}

func (x *GetHistoryReply) GetBotMessages() []*SessionMessage {
	if x != nil {
		return x.BotMessages
	}
	return nil
}

type DeleteHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *DeleteHistoryRequest) Reset() {
	*x = DeleteHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LLMMock_v1_LLM_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHistoryRequest) ProtoMessage() {}

func (x *DeleteHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_LLMMock_v1_LLM_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHistoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteHistoryRequest) Descriptor() ([]byte, []int) {
	return file_LLMMock_v1_LLM_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteHistoryRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type DeleteHistoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteHistoryReply) Reset() {
	*x = DeleteHistoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LLMMock_v1_LLM_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHistoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHistoryReply) ProtoMessage() {}

func (x *DeleteHistoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_LLMMock_v1_LLM_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHistoryReply.ProtoReflect.Descriptor instead.
func (*DeleteHistoryReply) Descriptor() ([]byte, []int) {
	return file_LLMMock_v1_LLM_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteHistoryReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_LLMMock_v1_LLM_proto protoreflect.FileDescriptor

var file_LLMMock_v1_LLM_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4c, 0x4c, 0x4d, 0x4d, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x4c, 0x4c, 0x4d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x25, 0x0a, 0x09, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6c, 0x0a, 0x12, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xba, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x10, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x62, 0x6f, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x62, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x48, 0x0a, 0x10,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x62, 0x6f, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x62, 0x6f, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x91, 0x03, 0x0a, 0x07, 0x4c, 0x4c, 0x4d, 0x4d, 0x6f, 0x63, 0x6b, 0x12, 0x4b, 0x0a, 0x08, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x09, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x03, 0x12, 0x01, 0x2f, 0x12, 0x63, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x22, 0x05, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x12, 0x60, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x72, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x23, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x42, 0x48, 0x0a, 0x19, 0x64, 0x65, 0x76, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x4c, 0x4d, 0x4d, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31,
	0x42, 0x0e, 0x4c, 0x4c, 0x4d, 0x4d, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x31,
	0x50, 0x01, 0x5a, 0x19, 0x4c, 0x4c, 0x4d, 0x4d, 0x6f, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x4c, 0x4c, 0x4d, 0x4d, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LLMMock_v1_LLM_proto_rawDescOnce sync.Once
	file_LLMMock_v1_LLM_proto_rawDescData = file_LLMMock_v1_LLM_proto_rawDesc
)

func file_LLMMock_v1_LLM_proto_rawDescGZIP() []byte {
	file_LLMMock_v1_LLM_proto_rawDescOnce.Do(func() {
		file_LLMMock_v1_LLM_proto_rawDescData = protoimpl.X.CompressGZIP(file_LLMMock_v1_LLM_proto_rawDescData)
	})
	return file_LLMMock_v1_LLM_proto_rawDescData
}

var file_LLMMock_v1_LLM_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_LLMMock_v1_LLM_proto_goTypes = []any{
	(*RootRequest)(nil),          // 0: helloworld.v1.RootRequest
	(*RootReply)(nil),            // 1: helloworld.v1.RootReply
	(*SendMessageRequest)(nil),   // 2: helloworld.v1.SendMessageRequest
	(*SessionMessage)(nil),       // 3: helloworld.v1.SessionMessage
	(*SendMessageReply)(nil),     // 4: helloworld.v1.SendMessageReply
	(*GetHistoryRequest)(nil),    // 5: helloworld.v1.GetHistoryRequest
	(*GetHistoryReply)(nil),      // 6: helloworld.v1.GetHistoryReply
	(*DeleteHistoryRequest)(nil), // 7: helloworld.v1.DeleteHistoryRequest
	(*DeleteHistoryReply)(nil),   // 8: helloworld.v1.DeleteHistoryReply
}
var file_LLMMock_v1_LLM_proto_depIdxs = []int32{
	3, // 0: helloworld.v1.SendMessageReply.session_messages:type_name -> helloworld.v1.SessionMessage
	3, // 1: helloworld.v1.SendMessageReply.bot_messages:type_name -> helloworld.v1.SessionMessage
	3, // 2: helloworld.v1.GetHistoryReply.session_messages:type_name -> helloworld.v1.SessionMessage
	3, // 3: helloworld.v1.GetHistoryReply.bot_messages:type_name -> helloworld.v1.SessionMessage
	0, // 4: helloworld.v1.LLMMock.SayHello:input_type -> helloworld.v1.RootRequest
	2, // 5: helloworld.v1.LLMMock.SendMessage:input_type -> helloworld.v1.SendMessageRequest
	5, // 6: helloworld.v1.LLMMock.GetHistory:input_type -> helloworld.v1.GetHistoryRequest
	7, // 7: helloworld.v1.LLMMock.DeleteHistory:input_type -> helloworld.v1.DeleteHistoryRequest
	1, // 8: helloworld.v1.LLMMock.SayHello:output_type -> helloworld.v1.RootReply
	4, // 9: helloworld.v1.LLMMock.SendMessage:output_type -> helloworld.v1.SendMessageReply
	6, // 10: helloworld.v1.LLMMock.GetHistory:output_type -> helloworld.v1.GetHistoryReply
	8, // 11: helloworld.v1.LLMMock.DeleteHistory:output_type -> helloworld.v1.DeleteHistoryReply
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_LLMMock_v1_LLM_proto_init() }
func file_LLMMock_v1_LLM_proto_init() {
	if File_LLMMock_v1_LLM_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LLMMock_v1_LLM_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RootRequest); i {
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
		file_LLMMock_v1_LLM_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RootReply); i {
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
		file_LLMMock_v1_LLM_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SendMessageRequest); i {
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
		file_LLMMock_v1_LLM_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SessionMessage); i {
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
		file_LLMMock_v1_LLM_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SendMessageReply); i {
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
		file_LLMMock_v1_LLM_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetHistoryRequest); i {
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
		file_LLMMock_v1_LLM_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetHistoryReply); i {
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
		file_LLMMock_v1_LLM_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteHistoryRequest); i {
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
		file_LLMMock_v1_LLM_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteHistoryReply); i {
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
			RawDescriptor: file_LLMMock_v1_LLM_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_LLMMock_v1_LLM_proto_goTypes,
		DependencyIndexes: file_LLMMock_v1_LLM_proto_depIdxs,
		MessageInfos:      file_LLMMock_v1_LLM_proto_msgTypes,
	}.Build()
	File_LLMMock_v1_LLM_proto = out.File
	file_LLMMock_v1_LLM_proto_rawDesc = nil
	file_LLMMock_v1_LLM_proto_goTypes = nil
	file_LLMMock_v1_LLM_proto_depIdxs = nil
}