// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: worldbuilder/entity/v1/entity_service.proto

package entityv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetEntityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEntityRequest) Reset() {
	*x = GetEntityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worldbuilder_entity_v1_entity_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntityRequest) ProtoMessage() {}

func (x *GetEntityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worldbuilder_entity_v1_entity_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntityRequest.ProtoReflect.Descriptor instead.
func (*GetEntityRequest) Descriptor() ([]byte, []int) {
	return file_worldbuilder_entity_v1_entity_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetEntityRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetEntityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *GetEntityResponse) Reset() {
	*x = GetEntityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worldbuilder_entity_v1_entity_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntityResponse) ProtoMessage() {}

func (x *GetEntityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worldbuilder_entity_v1_entity_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntityResponse.ProtoReflect.Descriptor instead.
func (*GetEntityResponse) Descriptor() ([]byte, []int) {
	return file_worldbuilder_entity_v1_entity_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetEntityResponse) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

type GetEntitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Criteria *Filter `protobuf:"bytes,1,opt,name=criteria,proto3" json:"criteria,omitempty"`
}

func (x *GetEntitiesRequest) Reset() {
	*x = GetEntitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worldbuilder_entity_v1_entity_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntitiesRequest) ProtoMessage() {}

func (x *GetEntitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worldbuilder_entity_v1_entity_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntitiesRequest.ProtoReflect.Descriptor instead.
func (*GetEntitiesRequest) Descriptor() ([]byte, []int) {
	return file_worldbuilder_entity_v1_entity_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetEntitiesRequest) GetCriteria() *Filter {
	if x != nil {
		return x.Criteria
	}
	return nil
}

type GetEntitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Criteria *Filter   `protobuf:"bytes,1,opt,name=criteria,proto3" json:"criteria,omitempty"`
	Entities []*Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *GetEntitiesResponse) Reset() {
	*x = GetEntitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worldbuilder_entity_v1_entity_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntitiesResponse) ProtoMessage() {}

func (x *GetEntitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worldbuilder_entity_v1_entity_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntitiesResponse.ProtoReflect.Descriptor instead.
func (*GetEntitiesResponse) Descriptor() ([]byte, []int) {
	return file_worldbuilder_entity_v1_entity_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetEntitiesResponse) GetCriteria() *Filter {
	if x != nil {
		return x.Criteria
	}
	return nil
}

func (x *GetEntitiesResponse) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

var File_worldbuilder_entity_v1_entity_service_proto protoreflect.FileDescriptor

var file_worldbuilder_entity_v1_entity_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x50, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x22, 0x8d,
	0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x12, 0x3a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x32, 0xdd,
	0x01, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x62, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x28, 0x2e,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xe8,
	0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x73, 0x61, 0x72, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x57, 0x45, 0x58, 0xaa, 0x02, 0x16, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x16, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5c,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x5c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x18, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_worldbuilder_entity_v1_entity_service_proto_rawDescOnce sync.Once
	file_worldbuilder_entity_v1_entity_service_proto_rawDescData = file_worldbuilder_entity_v1_entity_service_proto_rawDesc
)

func file_worldbuilder_entity_v1_entity_service_proto_rawDescGZIP() []byte {
	file_worldbuilder_entity_v1_entity_service_proto_rawDescOnce.Do(func() {
		file_worldbuilder_entity_v1_entity_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_worldbuilder_entity_v1_entity_service_proto_rawDescData)
	})
	return file_worldbuilder_entity_v1_entity_service_proto_rawDescData
}

var file_worldbuilder_entity_v1_entity_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_worldbuilder_entity_v1_entity_service_proto_goTypes = []interface{}{
	(*GetEntityRequest)(nil),    // 0: worldbuilder.entity.v1.GetEntityRequest
	(*GetEntityResponse)(nil),   // 1: worldbuilder.entity.v1.GetEntityResponse
	(*GetEntitiesRequest)(nil),  // 2: worldbuilder.entity.v1.GetEntitiesRequest
	(*GetEntitiesResponse)(nil), // 3: worldbuilder.entity.v1.GetEntitiesResponse
	(*Entity)(nil),              // 4: worldbuilder.entity.v1.Entity
	(*Filter)(nil),              // 5: worldbuilder.entity.v1.Filter
}
var file_worldbuilder_entity_v1_entity_service_proto_depIdxs = []int32{
	4, // 0: worldbuilder.entity.v1.GetEntityResponse.entity:type_name -> worldbuilder.entity.v1.Entity
	5, // 1: worldbuilder.entity.v1.GetEntitiesRequest.criteria:type_name -> worldbuilder.entity.v1.Filter
	5, // 2: worldbuilder.entity.v1.GetEntitiesResponse.criteria:type_name -> worldbuilder.entity.v1.Filter
	4, // 3: worldbuilder.entity.v1.GetEntitiesResponse.entities:type_name -> worldbuilder.entity.v1.Entity
	0, // 4: worldbuilder.entity.v1.EntityService.GetEntity:input_type -> worldbuilder.entity.v1.GetEntityRequest
	2, // 5: worldbuilder.entity.v1.EntityService.GetEntities:input_type -> worldbuilder.entity.v1.GetEntitiesRequest
	1, // 6: worldbuilder.entity.v1.EntityService.GetEntity:output_type -> worldbuilder.entity.v1.GetEntityResponse
	3, // 7: worldbuilder.entity.v1.EntityService.GetEntities:output_type -> worldbuilder.entity.v1.GetEntitiesResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_worldbuilder_entity_v1_entity_service_proto_init() }
func file_worldbuilder_entity_v1_entity_service_proto_init() {
	if File_worldbuilder_entity_v1_entity_service_proto != nil {
		return
	}
	file_worldbuilder_entity_v1_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_worldbuilder_entity_v1_entity_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntityRequest); i {
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
		file_worldbuilder_entity_v1_entity_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntityResponse); i {
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
		file_worldbuilder_entity_v1_entity_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntitiesRequest); i {
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
		file_worldbuilder_entity_v1_entity_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntitiesResponse); i {
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
			RawDescriptor: file_worldbuilder_entity_v1_entity_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worldbuilder_entity_v1_entity_service_proto_goTypes,
		DependencyIndexes: file_worldbuilder_entity_v1_entity_service_proto_depIdxs,
		MessageInfos:      file_worldbuilder_entity_v1_entity_service_proto_msgTypes,
	}.Build()
	File_worldbuilder_entity_v1_entity_service_proto = out.File
	file_worldbuilder_entity_v1_entity_service_proto_rawDesc = nil
	file_worldbuilder_entity_v1_entity_service_proto_goTypes = nil
	file_worldbuilder_entity_v1_entity_service_proto_depIdxs = nil
}
