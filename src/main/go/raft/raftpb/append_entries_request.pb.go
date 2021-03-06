// Code generated by protoc-gen-go. DO NOT EDIT.
// source: append_entries_request.proto

package raftpb

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

type AppendEntriesRequest struct {
	Term                 *uint64     `protobuf:"varint,1,req,name=Term" json:"Term,omitempty"`
	PrevLogIndex         *uint64     `protobuf:"varint,2,req,name=PrevLogIndex" json:"PrevLogIndex,omitempty"`
	PrevLogTerm          *uint64     `protobuf:"varint,3,req,name=PrevLogTerm" json:"PrevLogTerm,omitempty"`
	CommitIndex          *uint64     `protobuf:"varint,4,req,name=CommitIndex" json:"CommitIndex,omitempty"`
	LeaderName           *string     `protobuf:"bytes,5,req,name=LeaderName" json:"LeaderName,omitempty"`
	Entries              []*LogEntry `protobuf:"bytes,6,rep,name=Entries" json:"Entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AppendEntriesRequest) Reset()         { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()    {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92c1c73991de9c02, []int{0}
}

func (m *AppendEntriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntriesRequest.Unmarshal(m, b)
}
func (m *AppendEntriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntriesRequest.Marshal(b, m, deterministic)
}
func (m *AppendEntriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesRequest.Merge(m, src)
}
func (m *AppendEntriesRequest) XXX_Size() int {
	return xxx_messageInfo_AppendEntriesRequest.Size(m)
}
func (m *AppendEntriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesRequest proto.InternalMessageInfo

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if m != nil && m.PrevLogIndex != nil {
		return *m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if m != nil && m.PrevLogTerm != nil {
		return *m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesRequest) GetCommitIndex() uint64 {
	if m != nil && m.CommitIndex != nil {
		return *m.CommitIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeaderName() string {
	if m != nil && m.LeaderName != nil {
		return *m.LeaderName
	}
	return ""
}

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*AppendEntriesRequest)(nil), "raftpb.AppendEntriesRequest")
}

func init() { proto.RegisterFile("append_entries_request.proto", fileDescriptor_92c1c73991de9c02) }

var fileDescriptor_92c1c73991de9c02 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x2c, 0x28, 0x48,
	0xcd, 0x4b, 0x89, 0x4f, 0xcd, 0x2b, 0x29, 0xca, 0x4c, 0x2d, 0x8e, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x4a, 0x4c, 0x2b, 0x29, 0x48,
	0x92, 0xe2, 0xcf, 0xc9, 0x4f, 0x07, 0x2b, 0xa9, 0x84, 0x48, 0x28, 0xcd, 0x66, 0xe4, 0x12, 0x71,
	0x04, 0xeb, 0x74, 0x85, 0x68, 0x0c, 0x82, 0xe8, 0x13, 0xe2, 0xe1, 0x62, 0x09, 0x49, 0x2d, 0xca,
	0x95, 0x60, 0x54, 0x60, 0xd2, 0x60, 0x11, 0x12, 0xe1, 0xe2, 0x09, 0x28, 0x4a, 0x2d, 0xf3, 0xc9,
	0x4f, 0xf7, 0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60, 0x02, 0x8b, 0x0a, 0x73, 0x71, 0x43, 0x45, 0xc1,
	0x4a, 0x99, 0x61, 0x82, 0xce, 0xf9, 0xb9, 0xb9, 0x99, 0x25, 0x10, 0x95, 0x2c, 0x60, 0x41, 0x21,
	0x2e, 0x2e, 0x9f, 0xd4, 0xc4, 0x94, 0xd4, 0x22, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x56, 0x05, 0x26,
	0x0d, 0x4e, 0x21, 0x45, 0x2e, 0x76, 0xa8, 0x9d, 0x12, 0x6c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x02,
	0x7a, 0x10, 0x57, 0xea, 0xf9, 0xe4, 0xa7, 0x83, 0x64, 0x2a, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xa3, 0xb3, 0x46, 0xa9, 0xd5, 0x00, 0x00, 0x00,
}
