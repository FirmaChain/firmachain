// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: firmachain/token/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type QueryGetTokenDataRequest struct {
	TokenID string `protobuf:"bytes,1,opt,name=tokenID,proto3" json:"tokenID,omitempty"`
}

func (m *QueryGetTokenDataRequest) Reset()         { *m = QueryGetTokenDataRequest{} }
func (m *QueryGetTokenDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetTokenDataRequest) ProtoMessage()    {}
func (*QueryGetTokenDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4f2a16e589eeae, []int{0}
}
func (m *QueryGetTokenDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTokenDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTokenDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTokenDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTokenDataRequest.Merge(m, src)
}
func (m *QueryGetTokenDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTokenDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTokenDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTokenDataRequest proto.InternalMessageInfo

func (m *QueryGetTokenDataRequest) GetTokenID() string {
	if m != nil {
		return m.TokenID
	}
	return ""
}

type QueryGetTokenDataResponse struct {
	TokenData TokenData `protobuf:"bytes,1,opt,name=tokenData,proto3" json:"tokenData"`
}

func (m *QueryGetTokenDataResponse) Reset()         { *m = QueryGetTokenDataResponse{} }
func (m *QueryGetTokenDataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetTokenDataResponse) ProtoMessage()    {}
func (*QueryGetTokenDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4f2a16e589eeae, []int{1}
}
func (m *QueryGetTokenDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTokenDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTokenDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTokenDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTokenDataResponse.Merge(m, src)
}
func (m *QueryGetTokenDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTokenDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTokenDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTokenDataResponse proto.InternalMessageInfo

func (m *QueryGetTokenDataResponse) GetTokenData() TokenData {
	if m != nil {
		return m.TokenData
	}
	return TokenData{}
}

type QueryAllTokenDataRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTokenDataRequest) Reset()         { *m = QueryAllTokenDataRequest{} }
func (m *QueryAllTokenDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllTokenDataRequest) ProtoMessage()    {}
func (*QueryAllTokenDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4f2a16e589eeae, []int{2}
}
func (m *QueryAllTokenDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTokenDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTokenDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTokenDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTokenDataRequest.Merge(m, src)
}
func (m *QueryAllTokenDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTokenDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTokenDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTokenDataRequest proto.InternalMessageInfo

func (m *QueryAllTokenDataRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllTokenDataResponse struct {
	TokenData  []TokenData         `protobuf:"bytes,1,rep,name=tokenData,proto3" json:"tokenData"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTokenDataResponse) Reset()         { *m = QueryAllTokenDataResponse{} }
func (m *QueryAllTokenDataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllTokenDataResponse) ProtoMessage()    {}
func (*QueryAllTokenDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4f2a16e589eeae, []int{3}
}
func (m *QueryAllTokenDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTokenDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTokenDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTokenDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTokenDataResponse.Merge(m, src)
}
func (m *QueryAllTokenDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTokenDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTokenDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTokenDataResponse proto.InternalMessageInfo

func (m *QueryAllTokenDataResponse) GetTokenData() []TokenData {
	if m != nil {
		return m.TokenData
	}
	return nil
}

func (m *QueryAllTokenDataResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryGetTokenListRequest struct {
	OwnerAddress string `protobuf:"bytes,1,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
}

func (m *QueryGetTokenListRequest) Reset()         { *m = QueryGetTokenListRequest{} }
func (m *QueryGetTokenListRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetTokenListRequest) ProtoMessage()    {}
func (*QueryGetTokenListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4f2a16e589eeae, []int{4}
}
func (m *QueryGetTokenListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTokenListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTokenListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTokenListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTokenListRequest.Merge(m, src)
}
func (m *QueryGetTokenListRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTokenListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTokenListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTokenListRequest proto.InternalMessageInfo

func (m *QueryGetTokenListRequest) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

type QueryGetTokenListResponse struct {
	TokenID []string `protobuf:"bytes,1,rep,name=tokenID,proto3" json:"tokenID,omitempty"`
}

func (m *QueryGetTokenListResponse) Reset()         { *m = QueryGetTokenListResponse{} }
func (m *QueryGetTokenListResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetTokenListResponse) ProtoMessage()    {}
func (*QueryGetTokenListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4f2a16e589eeae, []int{5}
}
func (m *QueryGetTokenListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTokenListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTokenListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTokenListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTokenListResponse.Merge(m, src)
}
func (m *QueryGetTokenListResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTokenListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTokenListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTokenListResponse proto.InternalMessageInfo

func (m *QueryGetTokenListResponse) GetTokenID() []string {
	if m != nil {
		return m.TokenID
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetTokenDataRequest)(nil), "firmachain.token.QueryGetTokenDataRequest")
	proto.RegisterType((*QueryGetTokenDataResponse)(nil), "firmachain.token.QueryGetTokenDataResponse")
	proto.RegisterType((*QueryAllTokenDataRequest)(nil), "firmachain.token.QueryAllTokenDataRequest")
	proto.RegisterType((*QueryAllTokenDataResponse)(nil), "firmachain.token.QueryAllTokenDataResponse")
	proto.RegisterType((*QueryGetTokenListRequest)(nil), "firmachain.token.QueryGetTokenListRequest")
	proto.RegisterType((*QueryGetTokenListResponse)(nil), "firmachain.token.QueryGetTokenListResponse")
}

func init() { proto.RegisterFile("firmachain/token/query.proto", fileDescriptor_cb4f2a16e589eeae) }

var fileDescriptor_cb4f2a16e589eeae = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x8b, 0x13, 0x31,
	0x18, 0xc6, 0x9b, 0xad, 0x7f, 0x68, 0xec, 0x41, 0x82, 0x87, 0x5a, 0x97, 0x51, 0x73, 0x58, 0xd7,
	0x5d, 0x48, 0x6c, 0x55, 0xf0, 0xa4, 0x74, 0x59, 0x5c, 0x05, 0x0f, 0x5a, 0x3c, 0x89, 0x20, 0x99,
	0x36, 0x66, 0x07, 0xa7, 0xc9, 0xec, 0x24, 0x55, 0x17, 0xf1, 0xe2, 0x27, 0x10, 0x3c, 0x08, 0x1e,
	0xbc, 0xf9, 0x5d, 0xf6, 0xb8, 0xe0, 0xc5, 0x93, 0x4a, 0xeb, 0x07, 0x91, 0x49, 0xd2, 0x76, 0xa6,
	0x9d, 0xda, 0xf5, 0x52, 0xa6, 0x99, 0xf7, 0x79, 0xf3, 0x7b, 0x9e, 0xf7, 0x65, 0xe0, 0xfa, 0xcb,
	0x28, 0x1d, 0xb0, 0xde, 0x3e, 0x8b, 0x24, 0x35, 0xea, 0x15, 0x97, 0xf4, 0x60, 0xc8, 0xd3, 0x43,
	0x92, 0xa4, 0xca, 0x28, 0x74, 0x7e, 0xf6, 0x96, 0xd8, 0xb7, 0xcd, 0x75, 0xa1, 0x94, 0x88, 0x39,
	0x65, 0x49, 0x44, 0x99, 0x94, 0xca, 0x30, 0x13, 0x29, 0xa9, 0x5d, 0x7d, 0x73, 0xab, 0xa7, 0xf4,
	0x40, 0x69, 0x1a, 0x32, 0xcd, 0x5d, 0x23, 0xfa, 0xba, 0x15, 0x72, 0xc3, 0x5a, 0x34, 0x61, 0x22,
	0x92, 0xb6, 0xd8, 0xd7, 0x5e, 0x5d, 0xb8, 0xd9, 0xfe, 0xbe, 0xe8, 0x33, 0xc3, 0x7c, 0xc9, 0x05,
	0xa1, 0x84, 0xb2, 0x8f, 0x34, 0x7b, 0x72, 0xa7, 0xf8, 0x16, 0x6c, 0x3c, 0xc9, 0x5a, 0xef, 0x71,
	0xf3, 0x34, 0x53, 0xec, 0x32, 0xc3, 0xba, 0xfc, 0x60, 0xc8, 0xb5, 0x41, 0x0d, 0x78, 0xd6, 0x76,
	0x79, 0xb8, 0xdb, 0x00, 0x57, 0xc0, 0x66, 0xad, 0x3b, 0xf9, 0x8b, 0x9f, 0xc3, 0x8b, 0x25, 0x2a,
	0x9d, 0x28, 0xa9, 0x39, 0xba, 0x07, 0x6b, 0x66, 0x72, 0x68, 0x85, 0xe7, 0xda, 0x97, 0xc8, 0xbc,
	0x77, 0x32, 0xd5, 0xed, 0x9c, 0x3a, 0xfa, 0x79, 0xb9, 0xd2, 0x9d, 0x69, 0x70, 0xe8, 0x99, 0x3a,
	0x71, 0xbc, 0xc0, 0x74, 0x1f, 0xc2, 0x99, 0x79, 0xdf, 0x7d, 0x83, 0xb8, 0xa4, 0x48, 0x96, 0x14,
	0x71, 0x91, 0xfb, 0xa4, 0xc8, 0x63, 0x26, 0xb8, 0xd7, 0x76, 0x73, 0x4a, 0xfc, 0x0d, 0x78, 0x0b,
	0xc5, 0x4b, 0xca, 0x2d, 0x54, 0xff, 0xd7, 0x02, 0xda, 0x2b, 0x60, 0xae, 0x59, 0xcc, 0x6b, 0x2b,
	0x31, 0xdd, 0xed, 0x05, 0xce, 0xbb, 0x73, 0xf3, 0x79, 0x14, 0x69, 0x33, 0xc9, 0x02, 0xc3, 0xba,
	0x7a, 0x23, 0x79, 0xda, 0xe9, 0xf7, 0x53, 0xae, 0xb5, 0x1f, 0x52, 0xe1, 0x0c, 0xdf, 0x9e, 0x9b,
	0x94, 0xd3, 0x7b, 0x9b, 0x85, 0x01, 0x57, 0x73, 0x03, 0x6e, 0xff, 0xaa, 0xc2, 0xd3, 0x56, 0x87,
	0xbe, 0x02, 0x58, 0x9b, 0x1a, 0x45, 0x5b, 0x8b, 0x29, 0x2c, 0x5b, 0x9f, 0xe6, 0xf6, 0x89, 0x6a,
	0x1d, 0x0a, 0xbe, 0xf3, 0xe1, 0xfb, 0x9f, 0x4f, 0x6b, 0x6d, 0x74, 0x83, 0xe6, 0x36, 0xb9, 0x7c,
	0xa9, 0x33, 0x1d, 0x7d, 0xe7, 0x49, 0xdf, 0xa3, 0xcf, 0x00, 0xd6, 0xa7, 0xfd, 0x3a, 0x71, 0xbc,
	0x94, 0xb1, 0x64, 0x9d, 0x96, 0x32, 0x96, 0x6d, 0x05, 0x26, 0x96, 0x71, 0x13, 0x6d, 0x9c, 0x8c,
	0x11, 0x7d, 0x01, 0xb0, 0x9e, 0xcf, 0x7d, 0x65, 0x7a, 0xb9, 0xe1, 0xae, 0x4c, 0x2f, 0x3f, 0x48,
	0xdc, 0xb2, 0x64, 0xdb, 0xe8, 0xfa, 0x3f, 0xc9, 0x44, 0x4e, 0xba, 0xf3, 0xe0, 0x68, 0x14, 0x80,
	0xe3, 0x51, 0x00, 0x7e, 0x8f, 0x02, 0xf0, 0x71, 0x1c, 0x54, 0x8e, 0xc7, 0x41, 0xe5, 0xc7, 0x38,
	0xa8, 0x3c, 0x23, 0x22, 0x32, 0xfb, 0xc3, 0x90, 0xf4, 0xd4, 0x60, 0x49, 0xbb, 0xb7, 0x13, 0xab,
	0x87, 0x09, 0xd7, 0xe1, 0x19, 0xfb, 0x25, 0xb9, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xe0,
	0x88, 0x27, 0xfe, 0x04, 0x00, 0x00,
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
	// Queries a tokenData by index.
	TokenData(ctx context.Context, in *QueryGetTokenDataRequest, opts ...grpc.CallOption) (*QueryGetTokenDataResponse, error)
	// Queries a list of tokenData items.
	TokenDataAll(ctx context.Context, in *QueryAllTokenDataRequest, opts ...grpc.CallOption) (*QueryAllTokenDataResponse, error)
	// Queries a list of getTokenList items.
	GetTokenList(ctx context.Context, in *QueryGetTokenListRequest, opts ...grpc.CallOption) (*QueryGetTokenListResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) TokenData(ctx context.Context, in *QueryGetTokenDataRequest, opts ...grpc.CallOption) (*QueryGetTokenDataResponse, error) {
	out := new(QueryGetTokenDataResponse)
	err := c.cc.Invoke(ctx, "/firmachain.token.Query/TokenData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TokenDataAll(ctx context.Context, in *QueryAllTokenDataRequest, opts ...grpc.CallOption) (*QueryAllTokenDataResponse, error) {
	out := new(QueryAllTokenDataResponse)
	err := c.cc.Invoke(ctx, "/firmachain.token.Query/TokenDataAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetTokenList(ctx context.Context, in *QueryGetTokenListRequest, opts ...grpc.CallOption) (*QueryGetTokenListResponse, error) {
	out := new(QueryGetTokenListResponse)
	err := c.cc.Invoke(ctx, "/firmachain.token.Query/GetTokenList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a tokenData by index.
	TokenData(context.Context, *QueryGetTokenDataRequest) (*QueryGetTokenDataResponse, error)
	// Queries a list of tokenData items.
	TokenDataAll(context.Context, *QueryAllTokenDataRequest) (*QueryAllTokenDataResponse, error)
	// Queries a list of getTokenList items.
	GetTokenList(context.Context, *QueryGetTokenListRequest) (*QueryGetTokenListResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) TokenData(ctx context.Context, req *QueryGetTokenDataRequest) (*QueryGetTokenDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenData not implemented")
}
func (*UnimplementedQueryServer) TokenDataAll(ctx context.Context, req *QueryAllTokenDataRequest) (*QueryAllTokenDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenDataAll not implemented")
}
func (*UnimplementedQueryServer) GetTokenList(ctx context.Context, req *QueryGetTokenListRequest) (*QueryGetTokenListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenList not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_TokenData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTokenDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmachain.token.Query/TokenData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenData(ctx, req.(*QueryGetTokenDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TokenDataAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTokenDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TokenDataAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmachain.token.Query/TokenDataAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TokenDataAll(ctx, req.(*QueryAllTokenDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetTokenList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTokenListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTokenList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/firmachain.token.Query/GetTokenList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTokenList(ctx, req.(*QueryGetTokenListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "firmachain.token.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TokenData",
			Handler:    _Query_TokenData_Handler,
		},
		{
			MethodName: "TokenDataAll",
			Handler:    _Query_TokenDataAll_Handler,
		},
		{
			MethodName: "GetTokenList",
			Handler:    _Query_GetTokenList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "firmachain/token/query.proto",
}

func (m *QueryGetTokenDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTokenDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTokenDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenID) > 0 {
		i -= len(m.TokenID)
		copy(dAtA[i:], m.TokenID)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.TokenID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetTokenDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTokenDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTokenDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TokenData.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllTokenDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTokenDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTokenDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllTokenDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTokenDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTokenDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.TokenData) > 0 {
		for iNdEx := len(m.TokenData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryGetTokenListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTokenListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTokenListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetTokenListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTokenListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTokenListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenID) > 0 {
		for iNdEx := len(m.TokenID) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TokenID[iNdEx])
			copy(dAtA[i:], m.TokenID[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.TokenID[iNdEx])))
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
func (m *QueryGetTokenDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenID)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetTokenDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TokenData.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllTokenDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllTokenDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TokenData) > 0 {
		for _, e := range m.TokenData {
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

func (m *QueryGetTokenListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetTokenListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TokenID) > 0 {
		for _, s := range m.TokenID {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetTokenDataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetTokenDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTokenDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenID", wireType)
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
			m.TokenID = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetTokenDataResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetTokenDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTokenDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenData", wireType)
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
			if err := m.TokenData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllTokenDataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllTokenDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTokenDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
func (m *QueryAllTokenDataResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllTokenDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTokenDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenData", wireType)
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
			m.TokenData = append(m.TokenData, TokenData{})
			if err := m.TokenData[len(m.TokenData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGetTokenListRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetTokenListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTokenListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
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
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetTokenListResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetTokenListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTokenListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenID", wireType)
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
			m.TokenID = append(m.TokenID, string(dAtA[iNdEx:postIndex]))
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
