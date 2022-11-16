// Copyright 2017 The ObjectHash-Proto Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This is used for tests that ensure that the objecthash of well-known types is
// correctly calculated.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: well_known_types.proto

package proto3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KnownTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnyField         *anypb.Any              `protobuf:"bytes,1,opt,name=any_field,json=anyField,proto3" json:"any_field,omitempty"`
	BoolValueField   *wrapperspb.BoolValue   `protobuf:"bytes,2,opt,name=bool_value_field,json=boolValueField,proto3" json:"bool_value_field,omitempty"`
	BytesValueField  *wrapperspb.BytesValue  `protobuf:"bytes,3,opt,name=bytes_value_field,json=bytesValueField,proto3" json:"bytes_value_field,omitempty"`
	DoubleValueField *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=double_value_field,json=doubleValueField,proto3" json:"double_value_field,omitempty"`
	DurationField    *durationpb.Duration    `protobuf:"bytes,5,opt,name=duration_field,json=durationField,proto3" json:"duration_field,omitempty"`
	FloatValueField  *wrapperspb.FloatValue  `protobuf:"bytes,6,opt,name=float_value_field,json=floatValueField,proto3" json:"float_value_field,omitempty"`
	Int32ValueField  *wrapperspb.Int32Value  `protobuf:"bytes,7,opt,name=int32_value_field,json=int32ValueField,proto3" json:"int32_value_field,omitempty"`
	Int64ValueField  *wrapperspb.Int64Value  `protobuf:"bytes,8,opt,name=int64_value_field,json=int64ValueField,proto3" json:"int64_value_field,omitempty"`
	ListValueField   *structpb.ListValue     `protobuf:"bytes,9,opt,name=list_value_field,json=listValueField,proto3" json:"list_value_field,omitempty"`
	StringValueField *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=string_value_field,json=stringValueField,proto3" json:"string_value_field,omitempty"`
	StructField      *structpb.Struct        `protobuf:"bytes,11,opt,name=struct_field,json=structField,proto3" json:"struct_field,omitempty"`
	TimestampField   *timestamppb.Timestamp  `protobuf:"bytes,12,opt,name=timestamp_field,json=timestampField,proto3" json:"timestamp_field,omitempty"`
	Uint32ValueField *wrapperspb.UInt32Value `protobuf:"bytes,13,opt,name=uint32_value_field,json=uint32ValueField,proto3" json:"uint32_value_field,omitempty"`
	Uint64ValueField *wrapperspb.UInt64Value `protobuf:"bytes,14,opt,name=uint64_value_field,json=uint64ValueField,proto3" json:"uint64_value_field,omitempty"`
	ValueField       *structpb.Value         `protobuf:"bytes,15,opt,name=value_field,json=valueField,proto3" json:"value_field,omitempty"`
}

func (x *KnownTypes) Reset() {
	*x = KnownTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_well_known_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnownTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnownTypes) ProtoMessage() {}

func (x *KnownTypes) ProtoReflect() protoreflect.Message {
	mi := &file_well_known_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnownTypes.ProtoReflect.Descriptor instead.
func (*KnownTypes) Descriptor() ([]byte, []int) {
	return file_well_known_types_proto_rawDescGZIP(), []int{0}
}

func (x *KnownTypes) GetAnyField() *anypb.Any {
	if x != nil {
		return x.AnyField
	}
	return nil
}

func (x *KnownTypes) GetBoolValueField() *wrapperspb.BoolValue {
	if x != nil {
		return x.BoolValueField
	}
	return nil
}

func (x *KnownTypes) GetBytesValueField() *wrapperspb.BytesValue {
	if x != nil {
		return x.BytesValueField
	}
	return nil
}

func (x *KnownTypes) GetDoubleValueField() *wrapperspb.DoubleValue {
	if x != nil {
		return x.DoubleValueField
	}
	return nil
}

func (x *KnownTypes) GetDurationField() *durationpb.Duration {
	if x != nil {
		return x.DurationField
	}
	return nil
}

func (x *KnownTypes) GetFloatValueField() *wrapperspb.FloatValue {
	if x != nil {
		return x.FloatValueField
	}
	return nil
}

func (x *KnownTypes) GetInt32ValueField() *wrapperspb.Int32Value {
	if x != nil {
		return x.Int32ValueField
	}
	return nil
}

func (x *KnownTypes) GetInt64ValueField() *wrapperspb.Int64Value {
	if x != nil {
		return x.Int64ValueField
	}
	return nil
}

func (x *KnownTypes) GetListValueField() *structpb.ListValue {
	if x != nil {
		return x.ListValueField
	}
	return nil
}

func (x *KnownTypes) GetStringValueField() *wrapperspb.StringValue {
	if x != nil {
		return x.StringValueField
	}
	return nil
}

func (x *KnownTypes) GetStructField() *structpb.Struct {
	if x != nil {
		return x.StructField
	}
	return nil
}

func (x *KnownTypes) GetTimestampField() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampField
	}
	return nil
}

func (x *KnownTypes) GetUint32ValueField() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Uint32ValueField
	}
	return nil
}

func (x *KnownTypes) GetUint64ValueField() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Uint64ValueField
	}
	return nil
}

func (x *KnownTypes) GetValueField() *structpb.Value {
	if x != nil {
		return x.ValueField
	}
	return nil
}

var File_well_known_types_proto protoreflect.FileDescriptor

var file_well_known_types_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x65, 0x6c, 0x6c, 0x5f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9b, 0x08, 0x0a, 0x0a, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x31, 0x0a, 0x09, 0x61, 0x6e, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x61, 0x6e, 0x79, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x44, 0x0a, 0x10, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x62, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x47, 0x0a, 0x11, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x4a, 0x0a, 0x12, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x64, 0x6f,
	0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x40,
	0x0a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x47, 0x0a, 0x11, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x47, 0x0a, 0x11, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x47, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x44, 0x0a, 0x10, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x4a, 0x0a, 0x12, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3a, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x43, 0x0a, 0x0f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4a,
	0x0a, 0x12, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x4a, 0x0a, 0x12, 0x75, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42,
	0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x37, 0x72,
	0x33, 0x6e, 0x7a, 0x79, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x68, 0x61, 0x73, 0x68, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_well_known_types_proto_rawDescOnce sync.Once
	file_well_known_types_proto_rawDescData = file_well_known_types_proto_rawDesc
)

func file_well_known_types_proto_rawDescGZIP() []byte {
	file_well_known_types_proto_rawDescOnce.Do(func() {
		file_well_known_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_well_known_types_proto_rawDescData)
	})
	return file_well_known_types_proto_rawDescData
}

var file_well_known_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_well_known_types_proto_goTypes = []interface{}{
	(*KnownTypes)(nil),             // 0: schema.proto3.KnownTypes
	(*anypb.Any)(nil),              // 1: google.protobuf.Any
	(*wrapperspb.BoolValue)(nil),   // 2: google.protobuf.BoolValue
	(*wrapperspb.BytesValue)(nil),  // 3: google.protobuf.BytesValue
	(*wrapperspb.DoubleValue)(nil), // 4: google.protobuf.DoubleValue
	(*durationpb.Duration)(nil),    // 5: google.protobuf.Duration
	(*wrapperspb.FloatValue)(nil),  // 6: google.protobuf.FloatValue
	(*wrapperspb.Int32Value)(nil),  // 7: google.protobuf.Int32Value
	(*wrapperspb.Int64Value)(nil),  // 8: google.protobuf.Int64Value
	(*structpb.ListValue)(nil),     // 9: google.protobuf.ListValue
	(*wrapperspb.StringValue)(nil), // 10: google.protobuf.StringValue
	(*structpb.Struct)(nil),        // 11: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil),  // 12: google.protobuf.Timestamp
	(*wrapperspb.UInt32Value)(nil), // 13: google.protobuf.UInt32Value
	(*wrapperspb.UInt64Value)(nil), // 14: google.protobuf.UInt64Value
	(*structpb.Value)(nil),         // 15: google.protobuf.Value
}
var file_well_known_types_proto_depIdxs = []int32{
	1,  // 0: schema.proto3.KnownTypes.any_field:type_name -> google.protobuf.Any
	2,  // 1: schema.proto3.KnownTypes.bool_value_field:type_name -> google.protobuf.BoolValue
	3,  // 2: schema.proto3.KnownTypes.bytes_value_field:type_name -> google.protobuf.BytesValue
	4,  // 3: schema.proto3.KnownTypes.double_value_field:type_name -> google.protobuf.DoubleValue
	5,  // 4: schema.proto3.KnownTypes.duration_field:type_name -> google.protobuf.Duration
	6,  // 5: schema.proto3.KnownTypes.float_value_field:type_name -> google.protobuf.FloatValue
	7,  // 6: schema.proto3.KnownTypes.int32_value_field:type_name -> google.protobuf.Int32Value
	8,  // 7: schema.proto3.KnownTypes.int64_value_field:type_name -> google.protobuf.Int64Value
	9,  // 8: schema.proto3.KnownTypes.list_value_field:type_name -> google.protobuf.ListValue
	10, // 9: schema.proto3.KnownTypes.string_value_field:type_name -> google.protobuf.StringValue
	11, // 10: schema.proto3.KnownTypes.struct_field:type_name -> google.protobuf.Struct
	12, // 11: schema.proto3.KnownTypes.timestamp_field:type_name -> google.protobuf.Timestamp
	13, // 12: schema.proto3.KnownTypes.uint32_value_field:type_name -> google.protobuf.UInt32Value
	14, // 13: schema.proto3.KnownTypes.uint64_value_field:type_name -> google.protobuf.UInt64Value
	15, // 14: schema.proto3.KnownTypes.value_field:type_name -> google.protobuf.Value
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_well_known_types_proto_init() }
func file_well_known_types_proto_init() {
	if File_well_known_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_well_known_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnownTypes); i {
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
			RawDescriptor: file_well_known_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_well_known_types_proto_goTypes,
		DependencyIndexes: file_well_known_types_proto_depIdxs,
		MessageInfos:      file_well_known_types_proto_msgTypes,
	}.Build()
	File_well_known_types_proto = out.File
	file_well_known_types_proto_rawDesc = nil
	file_well_known_types_proto_goTypes = nil
	file_well_known_types_proto_depIdxs = nil
}
