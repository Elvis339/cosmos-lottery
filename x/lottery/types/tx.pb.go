// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmoslottery/lottery/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

type MsgPlaceBet struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Bet     uint64 `protobuf:"varint,2,opt,name=bet,proto3" json:"bet,omitempty"`
}

func (m *MsgPlaceBet) Reset()         { *m = MsgPlaceBet{} }
func (m *MsgPlaceBet) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceBet) ProtoMessage()    {}
func (*MsgPlaceBet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca6f23e6281d30, []int{0}
}
func (m *MsgPlaceBet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceBet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceBet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceBet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceBet.Merge(m, src)
}
func (m *MsgPlaceBet) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceBet) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceBet.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceBet proto.InternalMessageInfo

func (m *MsgPlaceBet) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgPlaceBet) GetBet() uint64 {
	if m != nil {
		return m.Bet
	}
	return 0
}

type MsgPlaceBetResponse struct {
}

func (m *MsgPlaceBetResponse) Reset()         { *m = MsgPlaceBetResponse{} }
func (m *MsgPlaceBetResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceBetResponse) ProtoMessage()    {}
func (*MsgPlaceBetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bca6f23e6281d30, []int{1}
}
func (m *MsgPlaceBetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceBetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceBetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceBetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceBetResponse.Merge(m, src)
}
func (m *MsgPlaceBetResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceBetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceBetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceBetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgPlaceBet)(nil), "cosmoslottery.lottery.MsgPlaceBet")
	proto.RegisterType((*MsgPlaceBetResponse)(nil), "cosmoslottery.lottery.MsgPlaceBetResponse")
}

func init() { proto.RegisterFile("cosmoslottery/lottery/tx.proto", fileDescriptor_3bca6f23e6281d30) }

var fileDescriptor_3bca6f23e6281d30 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xce, 0xc9, 0x2f, 0x29, 0x49, 0x2d, 0xaa, 0xd4, 0x87, 0xd1, 0x25, 0x15, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0xa2, 0x28, 0xf2, 0x7a, 0x50, 0x5a, 0xc9, 0x92, 0x8b, 0xdb, 0xb7,
	0x38, 0x3d, 0x20, 0x27, 0x31, 0x39, 0xd5, 0x29, 0xb5, 0x44, 0x48, 0x82, 0x8b, 0x3d, 0xb9, 0x28,
	0x35, 0xb1, 0x24, 0xbf, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x12, 0xe0,
	0x62, 0x4e, 0x4a, 0x2d, 0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x09, 0x02, 0x31, 0x95, 0x44, 0xb9,
	0x84, 0x91, 0xb4, 0x06, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0x25, 0x72, 0x31, 0xfb,
	0x16, 0xa7, 0x0b, 0x45, 0x71, 0x71, 0xc0, 0x4d, 0x55, 0xd2, 0xc3, 0x6a, 0xb9, 0x1e, 0x92, 0x76,
	0x29, 0x2d, 0xc2, 0x6a, 0x60, 0x56, 0x38, 0x59, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1,
	0x1c, 0x43, 0x14, 0x34, 0x14, 0x74, 0x61, 0xde, 0xaf, 0x40, 0x04, 0x44, 0x65, 0x41, 0x6a, 0x71,
	0x12, 0x1b, 0x38, 0x30, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85, 0x60, 0x07, 0xc1, 0x2e,
	0x01, 0x00, 0x00,
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
	PlaceBet(ctx context.Context, in *MsgPlaceBet, opts ...grpc.CallOption) (*MsgPlaceBetResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PlaceBet(ctx context.Context, in *MsgPlaceBet, opts ...grpc.CallOption) (*MsgPlaceBetResponse, error) {
	out := new(MsgPlaceBetResponse)
	err := c.cc.Invoke(ctx, "/cosmoslottery.lottery.Msg/PlaceBet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	PlaceBet(context.Context, *MsgPlaceBet) (*MsgPlaceBetResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PlaceBet(ctx context.Context, req *MsgPlaceBet) (*MsgPlaceBetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBet not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PlaceBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPlaceBet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PlaceBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmoslottery.lottery.Msg/PlaceBet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PlaceBet(ctx, req.(*MsgPlaceBet))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmoslottery.lottery.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceBet",
			Handler:    _Msg_PlaceBet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmoslottery/lottery/tx.proto",
}

func (m *MsgPlaceBet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceBet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceBet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bet != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Bet))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPlaceBetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceBetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceBetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgPlaceBet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Bet != 0 {
		n += 1 + sovTx(uint64(m.Bet))
	}
	return n
}

func (m *MsgPlaceBetResponse) Size() (n int) {
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
func (m *MsgPlaceBet) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgPlaceBet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceBet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bet", wireType)
			}
			m.Bet = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bet |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgPlaceBetResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgPlaceBetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceBetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
