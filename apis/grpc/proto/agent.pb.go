// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

package agentpb

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AgentStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentStatusRequest) Reset()         { *m = AgentStatusRequest{} }
func (m *AgentStatusRequest) String() string { return proto.CompactTextString(m) }
func (*AgentStatusRequest) ProtoMessage()    {}
func (*AgentStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_0f2881eb9c6d417c, []int{0}
}
func (m *AgentStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentStatusRequest.Unmarshal(m, b)
}
func (m *AgentStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentStatusRequest.Marshal(b, m, deterministic)
}
func (dst *AgentStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentStatusRequest.Merge(dst, src)
}
func (m *AgentStatusRequest) XXX_Size() int {
	return xxx_messageInfo_AgentStatusRequest.Size(m)
}
func (m *AgentStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AgentStatusRequest proto.InternalMessageInfo

type AgentStatusResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentStatusResponse) Reset()         { *m = AgentStatusResponse{} }
func (m *AgentStatusResponse) String() string { return proto.CompactTextString(m) }
func (*AgentStatusResponse) ProtoMessage()    {}
func (*AgentStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agent_0f2881eb9c6d417c, []int{1}
}
func (m *AgentStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentStatusResponse.Unmarshal(m, b)
}
func (m *AgentStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentStatusResponse.Marshal(b, m, deterministic)
}
func (dst *AgentStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentStatusResponse.Merge(dst, src)
}
func (m *AgentStatusResponse) XXX_Size() int {
	return xxx_messageInfo_AgentStatusResponse.Size(m)
}
func (m *AgentStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AgentStatusResponse proto.InternalMessageInfo

func (m *AgentStatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*AgentStatusRequest)(nil), "AgentStatusRequest")
	proto.RegisterType((*AgentStatusResponse)(nil), "AgentStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentStatusServiceClient is the client API for AgentStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentStatusServiceClient interface {
	Status(ctx context.Context, in *AgentStatusRequest, opts ...grpc.CallOption) (AgentStatusService_StatusClient, error)
}

type agentStatusServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentStatusServiceClient(cc *grpc.ClientConn) AgentStatusServiceClient {
	return &agentStatusServiceClient{cc}
}

func (c *agentStatusServiceClient) Status(ctx context.Context, in *AgentStatusRequest, opts ...grpc.CallOption) (AgentStatusService_StatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AgentStatusService_serviceDesc.Streams[0], "/AgentStatusService/Status", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentStatusServiceStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentStatusService_StatusClient interface {
	Recv() (*AgentStatusResponse, error)
	grpc.ClientStream
}

type agentStatusServiceStatusClient struct {
	grpc.ClientStream
}

func (x *agentStatusServiceStatusClient) Recv() (*AgentStatusResponse, error) {
	m := new(AgentStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentStatusServiceServer is the server API for AgentStatusService service.
type AgentStatusServiceServer interface {
	Status(*AgentStatusRequest, AgentStatusService_StatusServer) error
}

func RegisterAgentStatusServiceServer(s *grpc.Server, srv AgentStatusServiceServer) {
	s.RegisterService(&_AgentStatusService_serviceDesc, srv)
}

func _AgentStatusService_Status_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AgentStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentStatusServiceServer).Status(m, &agentStatusServiceStatusServer{stream})
}

type AgentStatusService_StatusServer interface {
	Send(*AgentStatusResponse) error
	grpc.ServerStream
}

type agentStatusServiceStatusServer struct {
	grpc.ServerStream
}

func (x *agentStatusServiceStatusServer) Send(m *AgentStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _AgentStatusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AgentStatusService",
	HandlerType: (*AgentStatusServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Status",
			Handler:       _AgentStatusService_Status_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "agent.proto",
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_agent_0f2881eb9c6d417c) }

var fileDescriptor_agent_0f2881eb9c6d417c = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0x4f, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x12, 0xe1, 0x12, 0x72, 0x04, 0x71, 0x83, 0x4b,
	0x12, 0x4b, 0x4a, 0x8b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x74, 0xb9, 0x84, 0x51,
	0x44, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0xc1, 0x22, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x91, 0x37, 0x8a, 0x21, 0xc1, 0xa9, 0x45, 0x65, 0x99,
	0xc9, 0xa9, 0x42, 0xa6, 0x5c, 0x6c, 0x10, 0x01, 0x21, 0x61, 0x3d, 0x4c, 0x3b, 0xa4, 0x44, 0xf4,
	0xb0, 0x58, 0x61, 0xc0, 0xe8, 0xc4, 0x19, 0xc5, 0x0e, 0x76, 0x60, 0x41, 0x52, 0x12, 0x1b, 0xd8,
	0x8d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x9c, 0x35, 0x01, 0xb2, 0x00, 0x00, 0x00,
}