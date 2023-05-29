// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/concentrated-liquidity/pool.proto

// This is a legacy package that requires additional migration logic
// in order to use the correct package. Decision made to use legacy package path
// until clear steps for migration logic and the unknowns for state breaking are
// investigated for changing proto package.

package model

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type Pool struct {
	// pool's address holding all liquidity tokens.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// address holding the incentives liquidity.
	IncentivesAddress string `protobuf:"bytes,2,opt,name=incentives_address,json=incentivesAddress,proto3" json:"incentives_address,omitempty" yaml:"incentives_address"`
	// address holding spread rewards from swaps.
	SpreadRewardsAddress string `protobuf:"bytes,3,opt,name=spread_rewards_address,json=spreadRewardsAddress,proto3" json:"spread_rewards_address,omitempty" yaml:"spread_rewards_address"`
	Id                   uint64 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	// Amount of total liquidity
	CurrentTickLiquidity github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=current_tick_liquidity,json=currentTickLiquidity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_tick_liquidity" yaml:"current_tick_liquidity"`
	Token0               string                                 `protobuf:"bytes,6,opt,name=token0,proto3" json:"token0,omitempty"`
	Token1               string                                 `protobuf:"bytes,7,opt,name=token1,proto3" json:"token1,omitempty"`
	CurrentSqrtPrice     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=current_sqrt_price,json=currentSqrtPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"current_sqrt_price" yaml:"spot_price"`
	CurrentTick          int64                                  `protobuf:"varint,9,opt,name=current_tick,json=currentTick,proto3" json:"current_tick,omitempty" yaml:"current_tick"`
	// tick_spacing must be one of the authorized_tick_spacing values set in the
	// concentrated-liquidity parameters
	TickSpacing        uint64 `protobuf:"varint,10,opt,name=tick_spacing,json=tickSpacing,proto3" json:"tick_spacing,omitempty" yaml:"tick_spacing"`
	ExponentAtPriceOne int64  `protobuf:"varint,11,opt,name=exponent_at_price_one,json=exponentAtPriceOne,proto3" json:"exponent_at_price_one,omitempty" yaml:"exponent_at_price_one"`
	// spread_factor is the ratio that is charged on the amount of token in.
	SpreadFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=spread_factor,json=spreadFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"spread_factor" yaml:"spread_factor"`
	// last_liquidity_update is the last time either the pool liquidity or the
	// active tick changed
	LastLiquidityUpdate time.Time `protobuf:"bytes,13,opt,name=last_liquidity_update,json=lastLiquidityUpdate,proto3,stdtime" json:"last_liquidity_update" yaml:"last_liquidity_update"`
}

func (m *Pool) Reset()      { *m = Pool{} }
func (*Pool) ProtoMessage() {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_3526ea5373d96c9a, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Pool)(nil), "osmosis.concentratedliquidity.v1beta1.Pool")
}

func init() {
	proto.RegisterFile("osmosis/concentrated-liquidity/pool.proto", fileDescriptor_3526ea5373d96c9a)
}

var fileDescriptor_3526ea5373d96c9a = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0xfb, 0xfb, 0x75, 0xd2, 0x56, 0x5f, 0xa7, 0x69, 0x71, 0x2b, 0x1a, 0x07, 0x4b, 0xa0,
	0x20, 0x11, 0x9b, 0x80, 0xd8, 0x74, 0xd7, 0x80, 0x2a, 0x21, 0x55, 0x6a, 0x35, 0x2d, 0x42, 0x42,
	0x48, 0xd6, 0xc4, 0x9e, 0x86, 0x51, 0x1c, 0x8f, 0xe3, 0x99, 0x94, 0xf6, 0x01, 0x90, 0x58, 0x76,
	0xc9, 0xb2, 0x0f, 0xc1, 0x43, 0x54, 0xac, 0xba, 0x44, 0x48, 0x18, 0xd4, 0xbe, 0x41, 0x9e, 0x00,
	0x79, 0x66, 0x9c, 0x18, 0x35, 0x2c, 0xba, 0x4a, 0xce, 0xb9, 0xe7, 0x9e, 0xb9, 0xf7, 0x7a, 0xee,
	0x80, 0xc7, 0x8c, 0xf7, 0x18, 0xa7, 0xdc, 0xf5, 0x59, 0xe4, 0x93, 0x48, 0x24, 0x58, 0x90, 0xa0,
	0x11, 0xd2, 0xfe, 0x80, 0x06, 0x54, 0x9c, 0xb9, 0x31, 0x63, 0xa1, 0x13, 0x27, 0x4c, 0x30, 0xf8,
	0x50, 0x4b, 0x9d, 0xa2, 0x74, 0xa4, 0x74, 0x4e, 0x9a, 0x6d, 0x22, 0x70, 0x73, 0x73, 0xc3, 0x97,
	0x3a, 0x4f, 0x26, 0xb9, 0x0a, 0x28, 0x87, 0xcd, 0x4a, 0x87, 0x75, 0x98, 0xe2, 0xb3, 0x7f, 0x9a,
	0xb5, 0x3a, 0x8c, 0x75, 0x42, 0xe2, 0x4a, 0xd4, 0x1e, 0x1c, 0xbb, 0x82, 0xf6, 0x08, 0x17, 0xb8,
	0x17, 0x2b, 0x81, 0xfd, 0x73, 0x1e, 0xcc, 0x1c, 0x30, 0x16, 0xc2, 0x27, 0x60, 0x1e, 0x07, 0x41,
	0x42, 0x38, 0x37, 0x8d, 0x9a, 0x51, 0x5f, 0x68, 0xc1, 0x61, 0x6a, 0x2d, 0x9f, 0xe1, 0x5e, 0xb8,
	0x6d, 0xeb, 0x80, 0x8d, 0x72, 0x09, 0xdc, 0x03, 0x90, 0xca, 0x42, 0xe9, 0x09, 0xe1, 0x5e, 0x9e,
	0x38, 0x25, 0x13, 0xb7, 0x86, 0xa9, 0xb5, 0xa1, 0x12, 0x6f, 0x6b, 0x6c, 0xb4, 0x32, 0x26, 0x77,
	0xb4, 0xdb, 0x5b, 0xb0, 0xce, 0xe3, 0x84, 0xe0, 0xc0, 0x4b, 0xc8, 0x47, 0x9c, 0x04, 0x63, 0xc7,
	0x69, 0xe9, 0xf8, 0x60, 0x98, 0x5a, 0x5b, 0xca, 0x71, 0xb2, 0xce, 0x46, 0x15, 0x15, 0x40, 0x8a,
	0xcf, 0x8d, 0x97, 0xc1, 0x14, 0x0d, 0xcc, 0x99, 0x9a, 0x51, 0x9f, 0x41, 0x53, 0x34, 0x80, 0x9f,
	0x0c, 0xb0, 0xee, 0x0f, 0x92, 0x84, 0x44, 0xc2, 0x13, 0xd4, 0xef, 0x7a, 0xa3, 0x11, 0x9b, 0xb3,
	0xf2, 0xa4, 0xfd, 0xcb, 0xd4, 0x2a, 0xfd, 0x48, 0xad, 0x47, 0x1d, 0x2a, 0x3e, 0x0c, 0xda, 0x8e,
	0xcf, 0x7a, 0x7a, 0xcc, 0xfa, 0xa7, 0xc1, 0x83, 0xae, 0x2b, 0xce, 0x62, 0xc2, 0x9d, 0x57, 0xc4,
	0x1f, 0xd7, 0x35, 0xd9, 0xd5, 0x46, 0x15, 0x1d, 0x38, 0xa2, 0x7e, 0x77, 0x2f, 0xa7, 0xe1, 0x3a,
	0x98, 0x13, 0xac, 0x4b, 0xa2, 0xa7, 0xe6, 0x5c, 0x76, 0x2c, 0xd2, 0x68, 0xc4, 0x37, 0xcd, 0xf9,
	0x02, 0xdf, 0x84, 0x7d, 0x00, 0xf3, 0x03, 0x78, 0x3f, 0x11, 0x5e, 0x9c, 0x50, 0x9f, 0x98, 0xff,
	0xc9, 0x92, 0x5f, 0xde, 0xb9, 0xe4, 0x95, 0x7c, 0x94, 0x4c, 0x3b, 0xd9, 0xe8, 0x7f, 0x6d, 0x7f,
	0xd8, 0x4f, 0xc4, 0x41, 0x46, 0xc1, 0x6d, 0xb0, 0x58, 0xec, 0xc9, 0x5c, 0xa8, 0x19, 0xf5, 0xe9,
	0xd6, 0xbd, 0x61, 0x6a, 0xad, 0xde, 0xee, 0xd8, 0x46, 0xe5, 0x42, 0x9f, 0x59, 0xae, 0x9c, 0x03,
	0x8f, 0xb1, 0x4f, 0xa3, 0x8e, 0x09, 0xb2, 0x0f, 0x50, 0xcc, 0x2d, 0x46, 0x6d, 0x54, 0xce, 0xe0,
	0xa1, 0x42, 0xf0, 0x10, 0xac, 0x91, 0xd3, 0x98, 0x45, 0x99, 0x35, 0xd6, 0xf5, 0x79, 0x2c, 0x22,
	0x66, 0x59, 0x16, 0x50, 0x1b, 0xa6, 0xd6, 0x7d, 0x65, 0x32, 0x51, 0x66, 0x23, 0x98, 0xf3, 0x3b,
	0xaa, 0x93, 0xfd, 0x88, 0xc0, 0x2e, 0x58, 0xd2, 0x17, 0xe7, 0x18, 0xfb, 0x82, 0x25, 0xe6, 0xa2,
	0x1c, 0xdd, 0xee, 0x9d, 0x47, 0x57, 0xf9, 0xeb, 0x16, 0x2a, 0x33, 0x1b, 0x2d, 0x2a, 0xbc, 0x2b,
	0x21, 0x3c, 0x05, 0x6b, 0x21, 0xe6, 0x62, 0x7c, 0x0b, 0xbc, 0x41, 0x1c, 0x60, 0x41, 0xcc, 0xa5,
	0x9a, 0x51, 0x2f, 0x3f, 0xdb, 0x74, 0xd4, 0x4e, 0x3a, 0xf9, 0x4e, 0x3a, 0x47, 0xf9, 0x4e, 0xb6,
	0xea, 0x59, 0x41, 0xe3, 0x0e, 0x27, 0xda, 0xd8, 0xe7, 0xbf, 0x2c, 0x03, 0xad, 0x66, 0xb1, 0xd1,
	0x85, 0x7a, 0x23, 0x23, 0xdb, 0x2b, 0x9f, 0x2f, 0xac, 0xd2, 0x97, 0x0b, 0xab, 0xf4, 0xed, 0x6b,
	0x63, 0x36, 0xdb, 0xea, 0xd7, 0xad, 0xf7, 0x97, 0xd7, 0x55, 0xe3, 0xea, 0xba, 0x6a, 0xfc, 0xbe,
	0xae, 0x1a, 0xe7, 0x37, 0xd5, 0xd2, 0xd5, 0x4d, 0xb5, 0xf4, 0xfd, 0xa6, 0x5a, 0x7a, 0xd7, 0x2a,
	0x34, 0xad, 0x5f, 0x9f, 0x46, 0x88, 0xdb, 0x3c, 0x07, 0xee, 0x49, 0xf3, 0x85, 0x7b, 0xfa, 0xaf,
	0xb7, 0xab, 0xc7, 0x02, 0x12, 0xb6, 0xe7, 0x64, 0x0f, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xd8, 0xba, 0x26, 0xce, 0xea, 0x04, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastLiquidityUpdate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastLiquidityUpdate):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPool(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x6a
	{
		size := m.SpreadFactor.Size()
		i -= size
		if _, err := m.SpreadFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.ExponentAtPriceOne != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.ExponentAtPriceOne))
		i--
		dAtA[i] = 0x58
	}
	if m.TickSpacing != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.TickSpacing))
		i--
		dAtA[i] = 0x50
	}
	if m.CurrentTick != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.CurrentTick))
		i--
		dAtA[i] = 0x48
	}
	{
		size := m.CurrentSqrtPrice.Size()
		i -= size
		if _, err := m.CurrentSqrtPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.Token1) > 0 {
		i -= len(m.Token1)
		copy(dAtA[i:], m.Token1)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Token1)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Token0) > 0 {
		i -= len(m.Token0)
		copy(dAtA[i:], m.Token0)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Token0)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.CurrentTickLiquidity.Size()
		i -= size
		if _, err := m.CurrentTickLiquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Id != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SpreadRewardsAddress) > 0 {
		i -= len(m.SpreadRewardsAddress)
		copy(dAtA[i:], m.SpreadRewardsAddress)
		i = encodeVarintPool(dAtA, i, uint64(len(m.SpreadRewardsAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.IncentivesAddress) > 0 {
		i -= len(m.IncentivesAddress)
		copy(dAtA[i:], m.IncentivesAddress)
		i = encodeVarintPool(dAtA, i, uint64(len(m.IncentivesAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.IncentivesAddress)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.SpreadRewardsAddress)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPool(uint64(m.Id))
	}
	l = m.CurrentTickLiquidity.Size()
	n += 1 + l + sovPool(uint64(l))
	l = len(m.Token0)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.Token1)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.CurrentSqrtPrice.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.CurrentTick != 0 {
		n += 1 + sovPool(uint64(m.CurrentTick))
	}
	if m.TickSpacing != 0 {
		n += 1 + sovPool(uint64(m.TickSpacing))
	}
	if m.ExponentAtPriceOne != 0 {
		n += 1 + sovPool(uint64(m.ExponentAtPriceOne))
	}
	l = m.SpreadFactor.Size()
	n += 1 + l + sovPool(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastLiquidityUpdate)
	n += 1 + l + sovPool(uint64(l))
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentivesAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncentivesAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpreadRewardsAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpreadRewardsAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTickLiquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentTickLiquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSqrtPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentSqrtPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTick", wireType)
			}
			m.CurrentTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentTick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickSpacing", wireType)
			}
			m.TickSpacing = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickSpacing |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExponentAtPriceOne", wireType)
			}
			m.ExponentAtPriceOne = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExponentAtPriceOne |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpreadFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SpreadFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastLiquidityUpdate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastLiquidityUpdate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)