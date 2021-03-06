// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	schema.proto

It has these top-level messages:
	KafkaEvent
	Event
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KafkaEvent struct {
	Data             *Event                     `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	UtcGeneratedTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=utc_generated_time,json=utcGeneratedTime" json:"utc_generated_time,omitempty"`
}

func (m *KafkaEvent) Reset()                    { *m = KafkaEvent{} }
func (m *KafkaEvent) String() string            { return proto.CompactTextString(m) }
func (*KafkaEvent) ProtoMessage()               {}
func (*KafkaEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KafkaEvent) GetData() *Event {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *KafkaEvent) GetUtcGeneratedTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.UtcGeneratedTime
	}
	return nil
}

type Event struct {
	Id        uint32                     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Account   uint32                     `protobuf:"varint,3,opt,name=account" json:"account,omitempty"`
	Field1    string                     `protobuf:"bytes,4,opt,name=field1" json:"field1,omitempty"`
	Field2    string                     `protobuf:"bytes,5,opt,name=field2" json:"field2,omitempty"`
	Field3    string                     `protobuf:"bytes,6,opt,name=field3" json:"field3,omitempty"`
	Field4    string                     `protobuf:"bytes,7,opt,name=field4" json:"field4,omitempty"`
	Field5    string                     `protobuf:"bytes,8,opt,name=field5" json:"field5,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Event) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Event) GetAccount() uint32 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *Event) GetField1() string {
	if m != nil {
		return m.Field1
	}
	return ""
}

func (m *Event) GetField2() string {
	if m != nil {
		return m.Field2
	}
	return ""
}

func (m *Event) GetField3() string {
	if m != nil {
		return m.Field3
	}
	return ""
}

func (m *Event) GetField4() string {
	if m != nil {
		return m.Field4
	}
	return ""
}

func (m *Event) GetField5() string {
	if m != nil {
		return m.Field5
	}
	return ""
}

func init() {
	proto.RegisterType((*KafkaEvent)(nil), "main.KafkaEvent")
	proto.RegisterType((*Event)(nil), "main.Event")
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0xd3, 0x5a, 0x8a, 0x0c, 0x6a, 0xcc, 0x1e, 0xcc, 0x84, 0x0b, 0x84, 0x13, 0xa7, 0x25,
	0xb6, 0x90, 0xf8, 0x02, 0x46, 0x13, 0x6f, 0x8d, 0x77, 0x32, 0x74, 0xb7, 0x75, 0x23, 0xed, 0x12,
	0x98, 0xea, 0x03, 0xfb, 0x22, 0xa6, 0x5b, 0xbb, 0xf6, 0xc8, 0x71, 0xbe, 0xf9, 0xe7, 0xff, 0x06,
	0x6e, 0xce, 0xf9, 0x87, 0xae, 0x48, 0x1e, 0x4f, 0x96, 0xad, 0x88, 0x2a, 0x32, 0xf5, 0x6c, 0x5e,
	0x5a, 0x5b, 0x1e, 0xf4, 0xda, 0xb1, 0x7d, 0x53, 0xac, 0xd9, 0x54, 0xfa, 0xcc, 0x54, 0x1d, 0xbb,
	0xd8, 0xf2, 0x1b, 0xe0, 0x8d, 0x8a, 0x4f, 0x7a, 0xfe, 0xd2, 0x35, 0x8b, 0x39, 0x44, 0x8a, 0x98,
	0x30, 0x58, 0x04, 0xab, 0x69, 0x32, 0x95, 0x6d, 0x87, 0x74, 0xab, 0xcc, 0x2d, 0xc4, 0x2b, 0x88,
	0x86, 0xf3, 0x5d, 0xa9, 0x6b, 0x7d, 0x22, 0xd6, 0x6a, 0xd7, 0xf6, 0x61, 0xe8, 0xe2, 0x33, 0xd9,
	0xc9, 0x64, 0x2f, 0x93, 0xef, 0xbd, 0x2c, 0xbb, 0x6f, 0x38, 0x7f, 0xe9, 0x8f, 0x5a, 0xbc, 0xfc,
	0x09, 0x60, 0xd4, 0x49, 0xef, 0x20, 0x34, 0xca, 0x29, 0x6f, 0xb3, 0xd0, 0x28, 0xf1, 0x04, 0x13,
	0xff, 0xe5, 0x05, 0xd5, 0xff, 0x61, 0x81, 0x30, 0xa6, 0x3c, 0xb7, 0x4d, 0xcd, 0x78, 0xe5, 0xea,
	0xfa, 0x51, 0x3c, 0x40, 0x5c, 0x18, 0x7d, 0x50, 0x8f, 0x18, 0x2d, 0x82, 0xd5, 0x24, 0xfb, 0x9b,
	0x3c, 0x4f, 0x70, 0x34, 0xe0, 0x89, 0xe7, 0x29, 0xc6, 0x03, 0x9e, 0x7a, 0xbe, 0xc1, 0xf1, 0x80,
	0x6f, 0x3c, 0xdf, 0xe2, 0xf5, 0x80, 0x6f, 0xf7, 0xb1, 0x7b, 0x38, 0xfd, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x09, 0x89, 0xf0, 0xcd, 0x9c, 0x01, 0x00, 0x00,
}
