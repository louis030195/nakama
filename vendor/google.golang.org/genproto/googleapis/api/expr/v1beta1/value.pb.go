// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/expr/v1beta1/value.proto

package expr // import "google.golang.org/genproto/googleapis/api/expr/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Represents a CEL value.
//
// This is similar to `google.protobuf.Value`, but can represent CEL's full
// range of values.
type Value struct {
	// Required. The valid kinds of values.
	//
	// Types that are valid to be assigned to Kind:
	//	*Value_NullValue
	//	*Value_BoolValue
	//	*Value_Int64Value
	//	*Value_Uint64Value
	//	*Value_DoubleValue
	//	*Value_StringValue
	//	*Value_BytesValue
	//	*Value_EnumValue
	//	*Value_ObjectValue
	//	*Value_MapValue
	//	*Value_ListValue
	//	*Value_TypeValue
	Kind                 isValue_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_6c2bb4fc7e0e374b, []int{0}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Kind interface {
	isValue_Kind()
}

type Value_NullValue struct {
	NullValue _struct.NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_Int64Value struct {
	Int64Value int64 `protobuf:"varint,3,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Value_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,4,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,6,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,7,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type Value_EnumValue struct {
	EnumValue *EnumValue `protobuf:"bytes,9,opt,name=enum_value,json=enumValue,proto3,oneof"`
}

type Value_ObjectValue struct {
	ObjectValue *any.Any `protobuf:"bytes,10,opt,name=object_value,json=objectValue,proto3,oneof"`
}

type Value_MapValue struct {
	MapValue *MapValue `protobuf:"bytes,11,opt,name=map_value,json=mapValue,proto3,oneof"`
}

type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,12,opt,name=list_value,json=listValue,proto3,oneof"`
}

type Value_TypeValue struct {
	TypeValue string `protobuf:"bytes,15,opt,name=type_value,json=typeValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_Int64Value) isValue_Kind() {}

func (*Value_Uint64Value) isValue_Kind() {}

func (*Value_DoubleValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_BytesValue) isValue_Kind() {}

func (*Value_EnumValue) isValue_Kind() {}

func (*Value_ObjectValue) isValue_Kind() {}

func (*Value_MapValue) isValue_Kind() {}

func (*Value_ListValue) isValue_Kind() {}

func (*Value_TypeValue) isValue_Kind() {}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Value) GetNullValue() _struct.NullValue {
	if x, ok := m.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return _struct.NullValue_NULL_VALUE
}

func (m *Value) GetBoolValue() bool {
	if x, ok := m.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Value) GetInt64Value() int64 {
	if x, ok := m.GetKind().(*Value_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *Value) GetUint64Value() uint64 {
	if x, ok := m.GetKind().(*Value_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetKind().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBytesValue() []byte {
	if x, ok := m.GetKind().(*Value_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *Value) GetEnumValue() *EnumValue {
	if x, ok := m.GetKind().(*Value_EnumValue); ok {
		return x.EnumValue
	}
	return nil
}

func (m *Value) GetObjectValue() *any.Any {
	if x, ok := m.GetKind().(*Value_ObjectValue); ok {
		return x.ObjectValue
	}
	return nil
}

func (m *Value) GetMapValue() *MapValue {
	if x, ok := m.GetKind().(*Value_MapValue); ok {
		return x.MapValue
	}
	return nil
}

func (m *Value) GetListValue() *ListValue {
	if x, ok := m.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

func (m *Value) GetTypeValue() string {
	if x, ok := m.GetKind().(*Value_TypeValue); ok {
		return x.TypeValue
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_NullValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_Int64Value)(nil),
		(*Value_Uint64Value)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BytesValue)(nil),
		(*Value_EnumValue)(nil),
		(*Value_ObjectValue)(nil),
		(*Value_MapValue)(nil),
		(*Value_ListValue)(nil),
		(*Value_TypeValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// kind
	switch x := m.Kind.(type) {
	case *Value_NullValue:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.NullValue))
	case *Value_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_Int64Value:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Value))
	case *Value_Uint64Value:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint64Value))
	case *Value_DoubleValue:
		b.EncodeVarint(5<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *Value_StringValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Value_BytesValue:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BytesValue)
	case *Value_EnumValue:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EnumValue); err != nil {
			return err
		}
	case *Value_ObjectValue:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ObjectValue); err != nil {
			return err
		}
	case *Value_MapValue:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MapValue); err != nil {
			return err
		}
	case *Value_ListValue:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListValue); err != nil {
			return err
		}
	case *Value_TypeValue:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.TypeValue)
	case nil:
	default:
		return fmt.Errorf("Value.Kind has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // kind.null_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_NullValue{_struct.NullValue(x)}
		return true, err
	case 2: // kind.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_BoolValue{x != 0}
		return true, err
	case 3: // kind.int64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_Int64Value{int64(x)}
		return true, err
	case 4: // kind.uint64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_Uint64Value{x}
		return true, err
	case 5: // kind.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Kind = &Value_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 6: // kind.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Kind = &Value_StringValue{x}
		return true, err
	case 7: // kind.bytes_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Kind = &Value_BytesValue{x}
		return true, err
	case 9: // kind.enum_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EnumValue)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_EnumValue{msg}
		return true, err
	case 10: // kind.object_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(any.Any)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_ObjectValue{msg}
		return true, err
	case 11: // kind.map_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MapValue)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_MapValue{msg}
		return true, err
	case 12: // kind.list_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListValue)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_ListValue{msg}
		return true, err
	case 15: // kind.type_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Kind = &Value_TypeValue{x}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// kind
	switch x := m.Kind.(type) {
	case *Value_NullValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.NullValue))
	case *Value_BoolValue:
		n += 1 // tag and wire
		n += 1
	case *Value_Int64Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Int64Value))
	case *Value_Uint64Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Uint64Value))
	case *Value_DoubleValue:
		n += 1 // tag and wire
		n += 8
	case *Value_StringValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Value_BytesValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.BytesValue)))
		n += len(x.BytesValue)
	case *Value_EnumValue:
		s := proto.Size(x.EnumValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_ObjectValue:
		s := proto.Size(x.ObjectValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_MapValue:
		s := proto.Size(x.MapValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_ListValue:
		s := proto.Size(x.ListValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_TypeValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.TypeValue)))
		n += len(x.TypeValue)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// An enum value.
type EnumValue struct {
	// The fully qualified name of the enum type.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The value of the enum.
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumValue) Reset()         { *m = EnumValue{} }
func (m *EnumValue) String() string { return proto.CompactTextString(m) }
func (*EnumValue) ProtoMessage()    {}
func (*EnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_6c2bb4fc7e0e374b, []int{1}
}
func (m *EnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumValue.Unmarshal(m, b)
}
func (m *EnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumValue.Marshal(b, m, deterministic)
}
func (dst *EnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumValue.Merge(dst, src)
}
func (m *EnumValue) XXX_Size() int {
	return xxx_messageInfo_EnumValue.Size(m)
}
func (m *EnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_EnumValue proto.InternalMessageInfo

func (m *EnumValue) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EnumValue) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// A list.
//
// Wrapped in a message so 'not set' and empty can be differentiated, which is
// required for use in a 'oneof'.
type ListValue struct {
	// The ordered values in the list.
	Values               []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListValue) Reset()         { *m = ListValue{} }
func (m *ListValue) String() string { return proto.CompactTextString(m) }
func (*ListValue) ProtoMessage()    {}
func (*ListValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_6c2bb4fc7e0e374b, []int{2}
}
func (m *ListValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListValue.Unmarshal(m, b)
}
func (m *ListValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListValue.Marshal(b, m, deterministic)
}
func (dst *ListValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListValue.Merge(dst, src)
}
func (m *ListValue) XXX_Size() int {
	return xxx_messageInfo_ListValue.Size(m)
}
func (m *ListValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListValue proto.InternalMessageInfo

func (m *ListValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// A map.
//
// Wrapped in a message so 'not set' and empty can be differentiated, which is
// required for use in a 'oneof'.
type MapValue struct {
	// The set of map entries.
	//
	// CEL has fewer restrictions on keys, so a protobuf map represenation
	// cannot be used.
	Entries              []*MapValue_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MapValue) Reset()         { *m = MapValue{} }
func (m *MapValue) String() string { return proto.CompactTextString(m) }
func (*MapValue) ProtoMessage()    {}
func (*MapValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_6c2bb4fc7e0e374b, []int{3}
}
func (m *MapValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapValue.Unmarshal(m, b)
}
func (m *MapValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapValue.Marshal(b, m, deterministic)
}
func (dst *MapValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapValue.Merge(dst, src)
}
func (m *MapValue) XXX_Size() int {
	return xxx_messageInfo_MapValue.Size(m)
}
func (m *MapValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MapValue.DiscardUnknown(m)
}

var xxx_messageInfo_MapValue proto.InternalMessageInfo

func (m *MapValue) GetEntries() []*MapValue_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// An entry in the map.
type MapValue_Entry struct {
	// The key.
	//
	// Must be unique with in the map.
	// Currently only boolean, int, uint, and string values can be keys.
	Key *Value `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value.
	Value                *Value   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapValue_Entry) Reset()         { *m = MapValue_Entry{} }
func (m *MapValue_Entry) String() string { return proto.CompactTextString(m) }
func (*MapValue_Entry) ProtoMessage()    {}
func (*MapValue_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_6c2bb4fc7e0e374b, []int{3, 0}
}
func (m *MapValue_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapValue_Entry.Unmarshal(m, b)
}
func (m *MapValue_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapValue_Entry.Marshal(b, m, deterministic)
}
func (dst *MapValue_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapValue_Entry.Merge(dst, src)
}
func (m *MapValue_Entry) XXX_Size() int {
	return xxx_messageInfo_MapValue_Entry.Size(m)
}
func (m *MapValue_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_MapValue_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_MapValue_Entry proto.InternalMessageInfo

func (m *MapValue_Entry) GetKey() *Value {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MapValue_Entry) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Value)(nil), "google.api.expr.v1beta1.Value")
	proto.RegisterType((*EnumValue)(nil), "google.api.expr.v1beta1.EnumValue")
	proto.RegisterType((*ListValue)(nil), "google.api.expr.v1beta1.ListValue")
	proto.RegisterType((*MapValue)(nil), "google.api.expr.v1beta1.MapValue")
	proto.RegisterType((*MapValue_Entry)(nil), "google.api.expr.v1beta1.MapValue.Entry")
}

func init() {
	proto.RegisterFile("google/api/expr/v1beta1/value.proto", fileDescriptor_value_6c2bb4fc7e0e374b)
}

var fileDescriptor_value_6c2bb4fc7e0e374b = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x6b, 0xfa, 0xb2, 0xe6, 0x69, 0x05, 0x92, 0x35, 0x89, 0x51, 0x10, 0x64, 0xdd, 0x81,
	0x9c, 0x1c, 0x56, 0xc6, 0x24, 0xd4, 0x0b, 0xeb, 0x34, 0xa9, 0x07, 0x40, 0x53, 0x0e, 0x1c, 0xb8,
	0xa0, 0xa4, 0x33, 0x51, 0xa8, 0x63, 0x47, 0x89, 0x3d, 0x91, 0x2f, 0xc7, 0x07, 0xe0, 0x13, 0x71,
	0x44, 0x7e, 0x0b, 0x63, 0x53, 0xd5, 0x5b, 0x9e, 0xbf, 0x7f, 0x7f, 0x3f, 0x2f, 0x7e, 0x14, 0x38,
	0xc9, 0x85, 0xc8, 0x19, 0x8d, 0xd3, 0xaa, 0x88, 0xe9, 0xcf, 0xaa, 0x8e, 0x6f, 0x4f, 0x33, 0x2a,
	0xd3, 0xd3, 0xf8, 0x36, 0x65, 0x8a, 0x92, 0xaa, 0x16, 0x52, 0xe0, 0xa7, 0x16, 0x22, 0x69, 0x55,
	0x10, 0x0d, 0x11, 0x07, 0xcd, 0x9e, 0x39, 0xb7, 0xc1, 0x32, 0xf5, 0x3d, 0x4e, 0x79, 0x6b, 0x3d,
	0xb3, 0x17, 0xf7, 0x8f, 0x1a, 0x59, 0xab, 0x8d, 0xb4, 0xa7, 0xf3, 0xdf, 0x03, 0x18, 0x7e, 0xd1,
	0x19, 0xf0, 0x12, 0x80, 0x2b, 0xc6, 0xbe, 0x99, 0x7c, 0x47, 0x28, 0x44, 0xd1, 0xe3, 0xc5, 0x8c,
	0xb8, 0x84, 0xde, 0x4c, 0x3e, 0x2b, 0xc6, 0x0c, 0xbf, 0xee, 0x25, 0x01, 0xf7, 0x01, 0x7e, 0x05,
	0x90, 0x09, 0xe1, 0xcd, 0x8f, 0x42, 0x14, 0x8d, 0x35, 0xa0, 0x35, 0x0b, 0x1c, 0xc3, 0xa4, 0xe0,
	0xf2, 0xfc, 0xcc, 0x11, 0xfd, 0x10, 0x45, 0xfd, 0x75, 0x2f, 0x01, 0x23, 0x5a, 0xe4, 0x04, 0xa6,
	0xea, 0x2e, 0x33, 0x08, 0x51, 0x34, 0x58, 0xf7, 0x92, 0x89, 0xfa, 0x1f, 0xba, 0x11, 0x2a, 0x63,
	0xd4, 0x41, 0xc3, 0x10, 0x45, 0x48, 0x43, 0x56, 0xed, 0xa0, 0x46, 0xd6, 0x05, 0xcf, 0x1d, 0x34,
	0x0a, 0x51, 0x14, 0x68, 0xc8, 0xaa, 0x5d, 0x45, 0x59, 0x2b, 0x69, 0xe3, 0x98, 0x83, 0x10, 0x45,
	0x53, 0x5d, 0x91, 0x11, 0x2d, 0x72, 0x09, 0x40, 0xb9, 0x2a, 0x1d, 0x11, 0x84, 0x28, 0x9a, 0x2c,
	0xe6, 0x64, 0xc7, 0x1b, 0x90, 0x2b, 0xae, 0xca, 0x6e, 0x34, 0xd4, 0x07, 0xf8, 0x3d, 0x4c, 0x45,
	0xf6, 0x83, 0x6e, 0xa4, 0xbb, 0x06, 0xcc, 0x35, 0x87, 0x0f, 0x26, 0x7b, 0xc1, 0x5b, 0x5d, 0xa2,
	0x65, 0xad, 0xf5, 0x03, 0x04, 0x65, 0x5a, 0x39, 0xdf, 0xc4, 0xf8, 0x8e, 0x77, 0xa6, 0xff, 0x94,
	0x56, 0x3e, 0xfb, 0xb8, 0x74, 0xdf, 0xba, 0x03, 0x56, 0x34, 0x3e, 0xf5, 0x74, 0x4f, 0x07, 0x1f,
	0x8b, 0x46, 0x76, 0x1d, 0x30, 0x1f, 0xe8, 0xc7, 0x95, 0x6d, 0xe5, 0x27, 0xfe, 0xc4, 0x0d, 0x33,
	0xd0, 0x9a, 0x01, 0x56, 0x23, 0x18, 0x6c, 0x0b, 0x7e, 0x33, 0x7f, 0x07, 0x41, 0x37, 0x04, 0x8c,
	0x61, 0xa0, 0x09, 0xb3, 0x49, 0x41, 0x62, 0xbe, 0xf1, 0x21, 0x0c, 0xff, 0x6d, 0xc8, 0x30, 0xb1,
	0xc1, 0xfc, 0x12, 0x82, 0x2e, 0x33, 0x3e, 0x87, 0x91, 0x51, 0x9b, 0x23, 0x14, 0xf6, 0xa3, 0xc9,
	0xe2, 0xe5, 0xce, 0x6a, 0x0d, 0x9f, 0x38, 0x7a, 0xfe, 0x0b, 0xc1, 0xd8, 0x8f, 0x00, 0x5f, 0xc0,
	0x01, 0xe5, 0xb2, 0x2e, 0xba, 0x5b, 0x5e, 0xef, 0x1d, 0x1b, 0xb9, 0xe2, 0xb2, 0x6e, 0x13, 0xef,
	0x9b, 0x09, 0x18, 0x1a, 0x05, 0xbf, 0x81, 0xfe, 0x96, 0xb6, 0xa6, 0x8d, 0xfd, 0xd5, 0x68, 0x14,
	0x9f, 0xdd, 0xed, 0x72, 0xbf, 0xc7, 0xc2, 0xab, 0x2d, 0x3c, 0xdf, 0x88, 0x72, 0x17, 0xbb, 0x02,
	0x03, 0x5f, 0xeb, 0x6d, 0xb9, 0x46, 0x5f, 0x97, 0x0e, 0xcb, 0x05, 0x4b, 0x79, 0x4e, 0x44, 0x9d,
	0xc7, 0x39, 0xe5, 0x66, 0x97, 0x62, 0x7b, 0x94, 0x56, 0x45, 0xf3, 0xe0, 0x67, 0xb2, 0xd4, 0xc1,
	0x1f, 0x84, 0xb2, 0x91, 0x41, 0xdf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x11, 0xf9, 0xd9,
	0x76, 0x04, 0x00, 0x00,
}