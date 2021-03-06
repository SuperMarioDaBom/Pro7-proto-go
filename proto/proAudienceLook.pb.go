// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: proAudienceLook.proto

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

type ProAudienceLook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid             *UUID                            `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name             string                           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ScreenLooks      []*ProAudienceLook_ProScreenLook `protobuf:"bytes,3,rep,name=screen_looks,json=screenLooks,proto3" json:"screen_looks,omitempty"`
	OriginalLookUuid *UUID                            `protobuf:"bytes,4,opt,name=original_look_uuid,json=originalLookUuid,proto3" json:"original_look_uuid,omitempty"`
}

func (x *ProAudienceLook) Reset() {
	*x = ProAudienceLook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proAudienceLook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProAudienceLook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProAudienceLook) ProtoMessage() {}

func (x *ProAudienceLook) ProtoReflect() protoreflect.Message {
	mi := &file_proAudienceLook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProAudienceLook.ProtoReflect.Descriptor instead.
func (*ProAudienceLook) Descriptor() ([]byte, []int) {
	return file_proAudienceLook_proto_rawDescGZIP(), []int{0}
}

func (x *ProAudienceLook) GetUuid() *UUID {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *ProAudienceLook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProAudienceLook) GetScreenLooks() []*ProAudienceLook_ProScreenLook {
	if x != nil {
		return x.ScreenLooks
	}
	return nil
}

func (x *ProAudienceLook) GetOriginalLookUuid() *UUID {
	if x != nil {
		return x.OriginalLookUuid
	}
	return nil
}

type ProAudienceLook_ProScreenLook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProScreenUuid                 *UUID `protobuf:"bytes,1,opt,name=pro_screen_uuid,json=proScreenUuid,proto3" json:"pro_screen_uuid,omitempty"`
	PropsEnabled                  bool  `protobuf:"varint,2,opt,name=props_enabled,json=propsEnabled,proto3" json:"props_enabled,omitempty"`
	LiveVideoEnabled              bool  `protobuf:"varint,3,opt,name=live_video_enabled,json=liveVideoEnabled,proto3" json:"live_video_enabled,omitempty"`
	PresentationBackgroundEnabled bool  `protobuf:"varint,4,opt,name=presentation_background_enabled,json=presentationBackgroundEnabled,proto3" json:"presentation_background_enabled,omitempty"`
	TemplateDocumentFilePath      *URL  `protobuf:"bytes,5,opt,name=template_document_file_path,json=templateDocumentFilePath,proto3" json:"template_document_file_path,omitempty"`
	TemplateSlideUuid             *UUID `protobuf:"bytes,6,opt,name=template_slide_uuid,json=templateSlideUuid,proto3" json:"template_slide_uuid,omitempty"`
	PresentationForegroundEnabled bool  `protobuf:"varint,7,opt,name=presentation_foreground_enabled,json=presentationForegroundEnabled,proto3" json:"presentation_foreground_enabled,omitempty"`
	MaskUuid                      *UUID `protobuf:"bytes,8,opt,name=mask_uuid,json=maskUuid,proto3" json:"mask_uuid,omitempty"`
	AnnouncementsEnabled          bool  `protobuf:"varint,9,opt,name=announcements_enabled,json=announcementsEnabled,proto3" json:"announcements_enabled,omitempty"`
	PropsLayerEnabled             bool  `protobuf:"varint,10,opt,name=props_layer_enabled,json=propsLayerEnabled,proto3" json:"props_layer_enabled,omitempty"`
	MessagesLayerEnabled          bool  `protobuf:"varint,11,opt,name=messages_layer_enabled,json=messagesLayerEnabled,proto3" json:"messages_layer_enabled,omitempty"`
}

func (x *ProAudienceLook_ProScreenLook) Reset() {
	*x = ProAudienceLook_ProScreenLook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proAudienceLook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProAudienceLook_ProScreenLook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProAudienceLook_ProScreenLook) ProtoMessage() {}

func (x *ProAudienceLook_ProScreenLook) ProtoReflect() protoreflect.Message {
	mi := &file_proAudienceLook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProAudienceLook_ProScreenLook.ProtoReflect.Descriptor instead.
func (*ProAudienceLook_ProScreenLook) Descriptor() ([]byte, []int) {
	return file_proAudienceLook_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProAudienceLook_ProScreenLook) GetProScreenUuid() *UUID {
	if x != nil {
		return x.ProScreenUuid
	}
	return nil
}

func (x *ProAudienceLook_ProScreenLook) GetPropsEnabled() bool {
	if x != nil {
		return x.PropsEnabled
	}
	return false
}

func (x *ProAudienceLook_ProScreenLook) GetLiveVideoEnabled() bool {
	if x != nil {
		return x.LiveVideoEnabled
	}
	return false
}

func (x *ProAudienceLook_ProScreenLook) GetPresentationBackgroundEnabled() bool {
	if x != nil {
		return x.PresentationBackgroundEnabled
	}
	return false
}

func (x *ProAudienceLook_ProScreenLook) GetTemplateDocumentFilePath() *URL {
	if x != nil {
		return x.TemplateDocumentFilePath
	}
	return nil
}

func (x *ProAudienceLook_ProScreenLook) GetTemplateSlideUuid() *UUID {
	if x != nil {
		return x.TemplateSlideUuid
	}
	return nil
}

func (x *ProAudienceLook_ProScreenLook) GetPresentationForegroundEnabled() bool {
	if x != nil {
		return x.PresentationForegroundEnabled
	}
	return false
}

func (x *ProAudienceLook_ProScreenLook) GetMaskUuid() *UUID {
	if x != nil {
		return x.MaskUuid
	}
	return nil
}

func (x *ProAudienceLook_ProScreenLook) GetAnnouncementsEnabled() bool {
	if x != nil {
		return x.AnnouncementsEnabled
	}
	return false
}

func (x *ProAudienceLook_ProScreenLook) GetPropsLayerEnabled() bool {
	if x != nil {
		return x.PropsLayerEnabled
	}
	return false
}

func (x *ProAudienceLook_ProScreenLook) GetMessagesLayerEnabled() bool {
	if x != nil {
		return x.MessagesLayerEnabled
	}
	return false
}

var File_proAudienceLook_proto protoreflect.FileDescriptor

var file_proAudienceLook_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x10, 0x62, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcf, 0x06, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a,
	0x0c, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x72,
	0x6f, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x2e, 0x50, 0x72,
	0x6f, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x4c, 0x6f, 0x6f, 0x6b, 0x52, 0x0b, 0x73, 0x63, 0x72,
	0x65, 0x65, 0x6e, 0x4c, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x3b, 0x0a, 0x12, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x6f, 0x6f,
	0x6b, 0x55, 0x75, 0x69, 0x64, 0x1a, 0xfc, 0x04, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x53, 0x63, 0x72,
	0x65, 0x65, 0x6e, 0x4c, 0x6f, 0x6f, 0x6b, 0x12, 0x35, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x5f, 0x73,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52,
	0x0d, 0x70, 0x72, 0x6f, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x6c, 0x69, 0x76, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x46, 0x0a, 0x1f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x4b, 0x0a, 0x1b, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x18, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3d, 0x0a, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x69, 0x64,
	0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x46, 0x0a, 0x1f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x65,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2a, 0x0a,
	0x09, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x72, 0x76, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52,
	0x08, 0x6d, 0x61, 0x73, 0x6b, 0x55, 0x75, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x15, 0x61, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2e,
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x70, 0x72, 0x6f,
	0x70, 0x73, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x34,
	0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x50, 0x72, 0x6f, 0x37, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proAudienceLook_proto_rawDescOnce sync.Once
	file_proAudienceLook_proto_rawDescData = file_proAudienceLook_proto_rawDesc
)

func file_proAudienceLook_proto_rawDescGZIP() []byte {
	file_proAudienceLook_proto_rawDescOnce.Do(func() {
		file_proAudienceLook_proto_rawDescData = protoimpl.X.CompressGZIP(file_proAudienceLook_proto_rawDescData)
	})
	return file_proAudienceLook_proto_rawDescData
}

var file_proAudienceLook_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proAudienceLook_proto_goTypes = []interface{}{
	(*ProAudienceLook)(nil),               // 0: rv.data.ProAudienceLook
	(*ProAudienceLook_ProScreenLook)(nil), // 1: rv.data.ProAudienceLook.ProScreenLook
	(*UUID)(nil),                          // 2: rv.data.UUID
	(*URL)(nil),                           // 3: rv.data.URL
}
var file_proAudienceLook_proto_depIdxs = []int32{
	2, // 0: rv.data.ProAudienceLook.uuid:type_name -> rv.data.UUID
	1, // 1: rv.data.ProAudienceLook.screen_looks:type_name -> rv.data.ProAudienceLook.ProScreenLook
	2, // 2: rv.data.ProAudienceLook.original_look_uuid:type_name -> rv.data.UUID
	2, // 3: rv.data.ProAudienceLook.ProScreenLook.pro_screen_uuid:type_name -> rv.data.UUID
	3, // 4: rv.data.ProAudienceLook.ProScreenLook.template_document_file_path:type_name -> rv.data.URL
	2, // 5: rv.data.ProAudienceLook.ProScreenLook.template_slide_uuid:type_name -> rv.data.UUID
	2, // 6: rv.data.ProAudienceLook.ProScreenLook.mask_uuid:type_name -> rv.data.UUID
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proAudienceLook_proto_init() }
func file_proAudienceLook_proto_init() {
	if File_proAudienceLook_proto != nil {
		return
	}
	file_basicTypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proAudienceLook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProAudienceLook); i {
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
		file_proAudienceLook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProAudienceLook_ProScreenLook); i {
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
			RawDescriptor: file_proAudienceLook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proAudienceLook_proto_goTypes,
		DependencyIndexes: file_proAudienceLook_proto_depIdxs,
		MessageInfos:      file_proAudienceLook_proto_msgTypes,
	}.Build()
	File_proAudienceLook_proto = out.File
	file_proAudienceLook_proto_rawDesc = nil
	file_proAudienceLook_proto_goTypes = nil
	file_proAudienceLook_proto_depIdxs = nil
}
