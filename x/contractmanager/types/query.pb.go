// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/contractmanager/query.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9524a427f219917, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9524a427f219917, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryFailuresRequest is request type for the Query/Failures RPC method.
type QueryFailuresRequest struct {
	// address of the contract which Sudo call failed.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// ID of the failure for the given contract.
	FailureId  uint64             `protobuf:"varint,2,opt,name=failure_id,json=failureId,proto3" json:"failure_id,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryFailuresRequest) Reset()         { *m = QueryFailuresRequest{} }
func (m *QueryFailuresRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFailuresRequest) ProtoMessage()    {}
func (*QueryFailuresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9524a427f219917, []int{2}
}
func (m *QueryFailuresRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFailuresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFailuresRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFailuresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFailuresRequest.Merge(m, src)
}
func (m *QueryFailuresRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFailuresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFailuresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFailuresRequest proto.InternalMessageInfo

func (m *QueryFailuresRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *QueryFailuresRequest) GetFailureId() uint64 {
	if m != nil {
		return m.FailureId
	}
	return 0
}

func (m *QueryFailuresRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryFailuresResponse is response type for the Query/Failures RPC method.
type QueryFailuresResponse struct {
	Failures   []Failure           `protobuf:"bytes,1,rep,name=failures,proto3" json:"failures"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryFailuresResponse) Reset()         { *m = QueryFailuresResponse{} }
func (m *QueryFailuresResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFailuresResponse) ProtoMessage()    {}
func (*QueryFailuresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9524a427f219917, []int{3}
}
func (m *QueryFailuresResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFailuresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFailuresResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFailuresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFailuresResponse.Merge(m, src)
}
func (m *QueryFailuresResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFailuresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFailuresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFailuresResponse proto.InternalMessageInfo

func (m *QueryFailuresResponse) GetFailures() []Failure {
	if m != nil {
		return m.Failures
	}
	return nil
}

func (m *QueryFailuresResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "neutron.contractmanager.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "neutron.contractmanager.QueryParamsResponse")
	proto.RegisterType((*QueryFailuresRequest)(nil), "neutron.contractmanager.QueryFailuresRequest")
	proto.RegisterType((*QueryFailuresResponse)(nil), "neutron.contractmanager.QueryFailuresResponse")
}

func init() {
	proto.RegisterFile("neutron/contractmanager/query.proto", fileDescriptor_f9524a427f219917)
}

var fileDescriptor_f9524a427f219917 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xc1, 0x6a, 0x13, 0x41,
	0x18, 0xc7, 0x33, 0x69, 0x1b, 0xdb, 0x29, 0x28, 0x8c, 0x11, 0x43, 0xd0, 0x4d, 0xba, 0x55, 0x1b,
	0xad, 0x99, 0xa1, 0x29, 0x88, 0x08, 0x82, 0xe6, 0x50, 0xf1, 0x16, 0x17, 0x4f, 0x5e, 0x64, 0x92,
	0x8c, 0xe3, 0x42, 0x33, 0xb3, 0x9d, 0x99, 0x15, 0x4b, 0xe9, 0xc5, 0x9b, 0xe0, 0x41, 0x50, 0xf0,
	0x05, 0xf4, 0x01, 0x7c, 0x8b, 0x1e, 0x0b, 0x5e, 0x3c, 0x89, 0x24, 0x3e, 0x88, 0x64, 0x66, 0xd2,
	0x36, 0x29, 0xdb, 0xc4, 0x83, 0xbd, 0xed, 0xce, 0xfe, 0xbf, 0xef, 0xff, 0x9b, 0xff, 0xf7, 0xb1,
	0x70, 0x55, 0xb0, 0xd4, 0x28, 0x29, 0x48, 0x47, 0x0a, 0xa3, 0x68, 0xc7, 0xf4, 0xa8, 0xa0, 0x9c,
	0x29, 0xb2, 0x93, 0x32, 0xb5, 0x8b, 0x13, 0x25, 0x8d, 0x44, 0x57, 0xbd, 0x08, 0x4f, 0x88, 0xca,
	0x45, 0x2e, 0xb9, 0xb4, 0x1a, 0x32, 0x7c, 0x72, 0xf2, 0xf2, 0x35, 0x2e, 0x25, 0xdf, 0x66, 0x84,
	0x26, 0x31, 0xa1, 0x42, 0x48, 0x43, 0x4d, 0x2c, 0x85, 0xf6, 0x5f, 0xef, 0x74, 0xa4, 0xee, 0x49,
	0x4d, 0xda, 0x54, 0x33, 0xe7, 0x42, 0xde, 0x6c, 0xb4, 0x99, 0xa1, 0x1b, 0x24, 0xa1, 0x3c, 0x16,
	0x56, 0xec, 0xb5, 0x37, 0xb2, 0xe8, 0x12, 0xaa, 0x68, 0x6f, 0xd4, 0xf1, 0x66, 0x96, 0xea, 0x15,
	0x8d, 0xb7, 0x53, 0xc5, 0x9c, 0x2c, 0x2c, 0x42, 0xf4, 0x6c, 0x68, 0xd7, 0xb2, 0xb5, 0x11, 0xdb,
	0x49, 0x99, 0x36, 0xe1, 0x73, 0x78, 0x79, 0xec, 0x54, 0x27, 0x52, 0x68, 0x86, 0x1e, 0xc2, 0x82,
	0xf3, 0x28, 0x81, 0x2a, 0xa8, 0x2d, 0x37, 0x2a, 0x38, 0x23, 0x03, 0xec, 0x0a, 0x9b, 0xf3, 0x07,
	0xbf, 0x2a, 0xb9, 0xc8, 0x17, 0x85, 0x5f, 0x00, 0x2c, 0xda, 0xb6, 0x5b, 0x0e, 0x61, 0x64, 0x87,
	0x4a, 0xf0, 0x02, 0xed, 0x76, 0x15, 0xd3, 0xae, 0xf1, 0x52, 0x34, 0x7a, 0x45, 0xd7, 0x21, 0xf4,
	0xbc, 0x2f, 0xe3, 0x6e, 0x29, 0x5f, 0x05, 0xb5, 0xf9, 0x68, 0xc9, 0x9f, 0x3c, 0xed, 0xa2, 0x2d,
	0x08, 0x8f, 0xe3, 0x29, 0xcd, 0x59, 0xa8, 0x5b, 0xd8, 0x65, 0x89, 0x87, 0x59, 0x62, 0x37, 0x31,
	0x9f, 0x25, 0x6e, 0x51, 0xce, 0xbc, 0x69, 0x74, 0xa2, 0x32, 0xfc, 0x0a, 0xe0, 0x95, 0x09, 0x32,
	0x7f, 0xe5, 0x26, 0x5c, 0xf4, 0x76, 0x43, 0xb6, 0xb9, 0xda, 0x72, 0xa3, 0x9a, 0x79, 0x69, 0x5f,
	0xec, 0x6f, 0x7d, 0x54, 0x87, 0x9e, 0x8c, 0x51, 0xe6, 0x2d, 0xe5, 0xda, 0x54, 0x4a, 0x07, 0x70,
	0x12, 0xb3, 0xf1, 0x7e, 0x01, 0x2e, 0x58, 0x4c, 0xf4, 0x01, 0xc0, 0x82, 0xcb, 0x18, 0xad, 0x67,
	0xf2, 0x9c, 0x1e, 0x6c, 0xf9, 0xee, 0x6c, 0x62, 0xe7, 0x1d, 0xae, 0xbd, 0xfb, 0xf1, 0xe7, 0x53,
	0x7e, 0x05, 0x55, 0xc8, 0xd9, 0x2b, 0x87, 0xbe, 0x03, 0x78, 0xf1, 0xb1, 0x1b, 0x99, 0x0f, 0x01,
	0xd5, 0xcf, 0x76, 0x9a, 0x58, 0x81, 0x32, 0x9e, 0x55, 0xee, 0xd1, 0x1e, 0x59, 0xb4, 0x07, 0xe8,
	0x3e, 0x99, 0xb2, 0xe7, 0x9a, 0xec, 0xf9, 0x65, 0xda, 0x27, 0x7b, 0xc7, 0xbb, 0xb4, 0x8f, 0xbe,
	0x01, 0x78, 0x69, 0x9c, 0x59, 0xff, 0x6f, 0xe8, 0x4d, 0x0b, 0x5d, 0x47, 0xeb, 0xff, 0x00, 0x8d,
	0x3e, 0x03, 0xb8, 0x78, 0x5e, 0x80, 0xb7, 0x2d, 0xe0, 0x2a, 0x5a, 0x99, 0x0a, 0xd8, 0x6c, 0x1d,
	0xf4, 0x03, 0x70, 0xd8, 0x0f, 0xc0, 0xef, 0x7e, 0x00, 0x3e, 0x0e, 0x82, 0xdc, 0xe1, 0x20, 0xc8,
	0xfd, 0x1c, 0x04, 0xb9, 0x17, 0xf7, 0x78, 0x6c, 0x5e, 0xa7, 0x6d, 0xdc, 0x91, 0xbd, 0x51, 0x9b,
	0xba, 0x54, 0xfc, 0xa8, 0xe5, 0xdb, 0x53, 0x4d, 0xcd, 0x6e, 0xc2, 0x74, 0xbb, 0x60, 0xff, 0x48,
	0x9b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xb5, 0xbb, 0x34, 0x7e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a Failure by contract address and failure ID.
	AddressFailure(ctx context.Context, in *QueryFailuresRequest, opts ...grpc.CallOption) (*QueryFailuresResponse, error)
	// Queries Failures by contract address.
	AddressFailures(ctx context.Context, in *QueryFailuresRequest, opts ...grpc.CallOption) (*QueryFailuresResponse, error)
	// Queries a list of Failures occurred on the network.
	Failures(ctx context.Context, in *QueryFailuresRequest, opts ...grpc.CallOption) (*QueryFailuresResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/neutron.contractmanager.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AddressFailure(ctx context.Context, in *QueryFailuresRequest, opts ...grpc.CallOption) (*QueryFailuresResponse, error) {
	out := new(QueryFailuresResponse)
	err := c.cc.Invoke(ctx, "/neutron.contractmanager.Query/AddressFailure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AddressFailures(ctx context.Context, in *QueryFailuresRequest, opts ...grpc.CallOption) (*QueryFailuresResponse, error) {
	out := new(QueryFailuresResponse)
	err := c.cc.Invoke(ctx, "/neutron.contractmanager.Query/AddressFailures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Failures(ctx context.Context, in *QueryFailuresRequest, opts ...grpc.CallOption) (*QueryFailuresResponse, error) {
	out := new(QueryFailuresResponse)
	err := c.cc.Invoke(ctx, "/neutron.contractmanager.Query/Failures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a Failure by contract address and failure ID.
	AddressFailure(context.Context, *QueryFailuresRequest) (*QueryFailuresResponse, error)
	// Queries Failures by contract address.
	AddressFailures(context.Context, *QueryFailuresRequest) (*QueryFailuresResponse, error)
	// Queries a list of Failures occurred on the network.
	Failures(context.Context, *QueryFailuresRequest) (*QueryFailuresResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) AddressFailure(ctx context.Context, req *QueryFailuresRequest) (*QueryFailuresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressFailure not implemented")
}
func (*UnimplementedQueryServer) AddressFailures(ctx context.Context, req *QueryFailuresRequest) (*QueryFailuresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressFailures not implemented")
}
func (*UnimplementedQueryServer) Failures(ctx context.Context, req *QueryFailuresRequest) (*QueryFailuresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Failures not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.contractmanager.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AddressFailure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFailuresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AddressFailure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.contractmanager.Query/AddressFailure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AddressFailure(ctx, req.(*QueryFailuresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AddressFailures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFailuresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AddressFailures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.contractmanager.Query/AddressFailures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AddressFailures(ctx, req.(*QueryFailuresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Failures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFailuresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Failures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.contractmanager.Query/Failures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Failures(ctx, req.(*QueryFailuresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "neutron.contractmanager.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "AddressFailure",
			Handler:    _Query_AddressFailure_Handler,
		},
		{
			MethodName: "AddressFailures",
			Handler:    _Query_AddressFailures_Handler,
		},
		{
			MethodName: "Failures",
			Handler:    _Query_Failures_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "neutron/contractmanager/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryFailuresRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFailuresRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFailuresRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.FailureId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.FailureId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFailuresResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFailuresResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFailuresResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Failures) > 0 {
		for iNdEx := len(m.Failures) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Failures[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryFailuresRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.FailureId != 0 {
		n += 1 + sovQuery(uint64(m.FailureId))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFailuresResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Failures) > 0 {
		for _, e := range m.Failures {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryFailuresRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryFailuresRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFailuresRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailureId", wireType)
			}
			m.FailureId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FailureId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryFailuresResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryFailuresResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFailuresResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Failures = append(m.Failures, Failure{})
			if err := m.Failures[len(m.Failures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
