// Code generated by protoc-gen-go. DO NOT EDIT.
// source: insonmnia.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TaskStatusReply_Status int32

const (
	TaskStatusReply_UNKNOWN  TaskStatusReply_Status = 0
	TaskStatusReply_SPOOLING TaskStatusReply_Status = 1
	TaskStatusReply_SPAWNING TaskStatusReply_Status = 2
	TaskStatusReply_RUNNING  TaskStatusReply_Status = 3
	TaskStatusReply_FINISHED TaskStatusReply_Status = 4
	TaskStatusReply_BROKEN   TaskStatusReply_Status = 5
)

var TaskStatusReply_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "SPOOLING",
	2: "SPAWNING",
	3: "RUNNING",
	4: "FINISHED",
	5: "BROKEN",
}
var TaskStatusReply_Status_value = map[string]int32{
	"UNKNOWN":  0,
	"SPOOLING": 1,
	"SPAWNING": 2,
	"RUNNING":  3,
	"FINISHED": 4,
	"BROKEN":   5,
}

func (x TaskStatusReply_Status) String() string {
	return proto.EnumName(TaskStatusReply_Status_name, int32(x))
}
func (TaskStatusReply_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{10, 0} }

type TaskLogsRequest_Type int32

const (
	TaskLogsRequest_STDOUT TaskLogsRequest_Type = 0
	TaskLogsRequest_STDERR TaskLogsRequest_Type = 1
	TaskLogsRequest_BOTH   TaskLogsRequest_Type = 2
)

var TaskLogsRequest_Type_name = map[int32]string{
	0: "STDOUT",
	1: "STDERR",
	2: "BOTH",
}
var TaskLogsRequest_Type_value = map[string]int32{
	"STDOUT": 0,
	"STDERR": 1,
	"BOTH":   2,
}

func (x TaskLogsRequest_Type) String() string {
	return proto.EnumName(TaskLogsRequest_Type_name, int32(x))
}
func (TaskLogsRequest_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{14, 0} }

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type PingReply struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *PingReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type CPUUsage struct {
	Total uint64 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
}

func (m *CPUUsage) Reset()                    { *m = CPUUsage{} }
func (m *CPUUsage) String() string            { return proto.CompactTextString(m) }
func (*CPUUsage) ProtoMessage()               {}
func (*CPUUsage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *CPUUsage) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type MemoryUsage struct {
	MaxUsage uint64 `protobuf:"varint,1,opt,name=maxUsage" json:"maxUsage,omitempty"`
}

func (m *MemoryUsage) Reset()                    { *m = MemoryUsage{} }
func (m *MemoryUsage) String() string            { return proto.CompactTextString(m) }
func (*MemoryUsage) ProtoMessage()               {}
func (*MemoryUsage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *MemoryUsage) GetMaxUsage() uint64 {
	if m != nil {
		return m.MaxUsage
	}
	return 0
}

type NetworkUsage struct {
	TxBytes   uint64 `protobuf:"varint,1,opt,name=txBytes" json:"txBytes,omitempty"`
	RxBytes   uint64 `protobuf:"varint,2,opt,name=rxBytes" json:"rxBytes,omitempty"`
	TxPackets uint64 `protobuf:"varint,3,opt,name=txPackets" json:"txPackets,omitempty"`
	RxPackets uint64 `protobuf:"varint,4,opt,name=rxPackets" json:"rxPackets,omitempty"`
	TxErrors  uint64 `protobuf:"varint,5,opt,name=txErrors" json:"txErrors,omitempty"`
	RxErrors  uint64 `protobuf:"varint,6,opt,name=rxErrors" json:"rxErrors,omitempty"`
	TxDropped uint64 `protobuf:"varint,7,opt,name=txDropped" json:"txDropped,omitempty"`
	RxDropped uint64 `protobuf:"varint,8,opt,name=rxDropped" json:"rxDropped,omitempty"`
}

func (m *NetworkUsage) Reset()                    { *m = NetworkUsage{} }
func (m *NetworkUsage) String() string            { return proto.CompactTextString(m) }
func (*NetworkUsage) ProtoMessage()               {}
func (*NetworkUsage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *NetworkUsage) GetTxBytes() uint64 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

func (m *NetworkUsage) GetRxBytes() uint64 {
	if m != nil {
		return m.RxBytes
	}
	return 0
}

func (m *NetworkUsage) GetTxPackets() uint64 {
	if m != nil {
		return m.TxPackets
	}
	return 0
}

func (m *NetworkUsage) GetRxPackets() uint64 {
	if m != nil {
		return m.RxPackets
	}
	return 0
}

func (m *NetworkUsage) GetTxErrors() uint64 {
	if m != nil {
		return m.TxErrors
	}
	return 0
}

func (m *NetworkUsage) GetRxErrors() uint64 {
	if m != nil {
		return m.RxErrors
	}
	return 0
}

func (m *NetworkUsage) GetTxDropped() uint64 {
	if m != nil {
		return m.TxDropped
	}
	return 0
}

func (m *NetworkUsage) GetRxDropped() uint64 {
	if m != nil {
		return m.RxDropped
	}
	return 0
}

type ResourceUsage struct {
	Cpu     *CPUUsage                `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	Memory  *MemoryUsage             `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
	Network map[string]*NetworkUsage `protobuf:"bytes,3,rep,name=network" json:"network,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ResourceUsage) Reset()                    { *m = ResourceUsage{} }
func (m *ResourceUsage) String() string            { return proto.CompactTextString(m) }
func (*ResourceUsage) ProtoMessage()               {}
func (*ResourceUsage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *ResourceUsage) GetCpu() *CPUUsage {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *ResourceUsage) GetMemory() *MemoryUsage {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *ResourceUsage) GetNetwork() map[string]*NetworkUsage {
	if m != nil {
		return m.Network
	}
	return nil
}

type InfoReply struct {
	Usage map[string]*ResourceUsage `protobuf:"bytes,1,rep,name=usage" json:"usage,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *InfoReply) Reset()                    { *m = InfoReply{} }
func (m *InfoReply) String() string            { return proto.CompactTextString(m) }
func (*InfoReply) ProtoMessage()               {}
func (*InfoReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *InfoReply) GetUsage() map[string]*ResourceUsage {
	if m != nil {
		return m.Usage
	}
	return nil
}

type StopTaskRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StopTaskRequest) Reset()                    { *m = StopTaskRequest{} }
func (m *StopTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*StopTaskRequest) ProtoMessage()               {}
func (*StopTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *StopTaskRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StopTaskReply struct {
}

func (m *StopTaskReply) Reset()                    { *m = StopTaskReply{} }
func (m *StopTaskReply) String() string            { return proto.CompactTextString(m) }
func (*StopTaskReply) ProtoMessage()               {}
func (*StopTaskReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

type TaskStatusRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *TaskStatusRequest) Reset()                    { *m = TaskStatusRequest{} }
func (m *TaskStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskStatusRequest) ProtoMessage()               {}
func (*TaskStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *TaskStatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TaskStatusReply struct {
	Status TaskStatusReply_Status `protobuf:"varint,1,opt,name=status,enum=sonm.TaskStatusReply_Status" json:"status,omitempty"`
}

func (m *TaskStatusReply) Reset()                    { *m = TaskStatusReply{} }
func (m *TaskStatusReply) String() string            { return proto.CompactTextString(m) }
func (*TaskStatusReply) ProtoMessage()               {}
func (*TaskStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *TaskStatusReply) GetStatus() TaskStatusReply_Status {
	if m != nil {
		return m.Status
	}
	return TaskStatusReply_UNKNOWN
}

type StatusMapReply struct {
	Statuses map[string]*TaskStatusReply `protobuf:"bytes,1,rep,name=statuses" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *StatusMapReply) Reset()                    { *m = StatusMapReply{} }
func (m *StatusMapReply) String() string            { return proto.CompactTextString(m) }
func (*StatusMapReply) ProtoMessage()               {}
func (*StatusMapReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *StatusMapReply) GetStatuses() map[string]*TaskStatusReply {
	if m != nil {
		return m.Statuses
	}
	return nil
}

type ContainerResources struct {
	Memory   int64 `protobuf:"varint,1,opt,name=memory" json:"memory,omitempty"`
	NanoCPUs int64 `protobuf:"varint,2,opt,name=nanoCPUs" json:"nanoCPUs,omitempty"`
}

func (m *ContainerResources) Reset()                    { *m = ContainerResources{} }
func (m *ContainerResources) String() string            { return proto.CompactTextString(m) }
func (*ContainerResources) ProtoMessage()               {}
func (*ContainerResources) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *ContainerResources) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *ContainerResources) GetNanoCPUs() int64 {
	if m != nil {
		return m.NanoCPUs
	}
	return 0
}

type ContainerRestartPolicy struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	MaximumRetryCount uint32 `protobuf:"varint,2,opt,name=maximumRetryCount" json:"maximumRetryCount,omitempty"`
}

func (m *ContainerRestartPolicy) Reset()                    { *m = ContainerRestartPolicy{} }
func (m *ContainerRestartPolicy) String() string            { return proto.CompactTextString(m) }
func (*ContainerRestartPolicy) ProtoMessage()               {}
func (*ContainerRestartPolicy) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *ContainerRestartPolicy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerRestartPolicy) GetMaximumRetryCount() uint32 {
	if m != nil {
		return m.MaximumRetryCount
	}
	return 0
}

type TaskLogsRequest struct {
	Type          TaskLogsRequest_Type `protobuf:"varint,1,opt,name=type,enum=sonm.TaskLogsRequest_Type" json:"type,omitempty"`
	Id            string               `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Since         string               `protobuf:"bytes,3,opt,name=since" json:"since,omitempty"`
	AddTimestamps bool                 `protobuf:"varint,4,opt,name=addTimestamps" json:"addTimestamps,omitempty"`
	Follow        bool                 `protobuf:"varint,5,opt,name=Follow" json:"Follow,omitempty"`
	Tail          string               `protobuf:"bytes,6,opt,name=Tail" json:"Tail,omitempty"`
	Details       bool                 `protobuf:"varint,7,opt,name=Details" json:"Details,omitempty"`
}

func (m *TaskLogsRequest) Reset()                    { *m = TaskLogsRequest{} }
func (m *TaskLogsRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskLogsRequest) ProtoMessage()               {}
func (*TaskLogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

func (m *TaskLogsRequest) GetType() TaskLogsRequest_Type {
	if m != nil {
		return m.Type
	}
	return TaskLogsRequest_STDOUT
}

func (m *TaskLogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskLogsRequest) GetSince() string {
	if m != nil {
		return m.Since
	}
	return ""
}

func (m *TaskLogsRequest) GetAddTimestamps() bool {
	if m != nil {
		return m.AddTimestamps
	}
	return false
}

func (m *TaskLogsRequest) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

func (m *TaskLogsRequest) GetTail() string {
	if m != nil {
		return m.Tail
	}
	return ""
}

func (m *TaskLogsRequest) GetDetails() bool {
	if m != nil {
		return m.Details
	}
	return false
}

type TaskLogsChunk struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TaskLogsChunk) Reset()                    { *m = TaskLogsChunk{} }
func (m *TaskLogsChunk) String() string            { return proto.CompactTextString(m) }
func (*TaskLogsChunk) ProtoMessage()               {}
func (*TaskLogsChunk) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{15} }

func (m *TaskLogsChunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "sonm.PingRequest")
	proto.RegisterType((*PingReply)(nil), "sonm.PingReply")
	proto.RegisterType((*CPUUsage)(nil), "sonm.CPUUsage")
	proto.RegisterType((*MemoryUsage)(nil), "sonm.MemoryUsage")
	proto.RegisterType((*NetworkUsage)(nil), "sonm.NetworkUsage")
	proto.RegisterType((*ResourceUsage)(nil), "sonm.ResourceUsage")
	proto.RegisterType((*InfoReply)(nil), "sonm.InfoReply")
	proto.RegisterType((*StopTaskRequest)(nil), "sonm.StopTaskRequest")
	proto.RegisterType((*StopTaskReply)(nil), "sonm.StopTaskReply")
	proto.RegisterType((*TaskStatusRequest)(nil), "sonm.TaskStatusRequest")
	proto.RegisterType((*TaskStatusReply)(nil), "sonm.TaskStatusReply")
	proto.RegisterType((*StatusMapReply)(nil), "sonm.StatusMapReply")
	proto.RegisterType((*ContainerResources)(nil), "sonm.ContainerResources")
	proto.RegisterType((*ContainerRestartPolicy)(nil), "sonm.ContainerRestartPolicy")
	proto.RegisterType((*TaskLogsRequest)(nil), "sonm.TaskLogsRequest")
	proto.RegisterType((*TaskLogsChunk)(nil), "sonm.TaskLogsChunk")
	proto.RegisterEnum("sonm.TaskStatusReply_Status", TaskStatusReply_Status_name, TaskStatusReply_Status_value)
	proto.RegisterEnum("sonm.TaskLogsRequest_Type", TaskLogsRequest_Type_name, TaskLogsRequest_Type_value)
}

func init() { proto.RegisterFile("insonmnia.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 806 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0x5d, 0x8f, 0xea, 0x44,
	0x18, 0x3e, 0xa5, 0x85, 0x2d, 0x2f, 0x0b, 0x74, 0x47, 0x3d, 0x21, 0xe4, 0x5c, 0x60, 0xd7, 0x0b,
	0x36, 0x9a, 0xc6, 0xa0, 0x17, 0xe6, 0x5c, 0x98, 0xb8, 0xc0, 0x71, 0xc9, 0x39, 0x5b, 0xc8, 0x00,
	0x39, 0xc6, 0xbb, 0x11, 0xc6, 0xb5, 0xa1, 0xed, 0xd4, 0x76, 0xea, 0xd2, 0xff, 0xe0, 0x9f, 0x30,
	0xf1, 0xc7, 0x79, 0xed, 0xbd, 0x89, 0x99, 0x8f, 0x96, 0xa2, 0x9c, 0xbb, 0x79, 0xde, 0xe7, 0x99,
	0x79, 0x3f, 0xfb, 0x16, 0xfa, 0x41, 0x9c, 0xb1, 0x38, 0x8a, 0x03, 0xe2, 0x25, 0x29, 0xe3, 0x0c,
	0x59, 0x02, 0xba, 0x5d, 0xe8, 0xac, 0x82, 0xf8, 0x09, 0xd3, 0x5f, 0x73, 0x9a, 0x71, 0xf7, 0x16,
	0xda, 0x0a, 0x26, 0x61, 0x81, 0x5e, 0x42, 0x2b, 0xe3, 0x84, 0xe7, 0xd9, 0xc0, 0x18, 0x19, 0xe3,
	0x36, 0xd6, 0xc8, 0x1d, 0x81, 0x3d, 0x5d, 0x6d, 0xb7, 0x19, 0x79, 0xa2, 0xe8, 0x63, 0x68, 0x72,
	0xc6, 0x49, 0x28, 0x25, 0x16, 0x56, 0xc0, 0xbd, 0x83, 0xce, 0x23, 0x8d, 0x58, 0x5a, 0x28, 0xd1,
	0x10, 0xec, 0x88, 0x1c, 0xe5, 0x59, 0xeb, 0x2a, 0xec, 0xfe, 0x6d, 0xc0, 0xb5, 0x4f, 0xf9, 0x33,
	0x4b, 0x0f, 0x4a, 0x3c, 0x80, 0x2b, 0x7e, 0xbc, 0x2f, 0x38, 0xcd, 0xb4, 0xb6, 0x84, 0x82, 0x49,
	0x35, 0xd3, 0x50, 0x8c, 0x86, 0xe8, 0x15, 0xb4, 0xf9, 0x71, 0x45, 0x76, 0x07, 0xca, 0xb3, 0x81,
	0x29, 0xb9, 0x93, 0x41, 0xb0, 0x69, 0xc5, 0x5a, 0x8a, 0xad, 0x0c, 0x22, 0x38, 0x7e, 0x9c, 0xa7,
	0x29, 0x4b, 0xb3, 0x41, 0x53, 0x05, 0x57, 0x62, 0xc1, 0xa5, 0x25, 0xd7, 0x52, 0x5c, 0x89, 0x95,
	0xcf, 0x59, 0xca, 0x92, 0x84, 0xee, 0x07, 0x57, 0xa5, 0x4f, 0x6d, 0x50, 0x3e, 0x4b, 0xd6, 0x2e,
	0x7d, 0x6a, 0x83, 0xfb, 0x97, 0x01, 0x5d, 0x4c, 0x33, 0x96, 0xa7, 0x3b, 0xaa, 0xb2, 0x1e, 0x81,
	0xb9, 0x4b, 0x72, 0x99, 0x71, 0x67, 0xd2, 0xf3, 0x44, 0x6f, 0xbc, 0xb2, 0xc8, 0x58, 0x50, 0xe8,
	0x0e, 0x5a, 0x91, 0xac, 0xa9, 0x4c, 0xbe, 0x33, 0xb9, 0x51, 0xa2, 0x5a, 0x9d, 0xb1, 0x16, 0xa0,
	0xd7, 0x70, 0x15, 0xab, 0x92, 0x0e, 0xcc, 0x91, 0x39, 0xee, 0x4c, 0x46, 0x4a, 0x7b, 0xe6, 0xd2,
	0xd3, 0x55, 0x9f, 0xc7, 0x3c, 0x2d, 0x70, 0x79, 0x61, 0xe8, 0x57, 0xed, 0x90, 0x04, 0x72, 0xc0,
	0x3c, 0xd0, 0x42, 0x4f, 0x80, 0x38, 0xa2, 0x31, 0x34, 0x7f, 0x23, 0x61, 0x4e, 0x75, 0x1c, 0x48,
	0xbd, 0x5d, 0xef, 0x21, 0x56, 0x82, 0xd7, 0x8d, 0x6f, 0x0c, 0xf7, 0x77, 0x03, 0xda, 0x8b, 0xf8,
	0x67, 0xa6, 0x46, 0xea, 0x4b, 0x68, 0xe6, 0x7a, 0x0c, 0x44, 0x5c, 0x43, 0x75, 0xb7, 0xe2, 0x3d,
	0x79, 0x5d, 0x45, 0xa4, 0x84, 0xc3, 0x47, 0x80, 0x93, 0xf1, 0x42, 0x34, 0x77, 0xe7, 0xd1, 0x7c,
	0x74, 0x21, 0xd3, 0x7a, 0x38, 0x9f, 0x42, 0x7f, 0xcd, 0x59, 0xb2, 0x21, 0xd9, 0x41, 0xcf, 0x3c,
	0xea, 0x41, 0x23, 0xd8, 0xeb, 0x27, 0x1b, 0xc1, 0xde, 0xed, 0x43, 0xf7, 0x24, 0x49, 0xc2, 0xc2,
	0xbd, 0x85, 0x1b, 0x01, 0xd6, 0x72, 0xfa, 0x3f, 0x74, 0xeb, 0x0f, 0x03, 0xfa, 0x75, 0x95, 0xc8,
	0xf6, 0xeb, 0xb3, 0x0f, 0xa8, 0x37, 0x79, 0xa5, 0x82, 0xfb, 0x8f, 0xcc, 0xd3, 0xe7, 0xf2, 0xf3,
	0xfa, 0x01, 0x5a, 0xca, 0x82, 0x3a, 0x70, 0xb5, 0xf5, 0xdf, 0xfa, 0xcb, 0xf7, 0xbe, 0xf3, 0x02,
	0x5d, 0x83, 0xbd, 0x5e, 0x2d, 0x97, 0xef, 0x16, 0xfe, 0xf7, 0x8e, 0xa1, 0xd0, 0x77, 0xef, 0x7d,
	0x81, 0x1a, 0x42, 0x88, 0xb7, 0xbe, 0x04, 0xa6, 0xa0, 0xde, 0x2c, 0xfc, 0xc5, 0xfa, 0x61, 0x3e,
	0x73, 0x2c, 0x04, 0xd0, 0xba, 0xc7, 0xcb, 0xb7, 0x73, 0xdf, 0x69, 0xba, 0x7f, 0x1a, 0xd0, 0x53,
	0x4f, 0x3f, 0x92, 0x44, 0x85, 0xf8, 0x2d, 0xd8, 0xca, 0xad, 0xfc, 0xdc, 0x44, 0x4f, 0x5c, 0x15,
	0xe4, 0xb9, 0x4e, 0x43, 0x9a, 0xa9, 0xde, 0x54, 0x77, 0x86, 0x58, 0x14, 0xab, 0x46, 0x5d, 0xe8,
	0xd0, 0xe7, 0xe7, 0x1d, 0xfa, 0xe4, 0x62, 0x11, 0xea, 0x3d, 0x7a, 0x00, 0x34, 0x65, 0x31, 0x27,
	0x41, 0x4c, 0xd3, 0xb2, 0x91, 0x99, 0xd8, 0x46, 0x7a, 0xfe, 0xc5, 0xdb, 0x66, 0x35, 0xec, 0x43,
	0xb0, 0x63, 0x12, 0xb3, 0xe9, 0x6a, 0xab, 0xd6, 0x82, 0x89, 0x2b, 0xec, 0xfe, 0x08, 0x2f, 0xeb,
	0x2f, 0x71, 0x92, 0xf2, 0x15, 0x0b, 0x83, 0x5d, 0x81, 0x10, 0x58, 0x31, 0x89, 0xa8, 0x8e, 0x53,
	0x9e, 0xd1, 0x17, 0x70, 0x13, 0x91, 0x63, 0x10, 0xe5, 0x11, 0xa6, 0x3c, 0x2d, 0xa6, 0x2c, 0x8f,
	0xb9, 0x7c, 0xb2, 0x8b, 0xff, 0x4f, 0xb8, 0xff, 0xe8, 0x86, 0xbf, 0x63, 0x4f, 0xd5, 0x50, 0x78,
	0x60, 0xf1, 0x22, 0xa1, 0xba, 0xdd, 0xc3, 0x53, 0xa6, 0x35, 0x91, 0xb7, 0x29, 0x12, 0x8a, 0xa5,
	0x4e, 0x0f, 0x51, 0xa3, 0x1c, 0x22, 0xb1, 0x4d, 0xb3, 0x20, 0xde, 0x51, 0xb9, 0xc3, 0xda, 0x58,
	0x01, 0xf4, 0x19, 0x74, 0xc9, 0x7e, 0xbf, 0x09, 0x22, 0x91, 0x41, 0x94, 0xa8, 0x1d, 0x66, 0xe3,
	0x73, 0xa3, 0xa8, 0xcf, 0x1b, 0x16, 0x86, 0xec, 0x59, 0x6e, 0x31, 0x1b, 0x6b, 0x24, 0x32, 0xdd,
	0x90, 0x20, 0x94, 0xfb, 0xab, 0x8d, 0xe5, 0x59, 0x6c, 0xd2, 0x19, 0xe5, 0x24, 0x08, 0x33, 0xb9,
	0xb9, 0x6c, 0x5c, 0x42, 0x77, 0x0c, 0x96, 0x88, 0x4f, 0x8c, 0xcd, 0x7a, 0x33, 0x5b, 0x6e, 0x37,
	0xce, 0x0b, 0x7d, 0x9e, 0x63, 0xec, 0x18, 0xc8, 0x06, 0xeb, 0x7e, 0xb9, 0x79, 0x70, 0x1a, 0xee,
	0x2d, 0x74, 0xcb, 0xcc, 0xa6, 0xbf, 0xe4, 0xf1, 0x41, 0x38, 0xda, 0x13, 0x4e, 0x64, 0xf2, 0xd7,
	0x58, 0x9e, 0x7f, 0x6a, 0xc9, 0x7f, 0xcd, 0x57, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x8d,
	0xd3, 0xfc, 0x7e, 0x06, 0x00, 0x00,
}