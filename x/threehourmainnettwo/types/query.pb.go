// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: threehourmainnettwo/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("threehourmainnettwo/query.proto", fileDescriptor_7bfb4e0530bb5582) }

var fileDescriptor_7bfb4e0530bb5582 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcf, 0x31, 0x4e, 0xc6, 0x30,
	0x0c, 0x05, 0xe0, 0x76, 0x00, 0xa4, 0x8e, 0x8c, 0x15, 0x0a, 0x3b, 0x43, 0xac, 0xc2, 0x09, 0xe0,
	0x06, 0xac, 0x30, 0x39, 0xad, 0x49, 0x23, 0xd1, 0x38, 0x24, 0x2e, 0xd0, 0x5b, 0x70, 0x2c, 0xc6,
	0x8e, 0x8c, 0xa8, 0xbd, 0xc8, 0xaf, 0xbf, 0x59, 0xb3, 0x3e, 0x7d, 0xf2, 0xf3, 0x6b, 0x6e, 0x65,
	0x8c, 0x44, 0x23, 0xcf, 0x71, 0x42, 0xe7, 0x3d, 0x89, 0x7c, 0x31, 0x7c, 0xcc, 0x14, 0x17, 0x1d,
	0x22, 0x0b, 0x5f, 0xeb, 0x37, 0x1c, 0x06, 0x14, 0x5d, 0x70, 0xa5, 0xac, 0xbd, 0xb1, 0xcc, 0xf6,
	0x9d, 0x00, 0x83, 0x03, 0xf4, 0x9e, 0x05, 0xc5, 0xb1, 0x4f, 0xf9, 0x5a, 0x7b, 0xd7, 0x73, 0x9a,
	0x38, 0x81, 0xc1, 0x44, 0xb9, 0x06, 0x3e, 0x3b, 0x43, 0x82, 0x1d, 0x04, 0xb4, 0xce, 0x1f, 0x38,
	0xdb, 0xfb, 0xab, 0xe6, 0xe2, 0xf9, 0x2c, 0x9e, 0x5e, 0x7f, 0x37, 0x55, 0xaf, 0x9b, 0xaa, 0xff,
	0x37, 0x55, 0xff, 0xec, 0xaa, 0x5a, 0x77, 0x55, 0xfd, 0xed, 0xaa, 0x7a, 0x79, 0xb4, 0x4e, 0xc6,
	0xd9, 0xe8, 0x9e, 0x27, 0xc8, 0x7f, 0x42, 0x69, 0xcf, 0x77, 0x31, 0x95, 0x25, 0x50, 0x32, 0x97,
	0x47, 0xd9, 0xc3, 0x29, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x7f, 0x3b, 0x0a, 0x09, 0x01, 0x00, 0x00,
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
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "faddat.threehourmainnettwo.threehourmainnettwo.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "threehourmainnettwo/query.proto",
}