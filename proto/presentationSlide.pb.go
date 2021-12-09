// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: presentationSlide.proto

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

type PresentationSlide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseSlide          *Slide                   `protobuf:"bytes,1,opt,name=base_slide,json=baseSlide,proto3" json:"base_slide,omitempty"`
	Notes              *PresentationSlide_Notes `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
	TemplateGuidelines []*AlignmentGuide        `protobuf:"bytes,3,rep,name=template_guidelines,json=templateGuidelines,proto3" json:"template_guidelines,omitempty"`
	ChordChart         *URL                     `protobuf:"bytes,4,opt,name=chord_chart,json=chordChart,proto3" json:"chord_chart,omitempty"`
	Transition         *Transition              `protobuf:"bytes,5,opt,name=transition,proto3" json:"transition,omitempty"`
}

func (x *PresentationSlide) Reset() {
	*x = PresentationSlide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentationSlide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresentationSlide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresentationSlide) ProtoMessage() {}

func (x *PresentationSlide) ProtoReflect() protoreflect.Message {
	mi := &file_presentationSlide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresentationSlide.ProtoReflect.Descriptor instead.
func (*PresentationSlide) Descriptor() ([]byte, []int) {
	return file_presentationSlide_proto_rawDescGZIP(), []int{0}
}

func (x *PresentationSlide) GetBaseSlide() *Slide {
	if x != nil {
		return x.BaseSlide
	}
	return nil
}

func (x *PresentationSlide) GetNotes() *PresentationSlide_Notes {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *PresentationSlide) GetTemplateGuidelines() []*AlignmentGuide {
	if x != nil {
		return x.TemplateGuidelines
	}
	return nil
}

func (x *PresentationSlide) GetChordChart() *URL {
	if x != nil {
		return x.ChordChart
	}
	return nil
}

func (x *PresentationSlide) GetTransition() *Transition {
	if x != nil {
		return x.Transition
	}
	return nil
}

type PresentationSlide_Notes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RtfData    []byte                    `protobuf:"bytes,1,opt,name=rtf_data,json=rtfData,proto3" json:"rtf_data,omitempty"`
	Attributes *Graphics_Text_Attributes `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *PresentationSlide_Notes) Reset() {
	*x = PresentationSlide_Notes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_presentationSlide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresentationSlide_Notes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresentationSlide_Notes) ProtoMessage() {}

func (x *PresentationSlide_Notes) ProtoReflect() protoreflect.Message {
	mi := &file_presentationSlide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresentationSlide_Notes.ProtoReflect.Descriptor instead.
func (*PresentationSlide_Notes) Descriptor() ([]byte, []int) {
	return file_presentationSlide_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PresentationSlide_Notes) GetRtfData() []byte {
	if x != nil {
		return x.RtfData
	}
	return nil
}

func (x *PresentationSlide_Notes) GetAttributes() *Graphics_Text_Attributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_presentationSlide_proto protoreflect.FileDescriptor

var file_presentationSlide_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6c,
	0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x76, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x0b, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x62, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x75, 0x69, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x03, 0x0a, 0x11, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6c, 0x69, 0x64, 0x65,
	0x12, 0x2d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53,
	0x6c, 0x69, 0x64, 0x65, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x12,
	0x36, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x73,
	0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x41,
	0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x75, 0x69, 0x64, 0x65, 0x52, 0x12, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x2d, 0x0a, 0x0b, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x55, 0x52, 0x4c, 0x52, 0x0a, 0x63, 0x68, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x12, 0x33, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x65, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x72, 0x74, 0x66, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x72, 0x74, 0x66, 0x44, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x73,
	0x2e, 0x54, 0x65, 0x78, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d,
	0x2f, 0x50, 0x72, 0x6f, 0x37, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_presentationSlide_proto_rawDescOnce sync.Once
	file_presentationSlide_proto_rawDescData = file_presentationSlide_proto_rawDesc
)

func file_presentationSlide_proto_rawDescGZIP() []byte {
	file_presentationSlide_proto_rawDescOnce.Do(func() {
		file_presentationSlide_proto_rawDescData = protoimpl.X.CompressGZIP(file_presentationSlide_proto_rawDescData)
	})
	return file_presentationSlide_proto_rawDescData
}

var file_presentationSlide_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_presentationSlide_proto_goTypes = []interface{}{
	(*PresentationSlide)(nil),        // 0: rv.data.PresentationSlide
	(*PresentationSlide_Notes)(nil),  // 1: rv.data.PresentationSlide.Notes
	(*Slide)(nil),                    // 2: rv.data.Slide
	(*AlignmentGuide)(nil),           // 3: rv.data.AlignmentGuide
	(*URL)(nil),                      // 4: rv.data.URL
	(*Transition)(nil),               // 5: rv.data.Transition
	(*Graphics_Text_Attributes)(nil), // 6: rv.data.Graphics.Text.Attributes
}
var file_presentationSlide_proto_depIdxs = []int32{
	2, // 0: rv.data.PresentationSlide.base_slide:type_name -> rv.data.Slide
	1, // 1: rv.data.PresentationSlide.notes:type_name -> rv.data.PresentationSlide.Notes
	3, // 2: rv.data.PresentationSlide.template_guidelines:type_name -> rv.data.AlignmentGuide
	4, // 3: rv.data.PresentationSlide.chord_chart:type_name -> rv.data.URL
	5, // 4: rv.data.PresentationSlide.transition:type_name -> rv.data.Transition
	6, // 5: rv.data.PresentationSlide.Notes.attributes:type_name -> rv.data.Graphics.Text.Attributes
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_presentationSlide_proto_init() }
func file_presentationSlide_proto_init() {
	if File_presentationSlide_proto != nil {
		return
	}
	file_slide_proto_init()
	file_basicTypes_proto_init()
	file_alignmentGuide_proto_init()
	file_effects_proto_init()
	file_graphicsData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_presentationSlide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresentationSlide); i {
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
		file_presentationSlide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresentationSlide_Notes); i {
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
			RawDescriptor: file_presentationSlide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_presentationSlide_proto_goTypes,
		DependencyIndexes: file_presentationSlide_proto_depIdxs,
		MessageInfos:      file_presentationSlide_proto_msgTypes,
	}.Build()
	File_presentationSlide_proto = out.File
	file_presentationSlide_proto_rawDesc = nil
	file_presentationSlide_proto_goTypes = nil
	file_presentationSlide_proto_depIdxs = nil
}
