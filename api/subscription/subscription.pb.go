// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: subscription.proto

package subscription

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SubsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sec    int64    `protobuf:"varint,1,opt,name=sec,proto3" json:"sec,omitempty"`
	Sports []string `protobuf:"bytes,2,rep,name=sports,proto3" json:"sports,omitempty"`
}

func (x *SubsRequest) Reset() {
	*x = SubsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubsRequest) ProtoMessage() {}

func (x *SubsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubsRequest.ProtoReflect.Descriptor instead.
func (*SubsRequest) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *SubsRequest) GetSec() int64 {
	if x != nil {
		return x.Sec
	}
	return 0
}

func (x *SubsRequest) GetSports() []string {
	if x != nil {
		return x.Sports
	}
	return nil
}

type SubsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sports map[string]float64 `protobuf:"bytes,1,rep,name=sports,proto3" json:"sports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *SubsResponse) Reset() {
	*x = SubsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubsResponse) ProtoMessage() {}

func (x *SubsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubsResponse.ProtoReflect.Descriptor instead.
func (*SubsResponse) Descriptor() ([]byte, []int) {
	return file_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *SubsResponse) GetSports() map[string]float64 {
	if x != nil {
		return x.Sports
	}
	return nil
}

var File_subscription_proto protoreflect.FileDescriptor

var file_subscription_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x73, 0x65, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x0c,
	0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b,
	0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x65, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x4f, 0x6e, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4c, 0x69, 0x6e, 0x65,
	0x73, 0x12, 0x19, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscription_proto_rawDescOnce sync.Once
	file_subscription_proto_rawDescData = file_subscription_proto_rawDesc
)

func file_subscription_proto_rawDescGZIP() []byte {
	file_subscription_proto_rawDescOnce.Do(func() {
		file_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscription_proto_rawDescData)
	})
	return file_subscription_proto_rawDescData
}

var file_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_subscription_proto_goTypes = []interface{}{
	(*SubsRequest)(nil),  // 0: subscription.SubsRequest
	(*SubsResponse)(nil), // 1: subscription.SubsResponse
	nil,                  // 2: subscription.SubsResponse.SportsEntry
}
var file_subscription_proto_depIdxs = []int32{
	2, // 0: subscription.SubsResponse.sports:type_name -> subscription.SubsResponse.SportsEntry
	0, // 1: subscription.Subscribtion.SubscribeOnSportsLines:input_type -> subscription.SubsRequest
	1, // 2: subscription.Subscribtion.SubscribeOnSportsLines:output_type -> subscription.SubsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_subscription_proto_init() }
func file_subscription_proto_init() {
	if File_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subscription_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subscription_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_subscription_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscription_proto_goTypes,
		DependencyIndexes: file_subscription_proto_depIdxs,
		MessageInfos:      file_subscription_proto_msgTypes,
	}.Build()
	File_subscription_proto = out.File
	file_subscription_proto_rawDesc = nil
	file_subscription_proto_goTypes = nil
	file_subscription_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SubscribtionClient is the client API for Subscribtion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscribtionClient interface {
	SubscribeOnSportsLines(ctx context.Context, opts ...grpc.CallOption) (Subscribtion_SubscribeOnSportsLinesClient, error)
}

type subscribtionClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscribtionClient(cc grpc.ClientConnInterface) SubscribtionClient {
	return &subscribtionClient{cc}
}

func (c *subscribtionClient) SubscribeOnSportsLines(ctx context.Context, opts ...grpc.CallOption) (Subscribtion_SubscribeOnSportsLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Subscribtion_serviceDesc.Streams[0], "/subscription.Subscribtion/SubscribeOnSportsLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscribtionSubscribeOnSportsLinesClient{stream}
	return x, nil
}

type Subscribtion_SubscribeOnSportsLinesClient interface {
	Send(*SubsRequest) error
	Recv() (*SubsResponse, error)
	grpc.ClientStream
}

type subscribtionSubscribeOnSportsLinesClient struct {
	grpc.ClientStream
}

func (x *subscribtionSubscribeOnSportsLinesClient) Send(m *SubsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *subscribtionSubscribeOnSportsLinesClient) Recv() (*SubsResponse, error) {
	m := new(SubsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SubscribtionServer is the server API for Subscribtion service.
type SubscribtionServer interface {
	SubscribeOnSportsLines(Subscribtion_SubscribeOnSportsLinesServer) error
}

// UnimplementedSubscribtionServer can be embedded to have forward compatible implementations.
type UnimplementedSubscribtionServer struct {
}

func (*UnimplementedSubscribtionServer) SubscribeOnSportsLines(Subscribtion_SubscribeOnSportsLinesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeOnSportsLines not implemented")
}

func RegisterSubscribtionServer(s *grpc.Server, srv SubscribtionServer) {
	s.RegisterService(&_Subscribtion_serviceDesc, srv)
}

func _Subscribtion_SubscribeOnSportsLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SubscribtionServer).SubscribeOnSportsLines(&subscribtionSubscribeOnSportsLinesServer{stream})
}

type Subscribtion_SubscribeOnSportsLinesServer interface {
	Send(*SubsResponse) error
	Recv() (*SubsRequest, error)
	grpc.ServerStream
}

type subscribtionSubscribeOnSportsLinesServer struct {
	grpc.ServerStream
}

func (x *subscribtionSubscribeOnSportsLinesServer) Send(m *SubsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *subscribtionSubscribeOnSportsLinesServer) Recv() (*SubsRequest, error) {
	m := new(SubsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Subscribtion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "subscription.Subscribtion",
	HandlerType: (*SubscribtionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeOnSportsLines",
			Handler:       _Subscribtion_SubscribeOnSportsLines_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "subscription.proto",
}
