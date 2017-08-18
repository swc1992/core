// Code generated by protoc-gen-go. DO NOT EDIT.
// source: miner.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MinerInfoRequest struct {
}

func (m *MinerInfoRequest) Reset()                    { *m = MinerInfoRequest{} }
func (m *MinerInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerInfoRequest) ProtoMessage()               {}
func (*MinerInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type MinerHandshakeRequest struct {
	Hub string `protobuf:"bytes,1,opt,name=hub" json:"hub,omitempty"`
}

func (m *MinerHandshakeRequest) Reset()                    { *m = MinerHandshakeRequest{} }
func (m *MinerHandshakeRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerHandshakeRequest) ProtoMessage()               {}
func (*MinerHandshakeRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *MinerHandshakeRequest) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

type MinerHandshakeReply struct {
	Miner        string        `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
	Capabilities *Capabilities `protobuf:"bytes,2,opt,name=capabilities" json:"capabilities,omitempty"`
}

func (m *MinerHandshakeReply) Reset()                    { *m = MinerHandshakeReply{} }
func (m *MinerHandshakeReply) String() string            { return proto.CompactTextString(m) }
func (*MinerHandshakeReply) ProtoMessage()               {}
func (*MinerHandshakeReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *MinerHandshakeReply) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *MinerHandshakeReply) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

type MinerStartRequest struct {
	Id            string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Registry      string                  `protobuf:"bytes,2,opt,name=registry" json:"registry,omitempty"`
	Image         string                  `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Auth          string                  `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	RestartPolicy *ContainerRestartPolicy `protobuf:"bytes,5,opt,name=restartPolicy" json:"restartPolicy,omitempty"`
	Resources     *ContainerResources     `protobuf:"bytes,6,opt,name=resources" json:"resources,omitempty"`
	PublicKeyData string                  `protobuf:"bytes,7,opt,name=PublicKeyData" json:"PublicKeyData,omitempty"`
}

func (m *MinerStartRequest) Reset()                    { *m = MinerStartRequest{} }
func (m *MinerStartRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerStartRequest) ProtoMessage()               {}
func (*MinerStartRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *MinerStartRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MinerStartRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *MinerStartRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *MinerStartRequest) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *MinerStartRequest) GetRestartPolicy() *ContainerRestartPolicy {
	if m != nil {
		return m.RestartPolicy
	}
	return nil
}

func (m *MinerStartRequest) GetResources() *ContainerResources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *MinerStartRequest) GetPublicKeyData() string {
	if m != nil {
		return m.PublicKeyData
	}
	return ""
}

type MinerStartReply struct {
	Container string                          `protobuf:"bytes,1,opt,name=container" json:"container,omitempty"`
	Ports     map[string]*MinerStartReplyPort `protobuf:"bytes,2,rep,name=Ports" json:"Ports,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MinerStartReply) Reset()                    { *m = MinerStartReply{} }
func (m *MinerStartReply) String() string            { return proto.CompactTextString(m) }
func (*MinerStartReply) ProtoMessage()               {}
func (*MinerStartReply) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *MinerStartReply) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *MinerStartReply) GetPorts() map[string]*MinerStartReplyPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

type MinerStartReplyPort struct {
	IP   string `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
}

func (m *MinerStartReplyPort) Reset()                    { *m = MinerStartReplyPort{} }
func (m *MinerStartReplyPort) String() string            { return proto.CompactTextString(m) }
func (*MinerStartReplyPort) ProtoMessage()               {}
func (*MinerStartReplyPort) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4, 0} }

func (m *MinerStartReplyPort) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *MinerStartReplyPort) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type MinerStatusMapRequest struct {
}

func (m *MinerStatusMapRequest) Reset()                    { *m = MinerStatusMapRequest{} }
func (m *MinerStatusMapRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerStatusMapRequest) ProtoMessage()               {}
func (*MinerStatusMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func init() {
	proto.RegisterType((*MinerInfoRequest)(nil), "sonm.MinerInfoRequest")
	proto.RegisterType((*MinerHandshakeRequest)(nil), "sonm.MinerHandshakeRequest")
	proto.RegisterType((*MinerHandshakeReply)(nil), "sonm.MinerHandshakeReply")
	proto.RegisterType((*MinerStartRequest)(nil), "sonm.MinerStartRequest")
	proto.RegisterType((*MinerStartReply)(nil), "sonm.MinerStartReply")
	proto.RegisterType((*MinerStartReplyPort)(nil), "sonm.MinerStartReply.port")
	proto.RegisterType((*MinerStatusMapRequest)(nil), "sonm.MinerStatusMapRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Miner service

type MinerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Info(ctx context.Context, in *MinerInfoRequest, opts ...grpc.CallOption) (*InfoReply, error)
	Handshake(ctx context.Context, in *MinerHandshakeRequest, opts ...grpc.CallOption) (*MinerHandshakeReply, error)
	Start(ctx context.Context, in *MinerStartRequest, opts ...grpc.CallOption) (*MinerStartReply, error)
	Stop(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error)
	TasksStatus(ctx context.Context, opts ...grpc.CallOption) (Miner_TasksStatusClient, error)
	TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Miner_TaskLogsClient, error)
}

type minerClient struct {
	cc *grpc.ClientConn
}

func NewMinerClient(cc *grpc.ClientConn) MinerClient {
	return &minerClient{cc}
}

func (c *minerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Info(ctx context.Context, in *MinerInfoRequest, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Handshake(ctx context.Context, in *MinerHandshakeRequest, opts ...grpc.CallOption) (*MinerHandshakeReply, error) {
	out := new(MinerHandshakeReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Handshake", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Start(ctx context.Context, in *MinerStartRequest, opts ...grpc.CallOption) (*MinerStartReply, error) {
	out := new(MinerStartReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Stop(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error) {
	out := new(StopTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) TasksStatus(ctx context.Context, opts ...grpc.CallOption) (Miner_TasksStatusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[0], c.cc, "/sonm.Miner/TasksStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerTasksStatusClient{stream}
	return x, nil
}

type Miner_TasksStatusClient interface {
	Send(*MinerStatusMapRequest) error
	Recv() (*StatusMapReply, error)
	grpc.ClientStream
}

type minerTasksStatusClient struct {
	grpc.ClientStream
}

func (x *minerTasksStatusClient) Send(m *MinerStatusMapRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *minerTasksStatusClient) Recv() (*StatusMapReply, error) {
	m := new(StatusMapReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minerClient) TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Miner_TaskLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[1], c.cc, "/sonm.Miner/TaskLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerTaskLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Miner_TaskLogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type minerTaskLogsClient struct {
	grpc.ClientStream
}

func (x *minerTaskLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Miner service

type MinerServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Info(context.Context, *MinerInfoRequest) (*InfoReply, error)
	Handshake(context.Context, *MinerHandshakeRequest) (*MinerHandshakeReply, error)
	Start(context.Context, *MinerStartRequest) (*MinerStartReply, error)
	Stop(context.Context, *StopTaskRequest) (*StopTaskReply, error)
	TasksStatus(Miner_TasksStatusServer) error
	TaskLogs(*TaskLogsRequest, Miner_TaskLogsServer) error
}

func RegisterMinerServer(s *grpc.Server, srv MinerServer) {
	s.RegisterService(&_Miner_serviceDesc, srv)
}

func _Miner_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Info(ctx, req.(*MinerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerHandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Handshake(ctx, req.(*MinerHandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Start(ctx, req.(*MinerStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Stop(ctx, req.(*StopTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_TasksStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MinerServer).TasksStatus(&minerTasksStatusServer{stream})
}

type Miner_TasksStatusServer interface {
	Send(*StatusMapReply) error
	Recv() (*MinerStatusMapRequest, error)
	grpc.ServerStream
}

type minerTasksStatusServer struct {
	grpc.ServerStream
}

func (x *minerTasksStatusServer) Send(m *StatusMapReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *minerTasksStatusServer) Recv() (*MinerStatusMapRequest, error) {
	m := new(MinerStatusMapRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Miner_TaskLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MinerServer).TaskLogs(m, &minerTaskLogsServer{stream})
}

type Miner_TaskLogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type minerTaskLogsServer struct {
	grpc.ServerStream
}

func (x *minerTaskLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Miner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Miner",
	HandlerType: (*MinerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Miner_Ping_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Miner_Info_Handler,
		},
		{
			MethodName: "Handshake",
			Handler:    _Miner_Handshake_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Miner_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Miner_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TasksStatus",
			Handler:       _Miner_TasksStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TaskLogs",
			Handler:       _Miner_TaskLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "miner.proto",
}

func init() { proto.RegisterFile("miner.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x5d, 0x6e, 0xd3, 0x40,
	0x18, 0xac, 0x13, 0xbb, 0x34, 0x5f, 0x28, 0x6d, 0xbf, 0xb6, 0xd4, 0x98, 0x3e, 0x44, 0x16, 0x0f,
	0x01, 0xa1, 0x28, 0x04, 0x14, 0x41, 0x1f, 0x29, 0x45, 0x44, 0x50, 0xc9, 0x72, 0x7a, 0x81, 0x8d,
	0xb3, 0x24, 0xab, 0x38, 0xb6, 0xf1, 0xae, 0x91, 0x7c, 0x06, 0xae, 0xc1, 0xe5, 0xb8, 0x05, 0xda,
	0x1f, 0x27, 0x8e, 0x9b, 0xa7, 0xec, 0xce, 0xce, 0xcc, 0x8e, 0xc7, 0x5f, 0x0c, 0xdd, 0x35, 0x4b,
	0x68, 0x3e, 0xc8, 0xf2, 0x54, 0xa4, 0x68, 0xf3, 0x34, 0x59, 0x7b, 0x27, 0x2c, 0x91, 0xbf, 0x09,
	0x23, 0x1a, 0xf6, 0x30, 0x22, 0x19, 0x99, 0xb1, 0x98, 0x09, 0x46, 0xb9, 0xc6, 0x7c, 0x84, 0xd3,
	0x7b, 0xa9, 0x9c, 0x24, 0x3f, 0xd3, 0x90, 0xfe, 0x2a, 0x28, 0x17, 0xfe, 0x6b, 0xb8, 0x54, 0xd8,
	0x37, 0x92, 0xcc, 0xf9, 0x92, 0xac, 0xa8, 0x39, 0xc0, 0x53, 0x68, 0x2f, 0x8b, 0x99, 0x6b, 0xf5,
	0xac, 0x7e, 0x27, 0x94, 0x4b, 0x3f, 0x82, 0xf3, 0x26, 0x35, 0x8b, 0x4b, 0xbc, 0x00, 0x47, 0xe5,
	0x31, 0x54, 0xbd, 0xc1, 0x31, 0x3c, 0xad, 0x27, 0x70, 0x5b, 0x3d, 0xab, 0xdf, 0x1d, 0xe1, 0x40,
	0xa6, 0x1c, 0xdc, 0xd6, 0x4e, 0xc2, 0x1d, 0x9e, 0xff, 0xa7, 0x05, 0x67, 0xea, 0x96, 0xa9, 0x20,
	0xb9, 0xa8, 0xc2, 0x3c, 0x83, 0x16, 0x9b, 0x9b, 0x0b, 0x5a, 0x6c, 0x8e, 0x1e, 0x1c, 0xe5, 0x74,
	0xc1, 0xb8, 0xc8, 0x4b, 0xe5, 0xdc, 0x09, 0x37, 0x7b, 0x99, 0x87, 0xad, 0xc9, 0x82, 0xba, 0x6d,
	0x9d, 0x47, 0x6d, 0x10, 0xc1, 0x26, 0x85, 0x58, 0xba, 0xb6, 0x02, 0xd5, 0x1a, 0x3f, 0xc3, 0x71,
	0x4e, 0xb9, 0xbc, 0x27, 0x48, 0x63, 0x16, 0x95, 0xae, 0xa3, 0x42, 0x5e, 0x9b, 0x90, 0x69, 0x22,
	0x88, 0x4c, 0x12, 0xd6, 0x39, 0xe1, 0xae, 0x04, 0xc7, 0xd0, 0xc9, 0x29, 0x4f, 0x8b, 0x3c, 0xa2,
	0xdc, 0x3d, 0x54, 0x7a, 0xf7, 0xb1, 0x5e, 0x9f, 0x87, 0x5b, 0x2a, 0xbe, 0x82, 0xe3, 0xa0, 0x98,
	0xc5, 0x2c, 0xfa, 0x4e, 0xcb, 0x2f, 0x44, 0x10, 0xf7, 0x89, 0x0a, 0xb6, 0x0b, 0xfa, 0xff, 0x2c,
	0x38, 0xa9, 0xb7, 0x21, 0xfb, 0xbe, 0x86, 0x4e, 0x54, 0x59, 0x9b, 0x4a, 0xb6, 0x00, 0x8e, 0xc1,
	0x09, 0xd2, 0x5c, 0xc8, 0xc2, 0xdb, 0xfd, 0xee, 0xa8, 0xa7, 0xb3, 0x34, 0x3c, 0x06, 0x8a, 0x72,
	0x97, 0x88, 0xbc, 0x0c, 0x35, 0xdd, 0x7b, 0x03, 0x76, 0x96, 0xe6, 0xaa, 0xe9, 0x49, 0x50, 0x35,
	0x3d, 0x09, 0x64, 0x6f, 0x12, 0x37, 0x2d, 0xab, 0xb5, 0xf7, 0x00, 0xb0, 0x35, 0x90, 0x83, 0xb2,
	0xa2, 0x65, 0x35, 0x28, 0x2b, 0x5a, 0xe2, 0x10, 0x9c, 0xdf, 0x24, 0x2e, 0xa8, 0x79, 0xe9, 0xde,
	0xfe, 0x0c, 0xd2, 0x2a, 0xd4, 0xc4, 0x9b, 0xd6, 0x47, 0xcb, 0xbf, 0x32, 0x93, 0x38, 0x15, 0x44,
	0x14, 0xfc, 0x9e, 0x64, 0xe6, 0xe5, 0x8f, 0xfe, 0xb6, 0xc1, 0x51, 0x27, 0xf8, 0x16, 0xec, 0x80,
	0x25, 0x0b, 0x3c, 0xd3, 0x8e, 0x72, 0x6d, 0x48, 0xde, 0x49, 0x1d, 0xca, 0xe2, 0xd2, 0x3f, 0xc0,
	0x77, 0x60, 0xcb, 0x49, 0xc7, 0xe7, 0xb5, 0xfb, 0x6b, 0xa3, 0x5f, 0x49, 0x34, 0xa4, 0x25, 0x77,
	0xd0, 0xd9, 0x4c, 0x37, 0xbe, 0xac, 0xe9, 0x9a, 0x7f, 0x0f, 0xef, 0xc5, 0xfe, 0x43, 0x6d, 0xf3,
	0x09, 0x1c, 0xf5, 0xa0, 0x78, 0xf5, 0xf8, 0xd1, 0xb5, 0xfc, 0x72, 0x6f, 0x27, 0xfe, 0x01, 0x7e,
	0x00, 0x7b, 0x2a, 0xd2, 0x0c, 0x0d, 0x41, 0xae, 0x1f, 0x08, 0x5f, 0x55, 0xba, 0xf3, 0x26, 0xac,
	0x55, 0x5f, 0xa1, 0x2b, 0xb7, 0x5c, 0x77, 0xb7, 0x93, 0xbc, 0x59, 0xa7, 0x77, 0x51, 0x59, 0x6c,
	0x70, 0xe5, 0xd1, 0xb7, 0x86, 0x16, 0xde, 0xc0, 0x91, 0xf4, 0xf9, 0x91, 0x2e, 0x78, 0x95, 0xa0,
	0xda, 0x37, 0x12, 0x54, 0xf0, 0xed, 0xb2, 0x48, 0x56, 0xfe, 0xc1, 0xd0, 0x9a, 0x1d, 0xaa, 0x8f,
	0xcc, 0xfb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xa4, 0x34, 0xf4, 0x9e, 0x04, 0x00, 0x00,
}