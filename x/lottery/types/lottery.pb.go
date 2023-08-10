// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmoslottery/lottery/lottery.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Lottery struct {
	Index string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Fee   types.Coin `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee"`
	Pool  types.Coin `protobuf:"bytes,3,opt,name=pool,proto3" json:"pool"`
}

func (m *Lottery) Reset()         { *m = Lottery{} }
func (m *Lottery) String() string { return proto.CompactTextString(m) }
func (*Lottery) ProtoMessage()    {}
func (*Lottery) Descriptor() ([]byte, []int) {
	return fileDescriptor_da23981c85921b2d, []int{0}
}
func (m *Lottery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lottery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lottery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lottery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lottery.Merge(m, src)
}
func (m *Lottery) XXX_Size() int {
	return m.Size()
}
func (m *Lottery) XXX_DiscardUnknown() {
	xxx_messageInfo_Lottery.DiscardUnknown(m)
}

var xxx_messageInfo_Lottery proto.InternalMessageInfo

func (m *Lottery) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Lottery) GetFee() types.Coin {
	if m != nil {
		return m.Fee
	}
	return types.Coin{}
}

func (m *Lottery) GetPool() types.Coin {
	if m != nil {
		return m.Pool
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Lottery)(nil), "cosmoslottery.lottery.Lottery")
}

func init() {
	proto.RegisterFile("cosmoslottery/lottery/lottery.proto", fileDescriptor_da23981c85921b2d)
}

var fileDescriptor_da23981c85921b2d = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xce, 0xc9, 0x2f, 0x29, 0x49, 0x2d, 0xaa, 0xd4, 0x47, 0xa3, 0xf5, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x44, 0x51, 0x14, 0xe9, 0x41, 0x69, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0,
	0x0a, 0x7d, 0x10, 0x0b, 0xa2, 0x58, 0x4a, 0x0e, 0xa2, 0x58, 0x3f, 0x29, 0xb1, 0x38, 0x55, 0xbf,
	0xcc, 0x30, 0x29, 0xb5, 0x24, 0xd1, 0x50, 0x3f, 0x39, 0x3f, 0x33, 0x0f, 0x2a, 0xaf, 0x8f, 0xd7,
	0xc6, 0xf8, 0x92, 0xa2, 0xc4, 0xbc, 0xe2, 0xc4, 0xe4, 0x92, 0xcc, 0x7c, 0xa8, 0x06, 0xa5, 0x76,
	0x46, 0x2e, 0x76, 0x1f, 0x88, 0xac, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x5e, 0x4a, 0x6a, 0x85, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x23, 0x64, 0xc8, 0xc5, 0x9c, 0x96, 0x9a, 0x2a, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa9, 0x07, 0xb1, 0x40, 0x0f, 0xe4, 0x00, 0x3d, 0xa8, 0x03,
	0xf4, 0x9c, 0xf3, 0x33, 0xf3, 0x9c, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x02, 0xa9, 0x15, 0x32,
	0xe6, 0x62, 0x29, 0xc8, 0xcf, 0xcf, 0x91, 0x60, 0x26, 0x4e, 0x0f, 0x58, 0xb1, 0x93, 0xc5, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x41, 0x3d, 0xad, 0x0b, 0xf3, 0x4d, 0x05,
	0xdc, 0x5f, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xaf, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xf0, 0xa0, 0x33, 0x4f, 0x6f, 0x01, 0x00, 0x00,
}

func (m *Lottery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lottery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lottery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Pool.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLottery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLottery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintLottery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLottery(dAtA []byte, offset int, v uint64) int {
	offset -= sovLottery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Lottery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovLottery(uint64(l))
	}
	l = m.Fee.Size()
	n += 1 + l + sovLottery(uint64(l))
	l = m.Pool.Size()
	n += 1 + l + sovLottery(uint64(l))
	return n
}

func sovLottery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLottery(x uint64) (n int) {
	return sovLottery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Lottery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLottery
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
			return fmt.Errorf("proto: Lottery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lottery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
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
				return ErrInvalidLengthLottery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLottery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
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
				return ErrInvalidLengthLottery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLottery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
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
				return ErrInvalidLengthLottery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLottery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLottery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLottery
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
func skipLottery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLottery
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
					return 0, ErrIntOverflowLottery
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
					return 0, ErrIntOverflowLottery
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
				return 0, ErrInvalidLengthLottery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLottery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLottery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLottery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLottery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLottery = fmt.Errorf("proto: unexpected end of group")
)