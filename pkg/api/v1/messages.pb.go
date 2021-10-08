// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/v1/messages.proto

package v1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Port range is the range definition that the rooms will use. If a scheduler
// defines its range as 0-1000 (start-end), it is guarantee that all rooms be
// within this range.
type PortRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Range start, it is inclusive.
	Start int32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// Range end, it is exclusive.
	End int32 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *PortRange) Reset() {
	*x = PortRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortRange) ProtoMessage() {}

func (x *PortRange) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortRange.ProtoReflect.Descriptor instead.
func (*PortRange) Descriptor() ([]byte, []int) {
	return file_api_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PortRange) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *PortRange) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

// Scheduler definition.
type Scheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is an unique identifier for the scheduler.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Game is the "scope" where the scheduler will be placed. It is required, and
	// games name are unique.
	Game string `protobuf:"bytes,2,opt,name=game,proto3" json:"game,omitempty"`
	// Current state of the scheduler.
	// TODO(gabrielcorado): move this to an ENUM.
	State string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	// Scheduler current version.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Port range for the scheduler rooms.
	PortRange *PortRange `protobuf:"bytes,5,opt,name=port_range,json=portRange,proto3" json:"port_range,omitempty"`
	// Time the scheduler was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Max surge of rooms
	MaxSurge string `protobuf:"bytes,7,opt,name=max_surge,json=maxSurge,proto3" json:"max_surge,omitempty"`
}

func (x *Scheduler) Reset() {
	*x = Scheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scheduler) ProtoMessage() {}

func (x *Scheduler) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scheduler.ProtoReflect.Descriptor instead.
func (*Scheduler) Descriptor() ([]byte, []int) {
	return file_api_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Scheduler) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scheduler) GetGame() string {
	if x != nil {
		return x.Game
	}
	return ""
}

func (x *Scheduler) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Scheduler) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Scheduler) GetPortRange() *PortRange {
	if x != nil {
		return x.PortRange
	}
	return nil
}

func (x *Scheduler) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Scheduler) GetMaxSurge() string {
	if x != nil {
		return x.MaxSurge
	}
	return ""
}

// The operation object representation
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the operation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Current status of the operation.
	// TODO(gabrielcorado): change to enum.
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// Name of the operation being executed.
	DefinitionName string `protobuf:"bytes,3,opt,name=definition_name,json=definitionName,proto3" json:"definition_name,omitempty"`
	// Scheduler indentifier that the operation is from.
	SchedulerName string `protobuf:"bytes,4,opt,name=scheduler_name,json=schedulerName,proto3" json:"scheduler_name,omitempty"`
	// Time the operation was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_api_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Operation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Operation) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Operation) GetDefinitionName() string {
	if x != nil {
		return x.DefinitionName
	}
	return ""
}

func (x *Operation) GetSchedulerName() string {
	if x != nil {
		return x.SchedulerName
	}
	return ""
}

func (x *Operation) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_api_v1_messages_proto protoreflect.FileDescriptor

var file_api_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x33, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0xed, 0x01, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x0a, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f,
	0x73, 0x75, 0x72, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x78,
	0x53, 0x75, 0x72, 0x67, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x51, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6f,
	0x70, 0x66, 0x72, 0x65, 0x65, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x65, 0x73, 0x74,
	0x72, 0x6f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x70, 0x66, 0x72, 0x65,
	0x65, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_v1_messages_proto_rawDescOnce sync.Once
	file_api_v1_messages_proto_rawDescData = file_api_v1_messages_proto_rawDesc
)

func file_api_v1_messages_proto_rawDescGZIP() []byte {
	file_api_v1_messages_proto_rawDescOnce.Do(func() {
		file_api_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_messages_proto_rawDescData)
	})
	return file_api_v1_messages_proto_rawDescData
}

var file_api_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_v1_messages_proto_goTypes = []interface{}{
	(*PortRange)(nil),           // 0: api.v1.PortRange
	(*Scheduler)(nil),           // 1: api.v1.Scheduler
	(*Operation)(nil),           // 2: api.v1.Operation
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_api_v1_messages_proto_depIdxs = []int32{
	0, // 0: api.v1.Scheduler.port_range:type_name -> api.v1.PortRange
	3, // 1: api.v1.Scheduler.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: api.v1.Operation.created_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_v1_messages_proto_init() }
func file_api_v1_messages_proto_init() {
	if File_api_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortRange); i {
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
		file_api_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scheduler); i {
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
		file_api_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
			RawDescriptor: file_api_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_messages_proto_goTypes,
		DependencyIndexes: file_api_v1_messages_proto_depIdxs,
		MessageInfos:      file_api_v1_messages_proto_msgTypes,
	}.Build()
	File_api_v1_messages_proto = out.File
	file_api_v1_messages_proto_rawDesc = nil
	file_api_v1_messages_proto_goTypes = nil
	file_api_v1_messages_proto_depIdxs = nil
}
