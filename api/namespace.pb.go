// Code generated by protoc-gen-go. DO NOT EDIT.
// source: namespace.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type ListNamespacesResponse struct {
	Count                int32        `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Namespaces           []*Namespace `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListNamespacesResponse) Reset()         { *m = ListNamespacesResponse{} }
func (m *ListNamespacesResponse) String() string { return proto.CompactTextString(m) }
func (*ListNamespacesResponse) ProtoMessage()    {}
func (*ListNamespacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{0}
}

func (m *ListNamespacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNamespacesResponse.Unmarshal(m, b)
}
func (m *ListNamespacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNamespacesResponse.Marshal(b, m, deterministic)
}
func (m *ListNamespacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNamespacesResponse.Merge(m, src)
}
func (m *ListNamespacesResponse) XXX_Size() int {
	return xxx_messageInfo_ListNamespacesResponse.Size(m)
}
func (m *ListNamespacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNamespacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNamespacesResponse proto.InternalMessageInfo

func (m *ListNamespacesResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListNamespacesResponse) GetNamespaces() []*Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type Namespace struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{1}
}

func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ListNamespacesResponse)(nil), "api.ListNamespacesResponse")
	proto.RegisterType((*Namespace)(nil), "api.Namespace")
}

func init() { proto.RegisterFile("namespace.proto", fileDescriptor_ecb1e126f615f5dd) }

var fileDescriptor_ecb1e126f615f5dd = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4b, 0xcc, 0x4d,
	0x2d, 0x2e, 0x48, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8,
	0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb,
	0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x91, 0x92, 0x86, 0xca, 0x82, 0x79,
	0x49, 0xa5, 0x69, 0xfa, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x10, 0x49, 0xa5, 0x38, 0x2e, 0x31, 0x9f,
	0xcc, 0xe2, 0x12, 0x3f, 0x98, 0xb1, 0xc5, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42,
	0x22, 0x5c, 0xac, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10,
	0x8e, 0x90, 0x1e, 0x17, 0x17, 0xdc, 0x09, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x7c,
	0x7a, 0x89, 0x05, 0x99, 0x7a, 0x70, 0x23, 0x82, 0x90, 0x54, 0x28, 0xc9, 0x73, 0x71, 0xc2, 0x25,
	0x84, 0x84, 0xb8, 0x58, 0x40, 0x52, 0x60, 0x13, 0x39, 0x83, 0xc0, 0x6c, 0xa3, 0x6a, 0x2e, 0x01,
	0xb8, 0x82, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0xa1, 0x74, 0x2e, 0x3e, 0x54, 0x47, 0x09,
	0x89, 0xe9, 0x41, 0x3c, 0xa1, 0x07, 0xf3, 0x84, 0x9e, 0x2b, 0xc8, 0x13, 0x52, 0xd2, 0x60, 0xab,
	0xb1, 0xfb, 0x40, 0x49, 0xa1, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x52, 0x42, 0x12, 0xa0, 0x80, 0x29,
	0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x47, 0xb8, 0x2e, 0x89, 0x0d, 0x6c, 0x9c,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb5, 0xa6, 0xf4, 0xfe, 0x57, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NamespaceServiceClient is the client API for NamespaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NamespaceServiceClient interface {
	ListNamespaces(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
}

type namespaceServiceClient struct {
	cc *grpc.ClientConn
}

func NewNamespaceServiceClient(cc *grpc.ClientConn) NamespaceServiceClient {
	return &namespaceServiceClient{cc}
}

func (c *namespaceServiceClient) ListNamespaces(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, "/api.NamespaceService/ListNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespaceServiceServer is the server API for NamespaceService service.
type NamespaceServiceServer interface {
	ListNamespaces(context.Context, *empty.Empty) (*ListNamespacesResponse, error)
}

// UnimplementedNamespaceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNamespaceServiceServer struct {
}

func (*UnimplementedNamespaceServiceServer) ListNamespaces(ctx context.Context, req *empty.Empty) (*ListNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}

func RegisterNamespaceServiceServer(s *grpc.Server, srv NamespaceServiceServer) {
	s.RegisterService(&_NamespaceService_serviceDesc, srv)
}

func _NamespaceService_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NamespaceService/ListNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).ListNamespaces(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _NamespaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.NamespaceService",
	HandlerType: (*NamespaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNamespaces",
			Handler:    _NamespaceService_ListNamespaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namespace.proto",
}