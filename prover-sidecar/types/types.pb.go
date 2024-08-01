// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	types "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TODO: FIGURE OUT HOW TO GET THESE IMPORTED INSTEAD, SOMETHING WEIRD ABOUT INCORRECT GO TYPE OR SOMETHING
type SignedPacketCommitmentsClaim struct {
	AttestatorId           []byte                 `protobuf:"bytes,1,opt,name=attestator_id,json=attestatorId,proto3" json:"attestator_id,omitempty"`
	PacketCommitmentsClaim PacketCommitmentsClaim `protobuf:"bytes,2,opt,name=packet_commitments_claim,json=packetCommitmentsClaim,proto3" json:"packet_commitments_claim"`
	Signature              []byte                 `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedPacketCommitmentsClaim) Reset()         { *m = SignedPacketCommitmentsClaim{} }
func (m *SignedPacketCommitmentsClaim) String() string { return proto.CompactTextString(m) }
func (*SignedPacketCommitmentsClaim) ProtoMessage()    {}
func (*SignedPacketCommitmentsClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_16b92cdb257a8883, []int{0}
}
func (m *SignedPacketCommitmentsClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignedPacketCommitmentsClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignedPacketCommitmentsClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignedPacketCommitmentsClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedPacketCommitmentsClaim.Merge(m, src)
}
func (m *SignedPacketCommitmentsClaim) XXX_Size() int {
	return m.Size()
}
func (m *SignedPacketCommitmentsClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedPacketCommitmentsClaim.DiscardUnknown(m)
}

var xxx_messageInfo_SignedPacketCommitmentsClaim proto.InternalMessageInfo

func (m *SignedPacketCommitmentsClaim) GetAttestatorId() []byte {
	if m != nil {
		return m.AttestatorId
	}
	return nil
}

func (m *SignedPacketCommitmentsClaim) GetPacketCommitmentsClaim() PacketCommitmentsClaim {
	if m != nil {
		return m.PacketCommitmentsClaim
	}
	return PacketCommitmentsClaim{}
}

func (m *SignedPacketCommitmentsClaim) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type PacketCommitmentsClaim struct {
	Height            types.Height `protobuf:"bytes,1,opt,name=height,proto3" json:"height"`
	Timestamp         time.Time    `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	PacketCommitments [][]byte     `protobuf:"bytes,3,rep,name=packet_commitments,json=packetCommitments,proto3" json:"packet_commitments,omitempty"`
}

func (m *PacketCommitmentsClaim) Reset()         { *m = PacketCommitmentsClaim{} }
func (m *PacketCommitmentsClaim) String() string { return proto.CompactTextString(m) }
func (*PacketCommitmentsClaim) ProtoMessage()    {}
func (*PacketCommitmentsClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_16b92cdb257a8883, []int{1}
}
func (m *PacketCommitmentsClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketCommitmentsClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketCommitmentsClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketCommitmentsClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketCommitmentsClaim.Merge(m, src)
}
func (m *PacketCommitmentsClaim) XXX_Size() int {
	return m.Size()
}
func (m *PacketCommitmentsClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketCommitmentsClaim.DiscardUnknown(m)
}

var xxx_messageInfo_PacketCommitmentsClaim proto.InternalMessageInfo

func (m *PacketCommitmentsClaim) GetHeight() types.Height {
	if m != nil {
		return m.Height
	}
	return types.Height{}
}

func (m *PacketCommitmentsClaim) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *PacketCommitmentsClaim) GetPacketCommitments() [][]byte {
	if m != nil {
		return m.PacketCommitments
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedPacketCommitmentsClaim)(nil), "SignedPacketCommitmentsClaim")
	proto.RegisterType((*PacketCommitmentsClaim)(nil), "PacketCommitmentsClaim")
}

func init() { proto.RegisterFile("server/types.proto", fileDescriptor_16b92cdb257a8883) }

var fileDescriptor_16b92cdb257a8883 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0xcf, 0x1c, 0xaa, 0xa8, 0x7b, 0x0c, 0x44, 0xa8, 0x44, 0xa7, 0x2a, 0x77, 0x2a, 0xcb,
	0x2d, 0xb5, 0xd5, 0xb2, 0x30, 0xa7, 0x0c, 0xb0, 0xa1, 0x80, 0x84, 0xc4, 0xc0, 0xc9, 0x71, 0x1e,
	0xbe, 0x07, 0x71, 0x1c, 0xd9, 0x2f, 0x91, 0xf8, 0x16, 0xfd, 0x3c, 0xec, 0x48, 0x1d, 0x3b, 0x32,
	0x01, 0xba, 0xfb, 0x22, 0x28, 0x69, 0x8e, 0x0c, 0x2d, 0xdb, 0xf3, 0xd3, 0xff, 0xff, 0xde, 0xdf,
	0x3f, 0x9b, 0x47, 0x01, 0x7c, 0x0b, 0x5e, 0xd2, 0xb7, 0x1a, 0x82, 0xa8, 0xbd, 0x23, 0x37, 0x5f,
	0x60, 0xae, 0xa5, 0x76, 0x1e, 0xa4, 0x2e, 0x11, 0x2a, 0x92, 0xed, 0xf9, 0x50, 0xed, 0x05, 0xc6,
	0x39, 0x53, 0x82, 0xec, 0x4f, 0x79, 0xf3, 0x59, 0x12, 0x5a, 0x08, 0xa4, 0x6c, 0x3d, 0x08, 0x9e,
	0x1a, 0x67, 0x5c, 0x5f, 0xca, 0xae, 0xba, 0xed, 0x9e, 0x7e, 0x67, 0xfc, 0xe4, 0x1d, 0x9a, 0x0a,
	0x8a, 0xb7, 0x4a, 0x7f, 0x05, 0xba, 0x74, 0xd6, 0x22, 0x59, 0xa8, 0x28, 0x5c, 0x96, 0x0a, 0x6d,
	0xf4, 0x9c, 0x3f, 0x56, 0x44, 0xdd, 0x24, 0x72, 0x7e, 0x8d, 0x45, 0xcc, 0x96, 0x6c, 0x35, 0xcb,
	0x66, 0x63, 0xf3, 0x4d, 0x11, 0x7d, 0xe0, 0x71, 0xdd, 0xdb, 0xd7, 0x7a, 0xf4, 0xaf, 0x75, 0x37,
	0x20, 0x7e, 0xb0, 0x64, 0xab, 0xa3, 0x8b, 0x67, 0xe2, 0xfe, 0xf9, 0xe9, 0xc3, 0xeb, 0x5f, 0x8b,
	0x49, 0x76, 0x5c, 0xdf, 0xbf, 0xfd, 0x84, 0x1f, 0x06, 0x34, 0x95, 0xa2, 0xc6, 0x43, 0x3c, 0xed,
	0x37, 0x8f, 0x8d, 0xd3, 0x1f, 0x8c, 0x1f, 0xff, 0x27, 0xf6, 0x4b, 0x7e, 0xb0, 0x01, 0x34, 0x1b,
	0xea, 0xf3, 0x1e, 0x5d, 0xcc, 0x05, 0xe6, 0x5a, 0x74, 0x00, 0xc5, 0x80, 0xad, 0x3d, 0x17, 0xaf,
	0x7b, 0xc5, 0x10, 0x61, 0xd0, 0x47, 0x29, 0x3f, 0xfc, 0x87, 0x6e, 0x08, 0x3f, 0x17, 0xb7, 0x70,
	0xc5, 0x1e, 0xae, 0x78, 0xbf, 0x57, 0xa4, 0x8f, 0x3a, 0xf3, 0xd5, 0xef, 0x05, 0xcb, 0x46, 0x5b,
	0x74, 0xc6, 0xa3, 0xbb, 0x3c, 0xe2, 0xe9, 0x72, 0xba, 0x9a, 0x65, 0x4f, 0xee, 0x5c, 0x35, 0xfd,
	0x74, 0xbd, 0x4d, 0xd8, 0xcd, 0x36, 0x61, 0x7f, 0xb6, 0x09, 0xbb, 0xda, 0x25, 0x93, 0x9b, 0x5d,
	0x32, 0xf9, 0xb9, 0x4b, 0x26, 0x1f, 0x5f, 0x19, 0xa4, 0x4d, 0x93, 0x0b, 0xed, 0xac, 0x34, 0x5f,
	0xc0, 0xdb, 0xa6, 0x2a, 0x8c, 0xf2, 0x2a, 0x57, 0xb2, 0x86, 0x10, 0xd0, 0x62, 0x20, 0xd4, 0x67,
	0xad, 0x2a, 0xb1, 0x50, 0x84, 0xae, 0xea, 0xde, 0xbf, 0x05, 0x1f, 0xb0, 0x00, 0xad, 0x86, 0x2f,
	0x94, 0x1f, 0xf4, 0xb9, 0x5f, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xff, 0x72, 0x8c, 0x23, 0x59,
	0x02, 0x00, 0x00,
}

func (m *SignedPacketCommitmentsClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignedPacketCommitmentsClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignedPacketCommitmentsClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.PacketCommitmentsClaim.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.AttestatorId) > 0 {
		i -= len(m.AttestatorId)
		copy(dAtA[i:], m.AttestatorId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AttestatorId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PacketCommitmentsClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketCommitmentsClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketCommitmentsClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PacketCommitments) > 0 {
		for iNdEx := len(m.PacketCommitments) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PacketCommitments[iNdEx])
			copy(dAtA[i:], m.PacketCommitments[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.PacketCommitments[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Height.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignedPacketCommitmentsClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AttestatorId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.PacketCommitmentsClaim.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *PacketCommitmentsClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Height.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovTypes(uint64(l))
	if len(m.PacketCommitments) > 0 {
		for _, b := range m.PacketCommitments {
			l = len(b)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignedPacketCommitmentsClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SignedPacketCommitmentsClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignedPacketCommitmentsClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttestatorId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttestatorId = append(m.AttestatorId[:0], dAtA[iNdEx:postIndex]...)
			if m.AttestatorId == nil {
				m.AttestatorId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketCommitmentsClaim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PacketCommitmentsClaim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *PacketCommitmentsClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PacketCommitmentsClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketCommitmentsClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Height.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketCommitments", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PacketCommitments = append(m.PacketCommitments, make([]byte, postIndex-iNdEx))
			copy(m.PacketCommitments[len(m.PacketCommitments)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
