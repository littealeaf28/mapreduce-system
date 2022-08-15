// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: pb/rpc.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type StatusType int32

const (
	StatusType_Incomplete StatusType = 0
	StatusType_Processing StatusType = 1
	StatusType_Complete   StatusType = 2
)

// Enum value maps for StatusType.
var (
	StatusType_name = map[int32]string{
		0: "Incomplete",
		1: "Processing",
		2: "Complete",
	}
	StatusType_value = map[string]int32{
		"Incomplete": 0,
		"Processing": 1,
		"Complete":   2,
	}
)

func (x StatusType) Enum() *StatusType {
	p := new(StatusType)
	*p = x
	return p
}

func (x StatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (StatusType) Type() protoreflect.EnumType {
	return &file_pb_rpc_proto_enumTypes[0]
}

func (x StatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusType.Descriptor instead.
func (StatusType) EnumDescriptor() ([]byte, []int) {
	return file_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type TaskType int32

const (
	TaskType_Empty  TaskType = 0
	TaskType_Map    TaskType = 1
	TaskType_Reduce TaskType = 2
)

// Enum value maps for TaskType.
var (
	TaskType_name = map[int32]string{
		0: "Empty",
		1: "Map",
		2: "Reduce",
	}
	TaskType_value = map[string]int32{
		"Empty":  0,
		"Map":    1,
		"Reduce": 2,
	}
)

func (x TaskType) Enum() *TaskType {
	p := new(TaskType)
	*p = x
	return p
}

func (x TaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_rpc_proto_enumTypes[1].Descriptor()
}

func (TaskType) Type() protoreflect.EnumType {
	return &file_pb_rpc_proto_enumTypes[1]
}

func (x TaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskType.Descriptor instead.
func (TaskType) EnumDescriptor() ([]byte, []int) {
	return file_pb_rpc_proto_rawDescGZIP(), []int{1}
}

type TaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       TaskType    `protobuf:"varint,1,opt,name=type,proto3,enum=pb.TaskType" json:"type,omitempty"`
	MapTask    *MapTask    `protobuf:"bytes,2,opt,name=mapTask,proto3" json:"mapTask,omitempty"`
	ReduceTask *ReduceTask `protobuf:"bytes,3,opt,name=reduceTask,proto3" json:"reduceTask,omitempty"`
}

func (x *TaskReply) Reset() {
	*x = TaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskReply) ProtoMessage() {}

func (x *TaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskReply.ProtoReflect.Descriptor instead.
func (*TaskReply) Descriptor() ([]byte, []int) {
	return file_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *TaskReply) GetType() TaskType {
	if x != nil {
		return x.Type
	}
	return TaskType_Empty
}

func (x *TaskReply) GetMapTask() *MapTask {
	if x != nil {
		return x.MapTask
	}
	return nil
}

func (x *TaskReply) GetReduceTask() *ReduceTask {
	if x != nil {
		return x.ReduceTask
	}
	return nil
}

type MapTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   StatusType `protobuf:"varint,1,opt,name=status,proto3,enum=pb.StatusType" json:"status,omitempty"`
	FileName string     `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	MapNum   int32      `protobuf:"varint,3,opt,name=mapNum,proto3" json:"mapNum,omitempty"`
	NReduce  int32      `protobuf:"varint,4,opt,name=nReduce,proto3" json:"nReduce,omitempty"`
}

func (x *MapTask) Reset() {
	*x = MapTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapTask) ProtoMessage() {}

func (x *MapTask) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapTask.ProtoReflect.Descriptor instead.
func (*MapTask) Descriptor() ([]byte, []int) {
	return file_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *MapTask) GetStatus() StatusType {
	if x != nil {
		return x.Status
	}
	return StatusType_Incomplete
}

func (x *MapTask) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *MapTask) GetMapNum() int32 {
	if x != nil {
		return x.MapNum
	}
	return 0
}

func (x *MapTask) GetNReduce() int32 {
	if x != nil {
		return x.NReduce
	}
	return 0
}

type ReduceTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    StatusType `protobuf:"varint,1,opt,name=status,proto3,enum=pb.StatusType" json:"status,omitempty"`
	ReduceNum int32      `protobuf:"varint,2,opt,name=reduceNum,proto3" json:"reduceNum,omitempty"`
}

func (x *ReduceTask) Reset() {
	*x = ReduceTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceTask) ProtoMessage() {}

func (x *ReduceTask) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceTask.ProtoReflect.Descriptor instead.
func (*ReduceTask) Descriptor() ([]byte, []int) {
	return file_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *ReduceTask) GetStatus() StatusType {
	if x != nil {
		return x.Status
	}
	return StatusType_Incomplete
}

func (x *ReduceTask) GetReduceNum() int32 {
	if x != nil {
		return x.ReduceNum
	}
	return 0
}

type TaskComplete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type TaskType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.TaskType" json:"type,omitempty"`
	Num  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *TaskComplete) Reset() {
	*x = TaskComplete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskComplete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskComplete) ProtoMessage() {}

func (x *TaskComplete) ProtoReflect() protoreflect.Message {
	mi := &file_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskComplete.ProtoReflect.Descriptor instead.
func (*TaskComplete) Descriptor() ([]byte, []int) {
	return file_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *TaskComplete) GetType() TaskType {
	if x != nil {
		return x.Type
	}
	return TaskType_Empty
}

func (x *TaskComplete) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_pb_rpc_proto protoreflect.FileDescriptor

var file_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x84, 0x01, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x25, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x07, 0x6d,
	0x61, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x0a, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x72, 0x65, 0x64, 0x75,
	0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x7f, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x4e, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x70, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x22, 0x52, 0x0a, 0x0a, 0x52, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x42, 0x0a, 0x0c, 0x54,
	0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x2a,
	0x3a, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x0a, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x02, 0x2a, 0x2a, 0x0a, 0x08, 0x54,
	0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x65, 0x64, 0x75, 0x63, 0x65, 0x10, 0x02, 0x32, 0x7d, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x6d, 0x72, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_rpc_proto_rawDescOnce sync.Once
	file_pb_rpc_proto_rawDescData = file_pb_rpc_proto_rawDesc
)

func file_pb_rpc_proto_rawDescGZIP() []byte {
	file_pb_rpc_proto_rawDescOnce.Do(func() {
		file_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_rpc_proto_rawDescData)
	})
	return file_pb_rpc_proto_rawDescData
}

var file_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_rpc_proto_goTypes = []interface{}{
	(StatusType)(0),      // 0: pb.StatusType
	(TaskType)(0),        // 1: pb.TaskType
	(*TaskReply)(nil),    // 2: pb.TaskReply
	(*MapTask)(nil),      // 3: pb.MapTask
	(*ReduceTask)(nil),   // 4: pb.ReduceTask
	(*TaskComplete)(nil), // 5: pb.TaskComplete
	(*empty.Empty)(nil),  // 6: google.protobuf.Empty
}
var file_pb_rpc_proto_depIdxs = []int32{
	1, // 0: pb.TaskReply.type:type_name -> pb.TaskType
	3, // 1: pb.TaskReply.mapTask:type_name -> pb.MapTask
	4, // 2: pb.TaskReply.reduceTask:type_name -> pb.ReduceTask
	0, // 3: pb.MapTask.status:type_name -> pb.StatusType
	0, // 4: pb.ReduceTask.status:type_name -> pb.StatusType
	1, // 5: pb.TaskComplete.type:type_name -> pb.TaskType
	6, // 6: pb.Coordinator.GetTask:input_type -> google.protobuf.Empty
	5, // 7: pb.Coordinator.CompleteTask:input_type -> pb.TaskComplete
	2, // 8: pb.Coordinator.GetTask:output_type -> pb.TaskReply
	6, // 9: pb.Coordinator.CompleteTask:output_type -> google.protobuf.Empty
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pb_rpc_proto_init() }
func file_pb_rpc_proto_init() {
	if File_pb_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskReply); i {
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
		file_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapTask); i {
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
		file_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceTask); i {
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
		file_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskComplete); i {
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
			RawDescriptor: file_pb_rpc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_rpc_proto_goTypes,
		DependencyIndexes: file_pb_rpc_proto_depIdxs,
		EnumInfos:         file_pb_rpc_proto_enumTypes,
		MessageInfos:      file_pb_rpc_proto_msgTypes,
	}.Build()
	File_pb_rpc_proto = out.File
	file_pb_rpc_proto_rawDesc = nil
	file_pb_rpc_proto_goTypes = nil
	file_pb_rpc_proto_depIdxs = nil
}
