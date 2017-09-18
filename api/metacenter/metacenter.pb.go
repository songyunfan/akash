// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/metacenter/metacenter.proto

/*
Package metacenter is a generated protocol buffer package.

It is generated from these files:
	api/metacenter/metacenter.proto

It has these top-level messages:
	UpdateRequest
	UpdateResponse
	StatusRequest
	StatusResponse
	EventsRequest
	Event
*/
package metacenter

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

type UpdateRequest struct {
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StatusRequest struct {
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StatusResponse struct {
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type EventsRequest struct {
}

func (m *EventsRequest) Reset()                    { *m = EventsRequest{} }
func (m *EventsRequest) String() string            { return proto.CompactTextString(m) }
func (*EventsRequest) ProtoMessage()               {}
func (*EventsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Event struct {
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*StatusRequest)(nil), "StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "StatusResponse")
	proto.RegisterType((*EventsRequest)(nil), "EventsRequest")
	proto.RegisterType((*Event)(nil), "Event")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metacenter service

type MetacenterClient interface {
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Events(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (Metacenter_EventsClient, error)
}

type metacenterClient struct {
	cc *grpc.ClientConn
}

func NewMetacenterClient(cc *grpc.ClientConn) MetacenterClient {
	return &metacenterClient{cc}
}

func (c *metacenterClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/Metacenter/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metacenterClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/Metacenter/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metacenterClient) Events(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (Metacenter_EventsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Metacenter_serviceDesc.Streams[0], c.cc, "/Metacenter/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &metacenterEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Metacenter_EventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type metacenterEventsClient struct {
	grpc.ClientStream
}

func (x *metacenterEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Metacenter service

type MetacenterServer interface {
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	Events(*EventsRequest, Metacenter_EventsServer) error
}

func RegisterMetacenterServer(s *grpc.Server, srv MetacenterServer) {
	s.RegisterService(&_Metacenter_serviceDesc, srv)
}

func _Metacenter_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetacenterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Metacenter/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetacenterServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metacenter_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetacenterServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Metacenter/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetacenterServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metacenter_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetacenterServer).Events(m, &metacenterEventsServer{stream})
}

type Metacenter_EventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type metacenterEventsServer struct {
	grpc.ServerStream
}

func (x *metacenterEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Metacenter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Metacenter",
	HandlerType: (*MetacenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Metacenter_Update_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Metacenter_Status_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _Metacenter_Events_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/metacenter/metacenter.proto",
}

func init() { proto.RegisterFile("api/metacenter/metacenter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x2c, 0xc8, 0xd4,
	0xcf, 0x4d, 0x2d, 0x49, 0x4c, 0x4e, 0xcd, 0x2b, 0x49, 0x2d, 0x42, 0x62, 0xea, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x2b, 0xf1, 0x73, 0xf1, 0x86, 0x16, 0xa4, 0x24, 0x96, 0xa4, 0x06, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x28, 0x09, 0x70, 0xf1, 0xc1, 0x04, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x41,
	0x4a, 0x82, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x91, 0x94, 0xc0, 0x04, 0x10, 0x4a, 0x5c, 0xcb, 0x52,
	0xf3, 0x4a, 0xe0, 0x4a, 0xd8, 0xb9, 0x58, 0xc1, 0x02, 0x46, 0x6d, 0x8c, 0x5c, 0x5c, 0xbe, 0x70,
	0x4b, 0x85, 0x34, 0xb9, 0xd8, 0x20, 0xa6, 0x0b, 0xf1, 0xe9, 0xa1, 0xd8, 0x2b, 0xc5, 0xaf, 0x87,
	0x6a, 0x2d, 0x48, 0x29, 0xc4, 0x16, 0x21, 0x3e, 0x3d, 0x14, 0xfb, 0xa5, 0xf8, 0xf5, 0x50, 0xad,
	0x17, 0x52, 0xe2, 0x62, 0x83, 0x58, 0x2f, 0xc4, 0xa7, 0x87, 0xe2, 0x0e, 0x29, 0x36, 0x08, 0xdf,
	0x80, 0x31, 0x89, 0x0d, 0xec, 0x5f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xe3, 0xdd,
	0x18, 0x12, 0x01, 0x00, 0x00,
}
