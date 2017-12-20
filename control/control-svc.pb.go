// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/paralin/go-p2pd/control/control-svc.proto

/*
Package control is a generated protocol buffer package.

It is generated from these files:
	github.com/paralin/go-p2pd/control/control-svc.proto

It has these top-level messages:
	CreateNodeRequest
	CreateNodeResponse
*/
package control

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

// CreateNodeRequest is the argument to CreateNode.
type CreateNodeRequest struct {
	// NodeId is the desired ID for the node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
}

func (m *CreateNodeRequest) Reset()                    { *m = CreateNodeRequest{} }
func (m *CreateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeRequest) ProtoMessage()               {}
func (*CreateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateNodeRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

// CreateNodeResponse is the response to CreateNode.
type CreateNodeResponse struct {
	// NodePeerId is the peer ID of the new node.
	NodePeerId string `protobuf:"bytes,1,opt,name=node_peer_id,json=nodePeerId" json:"node_peer_id,omitempty"`
}

func (m *CreateNodeResponse) Reset()                    { *m = CreateNodeResponse{} }
func (m *CreateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeResponse) ProtoMessage()               {}
func (*CreateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateNodeResponse) GetNodePeerId() string {
	if m != nil {
		return m.NodePeerId
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateNodeRequest)(nil), "control.CreateNodeRequest")
	proto.RegisterType((*CreateNodeResponse)(nil), "control.CreateNodeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ControlService service

type ControlServiceClient interface {
	// CreateNode creates a new node.
	CreateNode(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*CreateNodeResponse, error)
}

type controlServiceClient struct {
	cc *grpc.ClientConn
}

func NewControlServiceClient(cc *grpc.ClientConn) ControlServiceClient {
	return &controlServiceClient{cc}
}

func (c *controlServiceClient) CreateNode(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*CreateNodeResponse, error) {
	out := new(CreateNodeResponse)
	err := grpc.Invoke(ctx, "/control.ControlService/CreateNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ControlService service

type ControlServiceServer interface {
	// CreateNode creates a new node.
	CreateNode(context.Context, *CreateNodeRequest) (*CreateNodeResponse, error)
}

func RegisterControlServiceServer(s *grpc.Server, srv ControlServiceServer) {
	s.RegisterService(&_ControlService_serviceDesc, srv)
}

func _ControlService_CreateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).CreateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control.ControlService/CreateNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).CreateNode(ctx, req.(*CreateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ControlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "control.ControlService",
	HandlerType: (*ControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNode",
			Handler:    _ControlService_CreateNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/paralin/go-p2pd/control/control-svc.proto",
}

func init() {
	proto.RegisterFile("github.com/paralin/go-p2pd/control/control-svc.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0xc9, 0xcc, 0xd3, 0x4f,
	0xcf, 0xd7, 0x2d, 0x30, 0x2a, 0x48, 0xd1, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0xca, 0xcf, 0x81, 0xd1,
	0xba, 0xc5, 0x65, 0xc9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0x21, 0x25, 0x1d,
	0x2e, 0x41, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54, 0xbf, 0xfc, 0x94, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x71, 0x2e, 0xf6, 0xbc, 0xfc, 0x94, 0xd4, 0xf8, 0xcc, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x36, 0x10, 0xd7, 0x33, 0x45, 0xc9, 0x8c, 0x4b, 0x08, 0x59, 0x75,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x02, 0x17, 0x0f, 0x58, 0x79, 0x41, 0x6a, 0x6a, 0x11,
	0x42, 0x0f, 0x17, 0x48, 0x2c, 0x20, 0x35, 0xb5, 0xc8, 0x33, 0xc5, 0x28, 0x92, 0x8b, 0xcf, 0x19,
	0x62, 0x61, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x3b, 0x17, 0x17, 0xc2, 0x24, 0x21,
	0x29, 0x3d, 0xa8, 0x7b, 0xf4, 0x30, 0x1c, 0x23, 0x25, 0x8d, 0x55, 0x0e, 0x62, 0xb5, 0x12, 0x43,
	0x12, 0x1b, 0xd8, 0x43, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0x25, 0xf3, 0xab, 0x08,
	0x01, 0x00, 0x00,
}
