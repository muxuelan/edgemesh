// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tcpproxy.proto

package tcp_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TCPProxy_Type int32

const (
	TCPProxy_CONNECT TCPProxy_Type = 0
	TCPProxy_SUCCESS TCPProxy_Type = 1
	TCPProxy_FAILED  TCPProxy_Type = 2
)

var TCPProxy_Type_name = map[int32]string{
	0: "CONNECT",
	1: "SUCCESS",
	2: "FAILED",
}

var TCPProxy_Type_value = map[string]int32{
	"CONNECT": 0,
	"SUCCESS": 1,
	"FAILED":  2,
}

func (x TCPProxy_Type) Enum() *TCPProxy_Type {
	p := new(TCPProxy_Type)
	*p = x
	return p
}

func (x TCPProxy_Type) String() string {
	return proto.EnumName(TCPProxy_Type_name, int32(x))
}

func (x *TCPProxy_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TCPProxy_Type_value, data, "TCPProxy_Type")
	if err != nil {
		return err
	}
	*x = TCPProxy_Type(value)
	return nil
}

func (TCPProxy_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d8f4eea48351cdc7, []int{0, 0}
}

type TCPProxy struct {
	Type                 *TCPProxy_Type `protobuf:"varint,1,opt,name=type,enum=tcp.pb.TCPProxy_Type" json:"type,omitempty"`
	Nodename             *string        `protobuf:"bytes,2,opt,name=nodename" json:"nodename,omitempty"`
	Ip                   *string        `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	Port                 *int32         `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TCPProxy) Reset()         { *m = TCPProxy{} }
func (m *TCPProxy) String() string { return proto.CompactTextString(m) }
func (*TCPProxy) ProtoMessage()    {}
func (*TCPProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8f4eea48351cdc7, []int{0}
}
func (m *TCPProxy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TCPProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TCPProxy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TCPProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TCPProxy.Merge(m, src)
}
func (m *TCPProxy) XXX_Size() int {
	return m.Size()
}
func (m *TCPProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_TCPProxy.DiscardUnknown(m)
}

var xxx_messageInfo_TCPProxy proto.InternalMessageInfo

func (m *TCPProxy) GetType() TCPProxy_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return TCPProxy_CONNECT
}

func (m *TCPProxy) GetNodename() string {
	if m != nil && m.Nodename != nil {
		return *m.Nodename
	}
	return ""
}

func (m *TCPProxy) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *TCPProxy) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func init() {
	proto.RegisterEnum("tcp.pb.TCPProxy_Type", TCPProxy_Type_name, TCPProxy_Type_value)
	proto.RegisterType((*TCPProxy)(nil), "tcp.pb.TCPProxy")
}

func init() { proto.RegisterFile("tcpproxy.proto", fileDescriptor_d8f4eea48351cdc7) }

var fileDescriptor_d8f4eea48351cdc7 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0x2e, 0x28,
	0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2b, 0x49, 0x2e, 0xd0,
	0x2b, 0x48, 0x52, 0x5a, 0xcc, 0xc8, 0xc5, 0x11, 0xe2, 0x1c, 0x10, 0x00, 0x92, 0x12, 0xd2, 0xe4,
	0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x12, 0xd5, 0x83, 0xa8,
	0xd1, 0x83, 0xc9, 0xeb, 0x85, 0x54, 0x16, 0xa4, 0x06, 0x81, 0x95, 0x08, 0x49, 0x71, 0x71, 0xe4,
	0xe5, 0xa7, 0xa4, 0xe6, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9,
	0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x05, 0x12, 0xcc, 0x60, 0x51, 0xa6, 0xcc, 0x02, 0x21, 0x21, 0x2e,
	0x96, 0x82, 0xfc, 0xa2, 0x12, 0x09, 0x16, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x30, 0x5b, 0x49, 0x87,
	0x8b, 0x05, 0x64, 0x9a, 0x10, 0x37, 0x17, 0xbb, 0xb3, 0xbf, 0x9f, 0x9f, 0xab, 0x73, 0x88, 0x00,
	0x03, 0x88, 0x13, 0x1c, 0xea, 0xec, 0xec, 0x1a, 0x1c, 0x2c, 0xc0, 0x28, 0xc4, 0xc5, 0xc5, 0xe6,
	0xe6, 0xe8, 0xe9, 0xe3, 0xea, 0x22, 0xc0, 0xe4, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x00, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x6e, 0x03, 0x00, 0x0b, 0xd0, 0x00, 0x00, 0x00,
}

func (m *TCPProxy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TCPProxy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TCPProxy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Port != nil {
		i = encodeVarintTcpproxy(dAtA, i, uint64(*m.Port))
		i--
		dAtA[i] = 0x20
	}
	if m.Ip != nil {
		i -= len(*m.Ip)
		copy(dAtA[i:], *m.Ip)
		i = encodeVarintTcpproxy(dAtA, i, uint64(len(*m.Ip)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Nodename != nil {
		i -= len(*m.Nodename)
		copy(dAtA[i:], *m.Nodename)
		i = encodeVarintTcpproxy(dAtA, i, uint64(len(*m.Nodename)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != nil {
		i = encodeVarintTcpproxy(dAtA, i, uint64(*m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTcpproxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovTcpproxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TCPProxy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != nil {
		n += 1 + sovTcpproxy(uint64(*m.Type))
	}
	if m.Nodename != nil {
		l = len(*m.Nodename)
		n += 1 + l + sovTcpproxy(uint64(l))
	}
	if m.Ip != nil {
		l = len(*m.Ip)
		n += 1 + l + sovTcpproxy(uint64(l))
	}
	if m.Port != nil {
		n += 1 + sovTcpproxy(uint64(*m.Port))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTcpproxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTcpproxy(x uint64) (n int) {
	return sovTcpproxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TCPProxy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTcpproxy
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
			return fmt.Errorf("proto: TCPProxy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TCPProxy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var v TCPProxy_Type
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= TCPProxy_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Type = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodename", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpproxy
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
				return ErrInvalidLengthTcpproxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTcpproxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Nodename = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpproxy
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
				return ErrInvalidLengthTcpproxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTcpproxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Ip = &s
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTcpproxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Port = &v
		default:
			iNdEx = preIndex
			skippy, err := skipTcpproxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTcpproxy
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
func skipTcpproxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTcpproxy
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
					return 0, ErrIntOverflowTcpproxy
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
					return 0, ErrIntOverflowTcpproxy
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
				return 0, ErrInvalidLengthTcpproxy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTcpproxy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTcpproxy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTcpproxy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTcpproxy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTcpproxy = fmt.Errorf("proto: unexpected end of group")
)
