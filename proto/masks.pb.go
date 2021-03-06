// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: masks.proto

package Pro7Protobuf

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

type Mask_Mode int32

const (
	Mask_MODE_OVERLAY Mask_Mode = 0
	Mask_MODE_KEYHOLE Mask_Mode = 1
)

// Enum value maps for Mask_Mode.
var (
	Mask_Mode_name = map[int32]string{
		0: "MODE_OVERLAY",
		1: "MODE_KEYHOLE",
	}
	Mask_Mode_value = map[string]int32{
		"MODE_OVERLAY": 0,
		"MODE_KEYHOLE": 1,
	}
)

func (x Mask_Mode) Enum() *Mask_Mode {
	p := new(Mask_Mode)
	*p = x
	return p
}

func (x Mask_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mask_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_masks_proto_enumTypes[0].Descriptor()
}

func (Mask_Mode) Type() protoreflect.EnumType {
	return &file_masks_proto_enumTypes[0]
}

func (x Mask_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mask_Mode.Descriptor instead.
func (Mask_Mode) EnumDescriptor() ([]byte, []int) {
	return file_masks_proto_rawDescGZIP(), []int{0, 0}
}

type Mask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   *UUID               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name   string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Color  *Color              `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Mode   Mask_Mode           `protobuf:"varint,4,opt,name=mode,proto3,enum=rv.data.Mask_Mode" json:"mode,omitempty"`
	Shapes []*Graphics_Element `protobuf:"bytes,5,rep,name=shapes,proto3" json:"shapes,omitempty"`
}

func (x *Mask) Reset() {
	*x = Mask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_masks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mask) ProtoMessage() {}

func (x *Mask) ProtoReflect() protoreflect.Message {
	mi := &file_masks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mask.ProtoReflect.Descriptor instead.
func (*Mask) Descriptor() ([]byte, []int) {
	return file_masks_proto_rawDescGZIP(), []int{0}
}

func (x *Mask) GetUuid() *UUID {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *Mask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Mask) GetColor() *Color {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *Mask) GetMode() Mask_Mode {
	if x != nil {
		return x.Mode
	}
	return Mask_MODE_OVERLAY
}

func (x *Mask) GetShapes() []*Graphics_Element {
	if x != nil {
		return x.Shapes
	}
	return nil
}

var File_masks_proto protoreflect.FileDescriptor

var file_masks_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72,
	0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x12, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a,
	0x04, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x76,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x76, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73, 0x2e, 0x45, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x68, 0x61, 0x70, 0x65, 0x73, 0x22, 0x2a, 0x0a,
	0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x56,
	0x45, 0x52, 0x4c, 0x41, 0x59, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x4b, 0x45, 0x59, 0x48, 0x4f, 0x4c, 0x45, 0x10, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x50, 0x72,
	0x6f, 0x37, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_masks_proto_rawDescOnce sync.Once
	file_masks_proto_rawDescData = file_masks_proto_rawDesc
)

func file_masks_proto_rawDescGZIP() []byte {
	file_masks_proto_rawDescOnce.Do(func() {
		file_masks_proto_rawDescData = protoimpl.X.CompressGZIP(file_masks_proto_rawDescData)
	})
	return file_masks_proto_rawDescData
}

var file_masks_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_masks_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_masks_proto_goTypes = []interface{}{
	(Mask_Mode)(0),           // 0: rv.data.Mask.Mode
	(*Mask)(nil),             // 1: rv.data.Mask
	(*UUID)(nil),             // 2: rv.data.UUID
	(*Color)(nil),            // 3: rv.data.Color
	(*Graphics_Element)(nil), // 4: rv.data.Graphics.Element
}
var file_masks_proto_depIdxs = []int32{
	2, // 0: rv.data.Mask.uuid:type_name -> rv.data.UUID
	3, // 1: rv.data.Mask.color:type_name -> rv.data.Color
	0, // 2: rv.data.Mask.mode:type_name -> rv.data.Mask.Mode
	4, // 3: rv.data.Mask.shapes:type_name -> rv.data.Graphics.Element
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_masks_proto_init() }
func file_masks_proto_init() {
	if File_masks_proto != nil {
		return
	}
	file_graphicsData_proto_init()
	file_basicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_masks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mask); i {
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
			RawDescriptor: file_masks_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_masks_proto_goTypes,
		DependencyIndexes: file_masks_proto_depIdxs,
		EnumInfos:         file_masks_proto_enumTypes,
		MessageInfos:      file_masks_proto_msgTypes,
	}.Build()
	File_masks_proto = out.File
	file_masks_proto_rawDesc = nil
	file_masks_proto_goTypes = nil
	file_masks_proto_depIdxs = nil
}
