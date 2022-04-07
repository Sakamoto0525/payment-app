// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.18.1
// source: pay.proto

package api

import (
	context "context"
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

// カード決済に使用するフィールドを定義
type PayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Token       string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Amount      int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PayRequest) Reset() {
	*x = PayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayRequest) ProtoMessage() {}

func (x *PayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayRequest.ProtoReflect.Descriptor instead.
func (*PayRequest) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{0}
}

func (x *PayRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PayRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PayRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PayRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// カード決済後のレスポンスを定義
type PayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paid     bool  `protobuf:"varint,1,opt,name=paid,proto3" json:"paid,omitempty"`
	Amount   int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Captured bool  `protobuf:"varint,3,opt,name=captured,proto3" json:"captured,omitempty"`
}

func (x *PayResponse) Reset() {
	*x = PayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayResponse) ProtoMessage() {}

func (x *PayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayResponse.ProtoReflect.Descriptor instead.
func (*PayResponse) Descriptor() ([]byte, []int) {
	return file_pay_proto_rawDescGZIP(), []int{1}
}

func (x *PayResponse) GetPaid() bool {
	if x != nil {
		return x.Paid
	}
	return false
}

func (x *PayResponse) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PayResponse) GetCaptured() bool {
	if x != nil {
		return x.Captured
	}
	return false
}

var File_pay_proto protoreflect.FileDescriptor

var file_pay_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x22, 0x80, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x0b, 0x50, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x70, 0x61, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x32, 0x49, 0x0a, 0x0a, 0x50, 0x61, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x3b, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x70, 0x2e, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07,
	0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pay_proto_rawDescOnce sync.Once
	file_pay_proto_rawDescData = file_pay_proto_rawDesc
)

func file_pay_proto_rawDescGZIP() []byte {
	file_pay_proto_rawDescOnce.Do(func() {
		file_pay_proto_rawDescData = protoimpl.X.CompressGZIP(file_pay_proto_rawDescData)
	})
	return file_pay_proto_rawDescData
}

var file_pay_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pay_proto_goTypes = []interface{}{
	(*PayRequest)(nil),  // 0: paymentapp.PayRequest
	(*PayResponse)(nil), // 1: paymentapp.PayResponse
}
var file_pay_proto_depIdxs = []int32{
	0, // 0: paymentapp.PayManager.Charge:input_type -> paymentapp.PayRequest
	1, // 1: paymentapp.PayManager.Charge:output_type -> paymentapp.PayResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pay_proto_init() }
func file_pay_proto_init() {
	if File_pay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayRequest); i {
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
		file_pay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayResponse); i {
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
			RawDescriptor: file_pay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pay_proto_goTypes,
		DependencyIndexes: file_pay_proto_depIdxs,
		MessageInfos:      file_pay_proto_msgTypes,
	}.Build()
	File_pay_proto = out.File
	file_pay_proto_rawDesc = nil
	file_pay_proto_goTypes = nil
	file_pay_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PayManagerClient is the client API for PayManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PayManagerClient interface {
	// 支払いを行うサービスを定義
	Charge(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error)
}

type payManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewPayManagerClient(cc grpc.ClientConnInterface) PayManagerClient {
	return &payManagerClient{cc}
}

func (c *payManagerClient) Charge(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error) {
	out := new(PayResponse)
	err := c.cc.Invoke(ctx, "/paymentapp.PayManager/Charge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayManagerServer is the server API for PayManager service.
type PayManagerServer interface {
	// 支払いを行うサービスを定義
	Charge(context.Context, *PayRequest) (*PayResponse, error)
}

// UnimplementedPayManagerServer can be embedded to have forward compatible implementations.
type UnimplementedPayManagerServer struct {
}

func (*UnimplementedPayManagerServer) Charge(context.Context, *PayRequest) (*PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Charge not implemented")
}

func RegisterPayManagerServer(s *grpc.Server, srv PayManagerServer) {
	s.RegisterService(&_PayManager_serviceDesc, srv)
}

func _PayManager_Charge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayManagerServer).Charge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/paymentapp.PayManager/Charge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayManagerServer).Charge(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PayManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "paymentapp.PayManager",
	HandlerType: (*PayManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Charge",
			Handler:    _PayManager_Charge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay.proto",
}