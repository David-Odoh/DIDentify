// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identity/identity/identity.proto

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

type Identity struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator    string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Birthdate  string `protobuf:"bytes,4,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	NationalId string `protobuf:"bytes,5,opt,name=nationalId,proto3" json:"nationalId,omitempty"`
	Verified   bool   `protobuf:"varint,6,opt,name=verified,proto3" json:"verified,omitempty"`
	Did        string `protobuf:"bytes,7,opt,name=did,proto3" json:"did,omitempty"`
}

func (m *Identity) Reset()         { *m = Identity{} }
func (m *Identity) String() string { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()    {}
func (*Identity) Descriptor() ([]byte, []int) {
	return fileDescriptor_152066f6a7536b4d, []int{0}
}
func (m *Identity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Identity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Identity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Identity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Identity.Merge(m, src)
}
func (m *Identity) XXX_Size() int {
	return m.Size()
}
func (m *Identity) XXX_DiscardUnknown() {
	xxx_messageInfo_Identity.DiscardUnknown(m)
}

var xxx_messageInfo_Identity proto.InternalMessageInfo

func (m *Identity) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Identity) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Identity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Identity) GetBirthdate() string {
	if m != nil {
		return m.Birthdate
	}
	return ""
}

func (m *Identity) GetNationalId() string {
	if m != nil {
		return m.NationalId
	}
	return ""
}

func (m *Identity) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *Identity) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func init() {
	proto.RegisterType((*Identity)(nil), "identity.identity.Identity")
}

func init() { proto.RegisterFile("identity/identity/identity.proto", fileDescriptor_152066f6a7536b4d) }

var fileDescriptor_152066f6a7536b4d = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4e, 0xc3, 0x40,
	0x10, 0x45, 0x3d, 0x8e, 0x49, 0x9c, 0x29, 0x10, 0x4c, 0x35, 0x20, 0xb4, 0xb2, 0xa8, 0x5c, 0x41,
	0x91, 0x1b, 0xd0, 0xa5, 0x75, 0x49, 0xb7, 0x61, 0x16, 0x31, 0x12, 0xd8, 0xd1, 0x32, 0x42, 0xe4,
	0x16, 0x1c, 0x86, 0x43, 0x50, 0xa6, 0xa4, 0x44, 0xf6, 0x45, 0x10, 0x2b, 0x62, 0x23, 0xd1, 0xbd,
	0xf7, 0xbf, 0xa6, 0x98, 0x8f, 0x95, 0x4a, 0x68, 0x4d, 0x6d, 0x77, 0xfd, 0x0f, 0xae, 0xb6, 0xb1,
	0xb3, 0x8e, 0x4e, 0x47, 0x3f, 0xc0, 0xe5, 0x3b, 0x60, 0xb9, 0xfe, 0x15, 0x3a, 0xc6, 0x5c, 0x85,
	0xa1, 0x82, 0xba, 0x68, 0x72, 0x15, 0x62, 0x5c, 0xdc, 0xc5, 0xe0, 0xad, 0x8b, 0x9c, 0x57, 0x50,
	0x2f, 0x9b, 0x83, 0x12, 0x61, 0xd1, 0xfa, 0xa7, 0xc0, 0xb3, 0x14, 0x27, 0xa6, 0x0b, 0x5c, 0x6e,
	0x34, 0xda, 0x83, 0x78, 0x0b, 0x5c, 0xa4, 0x62, 0x0a, 0xc8, 0x21, 0xb6, 0xde, 0xb4, 0x6b, 0xfd,
	0xe3, 0x5a, 0xf8, 0x28, 0xd5, 0x7f, 0x12, 0x3a, 0xc7, 0xf2, 0x25, 0x44, 0xbd, 0xd7, 0x20, 0x3c,
	0xaf, 0xa0, 0x2e, 0x9b, 0xd1, 0xe9, 0x04, 0x67, 0xa2, 0xc2, 0x8b, 0x74, 0xf4, 0x83, 0x37, 0xab,
	0x8f, 0xde, 0xc1, 0xbe, 0x77, 0xf0, 0xd5, 0x3b, 0x78, 0x1b, 0x5c, 0xb6, 0x1f, 0x5c, 0xf6, 0x39,
	0xb8, 0xec, 0xf6, 0x6c, 0x7c, 0xfe, 0x75, 0xda, 0xc1, 0x76, 0xdb, 0xf0, 0xbc, 0x99, 0xa7, 0x15,
	0x56, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xca, 0x1d, 0x04, 0x29, 0x01, 0x00, 0x00,
}

func (m *Identity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Identity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Identity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Verified {
		i--
		if m.Verified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.NationalId) > 0 {
		i -= len(m.NationalId)
		copy(dAtA[i:], m.NationalId)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.NationalId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Birthdate) > 0 {
		i -= len(m.Birthdate)
		copy(dAtA[i:], m.Birthdate)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Birthdate)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintIdentity(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintIdentity(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIdentity(dAtA []byte, offset int, v uint64) int {
	offset -= sovIdentity(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Identity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovIdentity(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.Birthdate)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	l = len(m.NationalId)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	if m.Verified {
		n += 2
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovIdentity(uint64(l))
	}
	return n
}

func sovIdentity(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIdentity(x uint64) (n int) {
	return sovIdentity(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Identity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentity
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
			return fmt.Errorf("proto: Identity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Identity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Birthdate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Birthdate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NationalId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NationalId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Verified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Verified = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentity
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
				return ErrInvalidLengthIdentity
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentity
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentity(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdentity
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
func skipIdentity(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdentity
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
					return 0, ErrIntOverflowIdentity
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
					return 0, ErrIntOverflowIdentity
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
				return 0, ErrInvalidLengthIdentity
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIdentity
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIdentity
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIdentity        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdentity          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIdentity = fmt.Errorf("proto: unexpected end of group")
)
