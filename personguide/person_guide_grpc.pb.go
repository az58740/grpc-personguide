// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0--rc1
// source: person_guide.proto

package personguide

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PersonGuide_GetPhone_FullMethodName      = "/personguide.PersonGuide/GetPhone"
	PersonGuide_ListPersons_FullMethodName   = "/personguide.PersonGuide/ListPersons"
	PersonGuide_RecordPersons_FullMethodName = "/personguide.PersonGuide/RecordPersons"
	PersonGuide_RoutePhones_FullMethodName   = "/personguide.PersonGuide/RoutePhones"
)

// PersonGuideClient is the client API for PersonGuide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PersonGuideClient interface {
	// Unary RPC=> Obtains the PhoneNumber from the given Person.
	GetPhone(ctx context.Context, in *Person, opts ...grpc.CallOption) (*PhoneNumber, error)
	// Server_side streaming => Obtains the Persons related to the adress.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field).
	ListPersons(ctx context.Context, in *Adress, opts ...grpc.CallOption) (PersonGuide_ListPersonsClient, error)
	// Client_side streaming=> Accepts a stream of Persons on a route being traversed, returning a
	// AddressBook when traversal is completed.
	RecordPersons(ctx context.Context, opts ...grpc.CallOption) (PersonGuide_RecordPersonsClient, error)
	// A bidirectional streaming RPC => Accepts a stream of Person sent while a route is being traversed,
	// while receiving Phone Numbers (e.g. from other users).
	RoutePhones(ctx context.Context, opts ...grpc.CallOption) (PersonGuide_RoutePhonesClient, error)
}

type personGuideClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonGuideClient(cc grpc.ClientConnInterface) PersonGuideClient {
	return &personGuideClient{cc}
}

func (c *personGuideClient) GetPhone(ctx context.Context, in *Person, opts ...grpc.CallOption) (*PhoneNumber, error) {
	out := new(PhoneNumber)
	err := c.cc.Invoke(ctx, PersonGuide_GetPhone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personGuideClient) ListPersons(ctx context.Context, in *Adress, opts ...grpc.CallOption) (PersonGuide_ListPersonsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersonGuide_ServiceDesc.Streams[0], PersonGuide_ListPersons_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &personGuideListPersonsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PersonGuide_ListPersonsClient interface {
	Recv() (*Person, error)
	grpc.ClientStream
}

type personGuideListPersonsClient struct {
	grpc.ClientStream
}

func (x *personGuideListPersonsClient) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personGuideClient) RecordPersons(ctx context.Context, opts ...grpc.CallOption) (PersonGuide_RecordPersonsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersonGuide_ServiceDesc.Streams[1], PersonGuide_RecordPersons_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &personGuideRecordPersonsClient{stream}
	return x, nil
}

type PersonGuide_RecordPersonsClient interface {
	Send(*Person) error
	CloseAndRecv() (*AddressBook, error)
	grpc.ClientStream
}

type personGuideRecordPersonsClient struct {
	grpc.ClientStream
}

func (x *personGuideRecordPersonsClient) Send(m *Person) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personGuideRecordPersonsClient) CloseAndRecv() (*AddressBook, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddressBook)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *personGuideClient) RoutePhones(ctx context.Context, opts ...grpc.CallOption) (PersonGuide_RoutePhonesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PersonGuide_ServiceDesc.Streams[2], PersonGuide_RoutePhones_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &personGuideRoutePhonesClient{stream}
	return x, nil
}

type PersonGuide_RoutePhonesClient interface {
	Send(*Person) error
	Recv() (*PhoneNumber, error)
	grpc.ClientStream
}

type personGuideRoutePhonesClient struct {
	grpc.ClientStream
}

func (x *personGuideRoutePhonesClient) Send(m *Person) error {
	return x.ClientStream.SendMsg(m)
}

func (x *personGuideRoutePhonesClient) Recv() (*PhoneNumber, error) {
	m := new(PhoneNumber)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PersonGuideServer is the server API for PersonGuide service.
// All implementations must embed UnimplementedPersonGuideServer
// for forward compatibility
type PersonGuideServer interface {
	// Unary RPC=> Obtains the PhoneNumber from the given Person.
	GetPhone(context.Context, *Person) (*PhoneNumber, error)
	// Server_side streaming => Obtains the Persons related to the adress.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field).
	ListPersons(*Adress, PersonGuide_ListPersonsServer) error
	// Client_side streaming=> Accepts a stream of Persons on a route being traversed, returning a
	// AddressBook when traversal is completed.
	RecordPersons(PersonGuide_RecordPersonsServer) error
	// A bidirectional streaming RPC => Accepts a stream of Person sent while a route is being traversed,
	// while receiving Phone Numbers (e.g. from other users).
	RoutePhones(PersonGuide_RoutePhonesServer) error
	mustEmbedUnimplementedPersonGuideServer()
}

// UnimplementedPersonGuideServer must be embedded to have forward compatible implementations.
type UnimplementedPersonGuideServer struct {
}

func (UnimplementedPersonGuideServer) GetPhone(context.Context, *Person) (*PhoneNumber, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPhone not implemented")
}
func (UnimplementedPersonGuideServer) ListPersons(*Adress, PersonGuide_ListPersonsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPersons not implemented")
}
func (UnimplementedPersonGuideServer) RecordPersons(PersonGuide_RecordPersonsServer) error {
	return status.Errorf(codes.Unimplemented, "method RecordPersons not implemented")
}
func (UnimplementedPersonGuideServer) RoutePhones(PersonGuide_RoutePhonesServer) error {
	return status.Errorf(codes.Unimplemented, "method RoutePhones not implemented")
}
func (UnimplementedPersonGuideServer) mustEmbedUnimplementedPersonGuideServer() {}

// UnsafePersonGuideServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PersonGuideServer will
// result in compilation errors.
type UnsafePersonGuideServer interface {
	mustEmbedUnimplementedPersonGuideServer()
}

func RegisterPersonGuideServer(s grpc.ServiceRegistrar, srv PersonGuideServer) {
	s.RegisterService(&PersonGuide_ServiceDesc, srv)
}

func _PersonGuide_GetPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonGuideServer).GetPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PersonGuide_GetPhone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonGuideServer).GetPhone(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersonGuide_ListPersons_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Adress)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PersonGuideServer).ListPersons(m, &personGuideListPersonsServer{stream})
}

type PersonGuide_ListPersonsServer interface {
	Send(*Person) error
	grpc.ServerStream
}

type personGuideListPersonsServer struct {
	grpc.ServerStream
}

func (x *personGuideListPersonsServer) Send(m *Person) error {
	return x.ServerStream.SendMsg(m)
}

func _PersonGuide_RecordPersons_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonGuideServer).RecordPersons(&personGuideRecordPersonsServer{stream})
}

type PersonGuide_RecordPersonsServer interface {
	SendAndClose(*AddressBook) error
	Recv() (*Person, error)
	grpc.ServerStream
}

type personGuideRecordPersonsServer struct {
	grpc.ServerStream
}

func (x *personGuideRecordPersonsServer) SendAndClose(m *AddressBook) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personGuideRecordPersonsServer) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PersonGuide_RoutePhones_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PersonGuideServer).RoutePhones(&personGuideRoutePhonesServer{stream})
}

type PersonGuide_RoutePhonesServer interface {
	Send(*PhoneNumber) error
	Recv() (*Person, error)
	grpc.ServerStream
}

type personGuideRoutePhonesServer struct {
	grpc.ServerStream
}

func (x *personGuideRoutePhonesServer) Send(m *PhoneNumber) error {
	return x.ServerStream.SendMsg(m)
}

func (x *personGuideRoutePhonesServer) Recv() (*Person, error) {
	m := new(Person)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PersonGuide_ServiceDesc is the grpc.ServiceDesc for PersonGuide service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PersonGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "personguide.PersonGuide",
	HandlerType: (*PersonGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPhone",
			Handler:    _PersonGuide_GetPhone_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPersons",
			Handler:       _PersonGuide_ListPersons_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecordPersons",
			Handler:       _PersonGuide_RecordPersons_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "RoutePhones",
			Handler:       _PersonGuide_RoutePhones_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "person_guide.proto",
}
