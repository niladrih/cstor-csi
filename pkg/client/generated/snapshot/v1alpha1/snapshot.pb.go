/*
Copyright © 2018-2019 The OpenEBS Authors

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

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cstorvolume.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type VolumeSnapCreateRequest struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Volume               string   `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
	Snapname             string   `protobuf:"bytes,3,opt,name=snapname,proto3" json:"snapname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumeSnapCreateRequest) Reset()         { *m = VolumeSnapCreateRequest{} }
func (m *VolumeSnapCreateRequest) String() string { return proto.CompactTextString(m) }
func (*VolumeSnapCreateRequest) ProtoMessage()    {}
func (*VolumeSnapCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0200b336e961136a, []int{0}
}

func (m *VolumeSnapCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeSnapCreateRequest.Unmarshal(m, b)
}
func (m *VolumeSnapCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeSnapCreateRequest.Marshal(b, m, deterministic)
}
func (m *VolumeSnapCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeSnapCreateRequest.Merge(m, src)
}
func (m *VolumeSnapCreateRequest) XXX_Size() int {
	return xxx_messageInfo_VolumeSnapCreateRequest.Size(m)
}
func (m *VolumeSnapCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeSnapCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeSnapCreateRequest proto.InternalMessageInfo

func (m *VolumeSnapCreateRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *VolumeSnapCreateRequest) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *VolumeSnapCreateRequest) GetSnapname() string {
	if m != nil {
		return m.Snapname
	}
	return ""
}

type VolumeSnapCreateResponse struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Status               []byte   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumeSnapCreateResponse) Reset()         { *m = VolumeSnapCreateResponse{} }
func (m *VolumeSnapCreateResponse) String() string { return proto.CompactTextString(m) }
func (*VolumeSnapCreateResponse) ProtoMessage()    {}
func (*VolumeSnapCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0200b336e961136a, []int{1}
}

func (m *VolumeSnapCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeSnapCreateResponse.Unmarshal(m, b)
}
func (m *VolumeSnapCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeSnapCreateResponse.Marshal(b, m, deterministic)
}
func (m *VolumeSnapCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeSnapCreateResponse.Merge(m, src)
}
func (m *VolumeSnapCreateResponse) XXX_Size() int {
	return xxx_messageInfo_VolumeSnapCreateResponse.Size(m)
}
func (m *VolumeSnapCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeSnapCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeSnapCreateResponse proto.InternalMessageInfo

func (m *VolumeSnapCreateResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *VolumeSnapCreateResponse) GetStatus() []byte {
	if m != nil {
		return m.Status
	}
	return nil
}

type VolumeSnapDeleteRequest struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Volume               string   `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
	Snapname             string   `protobuf:"bytes,3,opt,name=snapname,proto3" json:"snapname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumeSnapDeleteRequest) Reset()         { *m = VolumeSnapDeleteRequest{} }
func (m *VolumeSnapDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*VolumeSnapDeleteRequest) ProtoMessage()    {}
func (*VolumeSnapDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0200b336e961136a, []int{2}
}

func (m *VolumeSnapDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeSnapDeleteRequest.Unmarshal(m, b)
}
func (m *VolumeSnapDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeSnapDeleteRequest.Marshal(b, m, deterministic)
}
func (m *VolumeSnapDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeSnapDeleteRequest.Merge(m, src)
}
func (m *VolumeSnapDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_VolumeSnapDeleteRequest.Size(m)
}
func (m *VolumeSnapDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeSnapDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeSnapDeleteRequest proto.InternalMessageInfo

func (m *VolumeSnapDeleteRequest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *VolumeSnapDeleteRequest) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *VolumeSnapDeleteRequest) GetSnapname() string {
	if m != nil {
		return m.Snapname
	}
	return ""
}

type VolumeSnapDeleteResponse struct {
	Version              int32    `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Status               []byte   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumeSnapDeleteResponse) Reset()         { *m = VolumeSnapDeleteResponse{} }
func (m *VolumeSnapDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*VolumeSnapDeleteResponse) ProtoMessage()    {}
func (*VolumeSnapDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0200b336e961136a, []int{3}
}

func (m *VolumeSnapDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeSnapDeleteResponse.Unmarshal(m, b)
}
func (m *VolumeSnapDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeSnapDeleteResponse.Marshal(b, m, deterministic)
}
func (m *VolumeSnapDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeSnapDeleteResponse.Merge(m, src)
}
func (m *VolumeSnapDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_VolumeSnapDeleteResponse.Size(m)
}
func (m *VolumeSnapDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeSnapDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeSnapDeleteResponse proto.InternalMessageInfo

func (m *VolumeSnapDeleteResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *VolumeSnapDeleteResponse) GetStatus() []byte {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*VolumeSnapCreateRequest)(nil), "v1alpha1.VolumeSnapCreateRequest")
	proto.RegisterType((*VolumeSnapCreateResponse)(nil), "v1alpha1.VolumeSnapCreateResponse")
	proto.RegisterType((*VolumeSnapDeleteRequest)(nil), "v1alpha1.VolumeSnapDeleteRequest")
	proto.RegisterType((*VolumeSnapDeleteResponse)(nil), "v1alpha1.VolumeSnapDeleteResponse")
}

func init() { proto.RegisterFile("cstorvolume.proto", fileDescriptor_0200b336e961136a) }

var fileDescriptor_0200b336e961136a = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x6f, 0x15, 0xcf, 0x73, 0x10, 0xc1, 0x14, 0x1a, 0xb6, 0x3a, 0x53, 0x5d, 0xb5, 0x70,
	0xfa, 0x13, 0xb4, 0xb4, 0x8a, 0x60, 0x3f, 0xea, 0xa0, 0xc2, 0xee, 0x24, 0x66, 0x92, 0xfd, 0xb9,
	0xfe, 0x16, 0x21, 0x7b, 0x11, 0xd6, 0xd5, 0x2d, 0x2c, 0xae, 0x7c, 0xbc, 0x90, 0xef, 0xe5, 0xbd,
	0xc0, 0xf9, 0xb3, 0x44, 0x17, 0x7a, 0xd7, 0xa6, 0x8e, 0x1a, 0x1f, 0x5c, 0x74, 0x6a, 0xd5, 0x6f,
	0xb1, 0xf5, 0x6f, 0xb8, 0x35, 0xaf, 0x70, 0xf9, 0x98, 0x9d, 0x07, 0x46, 0x7f, 0x1b, 0x08, 0x23,
	0x59, 0xfa, 0x48, 0x24, 0x51, 0x69, 0x38, 0xee, 0x29, 0xc8, 0xbb, 0x63, 0x5d, 0xad, 0xab, 0xcd,
	0x91, 0x2d, 0x52, 0x5d, 0xc0, 0x72, 0xb8, 0x4e, 0x1f, 0xac, 0xab, 0xcd, 0x89, 0xdd, 0x29, 0x55,
	0xc3, 0x4a, 0x18, 0x3d, 0x63, 0x47, 0xfa, 0x30, 0x3b, 0xdf, 0xda, 0xdc, 0x83, 0x9e, 0x82, 0xc4,
	0x3b, 0x16, 0x9a, 0x27, 0x49, 0xc4, 0x98, 0x24, 0x93, 0x4e, 0xed, 0x4e, 0x8d, 0x63, 0xdf, 0x51,
	0x4b, 0x7b, 0x89, 0x5d, 0x40, 0xff, 0x8d, 0x7d, 0xfd, 0x59, 0xc1, 0x99, 0x4d, 0x9c, 0x2b, 0x70,
	0x5d, 0x87, 0xfc, 0xa2, 0x08, 0x6a, 0x9b, 0xf8, 0x67, 0x35, 0xc5, 0xbd, 0x6a, 0xca, 0x52, 0xcd,
	0x1f, 0x33, 0xd5, 0x66, 0xee, 0xc8, 0x90, 0xd4, 0x2c, 0x26, 0x98, 0xe1, 0x29, 0xb3, 0x98, 0x51,
	0xad, 0xbf, 0x63, 0xc6, 0x85, 0x98, 0xc5, 0xd3, 0x32, 0xff, 0xaf, 0x9b, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x56, 0xaf, 0xc4, 0x1c, 0x74, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RunSnapCommandClient is the client API for RunSnapCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RunSnapCommandClient interface {
	RunVolumeSnapCreateCommand(ctx context.Context, in *VolumeSnapCreateRequest, opts ...grpc.CallOption) (*VolumeSnapCreateResponse, error)
	RunVolumeSnapDeleteCommand(ctx context.Context, in *VolumeSnapDeleteRequest, opts ...grpc.CallOption) (*VolumeSnapDeleteResponse, error)
}

type runSnapCommandClient struct {
	cc *grpc.ClientConn
}

func NewRunSnapCommandClient(cc *grpc.ClientConn) RunSnapCommandClient {
	return &runSnapCommandClient{cc}
}

func (c *runSnapCommandClient) RunVolumeSnapCreateCommand(ctx context.Context, in *VolumeSnapCreateRequest, opts ...grpc.CallOption) (*VolumeSnapCreateResponse, error) {
	out := new(VolumeSnapCreateResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.RunSnapCommand/RunVolumeSnapCreateCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runSnapCommandClient) RunVolumeSnapDeleteCommand(ctx context.Context, in *VolumeSnapDeleteRequest, opts ...grpc.CallOption) (*VolumeSnapDeleteResponse, error) {
	out := new(VolumeSnapDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.RunSnapCommand/RunVolumeSnapDeleteCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunSnapCommandServer is the server API for RunSnapCommand service.
type RunSnapCommandServer interface {
	RunVolumeSnapCreateCommand(context.Context, *VolumeSnapCreateRequest) (*VolumeSnapCreateResponse, error)
	RunVolumeSnapDeleteCommand(context.Context, *VolumeSnapDeleteRequest) (*VolumeSnapDeleteResponse, error)
}

func RegisterRunSnapCommandServer(s *grpc.Server, srv RunSnapCommandServer) {
	s.RegisterService(&_RunSnapCommand_serviceDesc, srv)
}

func _RunSnapCommand_RunVolumeSnapCreateCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeSnapCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunSnapCommandServer).RunVolumeSnapCreateCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.RunSnapCommand/RunVolumeSnapCreateCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunSnapCommandServer).RunVolumeSnapCreateCommand(ctx, req.(*VolumeSnapCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunSnapCommand_RunVolumeSnapDeleteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeSnapDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunSnapCommandServer).RunVolumeSnapDeleteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.RunSnapCommand/RunVolumeSnapDeleteCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunSnapCommandServer).RunVolumeSnapDeleteCommand(ctx, req.(*VolumeSnapDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RunSnapCommand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.RunSnapCommand",
	HandlerType: (*RunSnapCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunVolumeSnapCreateCommand",
			Handler:    _RunSnapCommand_RunVolumeSnapCreateCommand_Handler,
		},
		{
			MethodName: "RunVolumeSnapDeleteCommand",
			Handler:    _RunSnapCommand_RunVolumeSnapDeleteCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cstorvolume.proto",
}
