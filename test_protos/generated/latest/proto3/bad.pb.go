// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bad.proto

/*
Package schema_proto3 is a generated protocol buffer package.

It is generated from these files:
	bad.proto
	floats.proto
	integers.proto
	maps.proto
	people.proto
	planets.proto
	simple.proto
	well_known_types.proto

It has these top-level messages:
	DoubleMessage
	FloatMessage
	Fixed32Message
	Fixed64Message
	Int32Message
	Int64Message
	Sfixed32Message
	Sfixed64Message
	Sint32Message
	Sint64Message
	Uint32Message
	Uint64Message
	BoolMaps
	IntMaps
	StringMaps
	PersonV1
	PersonV2
	PersonV3
	PersonV4
	MyFavoritePlanetsV1
	MyFavoritePlanetsV2
	MyFavoritePlanetsV3
	Empty
	Simple
	Repetitive
	Singleton
	KnownTypes
*/
package schema_proto3

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

func init() { proto.RegisterFile("bad.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 49 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4a, 0x4c, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x84, 0xf0, 0x8c,
	0x93, 0xd8, 0x20, 0x34, 0x20, 0x00, 0x00, 0xff, 0xff, 0x51, 0xa8, 0x49, 0x8d, 0x22, 0x00, 0x00,
	0x00,
}
