// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/infobloxopen/atlas-app-toolkit/rpc/errfields/error_fields.proto

/*
Package errfields is a generated protocol buffer package.

It is generated from these files:
	github.com/infobloxopen/atlas-app-toolkit/rpc/errfields/error_fields.proto

It has these top-level messages:
	FieldInfo
	StringListValue
*/
package errfields

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// FieldsInfo is a default representation of field details that conforms
// REST API Syntax Specification
type FieldInfo struct {
	Fields map[string]*StringListValue `protobuf:"bytes,1,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FieldInfo) Reset()                    { *m = FieldInfo{} }
func (m *FieldInfo) String() string            { return proto.CompactTextString(m) }
func (*FieldInfo) ProtoMessage()               {}
func (*FieldInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FieldInfo) GetFields() map[string]*StringListValue {
	if m != nil {
		return m.Fields
	}
	return nil
}

type StringListValue struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *StringListValue) Reset()                    { *m = StringListValue{} }
func (m *StringListValue) String() string            { return proto.CompactTextString(m) }
func (*StringListValue) ProtoMessage()               {}
func (*StringListValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StringListValue) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*FieldInfo)(nil), "atlas.rpc.FieldInfo")
	proto.RegisterType((*StringListValue)(nil), "atlas.rpc.StringListValue")
}

func init() {
	proto.RegisterFile("github.com/infobloxopen/atlas-app-toolkit/rpc/errfields/error_fields.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcc, 0x4b, 0xcb, 0x4f, 0xca, 0xc9, 0xaf, 0xc8,
	0x2f, 0x48, 0xcd, 0xd3, 0x4f, 0x2c, 0xc9, 0x49, 0x2c, 0xd6, 0x4d, 0x2c, 0x28, 0xd0, 0x2d, 0xc9,
	0xcf, 0xcf, 0xc9, 0xce, 0x2c, 0xd1, 0x2f, 0x2a, 0x48, 0xd6, 0x4f, 0x2d, 0x2a, 0x4a, 0xcb, 0x4c,
	0xcd, 0x49, 0x29, 0x06, 0xb1, 0xf2, 0x8b, 0xe2, 0x21, 0x1c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c,
	0x21, 0x4e, 0xb0, 0x1e, 0xbd, 0xa2, 0x82, 0x64, 0xa5, 0x39, 0x8c, 0x5c, 0x9c, 0x6e, 0x20, 0x39,
	0xcf, 0xbc, 0xb4, 0x7c, 0x21, 0x0b, 0x2e, 0x36, 0x88, 0x42, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e,
	0x23, 0x05, 0x3d, 0xb8, 0x4a, 0x3d, 0xb8, 0x2a, 0x08, 0xab, 0xd8, 0x35, 0xaf, 0xa4, 0xa8, 0x32,
	0x08, 0xaa, 0x5e, 0x2a, 0x94, 0x8b, 0x1b, 0x49, 0x58, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x52,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14, 0x32, 0xe0, 0x62, 0x2d, 0x4b, 0xcc, 0x29,
	0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0x42, 0x32, 0x39, 0xb8, 0xa4, 0x28, 0x33,
	0x2f, 0xdd, 0x27, 0xb3, 0xb8, 0x24, 0x0c, 0xa4, 0x22, 0x08, 0xa2, 0xd0, 0x8a, 0xc9, 0x82, 0x51,
	0x49, 0x93, 0x8b, 0x1f, 0x4d, 0x56, 0x48, 0x8c, 0x8b, 0x0d, 0x2c, 0x0f, 0x71, 0x23, 0x67, 0x10,
	0x94, 0xe7, 0xe4, 0x1c, 0xe5, 0x48, 0x66, 0x10, 0x59, 0xc3, 0x59, 0x49, 0x6c, 0xe0, 0x00, 0x32,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x55, 0xd1, 0x5a, 0x6e, 0x01, 0x00, 0x00,
}
