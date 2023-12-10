// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/twap/v1beta1/twap_record.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// A TWAP record should be indexed in state by pool_id, (asset pair), timestamp
// The asset pair assets should be lexicographically sorted.
// Technically (pool_id, asset_0_denom, asset_1_denom, height) do not need to
// appear in the struct however we view this as the wrong performance tradeoff
// given SDK today. Would rather we optimize for readability and correctness,
// than an optimal state storage format. The system bottleneck is elsewhere for
// now.
type TwapRecord struct {
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// Lexicographically smaller denom of the pair
	Asset0Denom string `protobuf:"bytes,2,opt,name=asset0_denom,json=asset0Denom,proto3" json:"asset0_denom,omitempty"`
	// Lexicographically larger denom of the pair
	Asset1Denom string `protobuf:"bytes,3,opt,name=asset1_denom,json=asset1Denom,proto3" json:"asset1_denom,omitempty"`
	// height this record corresponds to, for debugging purposes
	Height int64 `protobuf:"varint,4,opt,name=height,proto3" json:"record_height" yaml:"record_height"`
	// This field should only exist until we have a global registry in the state
	// machine, mapping prior block heights within {TIME RANGE} to times.
	Time time.Time `protobuf:"bytes,5,opt,name=time,proto3,stdtime" json:"time" yaml:"record_time"`
	// We store the last spot prices in the struct, so that we can interpolate
	// accumulator values for times between when accumulator records are stored.
	P0LastSpotPrice             cosmossdk_io_math.LegacyDec `protobuf:"bytes,6,opt,name=p0_last_spot_price,json=p0LastSpotPrice,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"p0_last_spot_price"`
	P1LastSpotPrice             cosmossdk_io_math.LegacyDec `protobuf:"bytes,7,opt,name=p1_last_spot_price,json=p1LastSpotPrice,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"p1_last_spot_price"`
	P0ArithmeticTwapAccumulator cosmossdk_io_math.LegacyDec `protobuf:"bytes,8,opt,name=p0_arithmetic_twap_accumulator,json=p0ArithmeticTwapAccumulator,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"p0_arithmetic_twap_accumulator"`
	P1ArithmeticTwapAccumulator cosmossdk_io_math.LegacyDec `protobuf:"bytes,9,opt,name=p1_arithmetic_twap_accumulator,json=p1ArithmeticTwapAccumulator,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"p1_arithmetic_twap_accumulator"`
	GeometricTwapAccumulator    cosmossdk_io_math.LegacyDec `protobuf:"bytes,10,opt,name=geometric_twap_accumulator,json=geometricTwapAccumulator,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"geometric_twap_accumulator"`
	// This field contains the time in which the last spot price error occurred.
	// It is used to alert the caller if they are getting a potentially erroneous
	// TWAP, due to an unforeseen underlying error.
	LastErrorTime time.Time `protobuf:"bytes,11,opt,name=last_error_time,json=lastErrorTime,proto3,stdtime" json:"last_error_time" yaml:"last_error_time"`
}

func (m *TwapRecord) Reset()         { *m = TwapRecord{} }
func (m *TwapRecord) String() string { return proto.CompactTextString(m) }
func (*TwapRecord) ProtoMessage()    {}
func (*TwapRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbf5c78678e601aa, []int{0}
}
func (m *TwapRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TwapRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TwapRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TwapRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TwapRecord.Merge(m, src)
}
func (m *TwapRecord) XXX_Size() int {
	return m.Size()
}
func (m *TwapRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_TwapRecord.DiscardUnknown(m)
}

var xxx_messageInfo_TwapRecord proto.InternalMessageInfo

func (m *TwapRecord) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *TwapRecord) GetAsset0Denom() string {
	if m != nil {
		return m.Asset0Denom
	}
	return ""
}

func (m *TwapRecord) GetAsset1Denom() string {
	if m != nil {
		return m.Asset1Denom
	}
	return ""
}

func (m *TwapRecord) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TwapRecord) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *TwapRecord) GetLastErrorTime() time.Time {
	if m != nil {
		return m.LastErrorTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*TwapRecord)(nil), "osmosis.twap.v1beta1.TwapRecord")
}

func init() {
	proto.RegisterFile("osmosis/twap/v1beta1/twap_record.proto", fileDescriptor_dbf5c78678e601aa)
}

var fileDescriptor_dbf5c78678e601aa = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x63, 0x1a, 0x52, 0xba, 0xa1, 0xaa, 0x64, 0x45, 0x60, 0x52, 0xc9, 0x0e, 0x41, 0x42,
	0xe1, 0x80, 0x3f, 0xca, 0x8d, 0x5b, 0xa2, 0x72, 0x00, 0x55, 0xa8, 0x32, 0x3d, 0x71, 0xb1, 0xd6,
	0xce, 0xd6, 0xb6, 0xb0, 0x33, 0xab, 0xdd, 0x4d, 0x4b, 0xde, 0xa2, 0xaf, 0xc4, 0xad, 0xc7, 0x1e,
	0x11, 0x07, 0x83, 0x92, 0x1b, 0xc7, 0x3e, 0x01, 0xda, 0x5d, 0x27, 0x90, 0xf0, 0xd5, 0xde, 0x3c,
	0x33, 0xff, 0xf9, 0xfd, 0x77, 0xbc, 0xa3, 0x45, 0x4f, 0x81, 0x97, 0xc0, 0x73, 0xee, 0x89, 0x73,
	0x4c, 0xbd, 0xb3, 0x20, 0x26, 0x02, 0x07, 0x2a, 0x88, 0x18, 0x49, 0x80, 0x8d, 0x5d, 0xca, 0x40,
	0x80, 0xd9, 0xa9, 0x75, 0xae, 0x2c, 0xb9, 0xb5, 0xae, 0xdb, 0x49, 0x21, 0x05, 0x25, 0xf0, 0xe4,
	0x97, 0xd6, 0x76, 0x1f, 0xa5, 0x00, 0x69, 0x41, 0x3c, 0x15, 0xc5, 0xd3, 0x53, 0x0f, 0x4f, 0x66,
	0xcb, 0x52, 0xa2, 0x38, 0x91, 0xee, 0xd1, 0x41, 0x5d, 0xb2, 0x75, 0xe4, 0xc5, 0x98, 0x93, 0xd5,
	0x41, 0x12, 0xc8, 0x27, 0x75, 0xdd, 0xd9, 0xa4, 0x8a, 0xbc, 0x24, 0x5c, 0xe0, 0x92, 0x6a, 0x41,
	0xff, 0x53, 0x0b, 0xa1, 0x93, 0x73, 0x4c, 0x43, 0x75, 0x6e, 0xf3, 0x21, 0xda, 0xa6, 0x00, 0x45,
	0x94, 0x8f, 0x2d, 0xa3, 0x67, 0x0c, 0x9a, 0x61, 0x4b, 0x86, 0xaf, 0xc7, 0xe6, 0x63, 0x74, 0x1f,
	0x73, 0x4e, 0x84, 0x1f, 0x8d, 0xc9, 0x04, 0x4a, 0xeb, 0x4e, 0xcf, 0x18, 0xec, 0x84, 0x6d, 0x9d,
	0x3b, 0x94, 0xa9, 0x95, 0x24, 0xa8, 0x25, 0x5b, 0xbf, 0x48, 0x02, 0x2d, 0x19, 0xa2, 0x56, 0x46,
	0xf2, 0x34, 0x13, 0x56, 0xb3, 0x67, 0x0c, 0xb6, 0x46, 0xcf, 0xbe, 0x57, 0xce, 0xae, 0xfe, 0x65,
	0x91, 0x2e, 0x5c, 0x57, 0x4e, 0x67, 0x86, 0xcb, 0xe2, 0x65, 0x7f, 0x2d, 0xdd, 0x0f, 0xeb, 0x46,
	0xf3, 0x2d, 0x6a, 0xca, 0x19, 0xac, 0xbb, 0x3d, 0x63, 0xd0, 0x3e, 0xe8, 0xba, 0x7a, 0x40, 0x77,
	0x39, 0xa0, 0x7b, 0xb2, 0x1c, 0x70, 0x64, 0x5f, 0x56, 0x4e, 0xe3, 0xba, 0x72, 0xcc, 0x35, 0x9e,
	0x6c, 0xee, 0x5f, 0x7c, 0x75, 0x8c, 0x50, 0x71, 0xcc, 0x63, 0x64, 0x52, 0x3f, 0x2a, 0x30, 0x17,
	0x11, 0xa7, 0x20, 0x22, 0xca, 0xf2, 0x84, 0x58, 0x2d, 0x79, 0xf6, 0xd1, 0x13, 0x49, 0xf8, 0x52,
	0x39, 0xfb, 0xfa, 0x2f, 0xf3, 0xf1, 0x07, 0x37, 0x07, 0xaf, 0xc4, 0x22, 0x73, 0x8f, 0x48, 0x8a,
	0x93, 0xd9, 0x21, 0x49, 0xc2, 0x3d, 0xea, 0x1f, 0x61, 0x2e, 0xde, 0x51, 0x10, 0xc7, 0xb2, 0x57,
	0x11, 0x83, 0xdf, 0x88, 0xdb, 0xb7, 0x21, 0x06, 0xeb, 0xc4, 0x0c, 0xd9, 0xd4, 0x8f, 0x30, 0xcb,
	0x45, 0x56, 0x12, 0x91, 0x27, 0x91, 0x5a, 0x35, 0x9c, 0x24, 0xd3, 0x72, 0x5a, 0x60, 0x01, 0xcc,
	0xba, 0x77, 0x73, 0xfa, 0x3e, 0xf5, 0x87, 0x2b, 0x92, 0xbc, 0xfa, 0xe1, 0x4f, 0x8e, 0x72, 0x0a,
	0xfe, 0xe9, 0xb4, 0x73, 0x1b, 0xa7, 0xe0, 0xef, 0x4e, 0x18, 0x75, 0x53, 0x02, 0x25, 0x11, 0xec,
	0x4f, 0x2e, 0xe8, 0xe6, 0x2e, 0xd6, 0x0a, 0xb3, 0x69, 0x71, 0x8a, 0xf6, 0xd4, 0x2d, 0x10, 0xc6,
	0x80, 0xa9, 0x8b, 0xb7, 0xda, 0xff, 0xdd, 0x9a, 0x7e, 0xbd, 0x35, 0x0f, 0xf4, 0xd6, 0x6c, 0x00,
	0xf4, 0xe6, 0xec, 0xca, 0xec, 0x2b, 0x99, 0x94, 0x7d, 0xa3, 0x37, 0x97, 0x73, 0xdb, 0xb8, 0x9a,
	0xdb, 0xc6, 0xb7, 0xb9, 0x6d, 0x5c, 0x2c, 0xec, 0xc6, 0xd5, 0xc2, 0x6e, 0x7c, 0x5e, 0xd8, 0x8d,
	0xf7, 0x7e, 0x9a, 0x8b, 0x6c, 0x1a, 0xbb, 0x09, 0x94, 0x5e, 0xfd, 0x16, 0x3c, 0x2f, 0x70, 0xcc,
	0x97, 0x81, 0x77, 0x76, 0x10, 0x78, 0x1f, 0xf5, 0x33, 0x22, 0x66, 0x94, 0xf0, 0xb8, 0xa5, 0x8e,
	0xf4, 0xe2, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xdb, 0x71, 0xb6, 0x63, 0x04, 0x00, 0x00,
}

func (m *TwapRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TwapRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TwapRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.LastErrorTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastErrorTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTwapRecord(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x5a
	{
		size := m.GeometricTwapAccumulator.Size()
		i -= size
		if _, err := m.GeometricTwapAccumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTwapRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.P1ArithmeticTwapAccumulator.Size()
		i -= size
		if _, err := m.P1ArithmeticTwapAccumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTwapRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.P0ArithmeticTwapAccumulator.Size()
		i -= size
		if _, err := m.P0ArithmeticTwapAccumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTwapRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.P1LastSpotPrice.Size()
		i -= size
		if _, err := m.P1LastSpotPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTwapRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.P0LastSpotPrice.Size()
		i -= size
		if _, err := m.P0LastSpotPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTwapRecord(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTwapRecord(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if m.Height != 0 {
		i = encodeVarintTwapRecord(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Asset1Denom) > 0 {
		i -= len(m.Asset1Denom)
		copy(dAtA[i:], m.Asset1Denom)
		i = encodeVarintTwapRecord(dAtA, i, uint64(len(m.Asset1Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Asset0Denom) > 0 {
		i -= len(m.Asset0Denom)
		copy(dAtA[i:], m.Asset0Denom)
		i = encodeVarintTwapRecord(dAtA, i, uint64(len(m.Asset0Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintTwapRecord(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTwapRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovTwapRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TwapRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovTwapRecord(uint64(m.PoolId))
	}
	l = len(m.Asset0Denom)
	if l > 0 {
		n += 1 + l + sovTwapRecord(uint64(l))
	}
	l = len(m.Asset1Denom)
	if l > 0 {
		n += 1 + l + sovTwapRecord(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTwapRecord(uint64(m.Height))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTwapRecord(uint64(l))
	l = m.P0LastSpotPrice.Size()
	n += 1 + l + sovTwapRecord(uint64(l))
	l = m.P1LastSpotPrice.Size()
	n += 1 + l + sovTwapRecord(uint64(l))
	l = m.P0ArithmeticTwapAccumulator.Size()
	n += 1 + l + sovTwapRecord(uint64(l))
	l = m.P1ArithmeticTwapAccumulator.Size()
	n += 1 + l + sovTwapRecord(uint64(l))
	l = m.GeometricTwapAccumulator.Size()
	n += 1 + l + sovTwapRecord(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastErrorTime)
	n += 1 + l + sovTwapRecord(uint64(l))
	return n
}

func sovTwapRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTwapRecord(x uint64) (n int) {
	return sovTwapRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TwapRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTwapRecord
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
			return fmt.Errorf("proto: TwapRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TwapRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset0Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset0Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset1Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset1Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P0LastSpotPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P0LastSpotPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P1LastSpotPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P1LastSpotPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P0ArithmeticTwapAccumulator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P0ArithmeticTwapAccumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P1ArithmeticTwapAccumulator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P1ArithmeticTwapAccumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GeometricTwapAccumulator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GeometricTwapAccumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastErrorTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwapRecord
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
				return ErrInvalidLengthTwapRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTwapRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.LastErrorTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTwapRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTwapRecord
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
func skipTwapRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTwapRecord
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
					return 0, ErrIntOverflowTwapRecord
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
					return 0, ErrIntOverflowTwapRecord
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
				return 0, ErrInvalidLengthTwapRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTwapRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTwapRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTwapRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTwapRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTwapRecord = fmt.Errorf("proto: unexpected end of group")
)
