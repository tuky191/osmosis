// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/accum/v1beta1/accum.proto

package accum

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// AccumulatorContent is the state-entry for the global accumulator.
// It contains the name of the global accumulator and the total value of
// shares belonging to it from all positions.
type AccumulatorContent struct {
	AccumValue  github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=accum_value,json=accumValue,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"accum_value"`
	TotalShares github_com_cosmos_cosmos_sdk_types.Dec      `protobuf:"bytes,2,opt,name=total_shares,json=totalShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"total_shares"`
}

func (m *AccumulatorContent) Reset()         { *m = AccumulatorContent{} }
func (m *AccumulatorContent) String() string { return proto.CompactTextString(m) }
func (*AccumulatorContent) ProtoMessage()    {}
func (*AccumulatorContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4866f7c74a169dc2, []int{0}
}
func (m *AccumulatorContent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccumulatorContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccumulatorContent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccumulatorContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccumulatorContent.Merge(m, src)
}
func (m *AccumulatorContent) XXX_Size() int {
	return m.Size()
}
func (m *AccumulatorContent) XXX_DiscardUnknown() {
	xxx_messageInfo_AccumulatorContent.DiscardUnknown(m)
}

var xxx_messageInfo_AccumulatorContent proto.InternalMessageInfo

func (m *AccumulatorContent) GetAccumValue() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.AccumValue
	}
	return nil
}

type Options struct {
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_4866f7c74a169dc2, []int{1}
}
func (m *Options) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Options.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return m.Size()
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

// Record corresponds to an individual position value belonging to the
// global accumulator.
type Record struct {
	// num_shares is the number of shares belonging to the position associated
	// with this record.
	NumShares github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=num_shares,json=numShares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"num_shares"`
	// accum_value_per_share is the subset of coins per shar of the global
	// accumulator value that allows to infer how much a position is entitled to
	// per share that it owns.
	//
	// In the default case with no intervals, this value equals to the global
	// accumulator value at the time of the position creation, the last update or
	// reward claim.
	//
	// In the interval case such as concentrated liquidity, this value equals to
	// the global growth of rewards inside the interval during one of: the time of
	// the position creation, the last update or reward claim. Note, that
	// immediately prior to claiming or updating rewards, this value must be
	// updated to "the growth inside at the time of last update + the growth
	// outside at the time of the current block". This is so that the claiming
	// logic can subtract this updated value from the global accumulator value to
	// get the growth inside the interval from the time of last update up until
	// the current block time.
	AccumValuePerShare github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=accum_value_per_share,json=accumValuePerShare,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"accum_value_per_share"`
	// unclaimed_rewards_total is the total amount of unclaimed rewards that the
	// position is entitled to. This value is updated whenever shares are added or
	// removed from an existing position. We also expose API for manually updating
	// this value for some custom use cases such as merging pre-existing positions
	// into a single one.
	UnclaimedRewardsTotal github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,3,rep,name=unclaimed_rewards_total,json=unclaimedRewardsTotal,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"unclaimed_rewards_total"`
	Options               *Options                                    `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_4866f7c74a169dc2, []int{2}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetAccumValuePerShare() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.AccumValuePerShare
	}
	return nil
}

func (m *Record) GetUnclaimedRewardsTotal() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.UnclaimedRewardsTotal
	}
	return nil
}

func (m *Record) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*AccumulatorContent)(nil), "osmosis.accum.v1beta1.AccumulatorContent")
	proto.RegisterType((*Options)(nil), "osmosis.accum.v1beta1.Options")
	proto.RegisterType((*Record)(nil), "osmosis.accum.v1beta1.Record")
}

func init() { proto.RegisterFile("osmosis/accum/v1beta1/accum.proto", fileDescriptor_4866f7c74a169dc2) }

var fileDescriptor_4866f7c74a169dc2 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x3f, 0x8f, 0xd3, 0x30,
	0x1c, 0x86, 0xe3, 0x2b, 0xba, 0x53, 0x1d, 0x26, 0x8b, 0x8a, 0xe8, 0x84, 0xdc, 0xd2, 0x01, 0x55,
	0x42, 0xe7, 0xe8, 0xae, 0x0b, 0x2b, 0x2d, 0x0b, 0x03, 0x02, 0x02, 0x62, 0x60, 0x89, 0x1c, 0xc7,
	0x6a, 0x23, 0x12, 0x3b, 0xf2, 0x9f, 0x22, 0x76, 0x16, 0x36, 0x3e, 0x04, 0x13, 0x9f, 0xa4, 0x63,
	0x47, 0x04, 0x52, 0x41, 0xed, 0x17, 0x41, 0xb1, 0x93, 0xd2, 0x81, 0x01, 0x21, 0x75, 0x72, 0x9c,
	0xbc, 0x79, 0x9e, 0x9f, 0xde, 0x38, 0xf0, 0xbe, 0xd4, 0x95, 0xd4, 0x85, 0x8e, 0x29, 0x63, 0xb6,
	0x8a, 0x57, 0xd7, 0x19, 0x37, 0xf4, 0xda, 0xef, 0x48, 0xad, 0xa4, 0x91, 0x68, 0xd0, 0x46, 0x88,
	0xbf, 0xd9, 0x46, 0x2e, 0xef, 0x2c, 0xe4, 0x42, 0xba, 0x44, 0xdc, 0x5c, 0xf9, 0xf0, 0x25, 0x66,
	0x2e, 0x1d, 0x67, 0x54, 0xf3, 0x03, 0x8d, 0xc9, 0x42, 0xf8, 0xe7, 0xe3, 0x1f, 0x00, 0xa2, 0xc7,
	0x0d, 0xc7, 0x96, 0xd4, 0x48, 0x35, 0x97, 0xc2, 0x70, 0x61, 0x90, 0x82, 0xa1, 0xa3, 0xa7, 0x2b,
	0x5a, 0x5a, 0x1e, 0x81, 0x51, 0x6f, 0x12, 0xde, 0xdc, 0x23, 0x1e, 0x46, 0x1a, 0x58, 0xe7, 0x25,
	0x4f, 0x38, 0x9b, 0xcb, 0x42, 0xcc, 0xa6, 0xeb, 0xed, 0x30, 0xf8, 0xfa, 0x73, 0xf8, 0x70, 0x51,
	0x98, 0xa5, 0xcd, 0x08, 0x93, 0x55, 0xdc, 0xca, 0xfd, 0x72, 0xa5, 0xf3, 0x77, 0xb1, 0xf9, 0x50,
	0x73, 0xdd, 0xbd, 0xa3, 0x13, 0xe8, 0x2c, 0x6f, 0x1a, 0x09, 0x7a, 0x09, 0x6f, 0x1b, 0x69, 0x68,
	0x99, 0xea, 0x25, 0x55, 0x5c, 0x47, 0x67, 0x23, 0x30, 0xe9, 0xcf, 0x48, 0x83, 0xfd, 0xbe, 0x1d,
	0x3e, 0xf8, 0x37, 0x6c, 0x12, 0x3a, 0xc6, 0x2b, 0x87, 0x18, 0xf7, 0xe1, 0xc5, 0xf3, 0xda, 0x14,
	0x52, 0xe8, 0xf1, 0x97, 0x1e, 0x3c, 0x4f, 0x38, 0x93, 0x2a, 0x47, 0xcf, 0x20, 0x14, 0xb6, 0xea,
	0x34, 0xe0, 0xbf, 0x34, 0x7d, 0x61, 0x2b, 0x2f, 0x41, 0x1f, 0x01, 0x1c, 0x1c, 0x95, 0x95, 0xd6,
	0x5c, 0x79, 0x76, 0x74, 0x76, 0xaa, 0xda, 0xd0, 0x9f, 0xda, 0x5e, 0x70, 0xe5, 0xe6, 0x40, 0x9f,
	0x00, 0xbc, 0x6b, 0x05, 0x2b, 0x69, 0x51, 0xf1, 0x3c, 0x55, 0xfc, 0x3d, 0x55, 0xb9, 0x4e, 0x5d,
	0x1b, 0x51, 0xef, 0x54, 0x83, 0x0c, 0x0e, 0xc6, 0xc4, 0x0b, 0x5f, 0x37, 0x3e, 0xf4, 0x08, 0x5e,
	0x48, 0xdf, 0x7b, 0x74, 0x6b, 0x04, 0x26, 0xe1, 0x0d, 0x26, 0x7f, 0x3d, 0xb4, 0xa4, 0xfd, 0x3a,
	0x49, 0x17, 0x9f, 0x3d, 0x5d, 0xef, 0x30, 0xd8, 0xec, 0x30, 0xf8, 0xb5, 0xc3, 0xe0, 0xf3, 0x1e,
	0x07, 0x9b, 0x3d, 0x0e, 0xbe, 0xed, 0x71, 0xf0, 0x36, 0x3e, 0x9a, 0xab, 0x85, 0x5d, 0x95, 0x34,
	0xd3, 0xdd, 0xc6, 0xad, 0xd6, 0x14, 0x65, 0xfb, 0xef, 0x64, 0xe7, 0xee, 0x84, 0x4f, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xd7, 0xe4, 0xc9, 0x60, 0x53, 0x03, 0x00, 0x00,
}

func (m *AccumulatorContent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccumulatorContent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccumulatorContent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalShares.Size()
		i -= size
		if _, err := m.TotalShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccum(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.AccumValue) > 0 {
		for iNdEx := len(m.AccumValue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccumValue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccum(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Options) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Options) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Options) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Options != nil {
		{
			size, err := m.Options.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccum(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.UnclaimedRewardsTotal) > 0 {
		for iNdEx := len(m.UnclaimedRewardsTotal) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnclaimedRewardsTotal[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccum(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AccumValuePerShare) > 0 {
		for iNdEx := len(m.AccumValuePerShare) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccumValuePerShare[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccum(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.NumShares.Size()
		i -= size
		if _, err := m.NumShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAccum(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintAccum(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccum(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccumulatorContent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AccumValue) > 0 {
		for _, e := range m.AccumValue {
			l = e.Size()
			n += 1 + l + sovAccum(uint64(l))
		}
	}
	l = m.TotalShares.Size()
	n += 1 + l + sovAccum(uint64(l))
	return n
}

func (m *Options) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NumShares.Size()
	n += 1 + l + sovAccum(uint64(l))
	if len(m.AccumValuePerShare) > 0 {
		for _, e := range m.AccumValuePerShare {
			l = e.Size()
			n += 1 + l + sovAccum(uint64(l))
		}
	}
	if len(m.UnclaimedRewardsTotal) > 0 {
		for _, e := range m.UnclaimedRewardsTotal {
			l = e.Size()
			n += 1 + l + sovAccum(uint64(l))
		}
	}
	if m.Options != nil {
		l = m.Options.Size()
		n += 1 + l + sovAccum(uint64(l))
	}
	return n
}

func sovAccum(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccum(x uint64) (n int) {
	return sovAccum(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccumulatorContent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccum
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
			return fmt.Errorf("proto: AccumulatorContent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccumulatorContent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccumValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccum
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
				return ErrInvalidLengthAccum
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccumValue = append(m.AccumValue, types.DecCoin{})
			if err := m.AccumValue[len(m.AccumValue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccum
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
				return ErrInvalidLengthAccum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccum
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
func (m *Options) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccum
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
			return fmt.Errorf("proto: Options: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Options: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAccum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccum
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
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccum
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
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumShares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccum
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
				return ErrInvalidLengthAccum
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NumShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccumValuePerShare", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccum
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
				return ErrInvalidLengthAccum
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccumValuePerShare = append(m.AccumValuePerShare, types.DecCoin{})
			if err := m.AccumValuePerShare[len(m.AccumValuePerShare)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnclaimedRewardsTotal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccum
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
				return ErrInvalidLengthAccum
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnclaimedRewardsTotal = append(m.UnclaimedRewardsTotal, types.DecCoin{})
			if err := m.UnclaimedRewardsTotal[len(m.UnclaimedRewardsTotal)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccum
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
				return ErrInvalidLengthAccum
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccum
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = &Options{}
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccum(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccum
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
func skipAccum(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccum
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
					return 0, ErrIntOverflowAccum
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
					return 0, ErrIntOverflowAccum
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
				return 0, ErrInvalidLengthAccum
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccum
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccum
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccum        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccum          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccum = fmt.Errorf("proto: unexpected end of group")
)