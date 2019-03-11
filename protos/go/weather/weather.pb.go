// Code generated by protoc-gen-go. DO NOT EDIT.
// source: weather.proto

/*
Package weather is a generated protocol buffer package.

It is generated from these files:
	weather.proto

It has these top-level messages:
	WeatherRequest
	WeatherResponse
*/
package weather

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

type WeatherRequest struct {
	City    string `protobuf:"bytes,1,opt,name=city" json:"city,omitempty"`
	Country string `protobuf:"bytes,2,opt,name=country" json:"country,omitempty"`
}

func (m *WeatherRequest) Reset()                    { *m = WeatherRequest{} }
func (m *WeatherRequest) String() string            { return proto.CompactTextString(m) }
func (*WeatherRequest) ProtoMessage()               {}
func (*WeatherRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WeatherRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *WeatherRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type WeatherResponse struct {
	Temperature float32 `protobuf:"fixed32,1,opt,name=temperature" json:"temperature,omitempty"`
}

func (m *WeatherResponse) Reset()                    { *m = WeatherResponse{} }
func (m *WeatherResponse) String() string            { return proto.CompactTextString(m) }
func (*WeatherResponse) ProtoMessage()               {}
func (*WeatherResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WeatherResponse) GetTemperature() float32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

func init() {
	proto.RegisterType((*WeatherRequest)(nil), "weather.WeatherRequest")
	proto.RegisterType((*WeatherResponse)(nil), "weather.WeatherResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WeatherService service

type WeatherServiceClient interface {
	Get(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherResponse, error)
}

type weatherServiceClient struct {
	cc *grpc.ClientConn
}

func NewWeatherServiceClient(cc *grpc.ClientConn) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) Get(ctx context.Context, in *WeatherRequest, opts ...grpc.CallOption) (*WeatherResponse, error) {
	out := new(WeatherResponse)
	err := grpc.Invoke(ctx, "/weather.WeatherService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WeatherService service

type WeatherServiceServer interface {
	Get(context.Context, *WeatherRequest) (*WeatherResponse, error)
}

func RegisterWeatherServiceServer(s *grpc.Server, srv WeatherServiceServer) {
	s.RegisterService(&_WeatherService_serviceDesc, srv)
}

func _WeatherService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.WeatherService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServiceServer).Get(ctx, req.(*WeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WeatherService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weather.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _WeatherService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "weather.proto",
}

func init() { proto.RegisterFile("weather.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4f, 0x4d, 0x2c,
	0xc9, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xec, 0xb8,
	0xf8, 0xc2, 0x21, 0xcc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xe4,
	0xcc, 0x92, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x82, 0x8b, 0x3d,
	0x39, 0xbf, 0x34, 0xaf, 0xa4, 0xa8, 0x52, 0x82, 0x09, 0x2c, 0x0c, 0xe3, 0x2a, 0x19, 0x73, 0xf1,
	0xc3, 0xf5, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x29, 0x70, 0x71, 0x97, 0xa4, 0xe6, 0x16,
	0xa4, 0x16, 0x25, 0x96, 0x94, 0x16, 0xa5, 0x82, 0xcd, 0x61, 0x0a, 0x42, 0x16, 0x32, 0xf2, 0x82,
	0x5b, 0x1a, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0xc1, 0xc5, 0xec, 0x9e, 0x5a, 0x22,
	0x24, 0xae, 0x07, 0x73, 0x26, 0xaa, 0xa3, 0xa4, 0x24, 0x30, 0x25, 0x20, 0xb6, 0x39, 0x71, 0x46,
	0xc1, 0xfc, 0x92, 0xc4, 0x06, 0xf6, 0x9b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x38, 0x13,
	0xa5, 0xec, 0x00, 0x00, 0x00,
}
