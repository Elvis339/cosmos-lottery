// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmoslottery/lottery/active_lottery.proto

package types

import (
	fmt "fmt"
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

type ActiveLottery struct {
	LotteryId uint64 `protobuf:"varint,1,opt,name=lotteryId,proto3" json:"lotteryId,omitempty"`
}

func (m *ActiveLottery) Reset()         { *m = ActiveLottery{} }
func (m *ActiveLottery) String() string { return proto.CompactTextString(m) }
func (*ActiveLottery) ProtoMessage()    {}
func (*ActiveLottery) Descriptor() ([]byte, []int) {
	return fileDescriptor_730521fa2871072e, []int{0}
}
func (m *ActiveLottery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActiveLottery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActiveLottery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActiveLottery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActiveLottery.Merge(m, src)
}
func (m *ActiveLottery) XXX_Size() int {
	return m.Size()
}
func (m *ActiveLottery) XXX_DiscardUnknown() {
	xxx_messageInfo_ActiveLottery.DiscardUnknown(m)
}

var xxx_messageInfo_ActiveLottery proto.InternalMessageInfo

func (m *ActiveLottery) GetLotteryId() uint64 {
	if m != nil {
		return m.LotteryId
	}
	return 0
}

func init() {
	proto.RegisterType((*ActiveLottery)(nil), "cosmoslottery.lottery.ActiveLottery")
}

func init() {
	proto.RegisterFile("cosmoslottery/lottery/active_lottery.proto", fileDescriptor_730521fa2871072e)
}

var fileDescriptor_730521fa2871072e = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xce, 0xc9, 0x2f, 0x29, 0x49, 0x2d, 0xaa, 0xd4, 0x87, 0xd1, 0x89, 0xc9, 0x25, 0x99,
	0x65, 0xa9, 0xf1, 0x50, 0xae, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x28, 0x8a, 0x5a, 0x3d,
	0x28, 0xad, 0xa4, 0xcb, 0xc5, 0xeb, 0x08, 0x56, 0xee, 0x03, 0x11, 0x10, 0x92, 0xe1, 0xe2, 0x84,
	0xca, 0x79, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x21, 0x04, 0x9c, 0x2c, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x0e, 0x62, 0xbe, 0x2e, 0xcc, 0x11, 0x15,
	0x70, 0xe7, 0x94, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x9d, 0x61, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x40, 0xf8, 0x8c, 0xb4, 0x00, 0x00, 0x00,
}

func (m *ActiveLottery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActiveLottery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActiveLottery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LotteryId != 0 {
		i = encodeVarintActiveLottery(dAtA, i, uint64(m.LotteryId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintActiveLottery(dAtA []byte, offset int, v uint64) int {
	offset -= sovActiveLottery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActiveLottery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LotteryId != 0 {
		n += 1 + sovActiveLottery(uint64(m.LotteryId))
	}
	return n
}

func sovActiveLottery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozActiveLottery(x uint64) (n int) {
	return sovActiveLottery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActiveLottery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowActiveLottery
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
			return fmt.Errorf("proto: ActiveLottery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActiveLottery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LotteryId", wireType)
			}
			m.LotteryId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowActiveLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LotteryId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipActiveLottery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthActiveLottery
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
func skipActiveLottery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowActiveLottery
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
					return 0, ErrIntOverflowActiveLottery
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
					return 0, ErrIntOverflowActiveLottery
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
				return 0, ErrInvalidLengthActiveLottery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupActiveLottery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthActiveLottery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthActiveLottery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowActiveLottery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupActiveLottery = fmt.Errorf("proto: unexpected end of group")
)