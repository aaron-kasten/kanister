// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: pkg/kanx/kanx.proto

package kanx

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

type ProcessState int32

const (
	ProcessState_PROCESS_STATE_UNSPECIFIED ProcessState = 0
	ProcessState_PROCESS_STATE_RUNNING     ProcessState = 1
	ProcessState_PROCESS_STATE_SUCCEEDED   ProcessState = 2
	ProcessState_PROCESS_STATE_FAILED      ProcessState = 3
)

// Enum value maps for ProcessState.
var (
	ProcessState_name = map[int32]string{
		0: "PROCESS_STATE_UNSPECIFIED",
		1: "PROCESS_STATE_RUNNING",
		2: "PROCESS_STATE_SUCCEEDED",
		3: "PROCESS_STATE_FAILED",
	}
	ProcessState_value = map[string]int32{
		"PROCESS_STATE_UNSPECIFIED": 0,
		"PROCESS_STATE_RUNNING":     1,
		"PROCESS_STATE_SUCCEEDED":   2,
		"PROCESS_STATE_FAILED":      3,
	}
)

func (x ProcessState) Enum() *ProcessState {
	p := new(ProcessState)
	*p = x
	return p
}

func (x ProcessState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessState) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_kanx_kanx_proto_enumTypes[0].Descriptor()
}

func (ProcessState) Type() protoreflect.EnumType {
	return &file_pkg_kanx_kanx_proto_enumTypes[0]
}

func (x ProcessState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessState.Descriptor instead.
func (ProcessState) EnumDescriptor() ([]byte, []int) {
	return file_pkg_kanx_kanx_proto_rawDescGZIP(), []int{0}
}

type ProcessOutputStream int32

const (
	ProcessOutputStream_PROCESS_OUTPUT_UNSPECIFIED ProcessOutputStream = 0
	ProcessOutputStream_PROCESS_OUTPUT_STDOUT      ProcessOutputStream = 1
	ProcessOutputStream_PROCESS_OUTPUT_STDERR      ProcessOutputStream = 2
)

// Enum value maps for ProcessOutputStream.
var (
	ProcessOutputStream_name = map[int32]string{
		0: "PROCESS_OUTPUT_UNSPECIFIED",
		1: "PROCESS_OUTPUT_STDOUT",
		2: "PROCESS_OUTPUT_STDERR",
	}
	ProcessOutputStream_value = map[string]int32{
		"PROCESS_OUTPUT_UNSPECIFIED": 0,
		"PROCESS_OUTPUT_STDOUT":      1,
		"PROCESS_OUTPUT_STDERR":      2,
	}
)

func (x ProcessOutputStream) Enum() *ProcessOutputStream {
	p := new(ProcessOutputStream)
	*p = x
	return p
}

func (x ProcessOutputStream) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessOutputStream) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_kanx_kanx_proto_enumTypes[1].Descriptor()
}

func (ProcessOutputStream) Type() protoreflect.EnumType {
	return &file_pkg_kanx_kanx_proto_enumTypes[1]
}

func (x ProcessOutputStream) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessOutputStream.Descriptor instead.
func (ProcessOutputStream) EnumDescriptor() ([]byte, []int) {
	return file_pkg_kanx_kanx_proto_rawDescGZIP(), []int{1}
}

type CreateProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *CreateProcessRequest) Reset() {
	*x = CreateProcessRequest{}
	mi := &file_pkg_kanx_kanx_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProcessRequest) ProtoMessage() {}

func (x *CreateProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_kanx_kanx_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProcessRequest.ProtoReflect.Descriptor instead.
func (*CreateProcessRequest) Descriptor() ([]byte, []int) {
	return file_pkg_kanx_kanx_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProcessRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProcessRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

type ListProcessesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProcessesRequest) Reset() {
	*x = ListProcessesRequest{}
	mi := &file_pkg_kanx_kanx_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListProcessesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProcessesRequest) ProtoMessage() {}

func (x *ListProcessesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_kanx_kanx_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProcessesRequest.ProtoReflect.Descriptor instead.
func (*ListProcessesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_kanx_kanx_proto_rawDescGZIP(), []int{1}
}

type ProcessPidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int64 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *ProcessPidRequest) Reset() {
	*x = ProcessPidRequest{}
	mi := &file_pkg_kanx_kanx_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessPidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPidRequest) ProtoMessage() {}

func (x *ProcessPidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_kanx_kanx_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPidRequest.ProtoReflect.Descriptor instead.
func (*ProcessPidRequest) Descriptor() ([]byte, []int) {
	return file_pkg_kanx_kanx_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessPidRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid      int64        `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	State    ProcessState `protobuf:"varint,2,opt,name=state,proto3,enum=kanx.ProcessState" json:"state,omitempty"`
	ExitCode int64        `protobuf:"varint,3,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	ExitErr  string       `protobuf:"bytes,4,opt,name=exitErr,proto3" json:"exitErr,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	mi := &file_pkg_kanx_kanx_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_kanx_kanx_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_pkg_kanx_kanx_proto_rawDescGZIP(), []int{3}
}

func (x *Process) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Process) GetState() ProcessState {
	if x != nil {
		return x.State
	}
	return ProcessState_PROCESS_STATE_UNSPECIFIED
}

func (x *Process) GetExitCode() int64 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *Process) GetExitErr() string {
	if x != nil {
		return x.ExitErr
	}
	return ""
}

type StdoutOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream string `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *StdoutOutput) Reset() {
	*x = StdoutOutput{}
	mi := &file_pkg_kanx_kanx_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StdoutOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StdoutOutput) ProtoMessage() {}

func (x *StdoutOutput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_kanx_kanx_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StdoutOutput.ProtoReflect.Descriptor instead.
func (*StdoutOutput) Descriptor() ([]byte, []int) {
	return file_pkg_kanx_kanx_proto_rawDescGZIP(), []int{4}
}

func (x *StdoutOutput) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

type StderrOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream string `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *StderrOutput) Reset() {
	*x = StderrOutput{}
	mi := &file_pkg_kanx_kanx_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StderrOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StderrOutput) ProtoMessage() {}

func (x *StderrOutput) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_kanx_kanx_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StderrOutput.ProtoReflect.Descriptor instead.
func (*StderrOutput) Descriptor() ([]byte, []int) {
	return file_pkg_kanx_kanx_proto_rawDescGZIP(), []int{5}
}

func (x *StderrOutput) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

var File_pkg_kanx_kanx_proto protoreflect.FileDescriptor

var file_pkg_kanx_kanx_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x6b, 0x61, 0x6e, 0x78, 0x2f, 0x6b, 0x61, 0x6e, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6b, 0x61, 0x6e, 0x78, 0x22, 0x3e, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x69,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x7b, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6b, 0x61, 0x6e, 0x78, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x78, 0x69, 0x74, 0x45, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x78, 0x69, 0x74, 0x45, 0x72, 0x72, 0x22, 0x26, 0x0a, 0x0c, 0x53, 0x74, 0x64, 0x6f, 0x75,
	0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22,
	0x26, 0x0a, 0x0c, 0x53, 0x74, 0x64, 0x65, 0x72, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2a, 0x7f, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x52, 0x4f, 0x43, 0x45,
	0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18,
	0x0a, 0x14, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x6b, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x1e, 0x0a, 0x1a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55,
	0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55,
	0x54, 0x5f, 0x53, 0x54, 0x44, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52,
	0x4f, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x5f, 0x53, 0x54, 0x44,
	0x45, 0x52, 0x52, 0x10, 0x02, 0x32, 0xbc, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x6b, 0x61, 0x6e, 0x78,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6b, 0x61, 0x6e, 0x78, 0x2e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x17, 0x2e, 0x6b, 0x61, 0x6e, 0x78, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x50, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x6b, 0x61, 0x6e, 0x78, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x1a, 0x2e, 0x6b, 0x61, 0x6e, 0x78, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6b, 0x61,
	0x6e, 0x78, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39,
	0x0a, 0x06, 0x53, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x6b, 0x61, 0x6e, 0x78, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x50, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x6b, 0x61, 0x6e, 0x78, 0x2e, 0x53, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x06, 0x53, 0x74, 0x64,
	0x65, 0x72, 0x72, 0x12, 0x17, 0x2e, 0x6b, 0x61, 0x6e, 0x78, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6b,
	0x61, 0x6e, 0x78, 0x2e, 0x53, 0x74, 0x64, 0x65, 0x72, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x6e, 0x69, 0x73, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x2f, 0x6b, 0x61,
	0x6e, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6b, 0x61, 0x6e, 0x78, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_kanx_kanx_proto_rawDescOnce sync.Once
	file_pkg_kanx_kanx_proto_rawDescData = file_pkg_kanx_kanx_proto_rawDesc
)

func file_pkg_kanx_kanx_proto_rawDescGZIP() []byte {
	file_pkg_kanx_kanx_proto_rawDescOnce.Do(func() {
		file_pkg_kanx_kanx_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_kanx_kanx_proto_rawDescData)
	})
	return file_pkg_kanx_kanx_proto_rawDescData
}

var file_pkg_kanx_kanx_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_kanx_kanx_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_kanx_kanx_proto_goTypes = []any{
	(ProcessState)(0),            // 0: kanx.ProcessState
	(ProcessOutputStream)(0),     // 1: kanx.ProcessOutputStream
	(*CreateProcessRequest)(nil), // 2: kanx.CreateProcessRequest
	(*ListProcessesRequest)(nil), // 3: kanx.ListProcessesRequest
	(*ProcessPidRequest)(nil),    // 4: kanx.ProcessPidRequest
	(*Process)(nil),              // 5: kanx.Process
	(*StdoutOutput)(nil),         // 6: kanx.StdoutOutput
	(*StderrOutput)(nil),         // 7: kanx.StderrOutput
}
var file_pkg_kanx_kanx_proto_depIdxs = []int32{
	0, // 0: kanx.Process.state:type_name -> kanx.ProcessState
	2, // 1: kanx.ProcessService.CreateProcess:input_type -> kanx.CreateProcessRequest
	4, // 2: kanx.ProcessService.GetProcess:input_type -> kanx.ProcessPidRequest
	3, // 3: kanx.ProcessService.ListProcesses:input_type -> kanx.ListProcessesRequest
	4, // 4: kanx.ProcessService.Stdout:input_type -> kanx.ProcessPidRequest
	4, // 5: kanx.ProcessService.Stderr:input_type -> kanx.ProcessPidRequest
	5, // 6: kanx.ProcessService.CreateProcess:output_type -> kanx.Process
	5, // 7: kanx.ProcessService.GetProcess:output_type -> kanx.Process
	5, // 8: kanx.ProcessService.ListProcesses:output_type -> kanx.Process
	6, // 9: kanx.ProcessService.Stdout:output_type -> kanx.StdoutOutput
	7, // 10: kanx.ProcessService.Stderr:output_type -> kanx.StderrOutput
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_kanx_kanx_proto_init() }
func file_pkg_kanx_kanx_proto_init() {
	if File_pkg_kanx_kanx_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_kanx_kanx_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_kanx_kanx_proto_goTypes,
		DependencyIndexes: file_pkg_kanx_kanx_proto_depIdxs,
		EnumInfos:         file_pkg_kanx_kanx_proto_enumTypes,
		MessageInfos:      file_pkg_kanx_kanx_proto_msgTypes,
	}.Build()
	File_pkg_kanx_kanx_proto = out.File
	file_pkg_kanx_kanx_proto_rawDesc = nil
	file_pkg_kanx_kanx_proto_goTypes = nil
	file_pkg_kanx_kanx_proto_depIdxs = nil
}
