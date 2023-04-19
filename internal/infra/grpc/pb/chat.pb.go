// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/chat.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ChatRequest struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserMessage          string   `protobuf:"bytes,3,opt,name=user_message,json=userMessage,proto3" json:"user_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatRequest) Reset()         { *m = ChatRequest{} }
func (m *ChatRequest) String() string { return proto.CompactTextString(m) }
func (*ChatRequest) ProtoMessage()    {}
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{0}
}

func (m *ChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatRequest.Unmarshal(m, b)
}
func (m *ChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatRequest.Marshal(b, m, deterministic)
}
func (m *ChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatRequest.Merge(m, src)
}
func (m *ChatRequest) XXX_Size() int {
	return xxx_messageInfo_ChatRequest.Size(m)
}
func (m *ChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChatRequest proto.InternalMessageInfo

func (m *ChatRequest) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *ChatRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ChatRequest) GetUserMessage() string {
	if m != nil {
		return m.UserMessage
	}
	return ""
}

type ChatResponse struct {
	ChatId               string   `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatResponse) Reset()         { *m = ChatResponse{} }
func (m *ChatResponse) String() string { return proto.CompactTextString(m) }
func (*ChatResponse) ProtoMessage()    {}
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7e7dde45555b7d, []int{1}
}

func (m *ChatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatResponse.Unmarshal(m, b)
}
func (m *ChatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatResponse.Marshal(b, m, deterministic)
}
func (m *ChatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatResponse.Merge(m, src)
}
func (m *ChatResponse) XXX_Size() int {
	return xxx_messageInfo_ChatResponse.Size(m)
}
func (m *ChatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChatResponse proto.InternalMessageInfo

func (m *ChatResponse) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *ChatResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ChatResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*ChatRequest)(nil), "pb.ChatRequest")
	proto.RegisterType((*ChatResponse)(nil), "pb.ChatResponse")
}

func init() { proto.RegisterFile("proto/chat.proto", fileDescriptor_ed7e7dde45555b7d) }

var fileDescriptor_ed7e7dde45555b7d = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x31, 0x4f, 0xc4, 0x20,
	0x14, 0xc7, 0xed, 0x99, 0xdc, 0xc5, 0x77, 0x97, 0x78, 0x61, 0x50, 0xe2, 0xa4, 0x9d, 0x9c, 0x8a,
	0xf1, 0xbe, 0x41, 0x9d, 0x3a, 0xb8, 0xd4, 0xad, 0x8b, 0x01, 0xfa, 0x6c, 0x49, 0x2c, 0x20, 0xbc,
	0xfa, 0xf9, 0x0d, 0xa0, 0xd1, 0xd9, 0x8d, 0xdf, 0xff, 0x47, 0x78, 0x7f, 0x1e, 0x1c, 0x7d, 0x70,
	0xe4, 0x84, 0x9e, 0x25, 0x35, 0xf9, 0xc8, 0x36, 0x5e, 0xd5, 0x0a, 0xf6, 0x4f, 0xb3, 0xa4, 0x1e,
	0x3f, 0x56, 0x8c, 0xc4, 0xae, 0x61, 0x97, 0x2e, 0xbc, 0x9a, 0x91, 0x57, 0xb7, 0xd5, 0xfd, 0x45,
	0xbf, 0x4d, 0xd8, 0x8d, 0x49, 0xac, 0x11, 0x43, 0x12, 0x9b, 0x22, 0x12, 0x76, 0x23, 0xbb, 0x83,
	0x43, 0x16, 0x0b, 0xc6, 0x28, 0x27, 0xe4, 0xe7, 0xd9, 0xee, 0x53, 0xf6, 0x5c, 0xa2, 0x7a, 0x80,
	0x43, 0x99, 0x11, 0xbd, 0xb3, 0x11, 0xff, 0x31, 0x84, 0xc3, 0x4e, 0x3b, 0x4b, 0x68, 0xe9, 0xfb,
	0xfd, 0x1f, 0x7c, 0x6c, 0x4b, 0xff, 0x17, 0x0c, 0x9f, 0x46, 0x23, 0x3b, 0x01, 0x64, 0xa4, 0x80,
	0x72, 0x61, 0x97, 0x8d, 0x57, 0xcd, 0x9f, 0xef, 0xdd, 0x1c, 0x7f, 0x83, 0xd2, 0xa5, 0x3e, 0x7b,
	0xa8, 0x5a, 0x3e, 0x5c, 0x19, 0x4b, 0x18, 0xac, 0x7c, 0x17, 0xc6, 0xbe, 0x05, 0x29, 0xa6, 0xe0,
	0xb5, 0xf0, 0x4a, 0x6d, 0xf3, 0xa2, 0x4e, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x87, 0xd7,
	0x20, 0x3c, 0x01, 0x00, 0x00,
}
