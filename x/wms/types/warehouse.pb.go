// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kewr/wms/warehouse.proto

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

type Warehouse struct {
	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Capacity uint64 `protobuf:"varint,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Creator  string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Warehouse) Reset()         { *m = Warehouse{} }
func (m *Warehouse) String() string { return proto.CompactTextString(m) }
func (*Warehouse) ProtoMessage()    {}
func (*Warehouse) Descriptor() ([]byte, []int) {
	return fileDescriptor_692d046e0d5ed8c3, []int{0}
}
func (m *Warehouse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Warehouse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Warehouse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Warehouse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Warehouse.Merge(m, src)
}
func (m *Warehouse) XXX_Size() int {
	return m.Size()
}
func (m *Warehouse) XXX_DiscardUnknown() {
	xxx_messageInfo_Warehouse.DiscardUnknown(m)
}

var xxx_messageInfo_Warehouse proto.InternalMessageInfo

func (m *Warehouse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Warehouse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Warehouse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Warehouse) GetCapacity() uint64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Warehouse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Warehouse)(nil), "kewr.wms.Warehouse")
}

func init() { proto.RegisterFile("kewr/wms/warehouse.proto", fileDescriptor_692d046e0d5ed8c3) }

var fileDescriptor_692d046e0d5ed8c3 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0x4e, 0x2d, 0x2f,
	0xd2, 0x2f, 0xcf, 0x2d, 0xd6, 0x2f, 0x4f, 0x2c, 0x4a, 0xcd, 0xc8, 0x2f, 0x2d, 0x4e, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0xc9, 0xe8, 0x95, 0xe7, 0x16, 0x2b, 0x35, 0x32, 0x72,
	0x71, 0x86, 0xc3, 0x64, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x82, 0x98, 0x32, 0x53, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18,
	0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x29, 0x2e, 0x8e, 0x9c, 0xfc, 0xe4, 0xc4, 0x92, 0xcc, 0xfc,
	0x3c, 0x09, 0x66, 0xb0, 0x38, 0x9c, 0x0f, 0x92, 0x4b, 0x4e, 0x2c, 0x48, 0x4c, 0xce, 0x2c, 0xa9,
	0x94, 0x60, 0x01, 0x9b, 0x02, 0xe7, 0x0b, 0x49, 0x70, 0xb1, 0x27, 0x17, 0xa5, 0x26, 0x96, 0xe4,
	0x17, 0x49, 0xb0, 0x82, 0xb5, 0xc1, 0xb8, 0x4e, 0x5a, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0x25, 0x00, 0xf6, 0x41, 0x05, 0xd8, 0x0f, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49,
	0x6c, 0x60, 0x0f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x76, 0x89, 0xc9, 0xdc, 0x00,
	0x00, 0x00,
}

func (m *Warehouse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Warehouse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Warehouse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWarehouse(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Capacity != 0 {
		i = encodeVarintWarehouse(dAtA, i, uint64(m.Capacity))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(dAtA[i:], m.Location)
		i = encodeVarintWarehouse(dAtA, i, uint64(len(m.Location)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWarehouse(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintWarehouse(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintWarehouse(dAtA []byte, offset int, v uint64) int {
	offset -= sovWarehouse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Warehouse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovWarehouse(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWarehouse(uint64(l))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovWarehouse(uint64(l))
	}
	if m.Capacity != 0 {
		n += 1 + sovWarehouse(uint64(m.Capacity))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWarehouse(uint64(l))
	}
	return n
}

func sovWarehouse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWarehouse(x uint64) (n int) {
	return sovWarehouse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Warehouse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWarehouse
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
			return fmt.Errorf("proto: Warehouse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Warehouse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWarehouse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWarehouse
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
				return ErrInvalidLengthWarehouse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWarehouse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWarehouse
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
				return ErrInvalidLengthWarehouse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWarehouse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Capacity", wireType)
			}
			m.Capacity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWarehouse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Capacity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWarehouse
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
				return ErrInvalidLengthWarehouse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWarehouse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWarehouse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWarehouse
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
func skipWarehouse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWarehouse
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
					return 0, ErrIntOverflowWarehouse
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
					return 0, ErrIntOverflowWarehouse
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
				return 0, ErrInvalidLengthWarehouse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWarehouse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWarehouse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWarehouse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWarehouse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWarehouse = fmt.Errorf("proto: unexpected end of group")
)
