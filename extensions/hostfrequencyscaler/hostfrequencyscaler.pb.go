// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hostfrequencyscaler.proto

package hostfrequencyscaler

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Scaler_ScalerState int32

const (
	Scaler_NONE        Scaler_ScalerState = 0
	Scaler_POWER_SAVE  Scaler_ScalerState = 1
	Scaler_PERFORMANCE Scaler_ScalerState = 2
	Scaler_SCHEDUTIL   Scaler_ScalerState = 3
)

var Scaler_ScalerState_name = map[int32]string{
	0: "NONE",
	1: "POWER_SAVE",
	2: "PERFORMANCE",
	3: "SCHEDUTIL",
}

var Scaler_ScalerState_value = map[string]int32{
	"NONE":        0,
	"POWER_SAVE":  1,
	"PERFORMANCE": 2,
	"SCHEDUTIL":   3,
}

func (x Scaler_ScalerState) String() string {
	return proto.EnumName(Scaler_ScalerState_name, int32(x))
}

func (Scaler_ScalerState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3961b3c8262123e6, []int{0, 0}
}

type Scaler struct {
	State                Scaler_ScalerState `protobuf:"varint,1,opt,name=state,proto3,enum=HostFrequencyScaler.Scaler_ScalerState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Scaler) Reset()         { *m = Scaler{} }
func (m *Scaler) String() string { return proto.CompactTextString(m) }
func (*Scaler) ProtoMessage()    {}
func (*Scaler) Descriptor() ([]byte, []int) {
	return fileDescriptor_3961b3c8262123e6, []int{0}
}
func (m *Scaler) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Scaler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Scaler.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Scaler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scaler.Merge(m, src)
}
func (m *Scaler) XXX_Size() int {
	return m.Size()
}
func (m *Scaler) XXX_DiscardUnknown() {
	xxx_messageInfo_Scaler.DiscardUnknown(m)
}

var xxx_messageInfo_Scaler proto.InternalMessageInfo

func (m *Scaler) GetState() Scaler_ScalerState {
	if m != nil {
		return m.State
	}
	return Scaler_NONE
}

func (*Scaler) XXX_MessageName() string {
	return "HostFrequencyScaler.Scaler"
}
func init() {
	proto.RegisterEnum("HostFrequencyScaler.Scaler_ScalerState", Scaler_ScalerState_name, Scaler_ScalerState_value)
	golang_proto.RegisterEnum("HostFrequencyScaler.Scaler_ScalerState", Scaler_ScalerState_name, Scaler_ScalerState_value)
	proto.RegisterType((*Scaler)(nil), "HostFrequencyScaler.Scaler")
	golang_proto.RegisterType((*Scaler)(nil), "HostFrequencyScaler.Scaler")
}

func init() { proto.RegisterFile("hostfrequencyscaler.proto", fileDescriptor_3961b3c8262123e6) }
func init() { golang_proto.RegisterFile("hostfrequencyscaler.proto", fileDescriptor_3961b3c8262123e6) }

var fileDescriptor_3961b3c8262123e6 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xc8, 0x2f, 0x2e,
	0x49, 0x2b, 0x4a, 0x2d, 0x2c, 0x4d, 0xcd, 0x4b, 0xae, 0x2c, 0x4e, 0x4e, 0xcc, 0x49, 0x2d, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xf6, 0xc8, 0x2f, 0x2e, 0x71, 0x83, 0x49, 0x05, 0x83,
	0xa5, 0xa4, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3,
	0xd3, 0xf3, 0xf5, 0xc1, 0x6a, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x98, 0xa1,
	0x34, 0x81, 0x91, 0x8b, 0x0d, 0xa2, 0x53, 0xc8, 0x96, 0x8b, 0xb5, 0xb8, 0x24, 0xb1, 0x24, 0x55,
	0x82, 0x51, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x5d, 0x0f, 0x8b, 0xf1, 0x7a, 0x28, 0x54, 0x30, 0x48,
	0x79, 0x10, 0x44, 0x97, 0x92, 0x3b, 0x17, 0x37, 0x92, 0xa8, 0x10, 0x07, 0x17, 0x8b, 0x9f, 0xbf,
	0x9f, 0xab, 0x00, 0x83, 0x10, 0x1f, 0x17, 0x57, 0x80, 0x7f, 0xb8, 0x6b, 0x50, 0x7c, 0xb0, 0x63,
	0x98, 0xab, 0x00, 0xa3, 0x10, 0x3f, 0x17, 0x77, 0x80, 0x6b, 0x90, 0x9b, 0x7f, 0x90, 0xaf, 0xa3,
	0x9f, 0xb3, 0xab, 0x00, 0x93, 0x10, 0x2f, 0x17, 0x67, 0xb0, 0xb3, 0x87, 0xab, 0x4b, 0x68, 0x88,
	0xa7, 0x8f, 0x00, 0xb3, 0x93, 0xf6, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x78, 0xe0, 0xb1, 0x1c, 0xe3, 0x89, 0xc7, 0x72, 0x8c, 0x51, 0xa2, 0x7a, 0xd6, 0x58,
	0x42, 0x22, 0x89, 0x0d, 0xec, 0x0d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xf0, 0x13,
	0xa9, 0x27, 0x01, 0x00, 0x00,
}

func (m *Scaler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scaler) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Scaler) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.State != 0 {
		i = encodeVarintHostfrequencyscaler(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHostfrequencyscaler(dAtA []byte, offset int, v uint64) int {
	offset -= sovHostfrequencyscaler(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Scaler) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovHostfrequencyscaler(uint64(m.State))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHostfrequencyscaler(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHostfrequencyscaler(x uint64) (n int) {
	return sovHostfrequencyscaler(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Scaler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHostfrequencyscaler
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
			return fmt.Errorf("proto: Scaler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scaler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHostfrequencyscaler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Scaler_ScalerState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHostfrequencyscaler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHostfrequencyscaler
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHostfrequencyscaler(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHostfrequencyscaler
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
					return 0, ErrIntOverflowHostfrequencyscaler
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
					return 0, ErrIntOverflowHostfrequencyscaler
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
				return 0, ErrInvalidLengthHostfrequencyscaler
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHostfrequencyscaler
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHostfrequencyscaler
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHostfrequencyscaler        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHostfrequencyscaler          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHostfrequencyscaler = fmt.Errorf("proto: unexpected end of group")
)
