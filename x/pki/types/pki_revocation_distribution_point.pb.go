// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/pki_revocation_distribution_point.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type PkiRevocationDistributionPoint struct {
	Vid                  int32  `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty"`
	Label                string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	IssuerSubjectKeyID   string `protobuf:"bytes,3,opt,name=issuerSubjectKeyID,proto3" json:"issuerSubjectKeyID,omitempty"`
	Pid                  int32  `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
	IsPAA                bool   `protobuf:"varint,5,opt,name=isPAA,proto3" json:"isPAA,omitempty"`
	CrlSignerCertificate string `protobuf:"bytes,6,opt,name=crlSignerCertificate,proto3" json:"crlSignerCertificate,omitempty"`
	DataURL              string `protobuf:"bytes,7,opt,name=dataURL,proto3" json:"dataURL,omitempty"`
	DataFileSize         uint64 `protobuf:"varint,8,opt,name=dataFileSize,proto3" json:"dataFileSize,omitempty"`
	DataDigest           string `protobuf:"bytes,9,opt,name=dataDigest,proto3" json:"dataDigest,omitempty"`
	DataDigestType       uint32 `protobuf:"varint,10,opt,name=dataDigestType,proto3" json:"dataDigestType,omitempty"`
	RevocationType       uint32 `protobuf:"varint,11,opt,name=revocationType,proto3" json:"revocationType,omitempty"`
}

func (m *PkiRevocationDistributionPoint) Reset()         { *m = PkiRevocationDistributionPoint{} }
func (m *PkiRevocationDistributionPoint) String() string { return proto.CompactTextString(m) }
func (*PkiRevocationDistributionPoint) ProtoMessage()    {}
func (*PkiRevocationDistributionPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_35504fa19b856908, []int{0}
}
func (m *PkiRevocationDistributionPoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PkiRevocationDistributionPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PkiRevocationDistributionPoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PkiRevocationDistributionPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PkiRevocationDistributionPoint.Merge(m, src)
}
func (m *PkiRevocationDistributionPoint) XXX_Size() int {
	return m.Size()
}
func (m *PkiRevocationDistributionPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PkiRevocationDistributionPoint.DiscardUnknown(m)
}

var xxx_messageInfo_PkiRevocationDistributionPoint proto.InternalMessageInfo

func (m *PkiRevocationDistributionPoint) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *PkiRevocationDistributionPoint) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *PkiRevocationDistributionPoint) GetIssuerSubjectKeyID() string {
	if m != nil {
		return m.IssuerSubjectKeyID
	}
	return ""
}

func (m *PkiRevocationDistributionPoint) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *PkiRevocationDistributionPoint) GetIsPAA() bool {
	if m != nil {
		return m.IsPAA
	}
	return false
}

func (m *PkiRevocationDistributionPoint) GetCrlSignerCertificate() string {
	if m != nil {
		return m.CrlSignerCertificate
	}
	return ""
}

func (m *PkiRevocationDistributionPoint) GetDataURL() string {
	if m != nil {
		return m.DataURL
	}
	return ""
}

func (m *PkiRevocationDistributionPoint) GetDataFileSize() uint64 {
	if m != nil {
		return m.DataFileSize
	}
	return 0
}

func (m *PkiRevocationDistributionPoint) GetDataDigest() string {
	if m != nil {
		return m.DataDigest
	}
	return ""
}

func (m *PkiRevocationDistributionPoint) GetDataDigestType() uint32 {
	if m != nil {
		return m.DataDigestType
	}
	return 0
}

func (m *PkiRevocationDistributionPoint) GetRevocationType() uint32 {
	if m != nil {
		return m.RevocationType
	}
	return 0
}

func init() {
	proto.RegisterType((*PkiRevocationDistributionPoint)(nil), "zigbeealliance.distributedcomplianceledger.pki.PkiRevocationDistributionPoint")
}

func init() {
	proto.RegisterFile("pki/pki_revocation_distribution_point.proto", fileDescriptor_35504fa19b856908)
}

var fileDescriptor_35504fa19b856908 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbb, 0x8a, 0xdb, 0x40,
	0x14, 0x86, 0x3d, 0xf1, 0x7d, 0x72, 0x21, 0x0c, 0x2e, 0xa6, 0x1a, 0x84, 0x8b, 0x20, 0x08, 0x96,
	0x20, 0x79, 0x02, 0x27, 0x22, 0x10, 0x92, 0xc2, 0xc8, 0x49, 0x93, 0x22, 0x46, 0x97, 0x13, 0xe5,
	0xac, 0x64, 0x69, 0x18, 0x8d, 0xcc, 0xda, 0xe5, 0x3e, 0xc1, 0x3e, 0xd6, 0x96, 0x2e, 0xb7, 0x5c,
	0xec, 0x17, 0x59, 0xa4, 0xc1, 0xb7, 0xc5, 0xdd, 0xfc, 0xff, 0xfc, 0x7c, 0xcd, 0x77, 0xe8, 0x47,
	0x99, 0xa2, 0x2b, 0x53, 0x5c, 0x28, 0x58, 0x15, 0x51, 0xa0, 0xb1, 0xc8, 0x17, 0x31, 0x96, 0x5a,
	0x61, 0x58, 0x35, 0x41, 0x16, 0x98, 0x6b, 0x47, 0xaa, 0x42, 0x17, 0xcc, 0xd9, 0x60, 0x12, 0x02,
	0x04, 0x59, 0x86, 0x41, 0x1e, 0x81, 0x73, 0x1c, 0x42, 0x1c, 0x15, 0x4b, 0x69, 0xda, 0x0c, 0xe2,
	0x04, 0x94, 0x23, 0x53, 0x1c, 0xdf, 0xb5, 0xa9, 0x98, 0xa5, 0xe8, 0x1f, 0xd1, 0xde, 0x19, 0x79,
	0x56, 0x83, 0xd9, 0x7b, 0xda, 0x5e, 0x61, 0xcc, 0x89, 0x45, 0xec, 0xae, 0x5f, 0x3f, 0xd9, 0x88,
	0x76, 0xb3, 0x20, 0x84, 0x8c, 0xbf, 0xb2, 0x88, 0x3d, 0xf4, 0x4d, 0x60, 0x0e, 0x65, 0x58, 0x96,
	0x15, 0xa8, 0x79, 0x15, 0xde, 0x40, 0xa4, 0x7f, 0xc0, 0xfa, 0xbb, 0xc7, 0xdb, 0xcd, 0xe4, 0xca,
	0x4f, 0xcd, 0x95, 0x18, 0xf3, 0x8e, 0xe1, 0x4a, 0xc3, 0xc5, 0x72, 0x36, 0x9d, 0xf2, 0xae, 0x45,
	0xec, 0x81, 0x6f, 0x02, 0xfb, 0x44, 0x47, 0x91, 0xca, 0xe6, 0x98, 0xe4, 0xa0, 0xbe, 0x82, 0xd2,
	0xf8, 0x0f, 0xa3, 0x40, 0x03, 0xef, 0x35, 0xe4, 0xab, 0x7f, 0x8c, 0xd3, 0x7e, 0x1c, 0xe8, 0xe0,
	0xb7, 0xff, 0x93, 0xf7, 0x9b, 0xd9, 0x21, 0xb2, 0x31, 0x7d, 0x53, 0x3f, 0xbf, 0x61, 0x06, 0x73,
	0xdc, 0x00, 0x1f, 0x58, 0xc4, 0xee, 0xf8, 0x17, 0x1d, 0x13, 0x94, 0xd6, 0xd9, 0xc3, 0x04, 0x4a,
	0xcd, 0x87, 0x0d, 0xe0, 0xac, 0x61, 0x1f, 0xe8, 0xbb, 0x53, 0xfa, 0xb5, 0x96, 0xc0, 0xa9, 0x45,
	0xec, 0xb7, 0xfe, 0x8b, 0xb6, 0xde, 0x9d, 0x9c, 0x35, 0xbb, 0xd7, 0x66, 0x77, 0xd9, 0x7e, 0xf9,
	0xfb, 0xb0, 0x13, 0x64, 0xbb, 0x13, 0xe4, 0x69, 0x27, 0xc8, 0xfd, 0x5e, 0xb4, 0xb6, 0x7b, 0xd1,
	0x7a, 0xdc, 0x8b, 0xd6, 0x1f, 0x2f, 0x41, 0xfd, 0xbf, 0x0a, 0x9d, 0xa8, 0x58, 0xba, 0xc6, 0xec,
	0xe4, 0xa0, 0xd6, 0x3d, 0x53, 0x3b, 0x39, 0xb9, 0x9d, 0x18, 0xb9, 0xee, 0x6d, 0x7d, 0x32, 0xae,
	0x5e, 0x4b, 0x28, 0xc3, 0x5e, 0x73, 0x1b, 0x9f, 0x9f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xed, 0x64,
	0x00, 0x15, 0x4a, 0x02, 0x00, 0x00,
}

func (m *PkiRevocationDistributionPoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PkiRevocationDistributionPoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PkiRevocationDistributionPoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RevocationType != 0 {
		i = encodeVarintPkiRevocationDistributionPoint(dAtA, i, uint64(m.RevocationType))
		i--
		dAtA[i] = 0x58
	}
	if m.DataDigestType != 0 {
		i = encodeVarintPkiRevocationDistributionPoint(dAtA, i, uint64(m.DataDigestType))
		i--
		dAtA[i] = 0x50
	}
	if len(m.DataDigest) > 0 {
		i -= len(m.DataDigest)
		copy(dAtA[i:], m.DataDigest)
		i = encodeVarintPkiRevocationDistributionPoint(dAtA, i, uint64(len(m.DataDigest)))
		i--
		dAtA[i] = 0x4a
	}
	if m.DataFileSize != 0 {
		i = encodeVarintPkiRevocationDistributionPoint(dAtA, i, uint64(m.DataFileSize))
		i--
		dAtA[i] = 0x40
	}
	if len(m.DataURL) > 0 {
		i -= len(m.DataURL)
		copy(dAtA[i:], m.DataURL)
		i = encodeVarintPkiRevocationDistributionPoint(dAtA, i, uint64(len(m.DataURL)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.CrlSignerCertificate) > 0 {
		i -= len(m.CrlSignerCertificate)
		copy(dAtA[i:], m.CrlSignerCertificate)
		i = encodeVarintPkiRevocationDistributionPoint(dAtA, i, uint64(len(m.CrlSignerCertificate)))
		i--
		dAtA[i] = 0x32
	}
	if m.IsPAA {
		i--
		if m.IsPAA {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Pid != 0 {
		i = encodeVarintPkiRevocationDistributionPoint(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x20
	}
	if len(m.IssuerSubjectKeyID) > 0 {
		i -= len(m.IssuerSubjectKeyID)
		copy(dAtA[i:], m.IssuerSubjectKeyID)
		i = encodeVarintPkiRevocationDistributionPoint(dAtA, i, uint64(len(m.IssuerSubjectKeyID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintPkiRevocationDistributionPoint(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x12
	}
	if m.Vid != 0 {
		i = encodeVarintPkiRevocationDistributionPoint(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPkiRevocationDistributionPoint(dAtA []byte, offset int, v uint64) int {
	offset -= sovPkiRevocationDistributionPoint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PkiRevocationDistributionPoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vid != 0 {
		n += 1 + sovPkiRevocationDistributionPoint(uint64(m.Vid))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovPkiRevocationDistributionPoint(uint64(l))
	}
	l = len(m.IssuerSubjectKeyID)
	if l > 0 {
		n += 1 + l + sovPkiRevocationDistributionPoint(uint64(l))
	}
	if m.Pid != 0 {
		n += 1 + sovPkiRevocationDistributionPoint(uint64(m.Pid))
	}
	if m.IsPAA {
		n += 2
	}
	l = len(m.CrlSignerCertificate)
	if l > 0 {
		n += 1 + l + sovPkiRevocationDistributionPoint(uint64(l))
	}
	l = len(m.DataURL)
	if l > 0 {
		n += 1 + l + sovPkiRevocationDistributionPoint(uint64(l))
	}
	if m.DataFileSize != 0 {
		n += 1 + sovPkiRevocationDistributionPoint(uint64(m.DataFileSize))
	}
	l = len(m.DataDigest)
	if l > 0 {
		n += 1 + l + sovPkiRevocationDistributionPoint(uint64(l))
	}
	if m.DataDigestType != 0 {
		n += 1 + sovPkiRevocationDistributionPoint(uint64(m.DataDigestType))
	}
	if m.RevocationType != 0 {
		n += 1 + sovPkiRevocationDistributionPoint(uint64(m.RevocationType))
	}
	return n
}

func sovPkiRevocationDistributionPoint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPkiRevocationDistributionPoint(x uint64) (n int) {
	return sovPkiRevocationDistributionPoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PkiRevocationDistributionPoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPkiRevocationDistributionPoint
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
			return fmt.Errorf("proto: PkiRevocationDistributionPoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PkiRevocationDistributionPoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
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
				return ErrInvalidLengthPkiRevocationDistributionPoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerSubjectKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
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
				return ErrInvalidLengthPkiRevocationDistributionPoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerSubjectKeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPAA", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
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
			m.IsPAA = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrlSignerCertificate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
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
				return ErrInvalidLengthPkiRevocationDistributionPoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CrlSignerCertificate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
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
				return ErrInvalidLengthPkiRevocationDistributionPoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataFileSize", wireType)
			}
			m.DataFileSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataFileSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataDigest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
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
				return ErrInvalidLengthPkiRevocationDistributionPoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataDigest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataDigestType", wireType)
			}
			m.DataDigestType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataDigestType |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevocationType", wireType)
			}
			m.RevocationType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RevocationType |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPkiRevocationDistributionPoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPoint
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
func skipPkiRevocationDistributionPoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPkiRevocationDistributionPoint
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
					return 0, ErrIntOverflowPkiRevocationDistributionPoint
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
					return 0, ErrIntOverflowPkiRevocationDistributionPoint
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
				return 0, ErrInvalidLengthPkiRevocationDistributionPoint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPkiRevocationDistributionPoint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPkiRevocationDistributionPoint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPkiRevocationDistributionPoint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPkiRevocationDistributionPoint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPkiRevocationDistributionPoint = fmt.Errorf("proto: unexpected end of group")
)
