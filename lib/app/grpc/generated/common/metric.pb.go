// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/metric.proto

package common

import (
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

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	MType string `protobuf:"bytes,2,opt,name=MType,proto3" json:"MType,omitempty"`
	// Types that are assignable to OptionalDelta:
	//
	//	*Metric_Delta
	OptionalDelta isMetric_OptionalDelta `protobuf_oneof:"optional_delta"`
	// Types that are assignable to OptionalValue:
	//
	//	*Metric_Value
	OptionalValue isMetric_OptionalValue `protobuf_oneof:"optional_value"`
	// Types that are assignable to OptionalHash:
	//
	//	*Metric_Hash
	OptionalHash isMetric_OptionalHash `protobuf_oneof:"optional_hash"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_proto_metric_proto_rawDescGZIP(), []int{0}
}

func (x *Metric) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Metric) GetMType() string {
	if x != nil {
		return x.MType
	}
	return ""
}

func (m *Metric) GetOptionalDelta() isMetric_OptionalDelta {
	if m != nil {
		return m.OptionalDelta
	}
	return nil
}

func (x *Metric) GetDelta() int64 {
	if x, ok := x.GetOptionalDelta().(*Metric_Delta); ok {
		return x.Delta
	}
	return 0
}

func (m *Metric) GetOptionalValue() isMetric_OptionalValue {
	if m != nil {
		return m.OptionalValue
	}
	return nil
}

func (x *Metric) GetValue() float64 {
	if x, ok := x.GetOptionalValue().(*Metric_Value); ok {
		return x.Value
	}
	return 0
}

func (m *Metric) GetOptionalHash() isMetric_OptionalHash {
	if m != nil {
		return m.OptionalHash
	}
	return nil
}

func (x *Metric) GetHash() string {
	if x, ok := x.GetOptionalHash().(*Metric_Hash); ok {
		return x.Hash
	}
	return ""
}

type isMetric_OptionalDelta interface {
	isMetric_OptionalDelta()
}

type Metric_Delta struct {
	Delta int64 `protobuf:"varint,3,opt,name=Delta,proto3,oneof"`
}

func (*Metric_Delta) isMetric_OptionalDelta() {}

type isMetric_OptionalValue interface {
	isMetric_OptionalValue()
}

type Metric_Value struct {
	Value float64 `protobuf:"fixed64,4,opt,name=Value,proto3,oneof"`
}

func (*Metric_Value) isMetric_OptionalValue() {}

type isMetric_OptionalHash interface {
	isMetric_OptionalHash()
}

type Metric_Hash struct {
	Hash string `protobuf:"bytes,5,opt,name=Hash,proto3,oneof"`
}

func (*Metric_Hash) isMetric_OptionalHash() {}

var File_proto_metric_proto protoreflect.FileDescriptor

var file_proto_metric_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x22, 0xa9, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x04, 0x48, 0x61, 0x73, 0x68, 0x42, 0x10, 0x0a, 0x0e, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x42, 0x10, 0x0a, 0x0e, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0f, 0x0a,
	0x0d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x42, 0x4c,
	0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x69,
	0x67, 0x61, 0x6e, 0x67, 0x2f, 0x59, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x63, 0x75, 0x6d, 0x47, 0x6f, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x2f, 0x6c,
	0x69, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_metric_proto_rawDescOnce sync.Once
	file_proto_metric_proto_rawDescData = file_proto_metric_proto_rawDesc
)

func file_proto_metric_proto_rawDescGZIP() []byte {
	file_proto_metric_proto_rawDescOnce.Do(func() {
		file_proto_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_metric_proto_rawDescData)
	})
	return file_proto_metric_proto_rawDescData
}

var file_proto_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_metric_proto_goTypes = []interface{}{
	(*Metric)(nil), // 0: proto.common.Metric
}
var file_proto_metric_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_metric_proto_init() }
func file_proto_metric_proto_init() {
	if File_proto_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_metric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
	file_proto_metric_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Metric_Delta)(nil),
		(*Metric_Value)(nil),
		(*Metric_Hash)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_metric_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_metric_proto_goTypes,
		DependencyIndexes: file_proto_metric_proto_depIdxs,
		MessageInfos:      file_proto_metric_proto_msgTypes,
	}.Build()
	File_proto_metric_proto = out.File
	file_proto_metric_proto_rawDesc = nil
	file_proto_metric_proto_goTypes = nil
	file_proto_metric_proto_depIdxs = nil
}
