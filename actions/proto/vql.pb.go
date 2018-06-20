// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vql.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VQLRequest struct {
	VQL                  string   `protobuf:"bytes,1,opt,name=VQL,proto3" json:"VQL,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VQLRequest) Reset()         { *m = VQLRequest{} }
func (m *VQLRequest) String() string { return proto.CompactTextString(m) }
func (*VQLRequest) ProtoMessage()    {}
func (*VQLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_43607497564d8ce5, []int{0}
}
func (m *VQLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLRequest.Unmarshal(m, b)
}
func (m *VQLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLRequest.Marshal(b, m, deterministic)
}
func (dst *VQLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLRequest.Merge(dst, src)
}
func (m *VQLRequest) XXX_Size() int {
	return xxx_messageInfo_VQLRequest.Size(m)
}
func (m *VQLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VQLRequest proto.InternalMessageInfo

func (m *VQLRequest) GetVQL() string {
	if m != nil {
		return m.VQL
	}
	return ""
}

func (m *VQLRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type VQLCollectorArgs struct {
	Query                []*VQLRequest `protobuf:"bytes,2,rep,name=Query,proto3" json:"Query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VQLCollectorArgs) Reset()         { *m = VQLCollectorArgs{} }
func (m *VQLCollectorArgs) String() string { return proto.CompactTextString(m) }
func (*VQLCollectorArgs) ProtoMessage()    {}
func (*VQLCollectorArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_43607497564d8ce5, []int{1}
}
func (m *VQLCollectorArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLCollectorArgs.Unmarshal(m, b)
}
func (m *VQLCollectorArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLCollectorArgs.Marshal(b, m, deterministic)
}
func (dst *VQLCollectorArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLCollectorArgs.Merge(dst, src)
}
func (m *VQLCollectorArgs) XXX_Size() int {
	return xxx_messageInfo_VQLCollectorArgs.Size(m)
}
func (m *VQLCollectorArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLCollectorArgs.DiscardUnknown(m)
}

var xxx_messageInfo_VQLCollectorArgs proto.InternalMessageInfo

func (m *VQLCollectorArgs) GetQuery() []*VQLRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

type VQLResponse struct {
	Response             string      `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	Columns              []string    `protobuf:"bytes,2,rep,name=Columns,proto3" json:"Columns,omitempty"`
	Query                *VQLRequest `protobuf:"bytes,3,opt,name=Query,proto3" json:"Query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VQLResponse) Reset()         { *m = VQLResponse{} }
func (m *VQLResponse) String() string { return proto.CompactTextString(m) }
func (*VQLResponse) ProtoMessage()    {}
func (*VQLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_43607497564d8ce5, []int{2}
}
func (m *VQLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VQLResponse.Unmarshal(m, b)
}
func (m *VQLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VQLResponse.Marshal(b, m, deterministic)
}
func (dst *VQLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VQLResponse.Merge(dst, src)
}
func (m *VQLResponse) XXX_Size() int {
	return xxx_messageInfo_VQLResponse.Size(m)
}
func (m *VQLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VQLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VQLResponse proto.InternalMessageInfo

func (m *VQLResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func (m *VQLResponse) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *VQLResponse) GetQuery() *VQLRequest {
	if m != nil {
		return m.Query
	}
	return nil
}

// FIXME: We replicate a small subset of GRR's elaborate knowledgebase
// protos here because the GUI API plugins use this to construct the
// GRR APIs. When we re-implement the API plugins, refactor this into
// a more sane structure.
type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_43607497564d8ce5, []int{3}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Knowledgebase struct {
	Users                []*User  `protobuf:"bytes,32,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Knowledgebase) Reset()         { *m = Knowledgebase{} }
func (m *Knowledgebase) String() string { return proto.CompactTextString(m) }
func (*Knowledgebase) ProtoMessage()    {}
func (*Knowledgebase) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_43607497564d8ce5, []int{4}
}
func (m *Knowledgebase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Knowledgebase.Unmarshal(m, b)
}
func (m *Knowledgebase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Knowledgebase.Marshal(b, m, deterministic)
}
func (dst *Knowledgebase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Knowledgebase.Merge(dst, src)
}
func (m *Knowledgebase) XXX_Size() int {
	return xxx_messageInfo_Knowledgebase.Size(m)
}
func (m *Knowledgebase) XXX_DiscardUnknown() {
	xxx_messageInfo_Knowledgebase.DiscardUnknown(m)
}

var xxx_messageInfo_Knowledgebase proto.InternalMessageInfo

func (m *Knowledgebase) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type ClientInfo struct {
	Info []*VQLResponse `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	// GRR Keeps these separated so they can be searched on.
	LastTimestamp        uint64         `protobuf:"varint,2,opt,name=last_timestamp,json=lastTimestamp,proto3" json:"last_timestamp,omitempty"`
	Hostname             string         `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Fqdn                 string         `protobuf:"bytes,4,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	System               string         `protobuf:"bytes,5,opt,name=system,proto3" json:"system,omitempty"`
	Release              string         `protobuf:"bytes,6,opt,name=release,proto3" json:"release,omitempty"`
	Architecture         string         `protobuf:"bytes,7,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Usernames            string         `protobuf:"bytes,8,opt,name=usernames,proto3" json:"usernames,omitempty"`
	MacAddress           string         `protobuf:"bytes,9,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	IpAddress            string         `protobuf:"bytes,10,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Ping                 uint64         `protobuf:"varint,11,opt,name=ping,proto3" json:"ping,omitempty"`
	ClientVersion        string         `protobuf:"bytes,12,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	ClientName           string         `protobuf:"bytes,13,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	Knowledgebase        *Knowledgebase `protobuf:"bytes,14,opt,name=knowledgebase,proto3" json:"knowledgebase,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_vql_43607497564d8ce5, []int{5}
}
func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (dst *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(dst, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetInfo() []*VQLResponse {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ClientInfo) GetLastTimestamp() uint64 {
	if m != nil {
		return m.LastTimestamp
	}
	return 0
}

func (m *ClientInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ClientInfo) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *ClientInfo) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *ClientInfo) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *ClientInfo) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func (m *ClientInfo) GetUsernames() string {
	if m != nil {
		return m.Usernames
	}
	return ""
}

func (m *ClientInfo) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *ClientInfo) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *ClientInfo) GetPing() uint64 {
	if m != nil {
		return m.Ping
	}
	return 0
}

func (m *ClientInfo) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *ClientInfo) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *ClientInfo) GetKnowledgebase() *Knowledgebase {
	if m != nil {
		return m.Knowledgebase
	}
	return nil
}

func init() {
	proto.RegisterType((*VQLRequest)(nil), "proto.VQLRequest")
	proto.RegisterType((*VQLCollectorArgs)(nil), "proto.VQLCollectorArgs")
	proto.RegisterType((*VQLResponse)(nil), "proto.VQLResponse")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*Knowledgebase)(nil), "proto.Knowledgebase")
	proto.RegisterType((*ClientInfo)(nil), "proto.ClientInfo")
}

func init() { proto.RegisterFile("vql.proto", fileDescriptor_vql_43607497564d8ce5) }

var fileDescriptor_vql_43607497564d8ce5 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x1a, 0xa7, 0x6d, 0x26, 0x4d, 0x55, 0x56, 0xa8, 0x5a, 0x55, 0x7c, 0x2c, 0x96, 0x80,
	0x22, 0x24, 0x47, 0x94, 0x5b, 0x6f, 0x6d, 0x4f, 0x40, 0x54, 0x14, 0x53, 0x72, 0xe0, 0x52, 0x6d,
	0xed, 0x69, 0x62, 0x61, 0xef, 0x3a, 0xbb, 0xeb, 0xa6, 0xfd, 0x0f, 0xfc, 0x2c, 0xfe, 0x03, 0x77,
	0xf8, 0x1b, 0x1c, 0xd0, 0xce, 0xc6, 0xfd, 0x90, 0x38, 0x79, 0xe6, 0xbd, 0xd9, 0xe7, 0xb7, 0x6f,
	0x16, 0xfa, 0x57, 0x8b, 0x32, 0xa9, 0x8d, 0x76, 0x9a, 0xf5, 0xe8, 0xb3, 0x77, 0xb8, 0x5c, 0x2e,
	0x93, 0x2b, 0x2c, 0x75, 0x56, 0xe4, 0x78, 0x9d, 0x64, 0xba, 0x1a, 0xcd, 0x74, 0x29, 0xd5, 0x6c,
	0x14, 0x40, 0x23, 0x6b, 0xa7, 0xcd, 0x88, 0x86, 0x47, 0x16, 0x2b, 0xa9, 0x5c, 0x91, 0x05, 0x89,
	0xf8, 0x00, 0x60, 0x3a, 0x19, 0xa7, 0xb8, 0x68, 0xd0, 0x3a, 0xb6, 0x03, 0xdd, 0xe9, 0x64, 0xcc,
	0x3b, 0xa2, 0xb3, 0xdf, 0x4f, 0x7d, 0xc9, 0x18, 0x44, 0xa7, 0xb2, 0x42, 0xbe, 0x46, 0x10, 0xd5,
	0xb1, 0x82, 0x9d, 0xe9, 0x64, 0x7c, 0xa2, 0xcb, 0x12, 0x33, 0xa7, 0xcd, 0x91, 0x99, 0x59, 0xf6,
	0x0d, 0x7a, 0x93, 0x06, 0xcd, 0x0d, 0x5f, 0x13, 0xdd, 0xfd, 0xc1, 0xc1, 0xa3, 0x20, 0x9f, 0xdc,
	0x69, 0x1f, 0xbf, 0xfb, 0xfd, 0xf7, 0xcf, 0xcf, 0xce, 0x5b, 0xf6, 0xe6, 0x6c, 0x8e, 0x62, 0x3a,
	0x19, 0x8b, 0x45, 0x83, 0xa6, 0x40, 0x2b, 0x9c, 0x16, 0x78, 0x8d, 0x59, 0xe3, 0x50, 0x68, 0x25,
	0xdc, 0x1c, 0x45, 0x56, 0x16, 0xa8, 0x5c, 0x92, 0x06, 0xc9, 0xf8, 0x47, 0x07, 0x06, 0x24, 0x64,
	0x6b, 0xad, 0x2c, 0xb2, 0x43, 0xd8, 0x6c, 0xeb, 0x60, 0xf5, 0xf8, 0x19, 0x69, 0x73, 0xb6, 0xfb,
	0xf1, 0xcb, 0xe7, 0x53, 0x81, 0x2a, 0xd3, 0x39, 0xe6, 0xc2, 0xac, 0x86, 0x92, 0xf4, 0x76, 0x9e,
	0x71, 0xd8, 0x38, 0xd1, 0x65, 0x53, 0x29, 0x4b, 0x4e, 0xfb, 0x69, 0xdb, 0xb2, 0xd7, 0xed, 0x0d,
	0xba, 0xa2, 0xf3, 0xdf, 0x1b, 0xb4, 0x76, 0x62, 0x88, 0xbe, 0x5a, 0x34, 0x6c, 0x0f, 0x36, 0x1b,
	0x8b, 0x46, 0xf9, 0x78, 0x42, 0x62, 0xb7, 0x7d, 0x7c, 0x00, 0xc3, 0x4f, 0x4a, 0x2f, 0x4b, 0xcc,
	0x67, 0x78, 0x21, 0x2d, 0xb2, 0x17, 0xd0, 0xf3, 0xa4, 0xe5, 0x82, 0xf2, 0x19, 0xac, 0xd4, 0xbd,
	0x50, 0x1a, 0x98, 0xf8, 0x57, 0x17, 0xe0, 0x84, 0x6e, 0xfe, 0x41, 0x5d, 0x6a, 0xf6, 0x0a, 0xa2,
	0x42, 0x5d, 0x6a, 0xde, 0xa1, 0x03, 0xec, 0xbe, 0x9d, 0x70, 0x97, 0x94, 0x78, 0xf6, 0x12, 0xb6,
	0x4b, 0x69, 0xdd, 0xb9, 0x2b, 0x2a, 0xb4, 0x4e, 0x56, 0x35, 0xed, 0x2a, 0x4a, 0x87, 0x1e, 0x3d,
	0x6b, 0x41, 0xef, 0x76, 0xae, 0xad, 0x23, 0xb7, 0xdd, 0xe0, 0xb6, 0xed, 0xfd, 0x92, 0x2f, 0x17,
	0xb9, 0xe2, 0x51, 0x58, 0xb2, 0xaf, 0xd9, 0x2e, 0xac, 0xdb, 0x1b, 0xeb, 0xb0, 0xe2, 0x3d, 0x42,
	0x57, 0x9d, 0x0f, 0xd0, 0x60, 0x89, 0xd2, 0x22, 0x5f, 0x27, 0xa2, 0x6d, 0x59, 0x0c, 0x5b, 0xd2,
	0x64, 0xf3, 0xc2, 0x61, 0xe6, 0x1a, 0x83, 0x7c, 0x83, 0xe8, 0x07, 0x18, 0x7b, 0x02, 0xfd, 0x36,
	0x23, 0xcb, 0x37, 0x69, 0xe0, 0x0e, 0x60, 0xcf, 0x61, 0x50, 0xc9, 0xec, 0x5c, 0xe6, 0xb9, 0x41,
	0x6b, 0x79, 0x9f, 0x78, 0xa8, 0x64, 0x76, 0x14, 0x10, 0xf6, 0x14, 0xa0, 0xa8, 0x6f, 0x79, 0x08,
	0xe7, 0x8b, 0xba, 0xa5, 0x19, 0x44, 0x75, 0xa1, 0x66, 0x7c, 0x40, 0x01, 0x50, 0xed, 0xe3, 0x09,
	0xcf, 0xe9, 0xfc, 0x0a, 0x8d, 0x2d, 0xb4, 0xe2, 0x5b, 0x74, 0x6c, 0x18, 0xd0, 0x69, 0x00, 0xfd,
	0xaf, 0x57, 0x63, 0x94, 0xd0, 0x30, 0xfc, 0x3a, 0x40, 0xfe, 0xd1, 0xb3, 0x43, 0x18, 0x7e, 0xbf,
	0xbf, 0x51, 0xbe, 0x4d, 0xcf, 0xe4, 0xf1, 0x6a, 0x2f, 0x0f, 0xb6, 0x9d, 0x3e, 0x1c, 0xbd, 0x58,
	0xa7, 0x99, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x9d, 0x99, 0xd9, 0xbb, 0x03, 0x00,
	0x00,
}
