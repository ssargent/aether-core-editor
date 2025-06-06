// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: game/v1/capability_service.proto

package gamev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Capability
type Capability struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Id                 int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WorldId            int64                  `protobuf:"varint,2,opt,name=world_id,json=worldId,proto3" json:"world_id,omitempty"`
	ParentId           *int64                 `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	CapabilityType     string                 `protobuf:"bytes,4,opt,name=capability_type,json=capabilityType,proto3" json:"capability_type,omitempty"`
	Code               string                 `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	Name               string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Description        string                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Requirements       *structpb.Struct       `protobuf:"bytes,8,opt,name=requirements,proto3" json:"requirements,omitempty"`
	Actions            *structpb.Struct       `protobuf:"bytes,9,opt,name=actions,proto3,oneof" json:"actions,omitempty"`
	AccessRequirements *structpb.Struct       `protobuf:"bytes,10,opt,name=access_requirements,json=accessRequirements,proto3" json:"access_requirements,omitempty"`
	Tags               []string               `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	GameplayDefinition *structpb.Struct       `protobuf:"bytes,12,opt,name=gameplay_definition,json=gameplayDefinition,proto3,oneof" json:"gameplay_definition,omitempty"`
	Metadata           *Metadata              `protobuf:"bytes,13,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Capability) Reset() {
	*x = Capability{}
	mi := &file_game_v1_capability_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Capability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capability) ProtoMessage() {}

func (x *Capability) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_capability_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capability.ProtoReflect.Descriptor instead.
func (*Capability) Descriptor() ([]byte, []int) {
	return file_game_v1_capability_service_proto_rawDescGZIP(), []int{0}
}

func (x *Capability) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Capability) GetWorldId() int64 {
	if x != nil {
		return x.WorldId
	}
	return 0
}

func (x *Capability) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *Capability) GetCapabilityType() string {
	if x != nil {
		return x.CapabilityType
	}
	return ""
}

func (x *Capability) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Capability) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Capability) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Capability) GetRequirements() *structpb.Struct {
	if x != nil {
		return x.Requirements
	}
	return nil
}

func (x *Capability) GetActions() *structpb.Struct {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Capability) GetAccessRequirements() *structpb.Struct {
	if x != nil {
		return x.AccessRequirements
	}
	return nil
}

func (x *Capability) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Capability) GetGameplayDefinition() *structpb.Struct {
	if x != nil {
		return x.GameplayDefinition
	}
	return nil
}

func (x *Capability) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type CreateCapabilityRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	WorldId            int64                  `protobuf:"varint,1,opt,name=world_id,json=worldId,proto3" json:"world_id,omitempty"`
	ParentId           *int64                 `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	CapabilityType     string                 `protobuf:"bytes,3,opt,name=capability_type,json=capabilityType,proto3" json:"capability_type,omitempty"`
	Code               string                 `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Name               string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Description        string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Requirements       *structpb.Struct       `protobuf:"bytes,7,opt,name=requirements,proto3" json:"requirements,omitempty"`
	Actions            *structpb.Struct       `protobuf:"bytes,8,opt,name=actions,proto3,oneof" json:"actions,omitempty"`
	AccessRequirements *structpb.Struct       `protobuf:"bytes,9,opt,name=access_requirements,json=accessRequirements,proto3" json:"access_requirements,omitempty"`
	Tags               []string               `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	GameplayDefinition *structpb.Struct       `protobuf:"bytes,11,opt,name=gameplay_definition,json=gameplayDefinition,proto3,oneof" json:"gameplay_definition,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *CreateCapabilityRequest) Reset() {
	*x = CreateCapabilityRequest{}
	mi := &file_game_v1_capability_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCapabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCapabilityRequest) ProtoMessage() {}

func (x *CreateCapabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_capability_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCapabilityRequest.ProtoReflect.Descriptor instead.
func (*CreateCapabilityRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_capability_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCapabilityRequest) GetWorldId() int64 {
	if x != nil {
		return x.WorldId
	}
	return 0
}

func (x *CreateCapabilityRequest) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *CreateCapabilityRequest) GetCapabilityType() string {
	if x != nil {
		return x.CapabilityType
	}
	return ""
}

func (x *CreateCapabilityRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateCapabilityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCapabilityRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCapabilityRequest) GetRequirements() *structpb.Struct {
	if x != nil {
		return x.Requirements
	}
	return nil
}

func (x *CreateCapabilityRequest) GetActions() *structpb.Struct {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *CreateCapabilityRequest) GetAccessRequirements() *structpb.Struct {
	if x != nil {
		return x.AccessRequirements
	}
	return nil
}

func (x *CreateCapabilityRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateCapabilityRequest) GetGameplayDefinition() *structpb.Struct {
	if x != nil {
		return x.GameplayDefinition
	}
	return nil
}

type UpdateCapabilityRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Id                 int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId           *int64                 `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	CapabilityType     *string                `protobuf:"bytes,3,opt,name=capability_type,json=capabilityType,proto3,oneof" json:"capability_type,omitempty"`
	Code               *string                `protobuf:"bytes,4,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Name               *string                `protobuf:"bytes,5,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description        *string                `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	Requirements       *structpb.Struct       `protobuf:"bytes,7,opt,name=requirements,proto3,oneof" json:"requirements,omitempty"`
	Actions            *structpb.Struct       `protobuf:"bytes,8,opt,name=actions,proto3,oneof" json:"actions,omitempty"`
	AccessRequirements *structpb.Struct       `protobuf:"bytes,9,opt,name=access_requirements,json=accessRequirements,proto3,oneof" json:"access_requirements,omitempty"`
	Tags               []string               `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	GameplayDefinition *structpb.Struct       `protobuf:"bytes,11,opt,name=gameplay_definition,json=gameplayDefinition,proto3,oneof" json:"gameplay_definition,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *UpdateCapabilityRequest) Reset() {
	*x = UpdateCapabilityRequest{}
	mi := &file_game_v1_capability_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCapabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCapabilityRequest) ProtoMessage() {}

func (x *UpdateCapabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_capability_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCapabilityRequest.ProtoReflect.Descriptor instead.
func (*UpdateCapabilityRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_capability_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCapabilityRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCapabilityRequest) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *UpdateCapabilityRequest) GetCapabilityType() string {
	if x != nil && x.CapabilityType != nil {
		return *x.CapabilityType
	}
	return ""
}

func (x *UpdateCapabilityRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *UpdateCapabilityRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateCapabilityRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdateCapabilityRequest) GetRequirements() *structpb.Struct {
	if x != nil {
		return x.Requirements
	}
	return nil
}

func (x *UpdateCapabilityRequest) GetActions() *structpb.Struct {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *UpdateCapabilityRequest) GetAccessRequirements() *structpb.Struct {
	if x != nil {
		return x.AccessRequirements
	}
	return nil
}

func (x *UpdateCapabilityRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdateCapabilityRequest) GetGameplayDefinition() *structpb.Struct {
	if x != nil {
		return x.GameplayDefinition
	}
	return nil
}

type GetCapabilityRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCapabilityRequest) Reset() {
	*x = GetCapabilityRequest{}
	mi := &file_game_v1_capability_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCapabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCapabilityRequest) ProtoMessage() {}

func (x *GetCapabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_capability_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCapabilityRequest.ProtoReflect.Descriptor instead.
func (*GetCapabilityRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_capability_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCapabilityRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCapabilityRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCapabilityRequest) Reset() {
	*x = DeleteCapabilityRequest{}
	mi := &file_game_v1_capability_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCapabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCapabilityRequest) ProtoMessage() {}

func (x *DeleteCapabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_capability_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCapabilityRequest.ProtoReflect.Descriptor instead.
func (*DeleteCapabilityRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_capability_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCapabilityRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListCapabilitiesRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	WorldId        int64                  `protobuf:"varint,1,opt,name=world_id,json=worldId,proto3" json:"world_id,omitempty"`
	CapabilityType string                 `protobuf:"bytes,2,opt,name=capability_type,json=capabilityType,proto3" json:"capability_type,omitempty"`
	Code           string                 `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Name           string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ParentId       int64                  `protobuf:"varint,5,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Tags           []string               `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListCapabilitiesRequest) Reset() {
	*x = ListCapabilitiesRequest{}
	mi := &file_game_v1_capability_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCapabilitiesRequest) ProtoMessage() {}

func (x *ListCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_capability_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCapabilitiesRequest.ProtoReflect.Descriptor instead.
func (*ListCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_game_v1_capability_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListCapabilitiesRequest) GetWorldId() int64 {
	if x != nil {
		return x.WorldId
	}
	return 0
}

func (x *ListCapabilitiesRequest) GetCapabilityType() string {
	if x != nil {
		return x.CapabilityType
	}
	return ""
}

func (x *ListCapabilitiesRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ListCapabilitiesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListCapabilitiesRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *ListCapabilitiesRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CapabilityResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Capability    *Capability            `protobuf:"bytes,1,opt,name=capability,proto3" json:"capability,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CapabilityResponse) Reset() {
	*x = CapabilityResponse{}
	mi := &file_game_v1_capability_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CapabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilityResponse) ProtoMessage() {}

func (x *CapabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_capability_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilityResponse.ProtoReflect.Descriptor instead.
func (*CapabilityResponse) Descriptor() ([]byte, []int) {
	return file_game_v1_capability_service_proto_rawDescGZIP(), []int{6}
}

func (x *CapabilityResponse) GetCapability() *Capability {
	if x != nil {
		return x.Capability
	}
	return nil
}

type ListCapabilitiesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Capabilities  []*Capability          `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	Pagination    *PaginationResponse    `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCapabilitiesResponse) Reset() {
	*x = ListCapabilitiesResponse{}
	mi := &file_game_v1_capability_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCapabilitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCapabilitiesResponse) ProtoMessage() {}

func (x *ListCapabilitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_capability_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCapabilitiesResponse.ProtoReflect.Descriptor instead.
func (*ListCapabilitiesResponse) Descriptor() ([]byte, []int) {
	return file_game_v1_capability_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListCapabilitiesResponse) GetCapabilities() []*Capability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *ListCapabilitiesResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// Capability filter for advanced querying
type CapabilityFilter struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	WorldId        *int64                 `protobuf:"varint,1,opt,name=world_id,json=worldId,proto3,oneof" json:"world_id,omitempty"`
	CapabilityType *string                `protobuf:"bytes,2,opt,name=capability_type,json=capabilityType,proto3,oneof" json:"capability_type,omitempty"`
	Code           *string                `protobuf:"bytes,3,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Name           *string                `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	ParentId       *int64                 `protobuf:"varint,5,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	Tags           []string               `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CapabilityFilter) Reset() {
	*x = CapabilityFilter{}
	mi := &file_game_v1_capability_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CapabilityFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilityFilter) ProtoMessage() {}

func (x *CapabilityFilter) ProtoReflect() protoreflect.Message {
	mi := &file_game_v1_capability_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilityFilter.ProtoReflect.Descriptor instead.
func (*CapabilityFilter) Descriptor() ([]byte, []int) {
	return file_game_v1_capability_service_proto_rawDescGZIP(), []int{8}
}

func (x *CapabilityFilter) GetWorldId() int64 {
	if x != nil && x.WorldId != nil {
		return *x.WorldId
	}
	return 0
}

func (x *CapabilityFilter) GetCapabilityType() string {
	if x != nil && x.CapabilityType != nil {
		return *x.CapabilityType
	}
	return ""
}

func (x *CapabilityFilter) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *CapabilityFilter) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CapabilityFilter) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *CapabilityFilter) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_game_v1_capability_service_proto protoreflect.FileDescriptor

const file_game_v1_capability_service_proto_rawDesc = "" +
	"\n" +
	" game/v1/capability_service.proto\x12\agame.v1\x1a\x14game/v1/common.proto\x1a\x1cgoogle/protobuf/struct.proto\"\xcf\x04\n" +
	"\n" +
	"Capability\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x19\n" +
	"\bworld_id\x18\x02 \x01(\x03R\aworldId\x12 \n" +
	"\tparent_id\x18\x03 \x01(\x03H\x00R\bparentId\x88\x01\x01\x12'\n" +
	"\x0fcapability_type\x18\x04 \x01(\tR\x0ecapabilityType\x12\x12\n" +
	"\x04code\x18\x05 \x01(\tR\x04code\x12\x12\n" +
	"\x04name\x18\x06 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\a \x01(\tR\vdescription\x12;\n" +
	"\frequirements\x18\b \x01(\v2\x17.google.protobuf.StructR\frequirements\x126\n" +
	"\aactions\x18\t \x01(\v2\x17.google.protobuf.StructH\x01R\aactions\x88\x01\x01\x12H\n" +
	"\x13access_requirements\x18\n" +
	" \x01(\v2\x17.google.protobuf.StructR\x12accessRequirements\x12\x12\n" +
	"\x04tags\x18\v \x03(\tR\x04tags\x12M\n" +
	"\x13gameplay_definition\x18\f \x01(\v2\x17.google.protobuf.StructH\x02R\x12gameplayDefinition\x88\x01\x01\x12-\n" +
	"\bmetadata\x18\r \x01(\v2\x11.game.v1.MetadataR\bmetadataB\f\n" +
	"\n" +
	"_parent_idB\n" +
	"\n" +
	"\b_actionsB\x16\n" +
	"\x14_gameplay_definition\"\x9d\x04\n" +
	"\x17CreateCapabilityRequest\x12\x19\n" +
	"\bworld_id\x18\x01 \x01(\x03R\aworldId\x12 \n" +
	"\tparent_id\x18\x02 \x01(\x03H\x00R\bparentId\x88\x01\x01\x12'\n" +
	"\x0fcapability_type\x18\x03 \x01(\tR\x0ecapabilityType\x12\x12\n" +
	"\x04code\x18\x04 \x01(\tR\x04code\x12\x12\n" +
	"\x04name\x18\x05 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\x12;\n" +
	"\frequirements\x18\a \x01(\v2\x17.google.protobuf.StructR\frequirements\x126\n" +
	"\aactions\x18\b \x01(\v2\x17.google.protobuf.StructH\x01R\aactions\x88\x01\x01\x12H\n" +
	"\x13access_requirements\x18\t \x01(\v2\x17.google.protobuf.StructR\x12accessRequirements\x12\x12\n" +
	"\x04tags\x18\n" +
	" \x03(\tR\x04tags\x12M\n" +
	"\x13gameplay_definition\x18\v \x01(\v2\x17.google.protobuf.StructH\x02R\x12gameplayDefinition\x88\x01\x01B\f\n" +
	"\n" +
	"_parent_idB\n" +
	"\n" +
	"\b_actionsB\x16\n" +
	"\x14_gameplay_definition\"\x8f\x05\n" +
	"\x17UpdateCapabilityRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12 \n" +
	"\tparent_id\x18\x02 \x01(\x03H\x00R\bparentId\x88\x01\x01\x12,\n" +
	"\x0fcapability_type\x18\x03 \x01(\tH\x01R\x0ecapabilityType\x88\x01\x01\x12\x17\n" +
	"\x04code\x18\x04 \x01(\tH\x02R\x04code\x88\x01\x01\x12\x17\n" +
	"\x04name\x18\x05 \x01(\tH\x03R\x04name\x88\x01\x01\x12%\n" +
	"\vdescription\x18\x06 \x01(\tH\x04R\vdescription\x88\x01\x01\x12@\n" +
	"\frequirements\x18\a \x01(\v2\x17.google.protobuf.StructH\x05R\frequirements\x88\x01\x01\x126\n" +
	"\aactions\x18\b \x01(\v2\x17.google.protobuf.StructH\x06R\aactions\x88\x01\x01\x12M\n" +
	"\x13access_requirements\x18\t \x01(\v2\x17.google.protobuf.StructH\aR\x12accessRequirements\x88\x01\x01\x12\x12\n" +
	"\x04tags\x18\n" +
	" \x03(\tR\x04tags\x12M\n" +
	"\x13gameplay_definition\x18\v \x01(\v2\x17.google.protobuf.StructH\bR\x12gameplayDefinition\x88\x01\x01B\f\n" +
	"\n" +
	"_parent_idB\x12\n" +
	"\x10_capability_typeB\a\n" +
	"\x05_codeB\a\n" +
	"\x05_nameB\x0e\n" +
	"\f_descriptionB\x0f\n" +
	"\r_requirementsB\n" +
	"\n" +
	"\b_actionsB\x16\n" +
	"\x14_access_requirementsB\x16\n" +
	"\x14_gameplay_definition\"&\n" +
	"\x14GetCapabilityRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\")\n" +
	"\x17DeleteCapabilityRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"\xb6\x01\n" +
	"\x17ListCapabilitiesRequest\x12\x19\n" +
	"\bworld_id\x18\x01 \x01(\x03R\aworldId\x12'\n" +
	"\x0fcapability_type\x18\x02 \x01(\tR\x0ecapabilityType\x12\x12\n" +
	"\x04code\x18\x03 \x01(\tR\x04code\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12\x1b\n" +
	"\tparent_id\x18\x05 \x01(\x03R\bparentId\x12\x12\n" +
	"\x04tags\x18\x06 \x03(\tR\x04tags\"I\n" +
	"\x12CapabilityResponse\x123\n" +
	"\n" +
	"capability\x18\x01 \x01(\v2\x13.game.v1.CapabilityR\n" +
	"capability\"\x90\x01\n" +
	"\x18ListCapabilitiesResponse\x127\n" +
	"\fcapabilities\x18\x01 \x03(\v2\x13.game.v1.CapabilityR\fcapabilities\x12;\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2\x1b.game.v1.PaginationResponseR\n" +
	"pagination\"\x89\x02\n" +
	"\x10CapabilityFilter\x12\x1e\n" +
	"\bworld_id\x18\x01 \x01(\x03H\x00R\aworldId\x88\x01\x01\x12,\n" +
	"\x0fcapability_type\x18\x02 \x01(\tH\x01R\x0ecapabilityType\x88\x01\x01\x12\x17\n" +
	"\x04code\x18\x03 \x01(\tH\x02R\x04code\x88\x01\x01\x12\x17\n" +
	"\x04name\x18\x04 \x01(\tH\x03R\x04name\x88\x01\x01\x12 \n" +
	"\tparent_id\x18\x05 \x01(\x03H\x04R\bparentId\x88\x01\x01\x12\x12\n" +
	"\x04tags\x18\x06 \x03(\tR\x04tagsB\v\n" +
	"\t_world_idB\x12\n" +
	"\x10_capability_typeB\a\n" +
	"\x05_codeB\a\n" +
	"\x05_nameB\f\n" +
	"\n" +
	"_parent_id2\xae\x03\n" +
	"\x11CapabilityService\x12Q\n" +
	"\x10CreateCapability\x12 .game.v1.CreateCapabilityRequest\x1a\x1b.game.v1.CapabilityResponse\x12K\n" +
	"\rGetCapability\x12\x1d.game.v1.GetCapabilityRequest\x1a\x1b.game.v1.CapabilityResponse\x12Q\n" +
	"\x10UpdateCapability\x12 .game.v1.UpdateCapabilityRequest\x1a\x1b.game.v1.CapabilityResponse\x12M\n" +
	"\x10DeleteCapability\x12 .game.v1.DeleteCapabilityRequest\x1a\x17.game.v1.DeleteResponse\x12W\n" +
	"\x10ListCapabilities\x12 .game.v1.ListCapabilitiesRequest\x1a!.game.v1.ListCapabilitiesResponseB\x9d\x01\n" +
	"\vcom.game.v1B\x16CapabilityServiceProtoP\x01Z9github.com/ssargent/aether-core-editor/gen/game/v1;gamev1\xa2\x02\x03GXX\xaa\x02\aGame.V1\xca\x02\aGame\\V1\xe2\x02\x13Game\\V1\\GPBMetadata\xea\x02\bGame::V1b\x06proto3"

var (
	file_game_v1_capability_service_proto_rawDescOnce sync.Once
	file_game_v1_capability_service_proto_rawDescData []byte
)

func file_game_v1_capability_service_proto_rawDescGZIP() []byte {
	file_game_v1_capability_service_proto_rawDescOnce.Do(func() {
		file_game_v1_capability_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_game_v1_capability_service_proto_rawDesc), len(file_game_v1_capability_service_proto_rawDesc)))
	})
	return file_game_v1_capability_service_proto_rawDescData
}

var file_game_v1_capability_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_game_v1_capability_service_proto_goTypes = []any{
	(*Capability)(nil),               // 0: game.v1.Capability
	(*CreateCapabilityRequest)(nil),  // 1: game.v1.CreateCapabilityRequest
	(*UpdateCapabilityRequest)(nil),  // 2: game.v1.UpdateCapabilityRequest
	(*GetCapabilityRequest)(nil),     // 3: game.v1.GetCapabilityRequest
	(*DeleteCapabilityRequest)(nil),  // 4: game.v1.DeleteCapabilityRequest
	(*ListCapabilitiesRequest)(nil),  // 5: game.v1.ListCapabilitiesRequest
	(*CapabilityResponse)(nil),       // 6: game.v1.CapabilityResponse
	(*ListCapabilitiesResponse)(nil), // 7: game.v1.ListCapabilitiesResponse
	(*CapabilityFilter)(nil),         // 8: game.v1.CapabilityFilter
	(*structpb.Struct)(nil),          // 9: google.protobuf.Struct
	(*Metadata)(nil),                 // 10: game.v1.Metadata
	(*PaginationResponse)(nil),       // 11: game.v1.PaginationResponse
	(*DeleteResponse)(nil),           // 12: game.v1.DeleteResponse
}
var file_game_v1_capability_service_proto_depIdxs = []int32{
	9,  // 0: game.v1.Capability.requirements:type_name -> google.protobuf.Struct
	9,  // 1: game.v1.Capability.actions:type_name -> google.protobuf.Struct
	9,  // 2: game.v1.Capability.access_requirements:type_name -> google.protobuf.Struct
	9,  // 3: game.v1.Capability.gameplay_definition:type_name -> google.protobuf.Struct
	10, // 4: game.v1.Capability.metadata:type_name -> game.v1.Metadata
	9,  // 5: game.v1.CreateCapabilityRequest.requirements:type_name -> google.protobuf.Struct
	9,  // 6: game.v1.CreateCapabilityRequest.actions:type_name -> google.protobuf.Struct
	9,  // 7: game.v1.CreateCapabilityRequest.access_requirements:type_name -> google.protobuf.Struct
	9,  // 8: game.v1.CreateCapabilityRequest.gameplay_definition:type_name -> google.protobuf.Struct
	9,  // 9: game.v1.UpdateCapabilityRequest.requirements:type_name -> google.protobuf.Struct
	9,  // 10: game.v1.UpdateCapabilityRequest.actions:type_name -> google.protobuf.Struct
	9,  // 11: game.v1.UpdateCapabilityRequest.access_requirements:type_name -> google.protobuf.Struct
	9,  // 12: game.v1.UpdateCapabilityRequest.gameplay_definition:type_name -> google.protobuf.Struct
	0,  // 13: game.v1.CapabilityResponse.capability:type_name -> game.v1.Capability
	0,  // 14: game.v1.ListCapabilitiesResponse.capabilities:type_name -> game.v1.Capability
	11, // 15: game.v1.ListCapabilitiesResponse.pagination:type_name -> game.v1.PaginationResponse
	1,  // 16: game.v1.CapabilityService.CreateCapability:input_type -> game.v1.CreateCapabilityRequest
	3,  // 17: game.v1.CapabilityService.GetCapability:input_type -> game.v1.GetCapabilityRequest
	2,  // 18: game.v1.CapabilityService.UpdateCapability:input_type -> game.v1.UpdateCapabilityRequest
	4,  // 19: game.v1.CapabilityService.DeleteCapability:input_type -> game.v1.DeleteCapabilityRequest
	5,  // 20: game.v1.CapabilityService.ListCapabilities:input_type -> game.v1.ListCapabilitiesRequest
	6,  // 21: game.v1.CapabilityService.CreateCapability:output_type -> game.v1.CapabilityResponse
	6,  // 22: game.v1.CapabilityService.GetCapability:output_type -> game.v1.CapabilityResponse
	6,  // 23: game.v1.CapabilityService.UpdateCapability:output_type -> game.v1.CapabilityResponse
	12, // 24: game.v1.CapabilityService.DeleteCapability:output_type -> game.v1.DeleteResponse
	7,  // 25: game.v1.CapabilityService.ListCapabilities:output_type -> game.v1.ListCapabilitiesResponse
	21, // [21:26] is the sub-list for method output_type
	16, // [16:21] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_game_v1_capability_service_proto_init() }
func file_game_v1_capability_service_proto_init() {
	if File_game_v1_capability_service_proto != nil {
		return
	}
	file_game_v1_common_proto_init()
	file_game_v1_capability_service_proto_msgTypes[0].OneofWrappers = []any{}
	file_game_v1_capability_service_proto_msgTypes[1].OneofWrappers = []any{}
	file_game_v1_capability_service_proto_msgTypes[2].OneofWrappers = []any{}
	file_game_v1_capability_service_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_game_v1_capability_service_proto_rawDesc), len(file_game_v1_capability_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_v1_capability_service_proto_goTypes,
		DependencyIndexes: file_game_v1_capability_service_proto_depIdxs,
		MessageInfos:      file_game_v1_capability_service_proto_msgTypes,
	}.Build()
	File_game_v1_capability_service_proto = out.File
	file_game_v1_capability_service_proto_goTypes = nil
	file_game_v1_capability_service_proto_depIdxs = nil
}
