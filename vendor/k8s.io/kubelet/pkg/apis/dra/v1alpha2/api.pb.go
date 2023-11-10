/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package v1alpha2

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NodePrepareResourceRequest struct {
	// The ResourceClaim namespace (ResourceClaim.meta.Namespace).
	// This field is REQUIRED.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The UID of the Resource claim (ResourceClaim.meta.UUID).
	// This field is REQUIRED.
	ClaimUid string `protobuf:"bytes,2,opt,name=claim_uid,json=claimUid,proto3" json:"claim_uid,omitempty"`
	// The name of the Resource claim (ResourceClaim.meta.Name)
	// This field is REQUIRED.
	ClaimName string `protobuf:"bytes,3,opt,name=claim_name,json=claimName,proto3" json:"claim_name,omitempty"`
	// Resource handle (AllocationResult.ResourceHandles[*].Data)
	// This field is REQUIRED.
	ResourceHandle       string   `protobuf:"bytes,4,opt,name=resource_handle,json=resourceHandle,proto3" json:"resource_handle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodePrepareResourceRequest) Reset()      { *m = NodePrepareResourceRequest{} }
func (*NodePrepareResourceRequest) ProtoMessage() {}
func (*NodePrepareResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *NodePrepareResourceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodePrepareResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodePrepareResourceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodePrepareResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePrepareResourceRequest.Merge(m, src)
}
func (m *NodePrepareResourceRequest) XXX_Size() int {
	return m.Size()
}
func (m *NodePrepareResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePrepareResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodePrepareResourceRequest proto.InternalMessageInfo

func (m *NodePrepareResourceRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NodePrepareResourceRequest) GetClaimUid() string {
	if m != nil {
		return m.ClaimUid
	}
	return ""
}

func (m *NodePrepareResourceRequest) GetClaimName() string {
	if m != nil {
		return m.ClaimName
	}
	return ""
}

func (m *NodePrepareResourceRequest) GetResourceHandle() string {
	if m != nil {
		return m.ResourceHandle
	}
	return ""
}

type NodePrepareResourceResponse struct {
	// These are the additional devices that kubelet must
	// make available via the container runtime. A resource
	// may have zero or more devices.
	CdiDevices           []string `protobuf:"bytes,1,rep,name=cdi_devices,json=cdiDevices,proto3" json:"cdi_devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodePrepareResourceResponse) Reset()      { *m = NodePrepareResourceResponse{} }
func (*NodePrepareResourceResponse) ProtoMessage() {}
func (*NodePrepareResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *NodePrepareResourceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodePrepareResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodePrepareResourceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodePrepareResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodePrepareResourceResponse.Merge(m, src)
}
func (m *NodePrepareResourceResponse) XXX_Size() int {
	return m.Size()
}
func (m *NodePrepareResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodePrepareResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodePrepareResourceResponse proto.InternalMessageInfo

func (m *NodePrepareResourceResponse) GetCdiDevices() []string {
	if m != nil {
		return m.CdiDevices
	}
	return nil
}

type NodeUnprepareResourceRequest struct {
	// The ResourceClaim namespace (ResourceClaim.meta.Namespace).
	// This field is REQUIRED.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The UID of the Resource claim (ResourceClaim.meta.UUID).
	// This field is REQUIRED.
	ClaimUid string `protobuf:"bytes,2,opt,name=claim_uid,json=claimUid,proto3" json:"claim_uid,omitempty"`
	// The name of the Resource claim (ResourceClaim.meta.Name)
	// This field is REQUIRED.
	ClaimName string `protobuf:"bytes,3,opt,name=claim_name,json=claimName,proto3" json:"claim_name,omitempty"`
	// Resource handle (AllocationResult.ResourceHandles[*].Data)
	// This field is REQUIRED.
	ResourceHandle       string   `protobuf:"bytes,4,opt,name=resource_handle,json=resourceHandle,proto3" json:"resource_handle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeUnprepareResourceRequest) Reset()      { *m = NodeUnprepareResourceRequest{} }
func (*NodeUnprepareResourceRequest) ProtoMessage() {}
func (*NodeUnprepareResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}
func (m *NodeUnprepareResourceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeUnprepareResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeUnprepareResourceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeUnprepareResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeUnprepareResourceRequest.Merge(m, src)
}
func (m *NodeUnprepareResourceRequest) XXX_Size() int {
	return m.Size()
}
func (m *NodeUnprepareResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeUnprepareResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeUnprepareResourceRequest proto.InternalMessageInfo

func (m *NodeUnprepareResourceRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *NodeUnprepareResourceRequest) GetClaimUid() string {
	if m != nil {
		return m.ClaimUid
	}
	return ""
}

func (m *NodeUnprepareResourceRequest) GetClaimName() string {
	if m != nil {
		return m.ClaimName
	}
	return ""
}

func (m *NodeUnprepareResourceRequest) GetResourceHandle() string {
	if m != nil {
		return m.ResourceHandle
	}
	return ""
}

type NodeUnprepareResourceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeUnprepareResourceResponse) Reset()      { *m = NodeUnprepareResourceResponse{} }
func (*NodeUnprepareResourceResponse) ProtoMessage() {}
func (*NodeUnprepareResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}
func (m *NodeUnprepareResourceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeUnprepareResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeUnprepareResourceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeUnprepareResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeUnprepareResourceResponse.Merge(m, src)
}
func (m *NodeUnprepareResourceResponse) XXX_Size() int {
	return m.Size()
}
func (m *NodeUnprepareResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeUnprepareResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeUnprepareResourceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NodePrepareResourceRequest)(nil), "v1alpha2.NodePrepareResourceRequest")
	proto.RegisterType((*NodePrepareResourceResponse)(nil), "v1alpha2.NodePrepareResourceResponse")
	proto.RegisterType((*NodeUnprepareResourceRequest)(nil), "v1alpha2.NodeUnprepareResourceRequest")
	proto.RegisterType((*NodeUnprepareResourceResponse)(nil), "v1alpha2.NodeUnprepareResourceResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x52, 0xb1, 0x6e, 0xe2, 0x40,
	0x10, 0x65, 0x0f, 0x74, 0xc2, 0x7b, 0xd2, 0x9d, 0xb4, 0xa7, 0x93, 0x2c, 0x03, 0x06, 0x59, 0xdc,
	0x41, 0x73, 0xb6, 0x8e, 0x6b, 0xae, 0xba, 0x02, 0xa5, 0x48, 0x85, 0x22, 0x4b, 0x34, 0x69, 0xd0,
	0xda, 0x3b, 0x31, 0x1b, 0x6c, 0xef, 0xc6, 0x6b, 0x53, 0xe7, 0x13, 0xf2, 0x07, 0x51, 0xfe, 0x86,
	0x32, 0x25, 0x65, 0x70, 0x7e, 0x24, 0x62, 0x1d, 0x2b, 0x8a, 0x04, 0xa2, 0x4d, 0xb7, 0xf3, 0xe6,
	0xcd, 0xbc, 0x37, 0xb3, 0x83, 0x0d, 0x2a, 0xb9, 0x2b, 0x33, 0x91, 0x0b, 0xd2, 0x5e, 0xff, 0xa1,
	0xb1, 0x5c, 0xd2, 0x89, 0xf5, 0x3b, 0xe2, 0xf9, 0xb2, 0x08, 0xdc, 0x50, 0x24, 0x5e, 0x24, 0x22,
	0xe1, 0x69, 0x42, 0x50, 0x5c, 0xe9, 0x48, 0x07, 0xfa, 0x55, 0x15, 0x3a, 0xf7, 0x08, 0x5b, 0x33,
	0xc1, 0xe0, 0x22, 0x03, 0x49, 0x33, 0xf0, 0x41, 0x89, 0x22, 0x0b, 0xc1, 0x87, 0x9b, 0x02, 0x54,
	0x4e, 0xba, 0xd8, 0x48, 0x69, 0x02, 0x4a, 0xd2, 0x10, 0x4c, 0x34, 0x40, 0x63, 0xc3, 0x7f, 0x03,
	0x48, 0x07, 0x1b, 0x61, 0x4c, 0x79, 0xb2, 0x28, 0x38, 0x33, 0x3f, 0xe9, 0x6c, 0x5b, 0x03, 0x73,
	0xce, 0x48, 0x0f, 0xe3, 0x2a, 0xb9, 0xe7, 0x9b, 0xcd, 0xaa, 0x56, 0x23, 0x33, 0x9a, 0x00, 0x19,
	0xe1, 0x6f, 0xd9, 0xab, 0xd8, 0x62, 0x49, 0x53, 0x16, 0x83, 0xd9, 0xd2, 0x9c, 0xaf, 0x35, 0x7c,
	0xae, 0x51, 0xe7, 0x3f, 0xee, 0x1c, 0x34, 0xa8, 0xa4, 0x48, 0x15, 0x90, 0x3e, 0xfe, 0x12, 0x32,
	0xbe, 0x60, 0xb0, 0xe6, 0x21, 0x28, 0x13, 0x0d, 0x9a, 0x63, 0xc3, 0xc7, 0x21, 0xe3, 0x67, 0x15,
	0xe2, 0x3c, 0x20, 0xdc, 0xdd, 0x37, 0x98, 0xa7, 0xf2, 0xc3, 0xce, 0xd8, 0xc7, 0xbd, 0x23, 0x16,
	0xab, 0x29, 0x27, 0x5b, 0x84, 0x5b, 0x7b, 0x06, 0x61, 0xf8, 0xfb, 0x81, 0x6d, 0x90, 0xa1, 0x5b,
	0x1f, 0x80, 0x7b, 0xfc, 0x37, 0xad, 0x9f, 0x27, 0x58, 0x95, 0x98, 0xd3, 0x20, 0xd7, 0xf8, 0xc7,
	0x41, 0x3f, 0xe4, 0xd7, 0xfb, 0x0e, 0xc7, 0x76, 0x6a, 0x8d, 0x4e, 0xf2, 0x6a, 0xad, 0xe9, 0x74,
	0xb3, 0xb3, 0xd1, 0x76, 0x67, 0x37, 0x6e, 0x4b, 0x1b, 0x6d, 0x4a, 0x1b, 0x3d, 0x96, 0x36, 0x7a,
	0x2a, 0x6d, 0x74, 0xf7, 0x6c, 0x37, 0x2e, 0x87, 0xab, 0x7f, 0xca, 0xe5, 0xc2, 0x5b, 0x15, 0x01,
	0xc4, 0x90, 0x7b, 0x72, 0x15, 0x79, 0x54, 0x72, 0xe5, 0xb1, 0x8c, 0x7a, 0xb5, 0x46, 0xf0, 0x59,
	0x1f, 0xf3, 0xdf, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x2f, 0x77, 0x8e, 0x12, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	NodePrepareResource(ctx context.Context, in *NodePrepareResourceRequest, opts ...grpc.CallOption) (*NodePrepareResourceResponse, error)
	NodeUnprepareResource(ctx context.Context, in *NodeUnprepareResourceRequest, opts ...grpc.CallOption) (*NodeUnprepareResourceResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) NodePrepareResource(ctx context.Context, in *NodePrepareResourceRequest, opts ...grpc.CallOption) (*NodePrepareResourceResponse, error) {
	out := new(NodePrepareResourceResponse)
	err := c.cc.Invoke(ctx, "/v1alpha2.Node/NodePrepareResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) NodeUnprepareResource(ctx context.Context, in *NodeUnprepareResourceRequest, opts ...grpc.CallOption) (*NodeUnprepareResourceResponse, error) {
	out := new(NodeUnprepareResourceResponse)
	err := c.cc.Invoke(ctx, "/v1alpha2.Node/NodeUnprepareResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	NodePrepareResource(context.Context, *NodePrepareResourceRequest) (*NodePrepareResourceResponse, error)
	NodeUnprepareResource(context.Context, *NodeUnprepareResourceRequest) (*NodeUnprepareResourceResponse, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) NodePrepareResource(ctx context.Context, req *NodePrepareResourceRequest) (*NodePrepareResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodePrepareResource not implemented")
}
func (*UnimplementedNodeServer) NodeUnprepareResource(ctx context.Context, req *NodeUnprepareResourceRequest) (*NodeUnprepareResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeUnprepareResource not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_NodePrepareResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodePrepareResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).NodePrepareResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.Node/NodePrepareResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).NodePrepareResource(ctx, req.(*NodePrepareResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_NodeUnprepareResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeUnprepareResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).NodeUnprepareResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha2.Node/NodeUnprepareResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).NodeUnprepareResource(ctx, req.(*NodeUnprepareResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha2.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NodePrepareResource",
			Handler:    _Node_NodePrepareResource_Handler,
		},
		{
			MethodName: "NodeUnprepareResource",
			Handler:    _Node_NodeUnprepareResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *NodePrepareResourceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodePrepareResourceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodePrepareResourceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResourceHandle) > 0 {
		i -= len(m.ResourceHandle)
		copy(dAtA[i:], m.ResourceHandle)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ResourceHandle)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClaimName) > 0 {
		i -= len(m.ClaimName)
		copy(dAtA[i:], m.ClaimName)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ClaimName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClaimUid) > 0 {
		i -= len(m.ClaimUid)
		copy(dAtA[i:], m.ClaimUid)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ClaimUid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NodePrepareResourceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodePrepareResourceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodePrepareResourceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CdiDevices) > 0 {
		for iNdEx := len(m.CdiDevices) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.CdiDevices[iNdEx])
			copy(dAtA[i:], m.CdiDevices[iNdEx])
			i = encodeVarintApi(dAtA, i, uint64(len(m.CdiDevices[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *NodeUnprepareResourceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeUnprepareResourceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeUnprepareResourceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ResourceHandle) > 0 {
		i -= len(m.ResourceHandle)
		copy(dAtA[i:], m.ResourceHandle)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ResourceHandle)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClaimName) > 0 {
		i -= len(m.ClaimName)
		copy(dAtA[i:], m.ClaimName)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ClaimName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClaimUid) > 0 {
		i -= len(m.ClaimUid)
		copy(dAtA[i:], m.ClaimUid)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ClaimUid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NodeUnprepareResourceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeUnprepareResourceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeUnprepareResourceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NodePrepareResourceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ClaimUid)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ClaimName)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ResourceHandle)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func (m *NodePrepareResourceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CdiDevices) > 0 {
		for _, s := range m.CdiDevices {
			l = len(s)
			n += 1 + l + sovApi(uint64(l))
		}
	}
	return n
}

func (m *NodeUnprepareResourceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ClaimUid)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ClaimName)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ResourceHandle)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func (m *NodeUnprepareResourceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *NodePrepareResourceRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NodePrepareResourceRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`ClaimUid:` + fmt.Sprintf("%v", this.ClaimUid) + `,`,
		`ClaimName:` + fmt.Sprintf("%v", this.ClaimName) + `,`,
		`ResourceHandle:` + fmt.Sprintf("%v", this.ResourceHandle) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NodePrepareResourceResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NodePrepareResourceResponse{`,
		`CdiDevices:` + fmt.Sprintf("%v", this.CdiDevices) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NodeUnprepareResourceRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NodeUnprepareResourceRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`ClaimUid:` + fmt.Sprintf("%v", this.ClaimUid) + `,`,
		`ClaimName:` + fmt.Sprintf("%v", this.ClaimName) + `,`,
		`ResourceHandle:` + fmt.Sprintf("%v", this.ResourceHandle) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NodeUnprepareResourceResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NodeUnprepareResourceResponse{`,
		`}`,
	}, "")
	return s
}
func valueToStringApi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *NodePrepareResourceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodePrepareResourceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodePrepareResourceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimUid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimUid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceHandle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceHandle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodePrepareResourceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodePrepareResourceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodePrepareResourceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CdiDevices", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CdiDevices = append(m.CdiDevices, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeUnprepareResourceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeUnprepareResourceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeUnprepareResourceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimUid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimUid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceHandle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceHandle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeUnprepareResourceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeUnprepareResourceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeUnprepareResourceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
