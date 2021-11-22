// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: api/v1/schedulers.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// List scheduler request options.
type ListSchedulersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSchedulersRequest) Reset() {
	*x = ListSchedulersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schedulers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSchedulersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSchedulersRequest) ProtoMessage() {}

func (x *ListSchedulersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schedulers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSchedulersRequest.ProtoReflect.Descriptor instead.
func (*ListSchedulersRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_schedulers_proto_rawDescGZIP(), []int{0}
}

// The list schedulers reponse message.
type ListSchedulersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of schedulers fetched.
	Schedulers []*ListedScheduler `protobuf:"bytes,1,rep,name=schedulers,proto3" json:"schedulers,omitempty"`
}

func (x *ListSchedulersResponse) Reset() {
	*x = ListSchedulersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schedulers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSchedulersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSchedulersResponse) ProtoMessage() {}

func (x *ListSchedulersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schedulers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSchedulersResponse.ProtoReflect.Descriptor instead.
func (*ListSchedulersResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_schedulers_proto_rawDescGZIP(), []int{1}
}

func (x *ListSchedulersResponse) GetSchedulers() []*ListedScheduler {
	if x != nil {
		return x.Schedulers
	}
	return nil
}

// Response for the create scheduler.
type CreateSchedulerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Scheduler that was created.
	Scheduler *Scheduler `protobuf:"bytes,1,opt,name=scheduler,proto3" json:"scheduler,omitempty"`
}

func (x *CreateSchedulerResponse) Reset() {
	*x = CreateSchedulerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schedulers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSchedulerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSchedulerResponse) ProtoMessage() {}

func (x *CreateSchedulerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schedulers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSchedulerResponse.ProtoReflect.Descriptor instead.
func (*CreateSchedulerResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_schedulers_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSchedulerResponse) GetScheduler() *Scheduler {
	if x != nil {
		return x.Scheduler
	}
	return nil
}

// Scheduler is the struct that defines a maestro scheduler.
type CreateSchedulerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the scheduler.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Game the new scheduler will be part of.
	Game string `protobuf:"bytes,2,opt,name=game,proto3" json:"game,omitempty"`
	// ?.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// The game room termination grace period.
	TerminationGracePeriod int64 `protobuf:"varint,4,opt,name=termination_grace_period,json=terminationGracePeriod,proto3" json:"termination_grace_period,omitempty"`
	// The container object array defines all the game room container configurations.
	Containers []*Container `protobuf:"bytes,5,rep,name=containers,proto3" json:"containers,omitempty"`
	// The port range object describes what is the port range used to allocate game rooms.
	PortRange *PortRange `protobuf:"bytes,6,opt,name=port_range,json=portRange,proto3" json:"port_range,omitempty"`
	// Runtime game room toleration configuration.
	Toleration string `protobuf:"bytes,7,opt,name=toleration,proto3" json:"toleration,omitempty"`
	// Runtime game room affinity configuration.
	Affinity string `protobuf:"bytes,8,opt,name=affinity,proto3" json:"affinity,omitempty"`
	// The max surge of new rooms, used to scale and update
	MaxSurge string `protobuf:"bytes,9,opt,name=max_surge,json=maxSurge,proto3" json:"max_surge,omitempty"`
}

func (x *CreateSchedulerRequest) Reset() {
	*x = CreateSchedulerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schedulers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSchedulerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSchedulerRequest) ProtoMessage() {}

func (x *CreateSchedulerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schedulers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSchedulerRequest.ProtoReflect.Descriptor instead.
func (*CreateSchedulerRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_schedulers_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSchedulerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSchedulerRequest) GetGame() string {
	if x != nil {
		return x.Game
	}
	return ""
}

func (x *CreateSchedulerRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CreateSchedulerRequest) GetTerminationGracePeriod() int64 {
	if x != nil {
		return x.TerminationGracePeriod
	}
	return 0
}

func (x *CreateSchedulerRequest) GetContainers() []*Container {
	if x != nil {
		return x.Containers
	}
	return nil
}

func (x *CreateSchedulerRequest) GetPortRange() *PortRange {
	if x != nil {
		return x.PortRange
	}
	return nil
}

func (x *CreateSchedulerRequest) GetToleration() string {
	if x != nil {
		return x.Toleration
	}
	return ""
}

func (x *CreateSchedulerRequest) GetAffinity() string {
	if x != nil {
		return x.Affinity
	}
	return ""
}

func (x *CreateSchedulerRequest) GetMaxSurge() string {
	if x != nil {
		return x.MaxSurge
	}
	return ""
}

// Add rooms operation request payload + path parameters.
type AddRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Scheduler name where the rooms will be added.
	SchedulerName string `protobuf:"bytes,1,opt,name=scheduler_name,json=schedulerName,proto3" json:"scheduler_name,omitempty"`
	// Amount of rooms to be added.
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AddRoomsRequest) Reset() {
	*x = AddRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schedulers_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoomsRequest) ProtoMessage() {}

func (x *AddRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schedulers_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoomsRequest.ProtoReflect.Descriptor instead.
func (*AddRoomsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_schedulers_proto_rawDescGZIP(), []int{4}
}

func (x *AddRoomsRequest) GetSchedulerName() string {
	if x != nil {
		return x.SchedulerName
	}
	return ""
}

func (x *AddRoomsRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Add rooms operation response payload, empty.
type AddRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Add rooms operation ID, further this id will be used to consult its state.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (x *AddRoomsResponse) Reset() {
	*x = AddRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schedulers_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoomsResponse) ProtoMessage() {}

func (x *AddRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schedulers_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoomsResponse.ProtoReflect.Descriptor instead.
func (*AddRoomsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_schedulers_proto_rawDescGZIP(), []int{5}
}

func (x *AddRoomsResponse) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

// Remove rooms operation request payload + path parameters.
type RemoveRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Scheduler name from which the rooms will be removed.
	SchedulerName string `protobuf:"bytes,1,opt,name=scheduler_name,json=schedulerName,proto3" json:"scheduler_name,omitempty"`
	// Amount of rooms to be removed.
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *RemoveRoomsRequest) Reset() {
	*x = RemoveRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schedulers_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoomsRequest) ProtoMessage() {}

func (x *RemoveRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schedulers_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoomsRequest.ProtoReflect.Descriptor instead.
func (*RemoveRoomsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_schedulers_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveRoomsRequest) GetSchedulerName() string {
	if x != nil {
		return x.SchedulerName
	}
	return ""
}

func (x *RemoveRoomsRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Remove rooms operation response payload.
type RemoveRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Remove rooms operation ID, further this id can be used to consult its state.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (x *RemoveRoomsResponse) Reset() {
	*x = RemoveRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schedulers_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoomsResponse) ProtoMessage() {}

func (x *RemoveRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schedulers_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoomsResponse.ProtoReflect.Descriptor instead.
func (*RemoveRoomsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_schedulers_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveRoomsResponse) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

// Get Scheduler operation request
type GetSchedulerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Scheduler name where the rooms will be added.
	SchedulerName string `protobuf:"bytes,1,opt,name=scheduler_name,json=schedulerName,proto3" json:"scheduler_name,omitempty"`
	// Scheduler version to be queried (query param)
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetSchedulerRequest) Reset() {
	*x = GetSchedulerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schedulers_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchedulerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchedulerRequest) ProtoMessage() {}

func (x *GetSchedulerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schedulers_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchedulerRequest.ProtoReflect.Descriptor instead.
func (*GetSchedulerRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_schedulers_proto_rawDescGZIP(), []int{8}
}

func (x *GetSchedulerRequest) GetSchedulerName() string {
	if x != nil {
		return x.SchedulerName
	}
	return ""
}

func (x *GetSchedulerRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// The list schedulers reponse message.
type GetSchedulerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of schedulers fetched.
	Scheduler *Scheduler `protobuf:"bytes,1,opt,name=scheduler,proto3" json:"scheduler,omitempty"`
}

func (x *GetSchedulerResponse) Reset() {
	*x = GetSchedulerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_schedulers_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchedulerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchedulerResponse) ProtoMessage() {}

func (x *GetSchedulerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_schedulers_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchedulerResponse.ProtoReflect.Descriptor instead.
func (*GetSchedulerResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_schedulers_proto_rawDescGZIP(), []int{9}
}

func (x *GetSchedulerResponse) GetScheduler() *Scheduler {
	if x != nil {
		return x.Scheduler
	}
	return nil
}

var File_api_v1_schedulers_proto protoreflect.FileDescriptor

var file_api_v1_schedulers_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x51, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x64, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x73, 0x22, 0x4a, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0xd2,
	0x02, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x67, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x18, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x61, 0x63, 0x65, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x09, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f,
	0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x66,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x66,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x75,
	0x72, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x75,
	0x72, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6f, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x12,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x38, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x52, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x32, 0xcc, 0x04, 0x0a,
	0x11, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x64, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x71, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x2a, 0x7d, 0x12, 0x6a, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x72, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x52, 0x6f,
	0x6f, 0x6d, 0x73, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x22, 0x28,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x2a, 0x7d, 0x2f, 0x61,
	0x64, 0x64, 0x2d, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0b, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x22, 0x2b, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x2a, 0x7d, 0x2f, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x2d, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x51, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x6f, 0x70, 0x66, 0x72, 0x65, 0x65, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e,
	0x6d, 0x61, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x6f, 0x70, 0x66, 0x72, 0x65, 0x65, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x65, 0x73,
	0x74, 0x72, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_schedulers_proto_rawDescOnce sync.Once
	file_api_v1_schedulers_proto_rawDescData = file_api_v1_schedulers_proto_rawDesc
)

func file_api_v1_schedulers_proto_rawDescGZIP() []byte {
	file_api_v1_schedulers_proto_rawDescOnce.Do(func() {
		file_api_v1_schedulers_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_schedulers_proto_rawDescData)
	})
	return file_api_v1_schedulers_proto_rawDescData
}

var file_api_v1_schedulers_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1_schedulers_proto_goTypes = []interface{}{
	(*ListSchedulersRequest)(nil),   // 0: api.v1.ListSchedulersRequest
	(*ListSchedulersResponse)(nil),  // 1: api.v1.ListSchedulersResponse
	(*CreateSchedulerResponse)(nil), // 2: api.v1.CreateSchedulerResponse
	(*CreateSchedulerRequest)(nil),  // 3: api.v1.CreateSchedulerRequest
	(*AddRoomsRequest)(nil),         // 4: api.v1.AddRoomsRequest
	(*AddRoomsResponse)(nil),        // 5: api.v1.AddRoomsResponse
	(*RemoveRoomsRequest)(nil),      // 6: api.v1.RemoveRoomsRequest
	(*RemoveRoomsResponse)(nil),     // 7: api.v1.RemoveRoomsResponse
	(*GetSchedulerRequest)(nil),     // 8: api.v1.GetSchedulerRequest
	(*GetSchedulerResponse)(nil),    // 9: api.v1.GetSchedulerResponse
	(*ListedScheduler)(nil),         // 10: api.v1.ListedScheduler
	(*Scheduler)(nil),               // 11: api.v1.Scheduler
	(*Container)(nil),               // 12: api.v1.Container
	(*PortRange)(nil),               // 13: api.v1.PortRange
}
var file_api_v1_schedulers_proto_depIdxs = []int32{
	10, // 0: api.v1.ListSchedulersResponse.schedulers:type_name -> api.v1.ListedScheduler
	11, // 1: api.v1.CreateSchedulerResponse.scheduler:type_name -> api.v1.Scheduler
	12, // 2: api.v1.CreateSchedulerRequest.containers:type_name -> api.v1.Container
	13, // 3: api.v1.CreateSchedulerRequest.port_range:type_name -> api.v1.PortRange
	11, // 4: api.v1.GetSchedulerResponse.scheduler:type_name -> api.v1.Scheduler
	0,  // 5: api.v1.SchedulersService.ListSchedulers:input_type -> api.v1.ListSchedulersRequest
	8,  // 6: api.v1.SchedulersService.GetScheduler:input_type -> api.v1.GetSchedulerRequest
	3,  // 7: api.v1.SchedulersService.CreateScheduler:input_type -> api.v1.CreateSchedulerRequest
	4,  // 8: api.v1.SchedulersService.AddRooms:input_type -> api.v1.AddRoomsRequest
	6,  // 9: api.v1.SchedulersService.RemoveRooms:input_type -> api.v1.RemoveRoomsRequest
	1,  // 10: api.v1.SchedulersService.ListSchedulers:output_type -> api.v1.ListSchedulersResponse
	9,  // 11: api.v1.SchedulersService.GetScheduler:output_type -> api.v1.GetSchedulerResponse
	2,  // 12: api.v1.SchedulersService.CreateScheduler:output_type -> api.v1.CreateSchedulerResponse
	5,  // 13: api.v1.SchedulersService.AddRooms:output_type -> api.v1.AddRoomsResponse
	7,  // 14: api.v1.SchedulersService.RemoveRooms:output_type -> api.v1.RemoveRoomsResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_schedulers_proto_init() }
func file_api_v1_schedulers_proto_init() {
	if File_api_v1_schedulers_proto != nil {
		return
	}
	file_api_v1_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_schedulers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSchedulersRequest); i {
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
		file_api_v1_schedulers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSchedulersResponse); i {
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
		file_api_v1_schedulers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSchedulerResponse); i {
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
		file_api_v1_schedulers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSchedulerRequest); i {
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
		file_api_v1_schedulers_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoomsRequest); i {
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
		file_api_v1_schedulers_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoomsResponse); i {
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
		file_api_v1_schedulers_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRoomsRequest); i {
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
		file_api_v1_schedulers_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRoomsResponse); i {
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
		file_api_v1_schedulers_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchedulerRequest); i {
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
		file_api_v1_schedulers_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchedulerResponse); i {
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
			RawDescriptor: file_api_v1_schedulers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_schedulers_proto_goTypes,
		DependencyIndexes: file_api_v1_schedulers_proto_depIdxs,
		MessageInfos:      file_api_v1_schedulers_proto_msgTypes,
	}.Build()
	File_api_v1_schedulers_proto = out.File
	file_api_v1_schedulers_proto_rawDesc = nil
	file_api_v1_schedulers_proto_goTypes = nil
	file_api_v1_schedulers_proto_depIdxs = nil
}
