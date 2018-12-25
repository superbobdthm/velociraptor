// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifact.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "www.velocidex.com/golang/velociraptor/proto"
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

type ArtifactParameter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Default              string   `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactParameter) Reset()         { *m = ArtifactParameter{} }
func (m *ArtifactParameter) String() string { return proto.CompactTextString(m) }
func (*ArtifactParameter) ProtoMessage()    {}
func (*ArtifactParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{0}
}

func (m *ArtifactParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactParameter.Unmarshal(m, b)
}
func (m *ArtifactParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactParameter.Marshal(b, m, deterministic)
}
func (m *ArtifactParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactParameter.Merge(m, src)
}
func (m *ArtifactParameter) XXX_Size() int {
	return xxx_messageInfo_ArtifactParameter.Size(m)
}
func (m *ArtifactParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactParameter.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactParameter proto.InternalMessageInfo

func (m *ArtifactParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactParameter) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *ArtifactParameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ArtifactSource struct {
	Precondition         string   `protobuf:"bytes,1,opt,name=precondition,proto3" json:"precondition,omitempty"`
	Queries              []string `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactSource) Reset()         { *m = ArtifactSource{} }
func (m *ArtifactSource) String() string { return proto.CompactTextString(m) }
func (*ArtifactSource) ProtoMessage()    {}
func (*ArtifactSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{1}
}

func (m *ArtifactSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactSource.Unmarshal(m, b)
}
func (m *ArtifactSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactSource.Marshal(b, m, deterministic)
}
func (m *ArtifactSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactSource.Merge(m, src)
}
func (m *ArtifactSource) XXX_Size() int {
	return xxx_messageInfo_ArtifactSource.Size(m)
}
func (m *ArtifactSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactSource.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactSource proto.InternalMessageInfo

func (m *ArtifactSource) GetPrecondition() string {
	if m != nil {
		return m.Precondition
	}
	return ""
}

func (m *ArtifactSource) GetQueries() []string {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Artifact struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string               `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	Raw                  string               `protobuf:"bytes,7,opt,name=raw,proto3" json:"raw,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Reference            string               `protobuf:"bytes,5,opt,name=reference,proto3" json:"reference,omitempty"`
	Parameters           []*ArtifactParameter `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
	Sources              []*ArtifactSource    `protobuf:"bytes,4,rep,name=sources,proto3" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{2}
}

func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artifact) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Artifact) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

func (m *Artifact) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Artifact) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *Artifact) GetParameters() []*ArtifactParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Artifact) GetSources() []*ArtifactSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

type ArtifactDescriptors struct {
	Items                []*Artifact `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ArtifactDescriptors) Reset()         { *m = ArtifactDescriptors{} }
func (m *ArtifactDescriptors) String() string { return proto.CompactTextString(m) }
func (*ArtifactDescriptors) ProtoMessage()    {}
func (*ArtifactDescriptors) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{3}
}

func (m *ArtifactDescriptors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactDescriptors.Unmarshal(m, b)
}
func (m *ArtifactDescriptors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactDescriptors.Marshal(b, m, deterministic)
}
func (m *ArtifactDescriptors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactDescriptors.Merge(m, src)
}
func (m *ArtifactDescriptors) XXX_Size() int {
	return xxx_messageInfo_ArtifactDescriptors.Size(m)
}
func (m *ArtifactDescriptors) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactDescriptors.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactDescriptors proto.InternalMessageInfo

func (m *ArtifactDescriptors) GetItems() []*Artifact {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*ArtifactParameter)(nil), "proto.ArtifactParameter")
	proto.RegisterType((*ArtifactSource)(nil), "proto.ArtifactSource")
	proto.RegisterType((*Artifact)(nil), "proto.Artifact")
	proto.RegisterType((*ArtifactDescriptors)(nil), "proto.ArtifactDescriptors")
}

func init() { proto.RegisterFile("artifact.proto", fileDescriptor_a1932e98ed811590) }

var fileDescriptor_a1932e98ed811590 = []byte{
	// 876 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0x24, 0x35,
	0x10, 0x56, 0x67, 0x36, 0x09, 0xf1, 0xa2, 0x45, 0x18, 0x21, 0xb5, 0x38, 0xa0, 0x22, 0x2b, 0x20,
	0x41, 0x8b, 0x07, 0x85, 0xcb, 0x6a, 0xb4, 0x42, 0x9a, 0x21, 0x48, 0x08, 0x85, 0x4d, 0x32, 0x89,
	0x58, 0xb1, 0x37, 0x8f, 0xbb, 0x7a, 0xda, 0x5a, 0xb7, 0xdd, 0xeb, 0x9f, 0x74, 0x9a, 0x0b, 0xef,
	0x80, 0xc4, 0x09, 0x9e, 0x80, 0x1b, 0xef, 0x80, 0x78, 0x04, 0x24, 0x38, 0x21, 0x78, 0x0d, 0x0e,
	0xc8, 0xee, 0x69, 0x66, 0x66, 0xf7, 0xc0, 0x69, 0xa6, 0xcb, 0xf6, 0xf7, 0x55, 0x7d, 0x55, 0x5f,
	0x91, 0x7b, 0xdc, 0x7a, 0x59, 0x72, 0xe1, 0x59, 0x63, 0x8d, 0x37, 0x74, 0x37, 0xfd, 0xbc, 0x35,
	0x69, 0xdb, 0x96, 0xdd, 0xa0, 0x32, 0x42, 0x16, 0x78, 0xcb, 0x84, 0xa9, 0xc7, 0x4b, 0xa3, 0xb8,
	0x5e, 0x8e, 0xfb, 0xa0, 0xe5, 0x8d, 0x37, 0x76, 0x9c, 0x2e, 0x8f, 0x1d, 0xd6, 0x5c, 0x7b, 0x29,
	0x7a, 0x88, 0xc3, 0x5f, 0x77, 0xc8, 0xeb, 0xd3, 0x15, 0xea, 0x05, 0xb7, 0xbc, 0x46, 0x8f, 0x96,
	0x52, 0x72, 0x47, 0xf3, 0x1a, 0xf3, 0x0c, 0xb2, 0xa3, 0x83, 0x79, 0xfa, 0x4f, 0x73, 0xb2, 0x5f,
	0x60, 0xc9, 0x83, 0xf2, 0xf9, 0x4e, 0x0a, 0x0f, 0x9f, 0x14, 0xc8, 0xdd, 0x02, 0x9d, 0xb0, 0xb2,
	0xf1, 0xd2, 0xe8, 0x7c, 0x94, 0x4e, 0x37, 0x43, 0x93, 0x3f, 0xb3, 0x3f, 0xfe, 0xf9, 0xfb, 0x97,
	0xec, 0xf7, 0x8c, 0xfc, 0x96, 0x0d, 0x6c, 0x0e, 0x6a, 0xde, 0x01, 0x17, 0x02, 0x1b, 0x0f, 0xcd,
	0x40, 0xed, 0xa0, 0xad, 0xa4, 0xa8, 0x80, 0x5b, 0x04, 0x5e, 0x14, 0x58, 0x80, 0x37, 0xe0, 0x2b,
	0x04, 0x27, 0x4c, 0x83, 0xd0, 0x58, 0x69, 0x6c, 0x0c, 0xe1, 0x2d, 0x8a, 0x10, 0xb1, 0x19, 0x3c,
	0x3e, 0xbf, 0xfe, 0x6c, 0x02, 0x5c, 0xa9, 0x4d, 0x94, 0xf8, 0xde, 0x79, 0x2b, 0xf5, 0xd2, 0xc1,
	0x87, 0x20, 0x4b, 0xe8, 0x4c, 0x00, 0x8d, 0x58, 0x80, 0x33, 0x35, 0xfa, 0x4a, 0xea, 0x25, 0xa0,
	0x72, 0x98, 0xb0, 0x9f, 0x07, 0xb4, 0x1d, 0x08, 0xae, 0x21, 0xe8, 0x86, 0x8b, 0x67, 0x80, 0x6c,
	0xc9, 0xa0, 0xb4, 0xa6, 0x86, 0x2f, 0xae, 0xce, 0x1f, 0x43, 0x70, 0xf1, 0x7a, 0xbc, 0x19, 0x3f,
	0x2f, 0xb8, 0x75, 0x78, 0x74, 0x0c, 0x5f, 0x5d, 0x9e, 0x41, 0x19, 0xb4, 0x48, 0x59, 0x1c, 0xfe,
	0xb0, 0x43, 0xee, 0x0d, 0xa5, 0x5d, 0x99, 0x60, 0x05, 0xd2, 0x8a, 0xbc, 0xda, 0x58, 0x14, 0x46,
	0x17, 0x32, 0x09, 0x93, 0xd4, 0x9c, 0x9d, 0xfe, 0x15, 0xa5, 0xf8, 0x84, 0x3e, 0x9a, 0x26, 0x04,
	0xbc, 0x6d, 0x2c, 0x3a, 0x27, 0x8d, 0x8e, 0x65, 0x2d, 0x10, 0xf0, 0x86, 0xab, 0xc0, 0x3d, 0x16,
	0xeb, 0x62, 0x07, 0x7e, 0xe9, 0xc0, 0x25, 0x68, 0x36, 0xdf, 0x42, 0xa6, 0x9e, 0xec, 0xc7, 0x1a,
	0x24, 0xba, 0x7c, 0x07, 0x46, 0x47, 0x07, 0xb3, 0xa7, 0x89, 0xe4, 0x9a, 0xce, 0x2f, 0xfb, 0x30,
	0xf8, 0x8a, 0x7b, 0x68, 0xa5, 0x52, 0x60, 0x83, 0x06, 0xa9, 0xc1, 0xd8, 0x02, 0x2d, 0x83, 0x73,
	0xad, 0x3a, 0x30, 0xc1, 0x37, 0xc1, 0xf7, 0x05, 0xc7, 0x2a, 0x15, 0x77, 0x7e, 0x25, 0x4a, 0x7a,
	0xb2, 0x40, 0x10, 0x46, 0x29, 0x14, 0x1e, 0x0b, 0x36, 0x1f, 0xa8, 0x26, 0xc7, 0xa9, 0xa9, 0xf7,
	0xc9, 0x3b, 0x4f, 0x2a, 0xb4, 0xbd, 0x92, 0xc3, 0x80, 0xc2, 0x12, 0xbd, 0x03, 0xe9, 0x1d, 0x14,
	0xdc, 0x73, 0x76, 0xf8, 0xd3, 0x1e, 0x79, 0x65, 0x50, 0x87, 0xfe, 0x9c, 0x6d, 0x8e, 0xd7, 0xec,
	0xc7, 0x2c, 0x25, 0xfb, 0x7d, 0x46, 0xbf, 0xcb, 0xae, 0x2b, 0x84, 0x78, 0x04, 0xa6, 0xdc, 0xc2,
	0x63, 0x70, 0x55, 0x99, 0xa0, 0x8a, 0x98, 0x4a, 0xd0, 0xf2, 0x79, 0x40, 0xe0, 0xba, 0x48, 0x13,
	0x24, 0x8c, 0xf6, 0x5c, 0x6a, 0x28, 0x8c, 0x77, 0x0c, 0xa6, 0x10, 0x1c, 0x96, 0x41, 0x81, 0x13,
	0x15, 0xd6, 0x08, 0xd2, 0x25, 0x59, 0x2d, 0xf2, 0x67, 0x20, 0xb8, 0xc7, 0xa5, 0x49, 0x72, 0xb4,
	0xd2, 0x57, 0xf1, 0x45, 0xdf, 0xe9, 0x33, 0xa9, 0xc3, 0x2d, 0x9b, 0x59, 0xd3, 0x3a, 0xb4, 0x8e,
	0x7d, 0x5a, 0x59, 0x53, 0xe3, 0xe7, 0xd2, 0x79, 0x63, 0xbb, 0xd5, 0xf4, 0x9f, 0x91, 0x3b, 0x0d,
	0xf7, 0x55, 0xbe, 0x97, 0x52, 0x7e, 0x98, 0x32, 0x3e, 0xa1, 0x1f, 0xc5, 0x7c, 0x63, 0xbc, 0xcf,
	0x57, 0xba, 0xb5, 0x00, 0x52, 0x6f, 0x0b, 0x12, 0xc1, 0x90, 0xcd, 0x13, 0x0a, 0x7d, 0x48, 0x46,
	0x96, 0xb7, 0xf9, 0x7e, 0x02, 0x7b, 0x2f, 0x81, 0x01, 0x7d, 0x3b, 0x82, 0x59, 0xde, 0xc2, 0xd7,
	0xd3, 0x2f, 0xcf, 0x5e, 0x02, 0x64, 0xf3, 0xf8, 0x84, 0x5e, 0x6e, 0x7b, 0x2d, 0x39, 0x71, 0x36,
	0x4e, 0x08, 0xc7, 0xf4, 0xfd, 0x27, 0xb1, 0xcb, 0x5b, 0xb4, 0xb1, 0x7e, 0xdb, 0xa5, 0xf1, 0x31,
	0x43, 0x13, 0xd9, 0x96, 0x39, 0xe9, 0x29, 0x39, 0xb0, 0x58, 0xa2, 0x45, 0x2d, 0x30, 0xdf, 0xdd,
	0x4e, 0x69, 0x0a, 0xff, 0x1d, 0x41, 0x19, 0x87, 0x71, 0x3b, 0xa5, 0xf5, 0x43, 0xaa, 0x08, 0x59,
	0xdb, 0x2f, 0x1f, 0xc1, 0xe8, 0xe8, 0xee, 0x49, 0xde, 0x2f, 0x19, 0xf6, 0xd2, 0x82, 0x99, 0x9d,
	0x24, 0x82, 0x07, 0xf4, 0x83, 0x8b, 0xb5, 0x65, 0xfb, 0xf1, 0x6f, 0xac, 0xb9, 0x91, 0x1b, 0xbe,
	0x5f, 0x93, 0x6d, 0xe0, 0xd3, 0xa7, 0x64, 0xbf, 0x77, 0x82, 0xcb, 0xef, 0x24, 0xaa, 0x37, 0x5f,
	0xa0, 0xea, 0x2d, 0x38, 0x3b, 0x4e, 0x3c, 0xf7, 0xe9, 0xff, 0x8f, 0xe8, 0x7c, 0x00, 0x9c, 0x7c,
	0x9b, 0xc6, 0xba, 0x23, 0xed, 0x54, 0xaf, 0x2f, 0xb7, 0x96, 0x37, 0x0e, 0x78, 0xb2, 0x6c, 0xef,
	0x0c, 0xa9, 0xc1, 0x62, 0x70, 0x7c, 0xa1, 0xf0, 0x01, 0x14, 0x46, 0x84, 0x1a, 0x75, 0x74, 0x6d,
	0xcb, 0x3b, 0xb6, 0xde, 0x71, 0x69, 0x8f, 0x29, 0x05, 0x7c, 0x61, 0x82, 0x1f, 0x1a, 0xb0, 0x72,
	0x73, 0xdc, 0x4d, 0xda, 0x78, 0xe0, 0x9a, 0xab, 0xee, 0x9b, 0xd5, 0x8a, 0xa9, 0xd9, 0xe1, 0x23,
	0xf2, 0xc6, 0x00, 0x70, 0xba, 0xea, 0x93, 0xb1, 0x8e, 0xbe, 0x4b, 0x76, 0xa5, 0xc7, 0xda, 0xe5,
	0x59, 0xaa, 0xf8, 0xb5, 0x17, 0x2a, 0x9e, 0xf7, 0xa7, 0x8b, 0xbd, 0x14, 0xfe, 0xf8, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa9, 0x65, 0x8a, 0x2f, 0x2d, 0x06, 0x00, 0x00,
}