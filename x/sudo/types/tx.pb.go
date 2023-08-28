// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nibiru/sudo/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgEditSudoers: Msg to update the "Sudoers" state.
type MsgEditSudoers struct {
	// Action: identifier for the type of edit that will take place. Using this
	//   action field prevents us from needing to create several similar message
	//   types.
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	// Contracts: An input payload.
	Contracts []string `protobuf:"bytes,2,rep,name=contracts,proto3" json:"contracts,omitempty"`
	// Sender: Address for the signer of the transaction.
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgEditSudoers) Reset()         { *m = MsgEditSudoers{} }
func (m *MsgEditSudoers) String() string { return proto.CompactTextString(m) }
func (*MsgEditSudoers) ProtoMessage()    {}
func (*MsgEditSudoers) Descriptor() ([]byte, []int) {
	return fileDescriptor_a610e3c1609cdcbc, []int{0}
}
func (m *MsgEditSudoers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEditSudoers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEditSudoers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEditSudoers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEditSudoers.Merge(m, src)
}
func (m *MsgEditSudoers) XXX_Size() int {
	return m.Size()
}
func (m *MsgEditSudoers) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEditSudoers.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEditSudoers proto.InternalMessageInfo

func (m *MsgEditSudoers) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *MsgEditSudoers) GetContracts() []string {
	if m != nil {
		return m.Contracts
	}
	return nil
}

func (m *MsgEditSudoers) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgEditSudoersResponse indicates the successful execution of MsgEditSudeors.
type MsgEditSudoersResponse struct {
}

func (m *MsgEditSudoersResponse) Reset()         { *m = MsgEditSudoersResponse{} }
func (m *MsgEditSudoersResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEditSudoersResponse) ProtoMessage()    {}
func (*MsgEditSudoersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a610e3c1609cdcbc, []int{1}
}
func (m *MsgEditSudoersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEditSudoersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEditSudoersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEditSudoersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEditSudoersResponse.Merge(m, src)
}
func (m *MsgEditSudoersResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgEditSudoersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEditSudoersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEditSudoersResponse proto.InternalMessageInfo

// MsgChangeRoot: Msg to update the "Sudoers" state.
type MsgChangeRoot struct {
	// Sender: Address for the signer of the transaction.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// NewRoot: New root address.
	NewRoot string `protobuf:"bytes,2,opt,name=new_root,json=newRoot,proto3" json:"new_root,omitempty"`
}

func (m *MsgChangeRoot) Reset()         { *m = MsgChangeRoot{} }
func (m *MsgChangeRoot) String() string { return proto.CompactTextString(m) }
func (*MsgChangeRoot) ProtoMessage()    {}
func (*MsgChangeRoot) Descriptor() ([]byte, []int) {
	return fileDescriptor_a610e3c1609cdcbc, []int{2}
}
func (m *MsgChangeRoot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgChangeRoot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgChangeRoot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgChangeRoot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgChangeRoot.Merge(m, src)
}
func (m *MsgChangeRoot) XXX_Size() int {
	return m.Size()
}
func (m *MsgChangeRoot) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgChangeRoot.DiscardUnknown(m)
}

var xxx_messageInfo_MsgChangeRoot proto.InternalMessageInfo

func (m *MsgChangeRoot) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgChangeRoot) GetNewRoot() string {
	if m != nil {
		return m.NewRoot
	}
	return ""
}

// MsgChangeRootResponse indicates the successful execution of MsgChangeRoot.
type MsgChangeRootResponse struct {
}

func (m *MsgChangeRootResponse) Reset()         { *m = MsgChangeRootResponse{} }
func (m *MsgChangeRootResponse) String() string { return proto.CompactTextString(m) }
func (*MsgChangeRootResponse) ProtoMessage()    {}
func (*MsgChangeRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a610e3c1609cdcbc, []int{3}
}
func (m *MsgChangeRootResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgChangeRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgChangeRootResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgChangeRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgChangeRootResponse.Merge(m, src)
}
func (m *MsgChangeRootResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgChangeRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgChangeRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgChangeRootResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgEditSudoers)(nil), "nibiru.sudo.v1.MsgEditSudoers")
	proto.RegisterType((*MsgEditSudoersResponse)(nil), "nibiru.sudo.v1.MsgEditSudoersResponse")
	proto.RegisterType((*MsgChangeRoot)(nil), "nibiru.sudo.v1.MsgChangeRoot")
	proto.RegisterType((*MsgChangeRootResponse)(nil), "nibiru.sudo.v1.MsgChangeRootResponse")
}

func init() { proto.RegisterFile("nibiru/sudo/v1/tx.proto", fileDescriptor_a610e3c1609cdcbc) }

var fileDescriptor_a610e3c1609cdcbc = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x4e, 0xea, 0x40,
	0x18, 0x86, 0x29, 0x24, 0x9c, 0xc3, 0x9c, 0x1c, 0x16, 0x8d, 0x42, 0xa9, 0xd8, 0x60, 0x13, 0x0d,
	0x71, 0xd1, 0x09, 0x7a, 0x07, 0xa0, 0x4b, 0x5c, 0xd4, 0x9d, 0x0b, 0xc9, 0xd0, 0x4e, 0x86, 0x49,
	0x74, 0xbe, 0xa6, 0x33, 0x05, 0xdc, 0x7a, 0x05, 0x26, 0xde, 0x94, 0x4b, 0x12, 0x37, 0x2e, 0x0d,
	0x78, 0x0b, 0xee, 0x4d, 0xa7, 0xfc, 0xb4, 0x89, 0x61, 0xd7, 0xe9, 0xf3, 0xbd, 0xcf, 0x3b, 0x33,
	0x2d, 0x6a, 0x0a, 0x3e, 0xe6, 0x71, 0x82, 0x65, 0x12, 0x02, 0x9e, 0xf6, 0xb0, 0x9a, 0x7b, 0x51,
	0x0c, 0x0a, 0xcc, 0x7a, 0x06, 0xbc, 0x14, 0x78, 0xd3, 0x9e, 0x7d, 0xc0, 0x80, 0x81, 0x46, 0x38,
	0x7d, 0xca, 0xa6, 0xec, 0x36, 0x03, 0x60, 0x0f, 0x14, 0x93, 0x88, 0x63, 0x22, 0x04, 0x28, 0xa2,
	0x38, 0x08, 0x99, 0x51, 0xf7, 0x1e, 0xd5, 0x87, 0x92, 0x5d, 0x87, 0x5c, 0xdd, 0x26, 0x21, 0xd0,
	0x58, 0x9a, 0x0d, 0x54, 0x25, 0x41, 0x3a, 0x62, 0x19, 0x1d, 0xa3, 0x5b, 0xf3, 0xd7, 0x2b, 0xb3,
	0x8d, 0x6a, 0x01, 0x08, 0x15, 0x93, 0x40, 0x49, 0xab, 0xdc, 0xa9, 0x74, 0x6b, 0xfe, 0xee, 0x45,
	0x9a, 0x92, 0x54, 0x84, 0x34, 0xb6, 0x2a, 0x59, 0x2a, 0x5b, 0xb9, 0x16, 0x6a, 0x14, 0xfd, 0x3e,
	0x95, 0x11, 0x08, 0x49, 0xdd, 0x3e, 0xfa, 0x3f, 0x94, 0x6c, 0x30, 0x21, 0x82, 0x51, 0x1f, 0x40,
	0xe5, 0x14, 0x46, 0x5e, 0x61, 0xb6, 0xd0, 0x5f, 0x41, 0x67, 0xa3, 0x18, 0x40, 0x59, 0x65, 0x4d,
	0xfe, 0x08, 0x3a, 0x4b, 0x23, 0x6e, 0x13, 0x1d, 0x16, 0x1c, 0x1b, 0xf9, 0xc5, 0xb7, 0x81, 0x2a,
	0x43, 0xc9, 0xcc, 0x39, 0xfa, 0x97, 0x3f, 0x9b, 0xe3, 0x15, 0xaf, 0xcc, 0x2b, 0xee, 0xcd, 0x3e,
	0xdb, 0xcf, 0xb7, 0x7b, 0x3f, 0x79, 0x7e, 0xff, 0x7a, 0x2d, 0x1f, 0xb9, 0x2d, 0x9c, 0xff, 0x36,
	0x34, 0xe4, 0x6a, 0x24, 0xd7, 0x55, 0x0a, 0xa1, 0xdc, 0xd9, 0x8e, 0x7f, 0x11, 0xef, 0xb0, 0x7d,
	0xba, 0x17, 0x6f, 0x6b, 0x3b, 0xba, 0xd6, 0x76, 0xad, 0x42, 0x6d, 0xa0, 0x07, 0xf5, 0xfd, 0xf4,
	0xaf, 0xde, 0x96, 0x8e, 0xb1, 0x58, 0x3a, 0xc6, 0xe7, 0xd2, 0x31, 0x5e, 0x56, 0x4e, 0x69, 0xb1,
	0x72, 0x4a, 0x1f, 0x2b, 0xa7, 0x74, 0x77, 0xce, 0xb8, 0x9a, 0x24, 0x63, 0x2f, 0x80, 0x47, 0x7c,
	0xa3, 0xd3, 0x83, 0x09, 0xe1, 0x62, 0x63, 0x9a, 0x67, 0x2e, 0xf5, 0x14, 0x51, 0x39, 0xae, 0xea,
	0x7f, 0xe3, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x5c, 0x76, 0xe7, 0x7a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// EditSudoers updates the "Sudoers" state
	EditSudoers(ctx context.Context, in *MsgEditSudoers, opts ...grpc.CallOption) (*MsgEditSudoersResponse, error)
	ChangeRoot(ctx context.Context, in *MsgChangeRoot, opts ...grpc.CallOption) (*MsgChangeRootResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) EditSudoers(ctx context.Context, in *MsgEditSudoers, opts ...grpc.CallOption) (*MsgEditSudoersResponse, error) {
	out := new(MsgEditSudoersResponse)
	err := c.cc.Invoke(ctx, "/nibiru.sudo.v1.Msg/EditSudoers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ChangeRoot(ctx context.Context, in *MsgChangeRoot, opts ...grpc.CallOption) (*MsgChangeRootResponse, error) {
	out := new(MsgChangeRootResponse)
	err := c.cc.Invoke(ctx, "/nibiru.sudo.v1.Msg/ChangeRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// EditSudoers updates the "Sudoers" state
	EditSudoers(context.Context, *MsgEditSudoers) (*MsgEditSudoersResponse, error)
	ChangeRoot(context.Context, *MsgChangeRoot) (*MsgChangeRootResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) EditSudoers(ctx context.Context, req *MsgEditSudoers) (*MsgEditSudoersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditSudoers not implemented")
}
func (*UnimplementedMsgServer) ChangeRoot(ctx context.Context, req *MsgChangeRoot) (*MsgChangeRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeRoot not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_EditSudoers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditSudoers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditSudoers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nibiru.sudo.v1.Msg/EditSudoers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditSudoers(ctx, req.(*MsgEditSudoers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ChangeRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgChangeRoot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ChangeRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nibiru.sudo.v1.Msg/ChangeRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ChangeRoot(ctx, req.(*MsgChangeRoot))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nibiru.sudo.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EditSudoers",
			Handler:    _Msg_EditSudoers_Handler,
		},
		{
			MethodName: "ChangeRoot",
			Handler:    _Msg_ChangeRoot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nibiru/sudo/v1/tx.proto",
}

func (m *MsgEditSudoers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEditSudoers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEditSudoers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Contracts) > 0 {
		for iNdEx := len(m.Contracts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contracts[iNdEx])
			copy(dAtA[i:], m.Contracts[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Contracts[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Action) > 0 {
		i -= len(m.Action)
		copy(dAtA[i:], m.Action)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Action)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgEditSudoersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEditSudoersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEditSudoersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgChangeRoot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgChangeRoot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgChangeRoot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewRoot) > 0 {
		i -= len(m.NewRoot)
		copy(dAtA[i:], m.NewRoot)
		i = encodeVarintTx(dAtA, i, uint64(len(m.NewRoot)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgChangeRootResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgChangeRootResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgChangeRootResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgEditSudoers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Contracts) > 0 {
		for _, s := range m.Contracts {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgEditSudoersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgChangeRoot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.NewRoot)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgChangeRootResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgEditSudoers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEditSudoers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEditSudoers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contracts = append(m.Contracts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgEditSudoersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEditSudoersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEditSudoersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgChangeRoot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgChangeRoot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgChangeRoot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewRoot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewRoot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgChangeRootResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgChangeRootResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgChangeRootResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)