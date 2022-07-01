// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitsong/fantoken/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// QueryFanTokenRequest is request type for the Query/FanToken RPC method
type QueryFanTokenRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryFanTokenRequest) Reset()         { *m = QueryFanTokenRequest{} }
func (m *QueryFanTokenRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFanTokenRequest) ProtoMessage()    {}
func (*QueryFanTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e961b51f12bacf3d, []int{0}
}
func (m *QueryFanTokenRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFanTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFanTokenRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFanTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFanTokenRequest.Merge(m, src)
}
func (m *QueryFanTokenRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFanTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFanTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFanTokenRequest proto.InternalMessageInfo

func (m *QueryFanTokenRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

// QueryFanTokenResponse is response type for the Query/FanToken RPC method
type QueryFanTokenResponse struct {
	Fantoken *FanToken `protobuf:"bytes,1,opt,name=fantoken,proto3" json:"fantoken,omitempty"`
}

func (m *QueryFanTokenResponse) Reset()         { *m = QueryFanTokenResponse{} }
func (m *QueryFanTokenResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFanTokenResponse) ProtoMessage()    {}
func (*QueryFanTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e961b51f12bacf3d, []int{1}
}
func (m *QueryFanTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFanTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFanTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFanTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFanTokenResponse.Merge(m, src)
}
func (m *QueryFanTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFanTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFanTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFanTokenResponse proto.InternalMessageInfo

func (m *QueryFanTokenResponse) GetFantoken() *FanToken {
	if m != nil {
		return m.Fantoken
	}
	return nil
}

// QueryFanTokensRequest is request type for the Query/FanTokens RPC method
type QueryFanTokensRequest struct {
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryFanTokensRequest) Reset()         { *m = QueryFanTokensRequest{} }
func (m *QueryFanTokensRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFanTokensRequest) ProtoMessage()    {}
func (*QueryFanTokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e961b51f12bacf3d, []int{2}
}
func (m *QueryFanTokensRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFanTokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFanTokensRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFanTokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFanTokensRequest.Merge(m, src)
}
func (m *QueryFanTokensRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFanTokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFanTokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFanTokensRequest proto.InternalMessageInfo

func (m *QueryFanTokensRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *QueryFanTokensRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryFanTokensResponse is response type for the Query/FanTokens RPC method
type QueryFanTokensResponse struct {
	Fantokens  []*FanToken         `protobuf:"bytes,1,rep,name=fantokens,proto3" json:"fantokens,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryFanTokensResponse) Reset()         { *m = QueryFanTokensResponse{} }
func (m *QueryFanTokensResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFanTokensResponse) ProtoMessage()    {}
func (*QueryFanTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e961b51f12bacf3d, []int{3}
}
func (m *QueryFanTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFanTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFanTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFanTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFanTokensResponse.Merge(m, src)
}
func (m *QueryFanTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFanTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFanTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFanTokensResponse proto.InternalMessageInfo

func (m *QueryFanTokensResponse) GetFantokens() []*FanToken {
	if m != nil {
		return m.Fantokens
	}
	return nil
}

func (m *QueryFanTokensResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryParametersRequest is request type for the Query/Parameters RPC method
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e961b51f12bacf3d, []int{4}
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

// QueryParametersResponse is response type for the Query/Parameters RPC method
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e961b51f12bacf3d, []int{5}
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

func init() {
	proto.RegisterType((*QueryFanTokenRequest)(nil), "bitsong.fantoken.v1beta1.QueryFanTokenRequest")
	proto.RegisterType((*QueryFanTokenResponse)(nil), "bitsong.fantoken.v1beta1.QueryFanTokenResponse")
	proto.RegisterType((*QueryFanTokensRequest)(nil), "bitsong.fantoken.v1beta1.QueryFanTokensRequest")
	proto.RegisterType((*QueryFanTokensResponse)(nil), "bitsong.fantoken.v1beta1.QueryFanTokensResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "bitsong.fantoken.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "bitsong.fantoken.v1beta1.QueryParamsResponse")
}

func init() {
	proto.RegisterFile("bitsong/fantoken/v1beta1/query.proto", fileDescriptor_e961b51f12bacf3d)
}

var fileDescriptor_e961b51f12bacf3d = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6b, 0x13, 0x41,
	0x18, 0xcd, 0xb4, 0x36, 0x34, 0xe3, 0x6d, 0x8c, 0x12, 0x42, 0x59, 0xc3, 0x6a, 0x4d, 0xd5, 0x76,
	0xc6, 0x56, 0xf0, 0x28, 0xd2, 0x43, 0xbd, 0xc6, 0xa0, 0x08, 0xde, 0x66, 0xe3, 0x64, 0x3b, 0xd8,
	0xcc, 0xb7, 0xcd, 0x4c, 0xc4, 0x20, 0xbd, 0xf8, 0x0f, 0x28, 0x78, 0xd4, 0x9b, 0xff, 0x4c, 0x8f,
	0x05, 0x41, 0x3c, 0x89, 0x24, 0xfe, 0x21, 0x92, 0xf9, 0xb1, 0x69, 0x42, 0x97, 0xec, 0x29, 0x93,
	0x99, 0xf7, 0xde, 0xf7, 0xbe, 0xef, 0xcd, 0x2c, 0xbe, 0x9b, 0x48, 0xa3, 0x41, 0xa5, 0xac, 0xcf,
	0x95, 0x81, 0x77, 0x42, 0xb1, 0xf7, 0xfb, 0x89, 0x30, 0x7c, 0x9f, 0x9d, 0x8e, 0xc4, 0x70, 0x4c,
	0xb3, 0x21, 0x18, 0x20, 0x0d, 0x8f, 0xa2, 0x01, 0x45, 0x3d, 0xaa, 0x19, 0xf5, 0x40, 0x0f, 0x40,
	0xb3, 0x84, 0x6b, 0x91, 0x53, 0x7b, 0x20, 0x95, 0x63, 0x36, 0x1f, 0x5c, 0x3e, 0xb7, 0x92, 0x39,
	0x2a, 0xe3, 0xa9, 0x54, 0xdc, 0x48, 0x08, 0xd8, 0x7a, 0x0a, 0x29, 0xd8, 0x25, 0x9b, 0xad, 0xfc,
	0xee, 0x56, 0x0a, 0x90, 0x9e, 0x08, 0xc6, 0x33, 0xc9, 0xb8, 0x52, 0x60, 0x2c, 0x45, 0xfb, 0xd3,
	0x76, 0xa1, 0xff, 0xdc, 0xaa, 0x03, 0x6e, 0x17, 0x02, 0x33, 0x3e, 0xe4, 0x03, 0xaf, 0x17, 0xef,
	0xe2, 0xfa, 0x8b, 0x99, 0xcb, 0x23, 0xae, 0x5e, 0xce, 0x50, 0x5d, 0x71, 0x3a, 0x12, 0xda, 0x90,
	0x3a, 0xde, 0x78, 0x2b, 0x14, 0x0c, 0x1a, 0xa8, 0x85, 0x76, 0x6a, 0x5d, 0xf7, 0x27, 0x7e, 0x8d,
	0x6f, 0x2e, 0xa1, 0x75, 0x06, 0x4a, 0x0b, 0xf2, 0x14, 0x6f, 0x86, 0x3a, 0x96, 0x71, 0xfd, 0x20,
	0xa6, 0x45, 0x33, 0xa4, 0x39, 0x3b, 0xe7, 0xc4, 0x67, 0x4b, 0xc2, 0x3a, 0xf8, 0xd8, 0xc2, 0x35,
	0x3e, 0x32, 0xc7, 0x30, 0x94, 0x66, 0xec, 0xbd, 0xcc, 0x37, 0xc8, 0x11, 0xc6, 0xf3, 0xa9, 0x36,
	0xd6, 0x6c, 0xe1, 0x7b, 0xd4, 0x45, 0x40, 0x67, 0x11, 0x50, 0x97, 0x6a, 0xa8, 0xdc, 0xe1, 0xa9,
	0xf0, 0xca, 0xdd, 0x4b, 0xcc, 0xf8, 0x07, 0xc2, 0xb7, 0x96, 0xeb, 0xfb, 0xce, 0x9e, 0xe1, 0x5a,
	0x70, 0xa9, 0x1b, 0xa8, 0xb5, 0x5e, 0xb2, 0xb5, 0x39, 0x89, 0x3c, 0xbf, 0xc2, 0x64, 0x7b, 0xa5,
	0x49, 0x57, 0x7e, 0xc1, 0x65, 0x1d, 0x13, 0x6b, 0xb2, 0x63, 0x03, 0xf4, 0x7d, 0xc4, 0xaf, 0xf0,
	0x8d, 0x85, 0xdd, 0x3c, 0x91, 0xaa, 0x0b, 0xda, 0xe7, 0xd1, 0x2a, 0x36, 0xed, 0x98, 0x87, 0xd7,
	0xce, 0xff, 0xdc, 0xae, 0x74, 0x3d, 0xeb, 0xe0, 0xd7, 0x3a, 0xde, 0xb0, 0xba, 0xe4, 0x3b, 0xc2,
	0x9b, 0xa1, 0x2f, 0x42, 0x8b, 0x65, 0xae, 0xba, 0x47, 0x4d, 0x56, 0x1a, 0xef, 0x7c, 0xc7, 0xec,
	0xd3, 0xcf, 0x7f, 0x5f, 0xd7, 0xee, 0x93, 0x36, 0x2b, 0xbc, 0xc0, 0xf6, 0x2e, 0xb2, 0x8f, 0xf6,
	0xe7, 0x8c, 0x7c, 0x43, 0xb8, 0x96, 0xc7, 0x46, 0xca, 0xd6, 0x0b, 0xe3, 0x6b, 0x3e, 0x2a, 0x4f,
	0xf0, 0x0e, 0x1f, 0x5a, 0x87, 0xdb, 0xe4, 0x0e, 0x5b, 0xf9, 0x16, 0x35, 0xf9, 0x8c, 0x70, 0xd5,
	0xcd, 0x97, 0xec, 0xae, 0xa8, 0xb4, 0x10, 0x6b, 0x73, 0xaf, 0x24, 0xda, 0x9b, 0xda, 0xb1, 0xa6,
	0x62, 0xd2, 0x62, 0x2b, 0xde, 0xfd, 0x61, 0xe7, 0x7c, 0x12, 0xa1, 0x8b, 0x49, 0x84, 0xfe, 0x4e,
	0x22, 0xf4, 0x65, 0x1a, 0x55, 0x2e, 0xa6, 0x51, 0xe5, 0xf7, 0x34, 0xaa, 0xbc, 0x79, 0x92, 0x4a,
	0x73, 0x3c, 0x4a, 0x68, 0x0f, 0x06, 0x41, 0x05, 0xfa, 0x7d, 0xd9, 0x93, 0xfc, 0x84, 0xa5, 0xb0,
	0x17, 0x84, 0x3f, 0xcc, 0xa5, 0xcd, 0x38, 0x13, 0x3a, 0xa9, 0xda, 0x4f, 0xc9, 0xe3, 0xff, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb5, 0x3f, 0x7e, 0x32, 0x5c, 0x05, 0x00, 0x00,
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
	// FanToken returns fantoken with fantoken name
	FanToken(ctx context.Context, in *QueryFanTokenRequest, opts ...grpc.CallOption) (*QueryFanTokenResponse, error)
	// FanTokens returns the fantoken list
	FanTokens(ctx context.Context, in *QueryFanTokensRequest, opts ...grpc.CallOption) (*QueryFanTokensResponse, error)
	// Params queries the fantoken parameters
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) FanToken(ctx context.Context, in *QueryFanTokenRequest, opts ...grpc.CallOption) (*QueryFanTokenResponse, error) {
	out := new(QueryFanTokenResponse)
	err := c.cc.Invoke(ctx, "/bitsong.fantoken.v1beta1.Query/FanToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FanTokens(ctx context.Context, in *QueryFanTokensRequest, opts ...grpc.CallOption) (*QueryFanTokensResponse, error) {
	out := new(QueryFanTokensResponse)
	err := c.cc.Invoke(ctx, "/bitsong.fantoken.v1beta1.Query/FanTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/bitsong.fantoken.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// FanToken returns fantoken with fantoken name
	FanToken(context.Context, *QueryFanTokenRequest) (*QueryFanTokenResponse, error)
	// FanTokens returns the fantoken list
	FanTokens(context.Context, *QueryFanTokensRequest) (*QueryFanTokensResponse, error)
	// Params queries the fantoken parameters
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) FanToken(ctx context.Context, req *QueryFanTokenRequest) (*QueryFanTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FanToken not implemented")
}
func (*UnimplementedQueryServer) FanTokens(ctx context.Context, req *QueryFanTokensRequest) (*QueryFanTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FanTokens not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_FanToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFanTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FanToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitsong.fantoken.v1beta1.Query/FanToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FanToken(ctx, req.(*QueryFanTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FanTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFanTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FanTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitsong.fantoken.v1beta1.Query/FanTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FanTokens(ctx, req.(*QueryFanTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/bitsong.fantoken.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bitsong.fantoken.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FanToken",
			Handler:    _Query_FanToken_Handler,
		},
		{
			MethodName: "FanTokens",
			Handler:    _Query_FanTokens_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitsong/fantoken/v1beta1/query.proto",
}

func (m *QueryFanTokenRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFanTokenRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFanTokenRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFanTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFanTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFanTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fantoken != nil {
		{
			size, err := m.Fantoken.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFanTokensRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFanTokensRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFanTokensRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFanTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFanTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFanTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Fantokens) > 0 {
		for iNdEx := len(m.Fantokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fantokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryFanTokenRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFanTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Fantoken != nil {
		l = m.Fantoken.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFanTokensRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFanTokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Fantokens) > 0 {
		for _, e := range m.Fantokens {
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

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryFanTokenRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFanTokenRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFanTokenRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
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
func (m *QueryFanTokenResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFanTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFanTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fantoken", wireType)
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
			if m.Fantoken == nil {
				m.Fantoken = &FanToken{}
			}
			if err := m.Fantoken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryFanTokensRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFanTokensRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFanTokensRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
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
			m.Authority = string(dAtA[iNdEx:postIndex])
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
func (m *QueryFanTokensResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFanTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFanTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fantokens", wireType)
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
			m.Fantokens = append(m.Fantokens, &FanToken{})
			if err := m.Fantokens[len(m.Fantokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
